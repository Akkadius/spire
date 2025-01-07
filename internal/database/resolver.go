package database

import (
	"fmt"
	"github.com/Akkadius/spire/internal/encryption"
	"github.com/Akkadius/spire/internal/env"
	"github.com/Akkadius/spire/internal/http/request"
	"github.com/Akkadius/spire/internal/logger"
	"github.com/Akkadius/spire/internal/models"
	"github.com/labstack/echo/v4"
	gocache "github.com/patrickmn/go-cache"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"strings"
	"sync"
	"time"
)

// database resolver is used to dynamically resolve to external database instances that users
// have defined in the connections database
// when the users request comes in; a middleware function will lookup their database info
// from the memory pool of connections and re-use it if found, re-establish if not found or
// return an error if credentials no longer work

type Resolver struct {
	connections     *Connections                 // local connections
	remoteDatabases map[string]map[uint]*gorm.DB // remote databases only used when Spire resolves connections defined by users
	logger          *logger.AppLogger
	crypt           *encryption.Encrypter
	cache           *gocache.Cache
}

const (
	connectionTypeDefault = "default"       // default connection type
	connectionTypeContent = "eqemu_content" // content connection type
	connectionTypeLogs    = "eqemu_logs"    // logs connection type
)

// connection types
var connectionTypes = []string{
	connectionTypeDefault,
	connectionTypeContent,
	connectionTypeLogs,
}

func NewResolver(
	connections *Connections,
	logger *logger.AppLogger,
	crypt *encryption.Encrypter,
	cache *gocache.Cache,
) *Resolver {
	i := &Resolver{
		connections:     connections,
		remoteDatabases: map[string]map[uint]*gorm.DB{},
		logger:          logger,
		crypt:           crypt,
		cache:           cache,
	}

	go i.connectionKeepAlive()

	// initialize
	for c := range connectionTypes {
		i.remoteDatabases[connectionTypes[c]] = map[uint]*gorm.DB{}
	}

	return i
}

func (d *Resolver) Get(model models.Modelable, c echo.Context) *gorm.DB {
	user := request.GetUser(c)
	if user.ID > 0 {
		return d.ResolveUserEqemuConnection(model, user)
	}

	return d.connections.EqemuDb()
}

func (d *Resolver) GetSpireDb() *gorm.DB {
	return d.connections.SpireDb()
}

func (d *Resolver) GetEqemuDb() *gorm.DB {
	return d.connections.EqemuDb()
}

func (d *Resolver) GetEncKey(userId uint) string {
	return fmt.Sprintf("%v-%v", d.crypt.GetEncryptionKey(), userId)
}

// connectionCreationMutex is used to lock the creation of database connections
var connCreationMutex = sync.Mutex{}

