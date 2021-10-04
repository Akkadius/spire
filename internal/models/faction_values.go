package models

type FactionValue struct {
	CharId        int   `json:"char_id" gorm:"Column:char_id"`
	FactionId     int   `json:"faction_id" gorm:"Column:faction_id"`
	CurrentValue  int16 `json:"current_value" gorm:"Column:current_value"`
	Temp          int8  `json:"temp" gorm:"Column:temp"`
}

func (FactionValue) TableName() string {
    return "faction_values"
}

func (FactionValue) Relationships() []string {
    return []string{}
}

func (FactionValue) Connection() string {
    return ""
}
