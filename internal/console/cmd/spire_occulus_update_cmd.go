package cmd

import (
	"github.com/Akkadius/spire/internal/occulus"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// SpireOcculusUpdateCommand is the command to update Occulus
type SpireOcculusUpdateCommand struct {
	logger  *logrus.Logger
	command *cobra.Command
	occulus *occulus.ProcessManagement
}

// Command implementation of the Command interface
func (c *SpireOcculusUpdateCommand) Command() *cobra.Command {
	return c.command
}

// NewSpireOcculusUpdateCommand creates a new spire:init command
func NewSpireOcculusUpdateCommand(
	logger *logrus.Logger,
	occulus *occulus.ProcessManagement,
) *SpireOcculusUpdateCommand {
	i := &SpireOcculusUpdateCommand{
		logger:  logger,
		occulus: occulus,
		command: &cobra.Command{
			Use:   "spire:occulus-update",
			Short: "Updates or fetches Occulus",
		},
	}

	i.command.Run = i.Handle

	return i
}

// Handle implementation of the Command interface
func (c *SpireOcculusUpdateCommand) Handle(_ *cobra.Command, args []string) {
	path, err := c.occulus.FetchOcculusAndGetBinaryPath()
	if err != nil {
		c.logger.Error(err)
		return
	}

	c.logger.Infof("Occulus updated to [%v]", path)
}
