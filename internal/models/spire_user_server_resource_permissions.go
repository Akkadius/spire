package models

import (
	"time"
)

type UserServerResourcePermission struct {
	ID                         uint64                   `json:"id" gorm:"primary_key,AUTO_INCREMENT"`
	UserId                     uint                     `json:"user_id"`
	ServerDatabaseConnectionId uint                     `json:"server_database_connection_id"`
	ResourceName               string                   `json:"resource_name" db:"resource_name"`
	CanWrite                   uint8                    `json:"can_write" db:"can_write"`
	CanRead                    uint8                    `json:"can_read" db:"can_read"`
	CreatedAt                  time.Time                `json:"created_at" gorm:"Column:created_at"`
	ServerDatabaseConnection   ServerDatabaseConnection `json:"database_connection,omitempty" gorm:"foreignKey:ServerDatabaseConnectionId;association_foreignkey:Id"`
}

func (UserServerResourcePermission) TableName() string {
	return "spire_user_server_resource_permissions"
}

func (UserServerResourcePermission) Relationships() []string {
	return []string{}
}

func (UserServerResourcePermission) Connection() string {
	return "spire"
}
