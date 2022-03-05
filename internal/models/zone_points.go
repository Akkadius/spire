package models

import (
	"github.com/volatiletech/null/v8"
)

type ZonePoint struct {
	ID                     int          `json:"id" gorm:"Column:id"`
	Zone                   null.String  `json:"zone" gorm:"Column:zone"`
	Version                int          `json:"version" gorm:"Column:version"`
	Number                 uint16       `json:"number" gorm:"Column:number"`
	Y                      float32      `json:"y" gorm:"Column:y"`
	X                      float32      `json:"x" gorm:"Column:x"`
	Z                      float32      `json:"z" gorm:"Column:z"`
	Heading                float32      `json:"heading" gorm:"Column:heading"`
	TargetY                float32      `json:"target_y" gorm:"Column:target_y"`
	TargetX                float32      `json:"target_x" gorm:"Column:target_x"`
	TargetZ                float32      `json:"target_z" gorm:"Column:target_z"`
	TargetHeading          float32      `json:"target_heading" gorm:"Column:target_heading"`
	Zoneinst               null.Uint16  `json:"zoneinst" gorm:"Column:zoneinst"`
	TargetZoneId           uint         `json:"target_zone_id" gorm:"Column:target_zone_id"`
	TargetInstance         uint         `json:"target_instance" gorm:"Column:target_instance"`
	Buffer                 null.Float32 `json:"buffer" gorm:"Column:buffer"`
	ClientVersionMask      uint         `json:"client_version_mask" gorm:"Column:client_version_mask"`
	MinExpansion           int8         `json:"min_expansion" gorm:"Column:min_expansion"`
	MaxExpansion           int8         `json:"max_expansion" gorm:"Column:max_expansion"`
	ContentFlags           null.String  `json:"content_flags" gorm:"Column:content_flags"`
	ContentFlagsDisabled   null.String  `json:"content_flags_disabled" gorm:"Column:content_flags_disabled"`
	IsVirtual              int8         `json:"is_virtual" gorm:"Column:is_virtual"`
	Height                 int          `json:"height" gorm:"Column:height"`
	Width                  int          `json:"width" gorm:"Column:width"`
}

func (ZonePoint) TableName() string {
    return "zone_points"
}

func (ZonePoint) Relationships() []string {
    return []string{}
}

func (ZonePoint) Connection() string {
    return "eqemu_content"
}
