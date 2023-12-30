package eqemuserver

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Akkadius/spire/internal/download"
	"github.com/Akkadius/spire/internal/eqemuloginserver"
	"github.com/Akkadius/spire/internal/eqemuserverconfig"
	"github.com/Akkadius/spire/internal/password"
	"github.com/Akkadius/spire/internal/pathmgmt"
	"github.com/Akkadius/spire/internal/unzip"
	"github.com/go-git/go-git/v5"
	"github.com/google/go-github/v41/github"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/process"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type Installer struct {
	pathmanager   *pathmgmt.PathManagement
	config        *eqemuserverconfig.Config
	loginConfig   *eqemuloginserver.Config
	logger        *logrus.Logger
	stepTime      time.Time
	totalTime     time.Time
	installConfig *InstallConfig
}

func getLogger() *logrus.Logger {
	l := logrus.New()
	l.SetFormatter(&logrus.TextFormatter{
		DisableTimestamp:       true,
		ForceColors:            true,
		DisableLevelTruncation: false,
		PadLevelText:           true,
	})

	// base level
	l.SetLevel(logrus.InfoLevel)

	// debug logging
	if len(os.Getenv("DEBUG")) > 0 {
		l.SetLevel(logrus.DebugLevel)
	}

	return l
}

func NewInstaller() *Installer {
	// TODO: Clean this up
	logger := getLogger()
	pathmanager := pathmgmt.NewPathManagement(logger)
	i := &Installer{
		logger:        logger,
		pathmanager:   pathmanager,
		config:        eqemuserverconfig.NewConfig(logger, pathmanager),
		installConfig: &InstallConfig{},
		loginConfig:   eqemuloginserver.NewConfig(logger, pathmanager),
	}

	return i
}

type Task struct {
	function func() error
}

func checkForWinRanAsAdmin() bool {
	_, err := os.Open("\\\\.\\PHYSICALDRIVE0")
	if err != nil {
		return false
	}
	return true
}

func (a *Installer) Install() error {
	err := a.setInstallerPath()
	if err != nil {
		return err
	}

	fmt.Println(`
------------------------------------------------------------------------------------------
|                 > EverQuest Emulator Server Installer < Author @Akkadius               |
------------------------------------------------------------------------------------------
| > EverQuest Emulator • A Fan Based, Open-Sourced Emulation Project                     |
------------------------------------------------------------------------------------------
| • Developed and maintained by the EQEmulator Development team                          |
------------------------------------------------------------------------------------------
| • GitHub https://github.com/eqemu/server                                               |
------------------------------------------------------------------------------------------
| • Everquest is a registered trademark Daybreak Game Company LLC.                       |
------------------------------------------------------------------------------------------
| • EQEmulator is not associated or affiliated in any way with Daybreak Game Company LLC.|
------------------------------------------------------------------------------------------
| > To be installed <
------------------------------------------------------------------------------------------
| • Server running folder • Will be installed to the folder you ran this script
| • MariaDB x64 (MySQL) • Database engine
| • Heidi SQL (Comes with MariaDB)
| • Perl x64 5.24.4.1 • Scripting language for quest engines
| • LUA Configured • Scripting language for quest engines
| • Latest PEQ Database
| • Latest PEQ Quests
| • Latest Plugins repository
| • Automatically added Firewall rules
| • Maps (Latest V2) formats are loaded
| • New Path files are loaded
| • Optimized server binaries
------------------------------------------------------------------------------------------
| > This installer will walk you through the installation of the EQEmu Server
------------------------------------------------------------------------------------------`)

	if runtime.GOOS == "windows" {
		_, _ = a.Exec(ExecConfig{command: "mode", args: []string{"800"}})
		if !checkForWinRanAsAdmin() {
			return fmt.Errorf("please run this installer as Administrator")
		}
	}

	// check install config
	err = a.checkInstallConfig()
	if err != nil {
		return err
	}

	// install debian packages
	// install ubuntu packages (for ubuntu)
	a.totalTime = time.Now()
	if runtime.GOOS == "linux" {
		for _, t := range []func() error{
			a.installLinuxOsPackages,
			a.initLinuxMysql,
		} {
			err := t()
			if err != nil {
				return err
			}
		}
	}

	// install windows packages
	if runtime.GOOS == "windows" {
		for _, t := range []func() error{
			a.initWindowsCommandPrompt,
			a.initWindowsPerl,
			a.initWindowsMysql,
		} {
			err := t()
			if err != nil {
				return err
			}
		}
	}

	// run server install tasks
	for _, t := range []func() error{
		a.initializeDirectories,
		a.cloneEQEmuSource,
		a.initializeServerConfig,
		a.cloneEQEmuMaps,
		a.clonePeqQuests,
		a.sourcePeqDatabase,
		a.symlinkPatchFiles,
		a.symlinkOpcodeFiles,
		a.symlinkLoginOpcodeFiles,
		a.symlinkPluginsAndModules,
	} {
		err := t()
		if err != nil {
			return err
		}
	}

	// server binaries
	if runtime.GOOS == "linux" && a.installConfig.CompileBinaries {
		err := a.compileBinaries()
		if err != nil {
			return err
		}
	} else {
		err := a.installBinaries()
		if err != nil {
			return err
		}
	}

	if runtime.GOOS == "linux" {
		for _, t := range []func() error{
			a.createLinuxServerScripts,
			a.injectSpireStartCronJob,
		} {
			err := t()
			if err != nil {
				return err
			}
		}
	}
	if runtime.GOOS == "windows" {
		for _, t := range []func() error{
			a.disableQuickEdit,
			a.createWindowsServerScripts,
			a.initWindowsWget, // TODO: Remove this when perl utility script is deprecated from world
			a.setWindowsPerlPath,
			a.setWindowsMysqlPath,
			a.publishWindowsPorts,
		} {
			err := t()
			if err != nil {
				return err
			}
		}
	}

	for _, t := range []func() error{
		a.installSpireBinary,
		a.initSpire,
		a.startSpire,
		a.setPostInstallConfigValues,
		a.runSharedMemory,
		a.runWorldForDatabaseUpdates,
		a.runZoneForDataInjections,
		a.enableBots,
		a.enableMercenaries,
		a.initLoginServer,
	} {
		err := t()
		if err != nil {
			return err
		}
	}

	checkmark := "✅"
	if runtime.GOOS == "windows" {
		checkmark = "√"
	}

	fmt.Println("")
	fmt.Println("----------------------------------------")
	fmt.Printf("| %s | Installation Complete (%v)\n", checkmark, FormatDuration(time.Since(a.totalTime)))
	fmt.Println("----------------------------------------")

	if runtime.GOOS == "windows" {
		a.openWindowsPostInstallWindows()
	}

	fmt.Print("Press [Enter] to close...")
	_, _ = bufio.NewReader(os.Stdin).ReadBytes('\n')

	return nil
}

func (a *Installer) installLinuxOsPackages() error {
	a.Banner("Installing OS Packages")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	// apt-get update
	params := []string{"bash", "-c", "apt-get update && apt-get install -yq apt-utils; echo 'debconf debconf/frontend select Noninteractive' | debconf-set-selections"}
	cmd := exec.CommandContext(ctx, "sudo", params...)
	cmd.Env = os.Environ()
	cmd.Dir = a.pathmanager.GetEQEmuServerPath()
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("could not get stdout pipe: %v", err)
	}
	cmd.Stderr = cmd.Stdout
	err = cmd.Start()
	if err != nil {
		return err
	}

	merged := io.MultiReader(stdout)
	scanner := bufio.NewScanner(merged)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	// install packages
	var packages []string
	distro, _ := a.getLinuxDistribution()
	version, _ := a.getLinuxDistributionVersion()

	fmt.Println("----------------------------------------")
	fmt.Printf("| > Installing packages for distro [%v] version [%v]\n", distro, version)
	fmt.Println("----------------------------------------")

	if distro == "ubuntu" {
		packages = getUbuntuPackages()
	} else if distro == "debian" && version >= 11 {
		packages = getDebian11Packages()
	} else {
		return fmt.Errorf("Unsupported distribution: %v", distro)
	}

	// apt-get install -y
	params = []string{"apt-get", "install", "-yq", "-m", "--no-install-recommends"}
	params = append(params, packages...)
	cmd = exec.CommandContext(ctx, "sudo", params...)
	cmd.Env = os.Environ()

	cmd.Dir = a.pathmanager.GetEQEmuServerPath()
	stdout, err = cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("could not get stdout pipe: %v", err)
	}

	cmd.Stderr = cmd.Stdout

	err = cmd.Start()
	if err != nil {
		return err
	}

	merged = io.MultiReader(stdout)
	scanner = bufio.NewScanner(merged)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	a.DoneBanner("Installing OS Packages")

	return nil
}

