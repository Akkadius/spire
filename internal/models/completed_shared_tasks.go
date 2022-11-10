package models

import (
	"github.com/volatiletech/null/v8"
)

type CompletedSharedTask struct {
	ID              int64     `json:"id" gorm:"Column:id"`
	TaskId          null.Int  `json:"task_id" gorm:"Column:task_id"`
	AcceptedTime    null.Time `json:"accepted_time" gorm:"Column:accepted_time"`
	ExpireTime      null.Time `json:"expire_time" gorm:"Column:expire_time"`
	CompletionTime  null.Time `json:"completion_time" gorm:"Column:completion_time"`
	IsLocked        null.Int8 `json:"is_locked" gorm:"Column:is_locked"`
}

func (CompletedSharedTask) TableName() string {
    return "completed_shared_tasks"
}

func (CompletedSharedTask) Relationships() []string {
    return []string{}
}

func (CompletedSharedTask) Connection() string {
    return ""
}
