package cmd

import (
	"fmt"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/encryption"
	"github.com/Akkadius/spire/internal/models"
	"github.com/Akkadius/spire/internal/serverconfig"
	"github.com/k0kubun/pp/v3"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"time"
)

type UserCreateCommand struct {
	db           *database.DatabaseResolver
	logger       *logrus.Logger
	command      *cobra.Command
	serverconfig *serverconfig.EQEmuServerConfig
	crypt        *encryption.Encrypter
}

func (c *UserCreateCommand) Command() *cobra.Command {
	return c.command
}

func NewUserCreateCommand(db *database.DatabaseResolver, logger *logrus.Logger, crypt *encryption.Encrypter) *UserCreateCommand {
	i := &UserCreateCommand{
		db:     db,
		logger: logger,
		crypt:  crypt,
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

	username := args[0]
	password := args[1]

	encrypted := c.crypt.Encrypt(password, c.crypt.GetEncryptionKey())
	decrypted := c.crypt.Decrypt(encrypted, c.crypt.GetEncryptionKey())

	pp.Printf("Decrypted password is [%v]", decrypted)

	var users []models.User
	c.db.GetSpireDb().Where("user_name = ?", username).Find(&users)
	if len(users) > 0 {
		c.logger.Error("[user] User already exists")
		return
	}

	var newUser models.User
	c.db.GetSpireDb().FirstOrCreate(
		&newUser, models.User{
			UserName:  username,
			FullName:  username,
			Password:  c.crypt.Encrypt(password, c.crypt.GetEncryptionKey()),
			Provider:  "local",
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		},
	)

	if newUser.ID > 0 {
		c.logger.Infof("[user] Created user ID [%v]", newUser.ID)
	}
}

// Validate implementation of the Command interface
func (c *UserCreateCommand) Validate(_ *cobra.Command, _ []string) error {
	return nil
}
