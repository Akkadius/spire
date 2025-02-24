package models

import (
	"github.com/volatiletech/null/v8"
)

type DataBucket struct {
	ID           uint64      `json:"id" gorm:"Column:id"`
	Key          null.String `json:"key" gorm:"Column:key"`
	Value        null.String `json:"value" gorm:"Column:value"`
	Expires      null.Uint   `json:"expires" gorm:"Column:expires"`
	AccountId    null.Uint64 `json:"account_id" gorm:"Column:account_id"`
	CharacterId  uint64      `json:"character_id" gorm:"Column:character_id"`
	NpcId        uint        `json:"npc_id" gorm:"Column:npc_id"`
	BotId        uint        `json:"bot_id" gorm:"Column:bot_id"`
	ZoneId       uint16      `json:"zone_id" gorm:"Column:zone_id"`
	InstanceId   uint16      `json:"instance_id" gorm:"Column:instance_id"`
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
