package models

import (
	"time"
)

type UserServerDatabaseConnection struct {
	ID                         uint                     `json:"id" gorm:"primary_key,AUTO_INCREMENT"`
	UserId                     uint                     `json:"user_id"`
	ServerDatabaseConnectionId uint                     `json:"server_database_connection_id"`
	Active                     uint                     `json:"active" gorm:"default:0"`
	CreatedBy                  uint                     `json:"created_by" gorm:"Column:created_by;default:0"`
	CreatedAt                  time.Time                `json:"created_at" gorm:"Column:created_at"`
	UpdatedAt                  time.Time                `json:"updated_at" gorm:"Column:updated_at"`
	DeletedAt                  *time.Time               `json:"deleted_at" gorm:"Column:deleted_at"`
	ServerDatabaseConnection   ServerDatabaseConnection `json:"database_connection,omitempty" gorm:"foreignKey:ServerDatabaseConnectionId;association_foreignkey:Id"`
}

func (UserServerDatabaseConnection) TableName() string {
	return "user_server_database_connections"
}

func (UserServerDatabaseConnection) Relationships() []string {
	return []string{
		"ServerDatabaseConnection",
	}
}

func (UserServerDatabaseConnection) Connection() string {
	return "spire"
}
