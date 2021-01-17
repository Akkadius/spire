package models

type Timer struct {
	CharId   int    `json:"char_id" gorm:"Column:char_id"`
	Type     uint32 `json:"type" gorm:"Column:type"`
	Start    uint   `json:"start" gorm:"Column:start"`
	Duration uint   `json:"duration" gorm:"Column:duration"`
	Enable   int8   `json:"enable" gorm:"Column:enable"`
}

func (Timer) TableName() string {
    return "timers"
}

func (Timer) Relationships() []string {
    return []string{}
}

func (Timer) Connection() string {
    return ""
}
