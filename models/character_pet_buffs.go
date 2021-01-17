package models

type CharacterPetBuff struct {
	CharId         int    `json:"char_id" gorm:"Column:char_id"`
	Pet            int    `json:"pet" gorm:"Column:pet"`
	Slot           int    `json:"slot" gorm:"Column:slot"`
	SpellId        int    `json:"spell_id" gorm:"Column:spell_id"`
	CasterLevel    int8   `json:"caster_level" gorm:"Column:caster_level"`
	Castername     string `json:"castername" gorm:"Column:castername"`
	Ticsremaining  int    `json:"ticsremaining" gorm:"Column:ticsremaining"`
	Counters       int    `json:"counters" gorm:"Column:counters"`
	Numhits        int    `json:"numhits" gorm:"Column:numhits"`
	Rune           int    `json:"rune" gorm:"Column:rune"`
	InstrumentMod  uint8  `json:"instrument_mod" gorm:"Column:instrument_mod"`
}

func (CharacterPetBuff) TableName() string {
    return "character_pet_buffs"
}

func (CharacterPetBuff) Relationships() []string {
    return []string{}
}

func (CharacterPetBuff) Connection() string {
    return ""
}
