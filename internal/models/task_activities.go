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
		"NpcType.NpcEmotes",
		"NpcType.NpcFactions",
		"NpcType.NpcFactions.NpcFactionEntries",
		"NpcType.NpcSpells",
		"NpcType.NpcSpells.NpcSpellsEntries",
		"NpcType.NpcTypesTint",
		"NpcType.Spawnentries",
		"NpcType.Spawnentries.Spawngroup",
		"NpcType.Spawnentries.Spawngroup.Spawn2",
	}
}

func (TaskActivity) Connection() string {
    return "eqemu_content"
}
