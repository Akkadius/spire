package cmd

import (
	"github.com/Akkadius/spire/internal/console"
	"github.com/Akkadius/spire/internal/http"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/sirupsen/logrus"

	"errors"

	"github.com/spf13/cobra"
)

type HttpServeCommand struct {
	port    uint
	logger  *logrus.Logger
	command *cobra.Command
	router  *routes.Router
}

func (c *HttpServeCommand) Command() *cobra.Command {
	return c.command
}

func NewHttpServeCommand(logger *logrus.Logger, router *routes.Router) *HttpServeCommand {
	i := &HttpServeCommand{
		logger: logger,
		command: &cobra.Command{
			Use:   "http:serve",
			Short: "Starts the API listener",
		},
		router: router,
	}

	i.command.Flags().Uint("port", 3000, "Port that the HTTP server listens on")
	i.command.Args = i.Validate
	i.command.Run = i.Handle

	return i
}

// Handle implementation of the Command interface
func (c *HttpServeCommand) Handle(_ *cobra.Command, _ []string) {
	if err := http.Serve(c.port, c.logger, c.router); err != nil {
		c.logger.WithError(err).Fatal(err.Error())
	}
}

// Validate implementation of the Command interface
func (c *HttpServeCommand) Validate(cmd *cobra.Command, _ []string) error {
	port := console.UintFromFlags(cmd.Flags(), "port")
	if port < 0 || port > 99999 {
		return errors.New("port is out of range")
	}

	c.port = port

	return nil
}
