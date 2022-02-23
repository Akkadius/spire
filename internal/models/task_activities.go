package models

import (
	"github.com/volatiletech/null/v8"
)

type TaskActivity struct {
	Taskid               uint       `json:"taskid" gorm:"Column:taskid"`
	Activityid           uint       `json:"activityid" gorm:"Column:activityid"`
	Step                 uint       `json:"step" gorm:"Column:step"`
	Activitytype         uint8      `json:"activitytype" gorm:"Column:activitytype"`
	TargetName           string     `json:"target_name" gorm:"Column:target_name"`
	ItemList             string     `json:"item_list" gorm:"Column:item_list"`
	SkillList            string     `json:"skill_list" gorm:"Column:skill_list"`
	SpellList            string     `json:"spell_list" gorm:"Column:spell_list"`
	DescriptionOverride  string     `json:"description_override" gorm:"Column:description_override"`
	Goalid               uint       `json:"goalid" gorm:"Column:goalid"`
	Goalmethod           uint       `json:"goalmethod" gorm:"Column:goalmethod"`
	Goalcount            null.Int   `json:"goalcount" gorm:"Column:goalcount"`
	Delivertonpc         uint       `json:"delivertonpc" gorm:"Column:delivertonpc"`
	Zones                string     `json:"zones" gorm:"Column:zones"`
	Optional             int8       `json:"optional" gorm:"Column:optional"`
	NpcType              *NpcType   `json:"npc_type,omitempty" gorm:"foreignKey:delivertonpc;references:id"`
	Goallists            []Goallist `json:"goallists,omitempty" gorm:"foreignKey:listid;references:goalid"`
}

func (TaskActivity) TableName() string {
    return "task_activities"
}

func (TaskActivity) Relationships() []string {
    return []string{
		"Goallists",
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
		"NpcType.Merchantlists.Item.Fishings.Item",
		"NpcType.Merchantlists.Item.Fishings.NpcType",
		"NpcType.Merchantlists.Item.Fishings.Zone",
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
	}
}

func (TaskActivity) Connection() string {
    return "eqemu_content"
}
