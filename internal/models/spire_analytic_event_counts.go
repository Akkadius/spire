package models

import (
	"time"
)

type AnalyticEventCount struct {
	ID        uint      `json:"id" gorm:"primary_key,AUTO_INCREMENT,type:bigint(11)"`
	EventName string    `json:"event_name" gorm:"type:varchar(50);index:event_name;index:event_name_key"`
	EventKey  string    `json:"event_key" gorm:"type:varchar(120);index:event_name_key"`
	Count     uint64    `json:"count" gorm:"type:bigint(11);"`
	UpdatedAt time.Time `json:"updated_at" gorm:"Column:updated_at"`
}

func (AnalyticEventCount) TableName() string {
	return "spire_analytic_event_counts"
}

func (AnalyticEventCount) Relationships() []string {
	return []string{}
}

func (AnalyticEventCount) Connection() string {
	return "spire"
}
