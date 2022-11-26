package models

type BotCreateCombination struct {
	Race    uint `json:"race" gorm:"Column:race"`
	Classes uint `json:"classes" gorm:"Column:classes"`
}

func (BotCreateCombination) TableName() string {
    return "bot_create_combinations"
}

func (BotCreateCombination) Relationships() []string {
    return []string{}
}

func (BotCreateCombination) Connection() string {
    return ""
}
