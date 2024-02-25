package connection

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/Akkadius/spire/internal/connection/contexts"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/encryption"
	"github.com/Akkadius/spire/internal/models"
	"github.com/sirupsen/logrus"
)

type Create struct {
	db     *database.Resolver
	logger *logrus.Logger
	crypt  *encryption.Encrypter
}

func NewCreate(
	db *database.Resolver,
	logger *logrus.Logger,
	crypt *encryption.Encrypter,
) *Create {
	return &Create{db: db, logger: logger, crypt: crypt}
}

func (c *Create) GetEncKey(userId uint) string {
	return fmt.Sprintf("%v-%v", c.crypt.GetEncryptionKey(), userId)
}

func (c *Create) Handle(ctx *contexts.ConnectionCreateContext) error {

	// validate valid user before creating
	if ctx.UserId() == 0 {
		return errors.New("user must be logged in to create connections")
	}

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
	if ctx.CheckDbConnection() && ctx.ContentDbUsername() != "" {
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

	// if using a logs database connection also check it
	if ctx.CheckDbConnection() && ctx.LogsDbUsername() != "" {
		dsn := fmt.Sprintf(
			"%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local&timeout=1s",
			ctx.LogsDbUsername(),
			ctx.LogsDbPassword(),
			ctx.LogsDbHost(),
			ctx.LogsDbPort(),
			ctx.LogsDbName(),
		)

		db, err := sql.Open("mysql", dsn)
		if err != nil {
			return errors.New(fmt.Sprintf("logs database connection error: %v", err))
		}

		err = db.Ping()
		if err != nil {
			return errors.New(fmt.Sprintf("logs database connection error: %v", err))
		}
	}

	// not required so lets not encrypt an empty string
	contentDbPassword := ""
	if len(ctx.ContentDbPassword()) > 0 {
		contentDbPassword = c.crypt.Encrypt(ctx.ContentDbPassword(), c.GetEncKey(ctx.UserId()))
	}
	logsDbPassword := ""
	if len(ctx.LogsDbPassword()) > 0 {
		logsDbPassword = c.crypt.Encrypt(ctx.LogsDbPassword(), c.GetEncKey(ctx.UserId()))
	}

	// server database connection
	var connection models.ServerDatabaseConnection
	connection.Name = ctx.ConnectionName()
	connection.DbHost = ctx.DbHost()
	connection.DbPort = ctx.DbPort()
	connection.DbName = ctx.DbName()
	connection.DbUsername = ctx.DbUsername()
	connection.DbPassword = c.crypt.Encrypt(ctx.DbPassword(), c.GetEncKey(ctx.UserId()))

	// content database connection
	connection.ContentDbHost = ctx.ContentDbHost()
	connection.ContentDbPort = ctx.ContentDbPort()
	connection.ContentDbName = ctx.ContentDbName()
	connection.ContentDbUsername = ctx.ContentDbUsername()
	connection.ContentDbPassword = contentDbPassword

	// logs database connection
	connection.LogsDbHost = ctx.LogsDbHost()
	connection.LogsDbPort = ctx.LogsDbPort()
	connection.LogsDbName = ctx.LogsDbName()
	connection.LogsDbUsername = ctx.LogsDbUsername()
	connection.LogsDbPassword = logsDbPassword

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
