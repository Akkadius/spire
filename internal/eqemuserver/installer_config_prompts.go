package eqemuserver

import (
	"fmt"
	"github.com/Akkadius/spire/internal/promptui"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
)

type InstallConfig struct {
	ServerPath         string `yaml:"server_path"`
	CodePath           string `yaml:"code_path"`
	CompileBinaries    bool   `yaml:"compile_binaries"`
	CompileDevelop     bool   `yaml:"compile_develop"`
	MysqlUsername      string `yaml:"mysql_username"`
	MysqlPassword      string `yaml:"mysql_password"`
	MysqlDatabaseName  string `yaml:"mysql_database_name"`
	MysqlHost          string `yaml:"mysql_host"`
	MysqlPort          string `yaml:"mysql_port"`
	SpireAdminUser     string `yaml:"spire_admin_user"`
	SpireAdminPassword string `yaml:"spire_admin_password"`
	SpireWebPort       string `yaml:"spire_web_port"`
	BotsEnabled        bool   `yaml:"bots_enabled"`
	MercsEnabled       bool   `yaml:"mercs_enabled"`
}

const (
	installConfigFileName = "install_config.yaml"
)

func (a *Installer) checkInstallConfig() {
	a.Banner("Checking Install Config")

	// check if the install config exists
	if a.loadInstallConfigIfExists() {
		return // config file exists, we're done
	}

	homedir, err := os.UserHomeDir()
	if err != nil {
		a.logger.Error(err)
	}

	// prompt: server path
	serverPath, err := (&promptui.Prompt{
		Label:     "Server Installation Path (Recommended: ~/server)",
		Default:   filepath.Join(homedir, "server"),
		AllowEdit: true,
	}).Run()
	if err != nil {
		a.logger.Fatalf("Prompt failed %v\n", err)
	}
	a.installConfig.ServerPath = serverPath

	// create the server path if it doesn't exist
	_ = os.MkdirAll(serverPath, os.ModePerm)

	// set the installer path
	err = os.Chdir(serverPath)
	if err != nil {
		a.logger.Error(err)
	}

	// prompt: server code path
	codePath, err := (&promptui.Prompt{
		Label:     "Server Code Path (Recommended: ~/code)",
		Default:   filepath.Join(homedir, "code"),
		AllowEdit: true,
	}).Run()
	if err != nil {
		a.logger.Fatalf("Prompt failed %v\n", err)
	}
	a.installConfig.CodePath = codePath

	// prompt: compile binaries
	if runtime.GOOS == "linux" {
		compileBinaries, _ := (&promptui.Prompt{
			Label:     "Compile binaries? (Default: N) By default we use pre-compiled binaries. If you plan on developing on your server you might want this enabled.",
			Default:   "N",
			IsConfirm: true,
		}).Run()
		a.installConfig.CompileBinaries = strings.ToLower(compileBinaries) == "y"

		// prompt: compile for development
		if a.installConfig.CompileBinaries {
			compileDevelop, _ := (&promptui.Prompt{
				Label:     "Compile for development? (Default: N) Do not enable this in production. If you want compiles to be faster, enable this",
				Default:   "N",
				IsConfirm: true,
			}).Run()
			a.installConfig.CompileDevelop = strings.ToLower(compileDevelop) == "y"
		}
	}

	// set the installer path
	a.pathmanager.SetServerPath(serverPath)

	// prompt: mysql username
	useExistingMysqlInstallPrompt, _ := (&promptui.Prompt{
		Label:     "Use an existing MySQL server",
		Default:   "N",
		IsConfirm: true,
		AllowEdit: true,
	}).Run()

	// check if we are using an existing mysql install
	useExistingMysqlInstall := strings.Contains(strings.ToLower(useExistingMysqlInstallPrompt), "y")
	if useExistingMysqlInstall {
		fmt.Printf("Using existing MySQL install, please specify your MySQL connection details.")

		// prompt: mysql host
		mysqlHost, err := (&promptui.Prompt{
			Label:     "MySQL Host",
			Default:   "127.0.0.1",
			AllowEdit: true,
		}).Run()
		if err != nil {
			a.logger.Fatalf("Prompt failed %v\n", err)
		}

		// prompt: mysql port
		mysqlPort, err := (&promptui.Prompt{
			Label:     "MySQL Port",
			Default:   "3306",
			AllowEdit: true,
		}).Run()
		if err != nil {
			a.logger.Fatalf("Prompt failed %v\n", err)
		}

		// prompt: mysql database name
		mysqlDbName, err := (&promptui.Prompt{
			Label:     "MySQL Database Name (all lowercase, no special characters) (This will be a new database that doesn't exist already)",
			Default:   "peq",
			AllowEdit: true,
		}).Run()
		if err != nil {
			a.logger.Fatalf("Prompt failed %v\n", err)
		}

		// prompt: mysql username
		mysqlUsername, err := (&promptui.Prompt{
			Label:     "MySQL Username",
			Default:   "eqemu",
			AllowEdit: true,
		}).Run()
		if err != nil {
			a.logger.Fatalf("Prompt failed %v\n", err)
		}

		// prompt: mysql password
		mysqlPassword, err := (&promptui.Prompt{
			Label:     "MySQL Password",
			Default:   "eqemu",
			AllowEdit: true,
		}).Run()
		if err != nil {
			a.logger.Fatalf("Prompt failed %v\n", err)
		}

		// validate the mysql connection
		fmt.Printf("Validating MySQL connection...")
		err = a.validateMysqlConnection(
			mysqlHost,
			mysqlPort,
			mysqlDbName,
			mysqlUsername,
			mysqlPassword,
		)
		if err != nil {
			a.logger.Fatalf("Failed to validate MySQL connection: %v", err)
		}

		// set installation variables
		a.installConfig.MysqlHost = mysqlHost
		a.installConfig.MysqlPort = mysqlPort
		a.installConfig.MysqlDatabaseName = mysqlDbName
		a.installConfig.MysqlUsername = mysqlUsername
		a.installConfig.MysqlPassword = mysqlPassword
	} else {
		// if we are installing a new mysql server
		generatedPassword := a.GetRandomPassword()

		// prompt: mysql database name
		mysqlDbName, err := (&promptui.Prompt{
			Label:     "MySQL Database Name (all lowercase, no special characters)",
			Default:   "peq",
			AllowEdit: true,
		}).Run()
		if err != nil {
			a.logger.Fatalf("Prompt failed %v\n", err)
		}
		mysqlDbName = a.stripSpecialCharacters(mysqlDbName)
		a.installConfig.MysqlDatabaseName = mysqlDbName

		// prompt: mysql username
		mysqlUsername, err := (&promptui.Prompt{
			Label:     "MySQL Username",
			Default:   "eqemu",
			AllowEdit: true,
		}).Run()
		if err != nil {
			a.logger.Fatalf("Prompt failed %v\n", err)
		}
		a.installConfig.MysqlUsername = mysqlUsername

		// prompt: mysql password
		mysqlPassword, err := (&promptui.Prompt{
			Label:     "MySQL Password (Leave blank for random password)",
			Default:   generatedPassword,
			Mask:      '*',
			AllowEdit: true,
		}).Run()
		if err != nil {
			a.logger.Fatalf("Prompt failed %v\n", err)
		}

		// validate: passwords match if we manually entered it
		if mysqlPassword != generatedPassword {
			// prompt: mysql password (confirm)
			mysqlPasswordConfirm, err := (&promptui.Prompt{
				Label:     "MySQL Password (Confirm)",
				Mask:      '*',
				AllowEdit: true,
			}).Run()
			if err != nil {
				a.logger.Fatalf("Prompt failed %v\n", err)
			}

			if mysqlPassword != mysqlPasswordConfirm {
				a.logger.Fatalf("MySQL Passwords do not match")
			}
		}

		a.installConfig.MysqlPassword = mysqlPassword

		a.installConfig.MysqlHost = "127.0.0.1"
		a.installConfig.MysqlPort = "3306"
	}

	// prompt: spire admin user
	spireAdminUser, err := (&promptui.Prompt{
		Label:     "Spire Admin User",
		Default:   "admin",
		AllowEdit: true,
	}).Run()
	if err != nil {
		a.logger.Fatalf("Prompt failed %v\n", err)
	}
	a.installConfig.SpireAdminUser = spireAdminUser

	// prompt: spire admin password
	spireAdminPassword, err := (&promptui.Prompt{
		Label:     "Spire Admin Password (Leave blank for random password)",
		Default:   a.GetRandomPassword(),
		Mask:      '*',
		AllowEdit: true,
	}).Run()
	if err != nil {
		a.logger.Fatalf("Prompt failed %v\n", err)
	}
	a.installConfig.SpireAdminPassword = spireAdminPassword

	// prompt: spire web port
	spireWebPort, err := (&promptui.Prompt{
		Label:     "Spire Web Port (Server Admin UI, Editing Tools, Default: 3007)",
		Default:   "3007",
		AllowEdit: true,
	}).Run()
	if err != nil {
		a.logger.Fatalf("Prompt failed %v\n", err)
	}
	a.installConfig.SpireWebPort = spireWebPort

	// prompt: bots enabled
	botsEnabled, _ := (&promptui.Prompt{
		Label:     "Enable bots? (Default: N) Want to play on your server solo? You can build a group and/or raid of bots to play with! (These are not mercenaries)",
		Default:   "N",
		IsConfirm: true,
	}).Run()
	a.installConfig.BotsEnabled = strings.ToLower(botsEnabled) == "y"

	// prompt: mercs enabled
	mercsEnabled, _ := (&promptui.Prompt{
		Label:     "Enable mercenaries? (Default: N) Mercenaries (BETA) are NPC's that you can hire to help you in your adventures. They are not bots.",
		Default:   "N",
		IsConfirm: true,
	}).Run()
	a.installConfig.MercsEnabled = strings.ToLower(mercsEnabled) == "y"

	// write a.installConfig to yaml config file called install_config.yaml
	// marshal a.installConfig into yaml
	installConfigYaml, err := yaml.Marshal(a.installConfig)
	if err != nil {
		a.logger.Fatalf("error: %v", err)
	}

	// write yaml to install_config.yaml
	installConfigFile := filepath.Join(a.pathmanager.GetEQEmuServerPath(), installConfigFileName)
	err = os.WriteFile(installConfigFile, installConfigYaml, 0644)
	if err != nil {
		a.logger.Fatalf("could not write install_config.yaml: %v", err)
	}

	a.DoneBanner("Checking Install Config")
}

