package eqemuserver

import (
	"fmt"
	"github.com/Akkadius/spire/internal/eqemuserverconfig"
	"github.com/Akkadius/spire/internal/pathmgmt"
	"github.com/Akkadius/spire/internal/spire"
	"github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

type ProcessManager struct {
	logger       *logrus.Logger
	serverconfig *eqemuserverconfig.Config
	settings     *spire.Settings
	pathmgmt     *pathmgmt.PathManagement
}

func NewProcessManager(
	logger *logrus.Logger,
	serverconfig *eqemuserverconfig.Config,
	settings *spire.Settings,
	pathmgmt *pathmgmt.PathManagement,
) *ProcessManager {
	return &ProcessManager{
		logger:       logger,
		serverconfig: serverconfig,
		settings:     settings,
		pathmgmt:     pathmgmt,
	}
}

// getLauncherPath will find the launcher in the eqemu server bin path
func (p *ProcessManager) getLauncherPath() string {
	files, err := os.ReadDir(p.pathmgmt.GetEQEmuServerBinPath())
	if err != nil {
		p.logger.Fatal(err)
	}

	for _, file := range files {
		if strings.Contains(file.Name(), "occulus") {
			// execute the launcher
			launcherPath := filepath.Join(p.pathmgmt.GetEQEmuServerBinPath(), file.Name())
			binary, err := exec.LookPath(launcherPath)
			if err != nil {
				p.logger.Fatal(err)
			}

			return binary
		}
	}

	return ""
}

func (p *ProcessManager) Stop() {
	p.runLauncherCmd("stop-server")
}

func (p *ProcessManager) Start() {
	p.runLauncherCmd("server-launcher")
}

func (p *ProcessManager) Restart() {
	p.runLauncherCmd("restart-server")
}

func (p *ProcessManager) runLauncherCmd(s string) {
	launcher := p.getLauncherPath()
	if launcher == "" {
		p.logger.Fatal("Unable to find launcher")
	}

	cmd := exec.Command(launcher, s)
	cmd.Dir = p.pathmgmt.GetEQEmuServerBinPath()
	err := cmd.Start()
	if err != nil {
		p.logger.Fatal(err)
	}

	var outbuf, errbuf strings.Builder // or bytes.Buffer
	cmd.Stdout = &outbuf
	cmd.Stderr = &errbuf

	// sleep for 5 seconds to allow the server to start
	time.Sleep(5 * time.Second)

	if len(outbuf.String()) > 0 {
		fmt.Println(outbuf.String())
	}
	if len(errbuf.String()) > 0 {
		fmt.Println(errbuf.String())
	}

	// release the process
	err = cmd.Process.Release()
	if err != nil {
		p.logger.Fatal(err)
	}
}
