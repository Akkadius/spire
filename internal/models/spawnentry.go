package models

import (
	"github.com/volatiletech/null/v8"
)

type Spawnentry struct {
	SpawngroupID           int         `json:"spawngroup_id" gorm:"Column:spawngroupID"`
	NpcID                  int         `json:"npc_id" gorm:"Column:npcID"`
	Chance                 int16       `json:"chance" gorm:"Column:chance"`
	ConditionValueFilter   int32       `json:"condition_value_filter" gorm:"Column:condition_value_filter"`
	MinExpansion           int8        `json:"min_expansion" gorm:"Column:min_expansion"`
	MaxExpansion           int8        `json:"max_expansion" gorm:"Column:max_expansion"`
	ContentFlags           null.String `json:"content_flags" gorm:"Column:content_flags"`
	ContentFlagsDisabled   null.String `json:"content_flags_disabled" gorm:"Column:content_flags_disabled"`
	Spawngroup             *Spawngroup `json:"spawngroup,omitempty" gorm:"foreignKey:spawngroupID;references:id"`
	NpcType                *NpcType    `json:"npc_type,omitempty" gorm:"foreignKey:npcID;references:id"`
}

func (Spawnentry) TableName() string {
    return "spawnentry"
}

func (Spawnentry) Relationships() []string {
    return []string{
		"NpcType",
		"NpcType.AlternateCurrency",
		"NpcType.Merchantlists",
		"NpcType.Merchantlists.NpcType",
		"NpcType.NpcEmotes",
		"NpcType.NpcFactions",
		"NpcType.NpcFactions.NpcFactionEntries",
		"NpcType.NpcSpells",
		"NpcType.NpcSpells.NpcSpellsEntries",
		"NpcType.NpcTypesTint",
		"NpcType.Spawnentries",
		"Spawngroup",
		"Spawngroup.Spawn2",
		"Spawngroup.Spawn2.Spawnentries",
		"Spawngroup.Spawn2.Spawngroup",
	}
}

func (Spawnentry) Connection() string {
    return "eqemu_content"
}
