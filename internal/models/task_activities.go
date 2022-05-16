package models

import (
	"github.com/volatiletech/null/v8"
)

type TaskActivity struct {
	Taskid               uint        `json:"taskid" gorm:"Column:taskid"`
	Activityid           uint        `json:"activityid" gorm:"Column:activityid"`
	Step                 uint        `json:"step" gorm:"Column:step"`
	Activitytype         uint8       `json:"activitytype" gorm:"Column:activitytype"`
	TargetName           string      `json:"target_name" gorm:"Column:target_name"`
	ItemList             string      `json:"item_list" gorm:"Column:item_list"`
	SkillList            string      `json:"skill_list" gorm:"Column:skill_list"`
	SpellList            string      `json:"spell_list" gorm:"Column:spell_list"`
	DescriptionOverride  string      `json:"description_override" gorm:"Column:description_override"`
	Goalid               uint        `json:"goalid" gorm:"Column:goalid"`
	GoalMatchList        null.String `json:"goal_match_list" gorm:"Column:goal_match_list"`
	Goalmethod           uint        `json:"goalmethod" gorm:"Column:goalmethod"`
	Goalcount            null.Int    `json:"goalcount" gorm:"Column:goalcount"`
	Delivertonpc         uint        `json:"delivertonpc" gorm:"Column:delivertonpc"`
	Zones                string      `json:"zones" gorm:"Column:zones"`
	Optional             int8        `json:"optional" gorm:"Column:optional"`
	NpcType              *NpcType    `json:"npc_type,omitempty" gorm:"foreignKey:delivertonpc;references:id"`
	Goallists            []Goallist  `json:"goallists,omitempty" gorm:"foreignKey:listid;references:goalid"`
}

func (TaskActivity) TableName() string {
    return "task_activities"
}

func (TaskActivity) Relationships() []string {
    return []string{
		"Goallists",
		"NpcType",
		"NpcType.AlternateCurrency",
		"NpcType.Loottable",
		"NpcType.Loottable.LoottableEntries",
		"NpcType.Loottable.LoottableEntries.LootdropEntries",
		"NpcType.Loottable.LoottableEntries.LootdropEntries.Item",
		"NpcType.Loottable.LoottableEntries.LootdropEntries.Item.AlternateCurrencies",
		"NpcType.Loottable.LoottableEntries.LootdropEntries.Item.CharacterCorpseItems",
		"NpcType.Loottable.LoottableEntries.LootdropEntries.Item.DiscoveredItems",
		"NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Doors",
		"NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Doors.Item",
		"NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Fishings",
		"NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Fishings.Item",
		"NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Fishings.NpcType",
		"NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Fishings.Zone",
		"NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Forages",
		"NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Forages.Item",
		"NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Forages.Zone",
		"NpcType.Loottable.LoottableEntries.LootdropEntries.Item.GroundSpawns",
		"NpcType.Loottable.LoottableEntries.LootdropEntries.Item.GroundSpawns.Zone",
		"NpcType.Loottable.LoottableEntries.LootdropEntries.Item.ItemTicks",
		"NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Keyrings",
		"NpcType.Loottable.LoottableEntries.LootdropEntries.Item.LootdropEntries",
		"NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Merchantlists",
		"NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType",
		"NpcType.Loottable.LoottableEntries.LootdropEntries.Item.ObjectContents",
		"NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Objects",
		"NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Objects.Item",
		"NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Objects.Zone",
		"NpcType.Loottable.LoottableEntries.LootdropEntries.Item.StartingItems",
		"NpcType.Loottable.LoottableEntries.LootdropEntries.Item.StartingItems.Item",
		"NpcType.Loottable.LoottableEntries.LootdropEntries.Item.StartingItems.Zone",
		"NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Tasks",
		"NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities",
		"NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Tasks.Tasksets",
		"NpcType.Loottable.LoottableEntries.LootdropEntries.Item.TradeskillRecipeEntries",
		"NpcType.Loottable.LoottableEntries.LootdropEntries.Item.TradeskillRecipeEntries.TradeskillRecipe",
		"NpcType.Loottable.LoottableEntries.LootdropEntries.Item.TributeLevels",
		"NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop",
		"NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LootdropEntries",
		"NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LoottableEntries",
		"NpcType.Loottable.LoottableEntries.Loottable",
		"NpcType.Loottable.NpcTypes",
		"NpcType.Merchantlists",
		"NpcType.Merchantlists.NpcType",
		"NpcType.NpcEmotes",
		"NpcType.NpcFactions",
		"NpcType.NpcFactions.NpcFactionEntries",
		"NpcType.NpcSpells",
		"NpcType.NpcSpells.NpcSpellsEntries",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Aura",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Aura.SpellsNew",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.BlockedSpells",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Damageshieldtypes",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.AlternateCurrencies",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.CharacterCorpseItems",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.DiscoveredItems",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.Doors",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.Doors.Item",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.Fishings",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.Fishings.Item",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.Fishings.Zone",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.Forages",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.Forages.Item",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.Forages.Zone",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.GroundSpawns",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.GroundSpawns.Zone",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.ItemTicks",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.Keyrings",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.LootdropEntries",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Item",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LootdropEntries",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.LootdropEntries",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.LoottableEntries",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.Merchantlists",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcType",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.ObjectContents",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.Objects",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.Objects.Item",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.Objects.Zone",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.StartingItems",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.StartingItems.Item",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.StartingItems.Zone",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.Tasks",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.Tasks.TaskActivities",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.Tasks.Tasksets",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.TradeskillRecipeEntries",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.TradeskillRecipeEntries.TradeskillRecipe",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.Items.TributeLevels",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.NpcSpellsEntries",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.SpellBuckets",
		"NpcType.NpcSpells.NpcSpellsEntries.SpellsNew.SpellGlobals",
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