func (a *Installer) Banner(s string) {
	fmt.Println("")
	fmt.Println("------------------------------------------------------------------------------------------")
	fmt.Printf("| > %v\n", s)
	fmt.Println("------------------------------------------------------------------------------------------")
	//fmt.Println("")
	a.stepTime = time.Now()
}

func FormatDuration(d time.Duration) string {
	scale := 100 * time.Second
	// look for the max scale that is smaller than d
	for scale > d {
		scale = scale / 10
	}
	return d.Round(scale / 100).String()
}

func (a *Installer) DoneBanner(s string) {
	checkmark := "✅"
	if runtime.GOOS == "windows" {
		checkmark = "√"
	}

	fmt.Println("----------------------------------------")
	fmt.Printf("| %v | %v (%v)\n", checkmark, s, FormatDuration(time.Since(a.stepTime)))
	fmt.Println("----------------------------------------")
}

func (a *Installer) initializeDirectories() error {
	a.Banner("Initializing Directories")

	directories := []string{
		"bin",
		"assets",
		"assets/opcodes",
		"assets/patches",
		"maps",
		"mods",
		"quests",
		"shared",
		"logs",
	}

	for _, dir := range directories {
		fmt.Printf("Creating directory [%v]\n", dir)
		p := filepath.Join(a.pathmanager.GetEQEmuServerPath(), dir)
		err := os.MkdirAll(p, os.ModePerm)
		if err != nil {
			return fmt.Errorf("could not create directory: %v", err)
		}
	}

	a.DoneBanner("Initializing Directories")

	return nil
}

func (a *Installer) cloneEQEmuMaps() error {
	a.Banner("Initializing Server Maps")

	err := a.checkIfMapsAreUpToDate()
	if err == nil {
		a.DoneBanner("Initializing Server Maps")
		return nil
	}

	// zip file path
	dumpZip := filepath.Join(os.TempDir(), "/maps.zip")

	// download the zip file
	err = download.WithProgress(
		dumpZip,
		"https://github.com/Akkadius/eqemu-maps/releases/latest/download/maps.zip",
	)
	if err != nil {
		return err
	}

	// unzip the file
	mapsPath := filepath.Join(a.pathmanager.GetEQEmuServerPath(), "maps")
	fmt.Printf("Downloaded zip to [%v]\n", dumpZip)
	err = unzip.New(dumpZip, mapsPath).Extract()
	if err != nil {
		return fmt.Errorf("could not extract zip: %v", err)
	}

	// remove the zip file
	err = os.Remove(dumpZip)
	if err != nil {
		return fmt.Errorf("could not remove zip: %v", err)
	}

	a.DoneBanner("Initializing Server Maps")
	return nil
}

func (a *Installer) clonePeqQuests() error {
	a.Banner("Initializing Server Quests")

	fmt.Printf("Cloning Quests from github.com/ProjectEQ/projecteqquests.git\n")

	// clone the repository
	repoPath := filepath.Join(a.pathmanager.GetEQEmuServerPath(), "quests")
	_, err := git.PlainClone(repoPath, false, &git.CloneOptions{
		URL:      "https://github.com/ProjectEQ/projecteqquests.git",
		Progress: a.logger.Writer(),
	})

	if err != nil && !errors.Is(err, git.ErrRepositoryAlreadyExists) {
		return fmt.Errorf("Could not clone quests [%v]\n", err)
	}

	// if the repository already exists, update it instead
	if errors.Is(err, git.ErrRepositoryAlreadyExists) {
		fmt.Printf("Quest repo already exists, skipping clone and updating instead\n")

		// open the repository
		r, err := git.PlainOpen(repoPath)
		if err != nil {
			return fmt.Errorf("could not open repository: %v", err)
		}

		// Get the working directory for the repository
		w, err := r.Worktree()
		if err != nil {
			return fmt.Errorf("could not get worktree: %v", err)
		}

		// Pull the latest changes from the origin remote and merge into the current branch
		err = w.Pull(&git.PullOptions{RemoteName: "origin", Progress: a.logger.Writer()})
		if err != nil && !errors.Is(err, git.NoErrAlreadyUpToDate) {
			return fmt.Errorf("could not pull: %v", err)
		}

		fmt.Printf("Quests updated successfully!\n")
	}

	a.DoneBanner("Initializing Server Quests")
	return nil
}

type MysqlConfig struct {
	DatabaseName     string `json:"database_name"`
	DatabaseUser     string `json:"database_user"`
	DatabasePassword string `json:"database_password"`
	RootPassword     string `json:"root_password"`
}

