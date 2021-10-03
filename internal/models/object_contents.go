package models

import (
	"github.com/volatiletech/null/v8"
	"time"
)

type ObjectContent struct {
	Zoneid   uint        `json:"zoneid" gorm:"Column:zoneid"`
	Parentid uint        `json:"parentid" gorm:"Column:parentid"`
	Bagidx   uint        `json:"bagidx" gorm:"Column:bagidx"`
	Itemid   uint        `json:"itemid" gorm:"Column:itemid"`
	Charges  int16       `json:"charges" gorm:"Column:charges"`
	Droptime time.Time   `json:"droptime" gorm:"Column:droptime"`
	Augslot1 null.Uint32 `json:"augslot_1" gorm:"Column:augslot1"`
	Augslot2 null.Uint32 `json:"augslot_2" gorm:"Column:augslot2"`
	Augslot3 null.Uint32 `json:"augslot_3" gorm:"Column:augslot3"`
	Augslot4 null.Uint32 `json:"augslot_4" gorm:"Column:augslot4"`
	Augslot5 null.Uint32 `json:"augslot_5" gorm:"Column:augslot5"`
	Augslot6 int32       `json:"augslot_6" gorm:"Column:augslot6"`
}

func (ObjectContent) TableName() string {
    return "object_contents"
}

func (ObjectContent) Relationships() []string {
    return []string{}
}

func (ObjectContent) Connection() string {
    return ""
}
