package models

import (
	"github.com/volatiletech/null/v8"
)

type NpcSpellsEffect struct {
	ID                      uint                    `json:"id" gorm:"Column:id"`
	Name                    null.String             `json:"name" gorm:"Column:name"`
	ParentList              uint                    `json:"parent_list" gorm:"Column:parent_list"`
	NpcSpellsEffectsEntries []NpcSpellsEffectsEntry `json:"npc_spells_effects_entries,omitempty" gorm:"foreignKey:npc_spells_effects_id;references:id"`
}

func (NpcSpellsEffect) TableName() string {
    return "npc_spells_effects"
}

func (NpcSpellsEffect) Relationships() []string {
    return []string{
		"NpcSpellsEffectsEntries",
	}
}

func (NpcSpellsEffect) Connection() string {
    return "eqemu_content"
}
