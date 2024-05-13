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
	IsRezzed         null.Uint8 `json:"is_rezzed" gorm:"Column:is_rezzed"`
	IsBuried         int8       `json:"is_buried" gorm:"Column:is_buried"`
	WasAtGraveyard   int8       `json:"was_at_graveyard" gorm:"Column:was_at_graveyard"`
	IsLocked         null.Int8  `json:"is_locked" gorm:"Column:is_locked"`
	Exp              null.Uint  `json:"exp" gorm:"Column:exp"`
	Size             null.Uint  `json:"size" gorm:"Column:size"`
	Level            null.Uint  `json:"level" gorm:"Column:level"`
	Race             null.Uint  `json:"race" gorm:"Column:race"`
	Gender           null.Uint  `json:"gender" gorm:"Column:gender"`
	Class            null.Uint  `json:"class" gorm:"Column:class"`
	Deity            null.Uint  `json:"deity" gorm:"Column:deity"`
	Texture          null.Uint  `json:"texture" gorm:"Column:texture"`
	HelmTexture      null.Uint  `json:"helm_texture" gorm:"Column:helm_texture"`
	Copper           null.Uint  `json:"copper" gorm:"Column:copper"`
	Silver           null.Uint  `json:"silver" gorm:"Column:silver"`
	Gold             null.Uint  `json:"gold" gorm:"Column:gold"`
	Platinum         null.Uint  `json:"platinum" gorm:"Column:platinum"`
	HairColor        null.Uint  `json:"hair_color" gorm:"Column:hair_color"`
	BeardColor       null.Uint  `json:"beard_color" gorm:"Column:beard_color"`
	EyeColor1        null.Uint  `json:"eye_color_1" gorm:"Column:eye_color_1"`
	EyeColor2        null.Uint  `json:"eye_color_2" gorm:"Column:eye_color_2"`
	HairStyle        null.Uint  `json:"hair_style" gorm:"Column:hair_style"`
	Face             null.Uint  `json:"face" gorm:"Column:face"`
	Beard            null.Uint  `json:"beard" gorm:"Column:beard"`
	DrakkinHeritage  null.Uint  `json:"drakkin_heritage" gorm:"Column:drakkin_heritage"`
	DrakkinTattoo    null.Uint  `json:"drakkin_tattoo" gorm:"Column:drakkin_tattoo"`
	DrakkinDetails   null.Uint  `json:"drakkin_details" gorm:"Column:drakkin_details"`
	Wc1              null.Uint  `json:"wc_1" gorm:"Column:wc_1"`
	Wc2              null.Uint  `json:"wc_2" gorm:"Column:wc_2"`
	Wc3              null.Uint  `json:"wc_3" gorm:"Column:wc_3"`
	Wc4              null.Uint  `json:"wc_4" gorm:"Column:wc_4"`
	Wc5              null.Uint  `json:"wc_5" gorm:"Column:wc_5"`
	Wc6              null.Uint  `json:"wc_6" gorm:"Column:wc_6"`
	Wc7              null.Uint  `json:"wc_7" gorm:"Column:wc_7"`
	Wc8              null.Uint  `json:"wc_8" gorm:"Column:wc_8"`
	Wc9              null.Uint  `json:"wc_9" gorm:"Column:wc_9"`
	RezTime          uint       `json:"rez_time" gorm:"Column:rez_time"`
	GmExp            uint       `json:"gm_exp" gorm:"Column:gm_exp"`
	KilledBy         uint       `json:"killed_by" gorm:"Column:killed_by"`
	Rezzable         uint8      `json:"rezzable" gorm:"Column:rezzable"`
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
