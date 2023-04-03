package models

type ChatchannelReservedName struct {
	ID   int    `json:"id" gorm:"Column:id"`
	Name string `json:"name" gorm:"Column:name"`
}

func (ChatchannelReservedName) TableName() string {
    return "chatchannel_reserved_names"
}

func (ChatchannelReservedName) Relationships() []string {
    return []string{}
}

func (ChatchannelReservedName) Connection() string {
    return ""
}
