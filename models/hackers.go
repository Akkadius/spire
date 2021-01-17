package models

import (
	"github.com/volatiletech/null/v8"
	"time"
)

type Hacker struct {
	ID      int         `json:"id" gorm:"Column:id"`
	Account string      `json:"account" gorm:"Column:account"`
	Name    string      `json:"name" gorm:"Column:name"`
	Hacked  string      `json:"hacked" gorm:"Column:hacked"`
	Zone    null.String `json:"zone" gorm:"Column:zone"`
	Date    time.Time   `json:"date" gorm:"Column:date"`
}

func (Hacker) TableName() string {
    return "hackers"
}

func (Hacker) Relationships() []string {
    return []string{}
}

func (Hacker) Connection() string {
    return ""
}
