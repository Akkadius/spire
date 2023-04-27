package eqemuserver

import (
	"github.com/manifoldco/promptui"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type InstallConfig struct {
	ServerPath        string `yaml:"server_path"`
	CodePath          string `yaml:"code_path"`
	MysqlUsername     string `yaml:"mysql_username"`
	MysqlPassword     string `yaml:"mysql_password"`
	MysqlDatabaseName string `yaml:"mysql_database_name"`
	MysqlHost         string `yaml:"mysql_host"`
	MysqlPort         string `yaml:"mysql_port"`
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
		Label:   "Server Installation Path (Recommended: ~/server)",
		Default: filepath.Join(homedir, "server"),
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
		Label:   "Server Code Path (Recommended: ~/code)",
		Default: filepath.Join(homedir, "code"),
	}).Run()
	if err != nil {
		a.logger.Fatalf("Prompt failed %v\n", err)
	}
	a.installConfig.CodePath = codePath

	// set the installer path
	a.pathmanager.SetServerPath(serverPath)

	// prompt: mysql username
	useExistingMysqlInstallPrompt, _ := (&promptui.Prompt{
		Label:     "Use an existing MySQL server",
		Default:   "N",
		IsConfirm: true,
	}).Run()

	// check if we are using an existing mysql install
	useExistingMysqlInstall := strings.Contains(strings.ToLower(useExistingMysqlInstallPrompt), "y")
	if useExistingMysqlInstall {
		// TODO: handle
	}

	// if we are installing a new mysql server
	generatedPassword := a.GetRandomPassword()

	// prompt: mysql database name
	mysqlDbName, err := (&promptui.Prompt{
		Label:   "MySQL Database Name (all lowercase, no special characters)",
		Default: "peq",
	}).Run()
	if err != nil {
		a.logger.Fatalf("Prompt failed %v\n", err)
	}
	mysqlDbName = a.stripSpecialCharacters(mysqlDbName)
	a.installConfig.MysqlDatabaseName = mysqlDbName

	// prompt: mysql username
	mysqlUsername, err := (&promptui.Prompt{
		Label:   "MySQL Username",
		Default: "eqemu",
	}).Run()
	if err != nil {
		a.logger.Fatalf("Prompt failed %v\n", err)
	}
	a.installConfig.MysqlUsername = mysqlUsername

	// prompt: mysql password
	mysqlPassword, err := (&promptui.Prompt{
		Label:   "MySQL Password (Leave blank for random password)",
		Default: generatedPassword,
		Mask:    '*',
	}).Run()
	if err != nil {
		a.logger.Fatalf("Prompt failed %v\n", err)
	}

	// validate: passwords match if we manually entered it
	if mysqlPassword != generatedPassword {
		// prompt: mysql password (confirm)
		mysqlPasswordConfirm, err := (&promptui.Prompt{
			Label: "MySQL Password (Confirm)",
			Mask:  '*',
		}).Run()
		if err != nil {
			a.logger.Fatalf("Prompt failed %v\n", err)
		}

		if mysqlPassword != mysqlPasswordConfirm {
			a.logger.Fatalf("MySQL Passwords do not match")
		}
	}

	a.installConfig.MysqlPassword = mysqlPassword

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

	a.logger.Infof("Install config file already exists, loading it")
	// get contents of install config file
	installConfigContents, err := os.ReadFile(installConfigFile)
	if err != nil {
		a.logger.Fatalf("could not read install config file: %v", err)
	}

	a.logger.Infof("----------------------------------------\n")
	a.logger.Infof("%v\n", installConfigFile)
	a.logger.Infof("----------------------------------------\n")
	a.logger.Infof("")

	// print install config contents
	for _, s := range strings.Split(string(installConfigContents), "\n") {
		a.logger.Infof("%v", s)
	}

	a.logger.Infof("----------------------------------------\n")

	// mysql username
	useExistingConfig, _ := (&promptui.Prompt{
		Label:     "Use this configuration",
		Default:   "Y",
		IsConfirm: true,
	}).Run()

	// confirmation check
	if strings.Contains(strings.ToLower(useExistingConfig), "y") || len(useExistingConfig) == 0 {
		a.logger.Infof("Using existing install config")

		// load install config contents into struct
		err = yaml.Unmarshal(installConfigContents, &a.installConfig)
		if err != nil {
			a.logger.Fatalf("could not unmarshal install config: %v", err)
		}

		return true
	}

	return false // config file does not exist
}

func (a *Installer) stripSpecialCharacters(name string) string {
	name = regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(name, "")
	name = strings.ReplaceAll(name, " ", "_")
	return name
}
