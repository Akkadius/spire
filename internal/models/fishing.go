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
}

func (Fishing) TableName() string {
    return "fishing"
}

func (Fishing) Relationships() []string {
    return []string{}
}

func (Fishing) Connection() string {
    return "eqemu_content"
}
