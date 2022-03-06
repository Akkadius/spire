package models

import (
	"github.com/volatiletech/null/v8"
)

type CharacterDatum struct {
	ID                           uint                         `json:"id" gorm:"Column:id"`
	AccountId                    int                          `json:"account_id" gorm:"Column:account_id"`
	Name                         string                       `json:"name" gorm:"Column:name"`
	LastName                     string                       `json:"last_name" gorm:"Column:last_name"`
	Title                        string                       `json:"title" gorm:"Column:title"`
	Suffix                       string                       `json:"suffix" gorm:"Column:suffix"`
	ZoneId                       uint                         `json:"zone_id" gorm:"Column:zone_id"`
	ZoneInstance                 uint                         `json:"zone_instance" gorm:"Column:zone_instance"`
	Y                            float32                      `json:"y" gorm:"Column:y"`
	X                            float32                      `json:"x" gorm:"Column:x"`
	Z                            float32                      `json:"z" gorm:"Column:z"`
	Heading                      float32                      `json:"heading" gorm:"Column:heading"`
	Gender                       uint8                        `json:"gender" gorm:"Column:gender"`
	Race                         uint16                       `json:"race" gorm:"Column:race"`
	Class                        uint8                        `json:"class" gorm:"Column:class"`
	Level                        uint                         `json:"level" gorm:"Column:level"`
	Deity                        uint                         `json:"deity" gorm:"Column:deity"`
	Birthday                     uint                         `json:"birthday" gorm:"Column:birthday"`
	LastLogin                    uint                         `json:"last_login" gorm:"Column:last_login"`
	TimePlayed                   uint                         `json:"time_played" gorm:"Column:time_played"`
	Level2                       uint8                        `json:"level_2" gorm:"Column:level2"`
	Anon                         uint8                        `json:"anon" gorm:"Column:anon"`
	Gm                           uint8                        `json:"gm" gorm:"Column:gm"`
	Face                         uint                         `json:"face" gorm:"Column:face"`
	HairColor                    uint8                        `json:"hair_color" gorm:"Column:hair_color"`
	HairStyle                    uint8                        `json:"hair_style" gorm:"Column:hair_style"`
	Beard                        uint8                        `json:"beard" gorm:"Column:beard"`
	BeardColor                   uint8                        `json:"beard_color" gorm:"Column:beard_color"`
	EyeColor1                    uint8                        `json:"eye_color_1" gorm:"Column:eye_color_1"`
	EyeColor2                    uint8                        `json:"eye_color_2" gorm:"Column:eye_color_2"`
	DrakkinHeritage              uint                         `json:"drakkin_heritage" gorm:"Column:drakkin_heritage"`
	DrakkinTattoo                uint                         `json:"drakkin_tattoo" gorm:"Column:drakkin_tattoo"`
	DrakkinDetails               uint                         `json:"drakkin_details" gorm:"Column:drakkin_details"`
	AbilityTimeSeconds           uint8                        `json:"ability_time_seconds" gorm:"Column:ability_time_seconds"`
	AbilityNumber                uint8                        `json:"ability_number" gorm:"Column:ability_number"`
	AbilityTimeMinutes           uint8                        `json:"ability_time_minutes" gorm:"Column:ability_time_minutes"`
	AbilityTimeHours             uint8                        `json:"ability_time_hours" gorm:"Column:ability_time_hours"`
	Exp                          uint                         `json:"exp" gorm:"Column:exp"`
	AaPointsSpent                uint                         `json:"aa_points_spent" gorm:"Column:aa_points_spent"`
	AaExp                        uint                         `json:"aa_exp" gorm:"Column:aa_exp"`
	AaPoints                     uint                         `json:"aa_points" gorm:"Column:aa_points"`
	GroupLeadershipExp           uint                         `json:"group_leadership_exp" gorm:"Column:group_leadership_exp"`
	RaidLeadershipExp            uint                         `json:"raid_leadership_exp" gorm:"Column:raid_leadership_exp"`
	GroupLeadershipPoints        uint                         `json:"group_leadership_points" gorm:"Column:group_leadership_points"`
	RaidLeadershipPoints         uint                         `json:"raid_leadership_points" gorm:"Column:raid_leadership_points"`
	Points                       uint                         `json:"points" gorm:"Column:points"`
	CurHp                        uint                         `json:"cur_hp" gorm:"Column:cur_hp"`
	Mana                         uint                         `json:"mana" gorm:"Column:mana"`
	Endurance                    uint                         `json:"endurance" gorm:"Column:endurance"`
	Intoxication                 uint                         `json:"intoxication" gorm:"Column:intoxication"`
	Str                          uint                         `json:"str" gorm:"Column:str"`
	Sta                          uint                         `json:"sta" gorm:"Column:sta"`
	Cha                          uint                         `json:"cha" gorm:"Column:cha"`
	Dex                          uint                         `json:"dex" gorm:"Column:dex"`
	Int                          uint                         `json:"int" gorm:"Column:int"`
	Agi                          uint                         `json:"agi" gorm:"Column:agi"`
	Wis                          uint                         `json:"wis" gorm:"Column:wis"`
	ZoneChangeCount              uint                         `json:"zone_change_count" gorm:"Column:zone_change_count"`
	Toxicity                     uint                         `json:"toxicity" gorm:"Column:toxicity"`
	HungerLevel                  uint                         `json:"hunger_level" gorm:"Column:hunger_level"`
	ThirstLevel                  uint                         `json:"thirst_level" gorm:"Column:thirst_level"`
	AbilityUp                    uint                         `json:"ability_up" gorm:"Column:ability_up"`
	LdonPointsGuk                uint                         `json:"ldon_points_guk" gorm:"Column:ldon_points_guk"`
	LdonPointsMir                uint                         `json:"ldon_points_mir" gorm:"Column:ldon_points_mir"`
	LdonPointsMmc                uint                         `json:"ldon_points_mmc" gorm:"Column:ldon_points_mmc"`
	LdonPointsRuj                uint                         `json:"ldon_points_ruj" gorm:"Column:ldon_points_ruj"`
	LdonPointsTak                uint                         `json:"ldon_points_tak" gorm:"Column:ldon_points_tak"`
	LdonPointsAvailable          uint                         `json:"ldon_points_available" gorm:"Column:ldon_points_available"`
	TributeTimeRemaining         uint                         `json:"tribute_time_remaining" gorm:"Column:tribute_time_remaining"`
	CareerTributePoints          uint                         `json:"career_tribute_points" gorm:"Column:career_tribute_points"`
	TributePoints                uint                         `json:"tribute_points" gorm:"Column:tribute_points"`
	TributeActive                uint                         `json:"tribute_active" gorm:"Column:tribute_active"`
	PvpStatus                    uint8                        `json:"pvp_status" gorm:"Column:pvp_status"`
	PvpKills                     uint                         `json:"pvp_kills" gorm:"Column:pvp_kills"`
	PvpDeaths                    uint                         `json:"pvp_deaths" gorm:"Column:pvp_deaths"`
	PvpCurrentPoints             uint                         `json:"pvp_current_points" gorm:"Column:pvp_current_points"`
	PvpCareerPoints              uint                         `json:"pvp_career_points" gorm:"Column:pvp_career_points"`
	PvpBestKillStreak            uint                         `json:"pvp_best_kill_streak" gorm:"Column:pvp_best_kill_streak"`
	PvpWorstDeathStreak          uint                         `json:"pvp_worst_death_streak" gorm:"Column:pvp_worst_death_streak"`
	PvpCurrentKillStreak         uint                         `json:"pvp_current_kill_streak" gorm:"Column:pvp_current_kill_streak"`
	Pvp2                         uint                         `json:"pvp_2" gorm:"Column:pvp2"`
	PvpType                      uint                         `json:"pvp_type" gorm:"Column:pvp_type"`
	ShowHelm                     uint                         `json:"show_helm" gorm:"Column:show_helm"`
	GroupAutoConsent             uint8                        `json:"group_auto_consent" gorm:"Column:group_auto_consent"`
	RaidAutoConsent              uint8                        `json:"raid_auto_consent" gorm:"Column:raid_auto_consent"`
	GuildAutoConsent             uint8                        `json:"guild_auto_consent" gorm:"Column:guild_auto_consent"`
	LeadershipExpOn              uint8                        `json:"leadership_exp_on" gorm:"Column:leadership_exp_on"`
	RestTimer                    uint                         `json:"rest_timer" gorm:"Column:RestTimer"`
	AirRemaining                 uint                         `json:"air_remaining" gorm:"Column:air_remaining"`
	AutosplitEnabled             uint                         `json:"autosplit_enabled" gorm:"Column:autosplit_enabled"`
	Lfp                          uint8                        `json:"lfp" gorm:"Column:lfp"`
	Lfg                          uint8                        `json:"lfg" gorm:"Column:lfg"`
	Mailkey                      string                       `json:"mailkey" gorm:"Column:mailkey"`
	Xtargets                     uint8                        `json:"xtargets" gorm:"Column:xtargets"`
	Firstlogon                   int8                         `json:"firstlogon" gorm:"Column:firstlogon"`
	EAaEffects                   uint                         `json:"e_aa_effects" gorm:"Column:e_aa_effects"`
	EPercentToAa                 uint                         `json:"e_percent_to_aa" gorm:"Column:e_percent_to_aa"`
	EExpendedAaSpent             uint                         `json:"e_expended_aa_spent" gorm:"Column:e_expended_aa_spent"`
	AaPointsSpentOld             uint                         `json:"aa_points_spent_old" gorm:"Column:aa_points_spent_old"`
	AaPointsOld                  uint                         `json:"aa_points_old" gorm:"Column:aa_points_old"`
	ELastInvsnapshot             uint                         `json:"e_last_invsnapshot" gorm:"Column:e_last_invsnapshot"`
	DeletedAt                    null.Time                    `json:"deleted_at" gorm:"Column:deleted_at"`
	Guild                        *Guild                       `json:"guild,omitempty" gorm:"foreignKey:id;references:id"`
	CharRecipeLists              []CharRecipeList             `json:"char_recipe_lists,omitempty" gorm:"foreignKey:char_id;references:id"`
	CharacterAltCurrencies       []CharacterAltCurrency       `json:"character_alt_currencies,omitempty" gorm:"foreignKey:char_id;references:id"`
	CharacterPetBuffs            []CharacterPetBuff           `json:"character_pet_buffs,omitempty" gorm:"foreignKey:char_id;references:id"`
	CharacterPetInfos            []CharacterPetInfo           `json:"character_pet_infos,omitempty" gorm:"foreignKey:char_id;references:id"`
	CharacterPetInventories      []CharacterPetInventory      `json:"character_pet_inventories,omitempty" gorm:"foreignKey:char_id;references:id"`
	FactionValues                []FactionValue               `json:"faction_values,omitempty" gorm:"foreignKey:char_id;references:id"`
	GuildMembers                 []GuildMember                `json:"guild_members,omitempty" gorm:"foreignKey:char_id;references:id"`
	Keyrings                     []Keyring                    `json:"keyrings,omitempty" gorm:"foreignKey:char_id;references:id"`
	PlayerTitlesets              []PlayerTitleset             `json:"player_titlesets,omitempty" gorm:"foreignKey:char_id;references:id"`
	Titles                       []Title                      `json:"titles,omitempty" gorm:"foreignKey:char_id;references:id"`
	Traders                      []Trader                     `json:"traders,omitempty" gorm:"foreignKey:char_id;references:id"`
	CharacterBuffs               []CharacterBuff              `json:"character_buffs,omitempty" gorm:"foreignKey:character_id;references:id"`
	Buyers                       []Buyer                      `json:"buyers,omitempty" gorm:"foreignKey:charid;references:id"`
	CharacterActivities          []CharacterActivity          `json:"character_activities,omitempty" gorm:"foreignKey:charid;references:id"`
	CharacterEnabledtasks        []CharacterEnabledtask       `json:"character_enabledtasks,omitempty" gorm:"foreignKey:charid;references:id"`
	CharacterTasks               []CharacterTask              `json:"character_tasks,omitempty" gorm:"foreignKey:charid;references:id"`
	CompletedTasks               []CompletedTask              `json:"completed_tasks,omitempty" gorm:"foreignKey:charid;references:id"`
	Friends                      []Friend                     `json:"friends,omitempty" gorm:"foreignKey:charid;references:id"`
	Inventories                  []Inventory                  `json:"inventories,omitempty" gorm:"foreignKey:charid;references:id"`
	Mail                         []Mail                       `json:"mail,omitempty" gorm:"foreignKey:charid;references:id"`
	QuestGlobals                 []QuestGlobal                `json:"quest_globals,omitempty" gorm:"foreignKey:charid;references:id"`
	ZoneFlags                    []ZoneFlag                   `json:"zone_flags,omitempty" gorm:"foreignKey:charID;references:id"`
	CharacterAlternateAbilities  []CharacterAlternateAbility  `json:"character_alternate_abilities,omitempty" gorm:"foreignKey:id;references:id"`
	CharacterAuras               []CharacterAura              `json:"character_auras,omitempty" gorm:"foreignKey:id;references:id"`
	CharacterBandoliers          []CharacterBandolier         `json:"character_bandoliers,omitempty" gorm:"foreignKey:id;references:id"`
	CharacterBinds               []CharacterBind              `json:"character_binds,omitempty" gorm:"foreignKey:id;references:id"`
	CharacterCorpses             []CharacterCorpse            `json:"character_corpses,omitempty" gorm:"foreignKey:id;references:id"`
	CharacterCurrencies          []CharacterCurrency          `json:"character_currencies,omitempty" gorm:"foreignKey:id;references:id"`
	CharacterDisciplines         []CharacterDiscipline        `json:"character_disciplines,omitempty" gorm:"foreignKey:id;references:id"`
	CharacterInspectMessages     []CharacterInspectMessage    `json:"character_inspect_messages,omitempty" gorm:"foreignKey:id;references:id"`
	CharacterItemRecasts         []CharacterItemRecast        `json:"character_item_recasts,omitempty" gorm:"foreignKey:id;references:id"`
	CharacterLanguages           []CharacterLanguage          `json:"character_languages,omitempty" gorm:"foreignKey:id;references:id"`
	CharacterLeadershipAbilities []CharacterLeadershipAbility `json:"character_leadership_abilities,omitempty" gorm:"foreignKey:id;references:id"`
	CharacterMaterials           []CharacterMaterial          `json:"character_materials,omitempty" gorm:"foreignKey:id;references:id"`
	CharacterMemmedSpells        []CharacterMemmedSpell       `json:"character_memmed_spells,omitempty" gorm:"foreignKey:id;references:id"`
	CharacterPotionbelts         []CharacterPotionbelt        `json:"character_potionbelts,omitempty" gorm:"foreignKey:id;references:id"`
	CharacterSkills              []CharacterSkill             `json:"character_skills,omitempty" gorm:"foreignKey:id;references:id"`
	CharacterSpells              []CharacterSpell             `json:"character_spells,omitempty" gorm:"foreignKey:id;references:id"`
	CharacterTributes            []CharacterTribute           `json:"character_tributes,omitempty" gorm:"foreignKey:id;references:id"`
	DataBuckets                  []DataBucket                 `json:"data_buckets,omitempty" gorm:"foreignKey:id;references:id"`
	InstanceListPlayers          []InstanceListPlayer         `json:"instance_list_players,omitempty" gorm:"foreignKey:id;references:id"`
	AdventureStats               []AdventureStat              `json:"adventure_stats,omitempty" gorm:"foreignKey:player_id;references:id"`
	Timers                       []Timer                      `json:"timers,omitempty" gorm:"foreignKey:char_id;references:id"`
}

