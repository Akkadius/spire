package models

import (
	"github.com/volatiletech/null/v8"
)

type RespawnTime struct {
	ID          int       `json:"id" gorm:"Column:id"`
	Start       int       `json:"start" gorm:"Column:start"`
	Duration    int       `json:"duration" gorm:"Column:duration"`
	ExpireAt    null.Uint `json:"expire_at" gorm:"Column:expire_at"`
	InstanceId  int16     `json:"instance_id" gorm:"Column:instance_id"`
}

func (RespawnTime) TableName() string {
    return "respawn_times"
}

func (RespawnTime) Relationships() []string {
    return []string{}
}

func (RespawnTime) Connection() string {
    return ""
}
