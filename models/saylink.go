package models

type Saylink struct {
	ID     int    `json:"id" gorm:"Column:id"`
	Phrase string `json:"phrase" gorm:"Column:phrase"`
}

func (Saylink) TableName() string {
    return "saylink"
}

func (Saylink) Relationships() []string {
    return []string{}
}

func (Saylink) Connection() string {
    return ""
}
