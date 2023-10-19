package eqemuserver

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Akkadius/spire/internal/download"
	"github.com/Akkadius/spire/internal/eqemuserverconfig"
	"github.com/Akkadius/spire/internal/password"
	"github.com/Akkadius/spire/internal/pathmgmt"
	"github.com/Akkadius/spire/internal/unzip"
	"github.com/go-git/go-git/v5"
	"github.com/google/go-github/v41/github"
	"github.com/k0kubun/pp/v3"
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
	}

	return i
}

func (a *Installer) Install() {
	a.setInstallerPath()
	a.checkInstallConfig()

	// install debian packages
	// install ubuntu packages (for ubuntu)
	a.totalTime = time.Now()
	if runtime.GOOS == "linux" {
		a.installLinuxOsPackages()
		a.initLinuxMysql()
	}

	if runtime.GOOS == "windows" {
		a.initWindowsCommandPrompt()
		a.initWindowsPerl()
		a.initWindowsMysql()
	}

	a.initializeDirectories()
	a.cloneEQEmuSource()
	a.initializeServerConfig()
	a.cloneEQEmuMaps()
	a.clonePeqQuests()
	a.sourcePeqDatabase()
	a.installBinaries()
	a.symlinkPatchFiles()
	a.symlinkOpcodeFiles()
	a.symlinkLoginOpcodeFiles()
	a.symlinkPluginsAndModules()
	if runtime.GOOS == "linux" {
		a.createLinuxServerScripts()
		a.injectSpireStartCronJob()
	}
	if runtime.GOOS == "windows" {
		a.createWindowsServerScripts()
		// TODO: Remove this when perl utility script is deprecated from world
		a.initWindowsWget()

		// ensure these exist
		a.setWindowsPerlPath()
		a.setWindowsMysqlPath()
	}

	a.installSpireBinary()
	a.initSpire()
	a.startSpire()

	a.setPostInstallConfigValues()

	a.runSharedMemory()
	a.runWorldForDatabaseUpdates()

	// TODO: add existing MySQL installation

	checkmark := "✅"
	if runtime.GOOS == "windows" {
		checkmark = "√"
	}

	a.logger.Println("")
	a.logger.Println("----------------------------------------")
	a.logger.Printf("| %s | Installation Complete (%v)\n", checkmark, FormatDuration(time.Since(a.totalTime)))
	a.logger.Println("----------------------------------------")
}

func (a *Installer) installLinuxOsPackages() {
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
		a.logger.Fatalf("could not get stdout pipe: %v", err)
	}
	cmd.Stderr = cmd.Stdout
	err = cmd.Start()
	if err != nil {
		a.logger.Error(err)
	}

	merged := io.MultiReader(stdout)
	scanner := bufio.NewScanner(merged)
	for scanner.Scan() {
		a.logger.Infoln(scanner.Text())
	}

	// install packages
	var packages []string
	distro := a.getLinuxDistribution()
	if distro == "ubuntu" {
		packages = getUbuntuPackages()
	} else if distro == "debian" {
		packages = getDebianPackages()
	} else {
		a.logger.Fatalf("Unsupported distribution: %v", distro)
	}

	// get specific version of distribution
	//version := a.getLinuxDistributionVersion()

	// apt-get install -y
	params = []string{"apt-get", "install", "-yq", "-m", "--no-install-recommends"}
	params = append(params, packages...)
	cmd = exec.CommandContext(ctx, "sudo", params...)
	cmd.Env = os.Environ()

	cmd.Dir = a.pathmanager.GetEQEmuServerPath()
	stdout, err = cmd.StdoutPipe()
	if err != nil {
		a.logger.Fatalf("could not get stdout pipe: %v", err)
	}

	cmd.Stderr = cmd.Stdout

	err = cmd.Start()
	if err != nil {
		a.logger.Error(err)
	}

	merged = io.MultiReader(stdout)
	scanner = bufio.NewScanner(merged)
	for scanner.Scan() {
		a.logger.Println(scanner.Text())
	}

	a.DoneBanner("Installing OS Packages")
}

func (a *Installer) Banner(s string) {
	a.logger.Println("")
	a.logger.Println("----------------------------------------")
	a.logger.Printf("> %v\n", s)
	a.logger.Println("----------------------------------------")
	//a.logger.Println("")
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

	a.logger.Println("----------------------------------------")
	a.logger.Printf("| %v | %v (%v)\n", checkmark, s, FormatDuration(time.Since(a.stepTime)))
	a.logger.Println("----------------------------------------")
}

func (a *Installer) initializeDirectories() {
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
		a.logger.Printf("Creating directory [%v]", dir)
		p := filepath.Join(a.pathmanager.GetEQEmuServerPath(), dir)
		err := os.MkdirAll(p, os.ModePerm)
		if err != nil {
			a.logger.Fatalf("could not create directory: %v", err)
		}
	}

	a.DoneBanner("Initializing Directories")
}

