package models

type BotGroupMember struct {
	GroupMembersIndex   uint `json:"group_members_index" gorm:"Column:group_members_index"`
	GroupsIndex         uint `json:"groups_index" gorm:"Column:groups_index"`
	BotId               uint `json:"bot_id" gorm:"Column:bot_id"`
}

func (BotGroupMember) TableName() string {
    return "bot_group_members"
}

func (BotGroupMember) Relationships() []string {
    return []string{}
}

func (BotGroupMember) Connection() string {
    return ""
}
