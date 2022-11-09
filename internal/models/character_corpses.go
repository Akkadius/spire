package models

import (
	"github.com/volatiletech/null/v8"
	"time"
)

type CharacterCorpse struct {
	ID               uint       `json:"id" gorm:"Column:id"`
	Charid           uint       `json:"charid" gorm:"Column:charid"`
	Charname         string     `json:"charname" gorm:"Column:charname"`
	ZoneId           int16      `json:"zone_id" gorm:"Column:zone_id"`
	InstanceId       uint16     `json:"instance_id" gorm:"Column:instance_id"`
	X                float32    `json:"x" gorm:"Column:x"`
	Y                float32    `json:"y" gorm:"Column:y"`
	Z                float32    `json:"z" gorm:"Column:z"`
	Heading          float32    `json:"heading" gorm:"Column:heading"`
	TimeOfDeath      time.Time  `json:"time_of_death" gorm:"Column:time_of_death"`
	GuildConsentId   uint       `json:"guild_consent_id" gorm:"Column:guild_consent_id"`
	IsRezzed         uint8 `json:"is_rezzed" gorm:"Column:is_rezzed"`
	IsBuried         int8       `json:"is_buried" gorm:"Column:is_buried"`
	WasAtGraveyard   int8       `json:"was_at_graveyard" gorm:"Column:was_at_graveyard"`
	IsLocked         int8  `json:"is_locked" gorm:"Column:is_locked"`
	Exp              uint  `json:"exp" gorm:"Column:exp"`
	Size             uint  `json:"size" gorm:"Column:size"`
	Level            uint  `json:"level" gorm:"Column:level"`
	Race             uint  `json:"race" gorm:"Column:race"`
	Gender           uint  `json:"gender" gorm:"Column:gender"`
	Class            uint  `json:"class" gorm:"Column:class"`
	Deity            uint  `json:"deity" gorm:"Column:deity"`
	Texture          uint  `json:"texture" gorm:"Column:texture"`
	HelmTexture      uint  `json:"helm_texture" gorm:"Column:helm_texture"`
	Copper           uint  `json:"copper" gorm:"Column:copper"`
	Silver           uint  `json:"silver" gorm:"Column:silver"`
	Gold             uint  `json:"gold" gorm:"Column:gold"`
	Platinum         uint  `json:"platinum" gorm:"Column:platinum"`
	HairColor        uint  `json:"hair_color" gorm:"Column:hair_color"`
	BeardColor       uint  `json:"beard_color" gorm:"Column:beard_color"`
	EyeColor1        uint  `json:"eye_color_1" gorm:"Column:eye_color_1"`
	EyeColor2        uint  `json:"eye_color_2" gorm:"Column:eye_color_2"`
	HairStyle        uint  `json:"hair_style" gorm:"Column:hair_style"`
	Face             uint  `json:"face" gorm:"Column:face"`
	Beard            uint  `json:"beard" gorm:"Column:beard"`
	DrakkinHeritage  uint  `json:"drakkin_heritage" gorm:"Column:drakkin_heritage"`
	DrakkinTattoo    uint  `json:"drakkin_tattoo" gorm:"Column:drakkin_tattoo"`
	DrakkinDetails   uint  `json:"drakkin_details" gorm:"Column:drakkin_details"`
	Wc1              uint  `json:"wc_1" gorm:"Column:wc_1"`
	Wc2              uint  `json:"wc_2" gorm:"Column:wc_2"`
	Wc3              uint  `json:"wc_3" gorm:"Column:wc_3"`
	Wc4              uint  `json:"wc_4" gorm:"Column:wc_4"`
	Wc5              uint  `json:"wc_5" gorm:"Column:wc_5"`
	Wc6              uint  `json:"wc_6" gorm:"Column:wc_6"`
	Wc7              uint  `json:"wc_7" gorm:"Column:wc_7"`
	Wc8              uint  `json:"wc_8" gorm:"Column:wc_8"`
	Wc9              uint  `json:"wc_9" gorm:"Column:wc_9"`
}

func (CharacterCorpse) TableName() string {
    return "character_corpses"
}

func (CharacterCorpse) Relationships() []string {
    return []string{}
}

func (CharacterCorpse) Connection() string {
    return ""
}
