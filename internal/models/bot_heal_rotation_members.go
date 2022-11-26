package models

type BotHealRotationMember struct {
	MemberIndex         uint `json:"member_index" gorm:"Column:member_index"`
	HealRotationIndex   uint `json:"heal_rotation_index" gorm:"Column:heal_rotation_index"`
	BotId               uint `json:"bot_id" gorm:"Column:bot_id"`
}

func (BotHealRotationMember) TableName() string {
    return "bot_heal_rotation_members"
}

func (BotHealRotationMember) Relationships() []string {
    return []string{}
}

func (BotHealRotationMember) Connection() string {
    return ""
}
