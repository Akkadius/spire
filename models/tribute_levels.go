package models

type TributeLevel struct {
	TributeId  uint `json:"tribute_id" gorm:"Column:tribute_id"`
	Level      uint `json:"level" gorm:"Column:level"`
	Cost       uint `json:"cost" gorm:"Column:cost"`
	ItemId     uint `json:"item_id" gorm:"Column:item_id"`
}

func (TributeLevel) TableName() string {
    return "tribute_levels"
}

func (TributeLevel) Relationships() []string {
    return []string{}
}

func (TributeLevel) Connection() string {
    return "eqemu_content"
}
