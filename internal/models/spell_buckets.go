package models

import (
	"github.com/volatiletech/null/v8"
)

type SpellBucket struct {
	Spellid uint64      `json:"spellid" gorm:"Column:spellid"`
	Key     null.String `json:"key" gorm:"Column:key"`
	Value   null.String `json:"value" gorm:"Column:value"`
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
