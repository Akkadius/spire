package models

import (
	"github.com/volatiletech/null/v8"
)

type PlayerEventLog struct {
	ID              int64        `json:"id" gorm:"Column:id"`
	AccountId       null.Int64   `json:"account_id" gorm:"Column:account_id"`
	CharacterId     null.Int64   `json:"character_id" gorm:"Column:character_id"`
	ZoneId          null.Int     `json:"zone_id" gorm:"Column:zone_id"`
	InstanceId      null.Int     `json:"instance_id" gorm:"Column:instance_id"`
	X               null.Float32 `json:"x" gorm:"Column:x"`
	Y               null.Float32 `json:"y" gorm:"Column:y"`
	Z               null.Float32 `json:"z" gorm:"Column:z"`
	Heading         null.Float32 `json:"heading" gorm:"Column:heading"`
	EventTypeId     null.Int     `json:"event_type_id" gorm:"Column:event_type_id"`
	EventTypeName   null.String  `json:"event_type_name" gorm:"Column:event_type_name"`
	EventData       null.String  `json:"event_data" gorm:"Column:event_data"`
	CreatedAt       null.Time    `json:"created_at" gorm:"Column:created_at"`
}

func (PlayerEventLog) TableName() string {
    return "player_event_logs"
}

func (PlayerEventLog) Relationships() []string {
    return []string{}
}

func (PlayerEventLog) Connection() string {
    return ""
}
