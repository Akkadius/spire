package models

import (
	"github.com/volatiletech/null/v8"
)

type PlayerEventAaPurchase struct {
	ID            uint64    `json:"id" gorm:"Column:id"`
	AaAbilityId   null.Int  `json:"aa_ability_id" gorm:"Column:aa_ability_id"`
	Cost          null.Int  `json:"cost" gorm:"Column:cost"`
	PreviousId    null.Int  `json:"previous_id" gorm:"Column:previous_id"`
	NextId        null.Int  `json:"next_id" gorm:"Column:next_id"`
	CreatedAt     null.Time `json:"created_at" gorm:"Column:created_at"`
}

func (PlayerEventAaPurchase) TableName() string {
    return "player_event_aa_purchase"
}

func (PlayerEventAaPurchase) Relationships() []string {
    return []string{}
}

func (PlayerEventAaPurchase) Connection() string {
    return ""
}
