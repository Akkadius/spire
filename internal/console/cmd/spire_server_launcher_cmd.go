package cmd

import (
	"github.com/Akkadius/spire/internal/pathmgmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type SpireServerLauncherCommand struct {
	logger   *logrus.Logger
	command  *cobra.Command
	pathmgmt *pathmgmt.PathManagement
}

func (c *SpireServerLauncherCommand) Command() *cobra.Command {
	return c.command
}

func NewSpireServerLauncherCommand(
	logger *logrus.Logger,
	pathmgmt *pathmgmt.PathManagement,
) *SpireServerLauncherCommand {
	i := &SpireServerLauncherCommand{
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
func (c *SpireServerLauncherCommand) Handle(_ *cobra.Command, args []string) {
	if len(args) == 0 {
		_ = c.command.Help()
		return
	}

	// loop through files in c.pathmgmt.GetEQEmuServerBinPath() and find the launcher
	// if it exists, run it with the args
	files, err := os.ReadDir(c.pathmgmt.GetEQEmuServerBinPath())
	if err != nil {
		c.logger.Fatal(err)
	}

	// this mapping is for the occulus launcher
	// until the occulus launcher is entirely replaced
	argMappings := map[string]string{
		"start":   "server-launcher",
		"stop":    "stop-server",
		"restart": "restart-server",
	}

	var arg string
	for s := range argMappings {
		if s == args[0] {
			arg = argMappings[s]
		}
	}

	// if arg is empty, we didn't find a mapping
	if arg == "" {
		_ = c.command.Help()
		return
	}

	for _, file := range files {
		if strings.Contains(file.Name(), "occulus") {
			// execute the launcher
			launcherPath := filepath.Join(c.pathmgmt.GetEQEmuServerBinPath(), file.Name())
			binary, err := exec.LookPath(launcherPath)
			if err != nil {
				c.logger.Fatal(err)
			}

			c.logger.Infof("Running %s %s", binary, arg)

			// run the launcher
			cmd := exec.Command(binary, arg)
			cmd.Dir = c.pathmgmt.GetEQEmuServerBinPath()
			//cmd.Stdout = os.Stdout
			//cmd.Stderr = os.Stderr
			err = cmd.Start()
			if err != nil {
				c.logger.Fatal(err)
			}

			// release the process
			err = cmd.Process.Release()
			if err != nil {
				c.logger.Fatal(err)
			}
		}
	}
}
