package models

import (
	"github.com/volatiletech/null/v8"
)

type CharacterActivity struct {
	Charid     uint      `json:"charid" gorm:"Column:charid"`
	Taskid     uint      `json:"taskid" gorm:"Column:taskid"`
	Activityid uint      `json:"activityid" gorm:"Column:activityid"`
	Donecount  uint      `json:"donecount" gorm:"Column:donecount"`
	Completed  null.Int8 `json:"completed" gorm:"Column:completed"`
}

func (CharacterActivity) TableName() string {
    return "character_activities"
}

func (CharacterActivity) Relationships() []string {
    return []string{}
}

func (CharacterActivity) Connection() string {
    return ""
}
