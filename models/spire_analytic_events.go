package models

import (
	"time"
)

type AnalyticEvent struct {
	ID         uint      `json:"id" gorm:"primary_key,AUTO_INCREMENT,type:bigint(11)"`
	EventName  string    `json:"event_name" gorm:"type:varchar(50);"`
	EventValue string    `json:"event_value" gorm:"type:varchar(200);"`
	RequestUri string    `json:"request_uri" gorm:"type:varchar(250);"`
	IpAddress  string    `json:"ip_address" gorm:"type:varchar(20);"`
	UserID     uint      `json:"user_id" gorm:"type:bigint(11);default:0"`
	CreatedAt  time.Time `json:"updated_at" gorm:"Column:updated_at"`
}

func (AnalyticEvent) TableName() string {
	return "analytic_events"
}

func (AnalyticEvent) Relationships() []string {
	return []string{}
}

func (AnalyticEvent) Connection() string {
	return "spire"
}

func (AnalyticEvent) Indexes() map[string][]string {
	return map[string][]string{
		"event_name":            {"event_name"},
		"event_name_ip_address": {"event_name", "ip_address"},
	}
}
