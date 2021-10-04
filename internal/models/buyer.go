package models

type Buyer struct {
	Charid   int    `json:"charid" gorm:"Column:charid"`
	Buyslot  int    `json:"buyslot" gorm:"Column:buyslot"`
	Itemid   int    `json:"itemid" gorm:"Column:itemid"`
	Itemname string `json:"itemname" gorm:"Column:itemname"`
	Quantity int    `json:"quantity" gorm:"Column:quantity"`
	Price    int    `json:"price" gorm:"Column:price"`
}

func (Buyer) TableName() string {
    return "buyer"
}

func (Buyer) Relationships() []string {
    return []string{}
}

func (Buyer) Connection() string {
    return ""
}
