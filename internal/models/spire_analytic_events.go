package models

import (
	"time"
)

type AnalyticEvent struct {
	ID         uint      `json:"id" gorm:"primary_key,AUTO_INCREMENT,type:bigint(11)"`
	EventName  string    `json:"event_name" gorm:"type:varchar(50);index:event_name_ip_address;index:event_name"`
	EventValue string    `json:"event_value" gorm:"type:varchar(200);"`
	RequestUri string    `json:"request_uri" gorm:"type:varchar(250);"`
	IpAddress  string    `json:"ip_address" gorm:"type:varchar(20);index:event_name_ip_address"`
	UserID     uint      `json:"user_id" gorm:"type:bigint(11);default:0"`
	CreatedAt  time.Time `json:"updated_at" gorm:"Column:updated_at"`
}

func (AnalyticEvent) TableName() string {
	return "spire_analytic_events"
}

func (AnalyticEvent) Relationships() []string {
	return []string{}
}

func (AnalyticEvent) Connection() string {
	return "spire"
}
