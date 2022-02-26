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
		"LootdropEntries.Item.AlternateCurrencies",
		"LootdropEntries.Item.CharacterCorpseItems",
		"LootdropEntries.Item.DiscoveredItems",
		"LootdropEntries.Item.Doors",
		"LootdropEntries.Item.Doors.Item",
		"LootdropEntries.Item.Fishings",
		"LootdropEntries.Item.Fishings.Item",
		"LootdropEntries.Item.Fishings.NpcType",
		"LootdropEntries.Item.Fishings.NpcType.AlternateCurrency",
		"LootdropEntries.Item.Fishings.NpcType.Merchantlists",
		"LootdropEntries.Item.Fishings.NpcType.NpcEmotes",
		"LootdropEntries.Item.Fishings.NpcType.NpcFactions",
		"LootdropEntries.Item.Fishings.NpcType.NpcFactions.NpcFactionEntries",
		"LootdropEntries.Item.Fishings.NpcType.NpcSpells",
		"LootdropEntries.Item.Fishings.NpcType.NpcSpells.NpcSpellsEntries",
		"LootdropEntries.Item.Fishings.NpcType.NpcTypesTint",
		"LootdropEntries.Item.Fishings.NpcType.Spawnentries",
		"LootdropEntries.Item.Fishings.NpcType.Spawnentries.NpcType",
		"LootdropEntries.Item.Fishings.NpcType.Spawnentries.Spawngroup",
		"LootdropEntries.Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2",
		"LootdropEntries.Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries",
		"LootdropEntries.Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup",
		"LootdropEntries.Item.Fishings.Zone",
		"LootdropEntries.Item.Forages",
		"LootdropEntries.Item.Forages.Item",
		"LootdropEntries.Item.Forages.Zone",
		"LootdropEntries.Item.ItemTicks",
		"LootdropEntries.Item.Keyrings",
		"LootdropEntries.Item.LootdropEntries",
		"LootdropEntries.Item.ObjectContents",
		"LootdropEntries.Item.Objects",
		"LootdropEntries.Item.Objects.Item",
		"LootdropEntries.Item.Objects.Zone",
		"LootdropEntries.Item.StartingItems",
		"LootdropEntries.Item.StartingItems.Item",
		"LootdropEntries.Item.StartingItems.Zone",
		"LootdropEntries.Item.TradeskillRecipeEntries",
		"LootdropEntries.Item.TradeskillRecipeEntries.TradeskillRecipe",
		"LootdropEntries.Item.TributeLevels",
		"LootdropEntries.Lootdrop",
		"LootdropEntries.Lootdrop.LootdropEntries",
		"LootdropEntries.Lootdrop.LoottableEntries",
	}
}

func (LoottableEntry) Connection() string {
    return "eqemu_content"
}
