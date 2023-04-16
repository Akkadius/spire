package eqemuserver

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/Akkadius/spire/internal/download"
	"github.com/Akkadius/spire/internal/eqemuserverconfig"
	"github.com/Akkadius/spire/internal/pathmgmt"
	"github.com/Akkadius/spire/internal/unzip"
	"github.com/go-git/go-git/v5"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

const mapsRelease = "v1.0.0"

type Installer struct {
	pathmanager *pathmgmt.PathManagement
	config      *eqemuserverconfig.Config
	logger      *logrus.Logger
	stepTime    time.Time
	totalTime   time.Time
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
		logger:      logger,
		pathmanager: pathmanager,
		config:      eqemuserverconfig.NewConfig(logger, pathmanager),
	}

	cwd, err := os.Getwd()
	if err != nil {
		logger.Fatalf("could not get current working directory: %v", err)
	}
	i.pathmanager.SetServerPath(cwd)

	return i
}

func (a *Installer) Install() {
	// install prompt library for installation
	// install debian packages
	// install ubuntu packages (for ubuntu)
	a.totalTime = time.Now()
	a.installOsPackages()
	a.initMySQL()
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

	a.logger.Println("")
	a.logger.Println("----------------------------------------")
	a.logger.Printf("| ✅ | Installation Complete (%v)\n", FormatDuration(time.Since(a.totalTime)))
	a.logger.Println("----------------------------------------")
}

func (a *Installer) installOsPackages() {
	a.Banner("Installing OS Packages")

	// apt-get update
	params := []string{"apt-get", "update"}
	cmd := exec.Command("sudo", params...)
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

	// apt-get install -y
	params = []string{"apt-get", "install", "-y", "-m"}
	params = append(params, getDebianPackages()...)
	cmd = exec.Command("sudo", params...)
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
	a.logger.Println("----------------------------------------")
	a.logger.Printf("| ✅ | %v (%v)\n", s, FormatDuration(time.Since(a.stepTime)))
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
		path := filepath.Join(a.pathmanager.GetEQEmuServerPath(), dir)
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			a.logger.Fatalf("could not create directory: %v", err)
		}
	}

	a.DoneBanner("Initializing Directories")
}

