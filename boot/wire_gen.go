// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package boot

import (
	"github.com/Akkadius/spire/console/cmd"
	"github.com/Akkadius/spire/database"
	"github.com/Akkadius/spire/http/controllers"
	"github.com/Akkadius/spire/http/crudcontrollers"
	"github.com/Akkadius/spire/http/middleware"
	"github.com/Akkadius/spire/internal/connection"
	"github.com/Akkadius/spire/internal/desktop"
	"github.com/Akkadius/spire/internal/encryption"
	"github.com/Akkadius/spire/internal/github"
	"github.com/Akkadius/spire/internal/influx"
	"github.com/Akkadius/spire/questapi"
)

import (
	_ "github.com/go-sql-driver/mysql"
)

// Injectors from wire.go:

func InitializeApplication() (App, error) {
	db, err := provideEQEmuLocalDatabase()
	if err != nil {
		return App{}, err
	}
	logger, err := provideLogger()
	if err != nil {
		return App{}, err
	}
	cache := provideCache()
	helloWorldCommand := cmd.NewHelloWorldCommand(db, logger)
	generateModelsCommand := cmd.NewGenerateModelsCommand(db, logger)
	generateControllersCommand := cmd.NewGenerateControllersCommand(db, logger)
	generateVueFormsCommand := cmd.NewGenerateVueFormsCommand(db, logger)
	helloWorldController := controllers.NewHelloWorldController(db, logger)
	connections := provideAppDbConnections()
	encrypter := encryption.NewEncrypter()
	databaseResolver := database.NewDatabaseResolver(connections, logger, encrypter, cache)
	authController := controllers.NewAuthController(databaseResolver, logger)
	meController := controllers.NewMeController()
	client := influx.NewClient()
	analyticsController := controllers.NewAnalyticsController(logger, client, databaseResolver)
	dbConnectionCreateService := connection.NewDbConnectionCreateService(databaseResolver, logger, encrypter)
	dbConnectionCheckService := connection.NewDbConnectionCheckService(databaseResolver, logger, encrypter)
	connectionsController := controllers.NewConnectionsController(databaseResolver, logger, cache, dbConnectionCreateService, dbConnectionCheckService)
	docsController := controllers.NewDocsController(databaseResolver, logger)
	githubSourceDownloader := github.NewGithubSourceDownloader(logger)
	parseService := questapi.NewParseService(logger, cache, githubSourceDownloader)
	questExamplesGithubSourcer := questapi.NewQuestExamplesGithubSourcer(logger, cache, githubSourceDownloader)
	questApiController := controllers.NewQuestApiController(logger, parseService, questExamplesGithubSourcer)
	bootAppControllerGroups := provideControllers(helloWorldController, authController, meController, analyticsController, connectionsController, docsController, questApiController)
	aaAbilityController := crudcontrollers.NewAaAbilityController(databaseResolver, logger)
	aaRankController := crudcontrollers.NewAaRankController(databaseResolver, logger)
	accountController := crudcontrollers.NewAccountController(databaseResolver, logger)
	adventureDetailController := crudcontrollers.NewAdventureDetailController(databaseResolver, logger)
	adventureTemplateController := crudcontrollers.NewAdventureTemplateController(databaseResolver, logger)
	adventureTemplateEntryController := crudcontrollers.NewAdventureTemplateEntryController(databaseResolver, logger)
	adventureTemplateEntryFlavorController := crudcontrollers.NewAdventureTemplateEntryFlavorController(databaseResolver, logger)
	alternateCurrencyController := crudcontrollers.NewAlternateCurrencyController(databaseResolver, logger)
	blockedSpellController := crudcontrollers.NewBlockedSpellController(databaseResolver, logger)
	bugController := crudcontrollers.NewBugController(databaseResolver, logger)
	bugReportController := crudcontrollers.NewBugReportController(databaseResolver, logger)
	charCreatePointAllocationController := crudcontrollers.NewCharCreatePointAllocationController(databaseResolver, logger)
	characterAlternateAbilityController := crudcontrollers.NewCharacterAlternateAbilityController(databaseResolver, logger)
	characterAuraController := crudcontrollers.NewCharacterAuraController(databaseResolver, logger)
	characterBandolierController := crudcontrollers.NewCharacterBandolierController(databaseResolver, logger)
	characterBindController := crudcontrollers.NewCharacterBindController(databaseResolver, logger)
	characterCorpseController := crudcontrollers.NewCharacterCorpseController(databaseResolver, logger)
	characterCurrencyController := crudcontrollers.NewCharacterCurrencyController(databaseResolver, logger)
	characterDatumController := crudcontrollers.NewCharacterDatumController(databaseResolver, logger)
	characterDisciplineController := crudcontrollers.NewCharacterDisciplineController(databaseResolver, logger)
	characterInspectMessageController := crudcontrollers.NewCharacterInspectMessageController(databaseResolver, logger)
	characterItemRecastController := crudcontrollers.NewCharacterItemRecastController(databaseResolver, logger)
	characterLanguageController := crudcontrollers.NewCharacterLanguageController(databaseResolver, logger)
	characterLeadershipAbilityController := crudcontrollers.NewCharacterLeadershipAbilityController(databaseResolver, logger)
	characterMaterialController := crudcontrollers.NewCharacterMaterialController(databaseResolver, logger)
	characterMemmedSpellController := crudcontrollers.NewCharacterMemmedSpellController(databaseResolver, logger)
	characterPotionbeltController := crudcontrollers.NewCharacterPotionbeltController(databaseResolver, logger)
	characterSkillController := crudcontrollers.NewCharacterSkillController(databaseResolver, logger)
	characterSpellController := crudcontrollers.NewCharacterSpellController(databaseResolver, logger)
	contentFlagController := crudcontrollers.NewContentFlagController(databaseResolver, logger)
	dataBucketController := crudcontrollers.NewDataBucketController(databaseResolver, logger)
	dbStrController := crudcontrollers.NewDbStrController(databaseResolver, logger)
	doorController := crudcontrollers.NewDoorController(databaseResolver, logger)
	eventlogController := crudcontrollers.NewEventlogController(databaseResolver, logger)
	factionListController := crudcontrollers.NewFactionListController(databaseResolver, logger)
	factionListModController := crudcontrollers.NewFactionListModController(databaseResolver, logger)
	fishingController := crudcontrollers.NewFishingController(databaseResolver, logger)
	forageController := crudcontrollers.NewForageController(databaseResolver, logger)
	globalLootController := crudcontrollers.NewGlobalLootController(databaseResolver, logger)
	graveyardController := crudcontrollers.NewGraveyardController(databaseResolver, logger)
	gridController := crudcontrollers.NewGridController(databaseResolver, logger)
	groundSpawnController := crudcontrollers.NewGroundSpawnController(databaseResolver, logger)
	guildController := crudcontrollers.NewGuildController(databaseResolver, logger)
	hackerController := crudcontrollers.NewHackerController(databaseResolver, logger)
	instanceListController := crudcontrollers.NewInstanceListController(databaseResolver, logger)
	instanceListPlayerController := crudcontrollers.NewInstanceListPlayerController(databaseResolver, logger)
	itemController := crudcontrollers.NewItemController(databaseResolver, logger)
	ldonTrapEntryController := crudcontrollers.NewLdonTrapEntryController(databaseResolver, logger)
	ldonTrapTemplateController := crudcontrollers.NewLdonTrapTemplateController(databaseResolver, logger)
	loginAccountController := crudcontrollers.NewLoginAccountController(databaseResolver, logger)
	loginApiTokenController := crudcontrollers.NewLoginApiTokenController(databaseResolver, logger)
	loginServerAdminController := crudcontrollers.NewLoginServerAdminController(databaseResolver, logger)
	loginServerListTypeController := crudcontrollers.NewLoginServerListTypeController(databaseResolver, logger)
	loginWorldServerController := crudcontrollers.NewLoginWorldServerController(databaseResolver, logger)
	lootdropController := crudcontrollers.NewLootdropController(databaseResolver, logger)
	loottableController := crudcontrollers.NewLoottableController(databaseResolver, logger)
	nameFilterController := crudcontrollers.NewNameFilterController(databaseResolver, logger)
	npcEmoteController := crudcontrollers.NewNpcEmoteController(databaseResolver, logger)
	npcFactionController := crudcontrollers.NewNpcFactionController(databaseResolver, logger)
	npcSpellController := crudcontrollers.NewNpcSpellController(databaseResolver, logger)
	npcSpellsEffectController := crudcontrollers.NewNpcSpellsEffectController(databaseResolver, logger)
	npcSpellsEffectsEntryController := crudcontrollers.NewNpcSpellsEffectsEntryController(databaseResolver, logger)
	npcSpellsEntryController := crudcontrollers.NewNpcSpellsEntryController(databaseResolver, logger)
	npcTypeController := crudcontrollers.NewNpcTypeController(databaseResolver, logger)
	npcTypesTintController := crudcontrollers.NewNpcTypesTintController(databaseResolver, logger)
	objectController := crudcontrollers.NewObjectController(databaseResolver, logger)
	playerTitlesetController := crudcontrollers.NewPlayerTitlesetController(databaseResolver, logger)
	reportController := crudcontrollers.NewReportController(databaseResolver, logger)
	respawnTimeController := crudcontrollers.NewRespawnTimeController(databaseResolver, logger)
	saylinkController := crudcontrollers.NewSaylinkController(databaseResolver, logger)
	spawn2Controller := crudcontrollers.NewSpawn2Controller(databaseResolver, logger)
	spawnConditionController := crudcontrollers.NewSpawnConditionController(databaseResolver, logger)
	spawnConditionValueController := crudcontrollers.NewSpawnConditionValueController(databaseResolver, logger)
	spawnEventController := crudcontrollers.NewSpawnEventController(databaseResolver, logger)
	spawngroupController := crudcontrollers.NewSpawngroupController(databaseResolver, logger)
	spellsNewController := crudcontrollers.NewSpellsNewController(databaseResolver, logger)
	startingItemController := crudcontrollers.NewStartingItemController(databaseResolver, logger)
	taskController := crudcontrollers.NewTaskController(databaseResolver, logger)
	tasksetController := crudcontrollers.NewTasksetController(databaseResolver, logger)
	titleController := crudcontrollers.NewTitleController(databaseResolver, logger)
	tradeskillRecipeController := crudcontrollers.NewTradeskillRecipeController(databaseResolver, logger)
	tradeskillRecipeEntryController := crudcontrollers.NewTradeskillRecipeEntryController(databaseResolver, logger)
	trapController := crudcontrollers.NewTrapController(databaseResolver, logger)
	tributeController := crudcontrollers.NewTributeController(databaseResolver, logger)
	zoneController := crudcontrollers.NewZoneController(databaseResolver, logger)
	zonePointController := crudcontrollers.NewZonePointController(databaseResolver, logger)
	bootCrudControllers := provideCrudControllers(aaAbilityController, aaRankController, accountController, adventureDetailController, adventureTemplateController, adventureTemplateEntryController, adventureTemplateEntryFlavorController, alternateCurrencyController, blockedSpellController, bugController, bugReportController, charCreatePointAllocationController, characterAlternateAbilityController, characterAuraController, characterBandolierController, characterBindController, characterCorpseController, characterCurrencyController, characterDatumController, characterDisciplineController, characterInspectMessageController, characterItemRecastController, characterLanguageController, characterLeadershipAbilityController, characterMaterialController, characterMemmedSpellController, characterPotionbeltController, characterSkillController, characterSpellController, contentFlagController, dataBucketController, dbStrController, doorController, eventlogController, factionListController, factionListModController, fishingController, forageController, globalLootController, graveyardController, gridController, groundSpawnController, guildController, hackerController, instanceListController, instanceListPlayerController, itemController, ldonTrapEntryController, ldonTrapTemplateController, loginAccountController, loginApiTokenController, loginServerAdminController, loginServerListTypeController, loginWorldServerController, lootdropController, loottableController, nameFilterController, npcEmoteController, npcFactionController, npcSpellController, npcSpellsEffectController, npcSpellsEffectsEntryController, npcSpellsEntryController, npcTypeController, npcTypesTintController, objectController, playerTitlesetController, reportController, respawnTimeController, saylinkController, spawn2Controller, spawnConditionController, spawnConditionValueController, spawnEventController, spawngroupController, spellsNewController, startingItemController, taskController, tasksetController, titleController, tradeskillRecipeController, tradeskillRecipeEntryController, trapController, tributeController, zoneController, zonePointController)
	userContextMiddleware := middleware.NewUserContextMiddleware(databaseResolver, cache, logger)
	requestLogMiddleware := middleware.NewRequestLogMiddleware(client)
	router := NewRouter(bootAppControllerGroups, bootCrudControllers, userContextMiddleware, requestLogMiddleware)
	httpServeCommand := cmd.NewHttpServeCommand(logger, router)
	routesListCommand := cmd.NewRoutesListCommand(router, logger)
	generateConfigurationCommand := cmd.NewGenerateConfigurationCommand(databaseResolver, logger)
	spireMigrateCommand := cmd.NewSpireMigrateCommand(connections, logger)
	questApiParseCommand := cmd.NewQuestApiParseCommand(logger, parseService)
	questExampleTestCommand := cmd.NewQuestExampleTestCommand(logger, questExamplesGithubSourcer)
	v := ProvideCommands(helloWorldCommand, generateModelsCommand, generateControllersCommand, generateVueFormsCommand, httpServeCommand, routesListCommand, generateConfigurationCommand, spireMigrateCommand, questApiParseCommand, questExampleTestCommand)
	webBoot := desktop.NewWebBoot(logger, router)
	app := NewApplication(db, logger, cache, v, databaseResolver, connections, router, webBoot)
	return app, nil
}
