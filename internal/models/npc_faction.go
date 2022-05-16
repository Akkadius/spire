package models

import (
	"github.com/volatiletech/null/v8"
)

type NpcFaction struct {
	ID                    int               `json:"id" gorm:"Column:id"`
	Name                  null.String       `json:"name" gorm:"Column:name"`
	Primaryfaction        int               `json:"primaryfaction" gorm:"Column:primaryfaction"`
	IgnorePrimaryAssist   int8              `json:"ignore_primary_assist" gorm:"Column:ignore_primary_assist"`
	NpcFactionEntries     []NpcFactionEntry `json:"npc_faction_entries,omitempty" gorm:"foreignKey:npc_faction_id;references:id"`
}

func (NpcFaction) TableName() string {
    return "npc_faction"
}

func (NpcFaction) Relationships() []string {
    return []string{
		"NpcFactionEntries",
		"NpcFactionEntries.FactionList",
	}
}

func (NpcFaction) Connection() string {
    return "eqemu_content"
}
