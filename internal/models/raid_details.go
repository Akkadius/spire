package models

import (
	"github.com/volatiletech/null/v8"
)

type RaidDetail struct {
	Raidid       int          `json:"raidid" gorm:"Column:raidid"`
	Loottype     int          `json:"loottype" gorm:"Column:loottype"`
	Locked       int8         `json:"locked" gorm:"Column:locked"`
	Motd         null.String  `json:"motd" gorm:"Column:motd"`
	MarkedNpc1   uint16       `json:"marked_npc_1" gorm:"Column:marked_npc_1"`
	MarkedNpc2   uint16       `json:"marked_npc_2" gorm:"Column:marked_npc_2"`
	MarkedNpc3   uint16       `json:"marked_npc_3" gorm:"Column:marked_npc_3"`
	MarkedNPC1   uint16       `json:"marked_npc_1" gorm:"Column:markedNPC1"`
	MarkedNPC2   uint16       `json:"marked_npc_2" gorm:"Column:markedNPC2"`
	MarkedNPC3   uint16       `json:"marked_npc_3" gorm:"Column:markedNPC3"`
	RaidLeaders  []RaidLeader `json:"raid_leaders,omitempty" gorm:"foreignKey:rid;references:raidid"`
	RaidMembers  []RaidMember `json:"raid_members,omitempty" gorm:"foreignKey:raidid;references:raidid"`
}

func (RaidDetail) TableName() string {
    return "raid_details"
}

func (RaidDetail) Relationships() []string {
    return []string{
		"RaidLeaders",
		"RaidMembers",
	}
}

func (RaidDetail) Connection() string {
    return ""
}
