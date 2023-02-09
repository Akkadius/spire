package models

import (
	"github.com/volatiletech/null/v8"
)

type BotPet struct {
	PetsIndex  uint        `json:"pets_index" gorm:"Column:pets_index"`
	PetId      uint        `json:"pet_id" gorm:"Column:pet_id"`
	BotId      uint        `json:"bot_id" gorm:"Column:bot_id"`
	Name       null.String `json:"name" gorm:"Column:name"`
	Mana       int         `json:"mana" gorm:"Column:mana"`
	Hp         int         `json:"hp" gorm:"Column:hp"`
}

func (BotPet) TableName() string {
    return "bot_pets"
}

func (BotPet) Relationships() []string {
    return []string{}
}

func (BotPet) Connection() string {
    return ""
}
