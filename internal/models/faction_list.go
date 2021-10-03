package models

type FactionList struct {
	ID   int    `json:"id" gorm:"Column:id"`
	Name string `json:"name" gorm:"Column:name"`
	Base int16  `json:"base" gorm:"Column:base"`
}

func (FactionList) TableName() string {
    return "faction_list"
}

func (FactionList) Relationships() []string {
    return []string{}
}

func (FactionList) Connection() string {
    return "eqemu_content"
}
