package models

import (
	"github.com/volatiletech/null/v8"
)

type RaidDetail struct {
	Raidid      int          `json:"raidid" gorm:"Column:raidid"`
	Loottype    int          `json:"loottype" gorm:"Column:loottype"`
	Locked      int8         `json:"locked" gorm:"Column:locked"`
	Motd        null.String  `json:"motd" gorm:"Column:motd"`
	RaidLeaders []RaidLeader `json:"raid_leaders,omitempty" gorm:"foreignKey:rid;references:raidid"`
	RaidMembers []RaidMember `json:"raid_members,omitempty" gorm:"foreignKey:raidid;references:raidid"`
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
