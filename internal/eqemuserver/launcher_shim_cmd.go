package eqemuserver

import (
	"github.com/spf13/cobra"
	"log"
)

type LauncherShimCmd struct {
	command  *cobra.Command
	launcher *Launcher
}

func (c *LauncherShimCmd) Command() *cobra.Command {
	return c.command
}

func NewLauncherShimCmd(
	launcher *Launcher,
) *LauncherShimCmd {
	i := &LauncherShimCmd{
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
		log.Fatalf("invalid argument count\n")
	}

	switch args[0] {
	case "start":
		err := c.launcher.StartLauncherProcess()
		if err != nil {
			log.Fatal(err)
		}
	case "stop":
		err := c.launcher.Stop()
		if err != nil {
			log.Fatal(err)
		}
	case "restart":
		err := c.launcher.Restart()
		if err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatal("invalid argument")
	}
}
