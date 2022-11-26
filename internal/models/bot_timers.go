package models

type BotTimer struct {
	BotId       uint `json:"bot_id" gorm:"Column:bot_id"`
	TimerId     uint `json:"timer_id" gorm:"Column:timer_id"`
	TimerValue  uint `json:"timer_value" gorm:"Column:timer_value"`
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
