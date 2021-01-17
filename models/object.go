package models

import (
	"github.com/volatiletech/null/v8"
)

type Object struct {
	ID                     int         `json:"id" gorm:"Column:id"`
	Zoneid                 uint        `json:"zoneid" gorm:"Column:zoneid"`
	Version                int16       `json:"version" gorm:"Column:version"`
	Xpos                   float32     `json:"xpos" gorm:"Column:xpos"`
	Ypos                   float32     `json:"ypos" gorm:"Column:ypos"`
	Zpos                   float32     `json:"zpos" gorm:"Column:zpos"`
	Heading                float32     `json:"heading" gorm:"Column:heading"`
	Itemid                 int         `json:"itemid" gorm:"Column:itemid"`
	Charges                uint16      `json:"charges" gorm:"Column:charges"`
	Objectname             null.String `json:"objectname" gorm:"Column:objectname"`
	Type                   int         `json:"type" gorm:"Column:type"`
	Icon                   int         `json:"icon" gorm:"Column:icon"`
	Unknown08              int32       `json:"unknown_08" gorm:"Column:unknown08"`
	Unknown10              int32       `json:"unknown_10" gorm:"Column:unknown10"`
	Unknown20              int         `json:"unknown_20" gorm:"Column:unknown20"`
	Unknown24              int         `json:"unknown_24" gorm:"Column:unknown24"`
	Unknown60              int         `json:"unknown_60" gorm:"Column:unknown60"`
	Unknown64              int         `json:"unknown_64" gorm:"Column:unknown64"`
	Unknown68              int         `json:"unknown_68" gorm:"Column:unknown68"`
	Unknown72              int         `json:"unknown_72" gorm:"Column:unknown72"`
	Unknown76              int         `json:"unknown_76" gorm:"Column:unknown76"`
	Unknown84              int         `json:"unknown_84" gorm:"Column:unknown84"`
	Size                   float32     `json:"size" gorm:"Column:size"`
	TiltX                  float32     `json:"tilt_x" gorm:"Column:tilt_x"`
	TiltY                  float32     `json:"tilt_y" gorm:"Column:tilt_y"`
	DisplayName            null.String `json:"display_name" gorm:"Column:display_name"`
	MinExpansion           uint8       `json:"min_expansion" gorm:"Column:min_expansion"`
	MaxExpansion           uint8       `json:"max_expansion" gorm:"Column:max_expansion"`
	ContentFlags           null.String `json:"content_flags" gorm:"Column:content_flags"`
	ContentFlagsDisabled   null.String `json:"content_flags_disabled" gorm:"Column:content_flags_disabled"`
}

func (Object) TableName() string {
    return "object"
}

func (Object) Relationships() []string {
    return []string{}
}

func (Object) Connection() string {
    return "eqemu_content"
}
