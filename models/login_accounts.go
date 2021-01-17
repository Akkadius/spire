package models

import (
	"github.com/volatiletech/null/v8"
	"time"
)

type LoginAccount struct {
	ID                 uint        `json:"id" gorm:"Column:id"`
	AccountName        string      `json:"account_name" gorm:"Column:account_name"`
	AccountPassword    string      `json:"account_password" gorm:"Column:account_password"`
	AccountEmail       string      `json:"account_email" gorm:"Column:account_email"`
	SourceLoginserver  null.String `json:"source_loginserver" gorm:"Column:source_loginserver"`
	LastIpAddress      string      `json:"last_ip_address" gorm:"Column:last_ip_address"`
	LastLoginDate      time.Time   `json:"last_login_date" gorm:"Column:last_login_date"`
	CreatedAt          null.Time   `json:"created_at" gorm:"Column:created_at"`
	UpdatedAt          null.Time   `json:"updated_at" gorm:"Column:updated_at"`
}

func (LoginAccount) TableName() string {
    return "login_accounts"
}

func (LoginAccount) Relationships() []string {
    return []string{}
}

func (LoginAccount) Connection() string {
    return ""
}
