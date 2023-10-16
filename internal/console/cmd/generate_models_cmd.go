package cmd

import (
	"github.com/Akkadius/spire/internal/generators"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

type GenerateModelsCommand struct {
	db      *gorm.DB
	logger  *logrus.Logger
	command *cobra.Command
}

func (c *GenerateModelsCommand) Command() *cobra.Command {
	return c.command
}

func NewGenerateModelsCommand(db *gorm.DB, logger *logrus.Logger) *GenerateModelsCommand {
	i := &GenerateModelsCommand{
		db:     db,
		logger: logger,
		command: &cobra.Command{
			Use:   "generate:models [all|table_name]",
			Short: "Generates model definitions",
		},
	}
	i.command.Args = i.Validate
	i.command.Run = i.Handle

	return i
}

func (c *GenerateModelsCommand) Handle(cmd *cobra.Command, args []string) {
	tablesToGenerate := make([]string, 0)

	// pass in table as argument
	if len(args) > 0 {
		tablesToGenerate = append(tablesToGenerate, args[0])
	}

	generators.NewGenerateModel(
		generators.GenerateModelContext{
			TablesToGenerate: tablesToGenerate,
		},
		c.logger,
		c.db,
	).Generate()

}

func (c *GenerateModelsCommand) Validate(cmd *cobra.Command, _ []string) error {
	//if len(args) < 1 {
	//	return errors.New("Requires [all|table_name]")
	//}

	return nil
}
