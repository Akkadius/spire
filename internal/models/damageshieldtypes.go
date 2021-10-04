package models

type Damageshieldtype struct {
	Spellid uint  `json:"spellid" gorm:"Column:spellid"`
	Type    uint8 `json:"type" gorm:"Column:type"`
}

func (Damageshieldtype) TableName() string {
    return "damageshieldtypes"
}

func (Damageshieldtype) Relationships() []string {
    return []string{}
}

func (Damageshieldtype) Connection() string {
    return "eqemu_content"
}
