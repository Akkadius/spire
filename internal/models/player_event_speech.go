package models

import (
	"github.com/volatiletech/null/v8"
)

type PlayerEventSpeech struct {
	ID           uint64      `json:"id" gorm:"Column:id"`
	ToCharId     null.String `json:"to_char_id" gorm:"Column:to_char_id"`
	FromCharId   null.String `json:"from_char_id" gorm:"Column:from_char_id"`
	GuildId      null.Uint   `json:"guild_id" gorm:"Column:guild_id"`
	Type         null.Uint   `json:"type" gorm:"Column:type"`
	MinStatus    null.Uint   `json:"min_status" gorm:"Column:min_status"`
	Message      null.String `json:"message" gorm:"Column:message"`
	CreatedAt    null.Time   `json:"created_at" gorm:"Column:created_at"`
}

func (PlayerEventSpeech) TableName() string {
    return "player_event_speech"
}

func (PlayerEventSpeech) Relationships() []string {
    return []string{}
}

func (PlayerEventSpeech) Connection() string {
    return ""
}