func (a *Installer) cloneEQEmuMaps() {
	a.Banner("Initializing Server Maps")

	if a.checkIfMapsAreUpToDate() {
		a.DoneBanner("Initializing Server Maps")
		return
	}

	// zip file path
	dumpZip := filepath.Join(os.TempDir(), "/maps.zip")

	// download the zip file
	err := download.WithProgress(
		dumpZip,
		"https://github.com/Akkadius/eqemu-maps/releases/latest/download/maps.zip",
	)
	if err != nil {
		a.logger.Fatalln(err)
	}

	// unzip the file
	mapsPath := filepath.Join(a.pathmanager.GetEQEmuServerPath(), "maps")
	a.logger.Infof("Downloaded zip to [%v]\n", dumpZip)
	err = unzip.New(dumpZip, mapsPath, a.logger).Extract()
	if err != nil {
		a.logger.Fatalf("could not extract zip: %v", err)
	}

	// remove the zip file
	err = os.Remove(dumpZip)
	if err != nil {
		a.logger.Fatalf("could not remove zip: %v", err)
	}

	a.DoneBanner("Initializing Server Maps")
}

func (a *Installer) clonePeqQuests() {
	a.Banner("Initializing Server Quests")

	a.logger.Infof("Cloning Quests from github.com/ProjectEQ/projecteqquests.git\n")

	// clone the repository
	repoPath := filepath.Join(a.pathmanager.GetEQEmuServerPath(), "quests")
	_, err := git.PlainClone(repoPath, false, &git.CloneOptions{
		URL:      "https://github.com/ProjectEQ/projecteqquests.git",
		Progress: a.logger.Writer(),
	})

	if err != nil && err != git.ErrRepositoryAlreadyExists {
		a.logger.Errorf("Could not clone quests [%v]\n", err)
	}

	// if the repository already exists, update it instead
	if err == git.ErrRepositoryAlreadyExists {
		a.logger.Infof("Quest repo already exists, skipping clone and updating instead\n")

		// open the repository
		r, err := git.PlainOpen(repoPath)
		if err != nil {
			a.logger.Fatalf("could not open repository: %v", err)
		}

		// Get the working directory for the repository
		w, err := r.Worktree()
		if err != nil {
			a.logger.Fatalf("could not get worktree: %v", err)
		}

		// Pull the latest changes from the origin remote and merge into the current branch
		err = w.Pull(&git.PullOptions{RemoteName: "origin", Progress: a.logger.Writer()})
		if err != nil && err != git.NoErrAlreadyUpToDate {
			a.logger.Fatalf("could not pull: %v", err)
		}

		a.logger.Infof("Quests updated successfully!\n")
	}

	a.DoneBanner("Initializing Server Quests")
}

type MysqlConfig struct {
	DatabaseName     string `json:"database_name"`
	DatabaseUser     string `json:"database_user"`
	DatabasePassword string `json:"database_password"`
	RootPassword     string `json:"root_password"`
}

