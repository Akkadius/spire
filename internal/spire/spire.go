package spire

import (
	"fmt"
	"github.com/Akkadius/spire/internal/connection"
	"github.com/Akkadius/spire/internal/connection/contexts"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/env"
	"github.com/Akkadius/spire/internal/models"
	"github.com/Akkadius/spire/internal/serverconfig"
	"github.com/go-sql-driver/mysql"
	gocache "github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
	gormMysql "gorm.io/driver/mysql"
)

// SpireInit is a type responsible for bootstrapping and initializing the application
// where this lives may change
type SpireInit struct {
	connections               *database.Connections
	logger                    *logrus.Logger
	isInitialized             bool // determines if spire as an app is initialized or not
	serverconfig              *serverconfig.EQEmuServerConfig
	dbConnectionCreateService *connection.DbConnectionCreateService
	cache                     *gocache.Cache
	settings                  *Settings
}

func NewSpire(
	connections *database.Connections,
	serverconfig *serverconfig.EQEmuServerConfig,
	logger *logrus.Logger,
	settings *Settings,
	cache *gocache.Cache,
	dbConnectionCreateService *connection.DbConnectionCreateService,
) *SpireInit {
	i := &SpireInit{
		connections:               connections,
		logger:                    logger,
		isInitialized:             false,
		cache:                     cache,
		settings:                  settings,
		serverconfig:              serverconfig,
		dbConnectionCreateService: dbConnectionCreateService,
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

func (o *SpireInit) Init() {
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

func (o *SpireInit) IsInitialized() bool {
	return o.isInitialized
}

func (o *SpireInit) SetIsInitialized(isInitialized bool) {
	o.isInitialized = isInitialized
}

// GetConnectionInfo will get the connection info for Spire to display to the
// user while going through the initial onboarding and setup pages
func (o *SpireInit) GetConnectionInfo() *mysql.Config {
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
func (o *SpireInit) GetInstallationTables() []string {
	return o.connections.GetMigrationTables()
}

func (o *SpireInit) CheckIfAppInitialized() bool {
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

func (o *SpireInit) SourceSpireTables() error {
	err := o.connections.SpireMigrate(false)
	if err != nil {
		return err
	}
	return nil
}

// CreateDefaultDatabaseConnectionFromConfig injects database connection configuration
// into Spire tables from the emulator server configuration
func (o *SpireInit) CreateDefaultDatabaseConnectionFromConfig(user models.User) error {
	db := o.connections.SpireDb()

	// delete users associated to connection
	db.Where("user_id = ?", user.ID).Delete(models.UserServerDatabaseConnection{})
	// delete connection itself
	db.Where("created_by = ?", user.ID).Delete(models.ServerDatabaseConnection{})

	cfg := o.serverconfig.Get()

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

	// created address
	ctx.SetCreatedFromIp("127.0.0.1")
	err = o.dbConnectionCreateService.Handle(ctx)
	if err != nil {
		return err
	}

	// purge any connection caching if exists
	o.cache.Delete(fmt.Sprintf("active-connection-%v", user.ID))
	o.cache.Delete(fmt.Sprintf("active-connection-%v-default", user.ID))
	o.cache.Delete(fmt.Sprintf("active-connection-%v-eqemu_content", user.ID))

	// hack to reset the indexes back to 1
	// local spire instances won't have any more than one database connection anyhow
	// local spire instances won't be able to support having multiple connections to manage
	db.Exec(fmt.Sprintf("UPDATE %v SET server_database_connection_id = 1", models.UserServerDatabaseConnection{}.TableName()))
	db.Exec(fmt.Sprintf("UPDATE %v SET id = 1", models.UserServerDatabaseConnection{}.TableName()))
	db.Exec(fmt.Sprintf("UPDATE %v SET id = 1", models.ServerDatabaseConnection{}.TableName()))

	return nil
}
