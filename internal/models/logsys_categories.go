package models

import (
	"github.com/volatiletech/null/v8"
)

type LogsysCategory struct {
	LogCategoryId            int         `json:"log_category_id" gorm:"Column:log_category_id"`
	LogCategoryDescription   string `json:"log_category_description" gorm:"Column:log_category_description"`
	LogToConsole             int16  `json:"log_to_console" gorm:"Column:log_to_console"`
	LogToFile                int16  `json:"log_to_file" gorm:"Column:log_to_file"`
	LogToGmsay               int16  `json:"log_to_gmsay" gorm:"Column:log_to_gmsay"`
	LogToDiscord             int16  `json:"log_to_discord" gorm:"Column:log_to_discord"`
	DiscordWebhookId         int    `json:"discord_webhook_id" gorm:"Column:discord_webhook_id"`
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
