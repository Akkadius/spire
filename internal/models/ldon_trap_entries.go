package models

type LdonTrapEntry struct {
	ID      uint `json:"id" gorm:"Column:id"`
	TrapId  uint `json:"trap_id" gorm:"Column:trap_id"`
}

func (LdonTrapEntry) TableName() string {
    return "ldon_trap_entries"
}

func (LdonTrapEntry) Relationships() []string {
    return []string{}
}

func (LdonTrapEntry) Connection() string {
    return "eqemu_content"
}
