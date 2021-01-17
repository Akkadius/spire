package models

type CharacterCurrency struct {
	ID                      uint `json:"id" gorm:"Column:id"`
	Platinum                uint `json:"platinum" gorm:"Column:platinum"`
	Gold                    uint `json:"gold" gorm:"Column:gold"`
	Silver                  uint `json:"silver" gorm:"Column:silver"`
	Copper                  uint `json:"copper" gorm:"Column:copper"`
	PlatinumBank            uint `json:"platinum_bank" gorm:"Column:platinum_bank"`
	GoldBank                uint `json:"gold_bank" gorm:"Column:gold_bank"`
	SilverBank              uint `json:"silver_bank" gorm:"Column:silver_bank"`
	CopperBank              uint `json:"copper_bank" gorm:"Column:copper_bank"`
	PlatinumCursor          uint `json:"platinum_cursor" gorm:"Column:platinum_cursor"`
	GoldCursor              uint `json:"gold_cursor" gorm:"Column:gold_cursor"`
	SilverCursor            uint `json:"silver_cursor" gorm:"Column:silver_cursor"`
	CopperCursor            uint `json:"copper_cursor" gorm:"Column:copper_cursor"`
	RadiantCrystals         uint `json:"radiant_crystals" gorm:"Column:radiant_crystals"`
	CareerRadiantCrystals   uint `json:"career_radiant_crystals" gorm:"Column:career_radiant_crystals"`
	EbonCrystals            uint `json:"ebon_crystals" gorm:"Column:ebon_crystals"`
	CareerEbonCrystals      uint `json:"career_ebon_crystals" gorm:"Column:career_ebon_crystals"`
}

func (CharacterCurrency) TableName() string {
    return "character_currency"
}

func (CharacterCurrency) Relationships() []string {
    return []string{}
}

func (CharacterCurrency) Connection() string {
    return ""
}
