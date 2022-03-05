package models

import (
	"github.com/volatiletech/null/v8"
)

type Trap struct {
	ID                     int         `json:"id" gorm:"Column:id"`
	Zone                   string      `json:"zone" gorm:"Column:zone"`
	Version                uint16      `json:"version" gorm:"Column:version"`
	X                      int         `json:"x" gorm:"Column:x"`
	Y                      int         `json:"y" gorm:"Column:y"`
	Z                      int         `json:"z" gorm:"Column:z"`
	Chance                 int8        `json:"chance" gorm:"Column:chance"`
	Maxzdiff               float32     `json:"maxzdiff" gorm:"Column:maxzdiff"`
	Radius                 float32     `json:"radius" gorm:"Column:radius"`
	Effect                 int         `json:"effect" gorm:"Column:effect"`
	Effectvalue            int         `json:"effectvalue" gorm:"Column:effectvalue"`
	Effectvalue2           int         `json:"effectvalue_2" gorm:"Column:effectvalue2"`
	Message                string      `json:"message" gorm:"Column:message"`
	Skill                  int         `json:"skill" gorm:"Column:skill"`
	Level                  uint32      `json:"level" gorm:"Column:level"`
	RespawnTime            uint        `json:"respawn_time" gorm:"Column:respawn_time"`
	RespawnVar             uint        `json:"respawn_var" gorm:"Column:respawn_var"`
	TriggeredNumber        int8        `json:"triggered_number" gorm:"Column:triggered_number"`
	Group                  int8        `json:"group" gorm:"Column:group"`
	DespawnWhenTriggered   int8        `json:"despawn_when_triggered" gorm:"Column:despawn_when_triggered"`
	Undetectable           int8        `json:"undetectable" gorm:"Column:undetectable"`
	MinExpansion           int8        `json:"min_expansion" gorm:"Column:min_expansion"`
	MaxExpansion           int8        `json:"max_expansion" gorm:"Column:max_expansion"`
	ContentFlags           null.String `json:"content_flags" gorm:"Column:content_flags"`
	ContentFlagsDisabled   null.String `json:"content_flags_disabled" gorm:"Column:content_flags_disabled"`
}

func (Trap) TableName() string {
    return "traps"
}

func (Trap) Relationships() []string {
    return []string{}
}

func (Trap) Connection() string {
    return "eqemu_content"
}
