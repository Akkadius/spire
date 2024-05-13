package generators

import (
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/logger"
	"github.com/spf13/cobra"
	"os"
)

type ConfigurationCommand struct {
	db      *database.Resolver
	logger  *logger.AppLogger
	command *cobra.Command
}

func (c *ConfigurationCommand) Command() *cobra.Command {
	return c.command
}

func NewGenerateConfigurationCommand(
	db *database.Resolver,
	logger *logger.AppLogger,
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
		c.logger.Fatal().Err(err).Msg("failed to get eqemu db")
	}

	if err := NewDbSchemaConfig(
		db,
		c.logger,
	).Generate(os.Getenv("MYSQL_EQEMU_DATABASE")); err != nil {
		c.logger.Fatal().Err(err).Msg("failed to generate db schema config")
	}

	if err := NewDbSchemaConfig(db, c.logger).GenerateKeys(); err != nil {
		c.logger.Fatal().Err(err).Msg("failed to generate db schema keys config")
	}

}

func (c *ConfigurationCommand) Validate(_ *cobra.Command, _ []string) error {
	return nil
}
