package models

type GroupId struct {
	Groupid int    `json:"groupid" gorm:"Column:groupid"`
	Charid  int    `json:"charid" gorm:"Column:charid"`
	Name    string `json:"name" gorm:"Column:name"`
	Ismerc  int8   `json:"ismerc" gorm:"Column:ismerc"`
}

func (GroupId) TableName() string {
    return "group_id"
}

func (GroupId) Relationships() []string {
    return []string{}
}

func (GroupId) Connection() string {
    return ""
}
