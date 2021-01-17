package models

type CharacterLanguage struct {
	ID      uint   `json:"id" gorm:"Column:id"`
	LangId  uint16 `json:"lang_id" gorm:"Column:lang_id"`
	Value   uint16 `json:"value" gorm:"Column:value"`
}

func (CharacterLanguage) TableName() string {
    return "character_languages"
}

func (CharacterLanguage) Relationships() []string {
    return []string{}
}

func (CharacterLanguage) Connection() string {
    return ""
}
