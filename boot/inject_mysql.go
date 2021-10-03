package boot

import (
	"errors"
	"github.com/Akkadius/spire/env"
	"github.com/Akkadius/spire/internal/database"
	"gorm.io/driver/mysql"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"strconv"

	"fmt"

	"github.com/google/wire"
	"gorm.io/gorm"
)

// wire set for loading the stores.
var databaseSet = wire.NewSet(
	provideEQEmuLocalDatabase,
	provideAppDbConnections,
)

// we need to do this because
func provideAppDbConnections() *database.Connections {
	eqEmuLocalDatabase, err := provideEQEmuLocalDatabase()
	if err != nil {
		log.Fatal(err)
	}

	spireDatabase, err := provideSpireDatabase()
	if err != nil {
		log.Fatal(err)
	}

	return database.NewConnections(
		spireDatabase,
		eqEmuLocalDatabase,
	)
}

type MySQLConfig struct {
	Password           string
	Username           string
	Database           string
	Host               string
	MaxIdleConnections int
	MaxOpenConnections int
	EnableLogging      bool
	ConnMaxLifetime    int
	Port               int
}

// return mysql config
func getEQEmuLocalMySQLConfig() (*MySQLConfig, error) {
	m := &MySQLConfig{
		Password:           os.Getenv("MYSQL_EQEMU_PASSWORD"),
		Username:           os.Getenv("MYSQL_EQEMU_USERNAME"),
		Database:           os.Getenv("MYSQL_EQEMU_DATABASE"),
		Host:               os.Getenv("MYSQL_EQEMU_HOST"),
		Port:               env.GetInt("MYSQL_EQEMU_PORT", "3306"),
		MaxIdleConnections: env.GetInt("MYSQL_MAX_IDLE_CONNECTIONS", "10"),
		MaxOpenConnections: env.GetInt("MYSQL_MAX_OPEN_CONNECTIONS", "150"),
		EnableLogging:      env.GetBool("MYSQL_QUERY_LOGGING", "false"),
		ConnMaxLifetime:    env.GetInt("MYSQL_CONNECTION_MAX_LIFE_TIME", "5"),
	}

	// load eqemu config if exists
	config := getEQEmuConfig()
	if config.Server.Database.Db != "" {
		port, err := strconv.Atoi(config.Server.Database.Port)
		if err != nil {
			log.Fatalf("unable to convert string to integer error [%v]", err)
		}

		m.Username = config.Server.Database.Username
		m.Password = config.Server.Database.Password
		m.Host = config.Server.Database.Host
		m.Port = port
		m.Database = config.Server.Database.Db
	}

	const errorPrefix string = "eqemu Server Local"

	if len(m.Username) == 0 {
		return &MySQLConfig{}, errors.New(fmt.Sprintf("[%v] mysql username cannot be empty", errorPrefix))
	}
	if len(m.Database) == 0 {
		return &MySQLConfig{}, errors.New(fmt.Sprintf("[%v] mysql database cannot be empty", errorPrefix))
	}
	if len(m.Host) == 0 {
		return &MySQLConfig{}, errors.New(fmt.Sprintf("[%v] mysql host cannot be empty", errorPrefix))
	}
	if m.Port == 0 {
		return &MySQLConfig{}, errors.New(fmt.Sprintf("[%v] mysql database cannot be empty", errorPrefix))
	}

	return m, nil
}

// provideEQEmuLocalDatabase is a Wire provider function that provides a
// database connection, configured from the environment.
// eqemu Local database is what is used if a remote connection is not provided otherwise
// If someone has a local installation of spire pointed to their own database this connection would be used
func provideEQEmuLocalDatabase() (*gorm.DB, error) {
	config, err := getEQEmuLocalMySQLConfig()
	if err != nil {
		return nil, err
	}

	logMode := logger.Silent
	if config.EnableLogging {
		logMode = logger.Info
	}

	db, err := gorm.Open(mysql.Open(fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
	)),
		&gorm.Config{
			SkipDefaultTransaction:                   true,
			DisableForeignKeyConstraintWhenMigrating: true,
			DisableAutomaticPing:                     false,
			Logger: logger.Default.LogMode(logMode),
		},
	)

	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	if config.ConnMaxLifetime > 0 {
		sqlDB.SetMaxOpenConns(config.ConnMaxLifetime)
	}
	if config.MaxIdleConnections > 0 {
		sqlDB.SetMaxIdleConns(config.MaxIdleConnections)
	}
	if config.MaxOpenConnections > 0 {
		sqlDB.SetMaxOpenConns(config.MaxOpenConnections)
	}

	return db, nil
}

// return mysql config
func getSpireMySQLConfig() (*MySQLConfig, error) {
	m := &MySQLConfig{
		Password:           os.Getenv("MYSQL_SPIRE_PASSWORD"),
		Username:           os.Getenv("MYSQL_SPIRE_USERNAME"),
		Database:           os.Getenv("MYSQL_SPIRE_DATABASE"),
		Host:               os.Getenv("MYSQL_SPIRE_HOST"),
		Port:               env.GetInt("MYSQL_SPIRE_PORT", "3306"),
		MaxIdleConnections: env.GetInt("MYSQL_MAX_IDLE_CONNECTIONS", "10"),
		MaxOpenConnections: env.GetInt("MYSQL_MAX_OPEN_CONNECTIONS", "150"),
		EnableLogging:      env.GetBool("MYSQL_QUERY_LOGGING", "false"),
		ConnMaxLifetime:    env.GetInt("MYSQL_CONNECTION_MAX_LIFE_TIME", "5"),
	}

	const errorPrefix string = "eqemu Server Local"

	if len(m.Username) == 0 {
		return &MySQLConfig{}, errors.New(fmt.Sprintf("[%v] mysql username cannot be empty", errorPrefix))
	}
	if len(m.Database) == 0 {
		return &MySQLConfig{}, errors.New(fmt.Sprintf("[%v] mysql database cannot be empty", errorPrefix))
	}
	if len(m.Host) == 0 {
		return &MySQLConfig{}, errors.New(fmt.Sprintf("[%v] mysql host cannot be empty", errorPrefix))
	}
	if m.Port == 0 {
		return &MySQLConfig{}, errors.New(fmt.Sprintf("[%v] mysql database cannot be empty", errorPrefix))
	}

	return m, nil
}

// Local spire database connection
func provideSpireDatabase() (*gorm.DB, error) {
	config, err := getSpireMySQLConfig()
	if err != nil {
		return nil, err
	}

	// if booting from local server folder
	if getEQEmuConfig().Server.Database.Db != "" {
		return nil, nil
	}

	logMode := logger.Silent
	if config.EnableLogging {
		logMode = logger.Info
	}

	db, err := gorm.Open(mysql.Open(fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
	)),
		&gorm.Config{
			SkipDefaultTransaction:                   true,
			DisableForeignKeyConstraintWhenMigrating: true,
			DisableAutomaticPing:                     false,
			Logger: logger.Default.LogMode(logMode),
		},
	)

	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	if config.ConnMaxLifetime > 0 {
		sqlDB.SetMaxOpenConns(config.ConnMaxLifetime)
	}
	if config.MaxIdleConnections > 0 {
		sqlDB.SetMaxIdleConns(config.MaxIdleConnections)
	}
	if config.MaxOpenConnections > 0 {
		sqlDB.SetMaxOpenConns(config.MaxOpenConnections)
	}


	return db, nil
}
