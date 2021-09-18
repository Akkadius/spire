package cmd

import (
	"github.com/Akkadius/spire/database"
	"github.com/Akkadius/spire/generators"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

type GenerateConfigurationCommand struct {
	db      *database.DatabaseResolver
	logger  *logrus.Logger
	command *cobra.Command
}

func (c *GenerateConfigurationCommand) Command() *cobra.Command {
	return c.command
}

func NewGenerateConfigurationCommand(
	db *database.DatabaseResolver,
	logger *logrus.Logger,
) *GenerateConfigurationCommand {
	i := &GenerateConfigurationCommand{
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

func (c *GenerateConfigurationCommand) Handle(_ *cobra.Command, _ []string) {
	db, err := c.db.GetEqemuDb().DB()
	if err != nil {
		c.logger.Fatal(err)
	}

	if err := generators.NewGenerateDbSchemaConfig(
		db,
		c.logger,
	).Generate(os.Getenv("MYSQL_EQEMU_DATABASE")); err != nil {
		c.logger.Fatal(err)
	}

	if err := generators.NewGenerateDbSchemaConfig(db, c.logger).GenerateKeys(); err != nil {
		c.logger.Fatal(err)
	}

	if err := generators.NewGenerateVueForm(
		generators.GenerateVueFormContext{
			TablesToGenerate:      []string{},
			UseDatabaseSchemaDocs: true,
		},
		c.logger,
	).GenerateConfig(); err != nil {
		c.logger.Fatal(err)
	}
}

func (c *GenerateConfigurationCommand) Validate(_ *cobra.Command, _ []string) error {
	return nil
}
