package user

import (
	"fmt"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/encryption"
	"github.com/spf13/cobra"
)

type ChangePasswordCommand struct {
	db      *database.Resolver
	command *cobra.Command
	crypt   *encryption.Encrypter
	user    *User
}

func (c *ChangePasswordCommand) Command() *cobra.Command {
	return c.command
}

func NewChangePasswordCommand(
	db *database.Resolver,
	crypt *encryption.Encrypter,
	user *User,
) *ChangePasswordCommand {
	i := &ChangePasswordCommand{
		db:    db,
		crypt: crypt,
		user:  user,
		command: &cobra.Command{
			Use:   "user:change-password [username] [new-password]",
			Short: "Changes a local database user's password",
			Args:  cobra.MinimumNArgs(2),
		},
	}

	i.command.Run = i.Handle

	return i
}

func (c *ChangePasswordCommand) Handle(_ *cobra.Command, args []string) {
	// args
	username := args[0]
	password := args[1]

	// change password
	err := c.user.ChangeLocalUserPassword(username, password)
	if err != nil {
		fmt.Printf("Error changing password for user %s: %v\n", username, err)
		return
	}

	fmt.Printf("Password changed for user %s\n", username)
}
