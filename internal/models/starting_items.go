package models

import (
	"github.com/volatiletech/null/v8"
)

type StartingItem struct {
	ID                     uint        `json:"id" gorm:"Column:id"`
	ClassList              null.String `json:"class_list" gorm:"Column:class_list"`
	RaceList               null.String `json:"race_list" gorm:"Column:race_list"`
	DeityList              null.String `json:"deity_list" gorm:"Column:deity_list"`
	ZoneIdList             null.String `json:"zone_id_list" gorm:"Column:zone_id_list"`
	ItemId                 uint        `json:"item_id" gorm:"Column:item_id"`
	ItemCharges            uint8       `json:"item_charges" gorm:"Column:item_charges"`
	AugmentOne             uint        `json:"augment_one" gorm:"Column:augment_one"`
	AugmentTwo             uint        `json:"augment_two" gorm:"Column:augment_two"`
	AugmentThree           uint        `json:"augment_three" gorm:"Column:augment_three"`
	AugmentFour            uint        `json:"augment_four" gorm:"Column:augment_four"`
	AugmentFive            uint        `json:"augment_five" gorm:"Column:augment_five"`
	AugmentSix             uint        `json:"augment_six" gorm:"Column:augment_six"`
	Status                 int32       `json:"status" gorm:"Column:status"`
	InventorySlot          int32       `json:"inventory_slot" gorm:"Column:inventory_slot"`
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
