package models

import (
	"github.com/volatiletech/null/v8"
)

type Spawn2Disabled struct {
	ID          int64      `json:"id" gorm:"Column:id"`
	Spawn2Id    null.Int   `json:"spawn_2_id" gorm:"Column:spawn2_id"`
	InstanceId  null.Int   `json:"instance_id" gorm:"Column:instance_id"`
	Disabled    null.Int16 `json:"disabled" gorm:"Column:disabled"`
}

func (Spawn2Disabled) TableName() string {
    return "spawn2_disabled"
}

func (Spawn2Disabled) Relationships() []string {
    return []string{}
}

func (Spawn2Disabled) Connection() string {
    return ""
}
