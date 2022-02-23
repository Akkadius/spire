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
	MinExpansion           uint8        `json:"min_expansion" gorm:"Column:min_expansion"`
	MaxExpansion           uint8        `json:"max_expansion" gorm:"Column:max_expansion"`
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
		"Spawnentries.NpcType.Merchantlists",
		"Spawnentries.NpcType.Merchantlists.Item",
		"Spawnentries.NpcType.Merchantlists.Item.AlternateCurrencies",
		"Spawnentries.NpcType.Merchantlists.Item.CharacterCorpseItems",
		"Spawnentries.NpcType.Merchantlists.Item.DiscoveredItems",
		"Spawnentries.NpcType.Merchantlists.Item.Doors",
		"Spawnentries.NpcType.Merchantlists.Item.Doors.Item",
		"Spawnentries.NpcType.Merchantlists.Item.Doors.Zone",
		"Spawnentries.NpcType.Merchantlists.Item.Fishings",
		"Spawnentries.NpcType.Merchantlists.Item.Fishings.Item",
		"Spawnentries.NpcType.Merchantlists.Item.Fishings.NpcType",
		"Spawnentries.NpcType.Merchantlists.Item.Fishings.Zone",
		"Spawnentries.NpcType.Merchantlists.Item.Forages",
		"Spawnentries.NpcType.Merchantlists.Item.Forages.Item",
		"Spawnentries.NpcType.Merchantlists.Item.Forages.Zone",
		"Spawnentries.NpcType.Merchantlists.Item.GroundSpawns",
		"Spawnentries.NpcType.Merchantlists.Item.GroundSpawns.Item",
		"Spawnentries.NpcType.Merchantlists.Item.GroundSpawns.Zone",
		"Spawnentries.NpcType.Merchantlists.Item.ItemTicks",
		"Spawnentries.NpcType.Merchantlists.Item.Keyrings",
		"Spawnentries.NpcType.Merchantlists.Item.LootdropEntries",
		"Spawnentries.NpcType.Merchantlists.Item.LootdropEntries.Item",
		"Spawnentries.NpcType.Merchantlists.Item.LootdropEntries.Lootdrop",
		"Spawnentries.NpcType.Merchantlists.Item.LootdropEntries.Lootdrop.LootdropEntries",
		"Spawnentries.NpcType.Merchantlists.Item.LootdropEntries.Lootdrop.LoottableEntries",
		"Spawnentries.NpcType.Merchantlists.Item.LootdropEntries.Lootdrop.LoottableEntries.LootdropEntries",
		"Spawnentries.NpcType.Merchantlists.Item.Merchantlists",
		"Spawnentries.NpcType.Merchantlists.Item.ObjectContents",
		"Spawnentries.NpcType.Merchantlists.Item.Objects",
		"Spawnentries.NpcType.Merchantlists.Item.Objects.Item",
		"Spawnentries.NpcType.Merchantlists.Item.Objects.Zone",
		"Spawnentries.NpcType.Merchantlists.Item.StartingItems",
		"Spawnentries.NpcType.Merchantlists.Item.StartingItems.Item",
		"Spawnentries.NpcType.Merchantlists.Item.StartingItems.Zone",
		"Spawnentries.NpcType.Merchantlists.Item.TradeskillRecipeEntries",
		"Spawnentries.NpcType.Merchantlists.Item.TradeskillRecipeEntries.TradeskillRecipe",
		"Spawnentries.NpcType.Merchantlists.Item.TributeLevels",
		"Spawnentries.NpcType.NpcEmotes",
		"Spawnentries.NpcType.NpcFactions",
		"Spawnentries.NpcType.NpcFactions.NpcFactionEntries",
		"Spawnentries.NpcType.NpcSpells",
		"Spawnentries.NpcType.NpcSpells.NpcSpellsEntries",
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
