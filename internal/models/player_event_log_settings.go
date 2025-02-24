package models

import (
	"github.com/volatiletech/null/v8"
)

type PlayerEventLogSetting struct {
	ID                 int64       `json:"id" gorm:"Column:id"`
	EventName          null.String `json:"event_name" gorm:"Column:event_name"`
	EventEnabled       null.Int8   `json:"event_enabled" gorm:"Column:event_enabled"`
	RetentionDays      null.Int    `json:"retention_days" gorm:"Column:retention_days"`
	DiscordWebhookId   null.Int    `json:"discord_webhook_id" gorm:"Column:discord_webhook_id"`
	EtlEnabled         uint8       `json:"etl_enabled" gorm:"Column:etl_enabled"`
}

func (PlayerEventLogSetting) TableName() string {
    return "player_event_log_settings"
}

func (PlayerEventLogSetting) Relationships() []string {
    return []string{}
}

func (PlayerEventLogSetting) Connection() string {
    return "eqemu_logs"
}
