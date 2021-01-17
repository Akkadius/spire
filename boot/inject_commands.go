package boot

import (
	"github.com/Akkadius/spire/console/cmd"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"github.com/spf13/cobra"
)

// command set
var commandSet = wire.NewSet(
	cmd.NewHelloWorldCommand,
	cmd.NewGenerateModelsCommand,
	cmd.NewGenerateControllersCommand,
	cmd.NewGenerateVueFormsCommand,
	cmd.NewHttpServeCommand,
	cmd.NewQuestApiParseCommand,
	cmd.NewRoutesListCommand,
	cmd.NewGenerateConfigurationCommand,
	cmd.NewSpireMigrateCommand,
	ProvideCommands,
)

// commands provider
func ProvideCommands(
	helloWorldCommand *cmd.HelloWorldCommand,
	generateModelsCommand *cmd.GenerateModelsCommand,
	generateControllersCommand *cmd.GenerateControllersCommand,
	generateVueFormsCommand *cmd.GenerateVueFormsCommand,
	httpServeCommand *cmd.HttpServeCommand,
	routesListCommand *cmd.RoutesListCommand,
	generateConfigurationCommand *cmd.GenerateConfigurationCommand,
	spireMigrateCommand *cmd.SpireMigrateCommand,
	questApiParseCommand *cmd.QuestApiParseCommand,
) []*cobra.Command {
	return []*cobra.Command{
		helloWorldCommand.Command(),
		generateModelsCommand.Command(),
		generateControllersCommand.Command(),
		generateVueFormsCommand.Command(),
		httpServeCommand.Command(),
		routesListCommand.Command(),
		generateConfigurationCommand.Command(),
		spireMigrateCommand.Command(),
		questApiParseCommand.Command(),
	}
}
