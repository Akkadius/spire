package models

type CharacterTribute struct {
	ID      uint  `json:"id" gorm:"Column:id"`
	Tier    uint8 `json:"tier" gorm:"Column:tier"`
	Tribute uint  `json:"tribute" gorm:"Column:tribute"`
}

func (CharacterTribute) TableName() string {
    return "character_tribute"
}

func (CharacterTribute) Relationships() []string {
    return []string{}
}

func (CharacterTribute) Connection() string {
    return ""
}
