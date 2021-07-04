package models

type Spawnentry struct {
	SpawngroupID           int   `json:"spawngroup_id" gorm:"Column:spawngroupID"`
	NpcID                  int   `json:"npc_id" gorm:"Column:npcID"`
	Chance                 int16 `json:"chance" gorm:"Column:chance"`
	ConditionValueFilter   int32 `json:"condition_value_filter" gorm:"Column:condition_value_filter"`
}

func (Spawnentry) TableName() string {
    return "spawnentry"
}

func (Spawnentry) Relationships() []string {
    return []string{}
}

func (Spawnentry) Connection() string {
    return "eqemu_content"
}
