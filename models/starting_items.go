package models

import (
	"github.com/volatiletech/null/v8"
)

type StartingItem struct {
	ID                     uint        `json:"id" gorm:"Column:id"`
	Race                   int         `json:"race" gorm:"Column:race"`
	Class                  int         `json:"class" gorm:"Column:class"`
	Deityid                int         `json:"deityid" gorm:"Column:deityid"`
	Zoneid                 int         `json:"zoneid" gorm:"Column:zoneid"`
	Itemid                 int         `json:"itemid" gorm:"Column:itemid"`
	ItemCharges            uint8       `json:"item_charges" gorm:"Column:item_charges"`
	Gm                     int8        `json:"gm" gorm:"Column:gm"`
	Slot                   int32       `json:"slot" gorm:"Column:slot"`
	MinExpansion           uint8       `json:"min_expansion" gorm:"Column:min_expansion"`
	MaxExpansion           uint8       `json:"max_expansion" gorm:"Column:max_expansion"`
	ContentFlags           null.String `json:"content_flags" gorm:"Column:content_flags"`
	ContentFlagsDisabled   null.String `json:"content_flags_disabled" gorm:"Column:content_flags_disabled"`
}

func (StartingItem) TableName() string {
    return "starting_items"
}

func (StartingItem) Relationships() []string {
    return []string{}
}

func (StartingItem) Connection() string {
    return "eqemu_content"
}
