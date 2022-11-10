package models

import (
	"github.com/volatiletech/null/v8"
)

type IpExemption struct {
	ExemptionId      int         `json:"exemption_id" gorm:"Column:exemption_id"`
	ExemptionIp      null.String `json:"exemption_ip" gorm:"Column:exemption_ip"`
	ExemptionAmount  null.Int    `json:"exemption_amount" gorm:"Column:exemption_amount"`
}

func (IpExemption) TableName() string {
    return "ip_exemptions"
}

func (IpExemption) Relationships() []string {
    return []string{}
}

func (IpExemption) Connection() string {
    return ""
}
