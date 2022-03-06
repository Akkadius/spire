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
		} `json:"zones"`
		Qsdatabase struct {
			Host     string `json:"host"`
			Port     string `json:"port"`
			Username string `json:"username"`
			Password string `json:"password"`
			Db       string `json:"db"`
		} `json:"qsdatabase"`
		Chatserver struct {
			Port string `json:"port"`
			Host string `json:"host"`
		} `json:"chatserver"`
		Mailserver struct {
			Host string `json:"host"`
			Port string `json:"port"`
		} `json:"mailserver"`
		World struct {
			Loginserver1 struct {
				Account  string `json:"account"`
				Password string `json:"password"`
				Legacy   string `json:"legacy"`
				Host     string `json:"host"`
				Port     string `json:"port"`
			} `json:"loginserver1"`
			Loginserver2 struct {
				Port     string `json:"port"`
				Account  string `json:"account"`
				Password string `json:"password"`
				Host     string `json:"host"`
			} `json:"loginserver2"`
			TCP struct {
				IP   string `json:"ip"`
				Port string `json:"port"`
			} `json:"tcp"`
			Telnet struct {
				IP      string `json:"ip"`
				Port    string `json:"port"`
				Enabled string `json:"enabled"`
			} `json:"telnet"`
			Key       string `json:"key"`
			Shortname string `json:"shortname"`
			Longname  string `json:"longname"`
		} `json:"world"`
		Database struct {
			Db       string `json:"db"`
			Host     string `json:"host"`
			Port     string `json:"port"`
			Username string `json:"username"`
			Password string `json:"password"`
		} `json:"database"`
		Files struct {
			Opcodes     string `json:"opcodes"`
			MailOpcodes string `json:"mail_opcodes"`
		} `json:"files"`
		Directories struct {
			Patches string `json:"patches"`
			Opcodes string `json:"opcodes"`
		} `json:"directories"`
	} `json:"server"`
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
	if len(os.Getenv("DEBUG")) > 0 {
		if len(a) > 0 {
			m.logger.Debug("[eqemu_server_config.go] " + fmt.Sprintf(msg, a...) + "\n")
			return
		}
		m.logger.Debug("[eqemu_server_config.go] " + fmt.Sprintf(msg) + "\n")
	}
}
