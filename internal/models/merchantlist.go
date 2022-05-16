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
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Aura",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Aura.SpellsNew",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.BlockedSpells",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Damageshieldtypes",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.AlternateCurrencies",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.CharacterCorpseItems",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.DiscoveredItems",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.Doors",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.Doors.Item",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.Fishings",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.Fishings.Item",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.Fishings.Zone",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.Forages",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.Forages.Item",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.Forages.Zone",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.GroundSpawns",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.GroundSpawns.Zone",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.ItemTicks",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.Keyrings",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.LootdropEntries",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Item",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LootdropEntries",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.LootdropEntries",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.LoottableEntries",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.Merchantlists",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.ObjectContents",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.Objects",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.Objects.Item",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.Objects.Zone",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.StartingItems",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.StartingItems.Item",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.StartingItems.Zone",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.Tasks",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.Tasks.TaskActivities",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.Tasks.TaskActivities.Goallists",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.Tasks.TaskActivities.NpcType",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.Tasks.Tasksets",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.TradeskillRecipeEntries",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.TradeskillRecipeEntries.TradeskillRecipe",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.TributeLevels",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.NpcSpellsEntries",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.SpellBuckets",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.SpellGlobals",
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