func (a *Installer) cloneEQEmuMaps() {
	a.Banner("Initializing Server Maps")

	a.logger.Infof("Downloading EQEmuMaps release %v\n", mapsRelease)

	// zip file path
	dumpZip := filepath.Join(os.TempDir(), "/maps.zip")

	// download the zip file
	err := download.WithProgress(
		dumpZip,
		"https://github.com/Akkadius/EQEmuMaps/releases/download/v1.0.0/maps.zip",
	)
	if err != nil {
		a.logger.Fatalln(err)
	}

	mapsPath := filepath.Join(a.pathmanager.GetEQEmuServerPath(), "maps")
	a.logger.Infof("Downloaded zip to [%v]\n", dumpZip)
	err = unzip.New(dumpZip, mapsPath).Extract()
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
	path := filepath.Join(a.pathmanager.GetEQEmuServerPath(), "quests")
	_, err := git.PlainClone(path, false, &git.CloneOptions{
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
		r, err := git.PlainOpen(path)
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

func (a *Installer) initMySQL() {
	a.Banner("Initializing MySQL")

	// change root password
	a.Exec("sudo", []string{"apt", "install", "-y", "-m", "mariadb-server"})
	a.Exec("sudo", []string{"pkill", "-f", "-9", "mysql"})
	time.Sleep(1 * time.Second)
	go func() {
		a.Exec("sudo", []string{"mysqld_safe", "--skip-grant-tables", "--skip-networking"})
	}()
	time.Sleep(1 * time.Second)
	//a.Exec("sudo", []string{"bash", "-c", "ps aux | grep mysql"})

	c := MysqlConfig{
		DatabaseName:     "peq",
		DatabaseUser:     "peq",
		DatabasePassword: "peq",
	}

	// create a new database

	var sql string

	a.logger.Infof("Creating database [%v]\n", c.DatabaseName)
	a.DbExec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %v", c.DatabaseName))

	// create a new user
	a.logger.Infof("Creating user [%v]\n", c.DatabaseUser)

	// grant privileges to the new user
	a.logger.Infof("Granting privileges to user [%v]\n", c.DatabaseUser)
	sql += fmt.Sprintf("CREATE USER IF NOT EXISTS '%v'@'localhost' IDENTIFIED BY '%v';", c.DatabaseUser, c.DatabasePassword)
	sql += fmt.Sprintf("GRANT ALL PRIVILEGES ON %v.* TO '%v'@'localhost';", c.DatabaseName, c.DatabaseUser)

	// flush privileges
	a.logger.Infoln("Flushing privileges")
	a.DbExec(fmt.Sprintf("FLUSH PRIVILEGES; %v", sql))

	a.Exec("sudo", []string{"pkill", "-f", "-9", "mysql"})
	a.Exec("sudo", []string{"service", "mariadb", "start"})

	a.DoneBanner("Initializing MySQL")
}

func (a *Installer) DbExec(statement string) {
	a.Exec("mysql", []string{"-uroot", "-e", fmt.Sprintf("%v", statement)})
}

func (a *Installer) Exec(command string, args []string) {
	cmd := exec.Command(command, args...)
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
}

func (a *Installer) ExecPath(path string, command string, args []string) {
	cmd := exec.Command(command, args...)
	cmd.Env = os.Environ()
	cmd.Dir = path
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
}

func (a *Installer) cloneEQEmuSource() {
	a.Banner("Initializing Server Source")

	a.logger.Infof("Cloning from https://github.com/EQEmu/Server.git\n")

	// clone the repository
	// TODO: make this configurable
	path := filepath.Join(os.TempDir(), "code")
	_, err := git.PlainClone(path, false, &git.CloneOptions{
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
		r, err := git.PlainOpen(path)
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

	// TODO: make this configurable later
	c.Server.Database.Host = "127.0.0.1"
	c.Server.Database.Port = "3306"
	c.Server.Database.Username = "peq"
	c.Server.Database.Password = "peq"
	c.Server.Database.Db = "peq"

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
	err = unzip.New(dumpZip, filepath.Join(os.TempDir(), "/dump")).Extract()
	if err != nil {
		a.logger.Fatalf("could not extract zip: %v", err)
	}

	a.logger.Infof("Extracted zip to [%v]\n", filepath.Join(os.TempDir(), "/dump/peq-dump"))

	extractPath := filepath.Join(os.TempDir(), "/dump/peq-dump")
	a.logger.Infof("Sourcing database dump from [%v]\n", extractPath)
	a.ExecPath(extractPath, "mysql", []string{"-upeq", "-ppeq", "peq", "-e", "source create_all_tables.sql"})
	a.logger.Infof("Sourced database dump from [%v]\n", extractPath)
	a.Exec("mysql", []string{"-upeq", "-ppeq", "peq", "-e", "show tables"})

	// cleanup the temp folder
	a.logger.Infof("Cleaning up temp folder [%v]\n", filepath.Join(os.TempDir(), "/dump"))
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
		"https://github.com/eqemu/server/releases/latest/download/eqemu-server-linux-x64.zip",
	)
	if err != nil {
		a.logger.Fatalln(err)
	}

	// extract the zip
	extractTo := a.pathmanager.GetEQEmuServerBinPath()
	a.logger.Infof("Extracting zip to [%v]\n", extractTo)
	err = unzip.New(tempPath, extractTo).Extract()
	if err != nil {
		a.logger.Fatalf("could not extract zip: %v", err)
	}

	// cleanup the temp folder
	a.logger.Infof("Cleaning up temp folder [%v]\n", tempPath)
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

		a.logger.Infof("Making [%v] executable\n", path)

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

		// TODO: this is temp I will go back and fix this later
		sourcePatchPath := filepath.Join(os.TempDir(), "code", "utils", "patches", patchFile)

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

		// TODO: this is temp I will go back and fix this later
		sourcePatchPath := filepath.Join(os.TempDir(), "code", "utils", "patches", opcodeFile)

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
