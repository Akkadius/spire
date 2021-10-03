package models

type CharacterPotionbelt struct {
	ID        uint  `json:"id" gorm:"Column:id"`
	PotionId  uint8 `json:"potion_id" gorm:"Column:potion_id"`
	ItemId    uint  `json:"item_id" gorm:"Column:item_id"`
	Icon      uint  `json:"icon" gorm:"Column:icon"`
}

func (CharacterPotionbelt) TableName() string {
    return "character_potionbelt"
}

func (CharacterPotionbelt) Relationships() []string {
    return []string{}
}

func (CharacterPotionbelt) Connection() string {
    return ""
}
