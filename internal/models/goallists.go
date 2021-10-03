package models

type Goallist struct {
	Listid uint `json:"listid" gorm:"Column:listid"`
	Entry  uint `json:"entry" gorm:"Column:entry"`
}

func (Goallist) TableName() string {
    return "goallists"
}

func (Goallist) Relationships() []string {
    return []string{}
}

func (Goallist) Connection() string {
    return "eqemu_content"
}
