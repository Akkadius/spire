package models

type LoottableEntry struct {
	LoottableId     uint            `json:"loottable_id" gorm:"Column:loottable_id"`
	LootdropId      uint            `json:"lootdrop_id" gorm:"Column:lootdrop_id"`
	Multiplier      uint8           `json:"multiplier" gorm:"Column:multiplier"`
	Droplimit       uint8           `json:"droplimit" gorm:"Column:droplimit"`
	Mindrop         uint8           `json:"mindrop" gorm:"Column:mindrop"`
	Probability     float32         `json:"probability" gorm:"Column:probability"`
	LootdropEntries []LootdropEntry `json:"lootdrop_entries,omitempty" gorm:"foreignKey:lootdrop_id;references:lootdrop_id"`
}

func (LoottableEntry) TableName() string {
    return "loottable_entries"
}

func (LoottableEntry) Relationships() []string {
    return []string{
		"LootdropEntries",
		"LootdropEntries.Item",
		"LootdropEntries.Item.DiscoveredItems",
	}
}

func (LoottableEntry) Connection() string {
    return "eqemu_content"
}
