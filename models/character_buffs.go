package models

type CharacterBuff struct {
	CharacterId    uint   `json:"character_id" gorm:"Column:character_id"`
	SlotId         uint8  `json:"slot_id" gorm:"Column:slot_id"`
	SpellId        uint16 `json:"spell_id" gorm:"Column:spell_id"`
	CasterLevel    uint8  `json:"caster_level" gorm:"Column:caster_level"`
	CasterName     string `json:"caster_name" gorm:"Column:caster_name"`
	Ticsremaining  int    `json:"ticsremaining" gorm:"Column:ticsremaining"`
	Counters       uint   `json:"counters" gorm:"Column:counters"`
	Numhits        uint   `json:"numhits" gorm:"Column:numhits"`
	MeleeRune      uint   `json:"melee_rune" gorm:"Column:melee_rune"`
	MagicRune      uint   `json:"magic_rune" gorm:"Column:magic_rune"`
	Persistent     uint8  `json:"persistent" gorm:"Column:persistent"`
	DotRune        int    `json:"dot_rune" gorm:"Column:dot_rune"`
	CastonX        int    `json:"caston_x" gorm:"Column:caston_x"`
	CastonY        int    `json:"caston_y" gorm:"Column:caston_y"`
	CastonZ        int    `json:"caston_z" gorm:"Column:caston_z"`
	ExtraDIChance  int    `json:"extra_di_chance" gorm:"Column:ExtraDIChance"`
	InstrumentMod  int    `json:"instrument_mod" gorm:"Column:instrument_mod"`
}

func (CharacterBuff) TableName() string {
    return "character_buffs"
}

func (CharacterBuff) Relationships() []string {
    return []string{}
}

func (CharacterBuff) Connection() string {
    return ""
}
