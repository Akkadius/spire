package eqemuserver

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type LauncherShimCmd struct {
	logger   *logrus.Logger
	command  *cobra.Command
	launcher *Launcher
}

func (c *LauncherShimCmd) Command() *cobra.Command {
	return c.command
}

func NewLauncherShimCmd(
	logger *logrus.Logger,
	launcher *Launcher,
) *LauncherShimCmd {
	i := &LauncherShimCmd{
		logger:   logger,
		launcher: launcher,
		command: &cobra.Command{
			Use:   "spire:launcher [start|stop|restart]",
			Short: "Manages running the eqemu server launcher",
		},
	}

	i.command.Run = i.Handle

	return i
}

// Handle implementation of the Command interface
func (c *LauncherShimCmd) Handle(_ *cobra.Command, args []string) {
	if len(args) == 0 {
		_ = c.command.Help()
		return
	}

	if len(args) > 1 {
		c.logger.Fatal("too many arguments")
	}

	switch args[0] {
	case "start":
		err := c.launcher.StartLauncherProcess()
		if err != nil {
			c.logger.Fatal(err)
		}
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
