// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package boot

import (
	"github.com/Akkadius/spire/internal/assets"
	"github.com/Akkadius/spire/internal/auditlog"
	"github.com/Akkadius/spire/internal/backup"
	"github.com/Akkadius/spire/internal/clientfiles"
	"github.com/Akkadius/spire/internal/connection"
	"github.com/Akkadius/spire/internal/console/cmd"
	"github.com/Akkadius/spire/internal/database"
	"github.com/Akkadius/spire/internal/deploy"
	"github.com/Akkadius/spire/internal/desktop"
	"github.com/Akkadius/spire/internal/encryption"
	"github.com/Akkadius/spire/internal/eqemuanalytics"
	"github.com/Akkadius/spire/internal/eqemuchangelog"
	"github.com/Akkadius/spire/internal/eqemuserver"
	"github.com/Akkadius/spire/internal/github"
	"github.com/Akkadius/spire/internal/http"
	"github.com/Akkadius/spire/internal/http/controllers"
	"github.com/Akkadius/spire/internal/http/crudcontrollers"
	"github.com/Akkadius/spire/internal/http/middleware"
	"github.com/Akkadius/spire/internal/http/staticmaps"
	"github.com/Akkadius/spire/internal/influx"
	"github.com/Akkadius/spire/internal/occulus"
	"github.com/Akkadius/spire/internal/pathmgmt"
	"github.com/Akkadius/spire/internal/permissions"
	"github.com/Akkadius/spire/internal/questapi"
	"github.com/Akkadius/spire/internal/serverconfig"
	"github.com/Akkadius/spire/internal/spire"
	"github.com/Akkadius/spire/internal/telnet"
	"github.com/Akkadius/spire/internal/websocket"
	"github.com/gertd/go-pluralize"
)

import (
	_ "github.com/go-sql-driver/mysql"
)

// Injectors from wire.go:

