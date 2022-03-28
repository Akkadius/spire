package boot

import (
	"github.com/Akkadius/spire/internal/console/cmd"
	"github.com/Akkadius/spire/internal/permissions"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"github.com/spf13/cobra"
)

// command set
var commandSet = wire.NewSet(
	cmd.NewHelloWorldCommand,
	cmd.NewGenerateModelsCommand,
	cmd.NewGenerateControllersCommand,
	cmd.NewHttpServeCommand,
	cmd.NewQuestApiParseCommand,
	cmd.NewRoutesListCommand,
	cmd.NewGenerateConfigurationCommand,
	cmd.NewSpireMigrateCommand,
	cmd.NewQuestExampleTestCommand,
	cmd.NewGenerateRaceModelMapsCommand,
	permissions.NewResourceListCommand,
	ProvideCommands,
)

// commands provider
func ProvideCommands(
	helloWorldCommand *cmd.HelloWorldCommand,
	generateModelsCommand *cmd.GenerateModelsCommand,
	generateControllersCommand *cmd.GenerateControllersCommand,
	httpServeCommand *cmd.HttpServeCommand,
	routesListCommand *cmd.RoutesListCommand,
	generateConfigurationCommand *cmd.GenerateConfigurationCommand,
	spireMigrateCommand *cmd.SpireMigrateCommand,
	questApiParseCommand *cmd.QuestApiParseCommand,
	questExampleTestCommand *cmd.QuestExampleTestCommand,
	generateRaceModelMapsCommand *cmd.GenerateRaceModelMapsCommand,
	resourceListCmd *permissions.ResourceListCommand,
) []*cobra.Command {
	return []*cobra.Command{
		helloWorldCommand.Command(),
		generateModelsCommand.Command(),
		generateControllersCommand.Command(),
		httpServeCommand.Command(),
		routesListCommand.Command(),
		generateConfigurationCommand.Command(),
		spireMigrateCommand.Command(),
		questApiParseCommand.Command(),
		questExampleTestCommand.Command(),
		generateRaceModelMapsCommand.Command(),
		resourceListCmd.Command(),
	}
}
