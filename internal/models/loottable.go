package models

import (
	"github.com/volatiletech/null/v8"
)

type Loottable struct {
	ID                     uint             `json:"id" gorm:"Column:id"`
	Name                   string           `json:"name" gorm:"Column:name"`
	Mincash                uint             `json:"mincash" gorm:"Column:mincash"`
	Maxcash                uint             `json:"maxcash" gorm:"Column:maxcash"`
	Avgcoin                uint             `json:"avgcoin" gorm:"Column:avgcoin"`
	Done                   int8             `json:"done" gorm:"Column:done"`
	MinExpansion           uint8            `json:"min_expansion" gorm:"Column:min_expansion"`
	MaxExpansion           uint8            `json:"max_expansion" gorm:"Column:max_expansion"`
	ContentFlags           null.String      `json:"content_flags" gorm:"Column:content_flags"`
	ContentFlagsDisabled   null.String      `json:"content_flags_disabled" gorm:"Column:content_flags_disabled"`
	LoottableEntries       []LoottableEntry `json:"loottable_entries,omitempty" gorm:"foreignKey:loottable_id;references:id"`
	NpcTypes               []NpcType        `json:"npc_types,omitempty" gorm:"foreignKey:loottable_id;references:id"`
}

func (Loottable) TableName() string {
    return "loottable"
}

func (Loottable) Relationships() []string {
    return []string{
		"LoottableEntries",
		"LoottableEntries.LootdropEntries",
		"LoottableEntries.LootdropEntries.Item",
		"LoottableEntries.LootdropEntries.Item.AlternateCurrencies",
		"LoottableEntries.LootdropEntries.Item.CharacterCorpseItems",
		"LoottableEntries.LootdropEntries.Item.DiscoveredItems",
		"LoottableEntries.LootdropEntries.Item.Doors",
		"LoottableEntries.LootdropEntries.Item.Doors.Item",
		"LoottableEntries.LootdropEntries.Item.Doors.Zone",
		"LoottableEntries.LootdropEntries.Item.Fishings",
		"LoottableEntries.LootdropEntries.Item.Fishings.Item",
		"LoottableEntries.LootdropEntries.Item.Fishings.NpcType",
		"LoottableEntries.LootdropEntries.Item.Fishings.NpcType.AlternateCurrency",
		"LoottableEntries.LootdropEntries.Item.Fishings.NpcType.Merchantlists",
		"LoottableEntries.LootdropEntries.Item.Fishings.NpcType.Merchantlists.Item",
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
		"LoottableEntries.LootdropEntries.Item.GroundSpawns",
		"LoottableEntries.LootdropEntries.Item.GroundSpawns.Item",
		"LoottableEntries.LootdropEntries.Item.GroundSpawns.Zone",
		"LoottableEntries.LootdropEntries.Item.ItemTicks",
		"LoottableEntries.LootdropEntries.Item.Keyrings",
		"LoottableEntries.LootdropEntries.Item.LootdropEntries",
		"LoottableEntries.LootdropEntries.Item.Merchantlists",
		"LoottableEntries.LootdropEntries.Item.Merchantlists.Item",
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
		"LoottableEntries.LootdropEntries.Lootdrop.LootdropEntries",
		"LoottableEntries.LootdropEntries.Lootdrop.LoottableEntries",
		"NpcTypes",
		"NpcTypes.AlternateCurrency",
		"NpcTypes.Merchantlists",
		"NpcTypes.Merchantlists.Item",
		"NpcTypes.Merchantlists.Item.AlternateCurrencies",
		"NpcTypes.Merchantlists.Item.CharacterCorpseItems",
		"NpcTypes.Merchantlists.Item.DiscoveredItems",
		"NpcTypes.Merchantlists.Item.Doors",
		"NpcTypes.Merchantlists.Item.Doors.Item",
		"NpcTypes.Merchantlists.Item.Doors.Zone",
		"NpcTypes.Merchantlists.Item.Fishings",
		"NpcTypes.Merchantlists.Item.Fishings.Item",
		"NpcTypes.Merchantlists.Item.Fishings.NpcType",
		"NpcTypes.Merchantlists.Item.Fishings.Zone",
		"NpcTypes.Merchantlists.Item.Forages",
		"NpcTypes.Merchantlists.Item.Forages.Item",
		"NpcTypes.Merchantlists.Item.Forages.Zone",
		"NpcTypes.Merchantlists.Item.GroundSpawns",
		"NpcTypes.Merchantlists.Item.GroundSpawns.Item",
		"NpcTypes.Merchantlists.Item.GroundSpawns.Zone",
		"NpcTypes.Merchantlists.Item.ItemTicks",
		"NpcTypes.Merchantlists.Item.Keyrings",
		"NpcTypes.Merchantlists.Item.LootdropEntries",
		"NpcTypes.Merchantlists.Item.LootdropEntries.Item",
		"NpcTypes.Merchantlists.Item.LootdropEntries.Lootdrop",
		"NpcTypes.Merchantlists.Item.LootdropEntries.Lootdrop.LootdropEntries",
		"NpcTypes.Merchantlists.Item.LootdropEntries.Lootdrop.LoottableEntries",
		"NpcTypes.Merchantlists.Item.LootdropEntries.Lootdrop.LoottableEntries.LootdropEntries",
		"NpcTypes.Merchantlists.Item.Merchantlists",
		"NpcTypes.Merchantlists.Item.ObjectContents",
		"NpcTypes.Merchantlists.Item.Objects",
		"NpcTypes.Merchantlists.Item.Objects.Item",
		"NpcTypes.Merchantlists.Item.Objects.Zone",
		"NpcTypes.Merchantlists.Item.StartingItems",
		"NpcTypes.Merchantlists.Item.StartingItems.Item",
		"NpcTypes.Merchantlists.Item.StartingItems.Zone",
		"NpcTypes.Merchantlists.Item.TradeskillRecipeEntries",
		"NpcTypes.Merchantlists.Item.TradeskillRecipeEntries.TradeskillRecipe",
		"NpcTypes.Merchantlists.Item.TributeLevels",
		"NpcTypes.NpcEmotes",
		"NpcTypes.NpcFactions",
		"NpcTypes.NpcFactions.NpcFactionEntries",
		"NpcTypes.NpcSpells",
		"NpcTypes.NpcSpells.NpcSpellsEntries",
		"NpcTypes.NpcTypesTint",
		"NpcTypes.Spawnentries",
		"NpcTypes.Spawnentries.NpcType",
		"NpcTypes.Spawnentries.Spawngroup",
		"NpcTypes.Spawnentries.Spawngroup.Spawn2",
		"NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries",
		"NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup",
	}
}

func (Loottable) Connection() string {
    return "eqemu_content"
}
