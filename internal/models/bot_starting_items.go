package models

import (
	"github.com/volatiletech/null/v8"
)

type BotStartingItem struct {
	ID                     uint        `json:"id" gorm:"Column:id"`
	Races                  uint        `json:"races" gorm:"Column:races"`
	Classes                uint        `json:"classes" gorm:"Column:classes"`
	ItemId                 uint        `json:"item_id" gorm:"Column:item_id"`
	ItemCharges            uint8       `json:"item_charges" gorm:"Column:item_charges"`
	MinStatus              uint8       `json:"min_status" gorm:"Column:min_status"`
	SlotId                 int32       `json:"slot_id" gorm:"Column:slot_id"`
	MinExpansion           int8        `json:"min_expansion" gorm:"Column:min_expansion"`
	MaxExpansion           int8        `json:"max_expansion" gorm:"Column:max_expansion"`
	ContentFlags           null.String `json:"content_flags" gorm:"Column:content_flags"`
	ContentFlagsDisabled   null.String `json:"content_flags_disabled" gorm:"Column:content_flags_disabled"`
}

func (BotStartingItem) TableName() string {
    return "bot_starting_items"
}

func (BotStartingItem) Relationships() []string {
    return []string{}
}

func (BotStartingItem) Connection() string {
    return ""
}
