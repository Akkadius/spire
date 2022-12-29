package pathmgmt

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"path/filepath"
)

type PathManagement struct {
	logger *logrus.Logger
}

func NewPathManagement(logger *logrus.Logger) *PathManagement {
	m := &PathManagement{
		logger: logger,
	}

	return m
}

const eqemuConfigFileName = "eqemu_config.json"

func (m PathManagement) GetEQEmuServerPath() string {
	// get current path
	cwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	// first detect top level config
	topLevelPath := filepath.Join(cwd, eqemuConfigFileName)
	if _, err := os.Stat(topLevelPath); err == nil {
		m.debug("[GetEQEmuServerPath] top level config path found @ [%v]", topLevelPath)
		return filepath.Dir(topLevelPath)
	}

	// second detect one level up (we're in bin folder)
	parentLevelPath := filepath.Join(cwd, "..", eqemuConfigFileName)
	if _, err := os.Stat(parentLevelPath); err == nil {
		m.debug("[GetEQEmuServerPath] parent level config path found @ [%v]", parentLevelPath)
		return filepath.Dir(parentLevelPath)
	}

	m.debug("[GetEQEmuServerPath] server path not found...")

	return ""
}

func (m PathManagement) GetEQEmuServerConfigFilePath() string {
	path := filepath.Join(m.GetEQEmuServerPath(), eqemuConfigFileName)
	if _, err := os.Stat(path); err == nil {
		m.debug("[GetEQEmuServerConfigFilePath] found @ [%v]", path)
		return path
	}

	m.debug("[GetEQEmuServerConfigFilePath] path not found...")

	return ""
}

func (m *PathManagement) debug(msg string, a ...interface{}) {
	if len(os.Getenv("DEBUG")) >= 3 {
		if len(a) > 0 {
			m.logger.Debug("[pathmgmt.go] " + fmt.Sprintf(msg, a...) + "\n")
			return
		}
		m.logger.Debug("[pathmgmt.go] " + fmt.Sprintf(msg) + "\n")
	}
}
