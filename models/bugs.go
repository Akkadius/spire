package models

import (
	"github.com/volatiletech/null/v8"
	"time"
)

type Bug struct {
	ID     uint        `json:"id" gorm:"Column:id"`
	Zone   string      `json:"zone" gorm:"Column:zone"`
	Name   string      `json:"name" gorm:"Column:name"`
	Ui     string      `json:"ui" gorm:"Column:ui"`
	X      float32     `json:"x" gorm:"Column:x"`
	Y      float32     `json:"y" gorm:"Column:y"`
	Z      float32     `json:"z" gorm:"Column:z"`
	Type   string      `json:"type" gorm:"Column:type"`
	Flag   uint8       `json:"flag" gorm:"Column:flag"`
	Target null.String `json:"target" gorm:"Column:target"`
	Bug    string      `json:"bug" gorm:"Column:bug"`
	Date   time.Time   `json:"date" gorm:"Column:date"`
	Status uint8       `json:"status" gorm:"Column:status"`
}

func (Bug) TableName() string {
    return "bugs"
}

func (Bug) Relationships() []string {
    return []string{}
}

func (Bug) Connection() string {
    return ""
}
