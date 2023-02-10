package eqemuserver

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/download"
	"github.com/Akkadius/spire/internal/pathmgmt"
	"github.com/Akkadius/spire/internal/serverconfig"
	"github.com/Akkadius/spire/internal/spire"
	"github.com/Akkadius/spire/internal/unzip"
	"github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

type Updater struct {
	db           *database.DatabaseResolver
	logger       *logrus.Logger
	serverconfig *serverconfig.EQEmuServerConfig
	settings     *spire.Settings
	pathmgmt     *pathmgmt.PathManagement
}

func NewUpdater(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
	serverconfig *serverconfig.EQEmuServerConfig,
	settings *spire.Settings,
	pathmgmt *pathmgmt.PathManagement,

) *Updater {
	return &Updater{
		db:           db,
		logger:       logger,
		serverconfig: serverconfig,
		settings:     settings,
		pathmgmt:     pathmgmt,
	}
}

const (
	updateTypeSetting      = "SERVER_UPDATE_TYPE"
	updateTypeRelease      = "release"
	updateTypeSelfCompiled = "self-compiled"
)

func (u *Updater) SetUpdateType(updateType string) {
	u.settings.SetSetting(updateTypeSetting, updateType)
}

func (u *Updater) GetUpdateType() string {
	return u.settings.GetSetting(updateTypeSetting)
}

type ServerVersionInfo struct {
	BotsDatabaseVersion int    `json:"bots_database_version"`
	CompileDate         string `json:"compile_date"`
	CompileTime         string `json:"compile_time"`
	DatabaseVersion     int    `json:"database_version"`
	ServerVersion       string `json:"server_version"`
}

func (u *Updater) GetVersionInfo() (ServerVersionInfo, error) {
	binPath := filepath.Join(u.pathmgmt.GetEQEmuServerPath(), "bin")
	bin := "world"
	startCmd := ""
	if _, err := os.Stat(filepath.Join(binPath, bin)); err == nil {
		startCmd = filepath.Join(binPath, bin)
	} else if _, err := os.Stat(filepath.Join(startCmd, fmt.Sprintf("%v.exe", bin))); err == nil {
		startCmd = filepath.Join(binPath, fmt.Sprintf("%v.exe", bin))
	}

	cmd := exec.Command(startCmd, "world:version")
	cmd.Dir = u.pathmgmt.GetEQEmuServerPath()
	output, err := cmd.Output()
	if err != nil {
		return ServerVersionInfo{}, err
	}

	var v ServerVersionInfo
	err = json.Unmarshal(output, &v)
	if err != nil {
		return ServerVersionInfo{}, err
	}

	return v, err
}

func (u *Updater) InstallRelease(release string) error {
	file := path.Base(release)
	downloadPath := filepath.Join(os.TempDir(), file)
	err := download.WithProgress(downloadPath, release)
	if err != nil {
		return err
	}

	uz := unzip.New(downloadPath, u.pathmgmt.GetEQEmuServerBinPath())
	err = uz.Extract()
	if err != nil {
		return err
	}

	return nil
}

type BuildInfo struct {
	SourceDirectory string `json:"source_directory"`
	BuildTool       string `json:"build_tool"`
}

// GetBuildInfo tries to auto discovery source directory and returns build tool
func (u *Updater) GetBuildInfo() (BuildInfo, error) {
	dirname, err := os.UserHomeDir()
	if err != nil {
		return BuildInfo{}, err
	}

	foundPath := ""
	err = filepath.Walk(
		dirname,
		func(path string, info os.FileInfo, err error) error {
			if strings.Contains(path, "build/CMakeCache.txt") {
				foundPath = path
			}

			return nil
		},
	)
	if err != nil {
		return BuildInfo{}, err
	}

	buildTool := ""
	if len(foundPath) > 0 {
		file, err := os.Open(foundPath)
		if err != nil {
			return BuildInfo{}, err
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		// optionally, resize scanner's capacity for lines over 64K, see next example
		for scanner.Scan() {
			if strings.Contains(scanner.Text(), "CMAKE_MAKE_PROGRAM:FILEPATH=") {
				buildTool = strings.ReplaceAll(scanner.Text(), "CMAKE_MAKE_PROGRAM:FILEPATH=", "")
			}
		}

		if err := scanner.Err(); err != nil {
			return BuildInfo{}, err
		}
	}

	return BuildInfo{
		SourceDirectory: filepath.Dir(foundPath),
		BuildTool:       buildTool,
	}, nil
}