func (d *Resolver) ResolveUserEqemuConnection(model models.Modelable, user models.User) *gorm.DB {
	connCreationMutex.Lock()
	defer connCreationMutex.Unlock()

	// use default otherwise key off of another connection type
	connectionType := "default"
	if model.Connection() == connectionTypeContent {
		connectionType = connectionTypeContent
	} else if model.Connection() == connectionTypeLogs {
		connectionType = connectionTypeLogs
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

	// key for the users database connection identifier
	connectionIdKey := fmt.Sprintf("active-connection-%v", user.ID)

	// this holds the key for the database instance itself
	// keyed off of the connection type (default, content)
	connectionKey := fmt.Sprintf("active-connection-%v-%v", user.ID, connectionType)

	// find cached connection
	cachedConn, found := d.cache.Get(connectionKey)

	// if we didn't find a cached connection, lock the user mutex for
	// database connection creation incase user fires off multiple requests
	if !found {
		//fmt.Println("after mutex firing " + connectionKey + " for user " + fmt.Sprintf("%v", user.ID))
		cachedConn, found = d.cache.Get(connectionKey)
		if found {
			connectionId = cachedConn.(uint)
		}
	}

	// if we found a cached connection, return it
	if found {
		connectionId = cachedConn.(uint)

		// If existing connection exists, return it
		if _, ok := d.remoteDatabases[connectionType][connectionId]; ok {
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
		d.cache.Set(connectionKey, conn.ServerDatabaseConnection.ID, 10*time.Minute)
		d.cache.Set(connectionIdKey, conn.ServerDatabaseConnection.ID, 10*time.Minute)

		return d.connections.EqemuDb()
	}

	// check if we found a connection from the database
	if conn.ID > 0 {

		// add connection id to memory
		d.cache.Set(connectionKey, conn.ServerDatabaseConnection.ID, 10*time.Minute)
		d.cache.Set(connectionIdKey, conn.ServerDatabaseConnection.ID, 10*time.Minute)

		// If existing connection exists, return it
		if _, ok := d.remoteDatabases[connectionType][conn.ServerDatabaseConnection.ID]; ok {
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
	if model.Connection() == connectionTypeContent && conn.ServerDatabaseConnection.ContentDbUsername != "" {
		dbUsername = conn.ServerDatabaseConnection.ContentDbUsername
		dbPassword = d.crypt.Decrypt(conn.ServerDatabaseConnection.ContentDbPassword, d.GetEncKey(user.ID))
		dbHost = conn.ServerDatabaseConnection.ContentDbHost
		dbPort = conn.ServerDatabaseConnection.ContentDbPort
		dbName = conn.ServerDatabaseConnection.ContentDbName
	} else if model.Connection() == connectionTypeLogs && conn.ServerDatabaseConnection.LogsDbUsername != "" {
		dbUsername = conn.ServerDatabaseConnection.LogsDbUsername
		dbPassword = d.crypt.Decrypt(conn.ServerDatabaseConnection.LogsDbPassword, d.GetEncKey(user.ID))
		dbHost = conn.ServerDatabaseConnection.LogsDbHost
		dbPort = conn.ServerDatabaseConnection.LogsDbPort
		dbName = conn.ServerDatabaseConnection.LogsDbName
	}

	// init nested map if not set
	_, ok = d.remoteDatabases[connectionType]
	if !ok {
		d.remoteDatabases[connectionType] = map[uint]*gorm.DB{}
	}

	// create new connection since we don't have one
	dsn := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local&timeout=3s",
		dbUsername,
		dbPassword,
		dbHost,
		dbPort,
		dbName,
	)

	logMode := gormLogger.Silent
	if env.GetBool("MYSQL_QUERY_LOGGING", "false") {
		logMode = gormLogger.Info
	}

	//fmt.Printf("[database_resolver] creating new connection for user [%v] with dsn [%v]\n", user.ID, dsn)

	db, err := gorm.Open(
		mysql.New(mysql.Config{
			DSN:                  dsn,
			DisableWithReturning: true,
		}),
		&gorm.Config{
			SkipDefaultTransaction:                   true,
			DisableForeignKeyConstraintWhenMigrating: true,
			DisableAutomaticPing:                     false,
			FullSaveAssociations:                     false,
			Logger:                                   gormLogger.Default.LogMode(logMode),
		},
	)

	if err != nil {
		d.logger.Info().Any("connection", conn).Err(err).Msg("Failed to connect to user database connection")
		return nil
	}

	sqlDB, err := db.DB()
	if err != nil {
		d.logger.Info().Any("connection", conn).Err(err).Msg("Failed to connect to user database connection")
		return nil
	}

	sqlDB.SetConnMaxLifetime(time.Minute * 3)
	sqlDB.SetMaxIdleConns(env.GetInt("MYSQL_MAX_IDLE_CONNECTIONS", "10"))
	sqlDB.SetMaxOpenConns(env.GetInt("MYSQL_MAX_OPEN_CONNECTIONS", "150"))

	d.remoteDatabases[connectionType][conn.ServerDatabaseConnection.ID] = db

	return db
}

func (d *Resolver) handleWheres(query *gorm.DB, filter string) *gorm.DB {
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

func (d *Resolver) handleJsonWheres(query *gorm.DB, filter string) *gorm.DB {
	// parse where field = value
	wheres := strings.Split(filter, equalDelimiter)
	if len(wheres) > 1 {
		jsonDataField := wheres[0]
		jsonLookupField := wheres[1]
		lookupValue := wheres[2]
		if strings.Contains(jsonLookupField, "[*]") {
			query = query.Where(fmt.Sprintf("JSON_CONTAINS(JSON_EXTRACT(`%v`, '$%v'), ?, '$') = 1", jsonDataField, jsonLookupField), lookupValue)
		} else {
			query = query.Where(fmt.Sprintf("JSON_EXTRACT(`%v`, '$%v') = ?", jsonDataField, jsonLookupField), lookupValue)
		}
	}

	// parse where field contains value
	wheres = strings.Split(filter, likeDelimiter)
	if len(wheres) > 1 {
		jsonDataField := wheres[0]
		jsonLookupField := wheres[1]
		lookupValue := wheres[2]
		if strings.Contains(jsonLookupField, "[*]") {
			query = query.Where(fmt.Sprintf("JSON_SEARCH(JSON_EXTRACT(`%v`, '$%v'), 'all', ?, null, '$') IS NOT NULL", jsonDataField, jsonLookupField), "%"+lookupValue+"%")
		} else {
			query = query.Where(fmt.Sprintf("INSTR(JSON_EXTRACT(`%v`, '$%v'), ?) > 0", jsonDataField, jsonLookupField), lookupValue)
		}
	}

	return query
}

func (d *Resolver) handleOrWheres(query *gorm.DB, filter string) *gorm.DB {
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

func (d *Resolver) GetUserConnection(user models.User) models.UserServerDatabaseConnection {
	// cache
	key := fmt.Sprintf("active-user-db-connection-%v", user.ID)
	cachedConn, found := d.cache.Get(key)
	if found {
		return cachedConn.(models.UserServerDatabaseConnection)
	}

	// cache miss
	var conn models.UserServerDatabaseConnection
	relationships := models.UserServerDatabaseConnection{}.Relationships()
	query := d.GetSpireDb().Model(&models.UserServerDatabaseConnection{})
	for _, relationship := range relationships {
		query = query.Preload(relationship)
	}

	query.Where("user_id = ? and active = 1", user.ID).First(&conn)

	// set cache
	d.cache.Set(key, conn, 10*time.Minute)

	return conn
}

// failedConnectionAttempts holds the number of failed connection attempts
var failedConnectionAttempts = map[uint]uint{}

// connectionKeepAlive is used to keep connections alive
func (d *Resolver) connectionKeepAlive() {
	for {
		time.Sleep(1 * time.Second)

		// loop through all connections
		for _, connections := range d.remoteDatabases {
			for connectionId, connection := range connections {
				if connection != nil {
					// ping connection
					db, err := connection.DB()
					if err != nil {
						d.logger.Info().Any("connection", connectionId).Err(err).Msg("Failed to fetch database connection from MySQL")
						continue
					}

					// if we have a failed connection attempt
					if err := db.Ping(); err != nil {
						d.logger.Info().Any("connection", connectionId).Err(err).Msg("Failed to ping database connection")
						// if we have a failed connection attempt
						if _, ok := failedConnectionAttempts[connectionId]; ok {
							// increment
							failedConnectionAttempts[connectionId] = failedConnectionAttempts[connectionId] + 1

							// if we have 3 failed attempts, destroy connection
							if failedConnectionAttempts[connectionId] >= 3 {
								// delete map entry

								for c := range connectionTypes {
									if _, ok := d.remoteDatabases[connectionTypes[c]][connectionId]; ok {
										delete(d.remoteDatabases[connectionTypes[c]], connectionId)
									}
								}

								d.logger.Info().Any("connection", connectionId).Msg("Destroying connection after 3 failed attempts to ping")
							}
						} else {
							// init failed connection attempt
							failedConnectionAttempts[connectionId] = 1
						}
					}
				}
			}
		}
	}
}

// PurgeUserDbCache is used to purge the user database cache
func (d *Resolver) PurgeUserDbCache(id uint) {
	d.cache.Delete(fmt.Sprintf("active-connection-%v", id))
	for c := range connectionTypes {
		d.cache.Delete(fmt.Sprintf("active-connection-%v-%v", id, connectionTypes[c]))
	}
	d.cache.Delete(fmt.Sprintf("active-user-db-connection-%v", id))
}

// PurgeDatabaseConnections is used to purge the database connections
func (d *Resolver) PurgeDatabaseConnections() {
	for c := range connectionTypes {
		d.remoteDatabases[connectionTypes[c]] = map[uint]*gorm.DB{}
	}

	// clear cache
	for key, _ := range d.cache.Items() {
		if strings.Contains(key, "active-connection-") {
			d.cache.Delete(key)
		}
		if strings.Contains(key, "active-user-db-connection-") {
			d.cache.Delete(key)
		}
	}
}
