package models

import (
	"github.com/volatiletech/null/v8"
)

type Report struct {
	ID            uint        `json:"id" gorm:"Column:id"`
	Name          null.String `json:"name" gorm:"Column:name"`
	Reported      null.String `json:"reported" gorm:"Column:reported"`
	ReportedText  null.String `json:"reported_text" gorm:"Column:reported_text"`
}

func (Report) TableName() string {
    return "reports"
}

func (Report) Relationships() []string {
    return []string{}
}

func (Report) Connection() string {
    return ""
}
