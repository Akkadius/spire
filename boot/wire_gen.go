// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package boot

import (
	"github.com/Akkadius/spire/internal/analytics"
	"github.com/Akkadius/spire/internal/app"
	"github.com/Akkadius/spire/internal/assets"
	"github.com/Akkadius/spire/internal/auditlog"
	"github.com/Akkadius/spire/internal/auth"
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
	"github.com/Akkadius/spire/internal/eqemuserverconfig"
	"github.com/Akkadius/spire/internal/eqtraders"
	"github.com/Akkadius/spire/internal/generators"
	"github.com/Akkadius/spire/internal/github"
	"github.com/Akkadius/spire/internal/http"
	"github.com/Akkadius/spire/internal/http/controllers"
	"github.com/Akkadius/spire/internal/http/crudcontrollers"
	"github.com/Akkadius/spire/internal/http/middleware"
	"github.com/Akkadius/spire/internal/http/staticmaps"
	"github.com/Akkadius/spire/internal/influx"
	"github.com/Akkadius/spire/internal/logger"
	"github.com/Akkadius/spire/internal/pathmgmt"
	"github.com/Akkadius/spire/internal/permissions"
	"github.com/Akkadius/spire/internal/query"
	"github.com/Akkadius/spire/internal/questapi"
	"github.com/Akkadius/spire/internal/spire"
	"github.com/Akkadius/spire/internal/system"
	"github.com/Akkadius/spire/internal/telnet"
	"github.com/Akkadius/spire/internal/user"
	"github.com/Akkadius/spire/internal/websocket"
	"github.com/gertd/go-pluralize"
)

import (
	_ "github.com/go-sql-driver/mysql"
)

// Injectors from wire.go:

