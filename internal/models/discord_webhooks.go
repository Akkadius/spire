package models

import (
	"github.com/volatiletech/null/v8"
)

type DiscordWebhook struct {
	ID           int         `json:"id" gorm:"Column:id"`
	WebhookName  string `json:"webhook_name" gorm:"Column:webhook_name"`
	WebhookUrl   string `json:"webhook_url" gorm:"Column:webhook_url"`
	CreatedAt    time.Time   `json:"created_at" gorm:"Column:created_at"`
	DeletedAt    time.Time   `json:"deleted_at" gorm:"Column:deleted_at"`
}

func (DiscordWebhook) TableName() string {
    return "discord_webhooks"
}

func (DiscordWebhook) Relationships() []string {
    return []string{}
}

func (DiscordWebhook) Connection() string {
    return ""
}
