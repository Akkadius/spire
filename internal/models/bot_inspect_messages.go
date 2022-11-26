package models

type BotInspectMessage struct {
	BotId           uint   `json:"bot_id" gorm:"Column:bot_id"`
	InspectMessage  string `json:"inspect_message" gorm:"Column:inspect_message"`
}

func (BotInspectMessage) TableName() string {
    return "bot_inspect_messages"
}

func (BotInspectMessage) Relationships() []string {
    return []string{}
}

func (BotInspectMessage) Connection() string {
    return ""
}
