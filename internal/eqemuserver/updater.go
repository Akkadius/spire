package eqemuserver

import (
	"bufio"
	"encoding/json"
	"errors"
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
	return u.settings.GetSetting(updateTypeSetting).Value
}

type ServerVersionInfo struct {
	BotsDatabaseVersion int    `json:"bots_database_version"`
	CompileDate         string `json:"compile_date"`
	CompileTime         string `json:"compile_time"`
	DatabaseVersion     int    `json:"database_version"`
	ServerVersion       string `json:"server_version"`
}

func (u *Updater) GetVersionInfo() (ServerVersionInfo, error) {
	worldBin := u.pathmgmt.GetWorldBinPath()
	if _, err := os.Stat(worldBin); errors.Is(err, os.ErrNotExist) {
		return ServerVersionInfo{}, errors.New("Failed to find World binary to fetch version")
	}

	cmd := exec.Command(worldBin, "world:version")
	cmd.Dir = u.pathmgmt.GetEQEmuServerPath()
	output, err := cmd.Output()
	if err != nil {
		return ServerVersionInfo{}, err
	}

	// not all binaries simply output json alone
	// there was an output bug
	o := string(output)
	var n string
	startWatch := false
	for _, s := range strings.Split(o, "\n") {
		if strings.Contains(s, "{") {
			startWatch = true
		}
		if startWatch {
			n += s
		}
	}

	var v ServerVersionInfo
	err = json.Unmarshal([]byte(n), &v)
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
	var dirname string
	var err error
	s := u.settings.GetSetting("BUILD_LOCATION")
	if s.ID != 0 {
		dirname = s.Value
	}

	if len(dirname) == 0 {
		dirname, err = os.UserHomeDir()
		if err != nil {
			return BuildInfo{}, err
		}
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
