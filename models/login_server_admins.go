package models

import (
	"time"
)

type LoginServerAdmin struct {
	ID                      uint      `json:"id" gorm:"Column:id"`
	AccountName             string    `json:"account_name" gorm:"Column:account_name"`
	AccountPassword         string    `json:"account_password" gorm:"Column:account_password"`
	FirstName               string    `json:"first_name" gorm:"Column:first_name"`
	LastName                string    `json:"last_name" gorm:"Column:last_name"`
	Email                   string    `json:"email" gorm:"Column:email"`
	RegistrationDate        time.Time `json:"registration_date" gorm:"Column:registration_date"`
	RegistrationIpAddress   string    `json:"registration_ip_address" gorm:"Column:registration_ip_address"`
}

func (LoginServerAdmin) TableName() string {
    return "login_server_admins"
}

func (LoginServerAdmin) Relationships() []string {
    return []string{}
}

func (LoginServerAdmin) Connection() string {
    return ""
}
