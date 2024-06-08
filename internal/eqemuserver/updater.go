package eqemuserver

import (
	"bufio"
	"encoding/json"
	"errors"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/download"
	"github.com/Akkadius/spire/internal/eqemuserverconfig"
	"github.com/Akkadius/spire/internal/github"
	"github.com/Akkadius/spire/internal/logger"
	"github.com/Akkadius/spire/internal/pathmgmt"
	"github.com/Akkadius/spire/internal/spire"
	"github.com/Akkadius/spire/internal/unzip"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

type Updater struct {
	db           *database.Resolver
	logger       *logger.AppLogger
	serverconfig *eqemuserverconfig.Config
	settings     *spire.Settings
	pathmgmt     *pathmgmt.PathManagement
	unzipper     *unzip.Unzipper
}

func NewUpdater(
	db *database.Resolver,
	logger *logger.AppLogger,
	serverconfig *eqemuserverconfig.Config,
	settings *spire.Settings,
	pathmgmt *pathmgmt.PathManagement,
	unzipper *unzip.Unzipper,
) *Updater {
	return &Updater{
		db:           db,
		logger:       logger,
		serverconfig: serverconfig,
		settings:     settings,
		pathmgmt:     pathmgmt,
		unzipper:     unzipper,
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

func (u *Updater) GetBuildLocation() string {
	return u.settings.GetSetting(spire.SettingBuildLocation).Value
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

	err = u.unzipper.Extract(downloadPath, u.pathmgmt.GetEQEmuServerBinPath())
	if err != nil {
		return err
	}

	return nil
}

type BuildInfo struct {
	SourceDirectory string `json:"source_directory"`
	BuildTool       string `json:"build_tool"`
	BuildCores      string `json:"build_cores"`
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
		BuildCores:      u.settings.GetSetting(spire.SettingBuildCores).Value,
	}, nil
}

func (u *Updater) GetLatestReleaseVersion() (string, error) {
	// get latest version from https://github.com/eqemu/server
	type Release struct {
		TagName string `json:"tag_name"`
	}

	// get latest release version
	resp, err := http.Get("https://api.github.com/repos/eqemu/server/releases/latest")
	if err != nil {
		u.logger.Fatal().Err(err).Msg("could not get latest release version")
	}

	defer resp.Body.Close()

	// read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		u.logger.Fatal().Err(err).Msg("could not read response body")
	}

	err = github.HandleRateLimitBodyResponseFormatter(body, resp.Header)
	if err != nil {
		u.logger.Fatal().Err(err).Msg("could not handle rate limit")
	}

	// bind body to struct
	var release Release
	err = json.Unmarshal(body, &release)
	if err != nil {
		u.logger.Fatal().Err(err).Msg("could not unmarshal response body")
	}

	return strings.ReplaceAll(release.TagName, "v", ""), nil
}
