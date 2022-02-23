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
		"LoottableEntries.LootdropEntries.Item.Fishings",
		"LoottableEntries.LootdropEntries.Item.Forages",
		"LoottableEntries.LootdropEntries.Item.GroundSpawns",
		"LoottableEntries.LootdropEntries.Item.ItemTicks",
		"LoottableEntries.LootdropEntries.Item.Keyrings",
		"LoottableEntries.LootdropEntries.Lootdrop",
		"NpcTypes",
		"NpcTypes.AlternateCurrency",
		"NpcTypes.Merchantlists",
		"NpcTypes.NpcEmotes",
		"NpcTypes.NpcFactions",
		"NpcTypes.NpcFactions.NpcFactionEntries",
		"NpcTypes.NpcSpells",
		"NpcTypes.NpcSpells.NpcSpellsEntries",
		"NpcTypes.NpcTypesTint",
		"NpcTypes.Spawnentries",
		"NpcTypes.Spawnentries.Spawngroup",
		"NpcTypes.Spawnentries.Spawngroup.Spawn2",
	}
}

func (Loottable) Connection() string {
    return "eqemu_content"
}
