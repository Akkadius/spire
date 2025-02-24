package models

import (
	"github.com/volatiletech/null/v8"
)

type PlayerEventTradeEntry struct {
	ID                    uint64      `json:"id" gorm:"Column:id"`
	PlayerEventTradeId    null.Uint64 `json:"player_event_trade_id" gorm:"Column:player_event_trade_id"`
	CharId                null.Uint   `json:"char_id" gorm:"Column:char_id"`
	Slot                  null.Int16  `json:"slot" gorm:"Column:slot"`
	ItemId                null.Uint   `json:"item_id" gorm:"Column:item_id"`
	Charges               null.Int16  `json:"charges" gorm:"Column:charges"`
	Augment1Id            null.Uint   `json:"augment_1_id" gorm:"Column:augment_1_id"`
	Augment2Id            null.Uint   `json:"augment_2_id" gorm:"Column:augment_2_id"`
	Augment3Id            null.Uint   `json:"augment_3_id" gorm:"Column:augment_3_id"`
	Augment4Id            null.Uint   `json:"augment_4_id" gorm:"Column:augment_4_id"`
	Augment5Id            null.Uint   `json:"augment_5_id" gorm:"Column:augment_5_id"`
	Augment6Id            null.Uint   `json:"augment_6_id" gorm:"Column:augment_6_id"`
	InBag                 null.Int8   `json:"in_bag" gorm:"Column:in_bag"`
	CreatedAt             null.Time   `json:"created_at" gorm:"Column:created_at"`
}

func (PlayerEventTradeEntry) TableName() string {
    return "player_event_trade_entries"
}

func (PlayerEventTradeEntry) Relationships() []string {
    return []string{}
}

func (PlayerEventTradeEntry) Connection() string {
    return ""
}
