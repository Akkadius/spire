package cmd

import (
	"github.com/Akkadius/spire/internal/backup"
	"github.com/k0kubun/pp/v3"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

type HelloWorldCommand struct {
	db      *gorm.DB
	logger  *logrus.Logger
	command *cobra.Command
	backup  *backup.Mysql
}

func (c *HelloWorldCommand) Command() *cobra.Command {
	return c.command
}

func NewHelloWorldCommand(
	db *gorm.DB,
	logger *logrus.Logger,
	backup *backup.Mysql,
) *HelloWorldCommand {
	i := &HelloWorldCommand{
		db:     db,
		logger: logger,
		command: &cobra.Command{
			Use:   "hello:hello-world",
			Short: "Says hello world",
		},
		backup: backup,
	}

	i.command.Args = i.Validate
	i.command.Run = i.Handle

	return i
}

// Handle implementation of the Command interface
func (c *HelloWorldCommand) Handle(cmd *cobra.Command, args []string) {
	pp.Println(
		c.backup.Backup(backup.BackupRequest{
			DumpAllTables:   true,
			ContentTables:   false,
			PlayerTables:    false,
			BotTables:       false,
			StateTables:     false,
			SystemTables:    false,
			QueryServTables: false,
			LoginTables:     false,
		}),
	)
}

// Validate implementation of the Command interface
func (c *HelloWorldCommand) Validate(_ *cobra.Command, _ []string) error {
	return nil
}
