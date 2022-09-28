package database

import (
	"fmt"
	"github.com/Akkadius/spire/internal/encryption"
	"github.com/Akkadius/spire/internal/env"
	"github.com/Akkadius/spire/internal/http/request"
	"github.com/Akkadius/spire/internal/models"
	"github.com/labstack/echo/v4"
	gocache "github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"strings"
	"time"
)

// database resolver is used to dynamically resolve to external database instances that users
// have defined in the connections database
// when the users request comes in; a middleware function will lookup their database info
// from the memory pool of connections and re-use it if found, re-establish if not found or
// return an error if credentials no longer work
type DatabaseResolver struct {
	connections           *Connections                 // local connections
	remoteDatabases       map[string]map[uint]*gorm.DB // remote databases only used when Spire resolves connections defined by users
	logger                *logrus.Logger
	crypt                 *encryption.Encrypter
	cache                 *gocache.Cache
	contentConnectionName string
}

func NewDatabaseResolver(
	connections *Connections,
	logger *logrus.Logger,
	crypt *encryption.Encrypter,
	cache *gocache.Cache,
) *DatabaseResolver {
	i := &DatabaseResolver{
		connections:           connections,
		remoteDatabases:       map[string]map[uint]*gorm.DB{},
		logger:                logger,
		crypt:                 crypt,
		cache:                 cache,
		contentConnectionName: os.Getenv("MYSQL_EQEMU_CONTENT_DB_CONNECTION_NAME"),
	}

	// initialize ahead of time
	i.remoteDatabases["default"] = map[uint]*gorm.DB{}
	i.remoteDatabases["eqemu_content"] = map[uint]*gorm.DB{}

	return i
}

func (d *DatabaseResolver) Get(model models.Modelable, c echo.Context) *gorm.DB {
	user := request.GetUser(c)
	if user.ID > 0 {
		return d.ResolveUserEqemuConnection(model, user)
	}

	return d.connections.EqemuDb()
}

func (d *DatabaseResolver) GetSpireDb() *gorm.DB {
	return d.connections.SpireDb()
}

func (d *DatabaseResolver) GetEqemuDb() *gorm.DB {
	return d.connections.EqemuDb()
}

func (d *DatabaseResolver) GetEncKey(userId uint) string {
	return fmt.Sprintf("%v-%v", env.Get("APP_KEY", ""), userId)
}

