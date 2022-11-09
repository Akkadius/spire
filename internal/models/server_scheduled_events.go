package models

import (
	"github.com/volatiletech/null/v8"
)

type ServerScheduledEvent struct {
	ID              int         `json:"id" gorm:"Column:id"`
	Description     string `json:"description" gorm:"Column:description"`
	EventType       string `json:"event_type" gorm:"Column:event_type"`
	EventData       string `json:"event_data" gorm:"Column:event_data"`
	MinuteStart     int    `json:"minute_start" gorm:"Column:minute_start"`
	HourStart       int    `json:"hour_start" gorm:"Column:hour_start"`
	DayStart        int    `json:"day_start" gorm:"Column:day_start"`
	MonthStart      int    `json:"month_start" gorm:"Column:month_start"`
	YearStart       int    `json:"year_start" gorm:"Column:year_start"`
	MinuteEnd       int    `json:"minute_end" gorm:"Column:minute_end"`
	HourEnd         int    `json:"hour_end" gorm:"Column:hour_end"`
	DayEnd          int    `json:"day_end" gorm:"Column:day_end"`
	MonthEnd        int    `json:"month_end" gorm:"Column:month_end"`
	YearEnd         int    `json:"year_end" gorm:"Column:year_end"`
	CronExpression  string `json:"cron_expression" gorm:"Column:cron_expression"`
	CreatedAt       time.Time   `json:"created_at" gorm:"Column:created_at"`
	DeletedAt       time.Time   `json:"deleted_at" gorm:"Column:deleted_at"`
}

func (ServerScheduledEvent) TableName() string {
    return "server_scheduled_events"
}

func (ServerScheduledEvent) Relationships() []string {
    return []string{}
}

func (ServerScheduledEvent) Connection() string {
    return ""
}
