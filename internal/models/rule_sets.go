package models

type RuleSet struct {
	RulesetId  uint8       `json:"ruleset_id" gorm:"Column:ruleset_id"`
	Name       string      `json:"name" gorm:"Column:name"`
	RuleValues []RuleValue `json:"rule_values,omitempty" gorm:"foreignKey:ruleset_id;references:ruleset_id"`
}

func (RuleSet) TableName() string {
    return "rule_sets"
}

func (RuleSet) Relationships() []string {
    return []string{
		"RuleValues",
	}
}

func (RuleSet) Connection() string {
    return ""
}
