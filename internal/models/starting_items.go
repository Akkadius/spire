package models

import (
	"github.com/volatiletech/null/v8"
)

type StartingItem struct {
	ID                     uint        `json:"id" gorm:"Column:id"`
	Race                   int         `json:"race" gorm:"Column:race"`
	Class                  int         `json:"class" gorm:"Column:class"`
	Deityid                int         `json:"deityid" gorm:"Column:deityid"`
	Zoneid                 int         `json:"zoneid" gorm:"Column:zoneid"`
	Itemid                 int         `json:"itemid" gorm:"Column:itemid"`
	ItemCharges            uint8       `json:"item_charges" gorm:"Column:item_charges"`
	Gm                     int8        `json:"gm" gorm:"Column:gm"`
	Slot                   int32       `json:"slot" gorm:"Column:slot"`
	MinExpansion           int8        `json:"min_expansion" gorm:"Column:min_expansion"`
	MaxExpansion           int8        `json:"max_expansion" gorm:"Column:max_expansion"`
	ContentFlags           null.String `json:"content_flags" gorm:"Column:content_flags"`
	ContentFlagsDisabled   null.String `json:"content_flags_disabled" gorm:"Column:content_flags_disabled"`
	Zone                   *Zone       `json:"zone,omitempty" gorm:"foreignKey:zoneid;references:zoneidnumber"`
	Item                   *Item       `json:"item,omitempty" gorm:"foreignKey:itemid;references:id"`
}

func (StartingItem) TableName() string {
    return "starting_items"
}

func (StartingItem) Relationships() []string {
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
		"Item.GroundSpawns",
		"Item.GroundSpawns.Zone",
		"Item.ItemTicks",
		"Item.Keyrings",
		"Item.LootdropEntries",
		"Item.LootdropEntries.Item",
		"Item.LootdropEntries.Lootdrop",
		"Item.LootdropEntries.Lootdrop.LootdropEntries",
		"Item.LootdropEntries.Lootdrop.LoottableEntries",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.LootdropEntries",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.LoottableEntries",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcEmotes",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions.NpcFactionEntries",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpells",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpells.NpcSpellsEntries",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcTypesTint",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.NpcType",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup",
		"Item.ObjectContents",
		"Item.Objects",
		"Item.Objects.Item",
		"Item.Objects.Zone",
		"Item.StartingItems",
		"Item.TradeskillRecipeEntries",
		"Item.TradeskillRecipeEntries.TradeskillRecipe",
		"Item.TributeLevels",
		"Zone",
	}
}

func (StartingItem) Connection() string {
    return "eqemu_content"
}
