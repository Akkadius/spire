package eqemuchangelog

import (
	"fmt"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
	"strconv"
)

type ChangelogCommand struct {
	db        *gorm.DB
	command   *cobra.Command
	changelog *Changelog
}

func (c *ChangelogCommand) Command() *cobra.Command {
	return c.command
}

func NewChangelogCommand(db *gorm.DB, changelog *Changelog) *ChangelogCommand {
	i := &ChangelogCommand{
		db: db,
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
			fmt.Printf("Invalid argument [%v] must be an integer\n", args[0])
		}
		days = intVar
	}

	fmt.Printf("------------------------------------\n")
	fmt.Printf("     Listing back [%v] day(s)\n", days)
	fmt.Printf("      (Copy below the line)\n")
	fmt.Printf("------------------------------------\n\n")

	fmt.Println(c.changelog.BuildChangelog(c.changelog.getCommitsDaysBack()))

}

// Validate implementation of the Command interface
func (c *ChangelogCommand) Validate(_ *cobra.Command, _ []string) error {
	return nil
}
