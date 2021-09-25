package models

import (
	"github.com/volatiletech/null/v8"
)

type NpcSpell struct {
	ID                       uint             `json:"id" gorm:"Column:id"`
	Name                     null.String      `json:"name" gorm:"Column:name"`
	ParentList               uint             `json:"parent_list" gorm:"Column:parent_list"`
	AttackProc               int16            `json:"attack_proc" gorm:"Column:attack_proc"`
	ProcChance               int8             `json:"proc_chance" gorm:"Column:proc_chance"`
	RangeProc                int16            `json:"range_proc" gorm:"Column:range_proc"`
	RprocChance              int16            `json:"rproc_chance" gorm:"Column:rproc_chance"`
	DefensiveProc            int16            `json:"defensive_proc" gorm:"Column:defensive_proc"`
	DprocChance              int16            `json:"dproc_chance" gorm:"Column:dproc_chance"`
	FailRecast               uint             `json:"fail_recast" gorm:"Column:fail_recast"`
	EngagedNoSpRecastMin     uint             `json:"engaged_no_sp_recast_min" gorm:"Column:engaged_no_sp_recast_min"`
	EngagedNoSpRecastMax     uint             `json:"engaged_no_sp_recast_max" gorm:"Column:engaged_no_sp_recast_max"`
	EngagedBSelfChance       uint8            `json:"engaged_b_self_chance" gorm:"Column:engaged_b_self_chance"`
	EngagedBOtherChance      uint8            `json:"engaged_b_other_chance" gorm:"Column:engaged_b_other_chance"`
	EngagedDChance           uint8            `json:"engaged_d_chance" gorm:"Column:engaged_d_chance"`
	PursueNoSpRecastMin      uint             `json:"pursue_no_sp_recast_min" gorm:"Column:pursue_no_sp_recast_min"`
	PursueNoSpRecastMax      uint             `json:"pursue_no_sp_recast_max" gorm:"Column:pursue_no_sp_recast_max"`
	PursueDChance            uint8            `json:"pursue_d_chance" gorm:"Column:pursue_d_chance"`
	IdleNoSpRecastMin        uint             `json:"idle_no_sp_recast_min" gorm:"Column:idle_no_sp_recast_min"`
	IdleNoSpRecastMax        uint             `json:"idle_no_sp_recast_max" gorm:"Column:idle_no_sp_recast_max"`
	IdleBChance              uint8            `json:"idle_b_chance" gorm:"Column:idle_b_chance"`
	NpcSpellsEntries         []NpcSpellsEntry `json:"npc_spells_entries,omitempty" gorm:"foreignKey:npc_spells_id;references:id"`
}

func (NpcSpell) TableName() string {
    return "npc_spells"
}

func (NpcSpell) Relationships() []string {
    return []string{
		"NpcSpellsEntries",
	}
}

func (NpcSpell) Connection() string {
    return "eqemu_content"
}
