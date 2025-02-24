package models

import (
	"github.com/volatiletech/null/v8"
)

type Sharedbank struct {
	AccountId           uint        `json:"account_id" gorm:"Column:account_id"`
	SlotId              uint32      `json:"slot_id" gorm:"Column:slot_id"`
	ItemId              uint        `json:"item_id" gorm:"Column:item_id"`
	Charges             uint16      `json:"charges" gorm:"Column:charges"`
	Color               uint        `json:"color" gorm:"Column:color"`
	AugmentOne          uint32      `json:"augment_one" gorm:"Column:augment_one"`
	AugmentTwo          uint32      `json:"augment_two" gorm:"Column:augment_two"`
	AugmentThree        uint32      `json:"augment_three" gorm:"Column:augment_three"`
	AugmentFour         uint32      `json:"augment_four" gorm:"Column:augment_four"`
	AugmentFive         uint32      `json:"augment_five" gorm:"Column:augment_five"`
	AugmentSix          uint32      `json:"augment_six" gorm:"Column:augment_six"`
	CustomData          null.String `json:"custom_data" gorm:"Column:custom_data"`
	OrnamentIcon        uint        `json:"ornament_icon" gorm:"Column:ornament_icon"`
	OrnamentIdfile      uint        `json:"ornament_idfile" gorm:"Column:ornament_idfile"`
	OrnamentHeroModel   int         `json:"ornament_hero_model" gorm:"Column:ornament_hero_model"`
	Guid                uint64      `json:"guid" gorm:"Column:guid"`
}

func (Sharedbank) TableName() string {
    return "sharedbank"
}

func (Sharedbank) Relationships() []string {
    return []string{}
}

func (Sharedbank) Connection() string {
    return ""
}
