package models

import (
	"github.com/volatiletech/null/v8"
)

type InstanceList struct {
	ID                  int                  `json:"id" gorm:"Column:id"`
	Zone                uint                 `json:"zone" gorm:"Column:zone"`
	Version             uint8                `json:"version" gorm:"Column:version"`
	IsGlobal            uint8                `json:"is_global" gorm:"Column:is_global"`
	StartTime           uint                 `json:"start_time" gorm:"Column:start_time"`
	Duration            uint                 `json:"duration" gorm:"Column:duration"`
	ExpireAt            uint64               `json:"expire_at" gorm:"Column:expire_at"`
	NeverExpires        uint8                `json:"never_expires" gorm:"Column:never_expires"`
	Notes               null.String          `json:"notes" gorm:"Column:notes"`
	InstanceListPlayers []InstanceListPlayer `json:"instance_list_players,omitempty" gorm:"foreignKey:id;references:id"`
	Zones               []Zone               `json:"zones,omitempty" gorm:"foreignKey:zoneidunumber;references:zone"`
}

func (InstanceList) TableName() string {
    return "instance_list"
}

func (InstanceList) Relationships() []string {
    return []string{
		"InstanceListPlayers",
		"Zones",
	}
}

func (InstanceList) Connection() string {
    return ""
}
