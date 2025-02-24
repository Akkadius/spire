package models

import (
	"github.com/volatiletech/null/v8"
)

type BotBlockedBuff struct {
	BotId       uint       `json:"bot_id" gorm:"Column:bot_id"`
	SpellId     uint       `json:"spell_id" gorm:"Column:spell_id"`
	Blocked     null.Uint8 `json:"blocked" gorm:"Column:blocked"`
	BlockedPet  null.Uint8 `json:"blocked_pet" gorm:"Column:blocked_pet"`
}

func (BotBlockedBuff) TableName() string {
    return "bot_blocked_buffs"
}

func (BotBlockedBuff) Relationships() []string {
    return []string{}
}

func (BotBlockedBuff) Connection() string {
    return ""
}
