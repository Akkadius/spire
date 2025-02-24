package models

import (
	"github.com/volatiletech/null/v8"
)

type PlayerEventKilledRaidNpc struct {
	ID                            uint64      `json:"id" gorm:"Column:id"`
	NpcId                         null.Uint   `json:"npc_id" gorm:"Column:npc_id"`
	NpcName                       null.String `json:"npc_name" gorm:"Column:npc_name"`
	CombatTimeSeconds             null.Uint   `json:"combat_time_seconds" gorm:"Column:combat_time_seconds"`
	TotalDamagePerSecondTaken     null.Uint64 `json:"total_damage_per_second_taken" gorm:"Column:total_damage_per_second_taken"`
	TotalHealPerSecondTaken       null.Uint64 `json:"total_heal_per_second_taken" gorm:"Column:total_heal_per_second_taken"`
	CreatedAt                     null.Time   `json:"created_at" gorm:"Column:created_at"`
}

func (PlayerEventKilledRaidNpc) TableName() string {
    return "player_event_killed_raid_npc"
}

func (PlayerEventKilledRaidNpc) Relationships() []string {
    return []string{}
}

func (PlayerEventKilledRaidNpc) Connection() string {
    return ""
}
