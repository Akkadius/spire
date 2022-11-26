package models

type BotGroup struct {
	GroupsIndex     uint   `json:"groups_index" gorm:"Column:groups_index"`
	GroupLeaderId   uint   `json:"group_leader_id" gorm:"Column:group_leader_id"`
	GroupName       string `json:"group_name" gorm:"Column:group_name"`
	AutoSpawn       uint8  `json:"auto_spawn" gorm:"Column:auto_spawn"`
}

func (BotGroup) TableName() string {
    return "bot_groups"
}

func (BotGroup) Relationships() []string {
    return []string{}
}

func (BotGroup) Connection() string {
    return ""
}
