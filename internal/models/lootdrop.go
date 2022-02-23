package models

import (
	"github.com/volatiletech/null/v8"
)

type Lootdrop struct {
	ID                     uint             `json:"id" gorm:"Column:id"`
	Name                   string           `json:"name" gorm:"Column:name"`
	MinExpansion           uint8            `json:"min_expansion" gorm:"Column:min_expansion"`
	MaxExpansion           uint8            `json:"max_expansion" gorm:"Column:max_expansion"`
	ContentFlags           null.String      `json:"content_flags" gorm:"Column:content_flags"`
	ContentFlagsDisabled   null.String      `json:"content_flags_disabled" gorm:"Column:content_flags_disabled"`
	LootdropEntries        []LootdropEntry  `json:"lootdrop_entries,omitempty" gorm:"foreignKey:lootdrop_id;references:id"`
	LoottableEntries       []LoottableEntry `json:"loottable_entries,omitempty" gorm:"foreignKey:loottable_id;references:id"`
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
		"LootdropEntries.Item.Fishings.NpcType.Merchantlists",
		"LootdropEntries.Item.Fishings.NpcType.NpcEmotes",
		"LootdropEntries.Item.Fishings.NpcType.NpcFactions",
		"LootdropEntries.Item.Fishings.NpcType.NpcFactions.NpcFactionEntries",
		"LootdropEntries.Item.Fishings.NpcType.NpcSpells",
		"LootdropEntries.Item.Fishings.NpcType.NpcSpells.NpcSpellsEntries",
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
		"LootdropEntries.Item.ItemTicks",
		"LootdropEntries.Item.Keyrings",
		"LootdropEntries.Item.LootdropEntries",
		"LootdropEntries.Item.ObjectContents",
		"LootdropEntries.Item.Objects",
		"LootdropEntries.Item.Objects.Item",
		"LootdropEntries.Item.Objects.Zone",
		"LootdropEntries.Item.StartingItems",
		"LootdropEntries.Item.StartingItems.Item",
		"LootdropEntries.Item.StartingItems.Zone",
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
		"LoottableEntries.LootdropEntries.Item.Fishings.NpcType.Merchantlists",
		"LoottableEntries.LootdropEntries.Item.Fishings.NpcType.NpcEmotes",
		"LoottableEntries.LootdropEntries.Item.Fishings.NpcType.NpcFactions",
		"LoottableEntries.LootdropEntries.Item.Fishings.NpcType.NpcFactions.NpcFactionEntries",
		"LoottableEntries.LootdropEntries.Item.Fishings.NpcType.NpcSpells",
		"LoottableEntries.LootdropEntries.Item.Fishings.NpcType.NpcSpells.NpcSpellsEntries",
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
		"LoottableEntries.LootdropEntries.Item.ItemTicks",
		"LoottableEntries.LootdropEntries.Item.Keyrings",
		"LoottableEntries.LootdropEntries.Item.LootdropEntries",
		"LoottableEntries.LootdropEntries.Item.ObjectContents",
		"LoottableEntries.LootdropEntries.Item.Objects",
		"LoottableEntries.LootdropEntries.Item.Objects.Item",
		"LoottableEntries.LootdropEntries.Item.Objects.Zone",
		"LoottableEntries.LootdropEntries.Item.StartingItems",
		"LoottableEntries.LootdropEntries.Item.StartingItems.Item",
		"LoottableEntries.LootdropEntries.Item.StartingItems.Zone",
		"LoottableEntries.LootdropEntries.Item.TradeskillRecipeEntries",
		"LoottableEntries.LootdropEntries.Item.TradeskillRecipeEntries.TradeskillRecipe",
		"LoottableEntries.LootdropEntries.Item.TributeLevels",
		"LoottableEntries.LootdropEntries.Lootdrop",
	}
}

func (Lootdrop) Connection() string {
    return "eqemu_content"
}
