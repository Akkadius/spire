package models

import (
	"github.com/volatiletech/null/v8"
)

type Spawnentry struct {
	SpawngroupID           int         `json:"spawngroup_id" gorm:"Column:spawngroupID"`
	NpcID                  int         `json:"npc_id" gorm:"Column:npcID"`
	Chance                 int16       `json:"chance" gorm:"Column:chance"`
	ConditionValueFilter   int32       `json:"condition_value_filter" gorm:"Column:condition_value_filter"`
	MinExpansion           int8        `json:"min_expansion" gorm:"Column:min_expansion"`
	MaxExpansion           int8        `json:"max_expansion" gorm:"Column:max_expansion"`
	ContentFlags           null.String `json:"content_flags" gorm:"Column:content_flags"`
	ContentFlagsDisabled   null.String `json:"content_flags_disabled" gorm:"Column:content_flags_disabled"`
	Spawngroup             *Spawngroup `json:"spawngroup,omitempty" gorm:"foreignKey:spawngroupID;references:id"`
	NpcType                *NpcType    `json:"npc_type,omitempty" gorm:"foreignKey:npcID;references:id"`
}

func (Spawnentry) TableName() string {
    return "spawnentry"
}

func (Spawnentry) Relationships() []string {
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
		"NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType",
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
		"NpcType.Merchantlists.NpcType",
		"NpcType.NpcEmotes",
		"NpcType.NpcFactions",
		"NpcType.NpcFactions.NpcFactionEntries",
		"NpcType.NpcFactions.NpcFactionEntries.FactionList",
		"NpcType.NpcSpell",
		"NpcType.NpcSpell.NpcSpellsEntries",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Aura",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.AlternateCurrencies",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.CharacterCorpseItems",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.DiscoveredItems",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Doors",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Doors.Item",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.Item",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.Zone",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Forages",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Forages.Item",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Forages.Zone",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.GroundSpawns",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.GroundSpawns.Zone",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.ItemTicks",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Keyrings",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Item",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LootdropEntries",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.LootdropEntries",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.LoottableEntries",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcType",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.ObjectContents",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Objects",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Objects.Item",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Objects.Zone",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.StartingItems",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.StartingItems.Item",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.StartingItems.Zone",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Tasks",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Tasks.TaskActivities",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Tasks.TaskActivities.Goallists",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Tasks.TaskActivities.NpcType",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Tasks.Tasksets",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.TradeskillRecipeEntries",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.TradeskillRecipeEntries.TradeskillRecipe",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.TributeLevels",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets",
		"NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals",
		"NpcType.NpcTypesTint",
		"NpcType.Spawnentries",
		"Spawngroup",
		"Spawngroup.Spawn2",
		"Spawngroup.Spawn2.Spawnentries",
		"Spawngroup.Spawn2.Spawngroup",
	}
}

func (Spawnentry) Connection() string {
    return "eqemu_content"
}
