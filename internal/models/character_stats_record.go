package models

import (
	"github.com/volatiletech/null/v8"
)

type CharacterStatsRecord struct {
	CharacterId              int         `json:"character_id" gorm:"Column:character_id"`
	Name                     null.String `json:"name" gorm:"Column:name"`
	Status                   null.Int    `json:"status" gorm:"Column:status"`
	Level                    null.Int    `json:"level" gorm:"Column:level"`
	Class                    null.Int    `json:"class" gorm:"Column:class"`
	Race                     null.Int    `json:"race" gorm:"Column:race"`
	AaPoints                 null.Int    `json:"aa_points" gorm:"Column:aa_points"`
	Hp                       null.Int64  `json:"hp" gorm:"Column:hp"`
	Mana                     null.Int64  `json:"mana" gorm:"Column:mana"`
	Endurance                null.Int64  `json:"endurance" gorm:"Column:endurance"`
	Ac                       null.Int    `json:"ac" gorm:"Column:ac"`
	Strength                 null.Int    `json:"strength" gorm:"Column:strength"`
	Stamina                  null.Int    `json:"stamina" gorm:"Column:stamina"`
	Dexterity                null.Int    `json:"dexterity" gorm:"Column:dexterity"`
	Agility                  null.Int    `json:"agility" gorm:"Column:agility"`
	Intelligence             null.Int    `json:"intelligence" gorm:"Column:intelligence"`
	Wisdom                   null.Int    `json:"wisdom" gorm:"Column:wisdom"`
	Charisma                 null.Int    `json:"charisma" gorm:"Column:charisma"`
	MagicResist              null.Int    `json:"magic_resist" gorm:"Column:magic_resist"`
	FireResist               null.Int    `json:"fire_resist" gorm:"Column:fire_resist"`
	ColdResist               null.Int    `json:"cold_resist" gorm:"Column:cold_resist"`
	PoisonResist             null.Int    `json:"poison_resist" gorm:"Column:poison_resist"`
	DiseaseResist            null.Int    `json:"disease_resist" gorm:"Column:disease_resist"`
	CorruptionResist         null.Int    `json:"corruption_resist" gorm:"Column:corruption_resist"`
	HeroicStrength           null.Int    `json:"heroic_strength" gorm:"Column:heroic_strength"`
	HeroicStamina            null.Int    `json:"heroic_stamina" gorm:"Column:heroic_stamina"`
	HeroicDexterity          null.Int    `json:"heroic_dexterity" gorm:"Column:heroic_dexterity"`
	HeroicAgility            null.Int    `json:"heroic_agility" gorm:"Column:heroic_agility"`
	HeroicIntelligence       null.Int    `json:"heroic_intelligence" gorm:"Column:heroic_intelligence"`
	HeroicWisdom             null.Int    `json:"heroic_wisdom" gorm:"Column:heroic_wisdom"`
	HeroicCharisma           null.Int    `json:"heroic_charisma" gorm:"Column:heroic_charisma"`
	HeroicMagicResist        null.Int    `json:"heroic_magic_resist" gorm:"Column:heroic_magic_resist"`
	HeroicFireResist         null.Int    `json:"heroic_fire_resist" gorm:"Column:heroic_fire_resist"`
	HeroicColdResist         null.Int    `json:"heroic_cold_resist" gorm:"Column:heroic_cold_resist"`
	HeroicPoisonResist       null.Int    `json:"heroic_poison_resist" gorm:"Column:heroic_poison_resist"`
	HeroicDiseaseResist      null.Int    `json:"heroic_disease_resist" gorm:"Column:heroic_disease_resist"`
	HeroicCorruptionResist   null.Int    `json:"heroic_corruption_resist" gorm:"Column:heroic_corruption_resist"`
	Haste                    null.Int    `json:"haste" gorm:"Column:haste"`
	Accuracy                 null.Int    `json:"accuracy" gorm:"Column:accuracy"`
	Attack                   null.Int    `json:"attack" gorm:"Column:attack"`
	Avoidance                null.Int    `json:"avoidance" gorm:"Column:avoidance"`
	Clairvoyance             null.Int    `json:"clairvoyance" gorm:"Column:clairvoyance"`
	CombatEffects            null.Int    `json:"combat_effects" gorm:"Column:combat_effects"`
	DamageShieldMitigation   null.Int    `json:"damage_shield_mitigation" gorm:"Column:damage_shield_mitigation"`
	DamageShield             null.Int    `json:"damage_shield" gorm:"Column:damage_shield"`
	DotShielding             null.Int    `json:"dot_shielding" gorm:"Column:dot_shielding"`
	HpRegen                  null.Int    `json:"hp_regen" gorm:"Column:hp_regen"`
	ManaRegen                null.Int    `json:"mana_regen" gorm:"Column:mana_regen"`
	EnduranceRegen           null.Int    `json:"endurance_regen" gorm:"Column:endurance_regen"`
	Shielding                null.Int    `json:"shielding" gorm:"Column:shielding"`
	SpellDamage              null.Int    `json:"spell_damage" gorm:"Column:spell_damage"`
	SpellShielding           null.Int    `json:"spell_shielding" gorm:"Column:spell_shielding"`
	Strikethrough            null.Int    `json:"strikethrough" gorm:"Column:strikethrough"`
	StunResist               null.Int    `json:"stun_resist" gorm:"Column:stun_resist"`
	Backstab                 null.Int    `json:"backstab" gorm:"Column:backstab"`
	Wind                     null.Int    `json:"wind" gorm:"Column:wind"`
	Brass                    null.Int    `json:"brass" gorm:"Column:brass"`
	String                   null.Int    `json:"string" gorm:"Column:string"`
	Percussion               null.Int    `json:"percussion" gorm:"Column:percussion"`
	Singing                  null.Int    `json:"singing" gorm:"Column:singing"`
	Baking                   null.Int    `json:"baking" gorm:"Column:baking"`
	Alchemy                  null.Int    `json:"alchemy" gorm:"Column:alchemy"`
	Tailoring                null.Int    `json:"tailoring" gorm:"Column:tailoring"`
	Blacksmithing            null.Int    `json:"blacksmithing" gorm:"Column:blacksmithing"`
	Fletching                null.Int    `json:"fletching" gorm:"Column:fletching"`
	Brewing                  null.Int    `json:"brewing" gorm:"Column:brewing"`
	Jewelry                  null.Int    `json:"jewelry" gorm:"Column:jewelry"`
	Pottery                  null.Int    `json:"pottery" gorm:"Column:pottery"`
	Research                 null.Int    `json:"research" gorm:"Column:research"`
	Alcohol                  null.Int    `json:"alcohol" gorm:"Column:alcohol"`
	Fishing                  null.Int    `json:"fishing" gorm:"Column:fishing"`
	Tinkering                null.Int    `json:"tinkering" gorm:"Column:tinkering"`
	CreatedAt                null.Time   `json:"created_at" gorm:"Column:created_at"`
	UpdatedAt                null.Time   `json:"updated_at" gorm:"Column:updated_at"`
}

func (CharacterStatsRecord) TableName() string {
    return "character_stats_record"
}

func (CharacterStatsRecord) Relationships() []string {
    return []string{}
}

func (CharacterStatsRecord) Connection() string {
    return ""
}
