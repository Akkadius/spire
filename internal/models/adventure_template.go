package models

import (
	"github.com/volatiletech/null/v8"
)

type AdventureTemplate struct {
	ID                uint        `json:"id" gorm:"Column:id"`
	Zone              string      `json:"zone" gorm:"Column:zone"`
	ZoneVersion       uint8       `json:"zone_version" gorm:"Column:zone_version"`
	IsHard            uint8       `json:"is_hard" gorm:"Column:is_hard"`
	IsRaid            uint8       `json:"is_raid" gorm:"Column:is_raid"`
	MinLevel          uint8       `json:"min_level" gorm:"Column:min_level"`
	MaxLevel          uint8       `json:"max_level" gorm:"Column:max_level"`
	Type              uint8       `json:"type" gorm:"Column:type"`
	TypeData          uint        `json:"type_data" gorm:"Column:type_data"`
	TypeCount         uint16      `json:"type_count" gorm:"Column:type_count"`
	AssaX             float32     `json:"assa_x" gorm:"Column:assa_x"`
	AssaY             float32     `json:"assa_y" gorm:"Column:assa_y"`
	AssaZ             float32     `json:"assa_z" gorm:"Column:assa_z"`
	AssaH             float32     `json:"assa_h" gorm:"Column:assa_h"`
	Text              null.String `json:"text" gorm:"Column:text"`
	Duration          uint        `json:"duration" gorm:"Column:duration"`
	ZoneInTime        uint        `json:"zone_in_time" gorm:"Column:zone_in_time"`
	WinPoints         uint16      `json:"win_points" gorm:"Column:win_points"`
	LosePoints        uint16      `json:"lose_points" gorm:"Column:lose_points"`
	Theme             uint8       `json:"theme" gorm:"Column:theme"`
	ZoneInZoneId      uint16      `json:"zone_in_zone_id" gorm:"Column:zone_in_zone_id"`
	ZoneInX           float32     `json:"zone_in_x" gorm:"Column:zone_in_x"`
	ZoneInY           float32     `json:"zone_in_y" gorm:"Column:zone_in_y"`
	ZoneInObjectId    int16       `json:"zone_in_object_id" gorm:"Column:zone_in_object_id"`
	DestX             float32     `json:"dest_x" gorm:"Column:dest_x"`
	DestY             float32     `json:"dest_y" gorm:"Column:dest_y"`
	DestZ             float32     `json:"dest_z" gorm:"Column:dest_z"`
	DestH             float32     `json:"dest_h" gorm:"Column:dest_h"`
	GraveyardZoneId   uint        `json:"graveyard_zone_id" gorm:"Column:graveyard_zone_id"`
	GraveyardX        float32     `json:"graveyard_x" gorm:"Column:graveyard_x"`
	GraveyardY        float32     `json:"graveyard_y" gorm:"Column:graveyard_y"`
	GraveyardZ        float32     `json:"graveyard_z" gorm:"Column:graveyard_z"`
	GraveyardRadius   string      `json:"graveyard_radius" gorm:"Column:graveyard_radius"`
}

func (AdventureTemplate) TableName() string {
    return "adventure_template"
}

func (AdventureTemplate) Relationships() []string {
    return []string{}
}

func (AdventureTemplate) Connection() string {
    return "eqemu_content"
}
