package models

type CharacterItemRecast struct {
	ID          uint `json:"id" gorm:"Column:id"`
	RecastType  uint `json:"recast_type" gorm:"Column:recast_type"`
	Timestamp   uint `json:"timestamp" gorm:"Column:timestamp"`
}

func (CharacterItemRecast) TableName() string {
    return "character_item_recast"
}

func (CharacterItemRecast) Relationships() []string {
    return []string{}
}

func (CharacterItemRecast) Connection() string {
    return ""
}
