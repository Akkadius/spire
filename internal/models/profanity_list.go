package models

type ProfanityList struct {
	Word string `json:"word" gorm:"Column:word"`
}

func (ProfanityList) TableName() string {
    return "profanity_list"
}

func (ProfanityList) Relationships() []string {
    return []string{}
}

func (ProfanityList) Connection() string {
    return ""
}
