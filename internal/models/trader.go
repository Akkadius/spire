package models

type Trader struct {
	CharId       uint  `json:"char_id" gorm:"Column:char_id"`
	ItemId       uint  `json:"item_id" gorm:"Column:item_id"`
	Serialnumber uint  `json:"serialnumber" gorm:"Column:serialnumber"`
	Charges      int   `json:"charges" gorm:"Column:charges"`
	ItemCost     uint  `json:"item_cost" gorm:"Column:item_cost"`
	SlotId       uint8 `json:"slot_id" gorm:"Column:slot_id"`
}

func (Trader) TableName() string {
    return "trader"
}

func (Trader) Relationships() []string {
    return []string{}
}

func (Trader) Connection() string {
    return ""
}
