package models

type Friend struct {
	Charid uint   `json:"charid" gorm:"Column:charid"`
	Type   uint8  `json:"type" gorm:"Column:type"`
	Name   string `json:"name" gorm:"Column:name"`
}

func (Friend) TableName() string {
    return "friends"
}

func (Friend) Relationships() []string {
    return []string{}
}

func (Friend) Connection() string {
    return ""
}
