package models

import (
	"github.com/volatiletech/null/v8"
)

type Parcel struct {
	ID        uint        `json:"id" gorm:"Column:id"`
	ItemId    uint        `json:"item_id" gorm:"Column:item_id"`
	SlotId    uint        `json:"slot_id" gorm:"Column:slot_id"`
	Quantity  uint        `json:"quantity" gorm:"Column:quantity"`
	ToName    null.String `json:"to_name" gorm:"Column:to_name"`
	FromName  null.String `json:"from_name" gorm:"Column:from_name"`
	Note      null.String `json:"note" gorm:"Column:note"`
	SentDate  null.Time   `json:"sent_date" gorm:"Column:sent_date"`
}

func (Parcel) TableName() string {
    return "parcels"
}

func (Parcel) Relationships() []string {
    return []string{}
}

func (Parcel) Connection() string {
    return ""
}
