package connection

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/Akkadius/spire/internal/connection/contexts"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/encryption"
	"github.com/Akkadius/spire/internal/env"
	"github.com/Akkadius/spire/internal/models"
	"github.com/sirupsen/logrus"
)

type DbConnectionCreateService struct {
	db     *database.DatabaseResolver
	logger *logrus.Logger
	crypt  *encryption.Encrypter
}

func NewDbConnectionCreateService(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
	crypt *encryption.Encrypter,
) *DbConnectionCreateService {
	return &DbConnectionCreateService{db: db, logger: logger, crypt: crypt}
}

func (c *DbConnectionCreateService) GetEncKey(userId uint) string {
	return fmt.Sprintf("%v-%v", env.Get("APP_KEY", ""), userId)
}

func (c *DbConnectionCreateService) Handle(ctx *contexts.ConnectionCreateContext) error {

	// validate connection doesn't already exist
	var con models.ServerDatabaseConnection
	c.db.GetSpireDb().Where(
		"name = ? and db_host = ? and db_name = ? and db_port = ? and db_username = ?",
		ctx.ConnectionName(),
		ctx.DbHost(),
		ctx.DbName(),
		ctx.DbPort(),
		ctx.DbUsername(),
	).First(&con)
	if con.ID > 0 {
		return errors.New("connection already exists")
	}

	// check connection
	dsn := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local&timeout=1s",
		ctx.DbUsername(),
		ctx.DbPassword(),
		ctx.DbHost(),
		ctx.DbPort(),
		ctx.DbName(),
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return errors.New(fmt.Sprintf("database connection error: %v", err))
	}

	// If using a content database connection also check it
	if ctx.ContentDbUsername() != "" {
		dsn := fmt.Sprintf(
			"%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local&timeout=1s",
			ctx.ContentDbUsername(),
			ctx.ContentDbPassword(),
			ctx.ContentDbHost(),
			ctx.ContentDbPort(),
			ctx.ContentDbName(),
		)

		db, err := sql.Open("mysql", dsn)
		if err != nil {
			return errors.New(fmt.Sprintf("content database connection error: %v", err))
		}

		err = db.Ping()
		if err != nil {
			return errors.New(fmt.Sprintf("content database connection error: %v", err))
		}
	}

	// not required so lets not encrypt an empty string
	contentDbPassword := ""
	if len(ctx.ContentDbPassword()) > 0 {
		contentDbPassword = c.crypt.Encrypt(ctx.ContentDbPassword(), c.GetEncKey(ctx.UserId()))
	}

	// server database connection
	var connection models.ServerDatabaseConnection
	connection.Name = ctx.ConnectionName()
	connection.DbHost = ctx.DbHost()
	connection.DbPort = ctx.DbPort()
	connection.DbName = ctx.DbName()
	connection.DbUsername = ctx.DbUsername()
	connection.DbPassword = c.crypt.Encrypt(ctx.DbPassword(), c.GetEncKey(ctx.UserId()))
	connection.ContentDbHost = ctx.ContentDbHost()
	connection.ContentDbPort = ctx.ContentDbPort()
	connection.ContentDbName = ctx.ContentDbName()
	connection.ContentDbUsername = ctx.ContentDbUsername()
	connection.ContentDbPassword = contentDbPassword
	connection.CreatedFromIp = ctx.CreatedFromIp()
	connection.CreatedBy = ctx.UserId()
	err = c.db.GetSpireDb().Create(&connection).Error
	if err != nil {
		return err
	}

	// Set other connections inactive
	c.db.GetSpireDb().Model(models.UserServerDatabaseConnection{}).Where("user_id = ?", ctx.UserId()).Update(
		"active",
		"0",
	)

	// associate connection with user
	var userConnection models.UserServerDatabaseConnection
	userConnection.UserId = ctx.UserId()
	userConnection.Active = 1
	userConnection.ServerDatabaseConnectionId = connection.ID
	userConnection.CreatedBy = ctx.UserId()
	err = c.db.GetSpireDb().Create(&userConnection).Error
	if err != nil {
		return err
	}

	return nil
}