func (a *Installer) initLinuxMysql() error {
	a.Banner("Initializing MySQL")

	// change root password
	_, err := a.Exec(ExecConfig{command: "sudo", args: []string{"apt", "install", "-y", "-m", "mariadb-server"}})
	if err != nil {
		return err
	}
	_, err = a.Exec(ExecConfig{command: "sudo", args: []string{"pkill", "-f", "-9", "mysql"}})
	if err != nil {
		return err
	}

	time.Sleep(1 * time.Second)
	go func() {
		_, _ = a.Exec(ExecConfig{command: "sudo", args: []string{"mysqld_safe", "--skip-grant-tables", "--skip-networking"}})
	}()
	time.Sleep(1 * time.Second)

	c := MysqlConfig{
		DatabaseName:     a.installConfig.MysqlDatabaseName,
		DatabaseUser:     a.installConfig.MysqlUsername,
		DatabasePassword: a.installConfig.MysqlPassword,
	}

	// create a new database
	var sql string

	fmt.Printf("Creating database [%v]\n", c.DatabaseName)
	err = a.DbExec(DbExecConfig{statement: fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %v", c.DatabaseName)})
	if err != nil {
		return err
	}

	// create a new user
	fmt.Printf("Creating user [%v]\n", c.DatabaseUser)

	// grant privileges to the new user
	fmt.Printf("Granting privileges to user [%v]\n", c.DatabaseUser)
	sql += fmt.Sprintf("CREATE USER IF NOT EXISTS '%v'@'localhost' IDENTIFIED BY '%v'; ", c.DatabaseUser, c.DatabasePassword)
	sql += fmt.Sprintf("GRANT ALL PRIVILEGES ON %v.* TO '%v'@'localhost'", c.DatabaseName, c.DatabaseUser)

	// flush privileges
	fmt.Println("Flushing privileges")
	err = a.DbExec(DbExecConfig{statement: fmt.Sprintf("FLUSH PRIVILEGES; %v; FLUSH PRIVILEGES;", sql), hidestring: c.DatabasePassword})
	if err != nil {
		return err
	}

	_, err = a.Exec(ExecConfig{command: "sudo", args: []string{"pkill", "-f", "-9", "mysql"}})
	if err != nil {
		return err
	}
	_, err = a.Exec(ExecConfig{command: "sudo", args: []string{"service", "mariadb", "start"}})
	if err != nil {
		return err
	}

	a.DoneBanner("Initializing MySQL")

	return nil
}

type DbExecConfig struct {
	statement  string
	hidestring string
}

func (a *Installer) DbExec(c DbExecConfig) error {
	mysqlPath := "mysql"
	if runtime.GOOS == "windows" {
		mysqlPath = filepath.Join(a.getWindowsMysqlPath(), "mysql.exe")
	}

	// check if mysql is installed
	if len(mysqlPath) == 0 {
		return fmt.Errorf("could not find mysql executable")
	}

	_, err := a.Exec(
		ExecConfig{
			command:    mysqlPath,
			args:       []string{"-uroot", "-e", fmt.Sprintf("%v", c.statement)},
			hidestring: c.hidestring,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

type ExecConfig struct {
	command     string   // command to run
	args        []string // arguments to pass to the command
	hidestring  string   // string to hide from the output
	dieonoutput string   // string to die on if found in the output
	execpath    string   // path to execute the command in
	detach      bool     // detach the process
	silent      bool     // silent execution
}

func (a *Installer) Exec(c ExecConfig) (string, error) {
	// create a new context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	// create the command
	cmd := exec.CommandContext(ctx, c.command, c.args...)
	cmd.Env = os.Environ()

	cmd.Dir = a.pathmanager.GetEQEmuServerPath()

	// if we have an execpath, use it
	if c.execpath != "" {
		cmd.Dir = c.execpath
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", fmt.Errorf("could not get stdout pipe: %v", err)
	}

	// hide the password from the output
	argsPrint := strings.TrimSpace(strings.Join(c.args, " "))
	if len(c.hidestring) > 0 {
		hide := c.hidestring
		argsPrint = strings.ReplaceAll(argsPrint, hide, "********")
	}
	if len(argsPrint) > 0 {
		argsPrint = " " + argsPrint
	}

	if !c.silent {
		fmt.Printf("Running command [%v%v]\n", c.command, argsPrint)
	}

	cmd.Stderr = cmd.Stdout
	err = cmd.Start()
	if err != nil {
		return "", err
	}

	// if we are detaching, release the process
	if c.detach {
		fmt.Printf("Detaching process [%v%v]\n", c.command, argsPrint)
		err = cmd.Process.Release()
		if err != nil {
			return "", err
		}
		return "", nil
	}

	// merge stdout and stderr
	merged := io.MultiReader(stdout)
	scanner := bufio.NewScanner(merged)
	output := ""
	for scanner.Scan() {
		if len(c.dieonoutput) > 0 {
			if strings.Contains(scanner.Text(), c.dieonoutput) {
				fmt.Printf("Found [%v] in output, exiting process\n", c.dieonoutput)
				_ = cmd.Process.Kill()
				break
			}
		}

		// don't print the output if we are silent
		// still return
		if !c.silent {
			fmt.Println(scanner.Text())
		}
		output += scanner.Text() + "\n"
	}

	return output, nil
}

func (a *Installer) cloneEQEmuSource() error {
	a.Banner("Initializing Server Source")

	fmt.Printf("Cloning from https://github.com/EQEmu/Server.git\n")

	// clone the repository
	repoPath := a.installConfig.CodePath
	_, err := git.PlainClone(repoPath, false, &git.CloneOptions{
		URL:               "https://github.com/EQEmu/Server.git",
		Progress:          a.logger.Writer(),
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	})

	// if the repository already exists, update it instead
	if err != nil && err != git.ErrRepositoryAlreadyExists {
		return fmt.Errorf("Could not clone quests [%v]\n", err)
	}

	// if the repository already exists, update it instead
	if errors.Is(err, git.ErrRepositoryAlreadyExists) {
		fmt.Printf("repo already exists, skipping clone and updating instead\n")

		// open the repository
		r, err := git.PlainOpen(repoPath)
		if err != nil {
			return fmt.Errorf("could not open repository: %v", err)
		}

		// Get the working directory for the repository
		w, err := r.Worktree()
		if err != nil {
			return fmt.Errorf("could not get worktree: %v", err)
		}

		// Pull the latest changes from the origin remote and merge into the current branch
		err = w.Pull(&git.PullOptions{RemoteName: "origin", Progress: a.logger.Writer()})
		if err != nil && err != git.NoErrAlreadyUpToDate {
			return fmt.Errorf("could not pull: %v", err)
		}

		fmt.Printf("repo updated successfully!\n")
	}

	a.DoneBanner("Initializing Server Source")
	return nil
}

func (a *Installer) initializeServerConfig() error {
	a.Banner("Initializing Server Config")

	// check if eqemu_config.json exists
	if _, err := os.Stat(a.pathmanager.GetEQEmuServerConfigFilePath()); err == nil {
		fmt.Printf("eqemu_config.json already exists, skipping\n")
		return nil
	}

	// download the config file
	res, err := http.Get("https://raw.githubusercontent.com/Akkadius/eqemu-install-v2/master/eqemu_config_docker.json")
	if err != nil {
		return err
	}

	// read the response body
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	// unmarshal the json
	var c eqemuserverconfig.EQEmuConfigJson
	err = json.Unmarshal(b, &c)
	if err != nil {
		return err
	}

	// set the installation config
	c.Server.Database.Host = a.installConfig.MysqlHost
	c.Server.Database.Port = a.installConfig.MysqlPort
	c.Server.Database.Username = a.installConfig.MysqlUsername
	c.Server.Database.Password = a.installConfig.MysqlPassword
	c.Server.Database.Db = a.installConfig.MysqlDatabaseName

	// convert a.installConfig.SpireWebPort to int
	spireWebPort, err := strconv.Atoi(a.installConfig.SpireWebPort)
	if err != nil {
		return err
	}
	c.Spire.HttpPort = spireWebPort

	// save the config file
	err = a.config.Save(c)
	if err != nil {
		return err
	}

	fmt.Printf("Saved config to [%v]\n", a.pathmanager.GetEQEmuServerConfigFilePath())

	a.DoneBanner("Initializing Server Config")
	return nil
}

func (a *Installer) sourcePeqDatabase() error {
	a.Banner("Sourcing ProjectEQ Database")

	mysqlPath := "mysql"
	if runtime.GOOS == "windows" {
		mysqlPath = filepath.Join(a.getWindowsMysqlPath(), "mysql.exe")
	}

	tables, err := a.Exec(
		ExecConfig{
			command: mysqlPath,
			args: []string{
				fmt.Sprintf("-u%v", a.installConfig.MysqlUsername),
				fmt.Sprintf("-p%v", a.installConfig.MysqlPassword),
				a.installConfig.MysqlDatabaseName,
				"-e",
				"show tables",
			},
			hidestring: a.installConfig.MysqlPassword,
		},
	)
	if err != nil {
		return err
	}

	// get the table count
	// subtract 1 because it is the output header
	tableCount := len(strings.Split(tables, "\n")) - 1

	fmt.Printf(
		"Database [%v] has [%v] tables\n",
		a.installConfig.MysqlDatabaseName,
		tableCount,
	)

	if len(strings.Split(tables, "\n")) > 200 {
		// database already exists, skip
		fmt.Printf(
			"Database [%v] already exists with [%v] tables, skipping source\n",
			a.installConfig.MysqlDatabaseName,
			tableCount,
		)
		return nil
	}

	// zip file path
	dumpZip := filepath.Join(os.TempDir(), "/dump/peq.zip")

	// Create the temp folder
	fmt.Printf("Creating directory [%v]\n", filepath.Dir(dumpZip))
	err = os.MkdirAll(filepath.Dir(dumpZip), os.ModePerm)
	if err != nil {
		return fmt.Errorf("could not create directory: %v", err)
	}

	// download the latest database dump
	err = download.WithProgress(
		dumpZip,
		"http://db.projecteq.net/api/v1/dump/latest",
	)
	if err != nil {
		return err
	}

	fmt.Printf("Downloaded zip to [%v]\n", dumpZip)
	err = unzip.New(dumpZip, filepath.Join(os.TempDir(), "/dump")).Extract()
	if err != nil {
		return fmt.Errorf("could not extract zip: %v", err)
	}

	fmt.Printf("Extracted zip to [%v]\n", filepath.Join(os.TempDir(), "/dump/peq-dump"))

	extractPath := filepath.Join(os.TempDir(), "/dump/peq-dump")

	fmt.Printf("Sourcing database dump from [%v]\n", extractPath)

	_, err = a.Exec(
		ExecConfig{
			command: mysqlPath,
			args: []string{
				fmt.Sprintf("-u%v", a.installConfig.MysqlUsername),
				fmt.Sprintf("-p%v", a.installConfig.MysqlPassword),
				a.installConfig.MysqlDatabaseName,
				"-e",
				"source create_all_tables.sql",
			},
			hidestring: a.installConfig.MysqlPassword,
			execpath:   extractPath,
		},
	)
	if err != nil {
		return err
	}

	fmt.Printf("Sourced database dump from [%v]\n", extractPath)

	_, err = a.Exec(
		ExecConfig{
			command: mysqlPath,
			args: []string{
				fmt.Sprintf("-u%v", a.installConfig.MysqlUsername),
				fmt.Sprintf("-p%v", a.installConfig.MysqlPassword),
				a.installConfig.MysqlDatabaseName,
				"-e",
				"show tables",
			},
			hidestring: a.installConfig.MysqlPassword,
		},
	)
	if err != nil {
		return err
	}

	// cleanup the temp folder
	fmt.Printf("|-- Cleaning up temp folder [%v]\n", filepath.Join(os.TempDir(), "/dump"))
	err = os.RemoveAll(filepath.Join(os.TempDir(), "/dump"))
	if err != nil {
		return fmt.Errorf("could not remove directory: %v", err)
	}

	a.DoneBanner("Sourcing ProjectEQ Database")
	return nil
}

func (a *Installer) installBinaries() error {
	a.Banner("Installing EQEmu Server Binaries")

	// download the latest binaries
	tempPath := filepath.Join(os.TempDir(), "eqemu-server.zip")
	fmt.Printf("Downloading binaries to [%v]\n", tempPath)
	err := download.WithProgress(
		tempPath,
		fmt.Sprintf("https://github.com/eqemu/server/releases/latest/download/eqemu-server-%v-x64.zip", runtime.GOOS),
	)
	if err != nil {
		return err
	}

	// extract the zip
	extractTo := a.pathmanager.GetEQEmuServerBinPath()
	fmt.Printf("Extracting zip to [%v]\n", extractTo)
	err = unzip.New(tempPath, extractTo).Extract()
	if err != nil {
		return fmt.Errorf("could not extract zip: %v", err)
	}

	// cleanup the temp folder
	fmt.Printf("|-- Cleaning up temp folder [%v]\n", tempPath)
	err = os.RemoveAll(tempPath)
	if err != nil {
		return fmt.Errorf("could not remove directory: %v", err)
	}

	// make the binaries executable
	fmt.Printf("Making binaries executable\n")

	// loop through files in the bin folder
	err = filepath.Walk(extractTo, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// skip directories
		if info.IsDir() {
			return nil
		}

		fmt.Printf("|-- Making [%v] executable\n", path)

		// make the file executable
		err = os.Chmod(path, 0755)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("could not make binaries executable: %v", err)
	}

	a.DoneBanner("Installing EQEmu Server Binaries")
	return nil
}

func (a *Installer) symlinkPatchFiles() error {
	a.Banner("Symlinking Patch Files")

	// get the patch files
	patchFiles := []string{
		"patch_RoF.conf",
		"patch_RoF2.conf",
		"patch_SoD.conf",
		"patch_SoF.conf",
		"patch_Titanium.conf",
		"patch_UF.conf",
	}

	// symlink the patch files
	for _, patchFile := range patchFiles {
		// get the patch file name
		patchFileName := filepath.Base(patchFile)

		// get the symlink path
		symlinkPath := filepath.Join(a.pathmanager.GetEQEmuServerPath(), "assets", "patches", patchFileName)

		// remove the symlink
		_ = os.Remove(symlinkPath)

		// check if the symlink exists
		if _, err := os.Stat(symlinkPath); !os.IsNotExist(err) {
			fmt.Printf("Symlink [%v] already exists, skipping\n", symlinkPath)

			// remove the symlink
			err = os.Remove(symlinkPath)
			if err != nil {
				return fmt.Errorf("could not remove symlink: %v", err)
			}
		}

		sourcePatchPath := filepath.Join(a.installConfig.CodePath, "utils", "patches", patchFile)

		// create the symlink
		fmt.Printf("Creating symlink [%v] -> [%v]\n", symlinkPath, sourcePatchPath)
		err := os.Symlink(sourcePatchPath, symlinkPath)
		if err != nil {
			return fmt.Errorf("could not create symlink: %v", err)
		}
	}

	a.DoneBanner("Symlinking Patch Files")
	return nil
}

func (a *Installer) symlinkOpcodeFiles() error {
	a.Banner("Symlinking Opcode Files")

	// get the opcode files
	opcodeFiles := []string{
		"opcodes.conf",
		"mail_opcodes.conf",
	}

	// symlink the opcode files
	for _, opcodeFile := range opcodeFiles {
		// get the opcode file name
		opcodeFileName := filepath.Base(opcodeFile)

		// get the symlink path
		symlinkPath := filepath.Join(a.pathmanager.GetEQEmuServerPath(), "assets", "opcodes", opcodeFileName)

		// remove the symlink, ignore errors
		_ = os.Remove(symlinkPath)

		sourcePatchPath := filepath.Join(a.installConfig.CodePath, "utils", "patches", opcodeFile)

		// create the symlink
		fmt.Printf("Creating symlink [%v] -> [%v]\n", symlinkPath, sourcePatchPath)
		err := os.Symlink(sourcePatchPath, symlinkPath)
		if err != nil {
			return fmt.Errorf("could not create symlink: %v", err)
		}
	}

	a.DoneBanner("Symlinking Opcode Files")
	return nil
}

func (a *Installer) symlinkLoginOpcodeFiles() error {
	a.Banner("Symlinking Login Opcode Files")

	// get the opcode files
	opcodeFiles := []string{
		"login_opcodes.conf",
		"login_opcodes_sod.conf",
	}

	// symlink the opcode files
	for _, opcodeFile := range opcodeFiles {
		// get the opcode file name
		opcodeFileName := filepath.Base(opcodeFile)

		// get the symlink path
		symlinkPath := filepath.Join(a.pathmanager.GetEQEmuServerPath(), "assets", "opcodes", opcodeFileName)

		// remove the symlink, ignore errors
		_ = os.Remove(symlinkPath)

		sourcePatchPath := filepath.Join(a.installConfig.CodePath, "loginserver", "login_util", opcodeFile)

		// create the symlink
		fmt.Printf("Creating symlink [%v] -> [%v]\n", symlinkPath, sourcePatchPath)
		err := os.Symlink(sourcePatchPath, symlinkPath)
		if err != nil {
			return fmt.Errorf("could not create symlink: %v", err)
		}
	}

	a.DoneBanner("Symlinking Login Opcode Files")
	return nil
}

func (a *Installer) symlinkPluginsAndModules() error {
	a.Banner("Symlinking Plugins and Modules")

	// get the symlink paths
	sourceLuaModules := filepath.Join(a.pathmanager.GetEQEmuServerPath(), "quests", "lua_modules")
	targetLuaModules := filepath.Join(a.pathmanager.GetEQEmuServerPath(), "lua_modules")

	// get the symlink paths
	sourcePerlPlugins := filepath.Join(a.pathmanager.GetEQEmuServerPath(), "quests", "plugins")
	targetPerlPlugins := filepath.Join(a.pathmanager.GetEQEmuServerPath(), "plugins")

	// remove the symlink, ignore errors
	_ = os.Remove(targetLuaModules)
	_ = os.Remove(targetPerlPlugins)

	// create the symlink
	fmt.Printf("Creating symlink [%v] -> [%v]\n", targetLuaModules, sourceLuaModules)
	err := os.Symlink(sourceLuaModules, targetLuaModules)
	if err != nil {
		return fmt.Errorf("could not create symlink: %v", err)
	}

	// create the symlink
	fmt.Printf("Creating symlink [%v] -> [%v]\n", targetPerlPlugins, sourcePerlPlugins)
	err = os.Symlink(sourcePerlPlugins, targetPerlPlugins)
	if err != nil {
		return fmt.Errorf("could not create symlink: %v", err)
	}

	a.DoneBanner("Symlinking Plugins and Modules")
	return nil
}

func (a *Installer) GetRandomPassword() (string, error) {
	p, err := password.Generate(32, 10, 0, false, false)
	if err != nil {
		return "", fmt.Errorf("could not generate random password: %v", err)
	}

	return p, nil
}

func (a *Installer) checkIfMapsAreUpToDate() error {
	type Release struct {
		TagName string `json:"tag_name"`
	}

	// get latest release version
	resp, err := http.Get("https://api.github.com/repos/Akkadius/eqemu-maps/releases/latest")
	if err != nil {
		return fmt.Errorf("could not get latest release version: %v", err)
	}

	defer resp.Body.Close()

	// read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("could not read response body: %v", err)
	}

	// bind body to struct
	var release Release
	err = json.Unmarshal(body, &release)
	if err != nil {
		return fmt.Errorf("could not unmarshal response body: %v", err)
	}

	fmt.Printf("Downloading eqemu-maps release\n")

	type PackageJson struct {
		Version string `json:"version"`
	}

	// get current version from package.json
	file := filepath.Join(a.pathmanager.GetEQEmuServerPath(), "maps", "package.json")

	// check if file exists
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return errors.New("package.json does not exist for maps")
	}

	// read file package.json contents into PackageJson struct
	packageJson, err := os.ReadFile(file)
	if err != nil {
		return fmt.Errorf("could not read PackageJson file: %v", err)
	}

	// bind package.json to struct
	var packageJsonStruct PackageJson
	err = json.Unmarshal(packageJson, &packageJsonStruct)
	if err != nil {
		return fmt.Errorf("could not unmarshal package.json: %v", err)
	}

	// check if current version is the same as the latest release version
	remoteVersion := strings.ReplaceAll(release.TagName, "v", "")

	if len(remoteVersion) == 0 {
		fmt.Printf("Could not retrieve latest [eqemu-maps] version, possibly rate limited, skipping\n")
		return nil
	}

	if len(remoteVersion) > 0 && packageJsonStruct.Version == remoteVersion {
		fmt.Printf("Maps are up to date on version v%v\n", packageJsonStruct.Version)
		return nil
	}

	return errors.New("maps are not up to date")
}

func (a *Installer) setInstallerPath() error {
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("could not get current working directory: %v", err)
	}
	a.pathmanager.SetServerPath(cwd)
	return nil
}

func (a *Installer) createLinuxServerScripts() error {
	a.Banner("Creating Server Scripts")

	// create a map of scripts
	serverScripts := map[string]string{
		"start":   "./spire spire:launcher start && echo \"Server started\"",
		"stop":    "./spire spire:launcher stop; echo \"Server stopped\"",
		"restart": "./spire spire:launcher restart; echo \"Server restarting\"",
	}

	for s := range serverScripts {

		// get the f name
		file := filepath.Join(a.pathmanager.GetEQEmuServerPath(), s)

		fmt.Printf("Creating script [%v]\n", file)

		// create file
		f, err := os.Create(file)
		if err != nil {
			return fmt.Errorf("could not create f: %v", err)
		}

		// write contents to f
		contents := fmt.Sprintf("#!/usr/bin/env bash\n%v\n", serverScripts[s])
		_, err = f.WriteString(contents)
		if err != nil {
			return fmt.Errorf("could not write to f: %v", err)
		}

		// close file
		_ = f.Close()

		fmt.Printf("|-- Making file [%v] executable\n", file)

		// make file executable
		err = os.Chmod(file, 0755)
		if err != nil {
			return fmt.Errorf("could not chmod f: %v", err)
		}
	}

	a.DoneBanner("Creating Server Scripts")
	return nil
}

func (a *Installer) injectSpireStartCronJob() error {
	a.Banner("Injecting Spire Start Cron Job")

	// get spire path
	spirePath := filepath.Join(a.pathmanager.GetEQEmuServerPath(), "spire")

	cronInject := fmt.Sprintf("while true; do nohup %s > %s/logs/spire.log 2>&1; sleep 1; done &", spirePath, a.pathmanager.GetEQEmuServerPath())

	_, err := a.Exec(
		ExecConfig{
			command: "bash",
			args: []string{
				"-c",
				fmt.Sprintf("crontab -l | grep -qF 'spire' || (crontab -l 2>/dev/null; echo \"@reboot %v\") | crontab -", cronInject),
			},
		},
	)
	if err != nil {
		return err
	}

	a.DoneBanner("Injecting Spire Start Cron Job")

	return nil
}

func (a *Installer) installSpireBinary() error {
	a.Banner("Installing Spire")

	// check if spire is already installed
	spirePath, err := exec.LookPath(filepath.Join(a.pathmanager.GetEQEmuServerPath(), "spire"))
	if err != nil {
		fmt.Printf("could not find spire binary: %v", err)
	}

	if _, err := os.Stat(spirePath); !os.IsNotExist(err) {
		// spire is already installed
		fmt.Printf("Spire is already installed at [%v]\n", spirePath)
		a.DoneBanner("Installing Spire")
		return nil
	}

	// check if spire is already installed
	if _, err := os.Stat(spirePath); os.IsNotExist(err) {
		client := github.NewClient(&http.Client{Timeout: 5 * time.Second})
		release, _, err := client.Repositories.GetLatestRelease(
			context.Background(),
			"Akkadius",
			"spire",
		)
		if err != nil {
			fmt.Printf("could not get latest release: %v", err)
			return nil
		}

		for _, asset := range release.Assets {
			assetName := *asset.Name
			downloadUrl := *asset.BrowserDownloadURL
			targetFileNameZipped := fmt.Sprintf("spire-%s-%s.zip", runtime.GOOS, runtime.GOARCH)
			if runtime.GOOS == "windows" {
				targetFileNameZipped = fmt.Sprintf("spire-%s-%s.exe.zip", runtime.GOOS, runtime.GOARCH)
			}

			targetFileName := fmt.Sprintf("spire-%s-%s", runtime.GOOS, runtime.GOARCH)

			if assetName == targetFileNameZipped {
				fmt.Printf("Found matching release [%s]\n", assetName)

				// download
				file := path.Base(downloadUrl)
				downloadPath := filepath.Join(os.TempDir(), file)
				err := download.WithProgress(downloadPath, downloadUrl)
				if err != nil {
					return fmt.Errorf("could not download spire: %v", err)
				}

				// unzip
				tempFileZipped := fmt.Sprintf("%s/%s", os.TempDir(), targetFileNameZipped)
				unzipPathTemp := filepath.Join(fmt.Sprintf("%s/spire-download", os.TempDir()))
				uz := unzip.New(tempFileZipped, unzipPathTemp)
				fmt.Printf("|-- Unzipping file [%v] to [%v]\n", tempFileZipped, a.pathmanager.GetEQEmuServerPath())
				err = uz.Extract()
				if err != nil {
					return fmt.Errorf("could not unzip spire: %v", err)
				}

				// rename
				src, err := exec.LookPath(filepath.Join(unzipPathTemp, targetFileName))
				if err != nil {
					return fmt.Errorf("could not find spire: %v", err)
				}

				// new spire path
				newSpirePath := filepath.Join(a.pathmanager.GetEQEmuServerPath(), "spire")
				if runtime.GOOS == "windows" {
					newSpirePath = filepath.Join(a.pathmanager.GetEQEmuServerPath(), "spire.exe")
				}

				// destination
				dst := newSpirePath

				fmt.Printf("|-- Renaming file [%v] to [%v]\n", src, dst)

				// copy file from src to dst
				err = a.Copy(src, dst)
				if err != nil {
					return err
				}

				fmt.Printf("|-- Making file [%v] executable\n", dst)

				// make executable
				err = os.Chmod(dst, 0755)
				if err != nil {
					return fmt.Errorf("could not chmod spire: %v", err)
				}
			}
		}
	}

	a.DoneBanner("Installing Spire")
	return nil
}

func (a *Installer) initSpire() error {
	a.Banner("Initializing Spire")

	// check if spire is already installed
	spirePath, err := exec.LookPath(filepath.Join(a.pathmanager.GetEQEmuServerPath(), "spire"))
	if err != nil {
		return fmt.Errorf("could not find spire binary: %v", err)
	}

	args := []string{
		"spire:init",
		a.installConfig.SpireAdminUser,
		a.installConfig.SpireAdminPassword,
	}

	// if we're compiling binaries
	if a.installConfig.CompileBinaries {
		buildPath := filepath.Join(a.installConfig.CodePath, "build")
		args = append(args, "--compile-server=true")
		args = append(args, fmt.Sprintf("--compile-build-location=%v", buildPath))
	}

	_, err = a.Exec(ExecConfig{
		command:    spirePath,
		args:       args,
		hidestring: a.installConfig.SpireAdminPassword,
	})
	if err != nil {
		return err
	}

	_, err = a.Exec(ExecConfig{
		command:    spirePath,
		args:       []string{"spire:occulus-update"},
		hidestring: a.installConfig.SpireAdminPassword,
	})
	if err != nil {
		return err
	}

	a.DoneBanner("Initializing Spire")
	return nil
}

func (a *Installer) runSharedMemory() error {
	a.Banner("Running Shared Memory")

	_, err := a.Exec(ExecConfig{
		execpath: a.pathmanager.GetEQEmuServerPath(),
		command:  filepath.Join("bin", "shared_memory"),
	})
	if err != nil {
		return err
	}

	a.DoneBanner("Running Shared Memory")

	return nil
}

func (a *Installer) runWorldForDatabaseUpdates() error {
	a.Banner("Running World for Database Updates")

	worldPath, err := exec.LookPath(filepath.Join(a.pathmanager.GetEQEmuServerPath(), "bin", "world"))
	if err != nil {
		return fmt.Errorf("could not find world binary: %v", err)
	}

	_, err = a.Exec(ExecConfig{
		execpath:    a.pathmanager.GetEQEmuServerPath(),
		command:     worldPath,
		dieonoutput: "Server (TCP) listener started on port",
	})
	if err != nil {
		return err
	}

	a.DoneBanner("Running World for Database Updates")

	return nil
}

func (a *Installer) runZoneForDataInjections() error {
	a.Banner("Running Zone for Data Injections")

	_, err := a.Exec(ExecConfig{
		execpath:    a.pathmanager.GetEQEmuServerPath(),
		command:     filepath.Join("bin", "zone"),
		dieonoutput: "Entering sleep mode",
	})
	if err != nil {
		return err
	}

	a.DoneBanner("Running Zone for Data Injections")

	return nil
}

func (a *Installer) startSpire() error {
	a.Banner("Starting Spire")

	// kill any running spire processes
	processes, _ := process.Processes()
	for _, p := range processes {
		cmdline, err := p.Cmdline()
		if err != nil {
			fmt.Printf("could not get cmdline for process: %v\n", err)
			continue
		}

		// kill spire if it's running
		// ignore spire processes that are running from /tmp
		installerRan := strings.Contains(cmdline, "/tmp/") || strings.Contains(cmdline, "-install")
		if strings.Contains(cmdline, "spire") && !installerRan {
			err := p.Kill()
			if err != nil {
				return fmt.Errorf("could not kill process: %v\n", err)
			}
		}
	}

	fmt.Printf("starting spire [%v]", filepath.Join(a.pathmanager.GetEQEmuServerPath(), "spire"))

	// start spire in a loop
	if runtime.GOOS == "linux" {
		_, err := a.Exec(ExecConfig{
			command: "bash",
			args: []string{
				"-c",
				fmt.Sprintf(
					"while true; do nohup %s > %s/logs/spire.log 2>&1; sleep 1; done &",
					filepath.Join(a.pathmanager.GetEQEmuServerPath(), "spire"),
					a.pathmanager.GetEQEmuServerPath(),
				),
			},
			detach: true,
		})
		if err != nil {
			return err
		}
	}

	if runtime.GOOS == "windows" {
		spirePath, err := exec.LookPath(filepath.Join(a.pathmanager.GetEQEmuServerPath(), "spire"))
		if err != nil {
			return fmt.Errorf("could not find spire binary: %v", err)
		}

		_, err = a.Exec(ExecConfig{
			command: "cmd",
			args: []string{
				"/c",
				fmt.Sprintf(
					"start %s http:serve --port=%v > %s/logs/spire.log 2>&1",
					spirePath,
					a.installConfig.SpireWebPort,
					a.pathmanager.GetEQEmuServerPath(),
				),
			},
			detach: true,
		})
		if err != nil {
			return err
		}
	}

	a.DoneBanner("Starting Spire")
	return nil
}

// getLinuxDistribution returns the linux distribution
func (a *Installer) getLinuxDistribution() (string, error) {
	// determine whether we're on ubuntu or debian
	// read from /etc/os-release to determine which distro we're on
	if _, err := os.Stat("/etc/os-release"); os.IsNotExist(err) {
		return "", fmt.Errorf("could not find /etc/os-release")
	}

	// get contents of /etc/os-release
	osRelease, err := os.ReadFile("/etc/os-release")
	if err != nil {
		return "", fmt.Errorf("could not read /etc/os-release: %v", err)
	}

	// parse /etc/os-release
	for _, line := range strings.Split(string(osRelease), "\n") {

		// skip empty lines
		if len(line) == 0 {
			continue
		}

		// split on "="
		split := strings.Split(line, "=")
		key := strings.TrimSpace(split[0])
		value := strings.TrimSpace(split[1])

		// if we find the NAME key, return the value
		if key == "NAME" {
			if strings.Contains(value, "Ubuntu") {
				return "ubuntu", nil
			} else if strings.Contains(value, "Debian") {
				return "debian", nil
			} else {
				return "unknown", fmt.Errorf("unknown OS: %s", value)
			}
		}
	}

	return "unknown", nil
}

// getLinuxDistributionVersion returns the linux distribution version
func (a *Installer) getLinuxDistributionVersion() (int, error) {
	// read from /etc/os-release to determine which version of the distro we're on
	if _, err := os.Stat("/etc/os-release"); os.IsNotExist(err) {
		return 0, fmt.Errorf("could not find /etc/os-release")
	}

	// get contents of /etc/os-release
	osRelease, err := os.ReadFile("/etc/os-release")
	if err != nil {
		return 0, fmt.Errorf("could not read /etc/os-release: %v", err)
	}

	// parse /etc/os-release
	for _, line := range strings.Split(string(osRelease), "\n") {

		// skip empty lines
		if len(line) == 0 {
			continue
		}

		// split on "="
		split := strings.Split(line, "=")
		key := strings.TrimSpace(split[0])
		value := strings.TrimSpace(split[1])

		// if we find the VERSION_ID key, return the value
		if key == "VERSION_ID" {
			version := strings.ReplaceAll(value, "\"", "")
			// conver to int
			versionInt, err := strconv.Atoi(version)
			if err != nil {
				return 0, fmt.Errorf("could not convert version to int: %v", err)
			}

			return versionInt, nil
		}
	}

	return 0, nil
}

func (a *Installer) initWindowsMysql() error {
	a.Banner("Downloading MariaDB")

	// check if mysql folder exists
	if _, err := os.Stat(a.getWindowsMysqlPath()); err == nil {
		// login to mysql
		mysqlPath := filepath.Join(a.getWindowsMysqlPath(), "mysqladmin.exe")
		fmt.Printf("Logging into MySQL\n")
		res, err := a.Exec(ExecConfig{
			command: mysqlPath,
			args: []string{
				fmt.Sprintf("-u%v", a.installConfig.MysqlUsername),
				fmt.Sprintf("-p%v", a.installConfig.MysqlPassword),
				"ping",
			},
		})
		if err != nil {
			return err
		}

		// check if mysql is alive
		if strings.Contains(res, "mysqld is alive") {
			fmt.Printf("MySQL already installed, skipping")
			return nil
		}
	}

	// download mariadb
	// download the latest binaries
	tempPath := filepath.Join(os.TempDir(), "mariadb.msi")
	fmt.Printf("Downloading binaries to [%v]\n", tempPath)
	err := download.WithProgress(
		tempPath,
		"https://github.com/Akkadius/eqemu-install-v2/releases/download/static/mariadb-10.11.4-winx64.msi",
	)
	if err != nil {
		return err
	}

	// install mariadb
	// start /wait msiexec /i mariadb-10.0.21-winx64.msi SERVICENAME=MySQL PORT=3306 PASSWORD=eqemu /qn
	// TODO: make port configurable
	// TODO: split out root user and eqemu user passwords
	fmt.Printf("Installing MariaDB\n")
	_, err = a.Exec(ExecConfig{
		command: "msiexec",
		args: []string{
			"/i",
			tempPath,
			"SERVICENAME=MySQL",
			fmt.Sprintf("PORT=%v", a.installConfig.MysqlPort),
			"BUFFERPOOLSIZE=1024",
			fmt.Sprintf("PASSWORD=%s", a.installConfig.MysqlPassword),
			"/qn",
			"/l*v",
			"mariadb-install-log.txt",
		},
		hidestring: a.installConfig.MysqlPassword,
	})
	if err != nil {
		return err
	}

	c := MysqlConfig{
		DatabaseName:     a.installConfig.MysqlDatabaseName,
		DatabaseUser:     a.installConfig.MysqlUsername,
		DatabasePassword: a.installConfig.MysqlPassword,
	}

	// create a new database
	var sql string

	fmt.Printf("Creating database [%v]\n", c.DatabaseName)
	err = a.DbExec(DbExecConfig{statement: fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %v", c.DatabaseName)})
	if err != nil {
		return err
	}

	// create a new user
	fmt.Printf("Creating user [%v]\n", c.DatabaseUser)

	// grant privileges to the new user
	fmt.Printf("Granting privileges to user [%v]\n", c.DatabaseUser)
	sql += fmt.Sprintf("CREATE USER IF NOT EXISTS '%v'@'localhost' IDENTIFIED BY '%v'; ", c.DatabaseUser, c.DatabasePassword)
	sql += fmt.Sprintf("GRANT ALL PRIVILEGES ON %v.* TO '%v'@'localhost'", c.DatabaseName, c.DatabaseUser)

	// flush privileges
	fmt.Println("Flushing privileges")
	err = a.DbExec(DbExecConfig{statement: fmt.Sprintf("FLUSH PRIVILEGES; %v; FLUSH PRIVILEGES;", sql), hidestring: c.DatabasePassword})
	if err != nil {
		return err
	}

	err = a.setWindowsMysqlPath()
	if err != nil {
		return err
	}

	a.DoneBanner("Downloading MariaDB")
	return nil
}

func (a *Installer) initWindowsPerl() error {
	a.Banner("Downloading Perl")

	// check if a.getWindowsPerlPath() exists
	if _, err := os.Stat(a.getWindowsPerlPath()); err == nil {
		fmt.Printf("Perl already installed, skipping\n")
		return nil
	}

	// download mariadb
	// download the latest binaries
	tempPath := filepath.Join(os.TempDir(), "perl.msi")
	fmt.Printf("Downloading binaries to [%v]\n", tempPath)
	err := download.WithProgress(
		tempPath,
		"https://github.com/Akkadius/eqemu-install-v2/releases/download/static/strawberry-perl-5.24.4.1-64bit.msi",
	)
	if err != nil {
		return err
	}

	// install perl
	// start /wait msiexec /i strawberry-perl-5.24.4.1-64bit.msi PERL_PATH="Yes" /q
	// start /wait msiexec /i mariadb-10.0.21-winx64.msi SERVICENAME=MySQL PORT=3306 PASSWORD=eqemu /qn
	fmt.Printf("Installing Perl\n")
	_, err = a.Exec(ExecConfig{
		command: "msiexec",
		args: []string{
			"/i",
			tempPath,
			"PERL_PATH=Yes",
			"/q",
		},
	})
	if err != nil {
		return err
	}

	err = a.setWindowsPerlPath()
	if err != nil {
		return err
	}

	a.DoneBanner("Downloading Perl")
	return nil
}

// initWindowsWget downloads wget for windows (backwards compatibility)
// TODO: remove this in the future
func (a *Installer) initWindowsWget() error {
	a.Banner("Downloading Windows wget")

	downloadPath := filepath.Join(a.pathmanager.GetEQEmuServerBinPath(), "wget.exe")
	fmt.Printf("Downloading binaries to [%v]\n", downloadPath)
	err := download.WithProgress(
		downloadPath,
		"https://github.com/Akkadius/eqemu-install-v2/releases/download/static/wget.exe",
	)
	if err != nil {
		return err
	}

	a.DoneBanner("Downloading Windows wget")
	return nil
}

// initWindowsMysqlService installs and configures the mysql service
func (a *Installer) getWindowsProgramFilesPath() string {
	// get program files path
	s, _ := a.Exec(ExecConfig{command: "cmd", args: []string{"/c", "echo", "%programfiles%"}, silent: true})

	return strings.TrimSpace(s)
}

// getWindowsMysqlPath returns the path to the mysql installation
func (a *Installer) getWindowsMysqlPath() string {
	// get folders in folder using go
	entries, err := os.ReadDir(filepath.Join(a.getWindowsProgramFilesPath()))
	if err != nil {
		a.logger.Fatal(err)
	}

	// first look for mariadb
	for _, e := range entries {
		if strings.Contains(strings.ToLower(e.Name()), "mariadb") {
			return filepath.Join(a.getWindowsProgramFilesPath(), e.Name(), "bin")
		}
	}

	// second look for mysql
	for _, e := range entries {
		if strings.Contains(strings.ToLower(e.Name()), "mysql") {
			return filepath.Join(a.getWindowsProgramFilesPath(), e.Name(), "bin")
		}
	}

	return ""
}

// getWindowsPerlPath returns the path to the mysql installation
func (a *Installer) getWindowsPerlPath() string {
	// cwd
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	// get the first 3 letters of cwd
	drive := cwd[0:3]

	return filepath.Join(drive, "Strawberry", "perl", "bin")
}

// Copy copies a file from src to dst
func (a *Installer) Copy(src string, dst string) error {
	// read the file
	bytesRead, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	// write the file
	err = os.WriteFile(dst, bytesRead, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

// createWindowsServerScripts creates the server scripts
func (a *Installer) createWindowsServerScripts() error {
	a.Banner("Creating Server Scripts")

	// create a map of scripts
	serverScripts := map[string]string{
		"server_restart.bat":         "spire.exe spire:launcher restart\n@echo off\necho Server is restarting\ntimeout /T 3 /NOBREAK > nul",
		"server_start.bat":           "spire.exe spire:launcher start\n@echo off\necho Server is starting\ntimeout /T 3 /NOBREAK > nul",
		"server_stop.bat":            "spire.exe spire:launcher stop\n@echo off\necho Server is stopping\ntimeout /T 3 /NOBREAK > nul",
		"spire_start.bat":            "TASKKILL /IM spire.exe /F\nstart /min spire.exe > logs/spire.log 2>&1",
		"spire_stop.bat":             "TASKKILL /IM spire.exe /F",
		"spire_web.bat":              fmt.Sprintf("start http://localhost:%v", a.installConfig.SpireWebPort),
		"spire_web_server_admin.bat": fmt.Sprintf("start http://localhost:%v/admin", a.installConfig.SpireWebPort),
	}

	for s := range serverScripts {

		// get the f name
		file := filepath.Join(a.pathmanager.GetEQEmuServerPath(), s)

		fmt.Printf("Creating script [%v]\n", file)

		// create file
		f, err := os.Create(file)
		if err != nil {
			return fmt.Errorf("could not create f: %v", err)
		}

		// write contents to f
		contents := fmt.Sprintf("%v\n", serverScripts[s])
		_, err = f.WriteString(contents)
		if err != nil {
			return fmt.Errorf("could not write to f: %v", err)
		}

		// close file
		_ = f.Close()

		fmt.Printf("|-- Making file [%v] executable\n", file)

		// make file executable
		err = os.Chmod(file, 0755)
		if err != nil {
			return fmt.Errorf("could not chmod f: %v", err)
		}
	}

	a.DoneBanner("Creating Server Scripts")
	return nil
}

func (a *Installer) initWindowsCommandPrompt() error {
	cmd := exec.Command("chcp", "65001")
	cmd.Env = os.Environ()
	cmd.Dir = a.pathmanager.GetEQEmuServerPath()
	_, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("could not get stdout pipe: %v", err)
	}
	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func (a *Installer) setWindowsMysqlPath() error {
	// check if path already contains mysql
	if !strings.Contains(os.Getenv("Path"), a.getWindowsMysqlPath()) {
		fmt.Printf("Updating PATH to include [%v]\n", a.getWindowsMysqlPath())
		err := os.Setenv("Path", fmt.Sprintf("%v;%v", os.Getenv("Path"), a.getWindowsMysqlPath()))
		if err != nil {
			return err
		}

		cmd := exec.Command(
			"setx",
			"PATH",
			fmt.Sprintf(
				"%v",
				os.Getenv("PATH"),
			),
		)
		cmd.Env = os.Environ()
		cmd.Dir = a.pathmanager.GetEQEmuServerPath()

		// tie command stdout to os stdout
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err = cmd.Run()
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *Installer) setWindowsPerlPath() error {
	// update current notion of path temporarily since we don't have the updated path in the current cmd shell
	if !strings.Contains(os.Getenv("Path"), a.getWindowsPerlPath()) {
		fmt.Printf("Updating PATH to include [%v]\n", a.getWindowsPerlPath())
		err := os.Setenv("Path", fmt.Sprintf("%v;%v", os.Getenv("Path"), a.getWindowsPerlPath()))
		if err != nil {
			return err
		}

		// exec command setx PATH "%PATH%;%FOO%"
		cmd := exec.Command(
			"setx",
			"PATH",
			fmt.Sprintf(
				"%v",
				os.Getenv("PATH"),
			),
		)
		cmd.Env = os.Environ()
		cmd.Dir = a.pathmanager.GetEQEmuServerPath()

		// tie command stdout to os stdout
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err = cmd.Run()
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *Installer) setPostInstallConfigValues() error {
	// load the config
	config := a.config.Get()

	// set the post install config values
	if config.WebAdmin != nil {

		// initialize the launcher config if it doesn't exist
		if config.WebAdmin.Launcher == nil {
			config.WebAdmin.Launcher = &eqemuserverconfig.WebAdminLauncherConfig{}
		}

		config.WebAdmin.Launcher.MinZoneProcesses = 10
		config.WebAdmin.Launcher.RunSharedMemory = true
		// boat zones mainly
		config.WebAdmin.Launcher.StaticZones = "butcher,erudnext,freporte,qeynos,freeporte,oot,iceclad,nro,oasis,nedaria,abysmal,natimbi,timorous,abysmal,firiona,overthere"
	}

	err := a.config.Save(config)
	if err != nil {
		return fmt.Errorf("could not save config: %v", err)
	}
	return nil
}

func (a *Installer) initLoginServer() error {
	a.Banner("Initializing Login Server")

	// download the default login config
	url := "https://raw.githubusercontent.com/EQEmu/Server/master/loginserver/login_util/login.json"
	fmt.Printf("Downloading default login config from [%v]\n", url)
	err := download.WithProgress(a.pathmanager.GetEqemuLoginServerConfigPath(), url)
	if err != nil {
		return fmt.Errorf("could not download login config: %v", err)
	}

	// hyrdrate the config
	c := a.loginConfig.Get()
	c.Database.Host = a.installConfig.MysqlHost
	c.Database.Port = a.installConfig.MysqlPort
	c.Database.User = a.installConfig.MysqlUsername
	c.Database.Password = a.installConfig.MysqlPassword
	c.Database.Db = a.installConfig.MysqlDatabaseName
	c.ClientConfiguration.SodOpcodes = "assets/opcodes/login_opcodes_sod.conf"
	c.ClientConfiguration.TitaniumOpcodes = "assets/opcodes/login_opcodes.conf"

	// save the config
	err = a.loginConfig.Save(c)
	if err != nil {
		return fmt.Errorf("could not save login config: %v", err)
	}

	mysqlPath := "mysql"
	if runtime.GOOS == "windows" {
		mysqlPath = filepath.Join(a.getWindowsMysqlPath(), "mysql.exe")
	}

	// create the login server database tables
	url = "https://raw.githubusercontent.com/EQEmu/Server/master/loginserver/login_util/login_schema.sql"
	fmt.Printf("Downloading login schema from [%v]\n", url)
	err = download.WithProgress(filepath.Join(os.TempDir(), "login_schema.sql"), url)
	if err != nil {
		return fmt.Errorf("could not download login config: %v", err)
	}

	fmt.Printf("Creating login server database tables\n")

	_, err = a.Exec(
		ExecConfig{
			command: mysqlPath,
			args: []string{
				fmt.Sprintf("-u%v", a.installConfig.MysqlUsername),
				fmt.Sprintf("-p%v", a.installConfig.MysqlPassword),
				a.installConfig.MysqlDatabaseName,
				"-e",
				"source login_schema.sql",
			},
			hidestring: a.installConfig.MysqlPassword,
			execpath:   os.TempDir(),
		},
	)
	if err != nil {
		return err
	}

	a.DoneBanner("Initializing Login Server")
	return nil
}

// disableQuickEdit disables quick edit mode in the windows console
func (a *Installer) disableQuickEdit() error {
	a.Banner("Disabling Quick Edit")

	cmd := exec.Command(
		"cmd.exe",
		"/C",
		"REG",
		"ADD",
		"HKCU\\CONSOLE",
		"/v",
		"QuickEdit",
		"/t",
		"REG_DWORD",
		"/d",
		"0",
		"/f",
	)
	cmd.Env = os.Environ()
	cmd.Dir = a.pathmanager.GetEQEmuServerPath()

	// tie command stdout to os stdout
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return err
	}

	a.DoneBanner("Disabling Quick Edit")
	return nil
}

// enableBots enables bots on the server
func (a *Installer) enableBots() error {
	if !a.installConfig.BotsEnabled {
		return nil
	}

	a.Banner("Enabling Bots")

	worldPath, err := exec.LookPath(filepath.Join(a.pathmanager.GetEQEmuServerPath(), "bin", "world"))
	if err != nil {
		return fmt.Errorf("could not find world binary: %v\n", err)
	}

	_, err = a.Exec(ExecConfig{
		execpath: a.pathmanager.GetEQEmuServerPath(),
		command:  worldPath,
		args:     []string{"bots:enable"},
	})
	if err != nil {
		return err
	}

	a.DoneBanner("Enabling Bots")
	return nil
}

// enableMercs enables bots on the server
func (a *Installer) enableMercenaries() error {
	if !a.installConfig.MercsEnabled {
		return nil
	}

	a.Banner("Enabling Mercenaries")

	worldPath, err := exec.LookPath(filepath.Join(a.pathmanager.GetEQEmuServerPath(), "bin", "world"))
	if err != nil {
		return fmt.Errorf("could not find world binary: %v\n", err)
	}

	_, err = a.Exec(ExecConfig{
		execpath: a.pathmanager.GetEQEmuServerPath(),
		command:  worldPath,
		args:     []string{"mercs:enable"},
	})
	if err != nil {
		return err
	}

	a.DoneBanner("Enabling Mercenaries")
	return nil
}

// compileBinaries compiles the server binaries
func (a *Installer) compileBinaries() error {
	a.Banner("Compiling Binaries")

	// get the build path
	codeBuildPath := filepath.Join(a.installConfig.CodePath, "build")

	// create the build directory
	err := os.MkdirAll(codeBuildPath, 0755)
	if err != nil {
		return fmt.Errorf("could not create build directory: %v", err)
	}

	args := []string{
		"-DEQEMU_BUILD_LOGIN=ON",
		"-DEQEMU_BUILD_LUA=ON",
	}

	// get user home directory
	homedir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("could not get user home directory: %v", err)
	}

	// check if system has ccache installed
	hasCcache := false
	ccacheDir := filepath.Join(homedir, ".ccache")
	if _, err := os.Stat(ccacheDir); err == nil {
		hasCcache = true
	}

	// use ccache if available
	if hasCcache {
		args = append(args, "-DCMAKE_CXX_COMPILER_LAUNCHER=ccache")
	}

	// compile with debug symbols and no optimization
	if a.installConfig.CompileDevelop {
		args = append(args, "-DCMAKE_CXX_FLAGS_RELWITHDEBINFO:STRING='-O0 -g -DNDEBUG'")
	}

	args = append(args, "-G", "Unix Makefiles", "..")

	_, err = a.Exec(ExecConfig{
		execpath: codeBuildPath,
		command:  "cmake",
		args:     args,
	})
	if err != nil {
		return err
	}

	cores := 1

	// get system memory available
	memory, err := mem.VirtualMemory()
	if err != nil {
		a.logger.Fatal(err)
	}

	// get system memory available in GB
	memoryAvailableGb := memory.Available / 1024 / 1024 / 1024
	if memoryAvailableGb >= 10 {
		cores = runtime.NumCPU() - 4
		if cores < 1 {
			cores = 1
		}
	} else if memoryAvailableGb >= 6 {
		cores = 4
	}

	_, err = a.Exec(ExecConfig{
		execpath: filepath.Join(a.installConfig.CodePath, "build"),
		command:  "make",
		args:     []string{"-j", strconv.Itoa(cores)},
	})
	if err != nil {
		return err
	}

	a.DoneBanner("Compiling binaries")

	a.Banner("Symlinking compiled binaries")

	// symlink the binaries
	// loop through bin path
	binPath := filepath.Join(a.installConfig.CodePath, "build", "bin")
	err = filepath.Walk(binPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		// get the file name
		fileName := info.Name()

		// skip .a extension files
		if filepath.Ext(fileName) == ".a" {
			return nil
		}

		fmt.Printf("Symlinking [%v] to [%v] in server bin directory\n", fileName, filepath.Join(a.pathmanager.GetEQEmuServerPath(), "bin", fileName))

		source := filepath.Join(binPath, fileName)
		destination := filepath.Join(a.pathmanager.GetEQEmuServerPath(), "bin", fileName)

		// remove symlink
		_ = os.Remove(destination)

		// symlink the file into server bin directory
		err = os.Symlink(source, destination)
		if err != nil {
			return err
		}

		return nil
	})

	a.DoneBanner("Symlinking compiled binaries")
	return nil
}

func (a *Installer) publishWindowsPorts() error {
	a.Banner("Publishing Windows Firewall rules")
	// netsh advfirewall firewall add rule name="EQEmu Loginserver (Titanium) (5998) TCP" dir=in action=allow protocol=TCP localport=5998
	// write using cmd.exe
	_, _ = a.Exec(ExecConfig{command: "cmd.exe", args: []string{"/c", "netsh advfirewall firewall add rule name=\"EQEmu Loginserver (Titanium) (5998) TCP\" dir=in action=allow protocol=TCP localport=5998"}})
	_, _ = a.Exec(ExecConfig{command: "cmd.exe", args: []string{"/c", "netsh advfirewall firewall add rule name=\"EQEmu Loginserver (Titanium) (5998) UDP\" dir=in action=allow protocol=UDP localport=5998"}})
	_, _ = a.Exec(ExecConfig{command: "cmd.exe", args: []string{"/c", "netsh advfirewall firewall add rule name=\"EQEmu Loginserver (SoD+) (5999) TCP\" dir=in action=allow protocol=TCP localport=5999"}})
	_, _ = a.Exec(ExecConfig{command: "cmd.exe", args: []string{"/c", "netsh advfirewall firewall add rule name=\"EQEmu Loginserver (SoD+) (5999) UDP\" dir=in action=allow protocol=UDP localport=5999"}})
	// open ports 7000-7100 for zones
	_, _ = a.Exec(ExecConfig{command: "cmd.exe", args: []string{"/c", "netsh advfirewall firewall add rule name=\"EQEmu Zones (7000-7100) TCP\" dir=in action=allow protocol=TCP localport=7000-7100"}})
	_, _ = a.Exec(ExecConfig{command: "cmd.exe", args: []string{"/c", "netsh advfirewall firewall add rule name=\"EQEmu Zones (7000-7100) UDP\" dir=in action=allow protocol=UDP localport=7000-7100"}})
	// open ports 9000-9001 for world
	_, _ = a.Exec(ExecConfig{command: "cmd.exe", args: []string{"/c", "netsh advfirewall firewall add rule name=\"EQEmu World (9000-9001) TCP\" dir=in action=allow protocol=TCP localport=9000-9001"}})
	_, _ = a.Exec(ExecConfig{command: "cmd.exe", args: []string{"/c", "netsh advfirewall firewall add rule name=\"EQEmu World UDP (9000-9001)\" dir=in action=allow protocol=UDP localport=9000-9001"}})
	// open ports 7778 for ucs
	_, _ = a.Exec(ExecConfig{command: "cmd.exe", args: []string{"/c", "netsh advfirewall firewall add rule name=\"EQEmu UCS (7778) TCP\" dir=in action=allow protocol=TCP localport=7778"}})
	_, _ = a.Exec(ExecConfig{command: "cmd.exe", args: []string{"/c", "netsh advfirewall firewall add rule name=\"EQEmu UCS (7778) UDP\" dir=in action=allow protocol=UDP localport=7778"}})
	// open configurable port for spire (TCP)
	_, _ = a.Exec(ExecConfig{command: "cmd.exe", args: []string{"/c", fmt.Sprintf("netsh advfirewall firewall add rule name=\"EQEmu Spire (TCP) (%v)\" dir=in action=allow protocol=TCP localport=%v", a.installConfig.SpireWebPort, a.installConfig.SpireWebPort)}})

	a.DoneBanner("Publishing Windows Firewall rules")

	return nil
}

func (a *Installer) openWindowsPostInstallWindows() {
	installConfigFile := filepath.Join(a.pathmanager.GetEQEmuServerPath(), installConfigFileName)
	_, _ = a.Exec(ExecConfig{command: "notepad.exe", args: []string{installConfigFile}, detach: true})
	_, _ = a.Exec(ExecConfig{command: "explorer.exe", args: []string{a.pathmanager.GetEQEmuServerPath()}})
	_, _ = a.Exec(
		ExecConfig{
			command: "rundll32",
			args: []string{
				"url.dll,FileProtocolHandler",
				fmt.Sprintf(
					"http://localhost:%v/login?user=%v&password=%v&redirect=/admin",
					a.installConfig.SpireWebPort,
					a.installConfig.SpireAdminUser,
					a.installConfig.SpireAdminPassword,
				),
			},
			hidestring: a.installConfig.SpireAdminPassword,
		},
	)
}
