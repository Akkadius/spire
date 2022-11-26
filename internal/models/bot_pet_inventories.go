package models

type BotPetInventory struct {
	PetInventoriesIndex   uint `json:"pet_inventories_index" gorm:"Column:pet_inventories_index"`
	PetsIndex             uint `json:"pets_index" gorm:"Column:pets_index"`
	ItemId                uint `json:"item_id" gorm:"Column:item_id"`
}

func (BotPetInventory) TableName() string {
    return "bot_pet_inventories"
}

func (BotPetInventory) Relationships() []string {
    return []string{}
}

func (BotPetInventory) Connection() string {
    return ""
}
