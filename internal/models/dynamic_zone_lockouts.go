package models

import (
	"time"
)

type DynamicZoneLockout struct {
	ID                   uint      `json:"id" gorm:"Column:id"`
	DynamicZoneId        uint      `json:"dynamic_zone_id" gorm:"Column:dynamic_zone_id"`
	EventName            string    `json:"event_name" gorm:"Column:event_name"`
	ExpireTime           time.Time `json:"expire_time" gorm:"Column:expire_time"`
	Duration             uint      `json:"duration" gorm:"Column:duration"`
	FromExpeditionUuid   string    `json:"from_expedition_uuid" gorm:"Column:from_expedition_uuid"`
}

func (DynamicZoneLockout) TableName() string {
    return "dynamic_zone_lockouts"
}

func (DynamicZoneLockout) Relationships() []string {
    return []string{}
}

func (DynamicZoneLockout) Connection() string {
    return ""
}
