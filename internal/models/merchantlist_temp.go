package models

type MerchantlistTemp struct {
	Npcid   uint  `json:"npcid" gorm:"Column:npcid"`
	Slot    uint8 `json:"slot" gorm:"Column:slot"`
	Itemid  uint  `json:"itemid" gorm:"Column:itemid"`
	Charges uint  `json:"charges" gorm:"Column:charges"`
}

func (MerchantlistTemp) TableName() string {
    return "merchantlist_temp"
}

func (MerchantlistTemp) Relationships() []string {
    return []string{}
}

func (MerchantlistTemp) Connection() string {
    return ""
}
