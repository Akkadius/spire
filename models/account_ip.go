package models

import (
	"time"
)

type AccountIp struct {
	Accid    int       `json:"accid" gorm:"Column:accid"`
	Ip       string    `json:"ip" gorm:"Column:ip"`
	Count    int       `json:"count" gorm:"Column:count"`
	Lastused time.Time `json:"lastused" gorm:"Column:lastused"`
}

func (AccountIp) TableName() string {
    return "account_ip"
}

func (AccountIp) Relationships() []string {
    return []string{}
}

func (AccountIp) Connection() string {
    return ""
}
