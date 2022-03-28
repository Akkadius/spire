package permissions

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type ResourceListCommand struct {
	resourceList *ResourceList
	command      *cobra.Command
	logger       *logrus.Logger
}

func (c *ResourceListCommand) Command() *cobra.Command {
	return c.command
}

func NewResourceListCommand(resourceList *ResourceList, logger *logrus.Logger) *ResourceListCommand {
	i := &ResourceListCommand{
		logger: logger,
		command: &cobra.Command{
			Use:   "permissions:resource-list",
			Short: "Lists resources",
		},
		resourceList: resourceList,
	}

	i.command.Args = i.Validate
	i.command.Run = i.Handle

	return i
}

// Handle implementation of the Command interface
func (c *ResourceListCommand) Handle(_ *cobra.Command, _ []string) {
	c.resourceList.Get()
}

// Validate implementation of the Command interface
func (c *ResourceListCommand) Validate(_ *cobra.Command, _ []string) error {
	return nil
}
