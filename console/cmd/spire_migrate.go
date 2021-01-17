package cmd

import (
	"eoc/database"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type SpireMigrateCommand struct {
	connections *database.Connections
	logger      *logrus.Logger
	command     *cobra.Command
}

func (g *SpireMigrateCommand) Command() *cobra.Command {
	return g.command
}

var dropTables bool

// new instance of command
func NewSpireMigrateCommand(
	connections *database.Connections,
	logger *logrus.Logger,
) *SpireMigrateCommand {
	i := &SpireMigrateCommand{
		connections: connections,
		logger:      logger,
		command: &cobra.Command{
			Use:   "spire:migrate",
			Short: "Runs database migrations for spire local database",
		},
	}

	i.command.PersistentFlags().BoolVarP(&dropTables, "drop", "d", false, "Drop tables")

	i.command.Args = i.Validate
	i.command.Run = i.Handle

	return i
}

// Handle implementation of the Command interface
func (c *SpireMigrateCommand) Handle(_ *cobra.Command, args []string) {
	c.connections.SpireMigrate(dropTables)
}

// Validate
func (g *SpireMigrateCommand) Validate(_ *cobra.Command, args []string) error {
	// Validate

	return nil
}
