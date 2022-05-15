package models

import (
	"github.com/volatiletech/null/v8"
)

type Merchantlist struct {
	Merchantid             int         `json:"merchantid" gorm:"Column:merchantid"`
	Slot                   int         `json:"slot" gorm:"Column:slot"`
	Item                   int         `json:"item" gorm:"Column:item"`
	FactionRequired        int16       `json:"faction_required" gorm:"Column:faction_required"`
	LevelRequired          uint8       `json:"level_required" gorm:"Column:level_required"`
	AltCurrencyCost        uint16      `json:"alt_currency_cost" gorm:"Column:alt_currency_cost"`
	ClassesRequired        int         `json:"classes_required" gorm:"Column:classes_required"`
	Probability            int         `json:"probability" gorm:"Column:probability"`
	BucketName             string      `json:"bucket_name" gorm:"Column:bucket_name"`
	BucketValue            string      `json:"bucket_value" gorm:"Column:bucket_value"`
	BucketComparison       null.Uint8  `json:"bucket_comparison" gorm:"Column:bucket_comparison"`
	MinExpansion           int8        `json:"min_expansion" gorm:"Column:min_expansion"`
	MaxExpansion           int8        `json:"max_expansion" gorm:"Column:max_expansion"`
	ContentFlags           null.String `json:"content_flags" gorm:"Column:content_flags"`
	ContentFlagsDisabled   null.String `json:"content_flags_disabled" gorm:"Column:content_flags_disabled"`
	NpcType                *NpcType    `json:"npc_type,omitempty" gorm:"foreignKey:merchantid;references:merchant_id"`
}

func (Merchantlist) TableName() string {
    return "merchantlist"
}

func (Merchantlist) Relationships() []string {
    return []string{
		"NpcType",
		"NpcType.AlternateCurrency",
		"NpcType.Loottable",
		"NpcType.Loottable.LoottableEntries",
		"NpcType.Loottable.LoottableEntries.LootdropEntries",
		"NpcType.Loottable.LoottableEntries.LootdropEntries.Item",
		"NpcType.Loottable.LoottableEntries.LootdropEntries.Item.AlternateCurrencies",
		"NpcType.Loottable.LoottableEntries.LootdropEntries.Item.CharacterCorpseItems",
		"NpcType.Loottable.LoottableEntries.LootdropEntries.Item.DiscoveredItems",
		"NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Doors",
		"NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Doors.Item",
		"NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Fishings",
		"NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Fishings.Item",
		"NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Fishings.NpcType",
		"NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Fishings.Zone",
		"NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Forages",
		"NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Forages.Item",
		"NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Forages.Zone",
		"NpcType.Loottable.LoottableEntries.LootdropEntries.Item.GroundSpawns",
		"NpcType.Loottable.LoottableEntries.LootdropEntries.Item.GroundSpawns.Zone",
		"NpcType.Loottable.LoottableEntries.LootdropEntries.Item.ItemTicks",
		"NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Keyrings",
		"NpcType.Loottable.LoottableEntries.LootdropEntries.Item.LootdropEntries",
		"NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Merchantlists",
		"NpcType.Loottable.LoottableEntries.LootdropEntries.Item.ObjectContents",
		"NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Objects",
		"NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Objects.Item",
		"NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Objects.Zone",
		"NpcType.Loottable.LoottableEntries.LootdropEntries.Item.StartingItems",
		"NpcType.Loottable.LoottableEntries.LootdropEntries.Item.StartingItems.Item",
		"NpcType.Loottable.LoottableEntries.LootdropEntries.Item.StartingItems.Zone",
		"NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Tasks",
		"NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities",
		"NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.Goallists",
		"NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType",
		"NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Tasks.Tasksets",
		"NpcType.Loottable.LoottableEntries.LootdropEntries.Item.TradeskillRecipeEntries",
		"NpcType.Loottable.LoottableEntries.LootdropEntries.Item.TradeskillRecipeEntries.TradeskillRecipe",
		"NpcType.Loottable.LoottableEntries.LootdropEntries.Item.TributeLevels",
		"NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop",
		"NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LootdropEntries",
		"NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LoottableEntries",
		"NpcType.Loottable.LoottableEntries.Loottable",
		"NpcType.Loottable.NpcTypes",
		"NpcType.Merchantlists",
		"NpcType.NpcEmotes",
		"NpcType.NpcFactions",
		"NpcType.NpcFactions.NpcFactionEntries",
		"NpcType.NpcSpells",
		"NpcType.NpcSpells.NpcSpellsEntries",
		"NpcType.NpcTypesTint",
		"NpcType.Spawnentries",
		"NpcType.Spawnentries.NpcType",
		"NpcType.Spawnentries.Spawngroup",
		"NpcType.Spawnentries.Spawngroup.Spawn2",
		"NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries",
		"NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup",
	}
}

func (Merchantlist) Connection() string {
    return "eqemu_content"
}
