package boot

import (
	"github.com/Akkadius/spire/internal/console/cmd"
	"github.com/Akkadius/spire/internal/eqemuchangelog"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"github.com/spf13/cobra"
)

// commandSet is a Wire provider set that returns a slice of commands
var commandSet = wire.NewSet(
	cmd.NewHelloWorldCommand,
	cmd.NewAdminPingOcculus,
	cmd.NewUserCreateCommand,
	cmd.NewGenerateModelsCommand,
	cmd.NewGenerateControllersCommand,
	cmd.NewHttpServeCommand,
	cmd.NewQuestApiParseCommand,
	cmd.NewRoutesListCommand,
	cmd.NewGenerateConfigurationCommand,
	cmd.NewSpireMigrateCommand,
	cmd.NewQuestExampleTestCommand,
	cmd.NewGenerateRaceModelMapsCommand,
	cmd.NewTestFilesystemCommand,
	cmd.NewSpireInitCommand,
	cmd.NewUserChangePasswordCommand,
	cmd.NewSpireOcculusUpdateCommand,
	cmd.NewSpireServerLauncherCommand,
	cmd.NewSpireCrashAnalyticsCommand,
	cmd.NewEQEmuServerUpdateCommand,
	eqemuchangelog.NewChangelogCommand,
	ProvideCommands,
)

// ProvideCommands is a Wire provider function that returns a slice of commands
func ProvideCommands(
	helloWorldCommand *cmd.HelloWorldCommand,
	adminPingOcculus *cmd.AdminPingOcculus,
	userCreateCommand *cmd.UserCreateCommand,
	generateModelsCommand *cmd.GenerateModelsCommand,
	generateControllersCommand *cmd.GenerateControllersCommand,
	httpServeCommand *cmd.HttpServeCommand,
	routesListCommand *cmd.RoutesListCommand,
	generateConfigurationCommand *cmd.GenerateConfigurationCommand,
	spireMigrateCommand *cmd.SpireMigrateCommand,
	questApiParseCommand *cmd.QuestApiParseCommand,
	questExampleTestCommand *cmd.QuestExampleTestCommand,
	generateRaceModelMapsCommand *cmd.GenerateRaceModelMapsCommand,
	changelogCmd *eqemuchangelog.ChangelogCommand,
	testFilesystemCmd *cmd.TestFilesystemCommand,
	spireInstallCmd *cmd.SpireInitCommand,
	userChangePasswordCmd *cmd.UserChangePasswordCommand,
	spireOcculusUpdateCmd *cmd.SpireOcculusUpdateCommand,
	spireServerLauncherCmd *cmd.SpireServerLauncherCommand,
	spireCrashAnalyticsCommand *cmd.SpireCrashAnalyticsFingerprintBackfillCommand,
	eQEmuServerUpdateCommand *cmd.EQEmuServerUpdateCommand,
) []*cobra.Command {
	return []*cobra.Command{
		adminPingOcculus.Command(),
		helloWorldCommand.Command(),
		userCreateCommand.Command(),
		generateModelsCommand.Command(),
		generateControllersCommand.Command(),
		httpServeCommand.Command(),
		routesListCommand.Command(),
		generateConfigurationCommand.Command(),
		spireMigrateCommand.Command(),
		questApiParseCommand.Command(),
		questExampleTestCommand.Command(),
		generateRaceModelMapsCommand.Command(),
		changelogCmd.Command(),
		testFilesystemCmd.Command(),
		spireInstallCmd.Command(),
		userChangePasswordCmd.Command(),
		spireOcculusUpdateCmd.Command(),
		spireServerLauncherCmd.Command(),
		spireCrashAnalyticsCommand.Command(),
		eQEmuServerUpdateCommand.Command(),
	}
}
