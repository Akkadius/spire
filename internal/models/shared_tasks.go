package models

import (
	"github.com/volatiletech/null/v8"
)

type SharedTask struct {
	ID              int64     `json:"id" gorm:"Column:id"`
	TaskId          int  `json:"task_id" gorm:"Column:task_id"`
	AcceptedTime    time.Time `json:"accepted_time" gorm:"Column:accepted_time"`
	ExpireTime      time.Time `json:"expire_time" gorm:"Column:expire_time"`
	CompletionTime  time.Time `json:"completion_time" gorm:"Column:completion_time"`
	IsLocked        int8 `json:"is_locked" gorm:"Column:is_locked"`
}

func (SharedTask) TableName() string {
    return "shared_tasks"
}

func (SharedTask) Relationships() []string {
    return []string{}
}

func (SharedTask) Connection() string {
    return ""
}