func (a *Installer) initLinuxMysql() {
	a.Banner("Initializing MySQL")

	// change root password

	a.Exec(ExecConfig{command: "sudo", args: []string{"apt", "install", "-y", "-m", "mariadb-server"}})
	a.Exec(ExecConfig{command: "sudo", args: []string{"pkill", "-f", "-9", "mysql"}})

	time.Sleep(1 * time.Second)
	go func() {
		a.Exec(ExecConfig{command: "sudo", args: []string{"mysqld_safe", "--skip-grant-tables", "--skip-networking"}})
	}()
	time.Sleep(1 * time.Second)

	c := MysqlConfig{
		DatabaseName:     a.installConfig.MysqlDatabaseName,
		DatabaseUser:     a.installConfig.MysqlUsername,
		DatabasePassword: a.installConfig.MysqlPassword,
	}

	// create a new database
	var sql string

	a.logger.Infof("Creating database [%v]\n", c.DatabaseName)
	a.DbExec(DbExecConfig{statement: fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %v", c.DatabaseName)})

	// create a new user
	a.logger.Infof("Creating user [%v]\n", c.DatabaseUser)

	// grant privileges to the new user
	a.logger.Infof("Granting privileges to user [%v]\n", c.DatabaseUser)
	sql += fmt.Sprintf("CREATE USER IF NOT EXISTS '%v'@'localhost' IDENTIFIED BY '%v'; ", c.DatabaseUser, c.DatabasePassword)
	sql += fmt.Sprintf("GRANT ALL PRIVILEGES ON %v.* TO '%v'@'localhost'", c.DatabaseName, c.DatabaseUser)

	// flush privileges
	a.logger.Infoln("Flushing privileges")
	a.DbExec(DbExecConfig{statement: fmt.Sprintf("FLUSH PRIVILEGES; %v; FLUSH PRIVILEGES;", sql), hidestring: c.DatabasePassword})

	a.Exec(ExecConfig{command: "sudo", args: []string{"pkill", "-f", "-9", "mysql"}})
	a.Exec(ExecConfig{command: "sudo", args: []string{"service", "mariadb", "start"}})

	a.DoneBanner("Initializing MySQL")
}

type DbExecConfig struct {
	statement  string
	hidestring string
}

func (a *Installer) DbExec(c DbExecConfig) {
	mysqlPath := "mysql"
	if runtime.GOOS == "windows" {
		mysqlPath = filepath.Join(a.getWindowsMysqlPath(), "mysql.exe")
	}

	// check if mysql is installed
	if len(mysqlPath) == 0 {
		a.logger.Fatalf("could not find mysql executable")
	}

	a.Exec(
		ExecConfig{
			command:    mysqlPath,
			args:       []string{"-uroot", "-e", fmt.Sprintf("%v", c.statement)},
			hidestring: c.hidestring,
		},
	)
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

func (a *Installer) Exec(c ExecConfig) string {
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
		a.logger.Fatalf("could not get stdout pipe: %v", err)
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
		a.logger.Infof("Running command [%v%v]\n", c.command, argsPrint)
	}

	cmd.Stderr = cmd.Stdout
	err = cmd.Start()
	if err != nil {
		a.logger.Error(err)
	}

	// if we are detaching, release the process
	if c.detach {
		a.logger.Infof("Detaching process [%v%v]\n", c.command, argsPrint)
		err = cmd.Process.Release()
		if err != nil {
			a.logger.Error(err)
		}
		return ""
	}

	// merge stdout and stderr
	merged := io.MultiReader(stdout)
	scanner := bufio.NewScanner(merged)
	output := ""
	for scanner.Scan() {
		if len(c.dieonoutput) > 0 {
			if strings.Contains(scanner.Text(), c.dieonoutput) {
				a.logger.Infof("Found [%v] in output, exiting process", c.dieonoutput)
				_ = cmd.Process.Kill()
				break
			}
		}

		// don't print the output if we are silent
		// still return
		if !c.silent {
			a.logger.Infoln(scanner.Text())
		}
		output += scanner.Text() + "\n"
	}

	return output
}

func (a *Installer) cloneEQEmuSource() {
	a.Banner("Initializing Server Source")

	a.logger.Infof("Cloning from https://github.com/EQEmu/Server.git\n")

	// clone the repository
	repoPath := a.installConfig.CodePath
	_, err := git.PlainClone(repoPath, false, &git.CloneOptions{
		URL:      "https://github.com/EQEmu/Server.git",
		Progress: a.logger.Writer(),
	})

	// if the repository already exists, update it instead
	if err != nil && err != git.ErrRepositoryAlreadyExists {
		a.logger.Errorf("Could not clone quests [%v]\n", err)
	}

	// if the repository already exists, update it instead
	if err == git.ErrRepositoryAlreadyExists {
		a.logger.Infof("repo already exists, skipping clone and updating instead\n")

		// open the repository
		r, err := git.PlainOpen(repoPath)
		if err != nil {
			a.logger.Fatalf("could not open repository: %v", err)
		}

		// Get the working directory for the repository
		w, err := r.Worktree()
		if err != nil {
			a.logger.Fatalf("could not get worktree: %v", err)
		}

		// Pull the latest changes from the origin remote and merge into the current branch
		err = w.Pull(&git.PullOptions{RemoteName: "origin", Progress: a.logger.Writer()})
		if err != nil && err != git.NoErrAlreadyUpToDate {
			a.logger.Fatalf("could not pull: %v", err)
		}

		a.logger.Infof("repo updated successfully!\n")
	}

	a.DoneBanner("Initializing Server Source")
}

func (a *Installer) initializeServerConfig() {
	a.Banner("Initializing Server Config")

	// check if eqemu_config.json exists
	if _, err := os.Stat(a.pathmanager.GetEQEmuServerConfigFilePath()); err == nil {
		a.logger.Infof("eqemu_config.json already exists, skipping\n")
		return
	}

	// download the config file
	res, err := http.Get("https://raw.githubusercontent.com/Akkadius/eqemu-install-v2/master/eqemu_config_docker.json")
	if err != nil {
		a.logger.Fatalln(err)
	}

	// read the response body
	b, err := io.ReadAll(res.Body)
	if err != nil {
		a.logger.Fatalln(err)
	}

	// unmarshal the json
	var c eqemuserverconfig.EQEmuConfigJson
	err = json.Unmarshal(b, &c)
	if err != nil {
		a.logger.Fatalln(err)
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
		a.logger.Fatalln(err)
	}
	c.Spire.HttpPort = spireWebPort

	// save the config file
	err = a.config.Save(c)
	if err != nil {
		a.logger.Fatalln(err)
	}

	a.logger.Infof("Saved config to [%v]\n", a.pathmanager.GetEQEmuServerConfigFilePath())

	a.DoneBanner("Initializing Server Config")
}

func (a *Installer) sourcePeqDatabase() {
	a.Banner("Sourcing ProjectEQ Database")

	mysqlPath := "mysql"
	if runtime.GOOS == "windows" {
		mysqlPath = filepath.Join(a.getWindowsMysqlPath(), "mysql.exe")
	}

	tables := a.Exec(
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

	// get the table count
	// subtract 1 because it is the output header
	tableCount := len(strings.Split(tables, "\n")) - 1

	a.logger.Infof(
		"Database [%v] has [%v] tables\n",
		a.installConfig.MysqlDatabaseName,
		tableCount,
	)

	if len(strings.Split(tables, "\n")) > 200 {
		// database already exists, skip
		a.logger.Infof(
			"Database [%v] already exists with [%v] tables, skipping source\n",
			a.installConfig.MysqlDatabaseName,
			tableCount,
		)
		return
	}

	// zip file path
	dumpZip := filepath.Join(os.TempDir(), "/dump/peq.zip")

	// Create the temp folder
	a.logger.Infof("Creating directory [%v]\n", filepath.Dir(dumpZip))
	err := os.MkdirAll(filepath.Dir(dumpZip), os.ModePerm)
	if err != nil {
		a.logger.Fatalf("could not create directory: %v", err)
	}

	// download the latest database dump
	err = download.WithProgress(
		dumpZip,
		"http://db.projecteq.net/api/v1/dump/latest",
	)
	if err != nil {
		a.logger.Fatalln(err)
	}

	a.logger.Infof("Downloaded zip to [%v]\n", dumpZip)
	err = unzip.New(dumpZip, filepath.Join(os.TempDir(), "/dump"), a.logger).Extract()
	if err != nil {
		a.logger.Fatalf("could not extract zip: %v", err)
	}

	a.logger.Infof("Extracted zip to [%v]\n", filepath.Join(os.TempDir(), "/dump/peq-dump"))

	extractPath := filepath.Join(os.TempDir(), "/dump/peq-dump")

	a.logger.Infof("Sourcing database dump from [%v]\n", extractPath)

	a.Exec(
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

	a.logger.Infof("Sourced database dump from [%v]\n", extractPath)

	a.Exec(
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

	// cleanup the temp folder
	a.logger.Infof("|-- Cleaning up temp folder [%v]\n", filepath.Join(os.TempDir(), "/dump"))
	err = os.RemoveAll(filepath.Join(os.TempDir(), "/dump"))
	if err != nil {
		a.logger.Fatalf("could not remove directory: %v", err)
	}

	a.DoneBanner("Sourcing ProjectEQ Database")
}

func (a *Installer) installBinaries() {
	a.Banner("Installing EQEmu Server Binaries")

	// download the latest binaries
	tempPath := filepath.Join(os.TempDir(), "eqemu-server.zip")
	a.logger.Infof("Downloading binaries to [%v]\n", tempPath)
	err := download.WithProgress(
		tempPath,
		fmt.Sprintf("https://github.com/eqemu/server/releases/latest/download/eqemu-server-%v-x64.zip", runtime.GOOS),
	)
	if err != nil {
		a.logger.Fatalln(err)
	}

	// extract the zip
	extractTo := a.pathmanager.GetEQEmuServerBinPath()
	a.logger.Infof("Extracting zip to [%v]\n", extractTo)
	err = unzip.New(tempPath, extractTo, a.logger).Extract()
	if err != nil {
		a.logger.Fatalf("could not extract zip: %v", err)
	}

	// cleanup the temp folder
	a.logger.Infof("|-- Cleaning up temp folder [%v]\n", tempPath)
	err = os.RemoveAll(tempPath)
	if err != nil {
		a.logger.Fatalf("could not remove directory: %v", err)
	}

	// make the binaries executable
	a.logger.Infof("Making binaries executable\n")

	// loop through files in the bin folder
	err = filepath.Walk(extractTo, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// skip directories
		if info.IsDir() {
			return nil
		}

		a.logger.Infof("|-- Making [%v] executable\n", path)

		// make the file executable
		err = os.Chmod(path, 0755)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		a.logger.Fatalf("could not make binaries executable: %v", err)
	}

	a.DoneBanner("Installing EQEmu Server Binaries")
}

func (a *Installer) symlinkPatchFiles() {
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
			a.logger.Infof("Symlink [%v] already exists, skipping\n", symlinkPath)

			// remove the symlink
			err = os.Remove(symlinkPath)
			if err != nil {
				a.logger.Fatalf("could not remove symlink: %v", err)
			}
		}

		sourcePatchPath := filepath.Join(a.installConfig.CodePath, "utils", "patches", patchFile)

		// create the symlink
		a.logger.Infof("Creating symlink [%v] -> [%v]\n", symlinkPath, sourcePatchPath)
		err := os.Symlink(sourcePatchPath, symlinkPath)
		if err != nil {
			a.logger.Fatalf("could not create symlink: %v", err)
		}
	}

	a.DoneBanner("Symlinking Patch Files")
}

func (a *Installer) symlinkOpcodeFiles() {
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

		// check if the symlink exists
		if _, err := os.Stat(symlinkPath); !os.IsNotExist(err) {
			a.logger.Infof("Symlink [%v] already exists, skipping\n", symlinkPath)

			// remove the symlink
			err = os.Remove(symlinkPath)
			if err != nil {
				a.logger.Fatalf("could not remove symlink: %v", err)
			}
		}

		sourcePatchPath := filepath.Join(a.installConfig.CodePath, "utils", "patches", opcodeFile)

		// create the symlink
		a.logger.Infof("Creating symlink [%v] -> [%v]\n", symlinkPath, sourcePatchPath)
		err := os.Symlink(sourcePatchPath, symlinkPath)
		if err != nil {
			a.logger.Fatalf("could not create symlink: %v", err)
		}
	}

	a.DoneBanner("Symlinking Opcode Files")
}

func (a *Installer) symlinkLoginOpcodeFiles() {
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

		// check if the symlink exists
		if _, err := os.Stat(symlinkPath); !os.IsNotExist(err) {
			a.logger.Infof("Symlink [%v] already exists, skipping\n", symlinkPath)

			// remove the symlink
			err = os.Remove(symlinkPath)
			if err != nil {
				a.logger.Fatalf("could not remove symlink: %v", err)
			}
		}

		// TODO: this is temp I will go back and fix this later
		sourcePatchPath := filepath.Join(os.TempDir(), "code", "loginserver", "login_util", opcodeFile)

		// create the symlink
		a.logger.Infof("Creating symlink [%v] -> [%v]\n", symlinkPath, sourcePatchPath)
		err := os.Symlink(sourcePatchPath, symlinkPath)
		if err != nil {
			a.logger.Fatalf("could not create symlink: %v", err)
		}
	}

	a.DoneBanner("Symlinking Login Opcode Files")
}

func (a *Installer) symlinkPluginsAndModules() {
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
	a.logger.Infof("Creating symlink [%v] -> [%v]\n", targetLuaModules, sourceLuaModules)
	err := os.Symlink(sourceLuaModules, targetLuaModules)
	if err != nil {
		a.logger.Fatalf("could not create symlink: %v", err)
	}

	// create the symlink
	a.logger.Infof("Creating symlink [%v] -> [%v]\n", targetPerlPlugins, sourcePerlPlugins)
	err = os.Symlink(sourcePerlPlugins, targetPerlPlugins)
	if err != nil {
		a.logger.Fatalf("could not create symlink: %v", err)
	}

	a.DoneBanner("Symlinking Plugins and Modules")
}

func (a *Installer) GetRandomPassword() string {
	p, err := password.Generate(32, 10, 0, false, false)
	if err != nil {
		a.logger.Fatalf("could not generate random password: %v", err)
	}

	return p
}

func (a *Installer) checkIfMapsAreUpToDate() bool {
	type Release struct {
		TagName string `json:"tag_name"`
	}

	// get latest release version
	resp, err := http.Get("https://api.github.com/repos/Akkadius/eqemu-maps/releases/latest")
	if err != nil {
		a.logger.Fatalf("could not get latest release version: %v", err)
	}

	defer resp.Body.Close()

	// read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		a.logger.Fatalf("could not read response body: %v", err)
	}

	// bind body to struct
	var release Release
	err = json.Unmarshal(body, &release)
	if err != nil {
		a.logger.Fatalf("could not unmarshal response body: %v", err)
	}

	a.logger.Infof("Downloading eqemu-maps release\n")

	type PackageJson struct {
		Version string `json:"version"`
	}

	// get current version from package.json
	file := filepath.Join(a.pathmanager.GetEQEmuServerPath(), "maps", "package.json")

	// check if file exists
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}

	// read file package.json contents into PackageJson struct
	packageJson, err := os.ReadFile(file)
	if err != nil {
		a.logger.Fatalf("could not read PackageJson file: %v", err)
	}

	// bind package.json to struct
	var packageJsonStruct PackageJson
	err = json.Unmarshal(packageJson, &packageJsonStruct)
	if err != nil {
		a.logger.Fatalf("could not unmarshal package.json: %v", err)
	}

	// check if current version is the same as the latest release version
	remoteVersion := strings.ReplaceAll(release.TagName, "v", "")

	if len(remoteVersion) == 0 {
		a.logger.Infof("Could not retrieve latest [eqemu-maps] version, possibly rate limited, skipping\n")
		return true
	}

	if len(remoteVersion) > 0 && packageJsonStruct.Version == remoteVersion {
		a.logger.Infof("Maps are up to date on version v%v\n", packageJsonStruct.Version)
		return true
	}

	return false
}

