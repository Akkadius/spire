package models

import (
	"github.com/volatiletech/null/v8"
)

type ItemsEvolvingDetail struct {
	ID                uint        `json:"id" gorm:"Column:id"`
	ItemEvoId         null.Uint   `json:"item_evo_id" gorm:"Column:item_evo_id"`
	ItemEvolveLevel   null.Uint   `json:"item_evolve_level" gorm:"Column:item_evolve_level"`
	ItemId            null.Uint   `json:"item_id" gorm:"Column:item_id"`
	Type              null.Uint   `json:"type" gorm:"Column:type"`
	SubType           null.String `json:"sub_type" gorm:"Column:sub_type"`
	RequiredAmount    null.Int64  `json:"required_amount" gorm:"Column:required_amount"`
}

func (ItemsEvolvingDetail) TableName() string {
    return "items_evolving_details"
}

func (ItemsEvolvingDetail) Relationships() []string {
    return []string{}
}

func (ItemsEvolvingDetail) Connection() string {
    return ""
}
