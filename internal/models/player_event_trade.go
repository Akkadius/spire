package models

import (
	"github.com/volatiletech/null/v8"
)

type PlayerEventTrade struct {
	ID             uint        `json:"id" gorm:"Column:id"`
	Char1Id        null.Uint   `json:"char_1_id" gorm:"Column:char1_id"`
	Char2Id        null.Uint   `json:"char_2_id" gorm:"Column:char2_id"`
	Char1Copper    null.Uint64 `json:"char_1_copper" gorm:"Column:char1_copper"`
	Char1Silver    null.Uint64 `json:"char_1_silver" gorm:"Column:char1_silver"`
	Char1Gold      null.Uint64 `json:"char_1_gold" gorm:"Column:char1_gold"`
	Char1Platinum  null.Uint64 `json:"char_1_platinum" gorm:"Column:char1_platinum"`
	Char2Copper    null.Uint64 `json:"char_2_copper" gorm:"Column:char2_copper"`
	Char2Silver    null.Uint64 `json:"char_2_silver" gorm:"Column:char2_silver"`
	Char2Gold      null.Uint64 `json:"char_2_gold" gorm:"Column:char2_gold"`
	Char2Platinum  null.Uint64 `json:"char_2_platinum" gorm:"Column:char2_platinum"`
	CreatedAt      null.Time   `json:"created_at" gorm:"Column:created_at"`
}

func (PlayerEventTrade) TableName() string {
    return "player_event_trade"
}

func (PlayerEventTrade) Relationships() []string {
    return []string{}
}

func (PlayerEventTrade) Connection() string {
    return ""
}
