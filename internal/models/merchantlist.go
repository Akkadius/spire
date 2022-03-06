package models

import (
	"github.com/volatiletech/null/v8"
)

type Merchantlist struct {
	Merchantid             int         `json:"merchantid" gorm:"Column:merchantid"`
	Slot                   int         `json:"slot" gorm:"Column:slot"`
	Item                   int         `json:"item" gorm:"Column:item"`
	FactionRequired        int16       `json:"faction_required" gorm:"Column:faction_required"`
	LevelRequired          uint8       `json:"level_required" gorm:"Column:level_required"`
	AltCurrencyCost        uint16      `json:"alt_currency_cost" gorm:"Column:alt_currency_cost"`
	ClassesRequired        int         `json:"classes_required" gorm:"Column:classes_required"`
	Probability            int         `json:"probability" gorm:"Column:probability"`
	MinExpansion           int8        `json:"min_expansion" gorm:"Column:min_expansion"`
	MaxExpansion           int8        `json:"max_expansion" gorm:"Column:max_expansion"`
	ContentFlags           null.String `json:"content_flags" gorm:"Column:content_flags"`
	ContentFlagsDisabled   null.String `json:"content_flags_disabled" gorm:"Column:content_flags_disabled"`
	NpcType                *NpcType    `json:"npc_type,omitempty" gorm:"foreignKey:merchantid;references:merchant_id"`
}

func (Merchantlist) TableName() string {
    return "merchantlist"
}

func (Merchantlist) Relationships() []string {
    return []string{
		"NpcType",
		"NpcType.AlternateCurrency",
		"NpcType.Merchantlists",
		"NpcType.NpcEmotes",
		"NpcType.NpcFactions",
		"NpcType.NpcFactions.NpcFactionEntries",
		"NpcType.NpcSpells",
		"NpcType.NpcSpells.NpcSpellsEntries",
		"NpcType.NpcTypesTint",
		"NpcType.Spawnentries",
		"NpcType.Spawnentries.NpcType",
		"NpcType.Spawnentries.Spawngroup",
		"NpcType.Spawnentries.Spawngroup.Spawn2",
		"NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries",
		"NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup",
	}
}

func (Merchantlist) Connection() string {
    return "eqemu_content"
}
