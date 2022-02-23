package models

type Aura struct {
	Type       int        `json:"type" gorm:"Column:type"`
	NpcType    int        `json:"npc_type" gorm:"Column:npc_type"`
	Name       string     `json:"name" gorm:"Column:name"`
	SpellId    int        `json:"spell_id" gorm:"Column:spell_id"`
	Distance   int        `json:"distance" gorm:"Column:distance"`
	AuraType   int        `json:"aura_type" gorm:"Column:aura_type"`
	SpawnType  int        `json:"spawn_type" gorm:"Column:spawn_type"`
	Movement   int        `json:"movement" gorm:"Column:movement"`
	Duration   int        `json:"duration" gorm:"Column:duration"`
	Icon       int        `json:"icon" gorm:"Column:icon"`
	CastTime   int        `json:"cast_time" gorm:"Column:cast_time"`
	SpellsNew  *SpellsNew `json:"spells_new,omitempty" gorm:"foreignKey:spell_id;references:id"`
}

func (Aura) TableName() string {
    return "auras"
}

func (Aura) Relationships() []string {
    return []string{
		"SpellsNew",
		"SpellsNew.Aura",
		"SpellsNew.BlockedSpells",
		"SpellsNew.Damageshieldtypes",
		"SpellsNew.SpellBuckets",
		"SpellsNew.SpellGlobals",
	}
}

func (Aura) Connection() string {
    return "eqemu_content"
}
