package models

import (
	"github.com/volatiletech/null/v8"
)

type SpellBucket struct {
	Spellid uint64      `json:"spellid" gorm:"Column:spellid"`
	Key     string `json:"key" gorm:"Column:key"`
	Value   string `json:"value" gorm:"Column:value"`
}

func (SpellBucket) TableName() string {
    return "spell_buckets"
}

func (SpellBucket) Relationships() []string {
    return []string{}
}

func (SpellBucket) Connection() string {
    return ""
}
