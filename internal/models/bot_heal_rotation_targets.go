package models

type BotHealRotationTarget struct {
	TargetIndex         uint   `json:"target_index" gorm:"Column:target_index"`
	HealRotationIndex   uint   `json:"heal_rotation_index" gorm:"Column:heal_rotation_index"`
	TargetName          string `json:"target_name" gorm:"Column:target_name"`
}

func (BotHealRotationTarget) TableName() string {
    return "bot_heal_rotation_targets"
}

func (BotHealRotationTarget) Relationships() []string {
    return []string{}
}

func (BotHealRotationTarget) Connection() string {
    return ""
}
