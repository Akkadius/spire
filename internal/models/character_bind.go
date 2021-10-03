package models

type CharacterBind struct {
	ID          uint    `json:"id" gorm:"Column:id"`
	Slot        int     `json:"slot" gorm:"Column:slot"`
	ZoneId      uint16  `json:"zone_id" gorm:"Column:zone_id"`
	InstanceId  uint32  `json:"instance_id" gorm:"Column:instance_id"`
	X           float32 `json:"x" gorm:"Column:x"`
	Y           float32 `json:"y" gorm:"Column:y"`
	Z           float32 `json:"z" gorm:"Column:z"`
	Heading     float32 `json:"heading" gorm:"Column:heading"`
}

func (CharacterBind) TableName() string {
    return "character_bind"
}

func (CharacterBind) Relationships() []string {
    return []string{}
}

func (CharacterBind) Connection() string {
    return ""
}
