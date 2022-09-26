package models

import (
	"github.com/volatiletech/null/v8"
)

type Task struct {
	ID                    uint               `json:"id" gorm:"Column:id"`
	Type                  int8               `json:"type" gorm:"Column:type"`
	Duration              uint               `json:"duration" gorm:"Column:duration"`
	DurationCode          int8               `json:"duration_code" gorm:"Column:duration_code"`
	Title                 string             `json:"title" gorm:"Column:title"`
	Description           string             `json:"description" gorm:"Column:description"`
	RewardText            string             `json:"reward_text" gorm:"Column:reward_text"`
	RewardIdList          null.String        `json:"reward_id_list" gorm:"Column:reward_id_list"`
	CashReward            uint               `json:"cash_reward" gorm:"Column:cash_reward"`
	ExpReward             int                `json:"exp_reward" gorm:"Column:exp_reward"`
	RewardMethod          uint8              `json:"reward_method" gorm:"Column:reward_method"`
	RewardPoints          int                `json:"reward_points" gorm:"Column:reward_points"`
	RewardPointType       int                `json:"reward_point_type" gorm:"Column:reward_point_type"`
	MinLevel              uint8              `json:"min_level" gorm:"Column:min_level"`
	MaxLevel              uint8              `json:"max_level" gorm:"Column:max_level"`
	LevelSpread           uint               `json:"level_spread" gorm:"Column:level_spread"`
	MinPlayers            uint               `json:"min_players" gorm:"Column:min_players"`
	MaxPlayers            uint               `json:"max_players" gorm:"Column:max_players"`
	Repeatable            uint8              `json:"repeatable" gorm:"Column:repeatable"`
	FactionReward         int                `json:"faction_reward" gorm:"Column:faction_reward"`
	CompletionEmote       string             `json:"completion_emote" gorm:"Column:completion_emote"`
	ReplayTimerGroup      uint               `json:"replay_timer_group" gorm:"Column:replay_timer_group"`
	ReplayTimerSeconds    uint               `json:"replay_timer_seconds" gorm:"Column:replay_timer_seconds"`
	RequestTimerGroup     uint               `json:"request_timer_group" gorm:"Column:request_timer_group"`
	RequestTimerSeconds   uint               `json:"request_timer_seconds" gorm:"Column:request_timer_seconds"`
	DzTemplateId          uint               `json:"dz_template_id" gorm:"Column:dz_template_id"`
	LockActivityId        int                `json:"lock_activity_id" gorm:"Column:lock_activity_id"`
	FactionAmount         int                `json:"faction_amount" gorm:"Column:faction_amount"`
	TaskActivities        []TaskActivity     `json:"task_activities,omitempty" gorm:"foreignKey:taskid;references:id"`
	Tasksets              []Taskset          `json:"tasksets,omitempty" gorm:"foreignKey:taskid;references:id"`
	AlternateCurrency     *AlternateCurrency `json:"alternate_currency,omitempty" gorm:"foreignKey:reward_point_type;references:id"`
}

func (Task) TableName() string {
    return "tasks"
}

