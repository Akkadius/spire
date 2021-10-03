package models

import (
	"github.com/volatiletech/null/v8"
)

type CharacterCorpseItem struct {
	CorpseId   uint      `json:"corpse_id" gorm:"Column:corpse_id"`
	EquipSlot  uint      `json:"equip_slot" gorm:"Column:equip_slot"`
	ItemId     null.Uint `json:"item_id" gorm:"Column:item_id"`
	Charges    null.Uint `json:"charges" gorm:"Column:charges"`
	Aug1       null.Uint `json:"aug_1" gorm:"Column:aug_1"`
	Aug2       null.Uint `json:"aug_2" gorm:"Column:aug_2"`
	Aug3       null.Uint `json:"aug_3" gorm:"Column:aug_3"`
	Aug4       null.Uint `json:"aug_4" gorm:"Column:aug_4"`
	Aug5       null.Uint `json:"aug_5" gorm:"Column:aug_5"`
	Aug6       int       `json:"aug_6" gorm:"Column:aug_6"`
	Attuned    int16     `json:"attuned" gorm:"Column:attuned"`
}

func (CharacterCorpseItem) TableName() string {
    return "character_corpse_items"
}

func (CharacterCorpseItem) Relationships() []string {
    return []string{}
}

func (CharacterCorpseItem) Connection() string {
    return ""
}
