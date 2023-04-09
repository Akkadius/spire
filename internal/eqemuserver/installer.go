package eqemuserver

import (
	"bufio"
	"fmt"
	"github.com/Akkadius/spire/internal/pathmgmt"
	"github.com/go-git/go-git/v5"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

type Installer struct {
	pathmanager *pathmgmt.PathManagement
	logger      *logrus.Logger
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
	logger := getLogger()
	i := &Installer{
		logger:      logger,
		pathmanager: pathmgmt.NewPathManagement(logger),
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
	// install perl packages
	// install perl modules
	// install mysql server
	// initialize mysql database
	// seed mysql database
	// initialize config files
	// install opcodes
	// install binaries

	//a.installOsPackages()
	a.initMySQL()
	//a.initializeDirectories()
	//a.cloneEQEmuMaps()
	//a.clonePeqQuests()
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
}

func (a *Installer) DoneBanner(s string) {
	a.logger.Println("----------------------------------------")
	a.logger.Printf("| ✅ | %v\n", s)
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

	a.logger.Infof("Maps and navmesh files are large, please be patient if output is not updating\n")
	a.logger.Infof("Cloning EQEmuMaps from github.com/Akkadius/EQEmuMaps.git\n")

	// clone the repository
	path := filepath.Join(a.pathmanager.GetEQEmuServerPath(), "maps")
	_, err := git.PlainClone(path, false, &git.CloneOptions{
		URL:      "https://github.com/Akkadius/EQEmuMaps.git",
		Progress: a.logger.Writer(),
	})

	if err != nil && err != git.ErrRepositoryAlreadyExists {
		a.logger.Errorf("Could not clone EQEmuMaps [%v]\n", err)
	}

	// if the repository already exists, update it instead
	if err == git.ErrRepositoryAlreadyExists {
		a.logger.Infof("EQEmuMaps already exists, skipping clone and updating instead\n")

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

		a.logger.Infof("EQEmuMaps updated successfully!\n")
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
		RootPassword:     "root",
	}

	// update root password
	a.logger.Infof("Updating root user password\n")
	a.DbExecSafe(fmt.Sprintf("FLUSH PRIVILEGES; ALTER USER 'root'@'localhost' IDENTIFIED BY '%v'", c.RootPassword))

	// create a new database
	a.logger.Infof("Creating database [%v]\n", c.DatabaseName)
	a.DbExec(c, fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %v", c.DatabaseName))

	// create a new user
	a.logger.Infof("Creating user [%v]\n", c.DatabaseUser)
	a.DbExec(c, fmt.Sprintf("CREATE USER IF NOT EXISTS '%v'@'localhost' IDENTIFIED BY '%v'", c.DatabaseUser, c.DatabasePassword))

	// grant privileges to the new user
	a.logger.Infof("Granting privileges to user [%v]\n", c.DatabaseUser)
	a.DbExec(c, fmt.Sprintf("GRANT ALL PRIVILEGES ON %v.* TO '%v'@'localhost'", c.DatabaseName, c.DatabaseUser))

	// flush privileges
	a.logger.Infoln("Flushing privileges")
	a.DbExec(c, "FLUSH PRIVILEGES")

	a.Exec("sudo", []string{"pkill", "-f", "-9", "mysql"})
	a.Exec("sudo", []string{"service"})

	a.DoneBanner("Initializing MySQL")
}

func (a *Installer) DbExecSafe(statement string) {
	a.Exec("mysql", []string{"-uroot", "-e", fmt.Sprintf("%v", statement)})
}

func (a *Installer) DbExec(c MysqlConfig, statement string) {
	a.Exec(
		"mysql",
		[]string{
			"-uroot",
			fmt.Sprintf("-p%v", c.RootPassword),
			"-e",
			fmt.Sprintf("%v", statement),
		},
	)
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
