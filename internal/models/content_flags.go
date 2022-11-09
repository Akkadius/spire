package models

import (
	"github.com/volatiletech/null/v8"
)

type ContentFlag struct {
	ID        int         `json:"id" gorm:"Column:id"`
	FlagName  string `json:"flag_name" gorm:"Column:flag_name"`
	Enabled   int8   `json:"enabled" gorm:"Column:enabled"`
	Notes     string `json:"notes" gorm:"Column:notes"`
}

func (ContentFlag) TableName() string {
    return "content_flags"
}

func (ContentFlag) Relationships() []string {
    return []string{}
}

func (ContentFlag) Connection() string {
    return ""
}
