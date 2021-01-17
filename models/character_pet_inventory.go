package models

type CharacterPetInventory struct {
	CharId  int `json:"char_id" gorm:"Column:char_id"`
	Pet     int `json:"pet" gorm:"Column:pet"`
	Slot    int `json:"slot" gorm:"Column:slot"`
	ItemId  int `json:"item_id" gorm:"Column:item_id"`
}

func (CharacterPetInventory) TableName() string {
    return "character_pet_inventory"
}

func (CharacterPetInventory) Relationships() []string {
    return []string{}
}

func (CharacterPetInventory) Connection() string {
    return ""
}
