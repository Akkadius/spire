package models

type PetsEquipmentsetEntry struct {
	SetId   int `json:"set_id" gorm:"Column:set_id"`
	Slot    int `json:"slot" gorm:"Column:slot"`
	ItemId  int `json:"item_id" gorm:"Column:item_id"`
}

func (PetsEquipmentsetEntry) TableName() string {
    return "pets_equipmentset_entries"
}

func (PetsEquipmentsetEntry) Relationships() []string {
    return []string{}
}

func (PetsEquipmentsetEntry) Connection() string {
    return "eqemu_content"
}