func InitializeApplication() (App, error) {
	logger, err := provideLogger()
	if err != nil {
		return App{}, err
	}
	pathManagement := pathmgmt.NewPathManagement(logger)
	eqEmuServerConfig := serverconfig.NewEQEmuServerConfig(logger, pathManagement)
	db, err := provideEQEmuLocalDatabase(eqEmuServerConfig)
	if err != nil {
		return App{}, err
	}
	cache := provideCache()
	mysql := backup.NewMysql(logger, pathManagement)
	helloWorldCommand := cmd.NewHelloWorldCommand(db, logger, mysql)
	processManagement := occulus.NewProcessManagement(pathManagement, logger)
	proxy := occulus.NewProxy(logger, eqEmuServerConfig, processManagement)
	adminPingOcculus := cmd.NewAdminPingOcculus(db, logger, eqEmuServerConfig, proxy)
	connections := provideAppDbConnections(eqEmuServerConfig, logger)
	encrypter := encryption.NewEncrypter(logger, eqEmuServerConfig)
	databaseResolver := database.NewDatabaseResolver(connections, logger, encrypter, cache)
	userService := spire.NewUserService(databaseResolver, logger, encrypter, cache)
	userCreateCommand := cmd.NewUserCreateCommand(databaseResolver, logger, encrypter, userService)
	generateModelsCommand := cmd.NewGenerateModelsCommand(db, logger)
	generateControllersCommand := cmd.NewGenerateControllersCommand(db, logger)
	helloWorldController := controllers.NewHelloWorldController(db, logger)
	authController := controllers.NewAuthController(databaseResolver, logger, userService)
	meController := controllers.NewMeController()
	client := influx.NewClient()
	analyticsController := controllers.NewAnalyticsController(logger, client, databaseResolver)
	dbConnectionCreateService := connection.NewDbConnectionCreateService(databaseResolver, logger, encrypter)
	dbConnectionCheckService := connection.NewDbConnectionCheckService(databaseResolver, logger, encrypter)
	pluralizeClient := pluralize.NewClient()
	service := permissions.NewService(databaseResolver, cache, logger, pluralizeClient)
	settings := spire.NewSettings(connections, logger)
	init := spire.NewInit(connections, eqEmuServerConfig, logger, settings, cache, encrypter, dbConnectionCreateService, userService)
	connectionsController := controllers.NewConnectionsController(databaseResolver, logger, cache, dbConnectionCreateService, dbConnectionCheckService, service, init, userService)
	docsController := controllers.NewDocsController(databaseResolver, logger)
	githubSourceDownloader := github.NewGithubSourceDownloader(logger, cache)
	parseService := questapi.NewParseService(logger, cache, githubSourceDownloader)
	questExamplesGithubSourcer := questapi.NewQuestExamplesGithubSourcer(logger, cache, githubSourceDownloader)
	questApiController := questapi.NewQuestApiController(logger, parseService, questExamplesGithubSourcer)
	appController := controllers.NewAppController(cache, logger, init, userService, settings)
	queryController := controllers.NewQueryController(databaseResolver, logger)
	exporter := clientfiles.NewExporter(logger)
	importer := clientfiles.NewImporter(logger)
	clientFilesController := clientfiles.NewClientFilesController(logger, exporter, importer, databaseResolver)
	staticMapController := staticmaps.NewStaticMapController(databaseResolver, logger)
	releases := eqemuanalytics.NewReleases()
	eqemuanalyticsAnalyticsController := eqemuanalytics.NewAnalyticsController(logger, databaseResolver, releases)
	changelog := eqemuchangelog.NewChangelog()
	eqemuChangelogController := eqemuchangelog.NewEqemuChangelogController(logger, databaseResolver, changelog)
	deployController := deploy.NewDeployController(logger)
	assetsController := assets.NewAssetsController(logger, databaseResolver)
	permissionsController := permissions.NewPermissionsController(logger, databaseResolver, service)
	usersController := spire.NewUsersController(databaseResolver, logger, userService, encrypter)
	controller := occulus.NewController(logger, databaseResolver, proxy)
	telnetClient := telnet.NewClient(logger)
	eqemuserverClient := eqemuserver.NewClient(telnetClient)
	eqemuserverController := eqemuserver.NewController(databaseResolver, logger, eqemuserverClient, eqEmuServerConfig)
	serverconfigController := serverconfig.NewController(logger, eqEmuServerConfig)
	backupController := backup.NewController(logger, mysql, pathManagement)
	spireHandler := websocket.NewSpireHandler(logger, pathManagement)
	websocketController := websocket.NewController(logger, pathManagement, spireHandler)
	bootAppControllerGroups := provideControllers(helloWorldController, authController, meController, analyticsController, connectionsController, docsController, questApiController, appController, queryController, clientFilesController, staticMapController, eqemuanalyticsAnalyticsController, eqemuChangelogController, deployController, assetsController, permissionsController, usersController, controller, eqemuserverController, serverconfigController, backupController, websocketController)
	userEvent := auditlog.NewUserEvent(databaseResolver, logger, cache)
	aaAbilityController := crudcontrollers.NewAaAbilityController(databaseResolver, logger, userEvent)
	aaRankController := crudcontrollers.NewAaRankController(databaseResolver, logger, userEvent)
	aaRankEffectController := crudcontrollers.NewAaRankEffectController(databaseResolver, logger, userEvent)
	aaRankPrereqController := crudcontrollers.NewAaRankPrereqController(databaseResolver, logger, userEvent)
	accountController := crudcontrollers.NewAccountController(databaseResolver, logger, userEvent)
	accountFlagController := crudcontrollers.NewAccountFlagController(databaseResolver, logger, userEvent)
	accountIpController := crudcontrollers.NewAccountIpController(databaseResolver, logger, userEvent)
	accountRewardController := crudcontrollers.NewAccountRewardController(databaseResolver, logger, userEvent)
	adventureDetailController := crudcontrollers.NewAdventureDetailController(databaseResolver, logger, userEvent)
	adventureMemberController := crudcontrollers.NewAdventureMemberController(databaseResolver, logger, userEvent)
	adventureStatController := crudcontrollers.NewAdventureStatController(databaseResolver, logger, userEvent)
	adventureTemplateController := crudcontrollers.NewAdventureTemplateController(databaseResolver, logger, userEvent)
	adventureTemplateEntryController := crudcontrollers.NewAdventureTemplateEntryController(databaseResolver, logger, userEvent)
	adventureTemplateEntryFlavorController := crudcontrollers.NewAdventureTemplateEntryFlavorController(databaseResolver, logger, userEvent)
	alternateCurrencyController := crudcontrollers.NewAlternateCurrencyController(databaseResolver, logger, userEvent)
	auraController := crudcontrollers.NewAuraController(databaseResolver, logger, userEvent)
	baseDatumController := crudcontrollers.NewBaseDatumController(databaseResolver, logger, userEvent)
	blockedSpellController := crudcontrollers.NewBlockedSpellController(databaseResolver, logger, userEvent)
	bookController := crudcontrollers.NewBookController(databaseResolver, logger, userEvent)
	botBuffController := crudcontrollers.NewBotBuffController(databaseResolver, logger, userEvent)
	botCreateCombinationController := crudcontrollers.NewBotCreateCombinationController(databaseResolver, logger, userEvent)
	botDatumController := crudcontrollers.NewBotDatumController(databaseResolver, logger, userEvent)
	botGroupController := crudcontrollers.NewBotGroupController(databaseResolver, logger, userEvent)
	botGroupMemberController := crudcontrollers.NewBotGroupMemberController(databaseResolver, logger, userEvent)
	botGuildMemberController := crudcontrollers.NewBotGuildMemberController(databaseResolver, logger, userEvent)
	botHealRotationController := crudcontrollers.NewBotHealRotationController(databaseResolver, logger, userEvent)
	botHealRotationMemberController := crudcontrollers.NewBotHealRotationMemberController(databaseResolver, logger, userEvent)
	botHealRotationTargetController := crudcontrollers.NewBotHealRotationTargetController(databaseResolver, logger, userEvent)
	botInspectMessageController := crudcontrollers.NewBotInspectMessageController(databaseResolver, logger, userEvent)
	botInventoryController := crudcontrollers.NewBotInventoryController(databaseResolver, logger, userEvent)
	botOwnerOptionController := crudcontrollers.NewBotOwnerOptionController(databaseResolver, logger, userEvent)
	botPetBuffController := crudcontrollers.NewBotPetBuffController(databaseResolver, logger, userEvent)
	botPetController := crudcontrollers.NewBotPetController(databaseResolver, logger, userEvent)
	botPetInventoryController := crudcontrollers.NewBotPetInventoryController(databaseResolver, logger, userEvent)
	botSpellCastingChanceController := crudcontrollers.NewBotSpellCastingChanceController(databaseResolver, logger, userEvent)
	botSpellSettingController := crudcontrollers.NewBotSpellSettingController(databaseResolver, logger, userEvent)
	botSpellsEntryController := crudcontrollers.NewBotSpellsEntryController(databaseResolver, logger, userEvent)
	botStanceController := crudcontrollers.NewBotStanceController(databaseResolver, logger, userEvent)
	botTimerController := crudcontrollers.NewBotTimerController(databaseResolver, logger, userEvent)
	bugController := crudcontrollers.NewBugController(databaseResolver, logger, userEvent)
	bugReportController := crudcontrollers.NewBugReportController(databaseResolver, logger, userEvent)
	buyerController := crudcontrollers.NewBuyerController(databaseResolver, logger, userEvent)
	charCreateCombinationController := crudcontrollers.NewCharCreateCombinationController(databaseResolver, logger, userEvent)
	charCreatePointAllocationController := crudcontrollers.NewCharCreatePointAllocationController(databaseResolver, logger, userEvent)
	charRecipeListController := crudcontrollers.NewCharRecipeListController(databaseResolver, logger, userEvent)
	characterActivityController := crudcontrollers.NewCharacterActivityController(databaseResolver, logger, userEvent)
	characterAltCurrencyController := crudcontrollers.NewCharacterAltCurrencyController(databaseResolver, logger, userEvent)
	characterAlternateAbilityController := crudcontrollers.NewCharacterAlternateAbilityController(databaseResolver, logger, userEvent)
	characterAuraController := crudcontrollers.NewCharacterAuraController(databaseResolver, logger, userEvent)
	characterBandolierController := crudcontrollers.NewCharacterBandolierController(databaseResolver, logger, userEvent)
	characterBindController := crudcontrollers.NewCharacterBindController(databaseResolver, logger, userEvent)
	characterBuffController := crudcontrollers.NewCharacterBuffController(databaseResolver, logger, userEvent)
	characterCorpseController := crudcontrollers.NewCharacterCorpseController(databaseResolver, logger, userEvent)
	characterCorpseItemController := crudcontrollers.NewCharacterCorpseItemController(databaseResolver, logger, userEvent)
	characterCurrencyController := crudcontrollers.NewCharacterCurrencyController(databaseResolver, logger, userEvent)
	characterDatumController := crudcontrollers.NewCharacterDatumController(databaseResolver, logger, userEvent)
	characterDisciplineController := crudcontrollers.NewCharacterDisciplineController(databaseResolver, logger, userEvent)
	characterEnabledtaskController := crudcontrollers.NewCharacterEnabledtaskController(databaseResolver, logger, userEvent)
	characterExpModifierController := crudcontrollers.NewCharacterExpModifierController(databaseResolver, logger, userEvent)
	characterExpeditionLockoutController := crudcontrollers.NewCharacterExpeditionLockoutController(databaseResolver, logger, userEvent)
	characterInspectMessageController := crudcontrollers.NewCharacterInspectMessageController(databaseResolver, logger, userEvent)
	characterInstanceSafereturnController := crudcontrollers.NewCharacterInstanceSafereturnController(databaseResolver, logger, userEvent)
	characterItemRecastController := crudcontrollers.NewCharacterItemRecastController(databaseResolver, logger, userEvent)
	characterLanguageController := crudcontrollers.NewCharacterLanguageController(databaseResolver, logger, userEvent)
	characterLeadershipAbilityController := crudcontrollers.NewCharacterLeadershipAbilityController(databaseResolver, logger, userEvent)
	characterMaterialController := crudcontrollers.NewCharacterMaterialController(databaseResolver, logger, userEvent)
	characterMemmedSpellController := crudcontrollers.NewCharacterMemmedSpellController(databaseResolver, logger, userEvent)
	characterPeqzoneFlagController := crudcontrollers.NewCharacterPeqzoneFlagController(databaseResolver, logger, userEvent)
	characterPetBuffController := crudcontrollers.NewCharacterPetBuffController(databaseResolver, logger, userEvent)
	characterPetInfoController := crudcontrollers.NewCharacterPetInfoController(databaseResolver, logger, userEvent)
	characterPetInventoryController := crudcontrollers.NewCharacterPetInventoryController(databaseResolver, logger, userEvent)
	characterPotionbeltController := crudcontrollers.NewCharacterPotionbeltController(databaseResolver, logger, userEvent)
	characterSkillController := crudcontrollers.NewCharacterSkillController(databaseResolver, logger, userEvent)
	characterSpellController := crudcontrollers.NewCharacterSpellController(databaseResolver, logger, userEvent)
	characterTaskController := crudcontrollers.NewCharacterTaskController(databaseResolver, logger, userEvent)
	characterTaskTimerController := crudcontrollers.NewCharacterTaskTimerController(databaseResolver, logger, userEvent)
	chatchannelReservedNameController := crudcontrollers.NewChatchannelReservedNameController(databaseResolver, logger, userEvent)
	completedSharedTaskActivityStateController := crudcontrollers.NewCompletedSharedTaskActivityStateController(databaseResolver, logger, userEvent)
	completedSharedTaskController := crudcontrollers.NewCompletedSharedTaskController(databaseResolver, logger, userEvent)
	completedSharedTaskMemberController := crudcontrollers.NewCompletedSharedTaskMemberController(databaseResolver, logger, userEvent)
	completedTaskController := crudcontrollers.NewCompletedTaskController(databaseResolver, logger, userEvent)
	contentFlagController := crudcontrollers.NewContentFlagController(databaseResolver, logger, userEvent)
	damageshieldtypeController := crudcontrollers.NewDamageshieldtypeController(databaseResolver, logger, userEvent)
	dataBucketController := crudcontrollers.NewDataBucketController(databaseResolver, logger, userEvent)
	dbStrController := crudcontrollers.NewDbStrController(databaseResolver, logger, userEvent)
	discordWebhookController := crudcontrollers.NewDiscordWebhookController(databaseResolver, logger, userEvent)
	discoveredItemController := crudcontrollers.NewDiscoveredItemController(databaseResolver, logger, userEvent)
	doorController := crudcontrollers.NewDoorController(databaseResolver, logger, userEvent)
	dynamicZoneController := crudcontrollers.NewDynamicZoneController(databaseResolver, logger, userEvent)
	dynamicZoneMemberController := crudcontrollers.NewDynamicZoneMemberController(databaseResolver, logger, userEvent)
	dynamicZoneTemplateController := crudcontrollers.NewDynamicZoneTemplateController(databaseResolver, logger, userEvent)
	eventlogController := crudcontrollers.NewEventlogController(databaseResolver, logger, userEvent)
	expeditionController := crudcontrollers.NewExpeditionController(databaseResolver, logger, userEvent)
	expeditionLockoutController := crudcontrollers.NewExpeditionLockoutController(databaseResolver, logger, userEvent)
	expeditionMemberController := crudcontrollers.NewExpeditionMemberController(databaseResolver, logger)
	factionAssociationController := crudcontrollers.NewFactionAssociationController(databaseResolver, logger, userEvent)
	factionBaseDatumController := crudcontrollers.NewFactionBaseDatumController(databaseResolver, logger, userEvent)
	factionListController := crudcontrollers.NewFactionListController(databaseResolver, logger, userEvent)
	factionListModController := crudcontrollers.NewFactionListModController(databaseResolver, logger, userEvent)
	factionValueController := crudcontrollers.NewFactionValueController(databaseResolver, logger, userEvent)
	fishingController := crudcontrollers.NewFishingController(databaseResolver, logger, userEvent)
	forageController := crudcontrollers.NewForageController(databaseResolver, logger, userEvent)
	friendController := crudcontrollers.NewFriendController(databaseResolver, logger, userEvent)
	globalLootController := crudcontrollers.NewGlobalLootController(databaseResolver, logger, userEvent)
	gmIpController := crudcontrollers.NewGmIpController(databaseResolver, logger, userEvent)
	graveyardController := crudcontrollers.NewGraveyardController(databaseResolver, logger, userEvent)
	gridController := crudcontrollers.NewGridController(databaseResolver, logger, userEvent)
	gridEntryController := crudcontrollers.NewGridEntryController(databaseResolver, logger, userEvent)
	groundSpawnController := crudcontrollers.NewGroundSpawnController(databaseResolver, logger, userEvent)
	groupIdController := crudcontrollers.NewGroupIdController(databaseResolver, logger, userEvent)
	guildController := crudcontrollers.NewGuildController(databaseResolver, logger, userEvent)
	guildMemberController := crudcontrollers.NewGuildMemberController(databaseResolver, logger, userEvent)
	guildRankController := crudcontrollers.NewGuildRankController(databaseResolver, logger, userEvent)
	guildRelationController := crudcontrollers.NewGuildRelationController(databaseResolver, logger, userEvent)
	hackerController := crudcontrollers.NewHackerController(databaseResolver, logger, userEvent)
	horseController := crudcontrollers.NewHorseController(databaseResolver, logger, userEvent)
	instanceListController := crudcontrollers.NewInstanceListController(databaseResolver, logger, userEvent)
	instanceListPlayerController := crudcontrollers.NewInstanceListPlayerController(databaseResolver, logger, userEvent)
	inventoryController := crudcontrollers.NewInventoryController(databaseResolver, logger, userEvent)
	inventorySnapshotController := crudcontrollers.NewInventorySnapshotController(databaseResolver, logger, userEvent)
	ipExemptionController := crudcontrollers.NewIpExemptionController(databaseResolver, logger, userEvent)
	itemController := crudcontrollers.NewItemController(databaseResolver, logger, userEvent)
	itemTickController := crudcontrollers.NewItemTickController(databaseResolver, logger, userEvent)
	ldonTrapEntryController := crudcontrollers.NewLdonTrapEntryController(databaseResolver, logger, userEvent)
	ldonTrapTemplateController := crudcontrollers.NewLdonTrapTemplateController(databaseResolver, logger, userEvent)
	levelExpModController := crudcontrollers.NewLevelExpModController(databaseResolver, logger, userEvent)
	lfguildController := crudcontrollers.NewLfguildController(databaseResolver, logger, userEvent)
	loginAccountController := crudcontrollers.NewLoginAccountController(databaseResolver, logger, userEvent)
	loginApiTokenController := crudcontrollers.NewLoginApiTokenController(databaseResolver, logger, userEvent)
	loginServerAdminController := crudcontrollers.NewLoginServerAdminController(databaseResolver, logger, userEvent)
	loginServerListTypeController := crudcontrollers.NewLoginServerListTypeController(databaseResolver, logger, userEvent)
	loginWorldServerController := crudcontrollers.NewLoginWorldServerController(databaseResolver, logger, userEvent)
	logsysCategoryController := crudcontrollers.NewLogsysCategoryController(databaseResolver, logger, userEvent)
	lootdropController := crudcontrollers.NewLootdropController(databaseResolver, logger, userEvent)
	lootdropEntryController := crudcontrollers.NewLootdropEntryController(databaseResolver, logger, userEvent)
	loottableController := crudcontrollers.NewLoottableController(databaseResolver, logger, userEvent)
	loottableEntryController := crudcontrollers.NewLoottableEntryController(databaseResolver, logger, userEvent)
	mailController := crudcontrollers.NewMailController(databaseResolver, logger, userEvent)
	merchantlistController := crudcontrollers.NewMerchantlistController(databaseResolver, logger, userEvent)
	merchantlistTempController := crudcontrollers.NewMerchantlistTempController(databaseResolver, logger, userEvent)
	nameFilterController := crudcontrollers.NewNameFilterController(databaseResolver, logger, userEvent)
	npcEmoteController := crudcontrollers.NewNpcEmoteController(databaseResolver, logger, userEvent)
	npcFactionController := crudcontrollers.NewNpcFactionController(databaseResolver, logger, userEvent)
	npcFactionEntryController := crudcontrollers.NewNpcFactionEntryController(databaseResolver, logger, userEvent)
	npcScaleGlobalBaseController := crudcontrollers.NewNpcScaleGlobalBaseController(databaseResolver, logger, userEvent)
	npcSpellController := crudcontrollers.NewNpcSpellController(databaseResolver, logger, userEvent)
	npcSpellsEffectController := crudcontrollers.NewNpcSpellsEffectController(databaseResolver, logger, userEvent)
	npcSpellsEffectsEntryController := crudcontrollers.NewNpcSpellsEffectsEntryController(databaseResolver, logger, userEvent)
	npcSpellsEntryController := crudcontrollers.NewNpcSpellsEntryController(databaseResolver, logger, userEvent)
	npcTypeController := crudcontrollers.NewNpcTypeController(databaseResolver, logger, userEvent)
	npcTypesTintController := crudcontrollers.NewNpcTypesTintController(databaseResolver, logger, userEvent)
	objectContentController := crudcontrollers.NewObjectContentController(databaseResolver, logger, userEvent)
	objectController := crudcontrollers.NewObjectController(databaseResolver, logger, userEvent)
	perlEventExportSettingController := crudcontrollers.NewPerlEventExportSettingController(databaseResolver, logger, userEvent)
	petController := crudcontrollers.NewPetController(databaseResolver, logger, userEvent)
	petitionController := crudcontrollers.NewPetitionController(databaseResolver, logger, userEvent)
	petsBeastlordDatumController := crudcontrollers.NewPetsBeastlordDatumController(databaseResolver, logger, userEvent)
	petsEquipmentsetController := crudcontrollers.NewPetsEquipmentsetController(databaseResolver, logger, userEvent)
	petsEquipmentsetEntryController := crudcontrollers.NewPetsEquipmentsetEntryController(databaseResolver, logger, userEvent)
	playerEventLogController := crudcontrollers.NewPlayerEventLogController(databaseResolver, logger, userEvent)
	playerEventLogSettingController := crudcontrollers.NewPlayerEventLogSettingController(databaseResolver, logger, userEvent)
	playerTitlesetController := crudcontrollers.NewPlayerTitlesetController(databaseResolver, logger, userEvent)
	questGlobalController := crudcontrollers.NewQuestGlobalController(databaseResolver, logger, userEvent)
	raidDetailController := crudcontrollers.NewRaidDetailController(databaseResolver, logger, userEvent)
	raidMemberController := crudcontrollers.NewRaidMemberController(databaseResolver, logger, userEvent)
	reportController := crudcontrollers.NewReportController(databaseResolver, logger, userEvent)
	respawnTimeController := crudcontrollers.NewRespawnTimeController(databaseResolver, logger, userEvent)
	ruleSetController := crudcontrollers.NewRuleSetController(databaseResolver, logger, userEvent)
	ruleValueController := crudcontrollers.NewRuleValueController(databaseResolver, logger, userEvent)
	saylinkController := crudcontrollers.NewSaylinkController(databaseResolver, logger, userEvent)
	serverScheduledEventController := crudcontrollers.NewServerScheduledEventController(databaseResolver, logger, userEvent)
	sharedTaskActivityStateController := crudcontrollers.NewSharedTaskActivityStateController(databaseResolver, logger, userEvent)
	sharedTaskController := crudcontrollers.NewSharedTaskController(databaseResolver, logger, userEvent)
	sharedTaskDynamicZoneController := crudcontrollers.NewSharedTaskDynamicZoneController(databaseResolver, logger, userEvent)
	sharedTaskMemberController := crudcontrollers.NewSharedTaskMemberController(databaseResolver, logger, userEvent)
	skillCapController := crudcontrollers.NewSkillCapController(databaseResolver, logger, userEvent)
	spawn2Controller := crudcontrollers.NewSpawn2Controller(databaseResolver, logger, userEvent)
	spawnConditionController := crudcontrollers.NewSpawnConditionController(databaseResolver, logger, userEvent)
	spawnConditionValueController := crudcontrollers.NewSpawnConditionValueController(databaseResolver, logger, userEvent)
	spawnEventController := crudcontrollers.NewSpawnEventController(databaseResolver, logger, userEvent)
	spawnentryController := crudcontrollers.NewSpawnentryController(databaseResolver, logger, userEvent)
	spawngroupController := crudcontrollers.NewSpawngroupController(databaseResolver, logger, userEvent)
	spellBucketController := crudcontrollers.NewSpellBucketController(databaseResolver, logger, userEvent)
	spellGlobalController := crudcontrollers.NewSpellGlobalController(databaseResolver, logger, userEvent)
	spellsNewController := crudcontrollers.NewSpellsNewController(databaseResolver, logger, userEvent)
	startZoneController := crudcontrollers.NewStartZoneController(databaseResolver, logger, userEvent)
	startingItemController := crudcontrollers.NewStartingItemController(databaseResolver, logger, userEvent)
	taskActivityController := crudcontrollers.NewTaskActivityController(databaseResolver, logger, userEvent)
	taskController := crudcontrollers.NewTaskController(databaseResolver, logger, userEvent)
	tasksetController := crudcontrollers.NewTasksetController(databaseResolver, logger, userEvent)
	timerController := crudcontrollers.NewTimerController(databaseResolver, logger, userEvent)
	titleController := crudcontrollers.NewTitleController(databaseResolver, logger, userEvent)
	traderController := crudcontrollers.NewTraderController(databaseResolver, logger, userEvent)
	tradeskillRecipeController := crudcontrollers.NewTradeskillRecipeController(databaseResolver, logger, userEvent)
	tradeskillRecipeEntryController := crudcontrollers.NewTradeskillRecipeEntryController(databaseResolver, logger, userEvent)
	trapController := crudcontrollers.NewTrapController(databaseResolver, logger, userEvent)
	tributeController := crudcontrollers.NewTributeController(databaseResolver, logger, userEvent)
	tributeLevelController := crudcontrollers.NewTributeLevelController(databaseResolver, logger, userEvent)
	variableController := crudcontrollers.NewVariableController(databaseResolver, logger, userEvent)
	veteranRewardTemplateController := crudcontrollers.NewVeteranRewardTemplateController(databaseResolver, logger, userEvent)
	zoneController := crudcontrollers.NewZoneController(databaseResolver, logger, userEvent)
	zoneFlagController := crudcontrollers.NewZoneFlagController(databaseResolver, logger, userEvent)
	zonePointController := crudcontrollers.NewZonePointController(databaseResolver, logger, userEvent)
	bootCrudControllers := provideCrudControllers(aaAbilityController, aaRankController, aaRankEffectController, aaRankPrereqController, accountController, accountFlagController, accountIpController, accountRewardController, adventureDetailController, adventureMemberController, adventureStatController, adventureTemplateController, adventureTemplateEntryController, adventureTemplateEntryFlavorController, alternateCurrencyController, auraController, baseDatumController, blockedSpellController, bookController, botBuffController, botCreateCombinationController, botDatumController, botGroupController, botGroupMemberController, botGuildMemberController, botHealRotationController, botHealRotationMemberController, botHealRotationTargetController, botInspectMessageController, botInventoryController, botOwnerOptionController, botPetBuffController, botPetController, botPetInventoryController, botSpellCastingChanceController, botSpellSettingController, botSpellsEntryController, botStanceController, botTimerController, bugController, bugReportController, buyerController, charCreateCombinationController, charCreatePointAllocationController, charRecipeListController, characterActivityController, characterAltCurrencyController, characterAlternateAbilityController, characterAuraController, characterBandolierController, characterBindController, characterBuffController, characterCorpseController, characterCorpseItemController, characterCurrencyController, characterDatumController, characterDisciplineController, characterEnabledtaskController, characterExpModifierController, characterExpeditionLockoutController, characterInspectMessageController, characterInstanceSafereturnController, characterItemRecastController, characterLanguageController, characterLeadershipAbilityController, characterMaterialController, characterMemmedSpellController, characterPeqzoneFlagController, characterPetBuffController, characterPetInfoController, characterPetInventoryController, characterPotionbeltController, characterSkillController, characterSpellController, characterTaskController, characterTaskTimerController, chatchannelReservedNameController, completedSharedTaskActivityStateController, completedSharedTaskController, completedSharedTaskMemberController, completedTaskController, contentFlagController, damageshieldtypeController, dataBucketController, dbStrController, discordWebhookController, discoveredItemController, doorController, dynamicZoneController, dynamicZoneMemberController, dynamicZoneTemplateController, eventlogController, expeditionController, expeditionLockoutController, expeditionMemberController, factionAssociationController, factionBaseDatumController, factionListController, factionListModController, factionValueController, fishingController, forageController, friendController, globalLootController, gmIpController, graveyardController, gridController, gridEntryController, groundSpawnController, groupIdController, guildController, guildMemberController, guildRankController, guildRelationController, hackerController, horseController, instanceListController, instanceListPlayerController, inventoryController, inventorySnapshotController, ipExemptionController, itemController, itemTickController, ldonTrapEntryController, ldonTrapTemplateController, levelExpModController, lfguildController, loginAccountController, loginApiTokenController, loginServerAdminController, loginServerListTypeController, loginWorldServerController, logsysCategoryController, lootdropController, lootdropEntryController, loottableController, loottableEntryController, mailController, merchantlistController, merchantlistTempController, nameFilterController, npcEmoteController, npcFactionController, npcFactionEntryController, npcScaleGlobalBaseController, npcSpellController, npcSpellsEffectController, npcSpellsEffectsEntryController, npcSpellsEntryController, npcTypeController, npcTypesTintController, objectContentController, objectController, perlEventExportSettingController, petController, petitionController, petsBeastlordDatumController, petsEquipmentsetController, petsEquipmentsetEntryController, playerEventLogController, playerEventLogSettingController, playerTitlesetController, questGlobalController, raidDetailController, raidMemberController, reportController, respawnTimeController, ruleSetController, ruleValueController, saylinkController, serverScheduledEventController, sharedTaskActivityStateController, sharedTaskController, sharedTaskDynamicZoneController, sharedTaskMemberController, skillCapController, spawn2Controller, spawnConditionController, spawnConditionValueController, spawnEventController, spawnentryController, spawngroupController, spellBucketController, spellGlobalController, spellsNewController, startZoneController, startingItemController, taskActivityController, taskController, tasksetController, timerController, titleController, traderController, tradeskillRecipeController, tradeskillRecipeEntryController, trapController, tributeController, tributeLevelController, variableController, veteranRewardTemplateController, zoneController, zoneFlagController, zonePointController)
	userContextMiddleware := middleware.NewUserContextMiddleware(databaseResolver, cache, logger)
	readOnlyMiddleware := middleware.NewReadOnlyMiddleware(databaseResolver, logger)
	permissionsMiddleware := middleware.NewPermissionsMiddleware(databaseResolver, logger, cache, service)
	requestLogMiddleware := middleware.NewRequestLogMiddleware(client)
	localUserAuthMiddleware := middleware.NewLocalUserAuthMiddleware(databaseResolver, logger, cache, settings, init)
	spireAssets := assets.NewSpireAssets(logger, cache, githubSourceDownloader)
	router := NewRouter(bootAppControllerGroups, bootCrudControllers, userContextMiddleware, readOnlyMiddleware, permissionsMiddleware, requestLogMiddleware, localUserAuthMiddleware, spireAssets)
	server := http.NewServer(logger, router, processManagement)
	httpServeCommand := cmd.NewHttpServeCommand(logger, server)
	routesListCommand := cmd.NewRoutesListCommand(router, logger)
	generateConfigurationCommand := cmd.NewGenerateConfigurationCommand(databaseResolver, logger)
	spireMigrateCommand := cmd.NewSpireMigrateCommand(connections, logger)
	questApiParseCommand := cmd.NewQuestApiParseCommand(logger, parseService)
	questExampleTestCommand := cmd.NewQuestExampleTestCommand(logger, questExamplesGithubSourcer)
	generateRaceModelMapsCommand := cmd.NewGenerateRaceModelMapsCommand(logger)
	changelogCommand := eqemuchangelog.NewChangelogCommand(db, logger, changelog)
	v := ProvideCommands(helloWorldCommand, adminPingOcculus, userCreateCommand, generateModelsCommand, generateControllersCommand, httpServeCommand, routesListCommand, generateConfigurationCommand, spireMigrateCommand, questApiParseCommand, questExampleTestCommand, generateRaceModelMapsCommand, changelogCommand)
	webBoot := desktop.NewWebBoot(logger, server)
	app := NewApplication(db, logger, cache, v, databaseResolver, connections, router, webBoot, init)
	return app, nil
}
