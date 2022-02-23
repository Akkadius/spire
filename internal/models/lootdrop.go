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
		"LootdropEntries.Item.Fishings",
		"LootdropEntries.Item.Forages",
		"LootdropEntries.Item.GroundSpawns",
		"LootdropEntries.Item.ItemTicks",
		"LootdropEntries.Item.Keyrings",
		"LootdropEntries.Item.LootdropEntries",
		"LootdropEntries.Item.Merchantlists",
		"LootdropEntries.Item.ObjectContents",
		"LootdropEntries.Item.Objects",
		"LootdropEntries.Item.StartingItems",
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
		"LoottableEntries.LootdropEntries.Item.Fishings",
		"LoottableEntries.LootdropEntries.Item.Forages",
		"LoottableEntries.LootdropEntries.Item.GroundSpawns",
		"LoottableEntries.LootdropEntries.Item.ItemTicks",
		"LoottableEntries.LootdropEntries.Item.Keyrings",
		"LoottableEntries.LootdropEntries.Item.LootdropEntries",
		"LoottableEntries.LootdropEntries.Item.Merchantlists",
		"LoottableEntries.LootdropEntries.Item.ObjectContents",
		"LoottableEntries.LootdropEntries.Item.Objects",
		"LoottableEntries.LootdropEntries.Item.StartingItems",
		"LoottableEntries.LootdropEntries.Item.TradeskillRecipeEntries",
		"LoottableEntries.LootdropEntries.Item.TradeskillRecipeEntries.TradeskillRecipe",
		"LoottableEntries.LootdropEntries.Item.TributeLevels",
		"LoottableEntries.LootdropEntries.Lootdrop",
	}
}

func (Lootdrop) Connection() string {
    return "eqemu_content"
}
