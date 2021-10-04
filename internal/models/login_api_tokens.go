package models

import (
	"github.com/volatiletech/null/v8"
)

type LoginApiToken struct {
	ID         int         `json:"id" gorm:"Column:id"`
	Token      null.String `json:"token" gorm:"Column:token"`
	CanWrite   null.Int    `json:"can_write" gorm:"Column:can_write"`
	CanRead    null.Int    `json:"can_read" gorm:"Column:can_read"`
	CreatedAt  null.Time   `json:"created_at" gorm:"Column:created_at"`
	UpdatedAt  null.Time   `json:"updated_at" gorm:"Column:updated_at"`
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