func (a *Installer) loadInstallConfigIfExists() bool {
	// check if install config file exists
	installConfigFile := filepath.Join(a.pathmanager.GetEQEmuServerPath(), installConfigFileName)
	if _, err := os.Stat(installConfigFile); os.IsNotExist(err) {

		// try to load from home directory default
		homedir, err := os.UserHomeDir()
		if err != nil {
			a.logger.Error(err)
		}

		file := filepath.Join(homedir, "server", installConfigFileName)
		if _, err := os.Stat(file); os.IsNotExist(err) {
			return false // config file does not exist
		}

		// set the install config file
		installConfigFile = file
	}

	fmt.Printf("Install config file already exists, loading it\n")
	// get contents of install config file
	installConfigContents, err := os.ReadFile(installConfigFile)
	if err != nil {
		a.logger.Fatalf("could not read install config file: %v", err)
	}

	fmt.Printf("----------------------------------------\n")
	fmt.Printf("%v\n", installConfigFile)
	fmt.Printf("----------------------------------------\n")
	fmt.Printf("\n")

	// print install config contents
	for _, s := range strings.Split(string(installConfigContents), "\n") {
		fmt.Printf("%v\n", s)
	}

	fmt.Printf("----------------------------------------\n")

	// mysql username
	useExistingConfig, _ := (&promptui.Prompt{
		Label:     "Use this configuration",
		Default:   "Y",
		IsConfirm: true,
	}).Run()

	// confirmation check
	if strings.Contains(strings.ToLower(useExistingConfig), "y") || len(useExistingConfig) == 0 {
		fmt.Printf("Using existing install config")

		// load install config contents into struct
		err = yaml.Unmarshal(installConfigContents, &a.installConfig)
		if err != nil {
			a.logger.Fatalf("could not unmarshal install config: %v", err)
		}

		a.pathmanager.SetServerPath(a.installConfig.ServerPath)

		return true
	}

	return false // config file does not exist
}

