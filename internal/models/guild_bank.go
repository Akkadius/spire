package models

import (
	"github.com/volatiletech/null/v8"
)

type GuildBank struct {
	ID          uint        `json:"id" gorm:"Column:id"`
	Guildid     uint        `json:"guildid" gorm:"Column:guildid"`
	Area        uint8       `json:"area" gorm:"Column:area"`
	Slot        uint        `json:"slot" gorm:"Column:slot"`
	Itemid      uint        `json:"itemid" gorm:"Column:itemid"`
	Qty         int         `json:"qty" gorm:"Column:qty"`
	Donator     null.String `json:"donator" gorm:"Column:donator"`
	Permissions uint8       `json:"permissions" gorm:"Column:permissions"`
	Whofor      null.String `json:"whofor" gorm:"Column:whofor"`
}

func (GuildBank) TableName() string {
    return "guild_bank"
}

func (GuildBank) Relationships() []string {
    return []string{}
}

func (GuildBank) Connection() string {
    return ""
}
