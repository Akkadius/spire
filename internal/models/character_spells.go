package models

type CharacterSpell struct {
	ID       uint   `json:"id" gorm:"Column:id"`
	SlotId   uint16 `json:"slot_id" gorm:"Column:slot_id"`
	SpellId  uint16 `json:"spell_id" gorm:"Column:spell_id"`
}

func (CharacterSpell) TableName() string {
    return "character_spells"
}

func (CharacterSpell) Relationships() []string {
    return []string{}
}

func (CharacterSpell) Connection() string {
    return ""
}
