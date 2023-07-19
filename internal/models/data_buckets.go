package models

import (
	"github.com/volatiletech/null/v8"
)

type DataBucket struct {
	ID           uint64      `json:"id" gorm:"Column:id"`
	Key          null.String `json:"key" gorm:"Column:key"`
	Value        null.String `json:"value" gorm:"Column:value"`
	Expires      null.Uint   `json:"expires" gorm:"Column:expires"`
	CharacterId  int64       `json:"character_id" gorm:"Column:character_id"`
	NpcId        int64       `json:"npc_id" gorm:"Column:npc_id"`
	BotId        int64       `json:"bot_id" gorm:"Column:bot_id"`
}

func (DataBucket) TableName() string {
    return "data_buckets"
}

func (DataBucket) Relationships() []string {
    return []string{}
}

func (DataBucket) Connection() string {
    return ""
}
