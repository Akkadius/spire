package models

import (
	"github.com/volatiletech/null/v8"
)

type LogsysCategory struct {
	LogCategoryId            int         `json:"log_category_id" gorm:"Column:log_category_id"`
	LogCategoryDescription   null.String `json:"log_category_description" gorm:"Column:log_category_description"`
	LogToConsole             null.Int16  `json:"log_to_console" gorm:"Column:log_to_console"`
	LogToFile                null.Int16  `json:"log_to_file" gorm:"Column:log_to_file"`
	LogToGmsay               null.Int16  `json:"log_to_gmsay" gorm:"Column:log_to_gmsay"`
}

func (LogsysCategory) TableName() string {
    return "logsys_categories"
}

func (LogsysCategory) Relationships() []string {
    return []string{}
}

func (LogsysCategory) Connection() string {
    return ""
}
