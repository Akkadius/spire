package models

import (
	"github.com/volatiletech/null/v8"
)

type Sharedbank struct {
	Acctid      uint   `json:"acctid" gorm:"Column:acctid"`
	Slotid      uint32 `json:"slotid" gorm:"Column:slotid"`
	Itemid      uint   `json:"itemid" gorm:"Column:itemid"`
	Charges     uint16 `json:"charges" gorm:"Column:charges"`
	Augslot1    uint32      `json:"augslot_1" gorm:"Column:augslot1"`
	Augslot2    uint32      `json:"augslot_2" gorm:"Column:augslot2"`
	Augslot3    uint32      `json:"augslot_3" gorm:"Column:augslot3"`
	Augslot4    uint32      `json:"augslot_4" gorm:"Column:augslot4"`
	Augslot5    uint32      `json:"augslot_5" gorm:"Column:augslot5"`
	Augslot6    int32       `json:"augslot_6" gorm:"Column:augslot6"`
	CustomData  string `json:"custom_data" gorm:"Column:custom_data"`
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
