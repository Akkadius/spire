package models

type CharacterAlternateAbility struct {
	ID       uint   `json:"id" gorm:"Column:id"`
	AaId     uint16 `json:"aa_id" gorm:"Column:aa_id"`
	AaValue  uint16 `json:"aa_value" gorm:"Column:aa_value"`
	Charges  uint16 `json:"charges" gorm:"Column:charges"`
}

func (CharacterAlternateAbility) TableName() string {
    return "character_alternate_abilities"
}

func (CharacterAlternateAbility) Relationships() []string {
    return []string{}
}

func (CharacterAlternateAbility) Connection() string {
    return ""
}