func (a *Installer) setInstallerPath() {
	cwd, err := os.Getwd()
	if err != nil {
		a.logger.Fatalf("could not get current working directory: %v", err)
	}
	a.pathmanager.SetServerPath(cwd)
}

func (a *Installer) createLinuxServerScripts() {
	a.Banner("Creating Server Scripts")

	// create a map of scripts
	serverScripts := map[string]string{
		"start": "bash -c \"while true; do nohup $(find ./bin -name 'occulus*' | head -1) server-launcher >/dev/null 2>&1; sleep 1; done &\" && echo Server started",
		"stop":  "$(find ./bin -name 'occulus*' | head -1) stop-server; echo \"Server stopped\"",
	}

	for s := range serverScripts {

		// get the f name
		file := filepath.Join(a.pathmanager.GetEQEmuServerPath(), s)

		a.logger.Infof("Creating script [%v]\n", file)

		// create file
		f, err := os.Create(file)
		if err != nil {
			a.logger.Fatalf("could not create f: %v", err)
		}

		// write contents to f
		contents := fmt.Sprintf("#!/usr/bin/env bash\n%v\n", serverScripts[s])
		_, err = f.WriteString(contents)
		if err != nil {
			a.logger.Fatalf("could not write to f: %v", err)
		}

		// close file
		_ = f.Close()

		a.logger.Infof("|-- Making file [%v] executable\n", file)

		// make file executable
		err = os.Chmod(file, 0755)
		if err != nil {
			a.logger.Fatalf("could not chmod f: %v", err)
		}
	}

	a.DoneBanner("Creating Server Scripts")
}

