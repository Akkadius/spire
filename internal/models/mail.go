package models

type Mail struct {
	Msgid     uint   `json:"msgid" gorm:"Column:msgid"`
	Charid    uint   `json:"charid" gorm:"Column:charid"`
	Timestamp int    `json:"timestamp" gorm:"Column:timestamp"`
	From      string `json:"from" gorm:"Column:from"`
	Subject   string `json:"subject" gorm:"Column:subject"`
	Body      string `json:"body" gorm:"Column:body"`
	To        string `json:"to" gorm:"Column:to"`
	Status    int8   `json:"status" gorm:"Column:status"`
}

func (Mail) TableName() string {
    return "mail"
}

func (Mail) Relationships() []string {
    return []string{}
}

func (Mail) Connection() string {
    return ""
}
