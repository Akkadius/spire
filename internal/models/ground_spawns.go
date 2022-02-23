package models

import (
	"github.com/volatiletech/null/v8"
)

type GroundSpawn struct {
	ID                     uint        `json:"id" gorm:"Column:id"`
	Zoneid                 uint        `json:"zoneid" gorm:"Column:zoneid"`
	Version                int16       `json:"version" gorm:"Column:version"`
	MaxX                   float32     `json:"max_x" gorm:"Column:max_x"`
	MaxY                   float32     `json:"max_y" gorm:"Column:max_y"`
	MaxZ                   float32     `json:"max_z" gorm:"Column:max_z"`
	MinX                   float32     `json:"min_x" gorm:"Column:min_x"`
	MinY                   float32     `json:"min_y" gorm:"Column:min_y"`
	Heading                float32     `json:"heading" gorm:"Column:heading"`
	Name                   string      `json:"name" gorm:"Column:name"`
	Item                   uint        `json:"item" gorm:"Column:item"`
	MaxAllowed             uint        `json:"max_allowed" gorm:"Column:max_allowed"`
	Comment                string      `json:"comment" gorm:"Column:comment"`
	RespawnTimer           uint        `json:"respawn_timer" gorm:"Column:respawn_timer"`
	MinExpansion           uint8       `json:"min_expansion" gorm:"Column:min_expansion"`
	MaxExpansion           uint8       `json:"max_expansion" gorm:"Column:max_expansion"`
	ContentFlags           null.String `json:"content_flags" gorm:"Column:content_flags"`
	ContentFlagsDisabled   null.String `json:"content_flags_disabled" gorm:"Column:content_flags_disabled"`
	Zone                   *Zone       `json:"zone,omitempty" gorm:"foreignKey:zoneid;references:zoneidnumber"`
	ItemRelation           *Item       `json:"item,omitempty" gorm:"foreignKey:item;references:id"`
}

func (GroundSpawn) TableName() string {
    return "ground_spawns"
}

func (GroundSpawn) Relationships() []string {
    return []string{
		"Item",
		"Item.AlternateCurrencies",
		"Item.CharacterCorpseItems",
		"Item.DiscoveredItems",
		"Item.Doors",
		"Item.Doors.Item",
		"Item.Doors.Zone",
		"Item.Fishings",
		"Item.Fishings.Item",
		"Item.Fishings.NpcType",
		"Item.Fishings.NpcType.AlternateCurrency",
		"Item.Fishings.NpcType.Merchantlists",
		"Item.Fishings.NpcType.Merchantlists.Item",
		"Item.Fishings.NpcType.NpcEmotes",
		"Item.Fishings.NpcType.NpcFactions",
		"Item.Fishings.NpcType.NpcFactions.NpcFactionEntries",
		"Item.Fishings.NpcType.NpcSpells",
		"Item.Fishings.NpcType.NpcSpells.NpcSpellsEntries",
		"Item.Fishings.NpcType.NpcTypesTint",
		"Item.Fishings.NpcType.Spawnentries",
		"Item.Fishings.NpcType.Spawnentries.NpcType",
		"Item.Fishings.NpcType.Spawnentries.Spawngroup",
		"Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2",
		"Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries",
		"Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup",
		"Item.Fishings.Zone",
		"Item.Forages",
		"Item.Forages.Item",
		"Item.Forages.Zone",
		"Item.GroundSpawns",
		"Item.ItemTicks",
		"Item.Keyrings",
		"Item.LootdropEntries",
		"Item.LootdropEntries.Item",
		"Item.LootdropEntries.Lootdrop",
		"Item.LootdropEntries.Lootdrop.LootdropEntries",
		"Item.LootdropEntries.Lootdrop.LoottableEntries",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.LootdropEntries",
		"Item.Merchantlists",
		"Item.Merchantlists.Item",
		"Item.ObjectContents",
		"Item.Objects",
		"Item.Objects.Item",
		"Item.Objects.Zone",
		"Item.StartingItems",
		"Item.StartingItems.Item",
		"Item.StartingItems.Zone",
		"Item.TradeskillRecipeEntries",
		"Item.TradeskillRecipeEntries.TradeskillRecipe",
		"Item.TributeLevels",
		"Zone",
	}
}

func (GroundSpawn) Connection() string {
    return "eqemu_content"
}
