package generators

import (
	"github.com/Akkadius/spire/internal/database"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

type ConfigurationCommand struct {
	db      *database.Resolver
	logger  *logrus.Logger
	command *cobra.Command
}

func (c *ConfigurationCommand) Command() *cobra.Command {
	return c.command
}

func NewGenerateConfigurationCommand(
	db *database.Resolver,
	logger *logrus.Logger,
) *ConfigurationCommand {
	i := &ConfigurationCommand{
		db:     db,
		logger: logger,
		command: &cobra.Command{
			Use:   "generate:config",
			Short: "Generates base generator configurations",
		},
	}
	i.command.Args = i.Validate
	i.command.Run = i.Handle

	return i
}

func (c *ConfigurationCommand) Handle(_ *cobra.Command, _ []string) {
	db, err := c.db.GetEqemuDb().DB()
	if err != nil {
		c.logger.Fatal(err)
	}

	if err := NewDbSchemaConfig(
		db,
		c.logger,
	).Generate(os.Getenv("MYSQL_EQEMU_DATABASE")); err != nil {
		c.logger.Fatal(err)
	}

	if err := NewDbSchemaConfig(db, c.logger).GenerateKeys(); err != nil {
		c.logger.Fatal(err)
	}

}

func (c *ConfigurationCommand) Validate(_ *cobra.Command, _ []string) error {
	return nil
}
