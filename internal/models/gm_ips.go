package models

type GmIp struct {
	Name       string `json:"name" gorm:"Column:name"`
	AccountId  int    `json:"account_id" gorm:"Column:account_id"`
	IpAddress  string `json:"ip_address" gorm:"Column:ip_address"`
}

func (GmIp) TableName() string {
    return "gm_ips"
}

func (GmIp) Relationships() []string {
    return []string{}
}

func (GmIp) Connection() string {
    return ""
}
