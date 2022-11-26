package models

type BotHealRotation struct {
	HealRotationIndex   uint   `json:"heal_rotation_index" gorm:"Column:heal_rotation_index"`
	BotId               uint   `json:"bot_id" gorm:"Column:bot_id"`
	Interval            uint   `json:"interval" gorm:"Column:interval"`
	FastHeals           uint   `json:"fast_heals" gorm:"Column:fast_heals"`
	AdaptiveTargeting   uint   `json:"adaptive_targeting" gorm:"Column:adaptive_targeting"`
	CastingOverride     uint   `json:"casting_override" gorm:"Column:casting_override"`
	SafeHpBase          string `json:"safe_hp_base" gorm:"Column:safe_hp_base"`
	SafeHpCloth         string `json:"safe_hp_cloth" gorm:"Column:safe_hp_cloth"`
	SafeHpLeather       string `json:"safe_hp_leather" gorm:"Column:safe_hp_leather"`
	SafeHpChain         string `json:"safe_hp_chain" gorm:"Column:safe_hp_chain"`
	SafeHpPlate         string `json:"safe_hp_plate" gorm:"Column:safe_hp_plate"`
	CriticalHpBase      string `json:"critical_hp_base" gorm:"Column:critical_hp_base"`
	CriticalHpCloth     string `json:"critical_hp_cloth" gorm:"Column:critical_hp_cloth"`
	CriticalHpLeather   string `json:"critical_hp_leather" gorm:"Column:critical_hp_leather"`
	CriticalHpChain     string `json:"critical_hp_chain" gorm:"Column:critical_hp_chain"`
	CriticalHpPlate     string `json:"critical_hp_plate" gorm:"Column:critical_hp_plate"`
}

func (BotHealRotation) TableName() string {
    return "bot_heal_rotations"
}

func (BotHealRotation) Relationships() []string {
    return []string{}
}

func (BotHealRotation) Connection() string {
    return ""
}
