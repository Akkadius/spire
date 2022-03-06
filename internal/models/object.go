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
	Unknown08              int32       `json:"unknown_08" gorm:"Column:unknown08"`
	Unknown10              int32       `json:"unknown_10" gorm:"Column:unknown10"`
	Unknown20              int         `json:"unknown_20" gorm:"Column:unknown20"`
	Unknown24              int         `json:"unknown_24" gorm:"Column:unknown24"`
	Unknown60              int         `json:"unknown_60" gorm:"Column:unknown60"`
	Unknown64              int         `json:"unknown_64" gorm:"Column:unknown64"`
	Unknown68              int         `json:"unknown_68" gorm:"Column:unknown68"`
	Unknown72              int         `json:"unknown_72" gorm:"Column:unknown72"`
	Unknown76              int         `json:"unknown_76" gorm:"Column:unknown76"`
	Unknown84              int         `json:"unknown_84" gorm:"Column:unknown84"`
	Size                   float32     `json:"size" gorm:"Column:size"`
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
		"Item.CharacterCorpseItems",
		"Item.DiscoveredItems",
		"Item.Doors",
		"Item.Doors.Item",
		"Item.Fishings",
		"Item.Fishings.Item",
		"Item.Fishings.NpcType",
		"Item.Fishings.NpcType.AlternateCurrency",
		"Item.Fishings.NpcType.Loottable",
		"Item.Fishings.NpcType.Loottable.LoottableEntries",
		"Item.Fishings.NpcType.Loottable.LoottableEntries.LootdropEntries",
		"Item.Fishings.NpcType.Loottable.LoottableEntries.LootdropEntries.Item",
		"Item.Fishings.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop",
		"Item.Fishings.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LootdropEntries",
		"Item.Fishings.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LoottableEntries",
		"Item.Fishings.NpcType.Loottable.LoottableEntries.Loottable",
		"Item.Fishings.NpcType.Loottable.NpcTypes",
		"Item.Fishings.NpcType.Merchantlists",
		"Item.Fishings.NpcType.Merchantlists.NpcType",
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
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Loottable",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.NpcType",
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
		"Item.Merchantlists",
		"Item.Merchantlists.NpcType",
		"Item.Merchantlists.NpcType.AlternateCurrency",
		"Item.Merchantlists.NpcType.Loottable",
		"Item.Merchantlists.NpcType.Loottable.LoottableEntries",
		"Item.Merchantlists.NpcType.Loottable.LoottableEntries.LootdropEntries",
		"Item.Merchantlists.NpcType.Loottable.LoottableEntries.LootdropEntries.Item",
		"Item.Merchantlists.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop",
		"Item.Merchantlists.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LootdropEntries",
		"Item.Merchantlists.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LoottableEntries",
		"Item.Merchantlists.NpcType.Loottable.LoottableEntries.Loottable",
		"Item.Merchantlists.NpcType.Loottable.NpcTypes",
		"Item.Merchantlists.NpcType.Merchantlists",
		"Item.Merchantlists.NpcType.NpcEmotes",
		"Item.Merchantlists.NpcType.NpcFactions",
		"Item.Merchantlists.NpcType.NpcFactions.NpcFactionEntries",
		"Item.Merchantlists.NpcType.NpcSpells",
		"Item.Merchantlists.NpcType.NpcSpells.NpcSpellsEntries",
		"Item.Merchantlists.NpcType.NpcTypesTint",
		"Item.Merchantlists.NpcType.Spawnentries",
		"Item.Merchantlists.NpcType.Spawnentries.NpcType",
		"Item.Merchantlists.NpcType.Spawnentries.Spawngroup",
		"Item.Merchantlists.NpcType.Spawnentries.Spawngroup.Spawn2",
		"Item.Merchantlists.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries",
		"Item.Merchantlists.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup",
		"Item.ObjectContents",
		"Item.Objects",
		"Item.StartingItems",
		"Item.StartingItems.Item",
		"Item.StartingItems.Zone",
		"Item.Tasks",
		"Item.Tasks.TaskActivities",
		"Item.Tasks.TaskActivities.Goallists",
		"Item.Tasks.TaskActivities.NpcType",
		"Item.Tasks.TaskActivities.NpcType.AlternateCurrency",
		"Item.Tasks.TaskActivities.NpcType.Loottable",
		"Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries",
		"Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries",
		"Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries.Item",
		"Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop",
		"Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LootdropEntries",
		"Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LoottableEntries",
		"Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.Loottable",
		"Item.Tasks.TaskActivities.NpcType.Loottable.NpcTypes",
		"Item.Tasks.TaskActivities.NpcType.Merchantlists",
		"Item.Tasks.TaskActivities.NpcType.Merchantlists.NpcType",
		"Item.Tasks.TaskActivities.NpcType.NpcEmotes",
		"Item.Tasks.TaskActivities.NpcType.NpcFactions",
		"Item.Tasks.TaskActivities.NpcType.NpcFactions.NpcFactionEntries",
		"Item.Tasks.TaskActivities.NpcType.NpcSpells",
		"Item.Tasks.TaskActivities.NpcType.NpcSpells.NpcSpellsEntries",
		"Item.Tasks.TaskActivities.NpcType.NpcTypesTint",
		"Item.Tasks.TaskActivities.NpcType.Spawnentries",
		"Item.Tasks.TaskActivities.NpcType.Spawnentries.NpcType",
		"Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup",
		"Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2",
		"Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries",
		"Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup",
		"Item.Tasks.Tasksets",
		"Item.TradeskillRecipeEntries",
		"Item.TradeskillRecipeEntries.TradeskillRecipe",
		"Item.TributeLevels",
		"Zone",
	}
}

func (Object) Connection() string {
    return "eqemu_content"
}
