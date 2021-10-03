package models

type CharacterInspectMessage struct {
	ID              uint   `json:"id" gorm:"Column:id"`
	InspectMessage  string `json:"inspect_message" gorm:"Column:inspect_message"`
}

func (CharacterInspectMessage) TableName() string {
    return "character_inspect_messages"
}

func (CharacterInspectMessage) Relationships() []string {
    return []string{}
}

func (CharacterInspectMessage) Connection() string {
    return ""
}
