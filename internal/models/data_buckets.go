package models

import (
	"github.com/volatiletech/null/v8"
)

type DataBucket struct {
	ID      uint64      `json:"id" gorm:"Column:id"`
	Key     string `json:"key" gorm:"Column:key"`
	Value   string `json:"value" gorm:"Column:value"`
	Expires uint   `json:"expires" gorm:"Column:expires"`
}

func (DataBucket) TableName() string {
    return "data_buckets"
}

func (DataBucket) Relationships() []string {
    return []string{}
}

func (DataBucket) Connection() string {
    return ""
}
