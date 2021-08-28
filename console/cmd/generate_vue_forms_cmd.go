package cmd

import (
	"errors"
	"github.com/Akkadius/spire/generators"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type GenerateVueFormsCommand struct {
	db      *gorm.DB
	logger  *logrus.Logger
	command *cobra.Command
}

func (g *GenerateVueFormsCommand) Command() *cobra.Command {
	return g.command
}

// new instance of command
func NewGenerateVueFormsCommand(
	db *gorm.DB,
	logger *logrus.Logger,
) *GenerateVueFormsCommand {

	i := &GenerateVueFormsCommand{
		db:     db,
		logger: logger,
		command: &cobra.Command{
			Use:   "generate:vue-forms [all|table_name]",
			Short: "Generates Vue forms from backend models",
		},
	}

	i.command.Args = i.Validate
	i.command.Run = i.Handle

	return i
}

// Handle implementation of the Command interface
func (g *GenerateVueFormsCommand) Handle(_ *cobra.Command, args []string) {
	tablesToGenerate := make([]string, 0)

	// pass in table as argument
	if len(args) > 0 && args[0] != "all" {
		tablesToGenerate = append(tablesToGenerate, args[0])
	}

	err := generators.NewGenerateVueForm(
		generators.GenerateVueFormContext{
			TablesToGenerate: tablesToGenerate,
		},
		g.logger,
	).Generate()

	if err != nil {
		g.logger.Fatal(err)
	}
}

// Validate
func (g *GenerateVueFormsCommand) Validate(_ *cobra.Command, args []string) error {
	// Validate

	if len(args) == 0 {
		return errors.New("Invalid argument")
	}

	return nil
}
