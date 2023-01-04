package models

type BotSpellSetting struct {
	ID         uint  `json:"id" gorm:"Column:id"`
	BotId      int   `json:"bot_id" gorm:"Column:bot_id"`
	SpellId    int16 `json:"spell_id" gorm:"Column:spell_id"`
	Priority   int16 `json:"priority" gorm:"Column:priority"`
	MinHp      int16 `json:"min_hp" gorm:"Column:min_hp"`
	MaxHp      int16 `json:"max_hp" gorm:"Column:max_hp"`
	IsEnabled  uint8 `json:"is_enabled" gorm:"Column:is_enabled"`
}

func (BotSpellSetting) TableName() string {
    return "bot_spell_settings"
}

func (BotSpellSetting) Relationships() []string {
    return []string{}
}

func (BotSpellSetting) Connection() string {
    return ""
}
