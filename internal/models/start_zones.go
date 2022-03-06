package models

import (
	"github.com/volatiletech/null/v8"
)

type StartZone struct {
	X                      float32     `json:"x" gorm:"Column:x"`
	Y                      float32     `json:"y" gorm:"Column:y"`
	Z                      float32     `json:"z" gorm:"Column:z"`
	Heading                float32     `json:"heading" gorm:"Column:heading"`
	ZoneId                 int         `json:"zone_id" gorm:"Column:zone_id"`
	BindId                 int         `json:"bind_id" gorm:"Column:bind_id"`
	PlayerChoice           int         `json:"player_choice" gorm:"Column:player_choice"`
	PlayerClass            int         `json:"player_class" gorm:"Column:player_class"`
	PlayerDeity            int         `json:"player_deity" gorm:"Column:player_deity"`
	PlayerRace             int         `json:"player_race" gorm:"Column:player_race"`
	StartZone              int         `json:"start_zone" gorm:"Column:start_zone"`
	BindX                  float32     `json:"bind_x" gorm:"Column:bind_x"`
	BindY                  float32     `json:"bind_y" gorm:"Column:bind_y"`
	BindZ                  float32     `json:"bind_z" gorm:"Column:bind_z"`
	SelectRank             uint8       `json:"select_rank" gorm:"Column:select_rank"`
	MinExpansion           int8        `json:"min_expansion" gorm:"Column:min_expansion"`
	MaxExpansion           int8        `json:"max_expansion" gorm:"Column:max_expansion"`
	ContentFlags           null.String `json:"content_flags" gorm:"Column:content_flags"`
	ContentFlagsDisabled   null.String `json:"content_flags_disabled" gorm:"Column:content_flags_disabled"`
	Zone                   *Zone       `json:"zone,omitempty" gorm:"foreignKey:zone_id;references:zoneidnumber"`
}

func (StartZone) TableName() string {
    return "start_zones"
}

func (StartZone) Relationships() []string {
    return []string{
		"Zone",
	}
}

func (StartZone) Connection() string {
    return "eqemu_content"
}
