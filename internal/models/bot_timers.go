package models

type BotTimer struct {
	BotId       uint  `json:"bot_id" gorm:"Column:bot_id"`
	TimerId     uint  `json:"timer_id" gorm:"Column:timer_id"`
	TimerValue  uint  `json:"timer_value" gorm:"Column:timer_value"`
	RecastTime  uint  `json:"recast_time" gorm:"Column:recast_time"`
	IsSpell     uint8 `json:"is_spell" gorm:"Column:is_spell"`
	IsDisc      uint8 `json:"is_disc" gorm:"Column:is_disc"`
	SpellId     uint  `json:"spell_id" gorm:"Column:spell_id"`
	IsItem      uint8 `json:"is_item" gorm:"Column:is_item"`
	ItemId      uint  `json:"item_id" gorm:"Column:item_id"`
}

func (BotTimer) TableName() string {
    return "bot_timers"
}

func (BotTimer) Relationships() []string {
    return []string{}
}

func (BotTimer) Connection() string {
    return ""
}
