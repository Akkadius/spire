package models

type CharacterDiscipline struct {
	ID      uint   `json:"id" gorm:"Column:id"`
	SlotId  uint16 `json:"slot_id" gorm:"Column:slot_id"`
	DiscId  uint16 `json:"disc_id" gorm:"Column:disc_id"`
}

func (CharacterDiscipline) TableName() string {
    return "character_disciplines"
}

func (CharacterDiscipline) Relationships() []string {
    return []string{}
}

func (CharacterDiscipline) Connection() string {
    return ""
}
