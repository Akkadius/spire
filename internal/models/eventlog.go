package models

import (
	"github.com/volatiletech/null/v8"
	"time"
)

type Eventlog struct {
	ID              uint        `json:"id" gorm:"Column:id"`
	Accountname     string      `json:"accountname" gorm:"Column:accountname"`
	Accountid       null.Uint   `json:"accountid" gorm:"Column:accountid"`
	Status          int         `json:"status" gorm:"Column:status"`
	Charname        string      `json:"charname" gorm:"Column:charname"`
	Target          null.String `json:"target" gorm:"Column:target"`
	Time            time.Time   `json:"time" gorm:"Column:time"`
	Descriptiontype string      `json:"descriptiontype" gorm:"Column:descriptiontype"`
	Description     string      `json:"description" gorm:"Column:description"`
	EventNid        int         `json:"event_nid" gorm:"Column:event_nid"`
}

func (Eventlog) TableName() string {
    return "eventlog"
}

func (Eventlog) Relationships() []string {
    return []string{}
}

func (Eventlog) Connection() string {
    return ""
}
