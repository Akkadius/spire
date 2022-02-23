package models

import (
	"github.com/volatiletech/null/v8"
)

type GlobalLoot struct {
	ID                     int         `json:"id" gorm:"Column:id"`
	Description            null.String `json:"description" gorm:"Column:description"`
	LoottableId            int         `json:"loottable_id" gorm:"Column:loottable_id"`
	Enabled                int8        `json:"enabled" gorm:"Column:enabled"`
	MinLevel               int         `json:"min_level" gorm:"Column:min_level"`
	MaxLevel               int         `json:"max_level" gorm:"Column:max_level"`
	Rare                   null.Int8   `json:"rare" gorm:"Column:rare"`
	Raid                   null.Int8   `json:"raid" gorm:"Column:raid"`
	Race                   null.String `json:"race" gorm:"Column:race"`
	Class                  null.String `json:"class" gorm:"Column:class"`
	Bodytype               null.String `json:"bodytype" gorm:"Column:bodytype"`
	Zone                   null.String `json:"zone" gorm:"Column:zone"`
	HotZone                null.Int8   `json:"hot_zone" gorm:"Column:hot_zone"`
	MinExpansion           uint8       `json:"min_expansion" gorm:"Column:min_expansion"`
	MaxExpansion           uint8       `json:"max_expansion" gorm:"Column:max_expansion"`
	ContentFlags           null.String `json:"content_flags" gorm:"Column:content_flags"`
	ContentFlagsDisabled   null.String `json:"content_flags_disabled" gorm:"Column:content_flags_disabled"`
	Loottable              *Loottable  `json:"loottable,omitempty" gorm:"foreignKey:loottable_id;references:id"`
}

func (GlobalLoot) TableName() string {
    return "global_loot"
}

