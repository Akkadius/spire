package models

import (
	"github.com/volatiletech/null/v8"
)

type PlayerEventNpcHandinEntry struct {
	ID                         uint64    `json:"id" gorm:"Column:id"`
	PlayerEventNpcHandinId     uint64    `json:"player_event_npc_handin_id" gorm:"Column:player_event_npc_handin_id"`
	Type                       null.Uint `json:"type" gorm:"Column:type"`
	ItemId                     uint      `json:"item_id" gorm:"Column:item_id"`
	Charges                    int       `json:"charges" gorm:"Column:charges"`
	EvolveLevel                uint      `json:"evolve_level" gorm:"Column:evolve_level"`
	EvolveAmount               uint64    `json:"evolve_amount" gorm:"Column:evolve_amount"`
	Augment1Id                 uint      `json:"augment_1_id" gorm:"Column:augment_1_id"`
	Augment2Id                 uint      `json:"augment_2_id" gorm:"Column:augment_2_id"`
	Augment3Id                 uint      `json:"augment_3_id" gorm:"Column:augment_3_id"`
	Augment4Id                 uint      `json:"augment_4_id" gorm:"Column:augment_4_id"`
	Augment5Id                 uint      `json:"augment_5_id" gorm:"Column:augment_5_id"`
	Augment6Id                 uint      `json:"augment_6_id" gorm:"Column:augment_6_id"`
	CreatedAt                  null.Time `json:"created_at" gorm:"Column:created_at"`
}

func (PlayerEventNpcHandinEntry) TableName() string {
    return "player_event_npc_handin_entries"
}

func (PlayerEventNpcHandinEntry) Relationships() []string {
    return []string{}
}

func (PlayerEventNpcHandinEntry) Connection() string {
    return ""
}
