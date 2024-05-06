package spire

import (
	"github.com/Akkadius/spire/internal/pathmgmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
)

type ServerLauncherCommand struct {
	logger   *logrus.Logger
	command  *cobra.Command
	pathmgmt *pathmgmt.PathManagement
}

func (c *ServerLauncherCommand) Command() *cobra.Command {
	return c.command
}

func NewServerLauncherCommand(
	logger *logrus.Logger,
	pathmgmt *pathmgmt.PathManagement,
) *ServerLauncherCommand {
	i := &ServerLauncherCommand{
		logger:   logger,
		pathmgmt: pathmgmt,
		command: &cobra.Command{
			Use:   "spire:launcher [start|stop|restart]",
			Short: "Manages running the eqemu server launcher",
		},
	}

	i.command.Run = i.Handle

	return i
}

// Handle implementation of the Command interface
func (c *ServerLauncherCommand) Handle(_ *cobra.Command, args []string) {
	if len(args) == 0 {
		_ = c.command.Help()
		return
	}

	// Get the path to the currently running executable
	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}

	// Resolve any symbolic links to get the actual path
	exPath, err := filepath.EvalSymlinks(ex)
	if err != nil {
		log.Fatal(err)
	}

	bin, err := exec.LookPath(exPath)
	if err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command(bin, "eqemu-server:launcher", args[0])
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = c.pathmgmt.GetEQEmuServerPath()
	cmd.SysProcAttr = &syscall.SysProcAttr{Setsid: true}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	if args[0] == "start" {
		// we do this otherwise we get a zombie process
		go func() {
			_ = cmd.Wait()
		}()
		return
	}
	_ = cmd.Wait()
}
