package models

import (
	"github.com/volatiletech/null/v8"
)

type Object struct {
	ID                     int         `json:"id" gorm:"Column:id"`
	Zoneid                 uint        `json:"zoneid" gorm:"Column:zoneid"`
	Version                int16       `json:"version" gorm:"Column:version"`
	Xpos                   float32     `json:"xpos" gorm:"Column:xpos"`
	Ypos                   float32     `json:"ypos" gorm:"Column:ypos"`
	Zpos                   float32     `json:"zpos" gorm:"Column:zpos"`
	Heading                float32     `json:"heading" gorm:"Column:heading"`
	Itemid                 int         `json:"itemid" gorm:"Column:itemid"`
	Charges                uint16      `json:"charges" gorm:"Column:charges"`
	Objectname             null.String `json:"objectname" gorm:"Column:objectname"`
	Type                   int         `json:"type" gorm:"Column:type"`
	Icon                   int         `json:"icon" gorm:"Column:icon"`
	SizePercentage         float32     `json:"size_percentage" gorm:"Column:size_percentage"`
	Unknown24              int         `json:"unknown_24" gorm:"Column:unknown24"`
	Unknown60              int         `json:"unknown_60" gorm:"Column:unknown60"`
	Unknown64              int         `json:"unknown_64" gorm:"Column:unknown64"`
	Unknown68              int         `json:"unknown_68" gorm:"Column:unknown68"`
	Unknown72              int         `json:"unknown_72" gorm:"Column:unknown72"`
	Unknown76              int         `json:"unknown_76" gorm:"Column:unknown76"`
	Unknown84              int         `json:"unknown_84" gorm:"Column:unknown84"`
	Size                   float32     `json:"size" gorm:"Column:size"`
	SolidType              int32       `json:"solid_type" gorm:"Column:solid_type"`
	Incline                int         `json:"incline" gorm:"Column:incline"`
	TiltX                  float32     `json:"tilt_x" gorm:"Column:tilt_x"`
	TiltY                  float32     `json:"tilt_y" gorm:"Column:tilt_y"`
	DisplayName            null.String `json:"display_name" gorm:"Column:display_name"`
	MinExpansion           int8        `json:"min_expansion" gorm:"Column:min_expansion"`
	MaxExpansion           int8        `json:"max_expansion" gorm:"Column:max_expansion"`
	ContentFlags           null.String `json:"content_flags" gorm:"Column:content_flags"`
	ContentFlagsDisabled   null.String `json:"content_flags_disabled" gorm:"Column:content_flags_disabled"`
	Zone                   *Zone       `json:"zone,omitempty" gorm:"foreignKey:zoneid;references:zoneidnumber"`
	Item                   *Item       `json:"item,omitempty" gorm:"foreignKey:itemid;references:id"`
}

func (Object) TableName() string {
    return "object"
}

func (Object) Relationships() []string {
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
		"Item.Fishings.NpcType.NpcSpell.BotSpellsEntries",
		"Item.Fishings.NpcType.NpcSpell.BotSpellsEntries.NpcSpell",
		"Item.Fishings.NpcType.NpcSpell.BotSpellsEntries.SpellsNew",
		"Item.Fishings.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Aura",
		"Item.Fishings.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Aura.SpellsNew",
		"Item.Fishings.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.BlockedSpells",
		"Item.Fishings.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.BotSpellsEntries",
		"Item.Fishings.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Damageshieldtypes",
		"Item.Fishings.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items",
		"Item.Fishings.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.NpcSpellsEntries",
		"Item.Fishings.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.NpcSpellsEntries.SpellsNew",
		"Item.Fishings.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.SpellBuckets",
		"Item.Fishings.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.SpellGlobals",
		"Item.Fishings.NpcType.NpcSpell.NpcSpell",
		"Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries",
		"Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew",
		"Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Aura",
		"Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew",
		"Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells",
		"Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.BotSpellsEntries",
		"Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.BotSpellsEntries.NpcSpell",
		"Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.BotSpellsEntries.SpellsNew",
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
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.BotSpellsEntries",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.BotSpellsEntries.NpcSpell",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Aura",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Aura.SpellsNew",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.BlockedSpells",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.BotSpellsEntries",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Damageshieldtypes",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.NpcSpellsEntries",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.NpcSpellsEntries.SpellsNew",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.SpellBuckets",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.SpellGlobals",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpell",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.BotSpellsEntries",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.BotSpellsEntries.NpcSpell",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.BotSpellsEntries.SpellsNew",
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
		"Item.Merchantlists.NpcTypes.NpcSpell.BotSpellsEntries",
		"Item.Merchantlists.NpcTypes.NpcSpell.BotSpellsEntries.NpcSpell",
		"Item.Merchantlists.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew",
		"Item.Merchantlists.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Aura",
		"Item.Merchantlists.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Aura.SpellsNew",
		"Item.Merchantlists.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.BlockedSpells",
		"Item.Merchantlists.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.BotSpellsEntries",
		"Item.Merchantlists.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Damageshieldtypes",
		"Item.Merchantlists.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items",
		"Item.Merchantlists.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.NpcSpellsEntries",
		"Item.Merchantlists.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.NpcSpellsEntries.SpellsNew",
		"Item.Merchantlists.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.SpellBuckets",
		"Item.Merchantlists.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.SpellGlobals",
		"Item.Merchantlists.NpcTypes.NpcSpell.NpcSpell",
		"Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries",
		"Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew",
		"Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura",
		"Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew",
		"Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells",
		"Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.BotSpellsEntries",
		"Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.BotSpellsEntries.NpcSpell",
		"Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.BotSpellsEntries.SpellsNew",
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
		"Item.TradeskillRecipeEntries",
		"Item.TradeskillRecipeEntries.TradeskillRecipe",
		"Item.TradeskillRecipeEntries.TradeskillRecipe.TradeskillRecipeEntries",
		"Item.TributeLevels",
		"Zone",
	}
}

func (Object) Connection() string {
    return "eqemu_content"
}
