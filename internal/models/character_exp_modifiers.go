package models

type CharacterExpModifier struct {
	CharacterId      int     `json:"character_id" gorm:"Column:character_id"`
	ZoneId           int     `json:"zone_id" gorm:"Column:zone_id"`
	InstanceVersion  int     `json:"instance_version" gorm:"Column:instance_version"`
	AaModifier       float32 `json:"aa_modifier" gorm:"Column:aa_modifier"`
	ExpModifier      float32 `json:"exp_modifier" gorm:"Column:exp_modifier"`
}

func (CharacterExpModifier) TableName() string {
    return "character_exp_modifiers"
}

func (CharacterExpModifier) Relationships() []string {
    return []string{}
}

func (CharacterExpModifier) Connection() string {
    return ""
}
