package models

import (
	"github.com/volatiletech/null/v8"
)

type Door struct {
	ID                     int         `json:"id" gorm:"Column:id"`
	Doorid                 int16       `json:"doorid" gorm:"Column:doorid"`
	Zone                   null.String `json:"zone" gorm:"Column:zone"`
	Version                int16       `json:"version" gorm:"Column:version"`
	Name                   string      `json:"name" gorm:"Column:name"`
	PosY                   float32     `json:"pos_y" gorm:"Column:pos_y"`
	PosX                   float32     `json:"pos_x" gorm:"Column:pos_x"`
	PosZ                   float32     `json:"pos_z" gorm:"Column:pos_z"`
	Heading                float32     `json:"heading" gorm:"Column:heading"`
	Opentype               int16       `json:"opentype" gorm:"Column:opentype"`
	Guild                  int16       `json:"guild" gorm:"Column:guild"`
	Lockpick               int16       `json:"lockpick" gorm:"Column:lockpick"`
	Keyitem                int         `json:"keyitem" gorm:"Column:keyitem"`
	Nokeyring              uint8       `json:"nokeyring" gorm:"Column:nokeyring"`
	Triggerdoor            int16       `json:"triggerdoor" gorm:"Column:triggerdoor"`
	Triggertype            int16       `json:"triggertype" gorm:"Column:triggertype"`
	DisableTimer           int8        `json:"disable_timer" gorm:"Column:disable_timer"`
	Doorisopen             int16       `json:"doorisopen" gorm:"Column:doorisopen"`
	DoorParam              int         `json:"door_param" gorm:"Column:door_param"`
	DestZone               null.String `json:"dest_zone" gorm:"Column:dest_zone"`
	DestInstance           uint        `json:"dest_instance" gorm:"Column:dest_instance"`
	DestX                  float32     `json:"dest_x" gorm:"Column:dest_x"`
	DestY                  float32     `json:"dest_y" gorm:"Column:dest_y"`
	DestZ                  float32     `json:"dest_z" gorm:"Column:dest_z"`
	DestHeading            float32     `json:"dest_heading" gorm:"Column:dest_heading"`
	InvertState            int         `json:"invert_state" gorm:"Column:invert_state"`
	Incline                int         `json:"incline" gorm:"Column:incline"`
	Size                   uint16      `json:"size" gorm:"Column:size"`
	Buffer                 float32     `json:"buffer" gorm:"Column:buffer"`
	ClientVersionMask      uint        `json:"client_version_mask" gorm:"Column:client_version_mask"`
	IsLdonDoor             int16       `json:"is_ldon_door" gorm:"Column:is_ldon_door"`
	MinExpansion           uint8       `json:"min_expansion" gorm:"Column:min_expansion"`
	MaxExpansion           uint8       `json:"max_expansion" gorm:"Column:max_expansion"`
	ContentFlags           null.String `json:"content_flags" gorm:"Column:content_flags"`
	ContentFlagsDisabled   null.String `json:"content_flags_disabled" gorm:"Column:content_flags_disabled"`
	Item                   *Item       `json:"item,omitempty" gorm:"foreignKey:keyitem;references:id"`
}

func (Door) TableName() string {
    return "doors"
}

func (Door) Relationships() []string {
    return []string{
		"Item",
		"Item.AlternateCurrencies",
		"Item.CharacterCorpseItems",
		"Item.DiscoveredItems",
		"Item.Doors",
		"Item.Fishings",
		"Item.Fishings.Item",
		"Item.Fishings.NpcType",
		"Item.Fishings.NpcType.AlternateCurrency",
		"Item.Fishings.NpcType.Merchantlists",
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
		"Item.ItemTicks",
		"Item.Keyrings",
		"Item.LootdropEntries",
		"Item.LootdropEntries.Item",
		"Item.LootdropEntries.Lootdrop",
		"Item.LootdropEntries.Lootdrop.LootdropEntries",
		"Item.LootdropEntries.Lootdrop.LoottableEntries",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.LootdropEntries",
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
	}
}

func (Door) Connection() string {
    return "eqemu_content"
}
