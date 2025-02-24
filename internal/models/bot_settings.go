package models

import (
	"github.com/volatiletech/null/v8"
)

type BotSetting struct {
	CharacterId   uint        `json:"character_id" gorm:"Column:character_id"`
	BotId         uint        `json:"bot_id" gorm:"Column:bot_id"`
	Stance        uint8       `json:"stance" gorm:"Column:stance"`
	SettingId     uint16      `json:"setting_id" gorm:"Column:setting_id"`
	SettingType   uint8       `json:"setting_type" gorm:"Column:setting_type"`
	Value         int64       `json:"value" gorm:"Column:value"`
	CategoryName  null.String `json:"category_name" gorm:"Column:category_name"`
	SettingName   null.String `json:"setting_name" gorm:"Column:setting_name"`
}

func (BotSetting) TableName() string {
    return "bot_settings"
}

func (BotSetting) Relationships() []string {
    return []string{}
}

func (BotSetting) Connection() string {
    return ""
}
