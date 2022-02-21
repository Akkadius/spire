package models

import (
	"time"
)

type CharacterTaskTimer struct {
	ID           uint      `json:"id" gorm:"Column:id"`
	CharacterId  uint      `json:"character_id" gorm:"Column:character_id"`
	TaskId       uint      `json:"task_id" gorm:"Column:task_id"`
	TimerType    int       `json:"timer_type" gorm:"Column:timer_type"`
	ExpireTime   time.Time `json:"expire_time" gorm:"Column:expire_time"`
}

func (CharacterTaskTimer) TableName() string {
    return "character_task_timers"
}

func (CharacterTaskTimer) Relationships() []string {
    return []string{}
}

func (CharacterTaskTimer) Connection() string {
    return ""
}
