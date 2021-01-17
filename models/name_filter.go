package models

type NameFilter struct {
	ID   int    `json:"id" gorm:"Column:id"`
	Name string `json:"name" gorm:"Column:name"`
}

func (NameFilter) TableName() string {
    return "name_filter"
}

func (NameFilter) Relationships() []string {
    return []string{}
}

func (NameFilter) Connection() string {
    return ""
}
