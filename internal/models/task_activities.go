package models

import (
	"github.com/volatiletech/null/v8"
)

type TaskActivity struct {
	Taskid               uint        `json:"taskid" gorm:"Column:taskid"`
	Activityid           uint        `json:"activityid" gorm:"Column:activityid"`
	ReqActivityId        int         `json:"req_activity_id" gorm:"Column:req_activity_id"`
	Step                 int         `json:"step" gorm:"Column:step"`
	Activitytype         uint8       `json:"activitytype" gorm:"Column:activitytype"`
	TargetName           string      `json:"target_name" gorm:"Column:target_name"`
	Goalmethod           uint        `json:"goalmethod" gorm:"Column:goalmethod"`
	Goalcount            null.Int    `json:"goalcount" gorm:"Column:goalcount"`
	DescriptionOverride  string      `json:"description_override" gorm:"Column:description_override"`
	NpcMatchList         null.String `json:"npc_match_list" gorm:"Column:npc_match_list"`
	ItemIdList           null.String `json:"item_id_list" gorm:"Column:item_id_list"`
	ItemList             string      `json:"item_list" gorm:"Column:item_list"`
	DzSwitchId           int         `json:"dz_switch_id" gorm:"Column:dz_switch_id"`
	MinX                 float32     `json:"min_x" gorm:"Column:min_x"`
	MinY                 float32     `json:"min_y" gorm:"Column:min_y"`
	MinZ                 float32     `json:"min_z" gorm:"Column:min_z"`
	MaxX                 float32     `json:"max_x" gorm:"Column:max_x"`
	MaxY                 float32     `json:"max_y" gorm:"Column:max_y"`
	MaxZ                 float32     `json:"max_z" gorm:"Column:max_z"`
	SkillList            string      `json:"skill_list" gorm:"Column:skill_list"`
	SpellList            string      `json:"spell_list" gorm:"Column:spell_list"`
	Zones                string      `json:"zones" gorm:"Column:zones"`
	ZoneVersion          null.Int    `json:"zone_version" gorm:"Column:zone_version"`
	Optional             int8        `json:"optional" gorm:"Column:optional"`
	ListGroup            uint8       `json:"list_group" gorm:"Column:list_group"`
}

func (TaskActivity) TableName() string {
    return "task_activities"
}

func (TaskActivity) Relationships() []string {
    return []string{}
}

func (TaskActivity) Connection() string {
    return "eqemu_content"
}
