package eqemuserver

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type LauncherCmd struct {
	logger   *logrus.Logger
	command  *cobra.Command
	launcher *Launcher
}

func (c *LauncherCmd) Command() *cobra.Command {
	return c.command
}

func NewLauncherCmd(
	logger *logrus.Logger,
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

	if len(args) > 1 {
		c.logger.Fatal("too many arguments")
	}

	switch args[0] {
	case "start":
		err := c.launcher.Start()
		if err != nil {
			c.logger.Fatal(err)
		}
		c.launcher.Process()
	case "stop":
		err := c.launcher.Stop()
		if err != nil {
			c.logger.Fatal(err)
		}
	case "restart":
		err := c.launcher.Restart()
		if err != nil {
			c.logger.Fatal(err)
		}
	default:
		c.logger.Fatal("invalid argument")
	}
}
