package models

import (
	"github.com/volatiletech/null/v8"
)

type NpcScaleGlobalBase struct {
	Type              int         `json:"type" gorm:"Column:type"`
	Level             int         `json:"level" gorm:"Column:level"`
	Ac                null.Int    `json:"ac" gorm:"Column:ac"`
	Hp                null.Int    `json:"hp" gorm:"Column:hp"`
	Accuracy          null.Int    `json:"accuracy" gorm:"Column:accuracy"`
	SlowMitigation    null.Int    `json:"slow_mitigation" gorm:"Column:slow_mitigation"`
	Attack            null.Int    `json:"attack" gorm:"Column:attack"`
	Strength          null.Int    `json:"strength" gorm:"Column:strength"`
	Stamina           null.Int    `json:"stamina" gorm:"Column:stamina"`
	Dexterity         null.Int    `json:"dexterity" gorm:"Column:dexterity"`
	Agility           null.Int    `json:"agility" gorm:"Column:agility"`
	Intelligence      null.Int    `json:"intelligence" gorm:"Column:intelligence"`
	Wisdom            null.Int    `json:"wisdom" gorm:"Column:wisdom"`
	Charisma          null.Int    `json:"charisma" gorm:"Column:charisma"`
	MagicResist       null.Int    `json:"magic_resist" gorm:"Column:magic_resist"`
	ColdResist        null.Int    `json:"cold_resist" gorm:"Column:cold_resist"`
	FireResist        null.Int    `json:"fire_resist" gorm:"Column:fire_resist"`
	PoisonResist      null.Int    `json:"poison_resist" gorm:"Column:poison_resist"`
	DiseaseResist     null.Int    `json:"disease_resist" gorm:"Column:disease_resist"`
	CorruptionResist  null.Int    `json:"corruption_resist" gorm:"Column:corruption_resist"`
	PhysicalResist    null.Int    `json:"physical_resist" gorm:"Column:physical_resist"`
	MinDmg            null.Int    `json:"min_dmg" gorm:"Column:min_dmg"`
	MaxDmg            null.Int    `json:"max_dmg" gorm:"Column:max_dmg"`
	HpRegenRate       null.Int    `json:"hp_regen_rate" gorm:"Column:hp_regen_rate"`
	AttackDelay       null.Int    `json:"attack_delay" gorm:"Column:attack_delay"`
	SpellScale        null.Int    `json:"spell_scale" gorm:"Column:spell_scale"`
	HealScale         null.Int    `json:"heal_scale" gorm:"Column:heal_scale"`
	SpecialAbilities  null.String `json:"special_abilities" gorm:"Column:special_abilities"`
}

func (NpcScaleGlobalBase) TableName() string {
    return "npc_scale_global_base"
}

func (NpcScaleGlobalBase) Relationships() []string {
    return []string{}
}

func (NpcScaleGlobalBase) Connection() string {
    return "eqemu_content"
}
