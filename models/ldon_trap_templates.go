package models

type LdonTrapTemplate struct {
	ID       uint   `json:"id" gorm:"Column:id"`
	Type     uint8  `json:"type" gorm:"Column:type"`
	SpellId  uint16 `json:"spell_id" gorm:"Column:spell_id"`
	Skill    uint16 `json:"skill" gorm:"Column:skill"`
	Locked   uint8  `json:"locked" gorm:"Column:locked"`
}

func (LdonTrapTemplate) TableName() string {
    return "ldon_trap_templates"
}

func (LdonTrapTemplate) Relationships() []string {
    return []string{}
}

func (LdonTrapTemplate) Connection() string {
    return "eqemu_content"
}
