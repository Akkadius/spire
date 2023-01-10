package cmd

import (
	"github.com/Akkadius/spire/internal/eqemuserverapi"
	"github.com/k0kubun/pp/v3"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

type HelloWorldCommand struct {
	db             *gorm.DB
	logger         *logrus.Logger
	command        *cobra.Command
	eqemuserverapi *eqemuserverapi.Client
}

func (c *HelloWorldCommand) Command() *cobra.Command {
	return c.command
}

func NewHelloWorldCommand(db *gorm.DB, logger *logrus.Logger, eqemuserverapi *eqemuserverapi.Client) *HelloWorldCommand {
	i := &HelloWorldCommand{
		db:     db,
		logger: logger,
		command: &cobra.Command{
			Use:   "hello:hello-world",
			Short: "Says hello world",
		},
		eqemuserverapi: eqemuserverapi,
	}

	i.command.Args = i.Validate
	i.command.Run = i.Handle

	return i
}

// Handle implementation of the Command interface
func (c *HelloWorldCommand) Handle(cmd *cobra.Command, args []string) {
	r, err := c.eqemuserverapi.GetReloadTypes()
	pp.Println(r)

	if err != nil {
		c.logger.Error(err)
	}
}

// Validate implementation of the Command interface
func (c *HelloWorldCommand) Validate(_ *cobra.Command, _ []string) error {
	return nil
}
