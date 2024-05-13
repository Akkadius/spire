package spire

import (
	"github.com/Akkadius/spire/internal/database"
	"github.com/spf13/cobra"
)

type MigrateCommand struct {
	connections *database.Connections
	command     *cobra.Command
}

func (c *MigrateCommand) Command() *cobra.Command {
	return c.command
}

var dropTables bool

func NewMigrateCommand(
	connections *database.Connections,
) *MigrateCommand {
	i := &MigrateCommand{
		connections: connections,
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
func (c *MigrateCommand) Handle(_ *cobra.Command, args []string) {
	c.connections.SpireMigrate(dropTables)
}

func (c *MigrateCommand) Validate(_ *cobra.Command, args []string) error {
	// Validate

	return nil
}
