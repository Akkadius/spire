package generators

import (
	"github.com/Akkadius/spire/internal/logger"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

type ModelGeneratorCommand struct {
	db      *gorm.DB
	logger  *logger.AppLogger
	command *cobra.Command
}

func (c *ModelGeneratorCommand) Command() *cobra.Command {
	return c.command
}

func NewModelGeneratorCommand(db *gorm.DB, logger *logger.AppLogger) *ModelGeneratorCommand {
	i := &ModelGeneratorCommand{
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

func (c *ModelGeneratorCommand) Handle(cmd *cobra.Command, args []string) {
	tablesToGenerate := make([]string, 0)

	// pass in table as argument
	if len(args) > 0 {
		tablesToGenerate = append(tablesToGenerate, args[0])
	}

	NewGenerateModel(
		GenerateModelContext{
			TablesToGenerate: tablesToGenerate,
		},
		c.logger,
		c.db,
	).Generate()

}

func (c *ModelGeneratorCommand) Validate(cmd *cobra.Command, _ []string) error {
	//if len(args) < 1 {
	//	return errors.New("Requires [all|table_name]")
	//}

	return nil
}
