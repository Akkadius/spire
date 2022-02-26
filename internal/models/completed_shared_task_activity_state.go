package models

import (
	"github.com/volatiletech/null/v8"
)

type CompletedSharedTaskActivityState struct {
	SharedTaskId   int64     `json:"shared_task_id" gorm:"Column:shared_task_id"`
	ActivityId     int       `json:"activity_id" gorm:"Column:activity_id"`
	DoneCount      null.Int  `json:"done_count" gorm:"Column:done_count"`
	UpdatedTime    null.Time `json:"updated_time" gorm:"Column:updated_time"`
	CompletedTime  null.Time `json:"completed_time" gorm:"Column:completed_time"`
}

func (CompletedSharedTaskActivityState) TableName() string {
    return "completed_shared_task_activity_state"
}

func (CompletedSharedTaskActivityState) Relationships() []string {
    return []string{}
}

func (CompletedSharedTaskActivityState) Connection() string {
    return ""
}
