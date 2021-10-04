package models

import (
	"github.com/volatiletech/null/v8"
)

type SpawnEvent struct {
	ID          uint        `json:"id" gorm:"Column:id"`
	Zone        null.String `json:"zone" gorm:"Column:zone"`
	CondId      uint32      `json:"cond_id" gorm:"Column:cond_id"`
	Name        string      `json:"name" gorm:"Column:name"`
	Period      uint        `json:"period" gorm:"Column:period"`
	NextMinute  uint8       `json:"next_minute" gorm:"Column:next_minute"`
	NextHour    uint8       `json:"next_hour" gorm:"Column:next_hour"`
	NextDay     uint8       `json:"next_day" gorm:"Column:next_day"`
	NextMonth   uint8       `json:"next_month" gorm:"Column:next_month"`
	NextYear    uint        `json:"next_year" gorm:"Column:next_year"`
	Enabled     int8        `json:"enabled" gorm:"Column:enabled"`
	Action      uint8       `json:"action" gorm:"Column:action"`
	Argument    int32       `json:"argument" gorm:"Column:argument"`
	Strict      int8        `json:"strict" gorm:"Column:strict"`
}

func (SpawnEvent) TableName() string {
    return "spawn_events"
}

func (SpawnEvent) Relationships() []string {
    return []string{}
}

func (SpawnEvent) Connection() string {
    return ""
}
