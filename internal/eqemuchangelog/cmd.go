package eqemuchangelog

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
	"strconv"
	"time"
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
			Use:   "eqemu-server:changelog [days-back]",
			Short: "Generates eqemu changelog",
		},
		changelog: changelog,
	}

	i.command.Args = i.Validate
	i.command.Run = i.Handle

	return i
}

// Handle implementation of the Command interface
func (c *ChangelogCommand) Handle(cmd *cobra.Command, args []string) {
	days := 30
	if len(args) == 1 {
		intVar, err := strconv.Atoi(args[0])
		if err != nil {
			c.logger.Error(err)
		}
		days = intVar
	}

	fmt.Printf("------------------------------------\n")
	fmt.Printf("     Listing back [%v] day(s)\n", days)
	fmt.Printf("      (Copy below the line)\n")
	fmt.Printf("------------------------------------\n\n")

	fmt.Println(
		c.changelog.BuildChangelog(
			c.changelog.getCommitsDaysBack(time.Duration(days)),
		),
	)

}

// Validate implementation of the Command interface
func (c *ChangelogCommand) Validate(_ *cobra.Command, _ []string) error {
	return nil
}
