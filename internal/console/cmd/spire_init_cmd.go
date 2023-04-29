package cmd

import (
	"github.com/Akkadius/spire/internal/spire"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type SpireInitCommand struct {
	logger    *logrus.Logger
	command   *cobra.Command
	spireinit *spire.Init
}

func (c *SpireInitCommand) Command() *cobra.Command {
	return c.command
}

// flagAuthEnabled is a flag that can be passed to the spire:init command
var flagAuthEnabled bool

// NewSpireInitCommand creates a new spire:init command
func NewSpireInitCommand(
	logger *logrus.Logger,
	spireinit *spire.Init,
) *SpireInitCommand {
	i := &SpireInitCommand{
		logger:    logger,
		spireinit: spireinit,
		command: &cobra.Command{
			Use:   "spire:init [username] [password]",
			Short: "Bootstraps Spire from the command-line without using the web interface",
			Args:  cobra.MinimumNArgs(2),
		},
	}

	i.command.PersistentFlags().BoolVarP(&flagAuthEnabled, "auth-enabled", "a", true, "Whether or not Spire has authentication / authorization enabled")
	i.command.Run = i.Handle

	return i
}

// Handle implementation of the Command interface
func (c *SpireInitCommand) Handle(_ *cobra.Command, args []string) {
	username := args[0]
	password := args[1]

	// Convert bool to int
	var authEnabled int
	if flagAuthEnabled {
		authEnabled = 1
	}

	// Run the spire:init command
	err := c.spireinit.InitApp(
		&spire.InitAppRequest{
			AuthEnabled: authEnabled,
			Username:    username,
			Password:    password,
		},
	)
	if err != nil {
		c.logger.Fatal(err)
	}
}