func (a *Installer) stripSpecialCharacters(name string) string {
	name = regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(name, "")
	name = strings.ReplaceAll(name, " ", "_")
	return name
}

// validateMysqlConnection validates the mysql connection
func (a *Installer) validateMysqlConnection(host string, port string, name string, username string, mysqlPassword string) error {
	mysqlAdminBin := "mysqladmin"
	if runtime.GOOS == "windows" {
		mysqlAdminBin = filepath.Join(a.getWindowsMysqlPath(), "mysqladmin.exe")
	}

	mysqlBin := "mysql"
	if runtime.GOOS == "windows" {
		mysqlBin = filepath.Join(a.getWindowsMysqlPath(), "mysql.exe")
	}

	// check if mysql is running
	res := a.Exec(ExecConfig{
		command: mysqlAdminBin,
		args: []string{
			"ping",
			"-h" + host,
			"-P" + port,
			"-u" + username,
			"-p" + mysqlPassword,
		},
	})
	if !strings.Contains(res, "mysqld is alive") {
		return fmt.Errorf("could not connect to mysql: %v", res)
	}

	// check if database exists and has data
	res = a.Exec(ExecConfig{
		command: mysqlBin,
		args: []string{
			"-h" + host,
			"-P" + port,
			"-u" + username,
			"-p" + mysqlPassword,
			"-e",
			"show databases like '" + name + "';",
		},
	})
	if len(res) > 0 {
		return fmt.Errorf("database already exists: %v", res)
	}

	return nil
}
