package models

import (
	"time"
)

type ExpeditionLockout struct {
	ID                   uint      `json:"id" gorm:"Column:id"`
	ExpeditionId         uint      `json:"expedition_id" gorm:"Column:expedition_id"`
	EventName            string    `json:"event_name" gorm:"Column:event_name"`
	ExpireTime           time.Time `json:"expire_time" gorm:"Column:expire_time"`
	Duration             uint      `json:"duration" gorm:"Column:duration"`
	FromExpeditionUuid   string    `json:"from_expedition_uuid" gorm:"Column:from_expedition_uuid"`
}

func (ExpeditionLockout) TableName() string {
    return "expedition_lockouts"
}

func (ExpeditionLockout) Relationships() []string {
    return []string{}
}

func (ExpeditionLockout) Connection() string {
    return ""
}
