package boot

import (
	"github.com/Akkadius/spire/internal/http/crudcontrollers"
	"github.com/Akkadius/spire/internal/http/routes"
	"github.com/google/wire"
)

type crudControllers struct {
	routes []routes.Controller
}

var httpCrudControllerSet = wire.NewSet(
	crudcontrollers.NewAaAbilityController,
	crudcontrollers.NewAaRankController,
	crudcontrollers.NewAaRankEffectController,
	crudcontrollers.NewAaRankPrereqController,
	crudcontrollers.NewAccountController,
	crudcontrollers.NewAccountFlagController,
	crudcontrollers.NewAccountIpController,
	crudcontrollers.NewAccountRewardController,
	crudcontrollers.NewAdventureDetailController,
	crudcontrollers.NewAdventureMemberController,
	crudcontrollers.NewAdventureStatController,
	crudcontrollers.NewAdventureTemplateController,
	crudcontrollers.NewAdventureTemplateEntryController,
	crudcontrollers.NewAdventureTemplateEntryFlavorController,
	crudcontrollers.NewAlternateCurrencyController,
	crudcontrollers.NewAuraController,
	crudcontrollers.NewBaseDatumController,
	crudcontrollers.NewBlockedSpellController,
	crudcontrollers.NewBookController,
	crudcontrollers.NewBotBuffController,
	crudcontrollers.NewBotCreateCombinationController,
	crudcontrollers.NewBotDatumController,
	crudcontrollers.NewBotGroupController,
	crudcontrollers.NewBotGroupMemberController,
	crudcontrollers.NewBotGuildMemberController,
	crudcontrollers.NewBotHealRotationController,
	crudcontrollers.NewBotHealRotationMemberController,
	crudcontrollers.NewBotHealRotationTargetController,
	crudcontrollers.NewBotInspectMessageController,
	crudcontrollers.NewBotInventoryController,
	crudcontrollers.NewBotOwnerOptionController,
	crudcontrollers.NewBotPetBuffController,
	crudcontrollers.NewBotPetController,
	crudcontrollers.NewBotPetInventoryController,
	crudcontrollers.NewBotSpellCastingChanceController,
	crudcontrollers.NewBotSpellSettingController,
	crudcontrollers.NewBotSpellsEntryController,
	crudcontrollers.NewBotStanceController,
	crudcontrollers.NewBotTimerController,
	crudcontrollers.NewBugController,
	crudcontrollers.NewBugReportController,
	crudcontrollers.NewBuyerController,
	crudcontrollers.NewCharCreateCombinationController,
	crudcontrollers.NewCharCreatePointAllocationController,
	crudcontrollers.NewCharRecipeListController,
	crudcontrollers.NewCharacterActivityController,
	crudcontrollers.NewCharacterAltCurrencyController,
	crudcontrollers.NewCharacterAlternateAbilityController,
	crudcontrollers.NewCharacterAuraController,
	crudcontrollers.NewCharacterBandolierController,
	crudcontrollers.NewCharacterBindController,
	crudcontrollers.NewCharacterBuffController,
	crudcontrollers.NewCharacterCorpseController,
	crudcontrollers.NewCharacterCorpseItemController,
	crudcontrollers.NewCharacterCurrencyController,
	crudcontrollers.NewCharacterDatumController,
	crudcontrollers.NewCharacterDisciplineController,
	crudcontrollers.NewCharacterEnabledtaskController,
	crudcontrollers.NewCharacterExpModifierController,
	crudcontrollers.NewCharacterExpeditionLockoutController,
	crudcontrollers.NewCharacterInspectMessageController,
	crudcontrollers.NewCharacterInstanceSafereturnController,
	crudcontrollers.NewCharacterItemRecastController,
	crudcontrollers.NewCharacterLanguageController,
	crudcontrollers.NewCharacterLeadershipAbilityController,
	crudcontrollers.NewCharacterMaterialController,
	crudcontrollers.NewCharacterMemmedSpellController,
	crudcontrollers.NewCharacterPeqzoneFlagController,
	crudcontrollers.NewCharacterPetBuffController,
	crudcontrollers.NewCharacterPetInfoController,
	crudcontrollers.NewCharacterPetInventoryController,
	crudcontrollers.NewCharacterPotionbeltController,
	crudcontrollers.NewCharacterSkillController,
	crudcontrollers.NewCharacterSpellController,
	crudcontrollers.NewCharacterTaskController,
	crudcontrollers.NewCharacterTaskTimerController,
	crudcontrollers.NewCharacterTributeController,
	crudcontrollers.NewChatchannelController,
	crudcontrollers.NewChatchannelReservedNameController,
	crudcontrollers.NewCommandSubsettingController,
	crudcontrollers.NewCompletedSharedTaskActivityStateController,
	crudcontrollers.NewCompletedSharedTaskController,
	crudcontrollers.NewCompletedSharedTaskMemberController,
	crudcontrollers.NewCompletedTaskController,
	crudcontrollers.NewContentFlagController,
	crudcontrollers.NewDamageshieldtypeController,
	crudcontrollers.NewDataBucketController,
	crudcontrollers.NewDbStrController,
	crudcontrollers.NewDiscordWebhookController,
	crudcontrollers.NewDiscoveredItemController,
	crudcontrollers.NewDoorController,
	crudcontrollers.NewDynamicZoneController,
	crudcontrollers.NewDynamicZoneMemberController,
	crudcontrollers.NewDynamicZoneTemplateController,
	crudcontrollers.NewEventlogController,
	crudcontrollers.NewExpeditionController,
	crudcontrollers.NewExpeditionLockoutController,
	crudcontrollers.NewExpeditionMemberController,
	crudcontrollers.NewFactionAssociationController,
	crudcontrollers.NewFactionBaseDatumController,
	crudcontrollers.NewFactionListController,
	crudcontrollers.NewFactionListModController,
	crudcontrollers.NewFactionValueController,
	crudcontrollers.NewFishingController,
	crudcontrollers.NewForageController,
	crudcontrollers.NewFriendController,
	crudcontrollers.NewGlobalLootController,
	crudcontrollers.NewGmIpController,
	crudcontrollers.NewGraveyardController,
	crudcontrollers.NewGridController,
	crudcontrollers.NewGridEntryController,
	crudcontrollers.NewGroundSpawnController,
	crudcontrollers.NewGroupIdController,
	crudcontrollers.NewGuildController,
	crudcontrollers.NewGuildMemberController,
	crudcontrollers.NewGuildRankController,
	crudcontrollers.NewGuildRelationController,
	crudcontrollers.NewHorseController,
	crudcontrollers.NewInstanceListController,
	crudcontrollers.NewInstanceListPlayerController,
	crudcontrollers.NewInventoryController,
	crudcontrollers.NewInventorySnapshotController,
	crudcontrollers.NewIpExemptionController,
	crudcontrollers.NewItemController,
	crudcontrollers.NewItemTickController,
	crudcontrollers.NewLdonTrapEntryController,
	crudcontrollers.NewLdonTrapTemplateController,
	crudcontrollers.NewLevelExpModController,
	crudcontrollers.NewLfguildController,
	crudcontrollers.NewLoginAccountController,
	crudcontrollers.NewLoginApiTokenController,
	crudcontrollers.NewLoginServerAdminController,
	crudcontrollers.NewLoginServerListTypeController,
	crudcontrollers.NewLoginWorldServerController,
	crudcontrollers.NewLogsysCategoryController,
	crudcontrollers.NewLootdropController,
	crudcontrollers.NewLootdropEntryController,
	crudcontrollers.NewLoottableController,
	crudcontrollers.NewLoottableEntryController,
	crudcontrollers.NewMailController,
	crudcontrollers.NewMerchantlistController,
	crudcontrollers.NewMerchantlistTempController,
	crudcontrollers.NewNameFilterController,
	crudcontrollers.NewNpcEmoteController,
	crudcontrollers.NewNpcFactionController,
	crudcontrollers.NewNpcFactionEntryController,
	crudcontrollers.NewNpcScaleGlobalBaseController,
	crudcontrollers.NewNpcSpellController,
	crudcontrollers.NewNpcSpellsEffectController,
	crudcontrollers.NewNpcSpellsEffectsEntryController,
	crudcontrollers.NewNpcSpellsEntryController,
	crudcontrollers.NewNpcTypeController,
	crudcontrollers.NewNpcTypesTintController,
	crudcontrollers.NewObjectContentController,
	crudcontrollers.NewObjectController,
	crudcontrollers.NewPerlEventExportSettingController,
	crudcontrollers.NewPetController,
	crudcontrollers.NewPetitionController,
	crudcontrollers.NewPetsBeastlordDatumController,
	crudcontrollers.NewPetsEquipmentsetController,
	crudcontrollers.NewPetsEquipmentsetEntryController,
	crudcontrollers.NewPlayerEventLogController,
	crudcontrollers.NewPlayerEventLogSettingController,
	crudcontrollers.NewPlayerTitlesetController,
	crudcontrollers.NewQuestGlobalController,
	crudcontrollers.NewRaidDetailController,
	crudcontrollers.NewRaidMemberController,
	crudcontrollers.NewReportController,
	crudcontrollers.NewRespawnTimeController,
	crudcontrollers.NewRuleSetController,
	crudcontrollers.NewRuleValueController,
	crudcontrollers.NewSaylinkController,
	crudcontrollers.NewServerScheduledEventController,
	crudcontrollers.NewSharedTaskActivityStateController,
	crudcontrollers.NewSharedTaskController,
	crudcontrollers.NewSharedTaskDynamicZoneController,
	crudcontrollers.NewSharedTaskMemberController,
	crudcontrollers.NewSkillCapController,
	crudcontrollers.NewSpawn2Controller,
	crudcontrollers.NewSpawnConditionController,
	crudcontrollers.NewSpawnConditionValueController,
	crudcontrollers.NewSpawnEventController,
	crudcontrollers.NewSpawnentryController,
	crudcontrollers.NewSpawngroupController,
	crudcontrollers.NewSpellBucketController,
	crudcontrollers.NewSpellGlobalController,
	crudcontrollers.NewSpellsNewController,
	crudcontrollers.NewStartZoneController,
	crudcontrollers.NewStartingItemController,
	crudcontrollers.NewTaskActivityController,
	crudcontrollers.NewTaskController,
	crudcontrollers.NewTasksetController,
	crudcontrollers.NewTimerController,
	crudcontrollers.NewTitleController,
	crudcontrollers.NewTraderController,
	crudcontrollers.NewTradeskillRecipeController,
	crudcontrollers.NewTradeskillRecipeEntryController,
	crudcontrollers.NewTrapController,
	crudcontrollers.NewTributeController,
	crudcontrollers.NewTributeLevelController,
	crudcontrollers.NewVariableController,
	crudcontrollers.NewVeteranRewardTemplateController,
	crudcontrollers.NewZoneController,
	crudcontrollers.NewZoneFlagController,
	crudcontrollers.NewZonePointController,
	provideCrudControllers,
)

