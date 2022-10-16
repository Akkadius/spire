package cmd

import (
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/models"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

type HelloWorldCommand struct {
	db      *gorm.DB
	logger  *logrus.Logger
	command *cobra.Command
}

func (c *HelloWorldCommand) Command() *cobra.Command {
	return c.command
}

func NewHelloWorldCommand(db *gorm.DB, logger *logrus.Logger) *HelloWorldCommand {
	i := &HelloWorldCommand{
		db:     db,
		logger: logger,
		command: &cobra.Command{
			Use:   "hello:hello-world",
			Short: "Says hello world",
		},
	}

	i.command.Args = i.Validate
	i.command.Run = i.Handle

	return i
}

// Handle implementation of the Command interface
func (c *HelloWorldCommand) Handle(cmd *cobra.Command, args []string) {
	// grab first entry
	var source models.Task
	err := c.db.Where("id", 1).First(&source).Error
	if err != nil {
		c.logger.Println(err)
	}

	var dest = source
	dest.Type = 1
	dest.ExpReward = 30002
	dest.Description = "Some new description3"

	c.db.Model(&dest).Updates(database.ResultDifference(source, dest))
}

// Validate implementation of the Command interface
func (c *HelloWorldCommand) Validate(_ *cobra.Command, _ []string) error {
	return nil
}
