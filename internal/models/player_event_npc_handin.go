package models

import (
	"github.com/volatiletech/null/v8"
)

type PlayerEventNpcHandin struct {
	ID              uint64      `json:"id" gorm:"Column:id"`
	NpcId           null.Uint   `json:"npc_id" gorm:"Column:npc_id"`
	NpcName         null.String `json:"npc_name" gorm:"Column:npc_name"`
	HandinCopper    null.Uint64 `json:"handin_copper" gorm:"Column:handin_copper"`
	HandinSilver    null.Uint64 `json:"handin_silver" gorm:"Column:handin_silver"`
	HandinGold      null.Uint64 `json:"handin_gold" gorm:"Column:handin_gold"`
	HandinPlatinum  null.Uint64 `json:"handin_platinum" gorm:"Column:handin_platinum"`
	ReturnCopper    null.Uint64 `json:"return_copper" gorm:"Column:return_copper"`
	ReturnSilver    null.Uint64 `json:"return_silver" gorm:"Column:return_silver"`
	ReturnGold      null.Uint64 `json:"return_gold" gorm:"Column:return_gold"`
	ReturnPlatinum  null.Uint64 `json:"return_platinum" gorm:"Column:return_platinum"`
	IsQuestHandin   null.Uint8  `json:"is_quest_handin" gorm:"Column:is_quest_handin"`
	CreatedAt       null.Time   `json:"created_at" gorm:"Column:created_at"`
}

func (PlayerEventNpcHandin) TableName() string {
    return "player_event_npc_handin"
}

func (PlayerEventNpcHandin) Relationships() []string {
    return []string{}
}

func (PlayerEventNpcHandin) Connection() string {
    return ""
}
