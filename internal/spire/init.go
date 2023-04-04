package spire

import (
	"fmt"
	"github.com/Akkadius/spire/internal/connection"
	"github.com/Akkadius/spire/internal/connection/contexts"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/encryption"
	"github.com/Akkadius/spire/internal/env"
	"github.com/Akkadius/spire/internal/eqemuserverconfig"
	"github.com/Akkadius/spire/internal/models"
	"github.com/go-sql-driver/mysql"
	gocache "github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
	gormMysql "gorm.io/driver/mysql"
)

// Init is a type responsible for bootstrapping and initializing the application
// where this lives may change
type Init struct {
	connections               *database.Connections
	logger                    *logrus.Logger
	isInitialized             bool // determines if spire as an app is initialized or not
	serverconfig              *eqemuserverconfig.Config
	dbConnectionCreateService *connection.DbConnectionCreateService
	cache                     *gocache.Cache
	settings                  *Settings
	crypt                     *encryption.Encrypter
	spireuser                 *UserService
}

func NewInit(
	connections *database.Connections,
	serverconfig *eqemuserverconfig.Config,
	logger *logrus.Logger,
	settings *Settings,
	cache *gocache.Cache,
	crypt *encryption.Encrypter,
	dbConnectionCreateService *connection.DbConnectionCreateService,
	spireuser *UserService,
) *Init {
	i := &Init{
		connections:               connections,
		logger:                    logger,
		isInitialized:             false,
		cache:                     cache,
		settings:                  settings,
		serverconfig:              serverconfig,
		crypt:                     crypt,
		dbConnectionCreateService: dbConnectionCreateService,
		spireuser:                 spireuser,
	}

	if env.IsAppEnvLocal() {
		i.Init()
	}

	// if running in production
	if !serverconfig.Exists() && env.IsAppEnvProduction() {
		i.isInitialized = true
	}

	return i
}

func (o *Init) Init() {
	o.settings.LoadSettings()
	o.isInitialized = o.CheckIfAppInitialized()

	// if we've set the app up initially but new settings have been added
	// let's automatically run them
	// otherwise these settings get populated when a user first sets up their
	// spire instance
	if o.isInitialized {
		_ = o.connections.SpireMigrate(false)
		o.settings.InitSettings()

		// get admin user
		var adminUser models.User
		o.connections.SpireDbNoLog().Where("is_admin = 1").First(&adminUser)

		if adminUser.ID > 0 {
			o.CreateDefaultDatabaseConnectionFromConfig(adminUser)
		}
	}
}

func (o *Init) IsInitialized() bool {
	return o.isInitialized
}

func (o *Init) SetIsInitialized(isInitialized bool) {
	o.isInitialized = isInitialized
}

// GetConnectionInfo will get the connection info for Spire to display to the
// user while going through the initial onboarding and setup pages
func (o *Init) GetConnectionInfo() *mysql.Config {
	f, ok := o.connections.SpireDbNoLog().Config.Dialector.(*gormMysql.Dialector)
	if ok {
		p, err := mysql.ParseDSN(f.DSN)
		if err != nil {
			o.logger.Error(err)
		}

		return p
	}
	return nil
}

// GetInstallationTables will return installation tables
func (o *Init) GetInstallationTables() []string {
	return o.connections.GetMigrationTables()
}

func (o *Init) CheckIfAppInitialized() bool {
	if o.settings.IsSettingEnabled(SettingAuthEnabled) {
		var adminUser models.User
		o.connections.SpireDbNoLog().Where("is_admin = 1").First(&adminUser)
		return adminUser.ID != 0
	}

	if o.settings.IsSettingDisabled(SettingAuthEnabled) {
		return true
	}

	return false
}

func (o *Init) SourceSpireTables() error {
	err := o.connections.SpireMigrate(false)
	if err != nil {
		return err
	}
	return nil
}

// CreateDefaultDatabaseConnectionFromConfig injects database connection configuration
// into Spire tables from the emulator server configuration
func (o *Init) CreateDefaultDatabaseConnectionFromConfig(user models.User) error {
	db := o.connections.SpireDbNoLog()
	cfg := o.serverconfig.Get()

	// connection already exists, let's just update it
	var c models.ServerDatabaseConnection
	db.Where("id = 1").First(&c)
	if c.ID > 0 {
		c.DbHost = cfg.Server.Database.Host
		c.DbPort = cfg.Server.Database.Port
		c.DbName = cfg.Server.Database.Db
		c.DbUsername = cfg.Server.Database.Username
		c.DbPassword = o.crypt.Encrypt(cfg.Server.Database.Password, o.GetEncKey(user.ID))

		// content db if exists
		if cfg.Server.ContentDatabase != nil {
			if len(cfg.Server.ContentDatabase.Host) > 0 {
				c.ContentDbHost = cfg.Server.ContentDatabase.Host
			}
			if len(cfg.Server.ContentDatabase.Port) > 0 {
				c.ContentDbPort = cfg.Server.ContentDatabase.Port
			}
			if len(cfg.Server.ContentDatabase.Db) > 0 {
				c.ContentDbName = cfg.Server.ContentDatabase.Db
			}
			if len(cfg.Server.ContentDatabase.Username) > 0 {
				c.ContentDbUsername = cfg.Server.ContentDatabase.Username
			}
			if len(cfg.Server.ContentDatabase.Password) > 0 {
				c.ContentDbPassword = o.crypt.Encrypt(cfg.Server.ContentDatabase.Password, o.GetEncKey(user.ID))
			}
		}

		db.Save(&c)

		return nil
	}

	// context
	ctx, err := contexts.NewConnectionCreateContext(
		user.ID,
		cfg.Server.World.Longname,
		cfg.Server.Database.Db,
		cfg.Server.Database.Host,
		cfg.Server.Database.Port,
		cfg.Server.Database.Username,
		cfg.Server.Database.Password,
	)
	if err != nil {
		return err
	}

	// If any content params are set
	if cfg.Server.ContentDatabase != nil {
		if len(cfg.Server.ContentDatabase.Db) > 0 {
			ctx.SetContentDbName(cfg.Server.ContentDatabase.Db)
		}
		if len(cfg.Server.ContentDatabase.Host) > 0 {
			ctx.SetContentDbHost(cfg.Server.ContentDatabase.Host)
		}
		if len(cfg.Server.ContentDatabase.Username) > 0 {
			ctx.SetContentDbUsername(cfg.Server.ContentDatabase.Username)
		}
		if len(cfg.Server.ContentDatabase.Password) > 0 {
			ctx.SetContentDbPassword(cfg.Server.ContentDatabase.Password)
		}
		if len(cfg.Server.ContentDatabase.Port) > 0 {
			ctx.SetContentDbPort(cfg.Server.ContentDatabase.Port)
		}
	}

	// created address
	ctx.SetCreatedFromIp("127.0.0.1")
	err = o.dbConnectionCreateService.Handle(ctx)
	if err != nil {
		return err
	}

	// purge any connection caching if exists
	o.spireuser.PurgeUserCache(user.ID)

	return nil
}

func (o *Init) GetEncKey(userId uint) string {
	return fmt.Sprintf("%v-%v", o.crypt.GetEncryptionKey(), userId)
}

func (o *Init) IsAuthEnabled() bool {
	return o.settings.IsSettingEnabled(SettingAuthEnabled)
}
