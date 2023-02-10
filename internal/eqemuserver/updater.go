package eqemuserver

import (
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
