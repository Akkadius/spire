package models

type BotCommandSetting struct {
	BotCommand  string `json:"bot_command" gorm:"Column:bot_command"`
	Access      int    `json:"access" gorm:"Column:access"`
	Aliases     string `json:"aliases" gorm:"Column:aliases"`
}

func (BotCommandSetting) TableName() string {
    return "bot_command_settings"
}

func (BotCommandSetting) Relationships() []string {
    return []string{}
}

func (BotCommandSetting) Connection() string {
    return ""
}
