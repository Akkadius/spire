package cmd

import (
	"fmt"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/encryption"
	"github.com/Akkadius/spire/internal/models"
	"github.com/Akkadius/spire/internal/spire"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type UserCreateCommand struct {
	db      *database.DatabaseResolver
	logger  *logrus.Logger
	command *cobra.Command
	crypt   *encryption.Encrypter
	user    *spire.UserService
}

func (c *UserCreateCommand) Command() *cobra.Command {
	return c.command
}

func NewUserCreateCommand(db *database.DatabaseResolver, logger *logrus.Logger, crypt *encryption.Encrypter, user *spire.UserService) *UserCreateCommand {
	i := &UserCreateCommand{
		db:     db,
		logger: logger,
		crypt:  crypt,
		user:   user,
		command: &cobra.Command{
			Use:   "user:create [username] [password]",
			Short: "Creates a local database user",
		},
	}

	i.command.Args = i.Validate
	i.command.Run = i.Handle

	return i
}

func (c *UserCreateCommand) Handle(_ *cobra.Command, args []string) {
	if len(args) != 2 {
		fmt.Println("Usage: [username] [password]")
		return
	}

	// args
	username := args[0]
	password := args[1]

	user := models.User{
		UserName: username,
		FullName: username,
		Password: password,
		Provider: spire.LoginProviderLocal,
	}

	// new user
	newUser, err := c.user.CreateUser(user)
	if err != nil {
		c.logger.Error(err)
	}

	if newUser.ID > 0 {
		c.logger.Infof("[user] Created user ID [%v]", newUser.ID)
	}
}

// Validate implementation of the Command interface
func (c *UserCreateCommand) Validate(_ *cobra.Command, _ []string) error {
	return nil
}
