package models

type CharacterAura struct {
	ID       int  `json:"id" gorm:"Column:id"`
	Slot     int8 `json:"slot" gorm:"Column:slot"`
	SpellId  int  `json:"spell_id" gorm:"Column:spell_id"`
}

func (CharacterAura) TableName() string {
    return "character_auras"
}

func (CharacterAura) Relationships() []string {
    return []string{}
}

func (CharacterAura) Connection() string {
    return ""
}
