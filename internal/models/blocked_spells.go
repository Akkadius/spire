package models

import (
	"github.com/volatiletech/null/v8"
)

type BlockedSpell struct {
	ID                     int         `json:"id" gorm:"Column:id"`
	Spellid                uint32      `json:"spellid" gorm:"Column:spellid"`
	Type                   int8        `json:"type" gorm:"Column:type"`
	Zoneid                 int         `json:"zoneid" gorm:"Column:zoneid"`
	X                      float32     `json:"x" gorm:"Column:x"`
	Y                      float32     `json:"y" gorm:"Column:y"`
	Z                      float32     `json:"z" gorm:"Column:z"`
	XDiff                  float32     `json:"x_diff" gorm:"Column:x_diff"`
	YDiff                  float32     `json:"y_diff" gorm:"Column:y_diff"`
	ZDiff                  float32     `json:"z_diff" gorm:"Column:z_diff"`
	Message                string      `json:"message" gorm:"Column:message"`
	Description            string      `json:"description" gorm:"Column:description"`
	MinExpansion           int8        `json:"min_expansion" gorm:"Column:min_expansion"`
	MaxExpansion           int8        `json:"max_expansion" gorm:"Column:max_expansion"`
	ContentFlags           null.String `json:"content_flags" gorm:"Column:content_flags"`
	ContentFlagsDisabled   null.String `json:"content_flags_disabled" gorm:"Column:content_flags_disabled"`
}

func (BlockedSpell) TableName() string {
    return "blocked_spells"
}

func (BlockedSpell) Relationships() []string {
    return []string{}
}

func (BlockedSpell) Connection() string {
    return "eqemu_content"
}
