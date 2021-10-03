package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/Akkadius/spire/internal/models"
	"gorm.io/gorm"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
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
func (c *HelloWorldCommand) Handle(cmd *cobra.Command, _ []string) {
	var npcTypes []models.NpcType
	query := c.db.Model(&models.NpcType{})
	query = query.Limit(10)
	query = query.Find(&npcTypes)

	for _, npcType := range npcTypes {
		n, _ := json.Marshal(npcType)
		fmt.Println(string(n))
	}
}

// Validate implementation of the Command interface
func (c *HelloWorldCommand) Validate(_ *cobra.Command, _ []string) error {
	return nil
}
