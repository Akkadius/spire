package models

import (
	"github.com/volatiletech/null/v8"
)

type GroundSpawn struct {
	ID                     uint        `json:"id" gorm:"Column:id"`
	Zoneid                 uint        `json:"zoneid" gorm:"Column:zoneid"`
	Version                int16       `json:"version" gorm:"Column:version"`
	MaxX                   float32     `json:"max_x" gorm:"Column:max_x"`
	MaxY                   float32     `json:"max_y" gorm:"Column:max_y"`
	MaxZ                   float32     `json:"max_z" gorm:"Column:max_z"`
	MinX                   float32     `json:"min_x" gorm:"Column:min_x"`
	MinY                   float32     `json:"min_y" gorm:"Column:min_y"`
	Heading                float32     `json:"heading" gorm:"Column:heading"`
	Name                   string      `json:"name" gorm:"Column:name"`
	Item                   uint        `json:"item" gorm:"Column:item"`
	MaxAllowed             uint        `json:"max_allowed" gorm:"Column:max_allowed"`
	Comment                string      `json:"comment" gorm:"Column:comment"`
	RespawnTimer           uint        `json:"respawn_timer" gorm:"Column:respawn_timer"`
	MinExpansion           uint8       `json:"min_expansion" gorm:"Column:min_expansion"`
	MaxExpansion           uint8       `json:"max_expansion" gorm:"Column:max_expansion"`
	ContentFlags           null.String `json:"content_flags" gorm:"Column:content_flags"`
	ContentFlagsDisabled   null.String `json:"content_flags_disabled" gorm:"Column:content_flags_disabled"`
	Zone                   *Zone       `json:"zone,omitempty" gorm:"foreignKey:zoneid;references:zoneidnumber"`
}

func (GroundSpawn) TableName() string {
    return "ground_spawns"
}

func (GroundSpawn) Relationships() []string {
    return []string{
		"Zone",
	}
}

func (GroundSpawn) Connection() string {
    return "eqemu_content"
}
