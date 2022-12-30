package serverconfig

import (
	"encoding/json"
	"fmt"
	"github.com/Akkadius/spire/internal/pathmgmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

type EQEmuServerConfig struct {
	logger   *logrus.Logger
	pathmgmt *pathmgmt.PathManagement
	config   EQEmuConfigJson
}

func NewEQEmuServerConfig(logger *logrus.Logger, pathmgmt *pathmgmt.PathManagement) *EQEmuServerConfig {
	return &EQEmuServerConfig{
		logger:   logger,
		pathmgmt: pathmgmt,
	}
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
		Qsdatabase struct {
			Host     string `json:"host"`
			Port     string `json:"port"`
			Username string `json:"username"`
			Password string `json:"password"`
			Db       string `json:"db"`
		} `json:"qsdatabase,omitempty"`
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
		Database struct {
			Db       string `json:"db"`
			Host     string `json:"host"`
			Port     string `json:"port"`
			Username string `json:"username"`
			Password string `json:"password"`
		} `json:"database"`
		ContentDatabase struct {
			Db       string `json:"db"`
			Host     string `json:"host"`
			Port     string `json:"port"`
			Username string `json:"username"`
			Password string `json:"password"`
		} `json:"content_database,omitempty"`
		Files struct {
			Opcodes     string `json:"opcodes"`
			MailOpcodes string `json:"mail_opcodes"`
		} `json:"files"`
		Directories struct {
			Patches string `json:"patches"`
			Opcodes string `json:"opcodes"`
		} `json:"directories"`
	} `json:"server"`
	WebAdmin struct {
		Discord struct {
			CrashLogWebhook string `json:"crash_log_webhook,omitempty"`
		} `json:"discord,omitempty"`
		Application struct {
			Key   string `json:"key,omitempty"`
			Admin struct {
				Password string `json:"password,omitempty"`
			} `json:"admin,omitempty"`
		} `json:"application,omitempty"`
		Launcher struct {
			RunSharedMemory  bool `json:"runSharedMemory,omitempty"`
			RunLoginserver   bool `json:"runLoginserver,omitempty"`
			RunQueryServ     bool `json:"runQueryServ,omitempty"`
			IsRunning        bool `json:"isRunning,omitempty"`
			MinZoneProcesses int  `json:"minZoneProcesses,omitempty"`
		} `json:"launcher,omitempty"`
		ServerCodePath string `json:"serverCodePath,omitempty"`
	} `json:"web-admin,omitempty"`
	Spire struct {
		EncryptionKey string `json:"encryption_key,omitempty"`
	} `json:"spire,omitempty"`
}

func (e EQEmuServerConfig) Get() EQEmuConfigJson {
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

func (m *EQEmuServerConfig) debug(msg string, a ...interface{}) {
	if len(os.Getenv("DEBUG")) >= 3 {
		if len(a) > 0 {
			m.logger.Debug("[eqemu_server_config.go] " + fmt.Sprintf(msg, a...) + "\n")
			return
		}
		m.logger.Debug("[eqemu_server_config.go] " + fmt.Sprintf(msg) + "\n")
	}
}

func (e EQEmuServerConfig) Exists() bool {
	return len(e.pathmgmt.GetEQEmuServerConfigFilePath()) > 0
}

func (e EQEmuServerConfig) Save(c EQEmuConfigJson) {
	file, err := json.MarshalIndent(c, "", " ")
	if err != nil {
		e.logger.Error(err)
	}

	err = os.WriteFile(e.pathmgmt.GetEQEmuServerConfigFilePath(), file, 0755)
	if err != nil {
		e.logger.Error(err)
	}
}
