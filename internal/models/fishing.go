package models

import (
	"github.com/volatiletech/null/v8"
)

type Fishing struct {
	ID                     int         `json:"id" gorm:"Column:id"`
	Zoneid                 int         `json:"zoneid" gorm:"Column:zoneid"`
	Itemid                 int         `json:"itemid" gorm:"Column:Itemid"`
	SkillLevel             int16       `json:"skill_level" gorm:"Column:skill_level"`
	Chance                 int16       `json:"chance" gorm:"Column:chance"`
	NpcId                  int         `json:"npc_id" gorm:"Column:npc_id"`
	NpcChance              int         `json:"npc_chance" gorm:"Column:npc_chance"`
	MinExpansion           uint8       `json:"min_expansion" gorm:"Column:min_expansion"`
	MaxExpansion           uint8       `json:"max_expansion" gorm:"Column:max_expansion"`
	ContentFlags           null.String `json:"content_flags" gorm:"Column:content_flags"`
	ContentFlagsDisabled   null.String `json:"content_flags_disabled" gorm:"Column:content_flags_disabled"`
	Item                   *Item       `json:"item,omitempty" gorm:"foreignKey:Itemid;references:id"`
	Zone                   *Zone       `json:"zone,omitempty" gorm:"foreignKey:zoneid;references:zoneidnumber"`
	NpcType                *NpcType    `json:"npc_type,omitempty" gorm:"foreignKey:npc_id;references:id"`
}

func (Fishing) TableName() string {
    return "fishing"
}

func (Fishing) Relationships() []string {
    return []string{
		"Item",
		"Item.AlternateCurrencies",
		"Item.CharacterCorpseItems",
		"Item.DiscoveredItems",
		"Item.Doors",
		"Item.Doors.Item",
		"Item.Doors.Zone",
		"Item.Fishings",
		"Item.Forages",
		"Item.Forages.Item",
		"Item.Forages.Zone",
		"Item.GroundSpawns",
		"Item.GroundSpawns.Item",
		"Item.GroundSpawns.Zone",
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
		"NpcType",
		"NpcType.AlternateCurrency",
		"NpcType.Merchantlists",
		"NpcType.Merchantlists.Item",
		"NpcType.Merchantlists.Item.AlternateCurrencies",
		"NpcType.Merchantlists.Item.CharacterCorpseItems",
		"NpcType.Merchantlists.Item.DiscoveredItems",
		"NpcType.Merchantlists.Item.Doors",
		"NpcType.Merchantlists.Item.Doors.Item",
		"NpcType.Merchantlists.Item.Doors.Zone",
		"NpcType.Merchantlists.Item.Fishings",
		"NpcType.Merchantlists.Item.Forages",
		"NpcType.Merchantlists.Item.Forages.Item",
		"NpcType.Merchantlists.Item.Forages.Zone",
		"NpcType.Merchantlists.Item.GroundSpawns",
		"NpcType.Merchantlists.Item.GroundSpawns.Item",
		"NpcType.Merchantlists.Item.GroundSpawns.Zone",
		"NpcType.Merchantlists.Item.ItemTicks",
		"NpcType.Merchantlists.Item.Keyrings",
		"NpcType.Merchantlists.Item.LootdropEntries",
		"NpcType.Merchantlists.Item.LootdropEntries.Item",
		"NpcType.Merchantlists.Item.LootdropEntries.Lootdrop",
		"NpcType.Merchantlists.Item.LootdropEntries.Lootdrop.LootdropEntries",
		"NpcType.Merchantlists.Item.LootdropEntries.Lootdrop.LoottableEntries",
		"NpcType.Merchantlists.Item.LootdropEntries.Lootdrop.LoottableEntries.LootdropEntries",
		"NpcType.Merchantlists.Item.Merchantlists",
		"NpcType.Merchantlists.Item.ObjectContents",
		"NpcType.Merchantlists.Item.Objects",
		"NpcType.Merchantlists.Item.Objects.Item",
		"NpcType.Merchantlists.Item.Objects.Zone",
		"NpcType.Merchantlists.Item.StartingItems",
		"NpcType.Merchantlists.Item.StartingItems.Item",
		"NpcType.Merchantlists.Item.StartingItems.Zone",
		"NpcType.Merchantlists.Item.TradeskillRecipeEntries",
		"NpcType.Merchantlists.Item.TradeskillRecipeEntries.TradeskillRecipe",
		"NpcType.Merchantlists.Item.TributeLevels",
		"NpcType.NpcEmotes",
		"NpcType.NpcFactions",
		"NpcType.NpcFactions.NpcFactionEntries",
		"NpcType.NpcSpells",
		"NpcType.NpcSpells.NpcSpellsEntries",
		"NpcType.NpcTypesTint",
		"NpcType.Spawnentries",
		"NpcType.Spawnentries.NpcType",
		"NpcType.Spawnentries.Spawngroup",
		"NpcType.Spawnentries.Spawngroup.Spawn2",
		"NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries",
		"NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup",
		"Zone",
	}
}

func (Fishing) Connection() string {
    return "eqemu_content"
}
