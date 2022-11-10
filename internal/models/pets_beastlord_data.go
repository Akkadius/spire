package models

import (
	"github.com/volatiletech/null/v8"
)

type PetsBeastlordDatum struct {
	PlayerRace    uint        `json:"player_race" gorm:"Column:player_race"`
	PetRace       uint        `json:"pet_race" gorm:"Column:pet_race"`
	Texture       uint8       `json:"texture" gorm:"Column:texture"`
	HelmTexture   uint8       `json:"helm_texture" gorm:"Column:helm_texture"`
	Gender        uint8       `json:"gender" gorm:"Column:gender"`
	SizeModifier  null.String `json:"size_modifier" gorm:"Column:size_modifier"`
	Face          uint8       `json:"face" gorm:"Column:face"`
}

func (PetsBeastlordDatum) TableName() string {
    return "pets_beastlord_data"
}

func (PetsBeastlordDatum) Relationships() []string {
    return []string{}
}

func (PetsBeastlordDatum) Connection() string {
    return ""
}
