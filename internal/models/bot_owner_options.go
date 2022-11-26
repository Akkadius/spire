package models

import (
	"github.com/volatiletech/null/v8"
)

type BotOwnerOption struct {
	OwnerId      uint        `json:"owner_id" gorm:"Column:owner_id"`
	OptionType   uint16      `json:"option_type" gorm:"Column:option_type"`
	OptionValue  null.Uint16 `json:"option_value" gorm:"Column:option_value"`
}

func (BotOwnerOption) TableName() string {
    return "bot_owner_options"
}

func (BotOwnerOption) Relationships() []string {
    return []string{}
}

func (BotOwnerOption) Connection() string {
    return ""
}