func provideCrudControllers(
	aaAbilityController *crudcontrollers.AaAbilityController,
	aaRankController *crudcontrollers.AaRankController,
	aaRankEffectController *crudcontrollers.AaRankEffectController,
	aaRankPrereqController *crudcontrollers.AaRankPrereqController,
	accountController *crudcontrollers.AccountController,
	accountFlagController *crudcontrollers.AccountFlagController,
	accountIpController *crudcontrollers.AccountIpController,
	accountRewardController *crudcontrollers.AccountRewardController,
	adventureDetailController *crudcontrollers.AdventureDetailController,
	adventureMemberController *crudcontrollers.AdventureMemberController,
	adventureStatController *crudcontrollers.AdventureStatController,
	adventureTemplateController *crudcontrollers.AdventureTemplateController,
	adventureTemplateEntryController *crudcontrollers.AdventureTemplateEntryController,
	adventureTemplateEntryFlavorController *crudcontrollers.AdventureTemplateEntryFlavorController,
	alternateCurrencyController *crudcontrollers.AlternateCurrencyController,
	auraController *crudcontrollers.AuraController,
	baseDatumController *crudcontrollers.BaseDatumController,
	blockedSpellController *crudcontrollers.BlockedSpellController,
	bookController *crudcontrollers.BookController,
	botBuffController *crudcontrollers.BotBuffController,
	botCreateCombinationController *crudcontrollers.BotCreateCombinationController,
	botDatumController *crudcontrollers.BotDatumController,
	botGroupController *crudcontrollers.BotGroupController,
	botGroupMemberController *crudcontrollers.BotGroupMemberController,
	botGuildMemberController *crudcontrollers.BotGuildMemberController,
	botHealRotationController *crudcontrollers.BotHealRotationController,
	botHealRotationMemberController *crudcontrollers.BotHealRotationMemberController,
	botHealRotationTargetController *crudcontrollers.BotHealRotationTargetController,
	botInspectMessageController *crudcontrollers.BotInspectMessageController,
	botInventoryController *crudcontrollers.BotInventoryController,
	botOwnerOptionController *crudcontrollers.BotOwnerOptionController,
	botPetBuffController *crudcontrollers.BotPetBuffController,
	botPetController *crudcontrollers.BotPetController,
	botPetInventoryController *crudcontrollers.BotPetInventoryController,
	botSpellCastingChanceController *crudcontrollers.BotSpellCastingChanceController,
	botSpellSettingController *crudcontrollers.BotSpellSettingController,
	botSpellsEntryController *crudcontrollers.BotSpellsEntryController,
	botStanceController *crudcontrollers.BotStanceController,
	botTimerController *crudcontrollers.BotTimerController,
	bugController *crudcontrollers.BugController,
	bugReportController *crudcontrollers.BugReportController,
	buyerController *crudcontrollers.BuyerController,
	charCreateCombinationController *crudcontrollers.CharCreateCombinationController,
	charCreatePointAllocationController *crudcontrollers.CharCreatePointAllocationController,
	charRecipeListController *crudcontrollers.CharRecipeListController,
	characterActivityController *crudcontrollers.CharacterActivityController,
	characterAltCurrencyController *crudcontrollers.CharacterAltCurrencyController,
	characterAlternateAbilityController *crudcontrollers.CharacterAlternateAbilityController,
	characterAuraController *crudcontrollers.CharacterAuraController,
	characterBandolierController *crudcontrollers.CharacterBandolierController,
	characterBindController *crudcontrollers.CharacterBindController,
	characterBuffController *crudcontrollers.CharacterBuffController,
	characterCorpseController *crudcontrollers.CharacterCorpseController,
	characterCorpseItemController *crudcontrollers.CharacterCorpseItemController,
	characterCurrencyController *crudcontrollers.CharacterCurrencyController,
	characterDatumController *crudcontrollers.CharacterDatumController,
	characterDisciplineController *crudcontrollers.CharacterDisciplineController,
	characterEnabledtaskController *crudcontrollers.CharacterEnabledtaskController,
	characterExpModifierController *crudcontrollers.CharacterExpModifierController,
	characterExpeditionLockoutController *crudcontrollers.CharacterExpeditionLockoutController,
	characterInspectMessageController *crudcontrollers.CharacterInspectMessageController,
	characterInstanceSafereturnController *crudcontrollers.CharacterInstanceSafereturnController,
	characterItemRecastController *crudcontrollers.CharacterItemRecastController,
	characterLanguageController *crudcontrollers.CharacterLanguageController,
	characterLeadershipAbilityController *crudcontrollers.CharacterLeadershipAbilityController,
	characterMaterialController *crudcontrollers.CharacterMaterialController,
	characterMemmedSpellController *crudcontrollers.CharacterMemmedSpellController,
	characterPeqzoneFlagController *crudcontrollers.CharacterPeqzoneFlagController,
	characterPetBuffController *crudcontrollers.CharacterPetBuffController,
	characterPetInfoController *crudcontrollers.CharacterPetInfoController,
	characterPetInventoryController *crudcontrollers.CharacterPetInventoryController,
	characterPotionbeltController *crudcontrollers.CharacterPotionbeltController,
	characterSkillController *crudcontrollers.CharacterSkillController,
	characterSpellController *crudcontrollers.CharacterSpellController,
	characterTaskController *crudcontrollers.CharacterTaskController,
	characterTaskTimerController *crudcontrollers.CharacterTaskTimerController,
	characterTributeController *crudcontrollers.CharacterTributeController,
	chatchannelController *crudcontrollers.ChatchannelController,
	chatchannelReservedNameController *crudcontrollers.ChatchannelReservedNameController,
	commandSubsettingController *crudcontrollers.CommandSubsettingController,
	completedSharedTaskActivityStateController *crudcontrollers.CompletedSharedTaskActivityStateController,
	completedSharedTaskController *crudcontrollers.CompletedSharedTaskController,
	completedSharedTaskMemberController *crudcontrollers.CompletedSharedTaskMemberController,
	completedTaskController *crudcontrollers.CompletedTaskController,
	contentFlagController *crudcontrollers.ContentFlagController,
	damageshieldtypeController *crudcontrollers.DamageshieldtypeController,
	dataBucketController *crudcontrollers.DataBucketController,
	dbStrController *crudcontrollers.DbStrController,
	discordWebhookController *crudcontrollers.DiscordWebhookController,
	discoveredItemController *crudcontrollers.DiscoveredItemController,
	doorController *crudcontrollers.DoorController,
	dynamicZoneController *crudcontrollers.DynamicZoneController,
	dynamicZoneMemberController *crudcontrollers.DynamicZoneMemberController,
	dynamicZoneTemplateController *crudcontrollers.DynamicZoneTemplateController,
	eventlogController *crudcontrollers.EventlogController,
	expeditionController *crudcontrollers.ExpeditionController,
	expeditionLockoutController *crudcontrollers.ExpeditionLockoutController,
	expeditionMemberController *crudcontrollers.ExpeditionMemberController,
	factionAssociationController *crudcontrollers.FactionAssociationController,
	factionBaseDatumController *crudcontrollers.FactionBaseDatumController,
	factionListController *crudcontrollers.FactionListController,
	factionListModController *crudcontrollers.FactionListModController,
	factionValueController *crudcontrollers.FactionValueController,
	fishingController *crudcontrollers.FishingController,
	forageController *crudcontrollers.ForageController,
	friendController *crudcontrollers.FriendController,
	globalLootController *crudcontrollers.GlobalLootController,
	gmIpController *crudcontrollers.GmIpController,
	graveyardController *crudcontrollers.GraveyardController,
	gridController *crudcontrollers.GridController,
	gridEntryController *crudcontrollers.GridEntryController,
	groundSpawnController *crudcontrollers.GroundSpawnController,
	groupIdController *crudcontrollers.GroupIdController,
	guildController *crudcontrollers.GuildController,
	guildMemberController *crudcontrollers.GuildMemberController,
	guildRankController *crudcontrollers.GuildRankController,
	guildRelationController *crudcontrollers.GuildRelationController,
	horseController *crudcontrollers.HorseController,
	instanceListController *crudcontrollers.InstanceListController,
	instanceListPlayerController *crudcontrollers.InstanceListPlayerController,
	inventoryController *crudcontrollers.InventoryController,
	inventorySnapshotController *crudcontrollers.InventorySnapshotController,
	ipExemptionController *crudcontrollers.IpExemptionController,
	itemController *crudcontrollers.ItemController,
	itemTickController *crudcontrollers.ItemTickController,
	ldonTrapEntryController *crudcontrollers.LdonTrapEntryController,
	ldonTrapTemplateController *crudcontrollers.LdonTrapTemplateController,
	levelExpModController *crudcontrollers.LevelExpModController,
	lfguildController *crudcontrollers.LfguildController,
	loginAccountController *crudcontrollers.LoginAccountController,
	loginApiTokenController *crudcontrollers.LoginApiTokenController,
	loginServerAdminController *crudcontrollers.LoginServerAdminController,
	loginServerListTypeController *crudcontrollers.LoginServerListTypeController,
	loginWorldServerController *crudcontrollers.LoginWorldServerController,
	logsysCategoryController *crudcontrollers.LogsysCategoryController,
	lootdropController *crudcontrollers.LootdropController,
	lootdropEntryController *crudcontrollers.LootdropEntryController,
	loottableController *crudcontrollers.LoottableController,
	loottableEntryController *crudcontrollers.LoottableEntryController,
	mailController *crudcontrollers.MailController,
	merchantlistController *crudcontrollers.MerchantlistController,
	merchantlistTempController *crudcontrollers.MerchantlistTempController,
	nameFilterController *crudcontrollers.NameFilterController,
	npcEmoteController *crudcontrollers.NpcEmoteController,
	npcFactionController *crudcontrollers.NpcFactionController,
	npcFactionEntryController *crudcontrollers.NpcFactionEntryController,
	npcScaleGlobalBaseController *crudcontrollers.NpcScaleGlobalBaseController,
	npcSpellController *crudcontrollers.NpcSpellController,
	npcSpellsEffectController *crudcontrollers.NpcSpellsEffectController,
	npcSpellsEffectsEntryController *crudcontrollers.NpcSpellsEffectsEntryController,
	npcSpellsEntryController *crudcontrollers.NpcSpellsEntryController,
	npcTypeController *crudcontrollers.NpcTypeController,
	npcTypesTintController *crudcontrollers.NpcTypesTintController,
	objectContentController *crudcontrollers.ObjectContentController,
	objectController *crudcontrollers.ObjectController,
	perlEventExportSettingController *crudcontrollers.PerlEventExportSettingController,
	petController *crudcontrollers.PetController,
	petitionController *crudcontrollers.PetitionController,
	petsBeastlordDatumController *crudcontrollers.PetsBeastlordDatumController,
	petsEquipmentsetController *crudcontrollers.PetsEquipmentsetController,
	petsEquipmentsetEntryController *crudcontrollers.PetsEquipmentsetEntryController,
	playerEventLogController *crudcontrollers.PlayerEventLogController,
	playerEventLogSettingController *crudcontrollers.PlayerEventLogSettingController,
	playerTitlesetController *crudcontrollers.PlayerTitlesetController,
	questGlobalController *crudcontrollers.QuestGlobalController,
	raidDetailController *crudcontrollers.RaidDetailController,
	raidMemberController *crudcontrollers.RaidMemberController,
	reportController *crudcontrollers.ReportController,
	respawnTimeController *crudcontrollers.RespawnTimeController,
	ruleSetController *crudcontrollers.RuleSetController,
	ruleValueController *crudcontrollers.RuleValueController,
	saylinkController *crudcontrollers.SaylinkController,
	serverScheduledEventController *crudcontrollers.ServerScheduledEventController,
	sharedTaskActivityStateController *crudcontrollers.SharedTaskActivityStateController,
	sharedTaskController *crudcontrollers.SharedTaskController,
	sharedTaskDynamicZoneController *crudcontrollers.SharedTaskDynamicZoneController,
	sharedTaskMemberController *crudcontrollers.SharedTaskMemberController,
	skillCapController *crudcontrollers.SkillCapController,
	spawn2Controller *crudcontrollers.Spawn2Controller,
	spawnConditionController *crudcontrollers.SpawnConditionController,
	spawnConditionValueController *crudcontrollers.SpawnConditionValueController,
	spawnEventController *crudcontrollers.SpawnEventController,
	spawnentryController *crudcontrollers.SpawnentryController,
	spawngroupController *crudcontrollers.SpawngroupController,
	spellBucketController *crudcontrollers.SpellBucketController,
	spellGlobalController *crudcontrollers.SpellGlobalController,
	spellsNewController *crudcontrollers.SpellsNewController,
	startZoneController *crudcontrollers.StartZoneController,
	startingItemController *crudcontrollers.StartingItemController,
	taskActivityController *crudcontrollers.TaskActivityController,
	taskController *crudcontrollers.TaskController,
	tasksetController *crudcontrollers.TasksetController,
	timerController *crudcontrollers.TimerController,
	titleController *crudcontrollers.TitleController,
	traderController *crudcontrollers.TraderController,
	tradeskillRecipeController *crudcontrollers.TradeskillRecipeController,
	tradeskillRecipeEntryController *crudcontrollers.TradeskillRecipeEntryController,
	trapController *crudcontrollers.TrapController,
	tributeController *crudcontrollers.TributeController,
	tributeLevelController *crudcontrollers.TributeLevelController,
	variableController *crudcontrollers.VariableController,
	veteranRewardTemplateController *crudcontrollers.VeteranRewardTemplateController,
	zoneController *crudcontrollers.ZoneController,
	zoneFlagController *crudcontrollers.ZoneFlagController,
	zonePointController *crudcontrollers.ZonePointController,
) *crudControllers {
	return &crudControllers{
		routes: []routes.Controller{
			aaAbilityController,
			aaRankController,
			aaRankEffectController,
			aaRankPrereqController,
			accountController,
			accountFlagController,
			accountIpController,
			accountRewardController,
			adventureDetailController,
			adventureMemberController,
			adventureStatController,
			adventureTemplateController,
			adventureTemplateEntryController,
			adventureTemplateEntryFlavorController,
			alternateCurrencyController,
			auraController,
			baseDatumController,
			blockedSpellController,
			bookController,
			botBuffController,
			botCreateCombinationController,
			botDatumController,
			botGroupController,
			botGroupMemberController,
			botGuildMemberController,
			botHealRotationController,
			botHealRotationMemberController,
			botHealRotationTargetController,
			botInspectMessageController,
			botInventoryController,
			botOwnerOptionController,
			botPetBuffController,
			botPetController,
			botPetInventoryController,
			botSpellCastingChanceController,
			botSpellSettingController,
			botSpellsEntryController,
			botStanceController,
			botTimerController,
			bugController,
			bugReportController,
			buyerController,
			charCreateCombinationController,
			charCreatePointAllocationController,
			charRecipeListController,
			characterActivityController,
			characterAltCurrencyController,
			characterAlternateAbilityController,
			characterAuraController,
			characterBandolierController,
			characterBindController,
			characterBuffController,
			characterCorpseController,
			characterCorpseItemController,
			characterCurrencyController,
			characterDatumController,
			characterDisciplineController,
			characterEnabledtaskController,
			characterExpModifierController,
			characterExpeditionLockoutController,
			characterInspectMessageController,
			characterInstanceSafereturnController,
			characterItemRecastController,
			characterLanguageController,
			characterLeadershipAbilityController,
			characterMaterialController,
			characterMemmedSpellController,
			characterPeqzoneFlagController,
			characterPetBuffController,
			characterPetInfoController,
			characterPetInventoryController,
			characterPotionbeltController,
			characterSkillController,
			characterSpellController,
			characterTaskController,
			characterTaskTimerController,
			characterTributeController,
			chatchannelController,
			chatchannelReservedNameController,
			commandSubsettingController,
			completedSharedTaskActivityStateController,
			completedSharedTaskController,
			completedSharedTaskMemberController,
			completedTaskController,
			contentFlagController,
			damageshieldtypeController,
			dataBucketController,
			dbStrController,
			discordWebhookController,
			discoveredItemController,
			doorController,
			dynamicZoneController,
			dynamicZoneMemberController,
			dynamicZoneTemplateController,
			eventlogController,
			expeditionController,
			expeditionLockoutController,
			expeditionMemberController,
			factionAssociationController,
			factionBaseDatumController,
			factionListController,
			factionListModController,
			factionValueController,
			fishingController,
			forageController,
			friendController,
			globalLootController,
			gmIpController,
			graveyardController,
			gridController,
			gridEntryController,
			groundSpawnController,
			groupIdController,
			guildController,
			guildMemberController,
			guildRankController,
			guildRelationController,
			horseController,
			instanceListController,
			instanceListPlayerController,
			inventoryController,
			inventorySnapshotController,
			ipExemptionController,
			itemController,
			itemTickController,
			ldonTrapEntryController,
			ldonTrapTemplateController,
			levelExpModController,
			lfguildController,
			loginAccountController,
			loginApiTokenController,
			loginServerAdminController,
			loginServerListTypeController,
			loginWorldServerController,
			logsysCategoryController,
			lootdropController,
			lootdropEntryController,
			loottableController,
			loottableEntryController,
			mailController,
			merchantlistController,
			merchantlistTempController,
			nameFilterController,
			npcEmoteController,
			npcFactionController,
			npcFactionEntryController,
			npcScaleGlobalBaseController,
			npcSpellController,
			npcSpellsEffectController,
			npcSpellsEffectsEntryController,
			npcSpellsEntryController,
			npcTypeController,
			npcTypesTintController,
			objectContentController,
			objectController,
			perlEventExportSettingController,
			petController,
			petitionController,
			petsBeastlordDatumController,
			petsEquipmentsetController,
			petsEquipmentsetEntryController,
			playerEventLogController,
			playerEventLogSettingController,
			playerTitlesetController,
			questGlobalController,
			raidDetailController,
			raidMemberController,
			reportController,
			respawnTimeController,
			ruleSetController,
			ruleValueController,
			saylinkController,
			serverScheduledEventController,
			sharedTaskActivityStateController,
			sharedTaskController,
			sharedTaskDynamicZoneController,
			sharedTaskMemberController,
			skillCapController,
			spawn2Controller,
			spawnConditionController,
			spawnConditionValueController,
			spawnEventController,
			spawnentryController,
			spawngroupController,
			spellBucketController,
			spellGlobalController,
			spellsNewController,
			startZoneController,
			startingItemController,
			taskActivityController,
			taskController,
			tasksetController,
			timerController,
			titleController,
			traderController,
			tradeskillRecipeController,
			tradeskillRecipeEntryController,
			trapController,
			tributeController,
			tributeLevelController,
			variableController,
			veteranRewardTemplateController,
			zoneController,
			zoneFlagController,
			zonePointController,
		},
	}
}