func (GlobalLoot) Relationships() []string {
    return []string{
		"Loottable",
		"Loottable.LoottableEntries",
		"Loottable.LoottableEntries.LootdropEntries",
		"Loottable.LoottableEntries.LootdropEntries.Item",
		"Loottable.LoottableEntries.LootdropEntries.Item.AlternateCurrencies",
		"Loottable.LoottableEntries.LootdropEntries.Item.CharacterCorpseItems",
		"Loottable.LoottableEntries.LootdropEntries.Item.DiscoveredItems",
		"Loottable.LoottableEntries.LootdropEntries.Item.Doors",
		"Loottable.LoottableEntries.LootdropEntries.Item.Doors.Item",
		"Loottable.LoottableEntries.LootdropEntries.Item.Doors.Zone",
		"Loottable.LoottableEntries.LootdropEntries.Item.Fishings",
		"Loottable.LoottableEntries.LootdropEntries.Item.Fishings.Item",
		"Loottable.LoottableEntries.LootdropEntries.Item.Fishings.NpcType",
		"Loottable.LoottableEntries.LootdropEntries.Item.Fishings.NpcType.AlternateCurrency",
		"Loottable.LoottableEntries.LootdropEntries.Item.Fishings.NpcType.Merchantlists",
		"Loottable.LoottableEntries.LootdropEntries.Item.Fishings.NpcType.Merchantlists.Item",
		"Loottable.LoottableEntries.LootdropEntries.Item.Fishings.NpcType.NpcEmotes",
		"Loottable.LoottableEntries.LootdropEntries.Item.Fishings.NpcType.NpcFactions",
		"Loottable.LoottableEntries.LootdropEntries.Item.Fishings.NpcType.NpcFactions.NpcFactionEntries",
		"Loottable.LoottableEntries.LootdropEntries.Item.Fishings.NpcType.NpcSpells",
		"Loottable.LoottableEntries.LootdropEntries.Item.Fishings.NpcType.NpcSpells.NpcSpellsEntries",
		"Loottable.LoottableEntries.LootdropEntries.Item.Fishings.NpcType.NpcTypesTint",
		"Loottable.LoottableEntries.LootdropEntries.Item.Fishings.NpcType.Spawnentries",
		"Loottable.LoottableEntries.LootdropEntries.Item.Fishings.NpcType.Spawnentries.NpcType",
		"Loottable.LoottableEntries.LootdropEntries.Item.Fishings.NpcType.Spawnentries.Spawngroup",
		"Loottable.LoottableEntries.LootdropEntries.Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2",
		"Loottable.LoottableEntries.LootdropEntries.Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries",
		"Loottable.LoottableEntries.LootdropEntries.Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup",
		"Loottable.LoottableEntries.LootdropEntries.Item.Fishings.Zone",
		"Loottable.LoottableEntries.LootdropEntries.Item.Forages",
		"Loottable.LoottableEntries.LootdropEntries.Item.Forages.Item",
		"Loottable.LoottableEntries.LootdropEntries.Item.Forages.Zone",
		"Loottable.LoottableEntries.LootdropEntries.Item.GroundSpawns",
		"Loottable.LoottableEntries.LootdropEntries.Item.GroundSpawns.Item",
		"Loottable.LoottableEntries.LootdropEntries.Item.GroundSpawns.Zone",
		"Loottable.LoottableEntries.LootdropEntries.Item.ItemTicks",
		"Loottable.LoottableEntries.LootdropEntries.Item.Keyrings",
		"Loottable.LoottableEntries.LootdropEntries.Item.LootdropEntries",
		"Loottable.LoottableEntries.LootdropEntries.Item.Merchantlists",
		"Loottable.LoottableEntries.LootdropEntries.Item.Merchantlists.Item",
		"Loottable.LoottableEntries.LootdropEntries.Item.ObjectContents",
		"Loottable.LoottableEntries.LootdropEntries.Item.Objects",
		"Loottable.LoottableEntries.LootdropEntries.Item.Objects.Item",
		"Loottable.LoottableEntries.LootdropEntries.Item.Objects.Zone",
		"Loottable.LoottableEntries.LootdropEntries.Item.StartingItems",
		"Loottable.LoottableEntries.LootdropEntries.Item.StartingItems.Item",
		"Loottable.LoottableEntries.LootdropEntries.Item.StartingItems.Zone",
		"Loottable.LoottableEntries.LootdropEntries.Item.TradeskillRecipeEntries",
		"Loottable.LoottableEntries.LootdropEntries.Item.TradeskillRecipeEntries.TradeskillRecipe",
		"Loottable.LoottableEntries.LootdropEntries.Item.TributeLevels",
		"Loottable.LoottableEntries.LootdropEntries.Lootdrop",
		"Loottable.LoottableEntries.LootdropEntries.Lootdrop.LootdropEntries",
		"Loottable.LoottableEntries.LootdropEntries.Lootdrop.LoottableEntries",
		"Loottable.NpcTypes",
		"Loottable.NpcTypes.AlternateCurrency",
		"Loottable.NpcTypes.Merchantlists",
		"Loottable.NpcTypes.Merchantlists.Item",
		"Loottable.NpcTypes.Merchantlists.Item.AlternateCurrencies",
		"Loottable.NpcTypes.Merchantlists.Item.CharacterCorpseItems",
		"Loottable.NpcTypes.Merchantlists.Item.DiscoveredItems",
		"Loottable.NpcTypes.Merchantlists.Item.Doors",
		"Loottable.NpcTypes.Merchantlists.Item.Doors.Item",
		"Loottable.NpcTypes.Merchantlists.Item.Doors.Zone",
		"Loottable.NpcTypes.Merchantlists.Item.Fishings",
		"Loottable.NpcTypes.Merchantlists.Item.Fishings.Item",
		"Loottable.NpcTypes.Merchantlists.Item.Fishings.NpcType",
		"Loottable.NpcTypes.Merchantlists.Item.Fishings.Zone",
		"Loottable.NpcTypes.Merchantlists.Item.Forages",
		"Loottable.NpcTypes.Merchantlists.Item.Forages.Item",
		"Loottable.NpcTypes.Merchantlists.Item.Forages.Zone",
		"Loottable.NpcTypes.Merchantlists.Item.GroundSpawns",
		"Loottable.NpcTypes.Merchantlists.Item.GroundSpawns.Item",
		"Loottable.NpcTypes.Merchantlists.Item.GroundSpawns.Zone",
		"Loottable.NpcTypes.Merchantlists.Item.ItemTicks",
		"Loottable.NpcTypes.Merchantlists.Item.Keyrings",
		"Loottable.NpcTypes.Merchantlists.Item.LootdropEntries",
		"Loottable.NpcTypes.Merchantlists.Item.LootdropEntries.Item",
		"Loottable.NpcTypes.Merchantlists.Item.LootdropEntries.Lootdrop",
		"Loottable.NpcTypes.Merchantlists.Item.LootdropEntries.Lootdrop.LootdropEntries",
		"Loottable.NpcTypes.Merchantlists.Item.LootdropEntries.Lootdrop.LoottableEntries",
		"Loottable.NpcTypes.Merchantlists.Item.LootdropEntries.Lootdrop.LoottableEntries.LootdropEntries",
		"Loottable.NpcTypes.Merchantlists.Item.Merchantlists",
		"Loottable.NpcTypes.Merchantlists.Item.ObjectContents",
		"Loottable.NpcTypes.Merchantlists.Item.Objects",
		"Loottable.NpcTypes.Merchantlists.Item.Objects.Item",
		"Loottable.NpcTypes.Merchantlists.Item.Objects.Zone",
		"Loottable.NpcTypes.Merchantlists.Item.StartingItems",
		"Loottable.NpcTypes.Merchantlists.Item.StartingItems.Item",
		"Loottable.NpcTypes.Merchantlists.Item.StartingItems.Zone",
		"Loottable.NpcTypes.Merchantlists.Item.TradeskillRecipeEntries",
		"Loottable.NpcTypes.Merchantlists.Item.TradeskillRecipeEntries.TradeskillRecipe",
		"Loottable.NpcTypes.Merchantlists.Item.TributeLevels",
		"Loottable.NpcTypes.NpcEmotes",
		"Loottable.NpcTypes.NpcFactions",
		"Loottable.NpcTypes.NpcFactions.NpcFactionEntries",
		"Loottable.NpcTypes.NpcSpells",
		"Loottable.NpcTypes.NpcSpells.NpcSpellsEntries",
		"Loottable.NpcTypes.NpcTypesTint",
		"Loottable.NpcTypes.Spawnentries",
		"Loottable.NpcTypes.Spawnentries.NpcType",
		"Loottable.NpcTypes.Spawnentries.Spawngroup",
		"Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2",
		"Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries",
		"Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup",
	}
}

func (GlobalLoot) Connection() string {
    return "eqemu_content"
}
