package models

import (
	"github.com/volatiletech/null/v8"
)

type BannedIp struct {
	IpAddress  string      `json:"ip_address" gorm:"Column:ip_address"`
	Notes      null.String `json:"notes" gorm:"Column:notes"`
}

func (BannedIp) TableName() string {
    return "banned_ips"
}

func (BannedIp) Relationships() []string {
    return []string{}
}

func (BannedIp) Connection() string {
    return ""
}
