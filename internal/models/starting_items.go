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
		"Item.AlternateCurrencies.Item",
		"Item.CharacterCorpseItems",
		"Item.DiscoveredItems",
		"Item.Doors",
		"Item.Doors.Item",
		"Item.Fishings",
		"Item.Fishings.Item",
		"Item.Fishings.NpcType",
		"Item.Fishings.NpcType.AlternateCurrency",
		"Item.Fishings.NpcType.AlternateCurrency.Item",
		"Item.Fishings.NpcType.Loottable",
		"Item.Fishings.NpcType.Loottable.LoottableEntries",
		"Item.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop",
		"Item.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries",
		"Item.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item",
		"Item.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Lootdrop",
		"Item.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LoottableEntries",
		"Item.Fishings.NpcType.Loottable.LoottableEntries.Loottable",
		"Item.Fishings.NpcType.Loottable.NpcTypes",
		"Item.Fishings.NpcType.Merchantlists",
		"Item.Fishings.NpcType.Merchantlists.Items",
		"Item.Fishings.NpcType.Merchantlists.NpcTypes",
		"Item.Fishings.NpcType.NpcEmotes",
		"Item.Fishings.NpcType.NpcFactions",
		"Item.Fishings.NpcType.NpcFactions.NpcFactionEntries",
		"Item.Fishings.NpcType.NpcFactions.NpcFactionEntries.FactionList",
		"Item.Fishings.NpcType.NpcSpell",
		"Item.Fishings.NpcType.NpcSpell.NpcSpell",
		"Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries",
		"Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew",
		"Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Aura",
		"Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew",
		"Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells",
		"Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes",
		"Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items",
		"Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries",
		"Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets",
		"Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals",
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
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Lootdrop",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.LoottableEntries",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Loottable",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.NpcTypes",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcEmotes",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions.NpcFactionEntries",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions.NpcFactionEntries.FactionList",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpell",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcTypesTint",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.NpcType",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup",
		"Item.Merchantlists",
		"Item.Merchantlists.Items",
		"Item.Merchantlists.NpcTypes",
		"Item.Merchantlists.NpcTypes.AlternateCurrency",
		"Item.Merchantlists.NpcTypes.AlternateCurrency.Item",
		"Item.Merchantlists.NpcTypes.Loottable",
		"Item.Merchantlists.NpcTypes.Loottable.LoottableEntries",
		"Item.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop",
		"Item.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries",
		"Item.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item",
		"Item.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Lootdrop",
		"Item.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LoottableEntries",
		"Item.Merchantlists.NpcTypes.Loottable.LoottableEntries.Loottable",
		"Item.Merchantlists.NpcTypes.Loottable.NpcTypes",
		"Item.Merchantlists.NpcTypes.Merchantlists",
		"Item.Merchantlists.NpcTypes.NpcEmotes",
		"Item.Merchantlists.NpcTypes.NpcFactions",
		"Item.Merchantlists.NpcTypes.NpcFactions.NpcFactionEntries",
		"Item.Merchantlists.NpcTypes.NpcFactions.NpcFactionEntries.FactionList",
		"Item.Merchantlists.NpcTypes.NpcSpell",
		"Item.Merchantlists.NpcTypes.NpcSpell.NpcSpell",
		"Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries",
		"Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew",
		"Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura",
		"Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew",
		"Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells",
		"Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes",
		"Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items",
		"Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries",
		"Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets",
		"Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals",
		"Item.Merchantlists.NpcTypes.NpcTypesTint",
		"Item.Merchantlists.NpcTypes.Spawnentries",
		"Item.Merchantlists.NpcTypes.Spawnentries.NpcType",
		"Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup",
		"Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2",
		"Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries",
		"Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup",
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
