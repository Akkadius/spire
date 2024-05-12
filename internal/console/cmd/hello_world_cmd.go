package cmd

import (
	"github.com/Akkadius/spire/internal/backup"
	"github.com/Akkadius/spire/internal/pathmgmt"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

type HelloWorldCommand struct {
	db       *gorm.DB
	command  *cobra.Command
	backup   *backup.Mysql
	pathmgmt *pathmgmt.PathManagement
}

func (c *HelloWorldCommand) Command() *cobra.Command {
	return c.command
}

func NewHelloWorldCommand(
	db *gorm.DB,
	backup *backup.Mysql,
	pathmgmt *pathmgmt.PathManagement,
) *HelloWorldCommand {
	i := &HelloWorldCommand{
		db: db,
		command: &cobra.Command{
			Use:   "hello:hello-world",
			Short: "Says hello world",
		},
		backup:   backup,
		pathmgmt: pathmgmt,
	}

	i.command.Args = i.Validate
	i.command.Run = i.Handle

	return i
}

// Handle implementation of the Command interface
func (c *HelloWorldCommand) Handle(cmd *cobra.Command, args []string) {
}

// Validate implementation of the Command interface
func (c *HelloWorldCommand) Validate(_ *cobra.Command, _ []string) error {
	return nil
}
