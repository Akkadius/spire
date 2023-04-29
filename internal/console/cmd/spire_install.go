package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type SpireInstallCommand struct {
	logger  *logrus.Logger
	command *cobra.Command
}

func (c *SpireInstallCommand) Command() *cobra.Command {
	return c.command
}

func NewSpireInstallCommand(logger *logrus.Logger) *SpireInstallCommand {
	i := &SpireInstallCommand{
		logger: logger,
		command: &cobra.Command{
			Use:   "spire:migrate",
			Short: "Bootstraps Spire from the command-line without using the web interface",
		},
	}

	//i.command.PersistentFlags().BoolVarP(&dropTables, "drop", "d", false, "Drop tables")

	i.command.Args = i.Validate
	i.command.Run = i.Handle

	return i
}

// Handle implementation of the Command interface
func (c *SpireInstallCommand) Handle(_ *cobra.Command, args []string) {
	//
}

func (c *SpireInstallCommand) Validate(_ *cobra.Command, args []string) error {
	// Validate

	return nil
}
