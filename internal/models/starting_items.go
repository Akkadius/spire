package models

import (
	"github.com/volatiletech/null/v8"
)

type StartingItem struct {
	ID                     uint        `json:"id" gorm:"Column:id"`
	RaceList               null.String `json:"race_list" gorm:"Column:race_list"`
	ClassList              null.String `json:"class_list" gorm:"Column:class_list"`
	DeityList              null.String `json:"deity_list" gorm:"Column:deity_list"`
	ZoneIdList             null.String `json:"zone_id_list" gorm:"Column:zone_id_list"`
	ItemId                 uint        `json:"item_id" gorm:"Column:item_id"`
	ItemCharges            uint8       `json:"item_charges" gorm:"Column:item_charges"`
	Gm                     uint32      `json:"gm" gorm:"Column:gm"`
	Slot                   int32       `json:"slot" gorm:"Column:slot"`
	MinExpansion           int8        `json:"min_expansion" gorm:"Column:min_expansion"`
	MaxExpansion           int8        `json:"max_expansion" gorm:"Column:max_expansion"`
	ContentFlags           null.String `json:"content_flags" gorm:"Column:content_flags"`
	ContentFlagsDisabled   null.String `json:"content_flags_disabled" gorm:"Column:content_flags_disabled"`
	Zone                   *Zone       `json:"zone,omitempty" gorm:"foreignKey:zoneid;references:zoneidnumber"`
}

func (StartingItem) TableName() string {
    return "starting_items"
}

func (StartingItem) Relationships() []string {
    return []string{
		"Zone",
	}
}

func (StartingItem) Connection() string {
    return "eqemu_content"
}