func (a *Installer) injectSpireStartCronJob() {
	a.Banner("Injecting Spire Start Cron Job")

	// get spire path
	spirePath := filepath.Join(a.pathmanager.GetEQEmuServerPath(), "spire")

	cronInject := fmt.Sprintf("while true; do nohup %s > %s/logs/spire.log 2>&1; sleep 1; done &", spirePath, a.pathmanager.GetEQEmuServerPath())

	a.Exec(
		ExecConfig{
			command: "bash",
			args: []string{
				"-c",
				fmt.Sprintf("crontab -l | grep -qF 'spire' || (crontab -l 2>/dev/null; echo \"@reboot %v\") | crontab -", cronInject),
			},
		},
	)

	a.DoneBanner("Injecting Spire Start Cron Job")
}

func (a *Installer) installSpireBinary() {
	a.Banner("Installing Spire")

	// check if spire is already installed
	spirePath, err := exec.LookPath(filepath.Join(a.pathmanager.GetEQEmuServerPath(), "spire"))
	if err != nil {
		a.logger.Infof("could not find spire binary: %v", err)
	}

	if _, err := os.Stat(spirePath); !os.IsNotExist(err) {
		// spire is already installed
		a.logger.Infof("Spire is already installed at [%v]\n", spirePath)
		a.DoneBanner("Installing Spire")
		return
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
			a.logger.Infof("could not get latest release: %v", err)
			return
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
					a.logger.Fatalf("could not download spire: %v", err)
				}

				// unzip
				tempFileZipped := fmt.Sprintf("%s/%s", os.TempDir(), targetFileNameZipped)
				unzipPathTemp := filepath.Join(fmt.Sprintf("%s/spire-download", os.TempDir()))
				uz := unzip.New(tempFileZipped, unzipPathTemp, a.logger)
				a.logger.Infof("|-- Unzipping file [%v] to [%v]\n", tempFileZipped, a.pathmanager.GetEQEmuServerPath())
				err = uz.Extract()
				if err != nil {
					a.logger.Fatalf("could not unzip spire: %v", err)
				}

				// rename
				src, err := exec.LookPath(filepath.Join(unzipPathTemp, targetFileName))
				if err != nil {
					a.logger.Fatalf("could not find spire: %v", err)
				}

				// new spire path
				newSpirePath := filepath.Join(a.pathmanager.GetEQEmuServerPath(), "spire")
				if runtime.GOOS == "windows" {
					newSpirePath = filepath.Join(a.pathmanager.GetEQEmuServerPath(), "spire.exe")
				}

				// destination
				dst := newSpirePath

				a.logger.Infof("|-- Renaming file [%v] to [%v]\n", src, dst)

				// copy file from src to dst
				a.Copy(src, dst)

				a.logger.Infof("|-- Making file [%v] executable\n", dst)

				// make executable
				err = os.Chmod(dst, 0755)
				if err != nil {
					a.logger.Fatalf("could not chmod spire: %v", err)
				}
			}
		}
	}

	a.DoneBanner("Installing Spire")
}

