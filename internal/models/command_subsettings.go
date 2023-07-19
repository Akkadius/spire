package models

type CommandSubsetting struct {
	ID                uint   `json:"id" gorm:"Column:id"`
	ParentCommand     string `json:"parent_command" gorm:"Column:parent_command"`
	SubCommand        string `json:"sub_command" gorm:"Column:sub_command"`
	AccessLevel       uint   `json:"access_level" gorm:"Column:access_level"`
	TopLevelAliases   string `json:"top_level_aliases" gorm:"Column:top_level_aliases"`
}

func (CommandSubsetting) TableName() string {
    return "command_subsettings"
}

func (CommandSubsetting) Relationships() []string {
    return []string{}
}

func (CommandSubsetting) Connection() string {
    return ""
}
