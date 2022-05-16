package models

type NpcFactionEntry struct {
	NpcFactionId   uint         `json:"npc_faction_id" gorm:"Column:npc_faction_id"`
	FactionId      uint         `json:"faction_id" gorm:"Column:faction_id"`
	Value          int          `json:"value" gorm:"Column:value"`
	NpcValue       int8         `json:"npc_value" gorm:"Column:npc_value"`
	Temp           int8         `json:"temp" gorm:"Column:temp"`
	FactionList    *FactionList `json:"faction_list,omitempty" gorm:"foreignKey:faction_id;references:id"`
}

func (NpcFactionEntry) TableName() string {
    return "npc_faction_entries"
}

func (NpcFactionEntry) Relationships() []string {
    return []string{
		"FactionList",
	}
}

func (NpcFactionEntry) Connection() string {
    return "eqemu_content"
}