func (d *DatabaseResolver) ResolveUserEqemuConnection(model models.Modelable, user models.User) *gorm.DB {

	// use default otherwise key off of another connection type
	connectionType := "default"
	if model.Connection() == d.contentConnectionName {
		connectionType = model.Connection()
	}

	// init nested map if not set
	_, ok := d.remoteDatabases[connectionType]
	if !ok {
		d.remoteDatabases[connectionType] = map[uint]*gorm.DB{}
	}

	// If we don't have a user
	if user.ID == 0 {
		return d.connections.EqemuDb()
	}

	// fetch connection id from memory first if exists
	connectionId := uint(0)
	connectionIdKey := fmt.Sprintf("active-connection-%v-%v", user.ID, connectionType)
	cachedConn, found := d.cache.Get(connectionIdKey)
	if found {
		connectionId = cachedConn.(uint)

		// If existing connection exists, return it
		if _, ok := d.remoteDatabases[connectionType][connectionId]; ok {
			//fmt.Println("Returning cached lookup")
			db, err := d.remoteDatabases[connectionType][connectionId].DB()
			if err != nil {
				d.logger.Printf("Debug: MySQL ping err [%v]", err)
			}

			err = db.Ping()
			if err != nil {
				d.logger.Printf("Debug: MySQL ping err [%v]", err)
			}

			return d.remoteDatabases[connectionType][connectionId]
		}
	}

	// get servers from database
	conn := d.GetUserConnection(user)

	// if we don't have an active connection
	// this will then fallback to the locally defined eqemu instance
	if conn.ID == 0 {

		// set default local to connection pool for default fallback
		d.remoteDatabases[connectionType][conn.ID] = d.connections.EqemuDb()

		// add connection id to memory
		d.cache.Set(connectionIdKey, conn.ServerDatabaseConnection.ID, 10*time.Minute)

		return d.connections.EqemuDb()
	}

	// check if we found a connection from the database
	if conn.ID > 0 {

		// add connection id to memory
		d.cache.Set(connectionIdKey, conn.ServerDatabaseConnection.ID, 10*time.Minute)

		// If existing connection exists, return it
		if _, ok := d.remoteDatabases[connectionType][conn.ServerDatabaseConnection.ID]; ok {
			db, err := d.remoteDatabases[connectionType][conn.ServerDatabaseConnection.ID].DB()
			if err != nil {
				d.logger.Printf("Debug: MySQL ping err [%v]", err)
			}

			err = db.Ping()
			if err != nil {
				d.logger.Printf("Debug: MySQL ping err [%v]", err)
			}

			return d.remoteDatabases[connectionType][conn.ServerDatabaseConnection.ID]
		}
	}

	// eqemu server default
	dbUsername := conn.ServerDatabaseConnection.DbUsername
	dbPassword := d.crypt.Decrypt(conn.ServerDatabaseConnection.DbPassword, d.GetEncKey(user.ID))
	dbHost := conn.ServerDatabaseConnection.DbHost
	dbPort := conn.ServerDatabaseConnection.DbPort
	dbName := conn.ServerDatabaseConnection.DbName

	// content connection
	if model.Connection() == d.contentConnectionName && conn.ServerDatabaseConnection.ContentDbUsername != "" {
		dbUsername = conn.ServerDatabaseConnection.ContentDbUsername
		dbPassword = d.crypt.Decrypt(conn.ServerDatabaseConnection.ContentDbPassword, d.GetEncKey(user.ID))
		dbHost = conn.ServerDatabaseConnection.ContentDbHost
		dbPort = conn.ServerDatabaseConnection.ContentDbPort
		dbName = conn.ServerDatabaseConnection.ContentDbName
	}

	// init nested map if not set
	_, ok = d.remoteDatabases[connectionType]
	if !ok {
		d.remoteDatabases[connectionType] = map[uint]*gorm.DB{}
	}

	// create new connection since we don't have one
	dsn := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local&timeout=1s",
		dbUsername,
		dbPassword,
		dbHost,
		dbPort,
		dbName,
	)

	logMode := logger.Silent
	if env.GetBool("MYSQL_QUERY_LOGGING", "false") {
		logMode = logger.Info
	}

	db, err := gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{
			SkipDefaultTransaction:                   true,
			DisableForeignKeyConstraintWhenMigrating: true,
			DisableAutomaticPing:                     false,
			Logger:                                   logger.Default.LogMode(logMode),
		},
	)

	if err != nil {
		d.logger.Printf("[database_resolver] MySQL conn err [%v]", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		d.logger.Printf("[database_resolver] MySQL fetch err [%v]", err)
	}

	sqlDB.SetConnMaxLifetime(time.Minute * 3)
	sqlDB.SetMaxIdleConns(env.GetInt("MYSQL_MAX_IDLE_CONNECTIONS", "10"))
	sqlDB.SetMaxOpenConns(env.GetInt("MYSQL_MAX_OPEN_CONNECTIONS", "150"))

	// cache instance pointer to memory
	d.remoteDatabases[connectionType][conn.ServerDatabaseConnection.ID] = db

	return db
}

