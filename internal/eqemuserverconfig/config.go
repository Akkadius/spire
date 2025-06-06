package eqemuserverconfig

import (
	"encoding/json"
	"github.com/Akkadius/spire/internal/logger"
	"github.com/Akkadius/spire/internal/pathmgmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// Config is the struct type
type Config struct {
	logger   *logger.AppLogger
	pathmgmt *pathmgmt.PathManagement
	config   EQEmuConfigJson
}

// NewConfig creates a new Config struct
func NewConfig(logger *logger.AppLogger, pathmgmt *pathmgmt.PathManagement) *Config {
	return &Config{
		logger:   logger,
		pathmgmt: pathmgmt,
	}
}

// DatabaseConfig is the struct that represents the database configuration in eqemu_config.json
type DatabaseConfig struct {
	Db       string `json:"db,omitempty"`
	Host     string `json:"host,omitempty"`
	Port     string `json:"port,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type WebAdminLauncherConfig struct {
	RunSharedMemory             bool   `json:"runSharedMemory"`
	RunLoginserver              bool   `json:"runLoginserver"`
	RunQueryServ                bool   `json:"runQueryServ"`
	RunUcs                      bool   `json:"runUcs"`
	IsRunning                   bool   `json:"isRunning"`
	MinZoneProcesses            int    `json:"minZoneProcesses,omitempty"`
	StaticZones                 string `json:"staticZones,omitempty"`
	UpdateOpcodesOnStart        bool   `json:"updateOpcodesOnStart"`
	DeleteLogFilesOlderThanDays int    `json:"deleteLogFilesOlderThanDays"`

	// leaf nodes do not run any process but zones
	DistributedNodeType     string `json:"distributed_node_type,omitempty"`      // root or leaf
	DistributedMaxZoneCount int    `json:"distributed_max_zone_count,omitempty"` // max zone count for leaf nodes
}

type WebAdminQuestsConfig struct {
	HotReload bool `json:"hotReload"`
}

type WebAdminConfig struct { // Occulus
	Version string `json:"version,omitempty"`
	Discord *struct {
		CrashLogWebhook string `json:"crash_log_webhook,omitempty"`
	} `json:"discord,omitempty"`
	Application struct {
		Key   string `json:"key,omitempty"`
		Admin struct {
			Password string `json:"password,omitempty"`
		} `json:"admin,omitempty"`
	} `json:"application,omitempty"`
	Launcher       *WebAdminLauncherConfig `json:"launcher,omitempty"`
	Quests         WebAdminQuestsConfig    `json:"quests"`
	ServerCodePath string                  `json:"serverCodePath,omitempty"`
}

// EQEmuConfigJson is the struct that represents the eqemu_config.json file
type EQEmuConfigJson struct {
	Server struct {
		Discord *struct {
			Channelid    string `json:"channelid"`
			Itemurl      string `json:"itemurl"`
			Refreshrate  string `json:"refreshrate"`
			Clientid     string `json:"clientid"`
			Clientsecret string `json:"clientsecret"`
			Serverid     string `json:"serverid"`
			Username     string `json:"username"`
		} `json:"discord,omitempty"`
		Zones struct {
			Defaultstatus string `json:"defaultstatus"`
			Ports         struct {
				Low  string `json:"low"`
				High string `json:"high"`
			} `json:"ports"`
		} `json:"zones,omitempty"`
		Qsdatabase *DatabaseConfig `json:"qsdatabase,omitempty"`
		// deprecated
		Chatserver *struct {
			Port string `json:"port"`
			Host string `json:"host"`
		} `json:"chatserver,omitempty"`
		// deprecated
		Mailserver *struct {
			Host string `json:"host"`
			Port string `json:"port"`
		} `json:"mailserver,omitempty"`
		// queryserver connection
		Queryserver *struct {
			Host string `json:"host"`
			Port string `json:"port"`
		} `json:"queryserver,omitempty"`
		// replaces chatserver and mailserver
		Ucs *struct {
			Host string `json:"host"`
			Port string `json:"port"`
		} `json:"ucs,omitempty"`
		World struct {
			AutoDatabaseUpdates *bool `json:"auto_database_updates,omitempty"`
			Locked              bool  `json:"locked,omitempty"`
			API                 struct {
				Enabled bool `json:"enabled"`
			} `json:"api,omitempty"`
			Address      string `json:"address,omitempty"`
			Localaddress string `json:"localaddress,omitempty"`
			Loginserver1 *struct {
				Port     string `json:"port"`
				Account  string `json:"account"`
				Password string `json:"password"`
				Host     string `json:"host"`
				Legacy   string `json:"legacy,omitempty"`
			} `json:"loginserver1,omitempty"`
			Loginserver2 *struct {
				Account  string `json:"account"`
				Password string `json:"password"`
				Legacy   string `json:"legacy,omitempty"`
				Host     string `json:"host"`
				Port     string `json:"port"`
			} `json:"loginserver2,omitempty"`
			Loginserver3 *struct {
				Account  string `json:"account"`
				Password string `json:"password"`
				Legacy   string `json:"legacy,omitempty"`
				Host     string `json:"host"`
				Port     string `json:"port"`
			} `json:"loginserver3,omitempty"`
			Loginserver4 *struct {
				Account  string `json:"account"`
				Password string `json:"password"`
				Legacy   string `json:"legacy,omitempty"`
				Host     string `json:"host"`
				Port     string `json:"port"`
			} `json:"loginserver4,omitempty"`
			Loginserver5 *struct {
				Account  string `json:"account"`
				Password string `json:"password"`
				Legacy   string `json:"legacy,omitempty"`
				Host     string `json:"host"`
				Port     string `json:"port"`
			} `json:"loginserver5,omitempty"`
			TCP struct {
				IP   string `json:"ip"`
				Port string `json:"port"`
			} `json:"tcp,omitempty"`
			Telnet struct {
				IP      string `json:"ip"`
				Port    string `json:"port"`
				Enabled string `json:"enabled"`
			} `json:"telnet,omitempty"`
			Key       string `json:"key"`
			Shortname string `json:"shortname"`
			Longname  string `json:"longname"`
		} `json:"world"`
		Database        *DatabaseConfig `json:"database"`
		ContentDatabase *DatabaseConfig `json:"content_database,omitempty"`
		Files           struct {
			Opcodes     string `json:"opcodes"`
			MailOpcodes string `json:"mail_opcodes"`
		} `json:"files"`
		Directories struct {
			Quests         string   `json:"quests,omitempty"`
			Plugins        string   `json:"plugins,omitempty"`
			Patches        string   `json:"patches,omitempty"`
			LuaModules     string   `json:"lua_modules,omitempty"`
			Maps           string   `json:"maps,omitempty"`
			SharedMemory   string   `json:"shared_memory,omitempty"`
			Logs           string   `json:"logs,omitempty"`
			Opcodes        string   `json:"opcodes,omitempty"`
			QuestPaths     []string `json:"quest_paths,omitempty"`
			PluginPaths    []string `json:"plugin_paths,omitempty"`
			LuaModulePaths []string `json:"lua_module_paths,omitempty"`
		} `json:"directories"`
	} `json:"server"`
	WebAdmin *WebAdminConfig `json:"web-admin,omitempty"`
	Spire    struct {
		EncryptionKey      string `json:"encryption_key,omitempty"`
		HttpPort           int    `json:"http_port,omitempty"`
		LauncherStart      bool   `json:"launcher_start"`                 // starts server launcher
		DisableAutoUpdates bool   `json:"disable_auto_updates,omitempty"` // disable auto updates
	} `json:"spire,omitempty"`
}

var cachedConfig *EQEmuConfigJson
var lastModifiedTime time.Time
var mutex = &sync.Mutex{}

// Get returns the eqemu config json
// If the file has been modified since the last read, it will re-read the file
// We utilize a cache to prevent reading the file on every call
func (e *Config) Get() (EQEmuConfigJson, error) {
	configFile := e.pathmgmt.GetEQEmuServerConfigFilePath()
	if len(configFile) > 0 {
		stat, err := os.Stat(configFile)
		if err != nil {
			return EQEmuConfigJson{}, err
		}

		if stat.ModTime().After(lastModifiedTime) || lastModifiedTime.IsZero() {
			e.logger.Debug().Any("path", configFile).Msg("Reading eqemu config file")
			body, err := os.ReadFile(e.pathmgmt.GetEQEmuServerConfigFilePath())
			if err != nil {
				return EQEmuConfigJson{}, err
			}

			config := EQEmuConfigJson{}
			err = json.Unmarshal(body, &config)
			if err != nil {
				return EQEmuConfigJson{}, err
			}

			e.setConfigDefaults(&config)

			lastModifiedTime = stat.ModTime()
			mutex.Lock()
			cachedConfig = &config
			mutex.Unlock()

			return config, nil
		} else if cachedConfig != nil {
			return *cachedConfig, nil
		}
	}

	return EQEmuConfigJson{}, nil
}

// Exists will return true if the eqemu_config.json file exists
func (e *Config) Exists() bool {
	return len(e.pathmgmt.GetEQEmuServerConfigFilePath()) > 0
}

// Save will save the config to the eqemu_config.json file
func (e *Config) Save(c EQEmuConfigJson) error {
	if c.WebAdmin != nil && c.WebAdmin.Discord != nil {
		if len(c.WebAdmin.Discord.CrashLogWebhook) == 0 {
			c.WebAdmin.Discord = nil
		}
		if c.WebAdmin.Launcher != nil {
			if c.WebAdmin.Launcher.MinZoneProcesses == 0 {
				c.WebAdmin.Launcher.MinZoneProcesses = 10
			}
		}
	}

	file, err := json.MarshalIndent(c, "", " ")
	if err != nil {
		return err
	}

	path := filepath.Join(e.pathmgmt.GetEQEmuServerPath(), "eqemu_config.json")
	err = os.WriteFile(path, file, 0755)
	if err != nil {
		return err
	}

	cachedConfig = &c

	return nil
}

// setConfigDefaults will set the default values for the config
// especially for pointer values
func (e *Config) setConfigDefaults(c *EQEmuConfigJson) {
	save := false
	if c.WebAdmin == nil {
		e.logger.Debug().Msg("Setting default web-admin config")
		c.WebAdmin = &WebAdminConfig{
			Quests: WebAdminQuestsConfig{
				HotReload: true,
			},
		}
		save = true
	}
	if c.WebAdmin != nil {
		if c.WebAdmin.Launcher == nil {
			e.logger.Debug().Msg("Setting default web-admin launcher config")
			c.WebAdmin.Launcher = &WebAdminLauncherConfig{
				RunSharedMemory:             true,
				RunLoginserver:              false,
				RunQueryServ:                false,
				RunUcs:                      true,
				MinZoneProcesses:            10,
				UpdateOpcodesOnStart:        true,
				DeleteLogFilesOlderThanDays: 7,
			}
			save = true
		}
		if c.WebAdmin.Launcher != nil {
			if c.WebAdmin.Launcher.MinZoneProcesses == 0 {
				e.logger.Debug().Msg("Setting default web-admin launcher min zone processes")
				c.WebAdmin.Launcher.MinZoneProcesses = 10
				save = true
			}

			if c.WebAdmin.Version == "" {
				e.logger.Debug().Msg("Setting default web-admin config version")
				c.WebAdmin.Version = "1.0.0"
				c.WebAdmin.Launcher.RunUcs = true
				save = true
			}
		}
	}
	if save {
		e.Save(*c)
	}
}
