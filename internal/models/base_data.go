package models

type BaseDatum struct {
	Level     uint8   `json:"level" gorm:"Column:level"`
	Class     uint8   `json:"class" gorm:"Column:class"`
	Hp        float64 `json:"hp" gorm:"Column:hp"`
	Mana      float64 `json:"mana" gorm:"Column:mana"`
	End       float64 `json:"end" gorm:"Column:end"`
	HpRegen   float64 `json:"hp_regen" gorm:"Column:hp_regen"`
	EndRegen  float64 `json:"end_regen" gorm:"Column:end_regen"`
	HpFac     float64 `json:"hp_fac" gorm:"Column:hp_fac"`
	ManaFac   float64 `json:"mana_fac" gorm:"Column:mana_fac"`
	EndFac    float64 `json:"end_fac" gorm:"Column:end_fac"`
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
