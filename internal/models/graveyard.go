package models

type Graveyard struct {
	ID      int     `json:"id" gorm:"Column:id"`
	ZoneId  int     `json:"zone_id" gorm:"Column:zone_id"`
	X       float32 `json:"x" gorm:"Column:x"`
	Y       float32 `json:"y" gorm:"Column:y"`
	Z       float32 `json:"z" gorm:"Column:z"`
	Heading float32 `json:"heading" gorm:"Column:heading"`
}

func (Graveyard) TableName() string {
    return "graveyard"
}

func (Graveyard) Relationships() []string {
    return []string{}
}

func (Graveyard) Connection() string {
    return "eqemu_content"
}
