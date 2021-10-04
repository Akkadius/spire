package models

import (
	"github.com/volatiletech/null/v8"
)

type Horse struct {
	Filename   string      `json:"filename" gorm:"Column:filename"`
	Race       int16       `json:"race" gorm:"Column:race"`
	Gender     int8        `json:"gender" gorm:"Column:gender"`
	Texture    int8        `json:"texture" gorm:"Column:texture"`
	Mountspeed float32     `json:"mountspeed" gorm:"Column:mountspeed"`
	Notes      null.String `json:"notes" gorm:"Column:notes"`
}

func (Horse) TableName() string {
    return "horses"
}

func (Horse) Relationships() []string {
    return []string{}
}

func (Horse) Connection() string {
    return "eqemu_content"
}
