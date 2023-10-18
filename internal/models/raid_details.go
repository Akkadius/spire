package models

import (
	"github.com/volatiletech/null/v8"
)

type RaidDetail struct {
	Raidid                   int          `json:"raidid" gorm:"Column:raidid"`
	Loottype                 int          `json:"loottype" gorm:"Column:loottype"`
	Locked                   int8         `json:"locked" gorm:"Column:locked"`
	Motd                     null.String  `json:"motd" gorm:"Column:motd"`
	MarkedNpc1EntityId       uint         `json:"marked_npc_1_entity_id" gorm:"Column:marked_npc_1_entity_id"`
	MarkedNpc1ZoneId         uint         `json:"marked_npc_1_zone_id" gorm:"Column:marked_npc_1_zone_id"`
	MarkedNpc1InstanceId     uint         `json:"marked_npc_1_instance_id" gorm:"Column:marked_npc_1_instance_id"`
	MarkedNpc2EntityId       uint         `json:"marked_npc_2_entity_id" gorm:"Column:marked_npc_2_entity_id"`
	MarkedNpc2ZoneId         uint         `json:"marked_npc_2_zone_id" gorm:"Column:marked_npc_2_zone_id"`
	MarkedNpc2InstanceId     uint         `json:"marked_npc_2_instance_id" gorm:"Column:marked_npc_2_instance_id"`
	MarkedNpc3EntityId       uint         `json:"marked_npc_3_entity_id" gorm:"Column:marked_npc_3_entity_id"`
	MarkedNpc3ZoneId         uint         `json:"marked_npc_3_zone_id" gorm:"Column:marked_npc_3_zone_id"`
	MarkedNpc3InstanceId     uint         `json:"marked_npc_3_instance_id" gorm:"Column:marked_npc_3_instance_id"`
	RaidLeaders              []RaidLeader `json:"raid_leaders,omitempty" gorm:"foreignKey:rid;references:raidid"`
	RaidMembers              []RaidMember `json:"raid_members,omitempty" gorm:"foreignKey:raidid;references:raidid"`
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
