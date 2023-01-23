package eqemuchangelog

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

type ChangelogCommand struct {
	db        *gorm.DB
	logger    *logrus.Logger
	command   *cobra.Command
	changelog *Changelog
}

func (c *ChangelogCommand) Command() *cobra.Command {
	return c.command
}

func NewChangelogCommand(db *gorm.DB, logger *logrus.Logger, changelog *Changelog) *ChangelogCommand {
	i := &ChangelogCommand{
		db:     db,
		logger: logger,
		command: &cobra.Command{
			Use: "eqemu-server:changelog",
		},
		changelog: changelog,
	}

	i.command.Args = i.Validate
	i.command.Run = i.Handle

	return i
}

// Handle implementation of the Command interface
func (c *ChangelogCommand) Handle(cmd *cobra.Command, args []string) {
	fmt.Println(c.changelog.BuildChangelog())
}

// Validate implementation of the Command interface
func (c *ChangelogCommand) Validate(_ *cobra.Command, _ []string) error {
	return nil
}
