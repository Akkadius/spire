package models

import (
	"time"
)

type UserEventLog struct {
	ID                         uint                      `json:"id" gorm:"primary_key,AUTO_INCREMENT"`
	UserId                     uint                      `json:"user_id" db:"user_id"`
	ServerDatabaseConnectionId uint                      `json:"server_database_connection_id" db:"server_database_connection_id" gorm:"index:server_database_connection_id_event;index:server_database_connection_id"`
	EventName                  string                    `json:"event_name" db:"event_name" gorm:"index:server_database_connection_id_event"`
	Data                       string                    `json:"data" db:"data"`
	CreatedAt                  time.Time                 `json:"created_at" gorm:"Column:created_at"`
	ServerDatabaseConnection   *ServerDatabaseConnection `json:"database_connection,omitempty" gorm:"foreignKey:ServerDatabaseConnectionId;association_foreignkey:Id"`
	User                       *User                     `json:"user,omitempty" gorm:"foreignKey:ID;references:UserId"`
}

func (UserEventLog) TableName() string {
	return "spire_user_event_log"
}

func (UserEventLog) Relationships() []string {
	return []string{}
}

func (UserEventLog) Connection() string {
	return "spire"
}