func (d *DatabaseResolver) handleWheres(query *gorm.DB, filter string) *gorm.DB {
	// parse where field = value
	wheres := strings.Split(filter, equalDelimiter)
	if len(wheres) > 1 {
		query = query.Where(fmt.Sprintf("%v = ?", wheres[0]), wheres[1])
	}

	// parse where field != value
	wheres = strings.Split(filter, notEqualDelimiter)
	if len(wheres) > 1 {
		query = query.Where(fmt.Sprintf("%v != ?", wheres[0]), wheres[1])
	}

	// parse where bitwise field & value
	wheres = strings.Split(filter, bitwiseAnd)
	if len(wheres) > 1 {
		query = query.Where(fmt.Sprintf("%v & ?", wheres[0]), wheres[1])
	}

	// parse where field like '%value%'
	wheres = strings.Split(filter, notLikeDelimiter)
	if len(wheres) > 1 {
		query = query.Where(fmt.Sprintf("%v not like ?", wheres[0]), fmt.Sprintf("%%%v%%", wheres[1]))
	}

	// parse where field like '%value%'
	wheres = strings.Split(filter, likeDelimiter)
	if len(wheres) > 1 {
		query = query.Where(fmt.Sprintf("%v like ?", wheres[0]), fmt.Sprintf("%%%v%%", wheres[1]))
	}

	// parse where [value > x]
	wheres = strings.Split(filter, greaterThan)
	if len(wheres) > 1 {
		query = query.Where(fmt.Sprintf("%v > ?", wheres[0]), wheres[1])
	}

	// parse where [value < x]
	wheres = strings.Split(filter, lesserThan)
	if len(wheres) > 1 {
		query = query.Where(fmt.Sprintf("%v < ?", wheres[0]), wheres[1])
	}

	// parse where [value >= x]
	wheres = strings.Split(filter, greaterThanEqual)
	if len(wheres) > 1 {
		query = query.Where(fmt.Sprintf("%v >= ?", wheres[0]), wheres[1])
	}

	// parse where [value <= x]
	wheres = strings.Split(filter, lesserThanEqual)
	if len(wheres) > 1 {
		query = query.Where(fmt.Sprintf("%v <= ?", wheres[0]), wheres[1])
	}

	// parse where [value = x]
	wheres = strings.Split(filter, equal)
	if len(wheres) > 1 {
		query = query.Where(fmt.Sprintf("%v = ?", wheres[0]), wheres[1])
	}

	return query
}

func (d *DatabaseResolver) handleOrWheres(query *gorm.DB, filter string) *gorm.DB {
	// parse where field = value
	wheres := strings.Split(filter, equalDelimiter)
	if len(wheres) > 1 {
		query = query.Or(fmt.Sprintf("%v = ?", wheres[0]), wheres[1])
	}

	// parse where field != value
	wheres = strings.Split(filter, notEqualDelimiter)
	if len(wheres) > 1 {
		query = query.Or(fmt.Sprintf("%v != ?", wheres[0]), wheres[1])
	}

	// parse where bitwise field & value
	wheres = strings.Split(filter, bitwiseAnd)
	if len(wheres) > 1 {
		query = query.Or(fmt.Sprintf("%v & ?", wheres[0]), wheres[1])
	}

	// parse where field like '%value%'
	wheres = strings.Split(filter, notLikeDelimiter)
	if len(wheres) > 1 {
		query = query.Or(fmt.Sprintf("%v not like ?", wheres[0]), fmt.Sprintf("%%%v%%", wheres[1]))
	}

	// parse where field like '%value%'
	wheres = strings.Split(filter, likeDelimiter)
	if len(wheres) > 1 {
		query = query.Or(fmt.Sprintf("%v like ?", wheres[0]), fmt.Sprintf("%%%v%%", wheres[1]))
	}

	// parse where [value > x]
	wheres = strings.Split(filter, greaterThan)
	if len(wheres) > 1 {
		query = query.Or(fmt.Sprintf("%v > ?", wheres[0]), wheres[1])
	}

	// parse where [value < x]
	wheres = strings.Split(filter, lesserThan)
	if len(wheres) > 1 {
		query = query.Or(fmt.Sprintf("%v < ?", wheres[0]), wheres[1])
	}

	// parse where [value >= x]
	wheres = strings.Split(filter, greaterThanEqual)
	if len(wheres) > 1 {
		query = query.Or(fmt.Sprintf("%v >= ?", wheres[0]), wheres[1])
	}

	// parse where [value <= x]
	wheres = strings.Split(filter, lesserThanEqual)
	if len(wheres) > 1 {
		query = query.Or(fmt.Sprintf("%v <= ?", wheres[0]), wheres[1])
	}

	// parse where [value = x]
	wheres = strings.Split(filter, equal)
	if len(wheres) > 1 {
		query = query.Or(fmt.Sprintf("%v = ?", wheres[0]), wheres[1])
	}

	return query
}

func (d *DatabaseResolver) GetUserConnection(user models.User) models.UserServerDatabaseConnection {
	var conn models.UserServerDatabaseConnection
	relationships := models.UserServerDatabaseConnection{}.Relationships()
	query := d.GetSpireDb().Model(&models.UserServerDatabaseConnection{})
	for _, relationship := range relationships {
		query = query.Preload(relationship)
	}

	query.Where("user_id = ? and active = 1", user.ID).First(&conn)

	return conn
}
