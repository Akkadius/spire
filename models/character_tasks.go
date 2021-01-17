package models

import (
	"github.com/volatiletech/null/v8"
)

type CharacterTask struct {
	Charid       uint      `json:"charid" gorm:"Column:charid"`
	Taskid       uint      `json:"taskid" gorm:"Column:taskid"`
	Slot         uint      `json:"slot" gorm:"Column:slot"`
	Type         int8      `json:"type" gorm:"Column:type"`
	Acceptedtime null.Uint `json:"acceptedtime" gorm:"Column:acceptedtime"`
}

func (CharacterTask) TableName() string {
    return "character_tasks"
}

func (CharacterTask) Relationships() []string {
    return []string{}
}

func (CharacterTask) Connection() string {
    return ""
}
