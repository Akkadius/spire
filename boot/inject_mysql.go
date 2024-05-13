package boot

import (
	"errors"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/env"
	"github.com/Akkadius/spire/internal/eqemuserverconfig"
	"github.com/Akkadius/spire/internal/logger"
	"gorm.io/driver/mysql"
	gormLogger "gorm.io/gorm/logger"
	"log"
	"os"
	"strconv"
	"time"

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
func provideAppDbConnections(serverconfig *eqemuserverconfig.Config, logger *logger.AppLogger) *database.Connections {
	eqEmuLocalDatabase, err := provideEQEmuLocalDatabase(serverconfig)
	if err != nil {
		logger.Fatal().Err(err).Msg("unable to provide eqemu local database")
	}

	spireDatabase, err := provideSpireDatabase(serverconfig)
	if err != nil {
		logger.Fatal().Err(err).Msg("unable to provide spire database")
	}

	// if we're local, we assume our emulator database
	// is where we are housing our spire tables
	if env.IsAppEnvLocal() {
		if spireDatabase == nil {
			spireDatabase = eqEmuLocalDatabase
		}
	}

	return database.NewConnections(
		spireDatabase,
		eqEmuLocalDatabase,
		logger,
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
func getEQEmuLocalMySQLConfig(serverconfig *eqemuserverconfig.Config) (*MySQLConfig, error) {
	m := &MySQLConfig{
		Password:           os.Getenv("MYSQL_EQEMU_PASSWORD"),
		Username:           os.Getenv("MYSQL_EQEMU_USERNAME"),
		Database:           os.Getenv("MYSQL_EQEMU_DATABASE"),
		Host:               os.Getenv("MYSQL_EQEMU_HOST"),
		Port:               env.GetInt("MYSQL_EQEMU_PORT", "3306"),
		MaxIdleConnections: env.GetInt("MYSQL_MAX_IDLE_CONNECTIONS", "10"),
		MaxOpenConnections: env.GetInt("MYSQL_MAX_OPEN_CONNECTIONS", "150"),
		EnableLogging:      isQueryLoggingEnabled(),
		ConnMaxLifetime:    env.GetInt("MYSQL_CONNECTION_MAX_LIFE_TIME", "5"),
	}

	// load eqemu config if exists
	config := serverconfig.Get()
	if config.Server.Database != nil && config.Server.Database.Db != "" {
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

	const errorPrefix string = "eqemu server local"

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

func isQueryLoggingEnabled() bool {
	//if env.IsAppEnvLocal() {
	//	return true
	//}

	return env.GetBool("MYSQL_QUERY_LOGGING", "false")
}

// provideEQEmuLocalDatabase is a Wire provider function that provides a
// database connection, configured from the environment.
// eqemu Local database is what is used if a remote connection is not provided otherwise
// If someone has a local installation of spire pointed to their own database this connection would be used
func provideEQEmuLocalDatabase(serverconfig *eqemuserverconfig.Config) (*gorm.DB, error) {
	config, err := getEQEmuLocalMySQLConfig(serverconfig)
	if err != nil {
		return nil, err
	}

	logMode := gormLogger.Silent
	if config.EnableLogging {
		logMode = gormLogger.Info
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: fmt.Sprintf(
			"%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
			config.Username,
			config.Password,
			config.Host,
			config.Port,
			config.Database,
		),
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
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetConnMaxLifetime(time.Minute * 3)

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
		EnableLogging:      isQueryLoggingEnabled(),
		ConnMaxLifetime:    env.GetInt("MYSQL_CONNECTION_MAX_LIFE_TIME", "5"),
	}

	const errorPrefix string = "eqemu server spire"

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
func provideSpireDatabase(serverconfig *eqemuserverconfig.Config) (*gorm.DB, error) {
	// if booting from local server folder
	cfg := serverconfig.Get()
	if cfg.Server.Database != nil && cfg.Server.Database.Db != "" {
		return nil, nil
	}

	config, err := getSpireMySQLConfig()
	if err != nil {
		return nil, err
	}

	logMode := gormLogger.Silent
	if config.EnableLogging {
		logMode = gormLogger.Info
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: fmt.Sprintf(
			"%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
			config.Username,
			config.Password,
			config.Host,
			config.Port,
			config.Database,
		),
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
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetConnMaxLifetime(time.Minute * 3)

	if config.MaxIdleConnections > 0 {
		sqlDB.SetMaxIdleConns(config.MaxIdleConnections)
	}
	if config.MaxOpenConnections > 0 {
		sqlDB.SetMaxOpenConns(config.MaxOpenConnections)
	}

	return db, nil
}
