package user

import (
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/encryption"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type ChangePasswordCommand struct {
	db      *database.Resolver
	logger  *logrus.Logger
	command *cobra.Command
	crypt   *encryption.Encrypter
	user    *User
}

func (c *ChangePasswordCommand) Command() *cobra.Command {
	return c.command
}

func NewChangePasswordCommand(
	db *database.Resolver,
	logger *logrus.Logger,
	crypt *encryption.Encrypter,
	user *User,
) *ChangePasswordCommand {
	i := &ChangePasswordCommand{
		db:     db,
		logger: logger,
		crypt:  crypt,
		user:   user,
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
		c.logger.Errorf("Error changing password for user %s: %v", username, err)
		return
	}

	c.logger.Infof("Password changed for user %s", username)
}