func (a *Installer) initSpire() {
	a.Banner("Initializing Spire")

	spirePath, err := exec.LookPath(filepath.Join(a.pathmanager.GetEQEmuServerPath(), "spire"))
	if err != nil {
		a.logger.Fatalf("could not find spire binary: %v", err)
	}

	a.Exec(ExecConfig{
		command: spirePath,
		args: []string{
			"spire:init",
			a.installConfig.SpireAdminUser,
			a.installConfig.SpireAdminPassword,
		},
		hidestring: a.installConfig.SpireAdminPassword,
	})

	a.Exec(ExecConfig{
		command:    spirePath,
		args:       []string{"spire:occulus-update"},
		hidestring: a.installConfig.SpireAdminPassword,
	})

	a.DoneBanner("Initializing Spire")
}

func (a *Installer) runSharedMemory() {
	a.Banner("Running Shared Memory")

	a.Exec(ExecConfig{
		execpath: a.pathmanager.GetEQEmuServerPath(),
		command:  filepath.Join("bin", "shared_memory"),
	})

	a.DoneBanner("Running Shared Memory")
}

func (a *Installer) runWorldForDatabaseUpdates() {
	a.Banner("Running World for Database Updates")

	// TODO: Windows is not aware of Perl yet at this stage because the PATH is not updated

	a.Exec(ExecConfig{
		execpath:    a.pathmanager.GetEQEmuServerPath(),
		command:     filepath.Join("bin", "world"),
		dieonoutput: "Server (TCP) listener started on port",
	})

	a.DoneBanner("Running World for Database Updates")
}

