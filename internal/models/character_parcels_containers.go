package models

type CharacterParcelsContainer struct {
	ID         uint `json:"id" gorm:"Column:id"`
	ParcelsId  uint `json:"parcels_id" gorm:"Column:parcels_id"`
	SlotId     uint `json:"slot_id" gorm:"Column:slot_id"`
	ItemId     uint `json:"item_id" gorm:"Column:item_id"`
	AugSlot1   uint `json:"aug_slot_1" gorm:"Column:aug_slot_1"`
	AugSlot2   uint `json:"aug_slot_2" gorm:"Column:aug_slot_2"`
	AugSlot3   uint `json:"aug_slot_3" gorm:"Column:aug_slot_3"`
	AugSlot4   uint `json:"aug_slot_4" gorm:"Column:aug_slot_4"`
	AugSlot5   uint `json:"aug_slot_5" gorm:"Column:aug_slot_5"`
	AugSlot6   uint `json:"aug_slot_6" gorm:"Column:aug_slot_6"`
	Quantity   uint `json:"quantity" gorm:"Column:quantity"`
}

func (CharacterParcelsContainer) TableName() string {
    return "character_parcels_containers"
}

func (CharacterParcelsContainer) Relationships() []string {
    return []string{}
}

func (CharacterParcelsContainer) Connection() string {
    return ""
}
