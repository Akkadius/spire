package models

type GroupId struct {
	GroupId      uint   `json:"group_id" gorm:"Column:group_id"`
	Name         string `json:"name" gorm:"Column:name"`
	CharacterId  uint   `json:"character_id" gorm:"Column:character_id"`
	BotId        uint   `json:"bot_id" gorm:"Column:bot_id"`
	MercId       uint   `json:"merc_id" gorm:"Column:merc_id"`
}

func (GroupId) TableName() string {
    return "group_id"
}

func (GroupId) Relationships() []string {
    return []string{}
}

func (GroupId) Connection() string {
    return ""
}
