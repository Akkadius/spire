package models

import (
	"github.com/volatiletech/null/v8"
)

type BotGuildMember struct {
	BotId          int         `json:"bot_id" gorm:"Column:bot_id"`
	GuildId        uint32      `json:"guild_id" gorm:"Column:guild_id"`
	Rank           uint8       `json:"rank" gorm:"Column:rank"`
	TributeEnable  uint8       `json:"tribute_enable" gorm:"Column:tribute_enable"`
	TotalTribute   uint        `json:"total_tribute" gorm:"Column:total_tribute"`
	LastTribute    uint        `json:"last_tribute" gorm:"Column:last_tribute"`
	Banker         uint8       `json:"banker" gorm:"Column:banker"`
	PublicNote     null.String `json:"public_note" gorm:"Column:public_note"`
	Alt            uint8       `json:"alt" gorm:"Column:alt"`
}

func (BotGuildMember) TableName() string {
    return "bot_guild_members"
}

func (BotGuildMember) Relationships() []string {
    return []string{}
}

func (BotGuildMember) Connection() string {
    return ""
}
