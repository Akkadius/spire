package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID                           uint                           `db:"id" json:"id"`
	UserName                     string                         `db:"user_name" json:"user_name"`
	FullName                     string                         `db:"full_name" json:"full_name"`
	FirstName                    string                         `db:"first_name" json:"first_name"`
	LastName                     string                         `db:"last_name" json:"last_name"`
	Email                        string                         `db:"email" json:"email"`
	Avatar                       string                         `db:"avatar" json:"avatar"`
	Provider                     string                         `db:"provider" json:"provider"`
	Password                     string                         `db:"password" json:"-"`
	IsAdmin                      bool                           `db:"is_admin" json:"is_admin" gorm:"default:0"`
	CreatedAt                    time.Time                      `db:"created_at" json:"created_at"`
	UpdatedAt                    time.Time                      `db:"updated_at" json:"updated_at"`
	Deleted                      gorm.DeletedAt                 `db:"deleted_at" json:"deleted_at"`
	ServerDatabaseConnections    []ServerDatabaseConnection     `json:"owned_connections,omitempty" gorm:"foreignKey:CreatedBy;association_foreignkey:Id"`
	UserServerDatabaseConnection []UserServerDatabaseConnection `json:"user_connections,omitempty" gorm:"foreignKey:UserId;association_foreignkey:Id"`
}

func (User) TableName() string {
	return "spire_users"
}

func (User) Relationships() []string {
	return []string{
		"ServerDatabaseConnections",
		"UserServerDatabaseConnection",
		"UserServerDatabaseConnection.ServerDatabaseConnection",
	}
}

func (User) Connection() string {
	return "spire"
}
