package models

import (
	"github.com/volatiletech/null/v8"
)

type InventorySnapshot struct {
	TimeIndex           uint        `json:"time_index" gorm:"Column:time_index"`
	Charid              uint        `json:"charid" gorm:"Column:charid"`
	Slotid              uint32      `json:"slotid" gorm:"Column:slotid"`
	Itemid              null.Uint   `json:"itemid" gorm:"Column:itemid"`
	Charges             null.Uint16 `json:"charges" gorm:"Column:charges"`
	Color               uint        `json:"color" gorm:"Column:color"`
	Augslot1            uint32      `json:"augslot_1" gorm:"Column:augslot1"`
	Augslot2            uint32      `json:"augslot_2" gorm:"Column:augslot2"`
	Augslot3            uint32      `json:"augslot_3" gorm:"Column:augslot3"`
	Augslot4            uint32      `json:"augslot_4" gorm:"Column:augslot4"`
	Augslot5            null.Uint32 `json:"augslot_5" gorm:"Column:augslot5"`
	Augslot6            int32       `json:"augslot_6" gorm:"Column:augslot6"`
	Instnodrop          uint8       `json:"instnodrop" gorm:"Column:instnodrop"`
	CustomData          null.String `json:"custom_data" gorm:"Column:custom_data"`
	Ornamenticon        uint        `json:"ornamenticon" gorm:"Column:ornamenticon"`
	Ornamentidfile      uint        `json:"ornamentidfile" gorm:"Column:ornamentidfile"`
	OrnamentHeroModel   int         `json:"ornament_hero_model" gorm:"Column:ornament_hero_model"`
	Guid                null.Uint64 `json:"guid" gorm:"Column:guid"`
}

func (InventorySnapshot) TableName() string {
    return "inventory_snapshots"
}

func (InventorySnapshot) Relationships() []string {
    return []string{}
}

func (InventorySnapshot) Connection() string {
    return ""
}
