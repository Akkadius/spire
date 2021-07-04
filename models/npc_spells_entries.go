package models

import (
	"github.com/volatiletech/null/v8"
)

type NpcSpellsEntry struct {
	ID            uint       `json:"id" gorm:"Column:id"`
	NpcSpellsId   int        `json:"npc_spells_id" gorm:"Column:npc_spells_id"`
	Spellid       uint16     `json:"spellid" gorm:"Column:spellid"`
	Type          uint       `json:"type" gorm:"Column:type"`
	Minlevel      uint8      `json:"minlevel" gorm:"Column:minlevel"`
	Maxlevel      uint8      `json:"maxlevel" gorm:"Column:maxlevel"`
	Manacost      int16      `json:"manacost" gorm:"Column:manacost"`
	RecastDelay   int        `json:"recast_delay" gorm:"Column:recast_delay"`
	Priority      int16      `json:"priority" gorm:"Column:priority"`
	ResistAdjust  null.Int   `json:"resist_adjust" gorm:"Column:resist_adjust"`
	MinHp         null.Int16 `json:"min_hp" gorm:"Column:min_hp"`
	MaxHp         null.Int16 `json:"max_hp" gorm:"Column:max_hp"`
}

func (NpcSpellsEntry) TableName() string {
    return "npc_spells_entries"
}

func (NpcSpellsEntry) Relationships() []string {
    return []string{}
}

func (NpcSpellsEntry) Connection() string {
    return "eqemu_content"
}
