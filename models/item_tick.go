package models

type ItemTick struct {
	ItItemid   int    `json:"it_itemid" gorm:"Column:it_itemid"`
	ItChance   int    `json:"it_chance" gorm:"Column:it_chance"`
	ItLevel    int    `json:"it_level" gorm:"Column:it_level"`
	ItId       int    `json:"it_id" gorm:"Column:it_id"`
	ItQglobal  string `json:"it_qglobal" gorm:"Column:it_qglobal"`
	ItBagslot  int8   `json:"it_bagslot" gorm:"Column:it_bagslot"`
}

func (ItemTick) TableName() string {
    return "item_tick"
}

func (ItemTick) Relationships() []string {
    return []string{}
}

func (ItemTick) Connection() string {
    return ""
}
