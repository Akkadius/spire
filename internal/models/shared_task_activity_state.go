package models

import (
	"github.com/volatiletech/null/v8"
)

type SharedTaskActivityState struct {
	SharedTaskId   int64     `json:"shared_task_id" gorm:"Column:shared_task_id"`
	ActivityId     int       `json:"activity_id" gorm:"Column:activity_id"`
	DoneCount      null.Int  `json:"done_count" gorm:"Column:done_count"`
	UpdatedTime    null.Time `json:"updated_time" gorm:"Column:updated_time"`
	CompletedTime  null.Time `json:"completed_time" gorm:"Column:completed_time"`
}

func (SharedTaskActivityState) TableName() string {
    return "shared_task_activity_state"
}

func (SharedTaskActivityState) Relationships() []string {
    return []string{}
}

func (SharedTaskActivityState) Connection() string {
    return ""
}
