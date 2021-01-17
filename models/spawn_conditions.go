package models

type SpawnCondition struct {
	Zone     string `json:"zone" gorm:"Column:zone"`
	ID       uint32 `json:"id" gorm:"Column:id"`
	Value    int32  `json:"value" gorm:"Column:value"`
	Onchange uint8  `json:"onchange" gorm:"Column:onchange"`
	Name     string `json:"name" gorm:"Column:name"`
}

func (SpawnCondition) TableName() string {
    return "spawn_conditions"
}

func (SpawnCondition) Relationships() []string {
    return []string{}
}

func (SpawnCondition) Connection() string {
    return "eqemu_content"
}
