package models

import (
	"github.com/volatiletech/null/v8"
)

type CharacterParcel struct {
	ID         uint        `json:"id" gorm:"Column:id"`
	CharId     uint        `json:"char_id" gorm:"Column:char_id"`
	ItemId     uint        `json:"item_id" gorm:"Column:item_id"`
	AugSlot1   uint        `json:"aug_slot_1" gorm:"Column:aug_slot_1"`
	AugSlot2   uint        `json:"aug_slot_2" gorm:"Column:aug_slot_2"`
	AugSlot3   uint        `json:"aug_slot_3" gorm:"Column:aug_slot_3"`
	AugSlot4   uint        `json:"aug_slot_4" gorm:"Column:aug_slot_4"`
	AugSlot5   uint        `json:"aug_slot_5" gorm:"Column:aug_slot_5"`
	AugSlot6   uint        `json:"aug_slot_6" gorm:"Column:aug_slot_6"`
	SlotId     uint        `json:"slot_id" gorm:"Column:slot_id"`
	Quantity   uint        `json:"quantity" gorm:"Column:quantity"`
	FromName   null.String `json:"from_name" gorm:"Column:from_name"`
	Note       null.String `json:"note" gorm:"Column:note"`
	SentDate   null.Time   `json:"sent_date" gorm:"Column:sent_date"`
}

func (CharacterParcel) TableName() string {
    return "character_parcels"
}

func (CharacterParcel) Relationships() []string {
    return []string{}
}

func (CharacterParcel) Connection() string {
    return ""
}