func (CharacterDatum) TableName() string {
    return "character_data"
}

func (CharacterDatum) Relationships() []string {
    return []string{
		"AdventureStats",
		"Buyers",
		"CharRecipeLists",
		"CharacterActivities",
		"CharacterAltCurrencies",
		"CharacterAlternateAbilities",
		"CharacterAuras",
		"CharacterBandoliers",
		"CharacterBinds",
		"CharacterBuffs",
		"CharacterCorpses",
		"CharacterCurrencies",
		"CharacterDisciplines",
		"CharacterEnabledtasks",
		"CharacterInspectMessages",
		"CharacterItemRecasts",
		"CharacterLanguages",
		"CharacterLeadershipAbilities",
		"CharacterMaterials",
		"CharacterMemmedSpells",
		"CharacterPetBuffs",
		"CharacterPetInfos",
		"CharacterPetInventories",
		"CharacterPotionbelts",
		"CharacterSkills",
		"CharacterSpells",
		"CharacterTasks",
		"CharacterTributes",
		"CompletedTasks",
		"DataBuckets",
		"FactionValues",
		"Friends",
		"Guild",
		"Guild.GuildBanks",
		"Guild.GuildMembers",
		"Guild.GuildRanks",
		"GuildMembers",
		"InstanceListPlayers",
		"Inventories",
		"Inventories.Item",
		"Inventories.Item.AlternateCurrencies",
		"Inventories.Item.CharacterCorpseItems",
		"Inventories.Item.DiscoveredItems",
		"Inventories.Item.Doors",
		"Inventories.Item.Doors.Item",
		"Inventories.Item.Fishings",
		"Inventories.Item.Fishings.Item",
		"Inventories.Item.Fishings.NpcType",
		"Inventories.Item.Fishings.NpcType.AlternateCurrency",
		"Inventories.Item.Fishings.NpcType.Loottable",
		"Inventories.Item.Fishings.NpcType.Loottable.LoottableEntries",
		"Inventories.Item.Fishings.NpcType.Loottable.LoottableEntries.LootdropEntries",
		"Inventories.Item.Fishings.NpcType.Loottable.LoottableEntries.LootdropEntries.Item",
		"Inventories.Item.Fishings.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop",
		"Inventories.Item.Fishings.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LootdropEntries",
		"Inventories.Item.Fishings.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LoottableEntries",
		"Inventories.Item.Fishings.NpcType.Loottable.LoottableEntries.Loottable",
		"Inventories.Item.Fishings.NpcType.Loottable.NpcTypes",
		"Inventories.Item.Fishings.NpcType.Merchantlists",
		"Inventories.Item.Fishings.NpcType.Merchantlists.NpcType",
		"Inventories.Item.Fishings.NpcType.NpcEmotes",
		"Inventories.Item.Fishings.NpcType.NpcFactions",
		"Inventories.Item.Fishings.NpcType.NpcFactions.NpcFactionEntries",
		"Inventories.Item.Fishings.NpcType.NpcSpells",
		"Inventories.Item.Fishings.NpcType.NpcSpells.NpcSpellsEntries",
		"Inventories.Item.Fishings.NpcType.NpcTypesTint",
		"Inventories.Item.Fishings.NpcType.Spawnentries",
		"Inventories.Item.Fishings.NpcType.Spawnentries.NpcType",
		"Inventories.Item.Fishings.NpcType.Spawnentries.Spawngroup",
		"Inventories.Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2",
		"Inventories.Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries",
		"Inventories.Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup",
		"Inventories.Item.Fishings.Zone",
		"Inventories.Item.Forages",
		"Inventories.Item.Forages.Item",
		"Inventories.Item.Forages.Zone",
		"Inventories.Item.GroundSpawns",
		"Inventories.Item.GroundSpawns.Zone",
		"Inventories.Item.ItemTicks",
		"Inventories.Item.Keyrings",
		"Inventories.Item.LootdropEntries",
		"Inventories.Item.LootdropEntries.Item",
		"Inventories.Item.LootdropEntries.Lootdrop",
		"Inventories.Item.LootdropEntries.Lootdrop.LootdropEntries",
		"Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries",
		"Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.LootdropEntries",
		"Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable",
		"Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.LoottableEntries",
		"Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes",
		"Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency",
		"Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Loottable",
		"Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists",
		"Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.NpcType",
		"Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcEmotes",
		"Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions",
		"Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions.NpcFactionEntries",
		"Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpells",
		"Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpells.NpcSpellsEntries",
		"Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcTypesTint",
		"Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries",
		"Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.NpcType",
		"Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup",
		"Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2",
		"Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries",
		"Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup",
		"Inventories.Item.Merchantlists",
		"Inventories.Item.Merchantlists.NpcType",
		"Inventories.Item.Merchantlists.NpcType.AlternateCurrency",
		"Inventories.Item.Merchantlists.NpcType.Loottable",
		"Inventories.Item.Merchantlists.NpcType.Loottable.LoottableEntries",
		"Inventories.Item.Merchantlists.NpcType.Loottable.LoottableEntries.LootdropEntries",
		"Inventories.Item.Merchantlists.NpcType.Loottable.LoottableEntries.LootdropEntries.Item",
		"Inventories.Item.Merchantlists.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop",
		"Inventories.Item.Merchantlists.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LootdropEntries",
		"Inventories.Item.Merchantlists.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LoottableEntries",
		"Inventories.Item.Merchantlists.NpcType.Loottable.LoottableEntries.Loottable",
		"Inventories.Item.Merchantlists.NpcType.Loottable.NpcTypes",
		"Inventories.Item.Merchantlists.NpcType.Merchantlists",
		"Inventories.Item.Merchantlists.NpcType.NpcEmotes",
		"Inventories.Item.Merchantlists.NpcType.NpcFactions",
		"Inventories.Item.Merchantlists.NpcType.NpcFactions.NpcFactionEntries",
		"Inventories.Item.Merchantlists.NpcType.NpcSpells",
		"Inventories.Item.Merchantlists.NpcType.NpcSpells.NpcSpellsEntries",
		"Inventories.Item.Merchantlists.NpcType.NpcTypesTint",
		"Inventories.Item.Merchantlists.NpcType.Spawnentries",
		"Inventories.Item.Merchantlists.NpcType.Spawnentries.NpcType",
		"Inventories.Item.Merchantlists.NpcType.Spawnentries.Spawngroup",
		"Inventories.Item.Merchantlists.NpcType.Spawnentries.Spawngroup.Spawn2",
		"Inventories.Item.Merchantlists.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries",
		"Inventories.Item.Merchantlists.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup",
		"Inventories.Item.ObjectContents",
		"Inventories.Item.Objects",
		"Inventories.Item.Objects.Item",
		"Inventories.Item.Objects.Zone",
		"Inventories.Item.StartingItems",
		"Inventories.Item.StartingItems.Item",
		"Inventories.Item.StartingItems.Zone",
		"Inventories.Item.Tasks",
		"Inventories.Item.Tasks.TaskActivities",
		"Inventories.Item.Tasks.TaskActivities.Goallists",
		"Inventories.Item.Tasks.TaskActivities.NpcType",
		"Inventories.Item.Tasks.TaskActivities.NpcType.AlternateCurrency",
		"Inventories.Item.Tasks.TaskActivities.NpcType.Loottable",
		"Inventories.Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries",
		"Inventories.Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries",
		"Inventories.Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries.Item",
		"Inventories.Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop",
		"Inventories.Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LootdropEntries",
		"Inventories.Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LoottableEntries",
		"Inventories.Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.Loottable",
		"Inventories.Item.Tasks.TaskActivities.NpcType.Loottable.NpcTypes",
		"Inventories.Item.Tasks.TaskActivities.NpcType.Merchantlists",
		"Inventories.Item.Tasks.TaskActivities.NpcType.Merchantlists.NpcType",
		"Inventories.Item.Tasks.TaskActivities.NpcType.NpcEmotes",
		"Inventories.Item.Tasks.TaskActivities.NpcType.NpcFactions",
		"Inventories.Item.Tasks.TaskActivities.NpcType.NpcFactions.NpcFactionEntries",
		"Inventories.Item.Tasks.TaskActivities.NpcType.NpcSpells",
		"Inventories.Item.Tasks.TaskActivities.NpcType.NpcSpells.NpcSpellsEntries",
		"Inventories.Item.Tasks.TaskActivities.NpcType.NpcTypesTint",
		"Inventories.Item.Tasks.TaskActivities.NpcType.Spawnentries",
		"Inventories.Item.Tasks.TaskActivities.NpcType.Spawnentries.NpcType",
		"Inventories.Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup",
		"Inventories.Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2",
		"Inventories.Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries",
		"Inventories.Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup",
		"Inventories.Item.Tasks.Tasksets",
		"Inventories.Item.TradeskillRecipeEntries",
		"Inventories.Item.TradeskillRecipeEntries.TradeskillRecipe",
		"Inventories.Item.TributeLevels",
		"Keyrings",
		"Mail",
		"PlayerTitlesets",
		"QuestGlobals",
		"Timers",
		"Titles",
		"Traders",
		"ZoneFlags",
	}
}

func (CharacterDatum) Connection() string {
    return ""
}
