package user

import (
	"fmt"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/encryption"
	"github.com/Akkadius/spire/internal/logger"
	"github.com/Akkadius/spire/internal/models"
	"github.com/spf13/cobra"
)

type CreateCommand struct {
	db      *database.Resolver
	logger  *logger.AppLogger
	command *cobra.Command
	crypt   *encryption.Encrypter
	user    *User
}

func (c *CreateCommand) Command() *cobra.Command {
	return c.command
}

func NewCreateCommand(db *database.Resolver, logger *logger.AppLogger, crypt *encryption.Encrypter, user *User) *CreateCommand {
	i := &CreateCommand{
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

func (c *CreateCommand) Handle(_ *cobra.Command, args []string) {
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
		Provider: LoginProviderLocal,
	}

	// new user
	newUser, err := c.user.CreateUser(user)
	if err != nil {
		c.logger.Error().Err(err).Msg("Failed to create user")
	}

	if newUser.ID > 0 {
		c.logger.Info().Msgf("[user] Created user ID [%v]", newUser.ID)
	}
}

// Validate implementation of the Command interface
func (c *CreateCommand) Validate(_ *cobra.Command, _ []string) error {
	return nil
}
