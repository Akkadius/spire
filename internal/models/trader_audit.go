package models

import (
	"time"
)

type TraderAudit struct {
	Time      time.Time `json:"time" gorm:"Column:time"`
	Seller    string    `json:"seller" gorm:"Column:seller"`
	Buyer     string    `json:"buyer" gorm:"Column:buyer"`
	Itemname  string    `json:"itemname" gorm:"Column:itemname"`
	Quantity  int       `json:"quantity" gorm:"Column:quantity"`
	Totalcost int       `json:"totalcost" gorm:"Column:totalcost"`
	Trantype  int8      `json:"trantype" gorm:"Column:trantype"`
}

func (TraderAudit) TableName() string {
    return "trader_audit"
}

func (TraderAudit) Relationships() []string {
    return []string{}
}

func (TraderAudit) Connection() string {
    return ""
}
