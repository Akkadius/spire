package models

type Pet struct {
	ID           int      `json:"id" gorm:"Column:id"`
	Type         string   `json:"type" gorm:"Column:type"`
	Petpower     int      `json:"petpower" gorm:"Column:petpower"`
	NpcID        int      `json:"npc_id" gorm:"Column:npcID"`
	Temp         int8     `json:"temp" gorm:"Column:temp"`
	Petcontrol   int8     `json:"petcontrol" gorm:"Column:petcontrol"`
	Petnaming    int8     `json:"petnaming" gorm:"Column:petnaming"`
	Monsterflag  int8     `json:"monsterflag" gorm:"Column:monsterflag"`
	Equipmentset int      `json:"equipmentset" gorm:"Column:equipmentset"`
	NpcType      *NpcType `json:"npc_type,omitempty" gorm:"foreignKey:npcID;references:id"`
}

func (Pet) TableName() string {
    return "pets"
}

func (Pet) Relationships() []string {
    return []string{
		"NpcType",
		"NpcType.AlternateCurrency",
		"NpcType.Merchantlists",
		"NpcType.NpcEmotes",
		"NpcType.NpcFactions",
		"NpcType.NpcFactions.NpcFactionEntries",
		"NpcType.NpcSpells",
		"NpcType.NpcSpells.NpcSpellsEntries",
		"NpcType.NpcTypesTint",
		"NpcType.Spawnentries",
		"NpcType.Spawnentries.NpcType",
		"NpcType.Spawnentries.Spawngroup",
		"NpcType.Spawnentries.Spawngroup.Spawn2",
		"NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries",
		"NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup",
	}
}

func (Pet) Connection() string {
    return "eqemu_content"
}
