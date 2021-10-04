package models

import (
	"github.com/volatiletech/null/v8"
)

type GlobalLoot struct {
	ID                     int         `json:"id" gorm:"Column:id"`
	Description            null.String `json:"description" gorm:"Column:description"`
	LoottableId            int         `json:"loottable_id" gorm:"Column:loottable_id"`
	Enabled                int8        `json:"enabled" gorm:"Column:enabled"`
	MinLevel               int         `json:"min_level" gorm:"Column:min_level"`
	MaxLevel               int         `json:"max_level" gorm:"Column:max_level"`
	Rare                   null.Int8   `json:"rare" gorm:"Column:rare"`
	Raid                   null.Int8   `json:"raid" gorm:"Column:raid"`
	Race                   null.String `json:"race" gorm:"Column:race"`
	Class                  null.String `json:"class" gorm:"Column:class"`
	Bodytype               null.String `json:"bodytype" gorm:"Column:bodytype"`
	Zone                   null.String `json:"zone" gorm:"Column:zone"`
	HotZone                null.Int8   `json:"hot_zone" gorm:"Column:hot_zone"`
	MinExpansion           uint8       `json:"min_expansion" gorm:"Column:min_expansion"`
	MaxExpansion           uint8       `json:"max_expansion" gorm:"Column:max_expansion"`
	ContentFlags           null.String `json:"content_flags" gorm:"Column:content_flags"`
	ContentFlagsDisabled   null.String `json:"content_flags_disabled" gorm:"Column:content_flags_disabled"`
}

func (GlobalLoot) TableName() string {
    return "global_loot"
}

func (GlobalLoot) Relationships() []string {
    return []string{}
}

func (GlobalLoot) Connection() string {
    return "eqemu_content"
}
