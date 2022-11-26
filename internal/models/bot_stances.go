package models

type BotStance struct {
	BotId     uint  `json:"bot_id" gorm:"Column:bot_id"`
	StanceId  uint8 `json:"stance_id" gorm:"Column:stance_id"`
}

func (BotStance) TableName() string {
    return "bot_stances"
}

func (BotStance) Relationships() []string {
    return []string{}
}

func (BotStance) Connection() string {
    return ""
}
