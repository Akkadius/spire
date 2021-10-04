package models

import (
	"github.com/volatiletech/null/v8"
)

type QuestGlobal struct {
	Charid  int      `json:"charid" gorm:"Column:charid"`
	Npcid   int      `json:"npcid" gorm:"Column:npcid"`
	Zoneid  int      `json:"zoneid" gorm:"Column:zoneid"`
	Name    string   `json:"name" gorm:"Column:name"`
	Value   string   `json:"value" gorm:"Column:value"`
	Expdate null.Int `json:"expdate" gorm:"Column:expdate"`
}

func (QuestGlobal) TableName() string {
    return "quest_globals"
}

func (QuestGlobal) Relationships() []string {
    return []string{}
}

func (QuestGlobal) Connection() string {
    return ""
}
