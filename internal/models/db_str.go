package models

type DbStr struct {
	ID    int64  `json:"id" gorm:"Column:id"`
	Type  int    `json:"type" gorm:"Column:type"`
	Value string `json:"value" gorm:"Column:value"`
}

func (DbStr) TableName() string {
    return "db_str"
}

func (DbStr) Relationships() []string {
    return []string{}
}

func (DbStr) Connection() string {
    return ""
}
