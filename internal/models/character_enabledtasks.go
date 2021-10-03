package models

type CharacterEnabledtask struct {
	Charid uint `json:"charid" gorm:"Column:charid"`
	Taskid uint `json:"taskid" gorm:"Column:taskid"`
}

func (CharacterEnabledtask) TableName() string {
    return "character_enabledtasks"
}

func (CharacterEnabledtask) Relationships() []string {
    return []string{}
}

func (CharacterEnabledtask) Connection() string {
    return ""
}
