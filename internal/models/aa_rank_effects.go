package models

type AaRankEffect struct {
	RankId    uint `json:"rank_id" gorm:"Column:rank_id"`
	Slot      uint `json:"slot" gorm:"Column:slot"`
	EffectId  int  `json:"effect_id" gorm:"Column:effect_id"`
	Base1     int  `json:"base_1" gorm:"Column:base1"`
	Base2     int  `json:"base_2" gorm:"Column:base2"`
}

func (AaRankEffect) TableName() string {
    return "aa_rank_effects"
}

func (AaRankEffect) Relationships() []string {
    return []string{}
}

func (AaRankEffect) Connection() string {
    return "eqemu_content"
}