func (a *Installer) startSpire() {
	a.Banner("Starting Spire")

	// kill any running spire processes
	processes, _ := process.Processes()
	for _, p := range processes {
		cmdline, err := p.Cmdline()
		if err != nil {
			a.logger.Infof("could not get cmdline for process: %v", err)
			continue
		}

		// kill spire if it's running
		// ignore spire processes that are running from /tmp
		installerRan := strings.Contains(cmdline, "/tmp/") || strings.Contains(cmdline, "-install")
		if strings.Contains(cmdline, "spire") && !installerRan {
			err := p.Kill()
			if err != nil {
				a.logger.Errorf("could not kill process: %v", err)
			}
		}
	}

	a.logger.Infof("starting spire [%v]", filepath.Join(a.pathmanager.GetEQEmuServerPath(), "spire"))

	// start spire in a loop
	if runtime.GOOS == "linux" {
		a.Exec(ExecConfig{
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
	}

	if runtime.GOOS == "windows" {
		spirePath, err := exec.LookPath(filepath.Join(a.pathmanager.GetEQEmuServerPath(), "spire"))
		if err != nil {
			a.logger.Fatalf("could not find spire binary: %v", err)
		}

		pp.Println(fmt.Sprintf(
			"start /b %s > %s/logs/spire.log 2>&1",
			spirePath,
			a.pathmanager.GetEQEmuServerPath(),
		))

		a.Exec(ExecConfig{
			command: "cmd",
			args: []string{
				"/c",
				fmt.Sprintf(
					"start /b %s > %s/logs/spire.log 2>&1",
					spirePath,
					a.pathmanager.GetEQEmuServerPath(),
				),
			},
			detach: true,
		})
	}

	a.DoneBanner("Starting Spire")
}

// getLinuxDistribution returns the linux distribution
func (a *Installer) getLinuxDistribution() string {
	// determine whether we're on ubuntu or debian
	// read from /etc/os-release to determine which distro we're on
	if _, err := os.Stat("/etc/os-release"); os.IsNotExist(err) {
		a.logger.Fatalf("could not find /etc/os-release")
	}

	// get contents of /etc/os-release
	osRelease, err := os.ReadFile("/etc/os-release")
	if err != nil {
		a.logger.Fatalf("could not read /etc/os-release: %v", err)
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
				return "ubuntu"
			} else if strings.Contains(value, "Debian") {
				return "debian"
			} else {
				a.logger.Fatalf("unknown OS: %s", value)
				return "unknown"
			}
		}
	}

	return "unknown"
}

// getLinuxDistributionVersion returns the linux distribution version
func (a *Installer) getLinuxDistributionVersion() string {
	// read from /etc/os-release to determine which version of the distro we're on
	if _, err := os.Stat("/etc/os-release"); os.IsNotExist(err) {
		a.logger.Fatalf("could not find /etc/os-release")
	}

	// get contents of /etc/os-release
	osRelease, err := os.ReadFile("/etc/os-release")
	if err != nil {
		a.logger.Fatalf("could not read /etc/os-release: %v", err)
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
			return value
		}
	}

	return "unknown"
}

