package models

type CharacterLeadershipAbility struct {
	ID   uint   `json:"id" gorm:"Column:id"`
	Slot uint16 `json:"slot" gorm:"Column:slot"`
	Rank uint16 `json:"rank" gorm:"Column:rank"`
}

func (CharacterLeadershipAbility) TableName() string {
    return "character_leadership_abilities"
}

func (CharacterLeadershipAbility) Relationships() []string {
    return []string{}
}

func (CharacterLeadershipAbility) Connection() string {
    return ""
}
