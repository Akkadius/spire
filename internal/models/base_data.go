package models

type BaseDatum struct {
	Level    uint    `json:"level" gorm:"Column:level"`
	Class    uint    `json:"class" gorm:"Column:class"`
	Hp       float64 `json:"hp" gorm:"Column:hp"`
	Mana     float64 `json:"mana" gorm:"Column:mana"`
	End      float64 `json:"end" gorm:"Column:end"`
	Unk1     float64 `json:"unk_1" gorm:"Column:unk1"`
	Unk2     float64 `json:"unk_2" gorm:"Column:unk2"`
	HpFac    float64 `json:"hp_fac" gorm:"Column:hp_fac"`
	ManaFac  float64 `json:"mana_fac" gorm:"Column:mana_fac"`
	EndFac   float64 `json:"end_fac" gorm:"Column:end_fac"`
}

func (BaseDatum) TableName() string {
    return "base_data"
}

func (BaseDatum) Relationships() []string {
    return []string{}
}

func (BaseDatum) Connection() string {
    return "eqemu_content"
}
