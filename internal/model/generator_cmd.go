package model

import (
	"github.com/Akkadius/spire/internal/logger"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

// GeneratorCommand is a command that generates model definitions
type GeneratorCommand struct {
	db              *gorm.DB
	logger          *logger.AppLogger
	command         *cobra.Command
	withControllers bool
}

// Command returns the cobra command
func (c *GeneratorCommand) Command() *cobra.Command {
	return c.command
}

// NewGeneratorCommand creates a new generator command
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

	i.command.Flags().BoolVarP(
		&i.withControllers,
		"with-controllers",
		"c",
		false,
		"Generate controller stubs for each model",
	)

	return i
}

// Handle implementation of the Command interface
func (c *GeneratorCommand) Handle(cmd *cobra.Command, args []string) {
	tablesToGenerate := make([]string, 0)

	// pass in table as argument
	if len(args) > 0 {
		tablesToGenerate = append(tablesToGenerate, args[0])
	}

	generator := NewGenerator(c.logger, c.db)

	// generate http controller
	if c.withControllers {
		generator.SetWithControllers(true)
	}

	generator.Generate(tablesToGenerate)

	// sync controllers to injector
	if c.withControllers {
		generator.SyncControllersToInjector()
	}
}

// Validate validates the command arguments
func (c *GeneratorCommand) Validate(cmd *cobra.Command, _ []string) error {
	//if len(args) < 1 {
	//	return errors.New("Requires [all|table_name]")
	//}

	return nil
}
