package models

type CharacterMaterial struct {
	ID       uint  `json:"id" gorm:"Column:id"`
	Slot     uint8 `json:"slot" gorm:"Column:slot"`
	Blue     uint8 `json:"blue" gorm:"Column:blue"`
	Green    uint8 `json:"green" gorm:"Column:green"`
	Red      uint8 `json:"red" gorm:"Column:red"`
	UseTint  uint8 `json:"use_tint" gorm:"Column:use_tint"`
	Color    uint  `json:"color" gorm:"Column:color"`
}

func (CharacterMaterial) TableName() string {
    return "character_material"
}

func (CharacterMaterial) Relationships() []string {
    return []string{}
}

func (CharacterMaterial) Connection() string {
    return ""
}
