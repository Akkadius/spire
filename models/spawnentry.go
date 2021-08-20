package models

type Spawnentry struct {
	SpawngroupID           int         `json:"spawngroup_id" gorm:"Column:spawngroupID"`
	NpcID                  int         `json:"npc_id" gorm:"Column:npcID"`
	Chance                 int16       `json:"chance" gorm:"Column:chance"`
	ConditionValueFilter   int32       `json:"condition_value_filter" gorm:"Column:condition_value_filter"`
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
		"NpcType.NpcFactions",
		"NpcType.NpcFactions.NpcFactionEntries",
		"NpcType.NpcSpells",
		"NpcType.NpcSpells.NpcSpellsEntries",
		"Spawngroup",
		"Spawngroup.Spawn2",
	}
}

func (Spawnentry) Connection() string {
    return "eqemu_content"
}
