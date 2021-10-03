package models

type LootdropEntry struct {
	LootdropId        uint    `json:"lootdrop_id" gorm:"Column:lootdrop_id"`
	ItemId            int     `json:"item_id" gorm:"Column:item_id"`
	ItemCharges       uint16  `json:"item_charges" gorm:"Column:item_charges"`
	EquipItem         uint8   `json:"equip_item" gorm:"Column:equip_item"`
	Chance            float32 `json:"chance" gorm:"Column:chance"`
	DisabledChance    float32 `json:"disabled_chance" gorm:"Column:disabled_chance"`
	TrivialMinLevel   uint16  `json:"trivial_min_level" gorm:"Column:trivial_min_level"`
	TrivialMaxLevel   uint16  `json:"trivial_max_level" gorm:"Column:trivial_max_level"`
	Multiplier        uint8   `json:"multiplier" gorm:"Column:multiplier"`
	NpcMinLevel       uint16  `json:"npc_min_level" gorm:"Column:npc_min_level"`
	NpcMaxLevel       uint16  `json:"npc_max_level" gorm:"Column:npc_max_level"`
	Item              *Item   `json:"item,omitempty" gorm:"foreignKey:item_id;references:id"`
}

func (LootdropEntry) TableName() string {
    return "lootdrop_entries"
}

func (LootdropEntry) Relationships() []string {
    return []string{
		"Item",
		"Item.DiscoveredItems",
	}
}

func (LootdropEntry) Connection() string {
    return "eqemu_content"
}