func (Task) Relationships() []string {
    return []string{
		"AlternateCurrency",
		"AlternateCurrency.Item",
		"AlternateCurrency.Item.AlternateCurrencies",
		"AlternateCurrency.Item.CharacterCorpseItems",
		"AlternateCurrency.Item.DiscoveredItems",
		"AlternateCurrency.Item.Doors",
		"AlternateCurrency.Item.Doors.Item",
		"AlternateCurrency.Item.Fishings",
		"AlternateCurrency.Item.Fishings.Item",
		"AlternateCurrency.Item.Fishings.NpcType",
		"AlternateCurrency.Item.Fishings.NpcType.AlternateCurrency",
		"AlternateCurrency.Item.Fishings.NpcType.Loottable",
		"AlternateCurrency.Item.Fishings.NpcType.Loottable.LoottableEntries",
		"AlternateCurrency.Item.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop",
		"AlternateCurrency.Item.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries",
		"AlternateCurrency.Item.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item",
		"AlternateCurrency.Item.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Lootdrop",
		"AlternateCurrency.Item.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LoottableEntries",
		"AlternateCurrency.Item.Fishings.NpcType.Loottable.LoottableEntries.Loottable",
		"AlternateCurrency.Item.Fishings.NpcType.Loottable.NpcTypes",
		"AlternateCurrency.Item.Fishings.NpcType.Merchantlists",
		"AlternateCurrency.Item.Fishings.NpcType.Merchantlists.Items",
		"AlternateCurrency.Item.Fishings.NpcType.Merchantlists.NpcTypes",
		"AlternateCurrency.Item.Fishings.NpcType.NpcEmotes",
		"AlternateCurrency.Item.Fishings.NpcType.NpcFactions",
		"AlternateCurrency.Item.Fishings.NpcType.NpcFactions.NpcFactionEntries",
		"AlternateCurrency.Item.Fishings.NpcType.NpcFactions.NpcFactionEntries.FactionList",
		"AlternateCurrency.Item.Fishings.NpcType.NpcSpell",
		"AlternateCurrency.Item.Fishings.NpcType.NpcSpell.NpcSpell",
		"AlternateCurrency.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries",
		"AlternateCurrency.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew",
		"AlternateCurrency.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Aura",
		"AlternateCurrency.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew",
		"AlternateCurrency.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells",
		"AlternateCurrency.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes",
		"AlternateCurrency.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items",
		"AlternateCurrency.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries",
		"AlternateCurrency.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets",
		"AlternateCurrency.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals",
		"AlternateCurrency.Item.Fishings.NpcType.NpcTypesTint",
		"AlternateCurrency.Item.Fishings.NpcType.Spawnentries",
		"AlternateCurrency.Item.Fishings.NpcType.Spawnentries.NpcType",
		"AlternateCurrency.Item.Fishings.NpcType.Spawnentries.Spawngroup",
		"AlternateCurrency.Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2",
		"AlternateCurrency.Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries",
		"AlternateCurrency.Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup",
		"AlternateCurrency.Item.Fishings.Zone",
		"AlternateCurrency.Item.Forages",
		"AlternateCurrency.Item.Forages.Item",
		"AlternateCurrency.Item.Forages.Zone",
		"AlternateCurrency.Item.GroundSpawns",
		"AlternateCurrency.Item.GroundSpawns.Zone",
		"AlternateCurrency.Item.ItemTicks",
		"AlternateCurrency.Item.Keyrings",
		"AlternateCurrency.Item.LootdropEntries",
		"AlternateCurrency.Item.LootdropEntries.Item",
		"AlternateCurrency.Item.LootdropEntries.Lootdrop",
		"AlternateCurrency.Item.LootdropEntries.Lootdrop.LootdropEntries",
		"AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries",
		"AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Lootdrop",
		"AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable",
		"AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.LoottableEntries",
		"AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes",
		"AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency",
		"AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Loottable",
		"AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists",
		"AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items",
		"AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.NpcTypes",
		"AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcEmotes",
		"AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions",
		"AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions.NpcFactionEntries",
		"AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions.NpcFactionEntries.FactionList",
		"AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell",
		"AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpell",
		"AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries",
		"AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew",
		"AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura",
		"AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew",
		"AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells",
		"AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes",
		"AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items",
		"AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries",
		"AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets",
		"AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals",
		"AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcTypesTint",
		"AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries",
		"AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.NpcType",
		"AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup",
		"AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2",
		"AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries",
		"AlternateCurrency.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup",
		"AlternateCurrency.Item.Merchantlists",
		"AlternateCurrency.Item.Merchantlists.Items",
		"AlternateCurrency.Item.Merchantlists.NpcTypes",
		"AlternateCurrency.Item.Merchantlists.NpcTypes.AlternateCurrency",
		"AlternateCurrency.Item.Merchantlists.NpcTypes.Loottable",
		"AlternateCurrency.Item.Merchantlists.NpcTypes.Loottable.LoottableEntries",
		"AlternateCurrency.Item.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop",
		"AlternateCurrency.Item.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries",
		"AlternateCurrency.Item.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item",
		"AlternateCurrency.Item.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Lootdrop",
		"AlternateCurrency.Item.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LoottableEntries",
		"AlternateCurrency.Item.Merchantlists.NpcTypes.Loottable.LoottableEntries.Loottable",
		"AlternateCurrency.Item.Merchantlists.NpcTypes.Loottable.NpcTypes",
		"AlternateCurrency.Item.Merchantlists.NpcTypes.Merchantlists",
		"AlternateCurrency.Item.Merchantlists.NpcTypes.NpcEmotes",
		"AlternateCurrency.Item.Merchantlists.NpcTypes.NpcFactions",
		"AlternateCurrency.Item.Merchantlists.NpcTypes.NpcFactions.NpcFactionEntries",
		"AlternateCurrency.Item.Merchantlists.NpcTypes.NpcFactions.NpcFactionEntries.FactionList",
		"AlternateCurrency.Item.Merchantlists.NpcTypes.NpcSpell",
		"AlternateCurrency.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpell",
		"AlternateCurrency.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries",
		"AlternateCurrency.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew",
		"AlternateCurrency.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura",
		"AlternateCurrency.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew",
		"AlternateCurrency.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells",
		"AlternateCurrency.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes",
		"AlternateCurrency.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items",
		"AlternateCurrency.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries",
		"AlternateCurrency.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets",
		"AlternateCurrency.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals",
		"AlternateCurrency.Item.Merchantlists.NpcTypes.NpcTypesTint",
		"AlternateCurrency.Item.Merchantlists.NpcTypes.Spawnentries",
		"AlternateCurrency.Item.Merchantlists.NpcTypes.Spawnentries.NpcType",
		"AlternateCurrency.Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup",
		"AlternateCurrency.Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2",
		"AlternateCurrency.Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries",
		"AlternateCurrency.Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup",
		"AlternateCurrency.Item.ObjectContents",
		"AlternateCurrency.Item.Objects",
		"AlternateCurrency.Item.Objects.Item",
		"AlternateCurrency.Item.Objects.Zone",
		"AlternateCurrency.Item.StartingItems",
		"AlternateCurrency.Item.StartingItems.Item",
		"AlternateCurrency.Item.StartingItems.Zone",
		"AlternateCurrency.Item.TradeskillRecipeEntries",
		"AlternateCurrency.Item.TradeskillRecipeEntries.TradeskillRecipe",
		"AlternateCurrency.Item.TributeLevels",
		"TaskActivities",
		"Tasksets",
	}
}

func (Task) Connection() string {
    return "eqemu_content"
}
