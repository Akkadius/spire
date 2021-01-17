package models

type CharacterBandolier struct {
	ID             uint   `json:"id" gorm:"Column:id"`
	BandolierId    uint8  `json:"bandolier_id" gorm:"Column:bandolier_id"`
	BandolierSlot  uint8  `json:"bandolier_slot" gorm:"Column:bandolier_slot"`
	ItemId         uint   `json:"item_id" gorm:"Column:item_id"`
	Icon           uint   `json:"icon" gorm:"Column:icon"`
	BandolierName  string `json:"bandolier_name" gorm:"Column:bandolier_name"`
}

func (CharacterBandolier) TableName() string {
    return "character_bandolier"
}

func (CharacterBandolier) Relationships() []string {
    return []string{}
}

func (CharacterBandolier) Connection() string {
    return ""
}
