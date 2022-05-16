package models

import (
	"github.com/volatiletech/null/v8"
)

type Lootdrop struct {
	ID                     uint             `json:"id" gorm:"Column:id"`
	Name                   string           `json:"name" gorm:"Column:name"`
	MinExpansion           int8             `json:"min_expansion" gorm:"Column:min_expansion"`
	MaxExpansion           int8             `json:"max_expansion" gorm:"Column:max_expansion"`
	ContentFlags           null.String      `json:"content_flags" gorm:"Column:content_flags"`
	ContentFlagsDisabled   null.String      `json:"content_flags_disabled" gorm:"Column:content_flags_disabled"`
	LootdropEntries        []LootdropEntry  `json:"lootdrop_entries,omitempty" gorm:"foreignKey:lootdrop_id;references:id"`
	LoottableEntries       []LoottableEntry `json:"loottable_entries,omitempty" gorm:"foreignKey:lootdrop_id;references:id"`
}

func (Lootdrop) TableName() string {
    return "lootdrop"
}

func (Lootdrop) Relationships() []string {
    return []string{
		"LootdropEntries",
		"LootdropEntries.Item",
		"LootdropEntries.Item.AlternateCurrencies",
		"LootdropEntries.Item.CharacterCorpseItems",
		"LootdropEntries.Item.DiscoveredItems",
		"LootdropEntries.Item.Doors",
		"LootdropEntries.Item.Doors.Item",
		"LootdropEntries.Item.Fishings",
		"LootdropEntries.Item.Fishings.Item",
		"LootdropEntries.Item.Fishings.NpcType",
		"LootdropEntries.Item.Fishings.NpcType.AlternateCurrency",
		"LootdropEntries.Item.Fishings.NpcType.Loottable",
		"LootdropEntries.Item.Fishings.NpcType.Loottable.LoottableEntries",
		"LootdropEntries.Item.Fishings.NpcType.Loottable.LoottableEntries.LootdropEntries",
		"LootdropEntries.Item.Fishings.NpcType.Loottable.LoottableEntries.Loottable",
		"LootdropEntries.Item.Fishings.NpcType.Loottable.NpcTypes",
		"LootdropEntries.Item.Fishings.NpcType.Merchantlists",
		"LootdropEntries.Item.Fishings.NpcType.Merchantlists.NpcType",
		"LootdropEntries.Item.Fishings.NpcType.NpcEmotes",
		"LootdropEntries.Item.Fishings.NpcType.NpcFactions",
		"LootdropEntries.Item.Fishings.NpcType.NpcFactions.NpcFactionEntries",
		"LootdropEntries.Item.Fishings.NpcType.NpcSpells",
		"LootdropEntries.Item.Fishings.NpcType.NpcSpells.NpcSpellsEntries",
		"LootdropEntries.Item.Fishings.NpcType.NpcSpells.NpcSpellsEntries.SpellsNew",
		"LootdropEntries.Item.Fishings.NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Aura",
		"LootdropEntries.Item.Fishings.NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Aura.SpellsNew",
		"LootdropEntries.Item.Fishings.NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.BlockedSpells",
		"LootdropEntries.Item.Fishings.NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Damageshieldtypes",
		"LootdropEntries.Item.Fishings.NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items",
		"LootdropEntries.Item.Fishings.NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.NpcSpellsEntries",
		"LootdropEntries.Item.Fishings.NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.SpellBuckets",
		"LootdropEntries.Item.Fishings.NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.SpellGlobals",
		"LootdropEntries.Item.Fishings.NpcType.NpcTypesTint",
		"LootdropEntries.Item.Fishings.NpcType.Spawnentries",
		"LootdropEntries.Item.Fishings.NpcType.Spawnentries.NpcType",
		"LootdropEntries.Item.Fishings.NpcType.Spawnentries.Spawngroup",
		"LootdropEntries.Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2",
		"LootdropEntries.Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries",
		"LootdropEntries.Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup",
		"LootdropEntries.Item.Fishings.Zone",
		"LootdropEntries.Item.Forages",
		"LootdropEntries.Item.Forages.Item",
		"LootdropEntries.Item.Forages.Zone",
		"LootdropEntries.Item.GroundSpawns",
		"LootdropEntries.Item.GroundSpawns.Zone",
		"LootdropEntries.Item.ItemTicks",
		"LootdropEntries.Item.Keyrings",
		"LootdropEntries.Item.LootdropEntries",
		"LootdropEntries.Item.Merchantlists",
		"LootdropEntries.Item.Merchantlists.NpcType",
		"LootdropEntries.Item.Merchantlists.NpcType.AlternateCurrency",
		"LootdropEntries.Item.Merchantlists.NpcType.Loottable",
		"LootdropEntries.Item.Merchantlists.NpcType.Loottable.LoottableEntries",
		"LootdropEntries.Item.Merchantlists.NpcType.Loottable.LoottableEntries.LootdropEntries",
		"LootdropEntries.Item.Merchantlists.NpcType.Loottable.LoottableEntries.Loottable",
		"LootdropEntries.Item.Merchantlists.NpcType.Loottable.NpcTypes",
		"LootdropEntries.Item.Merchantlists.NpcType.Merchantlists",
		"LootdropEntries.Item.Merchantlists.NpcType.NpcEmotes",
		"LootdropEntries.Item.Merchantlists.NpcType.NpcFactions",
		"LootdropEntries.Item.Merchantlists.NpcType.NpcFactions.NpcFactionEntries",
		"LootdropEntries.Item.Merchantlists.NpcType.NpcSpells",
		"LootdropEntries.Item.Merchantlists.NpcType.NpcSpells.NpcSpellsEntries",
		"LootdropEntries.Item.Merchantlists.NpcType.NpcSpells.NpcSpellsEntries.SpellsNew",
		"LootdropEntries.Item.Merchantlists.NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Aura",
		"LootdropEntries.Item.Merchantlists.NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Aura.SpellsNew",
		"LootdropEntries.Item.Merchantlists.NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.BlockedSpells",
		"LootdropEntries.Item.Merchantlists.NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Damageshieldtypes",
		"LootdropEntries.Item.Merchantlists.NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items",
		"LootdropEntries.Item.Merchantlists.NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.NpcSpellsEntries",
		"LootdropEntries.Item.Merchantlists.NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.SpellBuckets",
		"LootdropEntries.Item.Merchantlists.NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.SpellGlobals",
		"LootdropEntries.Item.Merchantlists.NpcType.NpcTypesTint",
		"LootdropEntries.Item.Merchantlists.NpcType.Spawnentries",
		"LootdropEntries.Item.Merchantlists.NpcType.Spawnentries.NpcType",
		"LootdropEntries.Item.Merchantlists.NpcType.Spawnentries.Spawngroup",
		"LootdropEntries.Item.Merchantlists.NpcType.Spawnentries.Spawngroup.Spawn2",
		"LootdropEntries.Item.Merchantlists.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries",
		"LootdropEntries.Item.Merchantlists.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup",
		"LootdropEntries.Item.ObjectContents",
		"LootdropEntries.Item.Objects",
		"LootdropEntries.Item.Objects.Item",
		"LootdropEntries.Item.Objects.Zone",
		"LootdropEntries.Item.StartingItems",
		"LootdropEntries.Item.StartingItems.Item",
		"LootdropEntries.Item.StartingItems.Zone",
		"LootdropEntries.Item.Tasks",
		"LootdropEntries.Item.Tasks.TaskActivities",
		"LootdropEntries.Item.Tasks.TaskActivities.Goallists",
		"LootdropEntries.Item.Tasks.TaskActivities.NpcType",
		"LootdropEntries.Item.Tasks.TaskActivities.NpcType.AlternateCurrency",
		"LootdropEntries.Item.Tasks.TaskActivities.NpcType.Loottable",
		"LootdropEntries.Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries",
		"LootdropEntries.Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries",
		"LootdropEntries.Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.Loottable",
		"LootdropEntries.Item.Tasks.TaskActivities.NpcType.Loottable.NpcTypes",
		"LootdropEntries.Item.Tasks.TaskActivities.NpcType.Merchantlists",
		"LootdropEntries.Item.Tasks.TaskActivities.NpcType.Merchantlists.NpcType",
		"LootdropEntries.Item.Tasks.TaskActivities.NpcType.NpcEmotes",
		"LootdropEntries.Item.Tasks.TaskActivities.NpcType.NpcFactions",
		"LootdropEntries.Item.Tasks.TaskActivities.NpcType.NpcFactions.NpcFactionEntries",
		"LootdropEntries.Item.Tasks.TaskActivities.NpcType.NpcSpells",
		"LootdropEntries.Item.Tasks.TaskActivities.NpcType.NpcSpells.NpcSpellsEntries",
		"LootdropEntries.Item.Tasks.TaskActivities.NpcType.NpcSpells.NpcSpellsEntries.SpellsNew",
		"LootdropEntries.Item.Tasks.TaskActivities.NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Aura",
		"LootdropEntries.Item.Tasks.TaskActivities.NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Aura.SpellsNew",
		"LootdropEntries.Item.Tasks.TaskActivities.NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.BlockedSpells",
		"LootdropEntries.Item.Tasks.TaskActivities.NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Damageshieldtypes",
		"LootdropEntries.Item.Tasks.TaskActivities.NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items",
		"LootdropEntries.Item.Tasks.TaskActivities.NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.NpcSpellsEntries",
		"LootdropEntries.Item.Tasks.TaskActivities.NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.SpellBuckets",
		"LootdropEntries.Item.Tasks.TaskActivities.NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.SpellGlobals",
		"LootdropEntries.Item.Tasks.TaskActivities.NpcType.NpcTypesTint",
		"LootdropEntries.Item.Tasks.TaskActivities.NpcType.Spawnentries",
		"LootdropEntries.Item.Tasks.TaskActivities.NpcType.Spawnentries.NpcType",
		"LootdropEntries.Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup",
		"LootdropEntries.Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2",
		"LootdropEntries.Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries",
		"LootdropEntries.Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup",
		"LootdropEntries.Item.Tasks.Tasksets",
		"LootdropEntries.Item.TradeskillRecipeEntries",
		"LootdropEntries.Item.TradeskillRecipeEntries.TradeskillRecipe",
		"LootdropEntries.Item.TributeLevels",
		"LootdropEntries.Lootdrop",
		"LoottableEntries",
		"LoottableEntries.LootdropEntries",
		"LoottableEntries.LootdropEntries.Item",
		"LoottableEntries.LootdropEntries.Item.AlternateCurrencies",
		"LoottableEntries.LootdropEntries.Item.CharacterCorpseItems",
		"LoottableEntries.LootdropEntries.Item.DiscoveredItems",
		"LoottableEntries.LootdropEntries.Item.Doors",
		"LoottableEntries.LootdropEntries.Item.Doors.Item",
		"LoottableEntries.LootdropEntries.Item.Fishings",
		"LoottableEntries.LootdropEntries.Item.Fishings.Item",
		"LoottableEntries.LootdropEntries.Item.Fishings.NpcType",
		"LoottableEntries.LootdropEntries.Item.Fishings.NpcType.AlternateCurrency",
		"LoottableEntries.LootdropEntries.Item.Fishings.NpcType.Loottable",
		"LoottableEntries.LootdropEntries.Item.Fishings.NpcType.Loottable.LoottableEntries",
		"LoottableEntries.LootdropEntries.Item.Fishings.NpcType.Loottable.NpcTypes",
		"LoottableEntries.LootdropEntries.Item.Fishings.NpcType.Merchantlists",
		"LoottableEntries.LootdropEntries.Item.Fishings.NpcType.Merchantlists.NpcType",
		"LoottableEntries.LootdropEntries.Item.Fishings.NpcType.NpcEmotes",
		"LoottableEntries.LootdropEntries.Item.Fishings.NpcType.NpcFactions",
		"LoottableEntries.LootdropEntries.Item.Fishings.NpcType.NpcFactions.NpcFactionEntries",
		"LoottableEntries.LootdropEntries.Item.Fishings.NpcType.NpcSpells",
		"LoottableEntries.LootdropEntries.Item.Fishings.NpcType.NpcSpells.NpcSpellsEntries",
		"LoottableEntries.LootdropEntries.Item.Fishings.NpcType.NpcSpells.NpcSpellsEntries.SpellsNew",
		"LoottableEntries.LootdropEntries.Item.Fishings.NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Aura",
		"LoottableEntries.LootdropEntries.Item.Fishings.NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Aura.SpellsNew",
		"LoottableEntries.LootdropEntries.Item.Fishings.NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.BlockedSpells",
		"LoottableEntries.LootdropEntries.Item.Fishings.NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Damageshieldtypes",
		"LoottableEntries.LootdropEntries.Item.Fishings.NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items",
		"LoottableEntries.LootdropEntries.Item.Fishings.NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.NpcSpellsEntries",
		"LoottableEntries.LootdropEntries.Item.Fishings.NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.SpellBuckets",
		"LoottableEntries.LootdropEntries.Item.Fishings.NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.SpellGlobals",
		"LoottableEntries.LootdropEntries.Item.Fishings.NpcType.NpcTypesTint",
		"LoottableEntries.LootdropEntries.Item.Fishings.NpcType.Spawnentries",
		"LoottableEntries.LootdropEntries.Item.Fishings.NpcType.Spawnentries.NpcType",
		"LoottableEntries.LootdropEntries.Item.Fishings.NpcType.Spawnentries.Spawngroup",
		"LoottableEntries.LootdropEntries.Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2",
		"LoottableEntries.LootdropEntries.Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries",
		"LoottableEntries.LootdropEntries.Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup",
		"LoottableEntries.LootdropEntries.Item.Fishings.Zone",
		"LoottableEntries.LootdropEntries.Item.Forages",
		"LoottableEntries.LootdropEntries.Item.Forages.Item",
		"LoottableEntries.LootdropEntries.Item.Forages.Zone",
		"LoottableEntries.LootdropEntries.Item.GroundSpawns",
		"LoottableEntries.LootdropEntries.Item.GroundSpawns.Zone",
		"LoottableEntries.LootdropEntries.Item.ItemTicks",
		"LoottableEntries.LootdropEntries.Item.Keyrings",
		"LoottableEntries.LootdropEntries.Item.LootdropEntries",
		"LoottableEntries.LootdropEntries.Item.Merchantlists",
		"LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType",
		"LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.AlternateCurrency",
		"LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.Loottable",
		"LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.Loottable.LoottableEntries",
		"LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.Loottable.NpcTypes",
		"LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.Merchantlists",
		"LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.NpcEmotes",
		"LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.NpcFactions",
		"LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.NpcFactions.NpcFactionEntries",
		"LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.NpcSpells",
		"LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.NpcSpells.NpcSpellsEntries",
		"LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.NpcSpells.NpcSpellsEntries.SpellsNew",
		"LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Aura",
		"LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Aura.SpellsNew",
		"LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.BlockedSpells",
		"LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Damageshieldtypes",
		"LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items",
		"LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.NpcSpellsEntries",
		"LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.SpellBuckets",
		"LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.SpellGlobals",
		"LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.NpcTypesTint",
		"LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.Spawnentries",
		"LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.Spawnentries.NpcType",
		"LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.Spawnentries.Spawngroup",
		"LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.Spawnentries.Spawngroup.Spawn2",
		"LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries",
		"LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup",
		"LoottableEntries.LootdropEntries.Item.ObjectContents",
		"LoottableEntries.LootdropEntries.Item.Objects",
		"LoottableEntries.LootdropEntries.Item.Objects.Item",
		"LoottableEntries.LootdropEntries.Item.Objects.Zone",
		"LoottableEntries.LootdropEntries.Item.StartingItems",
		"LoottableEntries.LootdropEntries.Item.StartingItems.Item",
		"LoottableEntries.LootdropEntries.Item.StartingItems.Zone",
		"LoottableEntries.LootdropEntries.Item.Tasks",
		"LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities",
		"LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.Goallists",
		"LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType",
		"LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.AlternateCurrency",
		"LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.Loottable",
		"LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries",
		"LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.Loottable.NpcTypes",
		"LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.Merchantlists",
		"LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.Merchantlists.NpcType",
		"LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.NpcEmotes",
		"LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.NpcFactions",
		"LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.NpcFactions.NpcFactionEntries",
		"LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.NpcSpells",
		"LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.NpcSpells.NpcSpellsEntries",
		"LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.NpcSpells.NpcSpellsEntries.SpellsNew",
		"LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Aura",
		"LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Aura.SpellsNew",
		"LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.BlockedSpells",
		"LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Damageshieldtypes",
		"LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items",
		"LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.NpcSpellsEntries",
		"LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.SpellBuckets",
		"LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.SpellGlobals",
		"LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.NpcTypesTint",
		"LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.Spawnentries",
		"LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.Spawnentries.NpcType",
		"LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup",
		"LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2",
		"LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries",
		"LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup",
		"LoottableEntries.LootdropEntries.Item.Tasks.Tasksets",
		"LoottableEntries.LootdropEntries.Item.TradeskillRecipeEntries",
		"LoottableEntries.LootdropEntries.Item.TradeskillRecipeEntries.TradeskillRecipe",
		"LoottableEntries.LootdropEntries.Item.TributeLevels",
		"LoottableEntries.LootdropEntries.Lootdrop",
		"LoottableEntries.Loottable",
		"LoottableEntries.Loottable.LoottableEntries",
		"LoottableEntries.Loottable.NpcTypes",
		"LoottableEntries.Loottable.NpcTypes.AlternateCurrency",
		"LoottableEntries.Loottable.NpcTypes.Loottable",
		"LoottableEntries.Loottable.NpcTypes.Merchantlists",
		"LoottableEntries.Loottable.NpcTypes.Merchantlists.NpcType",
		"LoottableEntries.Loottable.NpcTypes.NpcEmotes",
		"LoottableEntries.Loottable.NpcTypes.NpcFactions",
		"LoottableEntries.Loottable.NpcTypes.NpcFactions.NpcFactionEntries",
		"LoottableEntries.Loottable.NpcTypes.NpcSpells",
		"LoottableEntries.Loottable.NpcTypes.NpcSpells.NpcSpellsEntries",
		"LoottableEntries.Loottable.NpcTypes.NpcSpells.NpcSpellsEntries.SpellsNew",
		"LoottableEntries.Loottable.NpcTypes.NpcSpells.NpcSpellsEntries.SpellsNew.Aura",
		"LoottableEntries.Loottable.NpcTypes.NpcSpells.NpcSpellsEntries.SpellsNew.Aura.SpellsNew",
		"LoottableEntries.Loottable.NpcTypes.NpcSpells.NpcSpellsEntries.SpellsNew.BlockedSpells",
		"LoottableEntries.Loottable.NpcTypes.NpcSpells.NpcSpellsEntries.SpellsNew.Damageshieldtypes",
		"LoottableEntries.Loottable.NpcTypes.NpcSpells.NpcSpellsEntries.SpellsNew.Items",
		"LoottableEntries.Loottable.NpcTypes.NpcSpells.NpcSpellsEntries.SpellsNew.Items.AlternateCurrencies",
		"LoottableEntries.Loottable.NpcTypes.NpcSpells.NpcSpellsEntries.SpellsNew.Items.CharacterCorpseItems",
		"LoottableEntries.Loottable.NpcTypes.NpcSpells.NpcSpellsEntries.SpellsNew.Items.DiscoveredItems",
		"LoottableEntries.Loottable.NpcTypes.NpcSpells.NpcSpellsEntries.SpellsNew.Items.Doors",
		"LoottableEntries.Loottable.NpcTypes.NpcSpells.NpcSpellsEntries.SpellsNew.Items.Doors.Item",
		"LoottableEntries.Loottable.NpcTypes.NpcSpells.NpcSpellsEntries.SpellsNew.Items.Fishings",
		"LoottableEntries.Loottable.NpcTypes.NpcSpells.NpcSpellsEntries.SpellsNew.Items.Fishings.Item",
		"LoottableEntries.Loottable.NpcTypes.NpcSpells.NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType",
		"LoottableEntries.Loottable.NpcTypes.NpcSpells.NpcSpellsEntries.SpellsNew.Items.Fishings.Zone",
		"LoottableEntries.Loottable.NpcTypes.NpcSpells.NpcSpellsEntries.SpellsNew.Items.Forages",
		"LoottableEntries.Loottable.NpcTypes.NpcSpells.NpcSpellsEntries.SpellsNew.Items.Forages.Item",
		"LoottableEntries.Loottable.NpcTypes.NpcSpells.NpcSpellsEntries.SpellsNew.Items.Forages.Zone",
		"LoottableEntries.Loottable.NpcTypes.NpcSpells.NpcSpellsEntries.SpellsNew.Items.GroundSpawns",
		"LoottableEntries.Loottable.NpcTypes.NpcSpells.NpcSpellsEntries.SpellsNew.Items.GroundSpawns.Zone",
		"LoottableEntries.Loottable.NpcTypes.NpcSpells.NpcSpellsEntries.SpellsNew.Items.ItemTicks",
		"LoottableEntries.Loottable.NpcTypes.NpcSpells.NpcSpellsEntries.SpellsNew.Items.Keyrings",
		"LoottableEntries.Loottable.NpcTypes.NpcSpells.NpcSpellsEntries.SpellsNew.Items.LootdropEntries",
		"LoottableEntries.Loottable.NpcTypes.NpcSpells.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Item",
		"LoottableEntries.Loottable.NpcTypes.NpcSpells.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop",
		"LoottableEntries.Loottable.NpcTypes.NpcSpells.NpcSpellsEntries.SpellsNew.Items.Merchantlists",
		"LoottableEntries.Loottable.NpcTypes.NpcSpells.NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcType",
		"LoottableEntries.Loottable.NpcTypes.NpcSpells.NpcSpellsEntries.SpellsNew.Items.ObjectContents",
		"LoottableEntries.Loottable.NpcTypes.NpcSpells.NpcSpellsEntries.SpellsNew.Items.Objects",
		"LoottableEntries.Loottable.NpcTypes.NpcSpells.NpcSpellsEntries.SpellsNew.Items.Objects.Item",
		"LoottableEntries.Loottable.NpcTypes.NpcSpells.NpcSpellsEntries.SpellsNew.Items.Objects.Zone",
		"LoottableEntries.Loottable.NpcTypes.NpcSpells.NpcSpellsEntries.SpellsNew.Items.StartingItems",
		"LoottableEntries.Loottable.NpcTypes.NpcSpells.NpcSpellsEntries.SpellsNew.Items.StartingItems.Item",
		"LoottableEntries.Loottable.NpcTypes.NpcSpells.NpcSpellsEntries.SpellsNew.Items.StartingItems.Zone",
		"LoottableEntries.Loottable.NpcTypes.NpcSpells.NpcSpellsEntries.SpellsNew.Items.Tasks",
		"LoottableEntries.Loottable.NpcTypes.NpcSpells.NpcSpellsEntries.SpellsNew.Items.Tasks.TaskActivities",
		"LoottableEntries.Loottable.NpcTypes.NpcSpells.NpcSpellsEntries.SpellsNew.Items.Tasks.TaskActivities.Goallists",
		"LoottableEntries.Loottable.NpcTypes.NpcSpells.NpcSpellsEntries.SpellsNew.Items.Tasks.TaskActivities.NpcType",
		"LoottableEntries.Loottable.NpcTypes.NpcSpells.NpcSpellsEntries.SpellsNew.Items.Tasks.Tasksets",
		"LoottableEntries.Loottable.NpcTypes.NpcSpells.NpcSpellsEntries.SpellsNew.Items.TradeskillRecipeEntries",
		"LoottableEntries.Loottable.NpcTypes.NpcSpells.NpcSpellsEntries.SpellsNew.Items.TradeskillRecipeEntries.TradeskillRecipe",
		"LoottableEntries.Loottable.NpcTypes.NpcSpells.NpcSpellsEntries.SpellsNew.Items.TributeLevels",
		"LoottableEntries.Loottable.NpcTypes.NpcSpells.NpcSpellsEntries.SpellsNew.NpcSpellsEntries",
		"LoottableEntries.Loottable.NpcTypes.NpcSpells.NpcSpellsEntries.SpellsNew.SpellBuckets",
		"LoottableEntries.Loottable.NpcTypes.NpcSpells.NpcSpellsEntries.SpellsNew.SpellGlobals",
		"LoottableEntries.Loottable.NpcTypes.NpcTypesTint",
		"LoottableEntries.Loottable.NpcTypes.Spawnentries",
		"LoottableEntries.Loottable.NpcTypes.Spawnentries.NpcType",
		"LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup",
		"LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2",
		"LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries",
		"LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup",
	}
}

func (Lootdrop) Connection() string {
    return "eqemu_content"
}
