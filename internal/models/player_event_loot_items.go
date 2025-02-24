package models

import (
	"github.com/volatiletech/null/v8"
)

type PlayerEventLootItem struct {
	ID           uint64      `json:"id" gorm:"Column:id"`
	ItemId       null.Uint   `json:"item_id" gorm:"Column:item_id"`
	ItemName     null.String `json:"item_name" gorm:"Column:item_name"`
	Charges      null.Int    `json:"charges" gorm:"Column:charges"`
	Augment1Id   null.Uint   `json:"augment_1_id" gorm:"Column:augment_1_id"`
	Augment2Id   null.Uint   `json:"augment_2_id" gorm:"Column:augment_2_id"`
	Augment3Id   null.Uint   `json:"augment_3_id" gorm:"Column:augment_3_id"`
	Augment4Id   null.Uint   `json:"augment_4_id" gorm:"Column:augment_4_id"`
	Augment5Id   null.Uint   `json:"augment_5_id" gorm:"Column:augment_5_id"`
	Augment6Id   null.Uint   `json:"augment_6_id" gorm:"Column:augment_6_id"`
	NpcId        null.Uint   `json:"npc_id" gorm:"Column:npc_id"`
	CorpseName   null.String `json:"corpse_name" gorm:"Column:corpse_name"`
	CreatedAt    null.Time   `json:"created_at" gorm:"Column:created_at"`
}

func (PlayerEventLootItem) TableName() string {
    return "player_event_loot_items"
}

func (PlayerEventLootItem) Relationships() []string {
    return []string{}
}

func (PlayerEventLootItem) Connection() string {
    return ""
}
