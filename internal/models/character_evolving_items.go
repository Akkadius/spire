package models

import (
	"github.com/volatiletech/null/v8"
)

type CharacterEvolvingItem struct {
	ID             uint64       `json:"id" gorm:"Column:id"`
	CharacterId    null.Uint    `json:"character_id" gorm:"Column:character_id"`
	ItemId         null.Uint    `json:"item_id" gorm:"Column:item_id"`
	Activated      null.Uint8   `json:"activated" gorm:"Column:activated"`
	Equipped       null.Uint8   `json:"equipped" gorm:"Column:equipped"`
	CurrentAmount  null.Int64   `json:"current_amount" gorm:"Column:current_amount"`
	Progression    null.Float64 `json:"progression" gorm:"Column:progression"`
	FinalItemId    null.Uint    `json:"final_item_id" gorm:"Column:final_item_id"`
	DeletedAt      null.Time    `json:"deleted_at" gorm:"Column:deleted_at"`
}

func (CharacterEvolvingItem) TableName() string {
    return "character_evolving_items"
}

func (CharacterEvolvingItem) Relationships() []string {
    return []string{}
}

func (CharacterEvolvingItem) Connection() string {
    return ""
}
