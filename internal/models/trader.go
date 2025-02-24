package models

import (
	"github.com/volatiletech/null/v8"
)

type Trader struct {
	ID                    uint64    `json:"id" gorm:"Column:id"`
	CharId                uint      `json:"char_id" gorm:"Column:char_id"`
	ItemId                uint      `json:"item_id" gorm:"Column:item_id"`
	AugSlot1              uint      `json:"aug_slot_1" gorm:"Column:aug_slot_1"`
	AugSlot2              uint      `json:"aug_slot_2" gorm:"Column:aug_slot_2"`
	AugSlot3              uint      `json:"aug_slot_3" gorm:"Column:aug_slot_3"`
	AugSlot4              uint      `json:"aug_slot_4" gorm:"Column:aug_slot_4"`
	AugSlot5              uint      `json:"aug_slot_5" gorm:"Column:aug_slot_5"`
	AugSlot6              uint      `json:"aug_slot_6" gorm:"Column:aug_slot_6"`
	ItemSn                uint      `json:"item_sn" gorm:"Column:item_sn"`
	ItemCharges           int       `json:"item_charges" gorm:"Column:item_charges"`
	ItemCost              uint      `json:"item_cost" gorm:"Column:item_cost"`
	SlotId                uint8     `json:"slot_id" gorm:"Column:slot_id"`
	CharEntityId          uint      `json:"char_entity_id" gorm:"Column:char_entity_id"`
	CharZoneId            uint      `json:"char_zone_id" gorm:"Column:char_zone_id"`
	CharZoneInstanceId    null.Int  `json:"char_zone_instance_id" gorm:"Column:char_zone_instance_id"`
	ActiveTransaction     uint8     `json:"active_transaction" gorm:"Column:active_transaction"`
	ListingDate           null.Time `json:"listing_date" gorm:"Column:listing_date"`
}

func (Trader) TableName() string {
    return "trader"
}

func (Trader) Relationships() []string {
    return []string{}
}

func (Trader) Connection() string {
    return ""
}
