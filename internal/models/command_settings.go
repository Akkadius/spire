package models

type CommandSetting struct {
	Command string `json:"command" gorm:"Column:command"`
	Access  int    `json:"access" gorm:"Column:access"`
	Aliases string `json:"aliases" gorm:"Column:aliases"`
}

func (CommandSetting) TableName() string {
    return "command_settings"
}

func (CommandSetting) Relationships() []string {
    return []string{}
}

func (CommandSetting) Connection() string {
    return ""
}
