package eqemuloginserver

import (
	"encoding/json"
	"github.com/Akkadius/spire/internal/logger"
	"github.com/Akkadius/spire/internal/pathmgmt"
	"os"
	"path/filepath"
)

// Config is the struct type
type Config struct {
	logger   *logger.AppLogger
	pathmgmt *pathmgmt.PathManagement
	config   LoginConfigJson
}

// NewConfig creates a new Config struct
func NewConfig(logger *logger.AppLogger, pathmgmt *pathmgmt.PathManagement) *Config {
	return &Config{
		logger:   logger,
		pathmgmt: pathmgmt,
	}
}

// DatabaseConfig is the struct that represents the database configuration in login.json
type DatabaseConfig struct {
	Db       string `json:"db,omitempty"`
	Host     string `json:"host,omitempty"`
	Port     string `json:"port,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

// LoginConfigJson is the struct that represents the login.json file
type LoginConfigJson struct {
	Database struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		Db       string `json:"db"`
		User     string `json:"user"`
		Password string `json:"password"`
	} `json:"database"`
	Account struct {
		AutoCreateAccounts bool `json:"auto_create_accounts"`
	} `json:"account"`
	Worldservers struct {
		UnregisteredAllowed             bool `json:"unregistered_allowed"`
		ShowPlayerCount                 bool `json:"show_player_count"`
		DevTestServersListBottom        bool `json:"dev_test_servers_list_bottom"`
		SpecialCharacterStartListBottom bool `json:"special_character_start_list_bottom"`
		RejectDuplicateServers          bool `json:"reject_duplicate_servers"`
	} `json:"worldservers"`
	WebAPI struct {
		Enabled bool `json:"enabled"`
		Port    int  `json:"port"`
	} `json:"web_api"`
	Security struct {
		Mode               int  `json:"mode"`
		AllowPasswordLogin bool `json:"allow_password_login"`
		AllowTokenLogin    bool `json:"allow_token_login"`
	} `json:"security"`
	Logging struct {
		Trace          bool `json:"trace"`
		WorldTrace     bool `json:"world_trace"`
		DumpPacketsIn  bool `json:"dump_packets_in"`
		DumpPacketsOut bool `json:"dump_packets_out"`
	} `json:"logging"`
	ClientConfiguration struct {
		TitaniumPort      int    `json:"titanium_port"`
		TitaniumOpcodes   string `json:"titanium_opcodes"`
		SodPort           int    `json:"sod_port"`
		SodOpcodes        string `json:"sod_opcodes"`
		DisplayExpansions bool   `json:"display_expansions"`
		MaxExpansionsMask int    `json:"max_expansions_mask"`
	} `json:"client_configuration"`
}

// Get returns the login config json
func (e *Config) Get() LoginConfigJson {
	cfg := e.pathmgmt.GetEqemuLoginServerConfigPath()
	e.debug("path [%v]", cfg)

	if len(cfg) > 0 {
		e.debug("Reading login config file [%v]", cfg)

		body, err := os.ReadFile(e.pathmgmt.GetEqemuLoginServerConfigPath())
		if err != nil {
			e.logger.Fatal().Err(err).Msg("unable to read login.json file")
		}

		config := LoginConfigJson{}
		err = json.Unmarshal(body, &config)
		if err != nil {
			e.logger.Fatal().Err(err).Any("body", string(body)).Msg("unable to unmarshal login.json file")
		}

		return config
	}

	return LoginConfigJson{}
}

func (e *Config) debug(msg string, a ...interface{}) {
	if len(os.Getenv("DEBUG")) >= 3 {
		e.logger.Debug().Msgf(msg, a...)
	}
}

// Exists will return true if the login.json file exists
func (e *Config) Exists() bool {
	return len(e.pathmgmt.GetEqemuLoginServerConfigPath()) > 0
}

// Save will save the config to the login.json file
func (e *Config) Save(c LoginConfigJson) error {
	file, err := json.MarshalIndent(c, "", " ")
	if err != nil {
		return err
	}

	path := filepath.Join(e.pathmgmt.GetEQEmuServerPath(), "login.json")
	err = os.WriteFile(path, file, 0755)
	if err != nil {
		return err
	}

	return nil
}
