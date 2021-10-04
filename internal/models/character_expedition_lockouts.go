package models

import (
	"time"
)

type CharacterExpeditionLockout struct {
	ID                   uint      `json:"id" gorm:"Column:id"`
	CharacterId          uint      `json:"character_id" gorm:"Column:character_id"`
	ExpeditionName       string    `json:"expedition_name" gorm:"Column:expedition_name"`
	EventName            string    `json:"event_name" gorm:"Column:event_name"`
	ExpireTime           time.Time `json:"expire_time" gorm:"Column:expire_time"`
	Duration             uint      `json:"duration" gorm:"Column:duration"`
	FromExpeditionUuid   string    `json:"from_expedition_uuid" gorm:"Column:from_expedition_uuid"`
}

func (CharacterExpeditionLockout) TableName() string {
    return "character_expedition_lockouts"
}

func (CharacterExpeditionLockout) Relationships() []string {
    return []string{}
}

func (CharacterExpeditionLockout) Connection() string {
    return ""
}
