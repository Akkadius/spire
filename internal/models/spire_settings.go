package models

import (
	"time"
)

type Setting struct {
	ID        uint      `json:"id" gorm:"primary_key,AUTO_INCREMENT,type:bigint(11)"`
	Setting   string    `json:"setting" gorm:"type:varchar(255);uniqueIndex"`
	Value     string    `json:"value" gorm:"type:varchar(255);"`
	CreatedAt time.Time `json:"created_at" gorm:"Column:created_at"`
}

func (Setting) TableName() string {
	return "spire_settings"
}

func (Setting) Relationships() []string {
	return []string{}
}

func (Setting) Connection() string {
	return "spire"
}
