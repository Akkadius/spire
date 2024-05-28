package eqemuserver

import (
	"github.com/Akkadius/spire/internal/env"
	"github.com/Akkadius/spire/internal/logger"
	"github.com/spf13/cobra"
)

type LauncherCmd struct {
	logger   *logger.AppLogger `json:"logger,omitempty"`
	command  *cobra.Command    `json:"command,omitempty"`
	launcher *Launcher         `json:"launcher,omitempty"`
}

func (c *LauncherCmd) Command() *cobra.Command {
	return c.command
}

func NewLauncherCmd(
	logger *logger.AppLogger,
	launcher *Launcher,
) *LauncherCmd {
	i := &LauncherCmd{
		logger:   logger,
		launcher: launcher,
		command: &cobra.Command{
			Use:   "eqemu-server:launcher [start|stop|restart]",
			Short: "Manages running the eqemu server launcher",
		},
	}

	i.command.Run = i.Handle

	return i
}

// Handle implementation of the Command interface
func (c *LauncherCmd) Handle(_ *cobra.Command, args []string) {
	if len(args) == 0 {
		_ = c.command.Help()
		return
	}

	env.SetAppModeCli()

	if len(args) > 1 {
		c.logger.Fatal().Msg("too many arguments")
	}

	switch args[0] {
	case "start":
		err := c.launcher.Start()
		if err != nil {
			c.logger.Fatal().Err(err).Msg("failed to start launcher")
		}
		c.launcher.Process()
	case "stop":
		err := c.launcher.Stop()
		if err != nil {
			c.logger.Fatal().Err(err).Msg("failed to stop launcher")
		}
	case "restart":
		err := c.launcher.Restart()
		if err != nil {
			c.logger.Fatal().Err(err).Msg("failed to restart launcher")
		}
	default:
		c.logger.Fatal().Msg("invalid command")
	}
}
