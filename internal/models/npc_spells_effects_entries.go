package models

type NpcSpellsEffectsEntry struct {
	ID                    uint  `json:"id" gorm:"Column:id"`
	NpcSpellsEffectsId    int   `json:"npc_spells_effects_id" gorm:"Column:npc_spells_effects_id"`
	SpellEffectId         int16 `json:"spell_effect_id" gorm:"Column:spell_effect_id"`
	Minlevel              uint8 `json:"minlevel" gorm:"Column:minlevel"`
	Maxlevel              uint8 `json:"maxlevel" gorm:"Column:maxlevel"`
	SeBase                int   `json:"se_base" gorm:"Column:se_base"`
	SeLimit               int   `json:"se_limit" gorm:"Column:se_limit"`
	SeMax                 int   `json:"se_max" gorm:"Column:se_max"`
}

func (NpcSpellsEffectsEntry) TableName() string {
    return "npc_spells_effects_entries"
}

func (NpcSpellsEffectsEntry) Relationships() []string {
    return []string{}
}

func (NpcSpellsEffectsEntry) Connection() string {
    return "eqemu_content"
}
