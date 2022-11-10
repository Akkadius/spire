package models

import (
	"github.com/volatiletech/null/v8"
)

type DataBucket struct {
	ID      uint64      `json:"id" gorm:"Column:id"`
	Key     null.String `json:"key" gorm:"Column:key"`
	Value   null.String `json:"value" gorm:"Column:value"`
	Expires null.Uint   `json:"expires" gorm:"Column:expires"`
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
