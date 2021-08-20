package models

import (
	"github.com/volatiletech/null/v8"
)

type ServerScheduledEvent struct {
	ID              int         `json:"id" gorm:"Column:id"`
	Description     null.String `json:"description" gorm:"Column:description"`
	EventType       null.String `json:"event_type" gorm:"Column:event_type"`
	EventData       null.String `json:"event_data" gorm:"Column:event_data"`
	MinuteStart     null.Int    `json:"minute_start" gorm:"Column:minute_start"`
	HourStart       null.Int    `json:"hour_start" gorm:"Column:hour_start"`
	DayStart        null.Int    `json:"day_start" gorm:"Column:day_start"`
	MonthStart      null.Int    `json:"month_start" gorm:"Column:month_start"`
	YearStart       null.Int    `json:"year_start" gorm:"Column:year_start"`
	MinuteEnd       null.Int    `json:"minute_end" gorm:"Column:minute_end"`
	HourEnd         null.Int    `json:"hour_end" gorm:"Column:hour_end"`
	DayEnd          null.Int    `json:"day_end" gorm:"Column:day_end"`
	MonthEnd        null.Int    `json:"month_end" gorm:"Column:month_end"`
	YearEnd         null.Int    `json:"year_end" gorm:"Column:year_end"`
	CronExpression  null.String `json:"cron_expression" gorm:"Column:cron_expression"`
	CreatedAt       null.Time   `json:"created_at" gorm:"Column:created_at"`
	DeletedAt       null.Time   `json:"deleted_at" gorm:"Column:deleted_at"`
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
