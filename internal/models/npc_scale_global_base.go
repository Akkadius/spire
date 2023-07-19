package models

type NpcScaleGlobalBase struct {
	Type                  int    `json:"type" gorm:"Column:type"`
	Level                 int    `json:"level" gorm:"Column:level"`
	ZoneId                uint   `json:"zone_id" gorm:"Column:zone_id"`
	InstanceVersion       uint   `json:"instance_version" gorm:"Column:instance_version"`
	ZoneIdList            string `json:"zone_id_list" gorm:"Column:zone_id_list"`
	InstanceVersionList   string `json:"instance_version_list" gorm:"Column:instance_version_list"`
	Ac                    int    `json:"ac" gorm:"Column:ac"`
	Hp                    int64  `json:"hp" gorm:"Column:hp"`
	Accuracy              int    `json:"accuracy" gorm:"Column:accuracy"`
	SlowMitigation        int    `json:"slow_mitigation" gorm:"Column:slow_mitigation"`
	Attack                int    `json:"attack" gorm:"Column:attack"`
	Strength              int    `json:"strength" gorm:"Column:strength"`
	Stamina               int    `json:"stamina" gorm:"Column:stamina"`
	Dexterity             int    `json:"dexterity" gorm:"Column:dexterity"`
	Agility               int    `json:"agility" gorm:"Column:agility"`
	Intelligence          int    `json:"intelligence" gorm:"Column:intelligence"`
	Wisdom                int    `json:"wisdom" gorm:"Column:wisdom"`
	Charisma              int    `json:"charisma" gorm:"Column:charisma"`
	MagicResist           int    `json:"magic_resist" gorm:"Column:magic_resist"`
	ColdResist            int    `json:"cold_resist" gorm:"Column:cold_resist"`
	FireResist            int    `json:"fire_resist" gorm:"Column:fire_resist"`
	PoisonResist          int    `json:"poison_resist" gorm:"Column:poison_resist"`
	DiseaseResist         int    `json:"disease_resist" gorm:"Column:disease_resist"`
	CorruptionResist      int    `json:"corruption_resist" gorm:"Column:corruption_resist"`
	PhysicalResist        int    `json:"physical_resist" gorm:"Column:physical_resist"`
	MinDmg                int    `json:"min_dmg" gorm:"Column:min_dmg"`
	MaxDmg                int    `json:"max_dmg" gorm:"Column:max_dmg"`
	HpRegenRate           int64  `json:"hp_regen_rate" gorm:"Column:hp_regen_rate"`
	HpRegenPerSecond      int64  `json:"hp_regen_per_second" gorm:"Column:hp_regen_per_second"`
	AttackDelay           int    `json:"attack_delay" gorm:"Column:attack_delay"`
	SpellScale            int    `json:"spell_scale" gorm:"Column:spell_scale"`
	HealScale             int    `json:"heal_scale" gorm:"Column:heal_scale"`
	Avoidance             uint   `json:"avoidance" gorm:"Column:avoidance"`
	HeroicStrikethrough   int    `json:"heroic_strikethrough" gorm:"Column:heroic_strikethrough"`
	SpecialAbilities      string `json:"special_abilities" gorm:"Column:special_abilities"`
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
