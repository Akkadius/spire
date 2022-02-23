package models

import (
	"github.com/volatiletech/null/v8"
)

type Inventory struct {
	Charid              uint        `json:"charid" gorm:"Column:charid"`
	Slotid              uint32      `json:"slotid" gorm:"Column:slotid"`
	Itemid              null.Uint   `json:"itemid" gorm:"Column:itemid"`
	Charges             null.Uint16 `json:"charges" gorm:"Column:charges"`
	Color               uint        `json:"color" gorm:"Column:color"`
	Augslot1            uint32      `json:"augslot_1" gorm:"Column:augslot1"`
	Augslot2            uint32      `json:"augslot_2" gorm:"Column:augslot2"`
	Augslot3            uint32      `json:"augslot_3" gorm:"Column:augslot3"`
	Augslot4            uint32      `json:"augslot_4" gorm:"Column:augslot4"`
	Augslot5            null.Uint32 `json:"augslot_5" gorm:"Column:augslot5"`
	Augslot6            int32       `json:"augslot_6" gorm:"Column:augslot6"`
	Instnodrop          uint8       `json:"instnodrop" gorm:"Column:instnodrop"`
	CustomData          null.String `json:"custom_data" gorm:"Column:custom_data"`
	Ornamenticon        uint        `json:"ornamenticon" gorm:"Column:ornamenticon"`
	Ornamentidfile      uint        `json:"ornamentidfile" gorm:"Column:ornamentidfile"`
	OrnamentHeroModel   int         `json:"ornament_hero_model" gorm:"Column:ornament_hero_model"`
	Item                *Item       `json:"item,omitempty" gorm:"foreignKey:itemid;references:id"`
}

func (Inventory) TableName() string {
    return "inventory"
}

func (Inventory) Relationships() []string {
    return []string{
		"Item",
		"Item.AlternateCurrencies",
		"Item.CharacterCorpseItems",
		"Item.DiscoveredItems",
		"Item.Doors",
		"Item.Doors.Item",
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

func (Inventory) Connection() string {
    return ""
}
