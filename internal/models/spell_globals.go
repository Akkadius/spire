package models

type SpellGlobal struct {
	Spellid    int    `json:"spellid" gorm:"Column:spellid"`
	SpellName  string `json:"spell_name" gorm:"Column:spell_name"`
	Qglobal    string `json:"qglobal" gorm:"Column:qglobal"`
	Value      string `json:"value" gorm:"Column:value"`
}

func (SpellGlobal) TableName() string {
    return "spell_globals"
}

func (SpellGlobal) Relationships() []string {
    return []string{}
}

func (SpellGlobal) Connection() string {
    return ""
}
