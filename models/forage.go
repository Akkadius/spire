package models

import (
	"github.com/volatiletech/null/v8"
)

type Forage struct {
	ID                     int         `json:"id" gorm:"Column:id"`
	Zoneid                 int         `json:"zoneid" gorm:"Column:zoneid"`
	Itemid                 int         `json:"itemid" gorm:"Column:Itemid"`
	Level                  int16       `json:"level" gorm:"Column:level"`
	Chance                 int16       `json:"chance" gorm:"Column:chance"`
	MinExpansion           uint8       `json:"min_expansion" gorm:"Column:min_expansion"`
	MaxExpansion           uint8       `json:"max_expansion" gorm:"Column:max_expansion"`
	ContentFlags           null.String `json:"content_flags" gorm:"Column:content_flags"`
	ContentFlagsDisabled   null.String `json:"content_flags_disabled" gorm:"Column:content_flags_disabled"`
}

func (Forage) TableName() string {
    return "forage"
}

func (Forage) Relationships() []string {
    return []string{}
}

func (Forage) Connection() string {
    return "eqemu_content"
}
