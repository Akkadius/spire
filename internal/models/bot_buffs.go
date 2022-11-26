package models

type BotBuff struct {
	BuffsIndex          uint  `json:"buffs_index" gorm:"Column:buffs_index"`
	BotId               uint  `json:"bot_id" gorm:"Column:bot_id"`
	SpellId             uint  `json:"spell_id" gorm:"Column:spell_id"`
	CasterLevel         uint8 `json:"caster_level" gorm:"Column:caster_level"`
	DurationFormula     uint  `json:"duration_formula" gorm:"Column:duration_formula"`
	TicsRemaining       uint  `json:"tics_remaining" gorm:"Column:tics_remaining"`
	PoisonCounters      uint  `json:"poison_counters" gorm:"Column:poison_counters"`
	DiseaseCounters     uint  `json:"disease_counters" gorm:"Column:disease_counters"`
	CurseCounters       uint  `json:"curse_counters" gorm:"Column:curse_counters"`
	CorruptionCounters  uint  `json:"corruption_counters" gorm:"Column:corruption_counters"`
	Numhits             uint  `json:"numhits" gorm:"Column:numhits"`
	MeleeRune           uint  `json:"melee_rune" gorm:"Column:melee_rune"`
	MagicRune           uint  `json:"magic_rune" gorm:"Column:magic_rune"`
	DotRune             uint  `json:"dot_rune" gorm:"Column:dot_rune"`
	Persistent          int8  `json:"persistent" gorm:"Column:persistent"`
	CastonX             int   `json:"caston_x" gorm:"Column:caston_x"`
	CastonY             int   `json:"caston_y" gorm:"Column:caston_y"`
	CastonZ             int   `json:"caston_z" gorm:"Column:caston_z"`
	ExtraDiChance       uint  `json:"extra_di_chance" gorm:"Column:extra_di_chance"`
	InstrumentMod       int   `json:"instrument_mod" gorm:"Column:instrument_mod"`
}

func (BotBuff) TableName() string {
    return "bot_buffs"
}

func (BotBuff) Relationships() []string {
    return []string{}
}

func (BotBuff) Connection() string {
    return ""
}
