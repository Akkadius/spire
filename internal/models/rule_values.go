package models

import (
	"github.com/volatiletech/null/v8"
)

type RuleValue struct {
	RulesetId  uint8       `json:"ruleset_id" gorm:"Column:ruleset_id"`
	RuleName   string      `json:"rule_name" gorm:"Column:rule_name"`
	RuleValue  null.String `json:"rule_value" gorm:"Column:rule_value"`
	Notes      null.String `json:"notes" gorm:"Column:notes"`
}

func (RuleValue) TableName() string {
    return "rule_values"
}

func (RuleValue) Relationships() []string {
    return []string{}
}

func (RuleValue) Connection() string {
    return ""
}
