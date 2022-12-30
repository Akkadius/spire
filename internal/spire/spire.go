package spire

import (
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/env"
	"github.com/Akkadius/spire/internal/models"
	"github.com/Akkadius/spire/internal/serverconfig"
	"github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	gormMysql "gorm.io/driver/mysql"
)

// SpireInit is a type responsible for bootstrapping and initializing the application
// where this lives may change
type SpireInit struct {
	connections   *database.Connections
	logger        *logrus.Logger
	isInitialized bool // determines if spire as an app is initialized or not
	serverconfig  *serverconfig.EQEmuServerConfig
	settings      *Settings
}

func NewSpire(
	connections *database.Connections,
	serverconfig *serverconfig.EQEmuServerConfig,
	logger *logrus.Logger,
	settings *Settings,
) *SpireInit {
	i := &SpireInit{
		connections:   connections,
		logger:        logger,
		isInitialized: false,
		settings:      settings,
	}

	// if running in production
	if !serverconfig.Exists() && env.IsAppEnvProduction() {
		i.isInitialized = true
	}

	return i
}

func (o *SpireInit) Init() {
	o.isInitialized = o.CheckIfAppInitialized()

	// if we've set the app up initially but new settings have been added
	// let's automatically run them
	// otherwise these settings get populated when a user first sets up their
	// spire instance
	if o.isInitialized {
		o.connections.SpireMigrate(false)
		o.settings.InitSettings()
	}

	o.settings.LoadSettings()
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
	f, ok := o.connections.SpireDb().Config.Dialector.(*gormMysql.Dialector)
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
	var settings []models.Setting
	o.connections.SpireDb().Find(&settings)
	var adminUser models.User
	o.connections.SpireDb().Where("is_admin = 1").First(&adminUser)

	return len(settings) > 0 && adminUser.ID != 0
}
