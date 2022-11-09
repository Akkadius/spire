package models

import (
	"github.com/volatiletech/null/v8"
)

type LoginApiToken struct {
	ID         int         `json:"id" gorm:"Column:id"`
	Token      string `json:"token" gorm:"Column:token"`
	CanWrite   int    `json:"can_write" gorm:"Column:can_write"`
	CanRead    int    `json:"can_read" gorm:"Column:can_read"`
	CreatedAt  time.Time   `json:"created_at" gorm:"Column:created_at"`
	UpdatedAt  time.Time   `json:"updated_at" gorm:"Column:updated_at"`
}

func (LoginApiToken) TableName() string {
    return "login_api_tokens"
}

func (LoginApiToken) Relationships() []string {
    return []string{}
}

func (LoginApiToken) Connection() string {
    return ""
}
