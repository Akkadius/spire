package models

import (
	"github.com/volatiletech/null/v8"
)

type DbVersion struct {
	Version      null.Int `json:"version" gorm:"Column:version"`
	BotsVersion  null.Int `json:"bots_version" gorm:"Column:bots_version"`
}

func (DbVersion) TableName() string {
    return "db_version"
}

func (DbVersion) Relationships() []string {
    return []string{}
}

func (DbVersion) Connection() string {
    return ""
}
