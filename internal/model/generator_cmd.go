package model

import (
	"github.com/Akkadius/spire/internal/logger"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

type GeneratorCommand struct {
	db      *gorm.DB
	logger  *logger.AppLogger
	command *cobra.Command
}

func (c *GeneratorCommand) Command() *cobra.Command {
	return c.command
}

func NewGeneratorCommand(db *gorm.DB, logger *logger.AppLogger) *GeneratorCommand {
	i := &GeneratorCommand{
		db:     db,
		logger: logger,
		command: &cobra.Command{
			Use:   "make:models [all|table_name]",
			Short: "Generates model definitions",
		},
	}
	i.command.Args = i.Validate
	i.command.Run = i.Handle

	return i
}

func (c *GeneratorCommand) Handle(cmd *cobra.Command, args []string) {
	tablesToGenerate := make([]string, 0)

	// pass in table as argument
	if len(args) > 0 {
		tablesToGenerate = append(tablesToGenerate, args[0])
	}

	NewGenerator(c.logger, c.db).Generate(tablesToGenerate)
}

func (c *GeneratorCommand) Validate(cmd *cobra.Command, _ []string) error {
	//if len(args) < 1 {
	//	return errors.New("Requires [all|table_name]")
	//}

	return nil
}