func InitializeApplication() (App, error) {
	appLogger := logger.ProvideAppLogger()
	logrusLogger, err := provideLogger()
	if err != nil {
		return App{}, err
	}
	pathManagement := pathmgmt.NewPathManagement(logrusLogger)
	config := eqemuserverconfig.NewConfig(appLogger, pathManagement)
	db, err := provideEQEmuLocalDatabase(config)
	if err != nil {
		return App{}, err
	}
	cache := provideCache()
	mysql := backup.NewMysql(logrusLogger, pathManagement)
	helloWorldCommand := cmd.NewHelloWorldCommand(db, logrusLogger, mysql, pathManagement)
	connections := provideAppDbConnections(config, appLogger)
	encrypter := encryption.NewEncrypter(logrusLogger, config)
	resolver := database.NewResolver(connections, appLogger, encrypter, cache)
	userUser := user.NewUser(resolver, logrusLogger, encrypter, cache)
	createCommand := user.NewCreateCommand(resolver, logrusLogger, encrypter, userUser)
	modelGeneratorCommand := generators.NewModelGeneratorCommand(db, logrusLogger)
	controllerGeneratorCmd := generators.NewControllerGeneratorCommand(db, logrusLogger)
	helloWorldController := controllers.NewHelloWorldController(db, logrusLogger)
	controller := auth.NewController(resolver, userUser, cache)
	meController := user.NewMeController()
	client := influx.NewClient()
	analyticsController := analytics.NewController(client, resolver)
	create := connection.NewCreate(resolver, logrusLogger, encrypter)
	check := connection.NewCheck(resolver, logrusLogger, encrypter)
	pluralizeClient := pluralize.NewClient()
	service := permissions.NewService(resolver, cache, logrusLogger, pluralizeClient)
	settings := spire.NewSettings(connections, logrusLogger)
	init := spire.NewInit(connections, config, logrusLogger, settings, cache, encrypter, create, userUser)
	connectionsController := controllers.NewConnectionsController(resolver, logrusLogger, cache, create, check, service, init, userUser)
	sourceDownloader := github.NewGithubSourceDownloader(logrusLogger, cache)
	parseService := questapi.NewParseService(logrusLogger, cache, sourceDownloader)
	examplesGithubSourcer := questapi.NewExamplesGithubSourcer(logrusLogger, cache, sourceDownloader)
	questapiController := questapi.NewController(logrusLogger, parseService, examplesGithubSourcer)
	appController := app.NewController(cache, init, userUser, settings, resolver)
	queryController := query.NewController(resolver, logrusLogger)
	exporter := clientfiles.NewExporter(logrusLogger)
	importer := clientfiles.NewImporter(logrusLogger)
	clientfilesController := clientfiles.NewController(logrusLogger, exporter, importer, resolver)
	staticMapController := staticmaps.NewStaticMapController(resolver, logrusLogger)
	releases := eqemuanalytics.NewReleases()
	eqemuanalyticsController := eqemuanalytics.NewController(logrusLogger, resolver, releases)
	authedController := eqemuanalytics.NewAuthedController(logrusLogger, resolver)
	changelog := eqemuchangelog.NewChangelog()
	eqemuchangelogController := eqemuchangelog.NewController(logrusLogger, resolver, changelog)
	deployController := deploy.NewDeployController(logrusLogger)
	assetsController := assets.NewController(logrusLogger, resolver)
	permissionsController := permissions.NewController(logrusLogger, resolver, service)
	userController := user.NewController(resolver, logrusLogger, userUser, encrypter)
	settingsController := spire.NewSettingController(resolver, logrusLogger, encrypter, settings)
	telnetClient := telnet.NewClient(logrusLogger)
	eqemuserverClient := eqemuserver.NewClient(telnetClient)
	updater := eqemuserver.NewUpdater(resolver, logrusLogger, config, settings, pathManagement)
	launcher := eqemuserver.NewLauncher(appLogger, config, settings, pathManagement, eqemuserverClient)
	eqemuserverController := eqemuserver.NewController(resolver, logrusLogger, eqemuserverClient, config, pathManagement, settings, updater, launcher)
	publicController := eqemuserver.NewPublicController(resolver, logrusLogger, eqemuserverClient, config, pathManagement, settings, updater)
	eqemuserverconfigController := eqemuserverconfig.NewController(logrusLogger, config)
	backupController := backup.NewController(logrusLogger, mysql, pathManagement)
	handler := websocket.NewHandler(logrusLogger, pathManagement)
	websocketController := websocket.NewController(pathManagement, handler)
	systemController := system.NewController(logrusLogger)
	bootAppControllerGroups := provideControllers(helloWorldController, controller, meController, analyticsController, connectionsController, questapiController, appController, queryController, clientfilesController, staticMapController, eqemuanalyticsController, authedController, eqemuchangelogController, deployController, assetsController, permissionsController, userController, settingsController, eqemuserverController, publicController, eqemuserverconfigController, backupController, websocketController, systemController)
	userEvent := auditlog.NewUserEvent(resolver, cache)
	aaAbilityController := crudcontrollers.NewAaAbilityController(resolver, userEvent)
	aaRankController := crudcontrollers.NewAaRankController(resolver, userEvent)
	aaRankEffectController := crudcontrollers.NewAaRankEffectController(resolver, userEvent)
	aaRankPrereqController := crudcontrollers.NewAaRankPrereqController(resolver, userEvent)
	accountController := crudcontrollers.NewAccountController(resolver, userEvent)
	accountFlagController := crudcontrollers.NewAccountFlagController(resolver, userEvent)
	accountIpController := crudcontrollers.NewAccountIpController(resolver, userEvent)
	accountRewardController := crudcontrollers.NewAccountRewardController(resolver, userEvent)
	adventureDetailController := crudcontrollers.NewAdventureDetailController(resolver, userEvent)
	adventureMemberController := crudcontrollers.NewAdventureMemberController(resolver, userEvent)
	adventureStatController := crudcontrollers.NewAdventureStatController(resolver, userEvent)
	adventureTemplateController := crudcontrollers.NewAdventureTemplateController(resolver, userEvent)
	adventureTemplateEntryController := crudcontrollers.NewAdventureTemplateEntryController(resolver, userEvent)
	adventureTemplateEntryFlavorController := crudcontrollers.NewAdventureTemplateEntryFlavorController(resolver, userEvent)
	alternateCurrencyController := crudcontrollers.NewAlternateCurrencyController(resolver, userEvent)
	auraController := crudcontrollers.NewAuraController(resolver, userEvent)
	baseDatumController := crudcontrollers.NewBaseDatumController(resolver, userEvent)
	blockedSpellController := crudcontrollers.NewBlockedSpellController(resolver, userEvent)
	bookController := crudcontrollers.NewBookController(resolver, userEvent)
	botBuffController := crudcontrollers.NewBotBuffController(resolver, userEvent)
	botCreateCombinationController := crudcontrollers.NewBotCreateCombinationController(resolver, userEvent)
	botDatumController := crudcontrollers.NewBotDatumController(resolver, userEvent)
	botGroupController := crudcontrollers.NewBotGroupController(resolver, logrusLogger, userEvent)
	botGroupMemberController := crudcontrollers.NewBotGroupMemberController(resolver, logrusLogger, userEvent)
	botGuildMemberController := crudcontrollers.NewBotGuildMemberController(resolver, logrusLogger, userEvent)
	botHealRotationController := crudcontrollers.NewBotHealRotationController(resolver, userEvent)
	botHealRotationMemberController := crudcontrollers.NewBotHealRotationMemberController(resolver, userEvent)
	botHealRotationTargetController := crudcontrollers.NewBotHealRotationTargetController(resolver, userEvent)
	botInspectMessageController := crudcontrollers.NewBotInspectMessageController(resolver, userEvent)
	botInventoryController := crudcontrollers.NewBotInventoryController(resolver, userEvent)
	botOwnerOptionController := crudcontrollers.NewBotOwnerOptionController(resolver, userEvent)
	botPetBuffController := crudcontrollers.NewBotPetBuffController(resolver, userEvent)
	botPetController := crudcontrollers.NewBotPetController(resolver, userEvent)
	botPetInventoryController := crudcontrollers.NewBotPetInventoryController(resolver, userEvent)
	botSpellCastingChanceController := crudcontrollers.NewBotSpellCastingChanceController(resolver, userEvent)
	botSpellSettingController := crudcontrollers.NewBotSpellSettingController(resolver, userEvent)
	botSpellsEntryController := crudcontrollers.NewBotSpellsEntryController(resolver, userEvent)
	botStanceController := crudcontrollers.NewBotStanceController(resolver, userEvent)
	botTimerController := crudcontrollers.NewBotTimerController(resolver, userEvent)
	bugController := crudcontrollers.NewBugController(resolver, userEvent)
	bugReportController := crudcontrollers.NewBugReportController(resolver, userEvent)
	buyerController := crudcontrollers.NewBuyerController(resolver, userEvent)
	charCreateCombinationController := crudcontrollers.NewCharCreateCombinationController(resolver, userEvent)
	charCreatePointAllocationController := crudcontrollers.NewCharCreatePointAllocationController(resolver, userEvent)
	charRecipeListController := crudcontrollers.NewCharRecipeListController(resolver, userEvent)
	characterActivityController := crudcontrollers.NewCharacterActivityController(resolver, userEvent)
	characterAltCurrencyController := crudcontrollers.NewCharacterAltCurrencyController(resolver, userEvent)
	characterAlternateAbilityController := crudcontrollers.NewCharacterAlternateAbilityController(resolver, userEvent)
	characterAuraController := crudcontrollers.NewCharacterAuraController(resolver, userEvent)
	characterBandolierController := crudcontrollers.NewCharacterBandolierController(resolver, userEvent)
	characterBindController := crudcontrollers.NewCharacterBindController(resolver, userEvent)
	characterBuffController := crudcontrollers.NewCharacterBuffController(resolver, userEvent)
	characterCorpseController := crudcontrollers.NewCharacterCorpseController(resolver, userEvent)
	characterCorpseItemController := crudcontrollers.NewCharacterCorpseItemController(resolver, userEvent)
	characterCurrencyController := crudcontrollers.NewCharacterCurrencyController(resolver, userEvent)
	characterDatumController := crudcontrollers.NewCharacterDatumController(resolver, userEvent)
	characterDisciplineController := crudcontrollers.NewCharacterDisciplineController(resolver, userEvent)
	characterEnabledtaskController := crudcontrollers.NewCharacterEnabledtaskController(resolver, userEvent)
	characterExpModifierController := crudcontrollers.NewCharacterExpModifierController(resolver, userEvent)
	characterExpeditionLockoutController := crudcontrollers.NewCharacterExpeditionLockoutController(resolver, userEvent)
	characterInspectMessageController := crudcontrollers.NewCharacterInspectMessageController(resolver, userEvent)
	characterInstanceSafereturnController := crudcontrollers.NewCharacterInstanceSafereturnController(resolver, userEvent)
	characterItemRecastController := crudcontrollers.NewCharacterItemRecastController(resolver, userEvent)
	characterLanguageController := crudcontrollers.NewCharacterLanguageController(resolver, userEvent)
	characterLeadershipAbilityController := crudcontrollers.NewCharacterLeadershipAbilityController(resolver, userEvent)
	characterMaterialController := crudcontrollers.NewCharacterMaterialController(resolver, userEvent)
	characterMemmedSpellController := crudcontrollers.NewCharacterMemmedSpellController(resolver, userEvent)
	characterPeqzoneFlagController := crudcontrollers.NewCharacterPeqzoneFlagController(resolver, userEvent)
	characterPetBuffController := crudcontrollers.NewCharacterPetBuffController(resolver, userEvent)
	characterPetInfoController := crudcontrollers.NewCharacterPetInfoController(resolver, userEvent)
	characterPetInventoryController := crudcontrollers.NewCharacterPetInventoryController(resolver, userEvent)
	characterPotionbeltController := crudcontrollers.NewCharacterPotionbeltController(resolver, userEvent)
	characterSkillController := crudcontrollers.NewCharacterSkillController(resolver, userEvent)
	characterSpellController := crudcontrollers.NewCharacterSpellController(resolver, userEvent)
	characterTaskController := crudcontrollers.NewCharacterTaskController(resolver, userEvent)
	characterTaskTimerController := crudcontrollers.NewCharacterTaskTimerController(resolver, userEvent)
	characterTributeController := crudcontrollers.NewCharacterTributeController(resolver, userEvent)
	chatchannelController := crudcontrollers.NewChatchannelController(resolver, userEvent)
	chatchannelReservedNameController := crudcontrollers.NewChatchannelReservedNameController(resolver, userEvent)
	commandSubsettingController := crudcontrollers.NewCommandSubsettingController(resolver, userEvent)
	completedSharedTaskActivityStateController := crudcontrollers.NewCompletedSharedTaskActivityStateController(resolver, userEvent)
	completedSharedTaskController := crudcontrollers.NewCompletedSharedTaskController(resolver, userEvent)
	completedSharedTaskMemberController := crudcontrollers.NewCompletedSharedTaskMemberController(resolver, userEvent)
	completedTaskController := crudcontrollers.NewCompletedTaskController(resolver, userEvent)
	contentFlagController := crudcontrollers.NewContentFlagController(resolver, userEvent)
	damageshieldtypeController := crudcontrollers.NewDamageshieldtypeController(resolver, userEvent)
	dataBucketController := crudcontrollers.NewDataBucketController(resolver, userEvent)
	dbStrController := crudcontrollers.NewDbStrController(resolver, userEvent)
	discordWebhookController := crudcontrollers.NewDiscordWebhookController(resolver, userEvent)
	discoveredItemController := crudcontrollers.NewDiscoveredItemController(resolver, userEvent)
	doorController := crudcontrollers.NewDoorController(resolver, userEvent)
	dynamicZoneController := crudcontrollers.NewDynamicZoneController(resolver, userEvent)
	dynamicZoneMemberController := crudcontrollers.NewDynamicZoneMemberController(resolver, userEvent)
	dynamicZoneTemplateController := crudcontrollers.NewDynamicZoneTemplateController(resolver, userEvent)
	eventlogController := crudcontrollers.NewEventlogController(resolver, logrusLogger, userEvent)
	expeditionController := crudcontrollers.NewExpeditionController(resolver, userEvent)
	expeditionLockoutController := crudcontrollers.NewExpeditionLockoutController(resolver, userEvent)
	expeditionMemberController := crudcontrollers.NewExpeditionMemberController(resolver, logrusLogger)
	factionAssociationController := crudcontrollers.NewFactionAssociationController(resolver, userEvent)
	factionBaseDatumController := crudcontrollers.NewFactionBaseDatumController(resolver, userEvent)
	factionListController := crudcontrollers.NewFactionListController(resolver, userEvent)
	factionListModController := crudcontrollers.NewFactionListModController(resolver, userEvent)
	factionValueController := crudcontrollers.NewFactionValueController(resolver, userEvent)
	fishingController := crudcontrollers.NewFishingController(resolver, userEvent)
	forageController := crudcontrollers.NewForageController(resolver, userEvent)
	friendController := crudcontrollers.NewFriendController(resolver, userEvent)
	globalLootController := crudcontrollers.NewGlobalLootController(resolver, userEvent)
	gmIpController := crudcontrollers.NewGmIpController(resolver, userEvent)
	graveyardController := crudcontrollers.NewGraveyardController(resolver, userEvent)
	gridController := crudcontrollers.NewGridController(resolver, userEvent)
	gridEntryController := crudcontrollers.NewGridEntryController(resolver, userEvent)
	groundSpawnController := crudcontrollers.NewGroundSpawnController(resolver, userEvent)
	groupIdController := crudcontrollers.NewGroupIdController(resolver, userEvent)
	guildController := crudcontrollers.NewGuildController(resolver, userEvent)
	guildMemberController := crudcontrollers.NewGuildMemberController(resolver, userEvent)
	guildRankController := crudcontrollers.NewGuildRankController(resolver, userEvent)
	guildRelationController := crudcontrollers.NewGuildRelationController(resolver, userEvent)
	hackerController := crudcontrollers.NewHackerController(resolver, logrusLogger, userEvent)
	horseController := crudcontrollers.NewHorseController(resolver, userEvent)
	instanceListController := crudcontrollers.NewInstanceListController(resolver, userEvent)
	instanceListPlayerController := crudcontrollers.NewInstanceListPlayerController(resolver, userEvent)
	inventoryController := crudcontrollers.NewInventoryController(resolver, userEvent)
	inventorySnapshotController := crudcontrollers.NewInventorySnapshotController(resolver, userEvent)
	ipExemptionController := crudcontrollers.NewIpExemptionController(resolver, userEvent)
	itemController := crudcontrollers.NewItemController(resolver, userEvent)
	itemTickController := crudcontrollers.NewItemTickController(resolver, userEvent)
	ldonTrapEntryController := crudcontrollers.NewLdonTrapEntryController(resolver, userEvent)
	ldonTrapTemplateController := crudcontrollers.NewLdonTrapTemplateController(resolver, userEvent)
	levelExpModController := crudcontrollers.NewLevelExpModController(resolver, userEvent)
	lfguildController := crudcontrollers.NewLfguildController(resolver, userEvent)
	loginAccountController := crudcontrollers.NewLoginAccountController(resolver, userEvent)
	loginApiTokenController := crudcontrollers.NewLoginApiTokenController(resolver, userEvent)
	loginServerAdminController := crudcontrollers.NewLoginServerAdminController(resolver, userEvent)
	loginServerListTypeController := crudcontrollers.NewLoginServerListTypeController(resolver, userEvent)
	loginWorldServerController := crudcontrollers.NewLoginWorldServerController(resolver, userEvent)
	logsysCategoryController := crudcontrollers.NewLogsysCategoryController(resolver, userEvent)
	lootdropController := crudcontrollers.NewLootdropController(resolver, userEvent)
	lootdropEntryController := crudcontrollers.NewLootdropEntryController(resolver, userEvent)
	loottableController := crudcontrollers.NewLoottableController(resolver, userEvent)
	loottableEntryController := crudcontrollers.NewLoottableEntryController(resolver, userEvent)
	mailController := crudcontrollers.NewMailController(resolver, userEvent)
	merchantlistController := crudcontrollers.NewMerchantlistController(resolver, userEvent)
	merchantlistTempController := crudcontrollers.NewMerchantlistTempController(resolver, userEvent)
	nameFilterController := crudcontrollers.NewNameFilterController(resolver, userEvent)
	npcEmoteController := crudcontrollers.NewNpcEmoteController(resolver, userEvent)
	npcFactionController := crudcontrollers.NewNpcFactionController(resolver, userEvent)
	npcFactionEntryController := crudcontrollers.NewNpcFactionEntryController(resolver, userEvent)
	npcScaleGlobalBaseController := crudcontrollers.NewNpcScaleGlobalBaseController(resolver, userEvent)
	npcSpellController := crudcontrollers.NewNpcSpellController(resolver, userEvent)
	npcSpellsEffectController := crudcontrollers.NewNpcSpellsEffectController(resolver, userEvent)
	npcSpellsEffectsEntryController := crudcontrollers.NewNpcSpellsEffectsEntryController(resolver, userEvent)
	npcSpellsEntryController := crudcontrollers.NewNpcSpellsEntryController(resolver, userEvent)
	npcTypeController := crudcontrollers.NewNpcTypeController(resolver, userEvent)
	npcTypesTintController := crudcontrollers.NewNpcTypesTintController(resolver, userEvent)
	objectContentController := crudcontrollers.NewObjectContentController(resolver, userEvent)
	objectController := crudcontrollers.NewObjectController(resolver, userEvent)
	perlEventExportSettingController := crudcontrollers.NewPerlEventExportSettingController(resolver, userEvent)
	petController := crudcontrollers.NewPetController(resolver, userEvent)
	petitionController := crudcontrollers.NewPetitionController(resolver, userEvent)
	petsBeastlordDatumController := crudcontrollers.NewPetsBeastlordDatumController(resolver, userEvent)
	petsEquipmentsetController := crudcontrollers.NewPetsEquipmentsetController(resolver, userEvent)
	petsEquipmentsetEntryController := crudcontrollers.NewPetsEquipmentsetEntryController(resolver, userEvent)
	playerEventLogController := crudcontrollers.NewPlayerEventLogController(resolver, userEvent)
	playerEventLogSettingController := crudcontrollers.NewPlayerEventLogSettingController(resolver, userEvent)
	playerTitlesetController := crudcontrollers.NewPlayerTitlesetController(resolver, userEvent)
	questGlobalController := crudcontrollers.NewQuestGlobalController(resolver, userEvent)
	raidDetailController := crudcontrollers.NewRaidDetailController(resolver, userEvent)
	raidMemberController := crudcontrollers.NewRaidMemberController(resolver, userEvent)
	reportController := crudcontrollers.NewReportController(resolver, userEvent)
	respawnTimeController := crudcontrollers.NewRespawnTimeController(resolver, userEvent)
	ruleSetController := crudcontrollers.NewRuleSetController(resolver, userEvent)
	ruleValueController := crudcontrollers.NewRuleValueController(resolver, userEvent)
	saylinkController := crudcontrollers.NewSaylinkController(resolver, userEvent)
	serverScheduledEventController := crudcontrollers.NewServerScheduledEventController(resolver, userEvent)
	sharedTaskActivityStateController := crudcontrollers.NewSharedTaskActivityStateController(resolver, userEvent)
	sharedTaskController := crudcontrollers.NewSharedTaskController(resolver, userEvent)
	sharedTaskDynamicZoneController := crudcontrollers.NewSharedTaskDynamicZoneController(resolver, userEvent)
	sharedTaskMemberController := crudcontrollers.NewSharedTaskMemberController(resolver, userEvent)
	skillCapController := crudcontrollers.NewSkillCapController(resolver, userEvent)
	spawn2Controller := crudcontrollers.NewSpawn2Controller(resolver, userEvent)
	spawnConditionController := crudcontrollers.NewSpawnConditionController(resolver, userEvent)
	spawnConditionValueController := crudcontrollers.NewSpawnConditionValueController(resolver, userEvent)
	spawnEventController := crudcontrollers.NewSpawnEventController(resolver, userEvent)
	spawnentryController := crudcontrollers.NewSpawnentryController(resolver, userEvent)
	spawngroupController := crudcontrollers.NewSpawngroupController(resolver, userEvent)
	spellBucketController := crudcontrollers.NewSpellBucketController(resolver, userEvent)
	spellGlobalController := crudcontrollers.NewSpellGlobalController(resolver, userEvent)
	spellsNewController := crudcontrollers.NewSpellsNewController(resolver, userEvent)
	startZoneController := crudcontrollers.NewStartZoneController(resolver, userEvent)
	startingItemController := crudcontrollers.NewStartingItemController(resolver, userEvent)
	taskActivityController := crudcontrollers.NewTaskActivityController(resolver, userEvent)
	taskController := crudcontrollers.NewTaskController(resolver, userEvent)
	tasksetController := crudcontrollers.NewTasksetController(resolver, userEvent)
	timerController := crudcontrollers.NewTimerController(resolver, userEvent)
	titleController := crudcontrollers.NewTitleController(resolver, userEvent)
	traderController := crudcontrollers.NewTraderController(resolver, userEvent)
	tradeskillRecipeController := crudcontrollers.NewTradeskillRecipeController(resolver, userEvent)
	tradeskillRecipeEntryController := crudcontrollers.NewTradeskillRecipeEntryController(resolver, userEvent)
	trapController := crudcontrollers.NewTrapController(resolver, userEvent)
	tributeController := crudcontrollers.NewTributeController(resolver, userEvent)
	tributeLevelController := crudcontrollers.NewTributeLevelController(resolver, userEvent)
	variableController := crudcontrollers.NewVariableController(resolver, userEvent)
	veteranRewardTemplateController := crudcontrollers.NewVeteranRewardTemplateController(resolver, userEvent)
	zoneController := crudcontrollers.NewZoneController(resolver, userEvent)
	zoneFlagController := crudcontrollers.NewZoneFlagController(resolver, userEvent)
	zonePointController := crudcontrollers.NewZonePointController(resolver, userEvent)
	bootCrudControllers := provideCrudControllers(aaAbilityController, aaRankController, aaRankEffectController, aaRankPrereqController, accountController, accountFlagController, accountIpController, accountRewardController, adventureDetailController, adventureMemberController, adventureStatController, adventureTemplateController, adventureTemplateEntryController, adventureTemplateEntryFlavorController, alternateCurrencyController, auraController, baseDatumController, blockedSpellController, bookController, botBuffController, botCreateCombinationController, botDatumController, botGroupController, botGroupMemberController, botGuildMemberController, botHealRotationController, botHealRotationMemberController, botHealRotationTargetController, botInspectMessageController, botInventoryController, botOwnerOptionController, botPetBuffController, botPetController, botPetInventoryController, botSpellCastingChanceController, botSpellSettingController, botSpellsEntryController, botStanceController, botTimerController, bugController, bugReportController, buyerController, charCreateCombinationController, charCreatePointAllocationController, charRecipeListController, characterActivityController, characterAltCurrencyController, characterAlternateAbilityController, characterAuraController, characterBandolierController, characterBindController, characterBuffController, characterCorpseController, characterCorpseItemController, characterCurrencyController, characterDatumController, characterDisciplineController, characterEnabledtaskController, characterExpModifierController, characterExpeditionLockoutController, characterInspectMessageController, characterInstanceSafereturnController, characterItemRecastController, characterLanguageController, characterLeadershipAbilityController, characterMaterialController, characterMemmedSpellController, characterPeqzoneFlagController, characterPetBuffController, characterPetInfoController, characterPetInventoryController, characterPotionbeltController, characterSkillController, characterSpellController, characterTaskController, characterTaskTimerController, characterTributeController, chatchannelController, chatchannelReservedNameController, commandSubsettingController, completedSharedTaskActivityStateController, completedSharedTaskController, completedSharedTaskMemberController, completedTaskController, contentFlagController, damageshieldtypeController, dataBucketController, dbStrController, discordWebhookController, discoveredItemController, doorController, dynamicZoneController, dynamicZoneMemberController, dynamicZoneTemplateController, eventlogController, expeditionController, expeditionLockoutController, expeditionMemberController, factionAssociationController, factionBaseDatumController, factionListController, factionListModController, factionValueController, fishingController, forageController, friendController, globalLootController, gmIpController, graveyardController, gridController, gridEntryController, groundSpawnController, groupIdController, guildController, guildMemberController, guildRankController, guildRelationController, hackerController, horseController, instanceListController, instanceListPlayerController, inventoryController, inventorySnapshotController, ipExemptionController, itemController, itemTickController, ldonTrapEntryController, ldonTrapTemplateController, levelExpModController, lfguildController, loginAccountController, loginApiTokenController, loginServerAdminController, loginServerListTypeController, loginWorldServerController, logsysCategoryController, lootdropController, lootdropEntryController, loottableController, loottableEntryController, mailController, merchantlistController, merchantlistTempController, nameFilterController, npcEmoteController, npcFactionController, npcFactionEntryController, npcScaleGlobalBaseController, npcSpellController, npcSpellsEffectController, npcSpellsEffectsEntryController, npcSpellsEntryController, npcTypeController, npcTypesTintController, objectContentController, objectController, perlEventExportSettingController, petController, petitionController, petsBeastlordDatumController, petsEquipmentsetController, petsEquipmentsetEntryController, playerEventLogController, playerEventLogSettingController, playerTitlesetController, questGlobalController, raidDetailController, raidMemberController, reportController, respawnTimeController, ruleSetController, ruleValueController, saylinkController, serverScheduledEventController, sharedTaskActivityStateController, sharedTaskController, sharedTaskDynamicZoneController, sharedTaskMemberController, skillCapController, spawn2Controller, spawnConditionController, spawnConditionValueController, spawnEventController, spawnentryController, spawngroupController, spellBucketController, spellGlobalController, spellsNewController, startZoneController, startingItemController, taskActivityController, taskController, tasksetController, timerController, titleController, traderController, tradeskillRecipeController, tradeskillRecipeEntryController, trapController, tributeController, tributeLevelController, variableController, veteranRewardTemplateController, zoneController, zoneFlagController, zonePointController)
	contextMiddleware := user.NewContextMiddleware(resolver, cache, logrusLogger)
	readOnlyMiddleware := middleware.NewReadOnlyMiddleware(resolver, logrusLogger)
	permissionsMiddleware := middleware.NewPermissionsMiddleware(resolver, appLogger, cache, service)
	requestLogMiddleware := middleware.NewRequestLogMiddleware(client)
	localUserAuthMiddleware := middleware.NewLocalUserAuthMiddleware(resolver, cache, settings, init)
	spireAssets := assets.NewSpireAssets(logrusLogger, pathManagement)
	router := NewRouter(bootAppControllerGroups, bootCrudControllers, contextMiddleware, readOnlyMiddleware, permissionsMiddleware, requestLogMiddleware, localUserAuthMiddleware, spireAssets)
	questHotReloadWatcher := eqemuserver.NewQuestHotReloadWatcher(appLogger, config, pathManagement, eqemuserverClient, resolver)
	server := http.NewServer(appLogger, router, questHotReloadWatcher)
	httpServeCommand := cmd.NewHttpServeCommand(logrusLogger, server)
	routesListCommand := cmd.NewRoutesListCommand(router, logrusLogger)
	configurationCommand := generators.NewGenerateConfigurationCommand(resolver, logrusLogger)
	migrateCommand := spire.NewMigrateCommand(connections, logrusLogger)
	parseCommand := questapi.NewParseCommand(logrusLogger, parseService)
	exampleTestCommand := questapi.NewExampleTestCommand(logrusLogger, examplesGithubSourcer)
	raceModelMapsCommand := generators.NewRaceModelMapsCommand(logrusLogger)
	changelogCommand := eqemuchangelog.NewChangelogCommand(db, logrusLogger, changelog)
	testFilesystemCommand := cmd.NewTestFilesystemCommand(logrusLogger, pathManagement)
	initCommand := spire.NewInitCommand(logrusLogger, init)
	changePasswordCommand := user.NewChangePasswordCommand(resolver, logrusLogger, encrypter, userUser)
	crashAnalyticsFingerprintBackfillCommand := spire.NewCrashAnalyticsCommand(logrusLogger, pathManagement, resolver)
	updateCommand := eqemuserver.NewUpdateCommand(logrusLogger, config, settings, pathManagement, launcher, updater)
	launcherCmd := eqemuserver.NewLauncherCmd(logrusLogger, launcher)
	launcherShimCmd := eqemuserver.NewLauncherShimCmd(logrusLogger, launcher)
	scrapeCommand := eqtraders.NewScrapeCommand(db, logrusLogger)
	importCommand := eqtraders.NewImportCommand(db, logrusLogger)
	v := ProvideCommands(helloWorldCommand, createCommand, modelGeneratorCommand, controllerGeneratorCmd, httpServeCommand, routesListCommand, configurationCommand, migrateCommand, parseCommand, exampleTestCommand, raceModelMapsCommand, changelogCommand, testFilesystemCommand, initCommand, changePasswordCommand, crashAnalyticsFingerprintBackfillCommand, updateCommand, launcherCmd, launcherShimCmd, scrapeCommand, importCommand)
	webBoot := desktop.NewWebBoot(logrusLogger, server, config)
	bootApp := NewApplication(db, logrusLogger, cache, v, resolver, connections, router, webBoot, init)
	return bootApp, nil
}
