package pathmgmt

import (
	"errors"
	"github.com/Akkadius/spire/internal/env"
	"github.com/Akkadius/spire/internal/logger"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type PathManagement struct {
	logger     *logger.AppLogger
	serverPath string
	debugging  bool
}

func NewPathManagement(logger *logger.AppLogger) *PathManagement {
	m := &PathManagement{
		logger:    logger,
		debugging: env.GetInt("PATH_MGMT_DEBUG", "0") == 1,
	}

	return m
}

const eqemuConfigFileName = "eqemu_config.json"

// SetServerPath sets the server path
// this is typically set by the installer
// this is not used by the server
func (m *PathManagement) SetServerPath(serverPath string) {
	m.serverPath = serverPath
}

func (m *PathManagement) GetEQEmuServerPath() string {
	// this is typically set by the installer
	// this is not used by the server
	if m.serverPath != "" {
		return m.serverPath
	}

	// get current path
	cwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	// first detect top level config
	topLevelPath := filepath.Join(cwd, eqemuConfigFileName)
	if _, err := os.Stat(topLevelPath); err == nil {
		m.debug("top level config path found @ [%v]", topLevelPath)
		return filepath.Dir(topLevelPath)
	}

	// second detect one level up (we're in bin folder)
	parentLevelPath := filepath.Join(cwd, "..", eqemuConfigFileName)
	if _, err := os.Stat(parentLevelPath); err == nil {
		m.debug("parent level config path found @ [%v]", parentLevelPath)
		return filepath.Dir(parentLevelPath)
	}

	m.debug("server path not found...")

	return ""
}

func (m *PathManagement) GetEQEmuServerConfigFilePath() string {
	path := filepath.Join(m.GetEQEmuServerPath(), eqemuConfigFileName)
	if _, err := os.Stat(path); err == nil {
		m.debug("found @ [%v]", path)
		return path
	}

	m.debug("path not found...")

	return ""
}

const loginConfigFile = "login.json"

func (m *PathManagement) GetEqemuLoginServerConfigPath() string {
	return filepath.Join(m.GetEQEmuServerPath(), loginConfigFile)
}

func (m *PathManagement) debug(msg string, a ...interface{}) {
	if m.debugging {
		m.logger.Debug().Msgf(msg, a...)
	}
}

const (
	binaryWorld             = "world"
	binaryZone              = "zone"
	binaryChatServer        = "ucs"
	binaryLoginserver       = "loginserver"
	binaryQueryserver       = "queryserv"
	binaryExportClientFiles = "export_client_files"
	binaryImportClientFiles = "import_client_files"
	binDir                  = "bin"
	logsDir                 = "logs"
	backupsDir              = "backups"
	questsDir               = "quests"
	mapsDir                 = "maps"
	exportDir               = "export"
	importDir               = "import"
)

// GetBinary returns the platform specific binary name
// Example linux = world, windows = world.exe
func GetBinary(s string) string {
	bin := s
	if runtime.GOOS == "windows" {
		bin += ".exe"
	}

	return bin
}

func (m *PathManagement) GetEQEmuServerBinPath() string {
	return filepath.Join(m.GetEQEmuServerPath(), binDir)
}

func (m *PathManagement) GetLogsDirPath() string {
	return filepath.Join(m.GetEQEmuServerPath(), logsDir)
}

func (m *PathManagement) GetWorldBinPath() string {
	return filepath.Join(m.GetEQEmuServerBinPath(), GetBinary(binaryWorld))
}

func (m *PathManagement) GetZoneBinPath() string {
	return filepath.Join(m.GetEQEmuServerBinPath(), GetBinary(binaryZone))
}

func (m *PathManagement) GetUCSBinPath() string {
	return filepath.Join(m.GetEQEmuServerBinPath(), GetBinary(binaryChatServer))
}

func (m *PathManagement) GetLoginserverBinPath() string {
	return filepath.Join(m.GetEQEmuServerBinPath(), GetBinary(binaryLoginserver))
}

func (m *PathManagement) GetQSBinPath() string {
	return filepath.Join(m.GetEQEmuServerBinPath(), GetBinary(binaryQueryserver))
}

func (m *PathManagement) GetExportClientFilesBinPath() string {
	return filepath.Join(m.GetEQEmuServerBinPath(), GetBinary(binaryExportClientFiles))
}

func (m *PathManagement) GetImportClientFilesBinPath() string {
	return filepath.Join(m.GetEQEmuServerBinPath(), GetBinary(binaryImportClientFiles))
}

func (m *PathManagement) GetExportDir() string {
	return filepath.Join(m.GetEQEmuServerPath(), exportDir)
}

func (m *PathManagement) GetBackupsDir() string {
	return filepath.Join(m.GetEQEmuServerPath(), backupsDir)
}

func (m *PathManagement) GetQuestsDir() string {
	return filepath.Join(m.GetEQEmuServerPath(), questsDir)
}

func (m *PathManagement) GetMapsDir() string {
	return filepath.Join(m.GetEQEmuServerPath(), mapsDir)
}

func (m *PathManagement) MakeExportDirIfNotExists() error {
	path := m.GetExportDir()
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}
