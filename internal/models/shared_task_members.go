package models

import (
	"github.com/volatiletech/null/v8"
)

type SharedTaskMember struct {
	SharedTaskId   int64     `json:"shared_task_id" gorm:"Column:shared_task_id"`
	CharacterId    int64     `json:"character_id" gorm:"Column:character_id"`
	IsLeader       null.Int8 `json:"is_leader" gorm:"Column:is_leader"`
}

func (SharedTaskMember) TableName() string {
    return "shared_task_members"
}

func (SharedTaskMember) Relationships() []string {
    return []string{}
}

func (SharedTaskMember) Connection() string {
    return ""
}
