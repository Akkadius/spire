package models

import (
	"github.com/volatiletech/null/v8"
)

type ZoneStateSpawn struct {
	ID                  int64       `json:"id" gorm:"Column:id"`
	ZoneId              null.Uint   `json:"zone_id" gorm:"Column:zone_id"`
	InstanceId          null.Uint   `json:"instance_id" gorm:"Column:instance_id"`
	IsCorpse            null.Int8   `json:"is_corpse" gorm:"Column:is_corpse"`
	IsZone              null.Int8   `json:"is_zone" gorm:"Column:is_zone"`
	DecayInSeconds      null.Int    `json:"decay_in_seconds" gorm:"Column:decay_in_seconds"`
	NpcId               null.Uint   `json:"npc_id" gorm:"Column:npc_id"`
	Spawn2Id            uint        `json:"spawn_2_id" gorm:"Column:spawn2_id"`
	SpawngroupId        uint        `json:"spawngroup_id" gorm:"Column:spawngroup_id"`
	X                   float32     `json:"x" gorm:"Column:x"`
	Y                   float32     `json:"y" gorm:"Column:y"`
	Z                   float32     `json:"z" gorm:"Column:z"`
	Heading             float32     `json:"heading" gorm:"Column:heading"`
	RespawnTime         uint        `json:"respawn_time" gorm:"Column:respawn_time"`
	Variance            uint        `json:"variance" gorm:"Column:variance"`
	Grid                null.Uint   `json:"grid" gorm:"Column:grid"`
	CurrentWaypoint     null.Int    `json:"current_waypoint" gorm:"Column:current_waypoint"`
	PathWhenZoneIdle    null.Int16  `json:"path_when_zone_idle" gorm:"Column:path_when_zone_idle"`
	ConditionId         null.Uint16 `json:"condition_id" gorm:"Column:condition_id"`
	ConditionMinValue   null.Int16  `json:"condition_min_value" gorm:"Column:condition_min_value"`
	Enabled             null.Int16  `json:"enabled" gorm:"Column:enabled"`
	Anim                null.Uint16 `json:"anim" gorm:"Column:anim"`
	LootData            null.String `json:"loot_data" gorm:"Column:loot_data"`
	EntityVariables     null.String `json:"entity_variables" gorm:"Column:entity_variables"`
	Buffs               null.String `json:"buffs" gorm:"Column:buffs"`
	Hp                  null.Int64  `json:"hp" gorm:"Column:hp"`
	Mana                null.Int64  `json:"mana" gorm:"Column:mana"`
	Endurance           null.Int64  `json:"endurance" gorm:"Column:endurance"`
	CreatedAt           null.Time   `json:"created_at" gorm:"Column:created_at"`
}

func (ZoneStateSpawn) TableName() string {
    return "zone_state_spawns"
}

func (ZoneStateSpawn) Relationships() []string {
    return []string{}
}

func (ZoneStateSpawn) Connection() string {
    return ""
}
