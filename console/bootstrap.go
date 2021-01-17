package console

import (
	"github.com/spf13/cobra"
)

// Run boots & executes the application
func Run(commands []*cobra.Command) error {
	cmd := &cobra.Command{Use: "spire"}
	cmd.SetHelpFunc(helpFunc)
	cmd.AddCommand(commands...)

	return cmd.Execute()
}
