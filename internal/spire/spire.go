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
	settings      []models.Setting
}

func NewSpire(connections *database.Connections, serverconfig *serverconfig.EQEmuServerConfig, logger *logrus.Logger) *SpireInit {
	i := &SpireInit{
		connections:   connections,
		logger:        logger,
		isInitialized: false,
	}

	// if running in production
	if !serverconfig.Exists() && env.IsAppEnvProduction() {
		i.isInitialized = true
	}

	return i
}

func (o *SpireInit) Init() {
	o.connections.SpireMigrate(false)
	o.InitSettings()
	o.LoadSettings()
}

func (o *SpireInit) IsInitialized() bool {
	return o.isInitialized
}

func (o *SpireInit) SetIsInitialized(isInitialized bool) {
	o.isInitialized = isInitialized
}

func (o *SpireInit) InitSettings() {
	defaultSettings := []models.Setting{
		models.Setting{
			Setting: "AUTH_ENABLED",
			Value:   "false",
		},
	}

	var currentSettings []models.Setting
	o.connections.SpireDb().Find(&currentSettings)

	// inject defaults
	var settingsToCreate []models.Setting
	for _, d := range defaultSettings {
		settingExists := false
		for _, s := range currentSettings {
			if s.Setting == d.Setting {
				settingExists = true
			}
		}
		if !settingExists {
			settingsToCreate = append(settingsToCreate, d)
		}
	}

	// if we don't have some of the default settings, inject them
	if len(settingsToCreate) > 0 {
		o.connections.SpireDb().Create(settingsToCreate)
	}
}

func (o *SpireInit) LoadSettings() {
	var currentSettings []models.Setting
	o.connections.SpireDb().Find(&currentSettings)
	o.settings = currentSettings
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
