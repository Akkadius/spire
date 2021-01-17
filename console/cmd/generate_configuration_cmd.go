package cmd

import (
	"eoc/database"
	"eoc/generators"
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

// new instance of command
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

// Handle
func (c *GenerateConfigurationCommand) Handle(_ *cobra.Command, _ []string) {
	if err := generators.NewGenerateDbSchemaConfig(
		c.db.GetEqemuDb().DB(),
		c.logger,
	).Generate(os.Getenv("MYSQL_EQEMU_DATABASE")); err != nil {
		c.logger.Fatal(err)
	}

	if err := generators.NewGenerateDbSchemaConfig(c.db.GetEqemuDb().DB(), c.logger).GenerateKeys(); err != nil {
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

// Validate
func (c *GenerateConfigurationCommand) Validate(_ *cobra.Command, _ []string) error {
	return nil
}
