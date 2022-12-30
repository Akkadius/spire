package spire

import (
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/models"
	"github.com/sirupsen/logrus"
)

const (
	SETTING_AUTH_ENABLED = "AUTH_ENABLED"
)

type Settings struct {
	connections *database.Connections
	logger      *logrus.Logger
	settings    []models.Setting
}

func NewSettings(
	connections *database.Connections,
	logger *logrus.Logger,
) *Settings {
	return &Settings{
		connections: connections,
		logger:      logger,
	}
}

// InitSettings will initialize default settings where they don't otherwise exist
// This will run during onboarding of the app (initialization) and will run
// every bootup after the app has been onboarded for the first time
func (o *Settings) InitSettings() {
	defaultSettings := []models.Setting{
		models.Setting{
			Setting: SETTING_AUTH_ENABLED,
			Value:   "false",
		},
	}

	var currentSettings []models.Setting
	o.connections.SpireDbNoLog().Find(&currentSettings)

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
		o.connections.SpireDbNoLog().Create(settingsToCreate)
	}
}

// LoadSettings will load settings into memory
func (o *Settings) LoadSettings() {
	var currentSettings []models.Setting
	o.connections.SpireDbNoLog().Find(&currentSettings)
	o.settings = currentSettings
}

// EnableSetting will enable settings that have true/false value
func (o *Settings) EnableSetting(setting string) {
	var s models.Setting
	o.connections.SpireDbNoLog().First(&s, "setting = ?", setting)
	if s.ID == 0 {
		s.Setting = setting
		s.Value = "true"
		o.connections.SpireDbNoLog().Create(&s)
	}
	if s.ID > 0 {
		s.Value = "true"
		o.connections.SpireDbNoLog().Save(&s)
	}
}

// DisableSetting will enable settings that have true/false value
func (o *Settings) DisableSetting(setting string) {
	var s models.Setting
	o.connections.SpireDbNoLog().First(&s, "setting = ?", setting)
	if s.ID == 0 {
		s.Setting = setting
		s.Value = "false"
		o.connections.SpireDbNoLog().Create(&s)
	}
	if s.ID > 0 {
		s.Value = "false"
		o.connections.SpireDbNoLog().Save(&s)
	}
}

// IsSettingEnabled checks if a setting is enabled (boolean)
// out of what is loaded into memory
func (o *Settings) IsSettingEnabled(setting string) bool {
	for _, s := range o.settings {
		if s.Setting == setting {
			return s.Value == "true"
		}
	}

	return false
}

func (o *Settings) GetSettings() []models.Setting {
	return o.settings
}
