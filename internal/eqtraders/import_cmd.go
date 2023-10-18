package eqtraders

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
	"os"
)

type ImportCommand struct {
	db      *gorm.DB
	logger  *logrus.Logger
	command *cobra.Command
}

func (c *ImportCommand) Command() *cobra.Command {
	return c.command
}

func NewImportCommand(
	db *gorm.DB,
	logger *logrus.Logger,
) *ImportCommand {
	i := &ImportCommand{
		db:     db,
		logger: logger,
		command: &cobra.Command{
			Use:   "eq-traders:import [expansion_number]",
			Short: "Imports data from eqtraders.com using data scraped via eq-traders:scrape",
		},
	}

	i.command.Args = cobra.MinimumNArgs(1)
	i.command.Run = i.Handle

	return i
}

func (c *ImportCommand) Handle(cmd *cobra.Command, args []string) {
	expansion := os.Args[2]
	if expansion == "all" {

	}

	fmt.Println("Importing expansion: " + expansion)
}
