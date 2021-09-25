package models

import (
	"github.com/volatiletech/null/v8"
)

type Lootdrop struct {
	ID                     uint            `json:"id" gorm:"Column:id"`
	Name                   string          `json:"name" gorm:"Column:name"`
	MinExpansion           uint8           `json:"min_expansion" gorm:"Column:min_expansion"`
	MaxExpansion           uint8           `json:"max_expansion" gorm:"Column:max_expansion"`
	ContentFlags           null.String     `json:"content_flags" gorm:"Column:content_flags"`
	ContentFlagsDisabled   null.String     `json:"content_flags_disabled" gorm:"Column:content_flags_disabled"`
	LootdropEntries        []LootdropEntry `json:"lootdrop_entries,omitempty" gorm:"foreignKey:lootdrop_id;references:id"`
}

func (Lootdrop) TableName() string {
    return "lootdrop"
}

func (Lootdrop) Relationships() []string {
    return []string{
		"LootdropEntries",
		"LootdropEntries.Item",
		"LootdropEntries.Item.DiscoveredItems",
	}
}

func (Lootdrop) Connection() string {
    return "eqemu_content"
}
