package models

import (
	"github.com/volatiletech/null/v8"
)

type SpawnConditionValue struct {
	ID          uint       `json:"id" gorm:"Column:id"`
	Value       null.Uint8 `json:"value" gorm:"Column:value"`
	Zone        string     `json:"zone" gorm:"Column:zone"`
	InstanceId  uint       `json:"instance_id" gorm:"Column:instance_id"`
}

func (SpawnConditionValue) TableName() string {
    return "spawn_condition_values"
}

func (SpawnConditionValue) Relationships() []string {
    return []string{}
}

func (SpawnConditionValue) Connection() string {
    return ""
}
