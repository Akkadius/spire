package eqemuserverconfig

import (
	"encoding/json"
	"fmt"
	"github.com/Akkadius/spire/internal/pathmgmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

type Config struct {
	logger   *logrus.Logger
	pathmgmt *pathmgmt.PathManagement
	config   EQEmuConfigJson
}

func NewConfig(logger *logrus.Logger, pathmgmt *pathmgmt.PathManagement) *Config {
	return &Config{
		logger:   logger,
		pathmgmt: pathmgmt,
	}
}

type DatabaseConfig struct {
	Db       string `json:"db,omitempty"`
	Host     string `json:"host,omitempty"`
	Port     string `json:"port,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type EQEmuConfigJson struct {
	Server struct {
		Zones struct {
			Defaultstatus string `json:"defaultstatus"`
			Ports         struct {
				Low  string `json:"low"`
				High string `json:"high"`
			} `json:"ports"`
		} `json:"zones,omitempty"`
		Qsdatabase *DatabaseConfig `json:"qsdatabase,omitempty"`
		Chatserver struct {
			Port string `json:"port"`
			Host string `json:"host"`
		} `json:"chatserver,omitempty"`
		Mailserver struct {
			Host string `json:"host"`
			Port string `json:"port"`
		} `json:"mailserver,omitempty"`
		World struct {
			API struct {
				Enabled bool `json:"enabled"`
			} `json:"api,omitempty"`
			Address      string `json:"address,omitempty"`
			Localaddress string `json:"localaddress,omitempty"`
			Loginserver1 struct {
				Port     string `json:"port"`
				Account  string `json:"account"`
				Password string `json:"password"`
				Host     string `json:"host"`
			} `json:"loginserver1,omitempty"`
			Loginserver2 struct {
				Account  string `json:"account"`
				Password string `json:"password"`
				Host     string `json:"host"`
				Port     string `json:"port"`
			} `json:"loginserver2,omitempty"`
			TCP struct {
				IP   string `json:"ip"`
				Port string `json:"port"`
			} `json:"tcp,omitempty"`
			Telnet struct {
				IP      string `json:"ip"`
				Port    string `json:"port"`
				Enabled string `json:"enabled"`
			} `json:"telnet,omitempty"`
			Key          string `json:"key"`
			Shortname    string `json:"shortname"`
			Longname     string `json:"longname"`
			Loginserver3 struct {
				Account  string `json:"account"`
				Password string `json:"password"`
				Legacy   string `json:"legacy"`
				Host     string `json:"host"`
				Port     string `json:"port"`
			} `json:"loginserver3,omitempty"`
		} `json:"world"`
		Database        *DatabaseConfig `json:"database"`
		ContentDatabase *DatabaseConfig `json:"content_database,omitempty"`
		Files           struct {
			Opcodes     string `json:"opcodes"`
			MailOpcodes string `json:"mail_opcodes"`
		} `json:"files"`
		Directories struct {
			Patches string `json:"patches"`
			Opcodes string `json:"opcodes"`
		} `json:"directories"`
	} `json:"server"`
	WebAdmin struct { // Occulus
		Discord *struct {
			CrashLogWebhook string `json:"crash_log_webhook,omitempty"`
		} `json:"discord,omitempty"`
		Application struct {
			Key   string `json:"key,omitempty"`
			Admin struct {
				Password string `json:"password,omitempty"`
			} `json:"admin,omitempty"`
		} `json:"application,omitempty"`
		Launcher *struct {
			RunSharedMemory  bool   `json:"runSharedMemory"`
			RunLoginserver   bool   `json:"runLoginserver"`
			RunQueryServ     bool   `json:"runQueryServ"`
			IsRunning        bool   `json:"isRunning"`
			MinZoneProcesses int    `json:"minZoneProcesses"`
			StaticZones      string `json:"staticZones,omitempty"`
		} `json:"launcher,omitempty"`
		Quests struct {
			HotReload bool `json:"hotReload"`
		} `json:"quests"`
		ServerCodePath string `json:"serverCodePath,omitempty"`
	} `json:"web-admin,omitempty"`
	Spire struct {
		EncryptionKey string `json:"encryption_key,omitempty"`
		HttpPort      int    `json:"http_port,omitempty"`
	} `json:"spire,omitempty"`
}

func (e *Config) Get() EQEmuConfigJson {
	cfg := e.pathmgmt.GetEQEmuServerConfigFilePath()
	e.debug("path [%v]", cfg)

	if len(cfg) > 0 {
		e.debug("Reading eqemu config file [%v]", cfg)

		body, err := ioutil.ReadFile(e.pathmgmt.GetEQEmuServerConfigFilePath())
		if err != nil {
			e.logger.Fatalf("unable to read file: %v", err)
		}

		config := EQEmuConfigJson{}
		err = json.Unmarshal(body, &config)
		if err != nil {
			e.logger.Fatalf("unable to unmarshal file [%v] error [%v]", cfg, err)
		}

		return config
	}

	return EQEmuConfigJson{}
}

func (e *Config) debug(msg string, a ...interface{}) {
	if len(os.Getenv("DEBUG")) >= 3 {
		if len(a) > 0 {
			e.logger.Debug("[eqemu_server_config.go] " + fmt.Sprintf(msg, a...) + "\n")
			return
		}
		e.logger.Debug("[eqemu_server_config.go] " + fmt.Sprintf(msg) + "\n")
	}
}

// Exists returns true if the *Config file exists
func (e *Config) Exists() bool {
	return len(e.pathmgmt.GetEQEmuServerConfigFilePath()) > 0
}

// Save saves the *Config to disk
func (e *Config) Save(c EQEmuConfigJson) error {
	if c.WebAdmin.Discord != nil {
		if len(c.WebAdmin.Discord.CrashLogWebhook) == 0 {
			c.WebAdmin.Discord = nil
		}
	}
	if c.WebAdmin.Launcher != nil {
		if c.WebAdmin.Launcher.MinZoneProcesses == 0 {
			c.WebAdmin.Launcher.MinZoneProcesses = 10
		}
	}

	file, err := json.MarshalIndent(c, "", " ")
	if err != nil {
		return err
	}

	err = os.WriteFile(e.pathmgmt.GetEQEmuServerConfigFilePath(), file, 0755)
	if err != nil {
		return err
	}

	return nil
}
