package models

import (
	"github.com/volatiletech/null/v8"
)

type Spawn2 struct {
	ID                     int          `json:"id" gorm:"Column:id"`
	SpawngroupID           int          `json:"spawngroup_id" gorm:"Column:spawngroupID"`
	Zone                   null.String  `json:"zone" gorm:"Column:zone"`
	Version                int16        `json:"version" gorm:"Column:version"`
	X                      float32      `json:"x" gorm:"Column:x"`
	Y                      float32      `json:"y" gorm:"Column:y"`
	Z                      float32      `json:"z" gorm:"Column:z"`
	Heading                float32      `json:"heading" gorm:"Column:heading"`
	Respawntime            int          `json:"respawntime" gorm:"Column:respawntime"`
	Variance               int          `json:"variance" gorm:"Column:variance"`
	Pathgrid               int          `json:"pathgrid" gorm:"Column:pathgrid"`
	PathWhenZoneIdle       int8         `json:"path_when_zone_idle" gorm:"Column:path_when_zone_idle"`
	Condition              uint32       `json:"_condition" gorm:"Column:_condition"`
	CondValue              int32        `json:"cond_value" gorm:"Column:cond_value"`
	Enabled                uint8        `json:"enabled" gorm:"Column:enabled"`
	Animation              uint8        `json:"animation" gorm:"Column:animation"`
	MinExpansion           int8         `json:"min_expansion" gorm:"Column:min_expansion"`
	MaxExpansion           int8         `json:"max_expansion" gorm:"Column:max_expansion"`
	ContentFlags           null.String  `json:"content_flags" gorm:"Column:content_flags"`
	ContentFlagsDisabled   null.String  `json:"content_flags_disabled" gorm:"Column:content_flags_disabled"`
	Spawngroup             *Spawngroup  `json:"spawngroup,omitempty" gorm:"foreignKey:spawngroupID;references:id"`
	Spawnentries           []Spawnentry `json:"spawnentries,omitempty" gorm:"foreignKey:spawngroupID;references:spawngroupID"`
}

func (Spawn2) TableName() string {
    return "spawn2"
}

func (Spawn2) Relationships() []string {
    return []string{
		"Spawnentries",
		"Spawnentries.NpcType",
		"Spawnentries.NpcType.AlternateCurrency",
		"Spawnentries.NpcType.Loottable",
		"Spawnentries.NpcType.Loottable.LoottableEntries",
		"Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries",
		"Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item",
		"Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.AlternateCurrencies",
		"Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.CharacterCorpseItems",
		"Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.DiscoveredItems",
		"Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Doors",
		"Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Doors.Item",
		"Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Fishings",
		"Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Fishings.Item",
		"Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Fishings.NpcType",
		"Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Fishings.Zone",
		"Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Forages",
		"Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Forages.Item",
		"Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Forages.Zone",
		"Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.GroundSpawns",
		"Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.GroundSpawns.Zone",
		"Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.ItemTicks",
		"Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Keyrings",
		"Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.LootdropEntries",
		"Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Merchantlists",
		"Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType",
		"Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.ObjectContents",
		"Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Objects",
		"Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Objects.Item",
		"Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Objects.Zone",
		"Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.StartingItems",
		"Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.StartingItems.Item",
		"Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.StartingItems.Zone",
		"Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Tasks",
		"Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities",
		"Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.Goallists",
		"Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType",
		"Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Tasks.Tasksets",
		"Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.TradeskillRecipeEntries",
		"Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.TradeskillRecipeEntries.TradeskillRecipe",
		"Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.TributeLevels",
		"Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop",
		"Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LootdropEntries",
		"Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LoottableEntries",
		"Spawnentries.NpcType.Loottable.LoottableEntries.Loottable",
		"Spawnentries.NpcType.Loottable.NpcTypes",
		"Spawnentries.NpcType.Merchantlists",
		"Spawnentries.NpcType.Merchantlists.NpcType",
		"Spawnentries.NpcType.NpcEmotes",
		"Spawnentries.NpcType.NpcFactions",
		"Spawnentries.NpcType.NpcFactions.NpcFactionEntries",
		"Spawnentries.NpcType.NpcFactions.NpcFactionEntries.FactionList",
		"Spawnentries.NpcType.NpcSpell",
		"Spawnentries.NpcType.NpcSpell.NpcSpellsEntries",
		"Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew",
		"Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Aura",
		"Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew",
		"Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells",
		"Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes",
		"Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items",
		"Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.AlternateCurrencies",
		"Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.CharacterCorpseItems",
		"Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.DiscoveredItems",
		"Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Doors",
		"Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Doors.Item",
		"Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings",
		"Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.Item",
		"Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType",
		"Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.Zone",
		"Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Forages",
		"Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Forages.Item",
		"Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Forages.Zone",
		"Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.GroundSpawns",
		"Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.GroundSpawns.Zone",
		"Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.ItemTicks",
		"Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Keyrings",
		"Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries",
		"Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Item",
		"Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop",
		"Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LootdropEntries",
		"Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries",
		"Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.LootdropEntries",
		"Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable",
		"Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.LoottableEntries",
		"Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes",
		"Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists",
		"Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcType",
		"Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.ObjectContents",
		"Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Objects",
		"Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Objects.Item",
		"Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Objects.Zone",
		"Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.StartingItems",
		"Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.StartingItems.Item",
		"Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.StartingItems.Zone",
		"Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Tasks",
		"Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Tasks.TaskActivities",
		"Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Tasks.TaskActivities.Goallists",
		"Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Tasks.TaskActivities.NpcType",
		"Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Tasks.Tasksets",
		"Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.TradeskillRecipeEntries",
		"Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.TradeskillRecipeEntries.TradeskillRecipe",
		"Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items.TributeLevels",
		"Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries",
		"Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets",
		"Spawnentries.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals",
		"Spawnentries.NpcType.NpcTypesTint",
		"Spawnentries.NpcType.Spawnentries",
		"Spawnentries.Spawngroup",
		"Spawnentries.Spawngroup.Spawn2",
		"Spawngroup",
		"Spawngroup.Spawn2",
	}
}

func (Spawn2) Connection() string {
    return "eqemu_content"
}
