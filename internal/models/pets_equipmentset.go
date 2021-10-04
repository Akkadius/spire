package models

type PetsEquipmentset struct {
	SetId                   int                     `json:"set_id" gorm:"Column:set_id"`
	Setname                 string                  `json:"setname" gorm:"Column:setname"`
	NestedSet               int                     `json:"nested_set" gorm:"Column:nested_set"`
	PetsEquipmentsetEntries []PetsEquipmentsetEntry `json:"pets_equipmentset_entries,omitempty" gorm:"foreignKey:set_id;references:set_id"`
}

func (PetsEquipmentset) TableName() string {
    return "pets_equipmentset"
}

func (PetsEquipmentset) Relationships() []string {
    return []string{
		"PetsEquipmentsetEntries",
	}
}

func (PetsEquipmentset) Connection() string {
    return "eqemu_content"
}
