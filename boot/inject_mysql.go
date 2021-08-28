package boot

import (
	"errors"
	"github.com/Akkadius/spire/database"
	"github.com/Akkadius/spire/env"
	"log"
	"os"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
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

	mysql, err := gorm.Open(
		"mysql",
		fmt.Sprintf(
			"%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
			config.Username,
			config.Password,
			config.Host,
			config.Port,
			config.Database,
		),
	)
	if err != nil {
		return nil, err
	}
	if config.ConnMaxLifetime > 0 {
		mysql.DB().SetMaxOpenConns(config.ConnMaxLifetime)
	}
	if config.MaxIdleConnections > 0 {
		mysql.DB().SetMaxIdleConns(config.MaxIdleConnections)
	}
	if config.MaxOpenConnections > 0 {
		mysql.DB().SetMaxOpenConns(config.MaxOpenConnections)
	}
	if config.EnableLogging {
		mysql.LogMode(config.EnableLogging)
	}

	return mysql, nil
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

	mysql, err := gorm.Open(
		"mysql",
		fmt.Sprintf(
			"%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
			config.Username,
			config.Password,
			config.Host,
			config.Port,
			config.Database,
		),
	)
	if err != nil {
		return nil, err
	}
	if config.ConnMaxLifetime > 0 {
		mysql.DB().SetMaxOpenConns(config.ConnMaxLifetime)
	}
	if config.MaxIdleConnections > 0 {
		mysql.DB().SetMaxIdleConns(config.MaxIdleConnections)
	}
	if config.MaxOpenConnections > 0 {
		mysql.DB().SetMaxOpenConns(config.MaxOpenConnections)
	}
	if config.EnableLogging {
		mysql.LogMode(config.EnableLogging)
	}

	return mysql, nil
}
