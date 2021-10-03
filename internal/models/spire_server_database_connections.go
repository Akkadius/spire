package models

import (
	"time"
)

type ServerDatabaseConnection struct {
	ID                uint       `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Name              string     `json:"name" gorm:"type:varchar(255);"`
	DbHost            string     `json:"db_host" gorm:"type:varchar(50);"`
	DbPort            string     `json:"db_port" gorm:"type:varchar(50);"`
	DbName            string     `json:"db_name" gorm:"type:varchar(50);"`
	DbUsername        string     `json:"db_username" gorm:"type:varchar(50);"`
	DbPassword        string     `json:"db_password" gorm:"type:varchar(250);"`
	ContentDbHost     string     `json:"content_db_host" gorm:"type:varchar(50);"`
	ContentDbPort     string     `json:"content_db_port" gorm:"type:varchar(50);"`
	ContentDbName     string     `json:"content_db_name" gorm:"type:varchar(50);"`
	ContentDbUsername string     `json:"content_db_username" gorm:"type:varchar(50);"`
	ContentDbPassword string     `json:"content_db_password" gorm:"type:varchar(250);"`
	CreatedFromIp     string     `json:"created_from_ip" gorm:"type:varchar(50);"`
	CreatedBy         uint       `json:"created_by" gorm:"Column:created_by;default:0"`
	CreatedAt         time.Time  `json:"created_at" gorm:"Column:created_at"`
	UpdatedAt         time.Time  `json:"updated_at" gorm:"Column:updated_at"`
	DeletedAt         *time.Time `json:"deleted_at" gorm:"Column:deleted_at;null"`
}

func (ServerDatabaseConnection) TableName() string {
	return "server_database_connections"
}

func (ServerDatabaseConnection) Relationships() []string {
	return []string{}
}

func (ServerDatabaseConnection) Connection() string {
	return "spire"
}
