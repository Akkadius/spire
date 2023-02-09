package models

import (
	"time"
)

type Variable struct {
	ID          int       `json:"id" gorm:"Column:id"`
	Varname     string    `json:"varname" gorm:"Column:varname"`
	Value       string    `json:"value" gorm:"Column:value"`
	Information string    `json:"information" gorm:"Column:information"`
	Ts          time.Time `json:"ts" gorm:"Column:ts"`
}

func (Variable) TableName() string {
    return "variables"
}

func (Variable) Relationships() []string {
    return []string{}
}

func (Variable) Connection() string {
    return ""
}