func (a *Installer) initWindowsMysql() {
	a.Banner("Downloading MariaDB")

	// check if a.getWindowsMysqlPath() exists
	if _, err := os.Stat(a.getWindowsMysqlPath()); err == nil {
		a.logger.Infof("MySQL already installed, skipping")
		return
	}

	// download mariadb
	// download the latest binaries
	tempPath := filepath.Join(os.TempDir(), "mariadb.msi")
	a.logger.Infof("Downloading binaries to [%v]\n", tempPath)
	err := download.WithProgress(
		tempPath,
		"https://github.com/Akkadius/eqemu-install-v2/releases/download/static/mariadb-10.11.2-winx64.msi",
	)
	if err != nil {
		a.logger.Fatalln(err)
	}

	// install mariadb
	// start /wait msiexec /i mariadb-10.0.21-winx64.msi SERVICENAME=MySQL PORT=3306 PASSWORD=eqemu /qn
	// TODO: make port configurable
	// TODO: split out root user and eqemu user passwords
	a.logger.Infof("Installing MariaDB")
	a.Exec(ExecConfig{
		command: "msiexec",
		args: []string{
			"/i",
			tempPath,
			"SERVICENAME=MySQL",
			fmt.Sprintf("PORT=%v", a.installConfig.MysqlPort),
			"BUFFERPOOLSIZE=1024M",
			fmt.Sprintf("PASSWORD=%s", a.installConfig.MysqlPassword),
			"/qn",
		},
	})

	c := MysqlConfig{
		DatabaseName:     a.installConfig.MysqlDatabaseName,
		DatabaseUser:     a.installConfig.MysqlUsername,
		DatabasePassword: a.installConfig.MysqlPassword,
	}

	// create a new database
	var sql string

	a.logger.Infof("Creating database [%v]\n", c.DatabaseName)
	a.DbExec(DbExecConfig{statement: fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %v", c.DatabaseName)})

	// create a new user
	a.logger.Infof("Creating user [%v]\n", c.DatabaseUser)

	// grant privileges to the new user
	a.logger.Infof("Granting privileges to user [%v]\n", c.DatabaseUser)
	sql += fmt.Sprintf("CREATE USER IF NOT EXISTS '%v'@'localhost' IDENTIFIED BY '%v'; ", c.DatabaseUser, c.DatabasePassword)
	sql += fmt.Sprintf("GRANT ALL PRIVILEGES ON %v.* TO '%v'@'localhost'", c.DatabaseName, c.DatabaseUser)

	// flush privileges
	a.logger.Infoln("Flushing privileges")
	a.DbExec(DbExecConfig{statement: fmt.Sprintf("FLUSH PRIVILEGES; %v; FLUSH PRIVILEGES;", sql), hidestring: c.DatabasePassword})

	a.setWindowsMysqlPath()

	a.DoneBanner("Downloading MariaDB")
}

func (a *Installer) initWindowsPerl() {
	a.Banner("Downloading Perl")

	// check if a.getWindowsPerlPath() exists
	if _, err := os.Stat(a.getWindowsPerlPath()); err == nil {
		a.logger.Infof("Perl already installed, skipping")
		return
	}

	// download mariadb
	// download the latest binaries
	tempPath := filepath.Join(os.TempDir(), "perl.msi")
	a.logger.Infof("Downloading binaries to [%v]\n", tempPath)
	err := download.WithProgress(
		tempPath,
		"https://github.com/Akkadius/eqemu-install-v2/releases/download/static/strawberry-perl-5.24.4.1-64bit.msi",
	)
	if err != nil {
		a.logger.Fatalln(err)
	}

	// install perl
	// start /wait msiexec /i strawberry-perl-5.24.4.1-64bit.msi PERL_PATH="Yes" /q
	// start /wait msiexec /i mariadb-10.0.21-winx64.msi SERVICENAME=MySQL PORT=3306 PASSWORD=eqemu /qn
	a.logger.Infof("Installing Perl")
	a.Exec(ExecConfig{
		command: "msiexec",
		args: []string{
			"/i",
			tempPath,
			"PERL_PATH=Yes",
			"/q",
		},
	})

	a.setWindowsPerlPath()

	a.DoneBanner("Downloading Perl")
}

// initWindowsWget downloads wget for windows (backwards compatibility)
// TODO: remove this in the future
func (a *Installer) initWindowsWget() {
	a.Banner("Downloading Windows wget")

	downloadPath := filepath.Join(a.pathmanager.GetEQEmuServerBinPath(), "wget.exe")
	a.logger.Infof("Downloading binaries to [%v]\n", downloadPath)
	err := download.WithProgress(
		downloadPath,
		"https://github.com/Akkadius/eqemu-install-v2/releases/download/static/wget.exe",
	)
	if err != nil {
		a.logger.Fatalln(err)
	}

	a.DoneBanner("Downloading Windows wget")
}

// initWindowsMysqlService installs and configures the mysql service
func (a *Installer) getWindowsProgramFilesPath() string {
	// get program files path
	return strings.TrimSpace(a.Exec(ExecConfig{command: "cmd", args: []string{"/c", "echo", "%programfiles%"}, silent: true}))
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
func (a *Installer) Copy(src string, dst string) {
	// read the file
	bytesRead, err := os.ReadFile(src)
	if err != nil {
		a.logger.Fatal(err)
	}

	// write the file
	err = os.WriteFile(dst, bytesRead, os.ModePerm)
	if err != nil {
		a.logger.Fatal(err)
	}
}

// createWindowsServerScripts creates the server scripts
func (a *Installer) createWindowsServerScripts() {
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

		a.logger.Infof("Creating script [%v]\n", file)

		// create file
		f, err := os.Create(file)
		if err != nil {
			a.logger.Fatalf("could not create f: %v", err)
		}

		// write contents to f
		contents := fmt.Sprintf("%v\n", serverScripts[s])
		_, err = f.WriteString(contents)
		if err != nil {
			a.logger.Fatalf("could not write to f: %v", err)
		}

		// close file
		_ = f.Close()

		a.logger.Infof("|-- Making file [%v] executable\n", file)

		// make file executable
		err = os.Chmod(file, 0755)
		if err != nil {
			a.logger.Fatalf("could not chmod f: %v", err)
		}
	}

	a.DoneBanner("Creating Server Scripts")
}

func (a *Installer) initWindowsCommandPrompt() {
	cmd := exec.Command("chcp", "65001")
	cmd.Env = os.Environ()
	cmd.Dir = a.pathmanager.GetEQEmuServerPath()
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		a.logger.Fatalf("could not get stdout pipe: %v", err)
	}
	err = cmd.Run()
	if err != nil {
		a.logger.Error(err)
	}

	fmt.Println(stdout)
}

func (a *Installer) setWindowsMysqlPath() {
	// check if path already contains mysql
	if !strings.Contains(os.Getenv("Path"), a.getWindowsMysqlPath()) {
		a.logger.Infof("Updating PATH to include [%v]\n", a.getWindowsMysqlPath())
		err := os.Setenv("Path", fmt.Sprintf("%v;%v", os.Getenv("Path"), a.getWindowsMysqlPath()))
		if err != nil {
			a.logger.Fatalln(err)
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
			a.logger.Error(err)
		}
	}
}

func (a *Installer) setWindowsPerlPath() {
	// update current notion of path temporarily since we don't have the updated path in the current cmd shell
	if !strings.Contains(os.Getenv("Path"), a.getWindowsPerlPath()) {
		a.logger.Infof("Updating PATH to include [%v]\n", a.getWindowsPerlPath())
		err := os.Setenv("Path", fmt.Sprintf("%v;%v", os.Getenv("Path"), a.getWindowsPerlPath()))
		if err != nil {
			a.logger.Fatalln(err)
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
			a.logger.Error(err)
		}
	}
}

func (a *Installer) setPostInstallConfigValues() {
	// load the config
	config := a.config.Get()

	// set the post install config values
	if config.WebAdmin != nil {
		config.WebAdmin.Launcher.MinZoneProcesses = 10
		config.WebAdmin.Launcher.RunSharedMemory = true
		// boat zones mainly
		config.WebAdmin.Launcher.StaticZones = "butcher,erudnext,freporte,qeynos,freeporte,oot,iceclad,nro,oasis,nedaria,abysmal,natimbi,timorous,abysmal,firiona,overthere"
	}

	err := a.config.Save(config)
	if err != nil {
		a.logger.Fatalf("could not save config: %v", err)
	}
}
