package models

type Spawnentry struct {
	SpawngroupID           int         `json:"spawngroup_id" gorm:"Column:spawngroupID"`
	NpcID                  int         `json:"npc_id" gorm:"Column:npcID"`
	Chance                 int16       `json:"chance" gorm:"Column:chance"`
	ConditionValueFilter   int32       `json:"condition_value_filter" gorm:"Column:condition_value_filter"`
	Spawngroup             *Spawngroup `json:"spawngroup,omitempty" gorm:"foreignKey:spawngroupID;references:id"`
	NpcType                *NpcType    `json:"npc_type,omitempty" gorm:"foreignKey:npcID;references:id"`
}

func (Spawnentry) TableName() string {
    return "spawnentry"
}

func (Spawnentry) Relationships() []string {
    return []string{
		"NpcType",
		"NpcType.AlternateCurrency",
		"NpcType.Merchantlists",
		"NpcType.Merchantlists.Item",
		"NpcType.Merchantlists.Item.AlternateCurrencies",
		"NpcType.Merchantlists.Item.CharacterCorpseItems",
		"NpcType.Merchantlists.Item.DiscoveredItems",
		"NpcType.Merchantlists.Item.Doors",
		"NpcType.Merchantlists.Item.Doors.Item",
		"NpcType.Merchantlists.Item.Doors.Zone",
		"NpcType.Merchantlists.Item.Fishings",
		"NpcType.Merchantlists.Item.Fishings.Item",
		"NpcType.Merchantlists.Item.Fishings.NpcType",
		"NpcType.Merchantlists.Item.Fishings.Zone",
		"NpcType.Merchantlists.Item.Forages",
		"NpcType.Merchantlists.Item.Forages.Item",
		"NpcType.Merchantlists.Item.Forages.Zone",
		"NpcType.Merchantlists.Item.GroundSpawns",
		"NpcType.Merchantlists.Item.GroundSpawns.Item",
		"NpcType.Merchantlists.Item.GroundSpawns.Zone",
		"NpcType.Merchantlists.Item.ItemTicks",
		"NpcType.Merchantlists.Item.Keyrings",
		"NpcType.Merchantlists.Item.LootdropEntries",
		"NpcType.Merchantlists.Item.LootdropEntries.Item",
		"NpcType.Merchantlists.Item.LootdropEntries.Lootdrop",
		"NpcType.Merchantlists.Item.LootdropEntries.Lootdrop.LootdropEntries",
		"NpcType.Merchantlists.Item.LootdropEntries.Lootdrop.LoottableEntries",
		"NpcType.Merchantlists.Item.LootdropEntries.Lootdrop.LoottableEntries.LootdropEntries",
		"NpcType.Merchantlists.Item.Merchantlists",
		"NpcType.Merchantlists.Item.ObjectContents",
		"NpcType.Merchantlists.Item.Objects",
		"NpcType.Merchantlists.Item.Objects.Item",
		"NpcType.Merchantlists.Item.Objects.Zone",
		"NpcType.Merchantlists.Item.StartingItems",
		"NpcType.Merchantlists.Item.StartingItems.Item",
		"NpcType.Merchantlists.Item.StartingItems.Zone",
		"NpcType.Merchantlists.Item.TradeskillRecipeEntries",
		"NpcType.Merchantlists.Item.TradeskillRecipeEntries.TradeskillRecipe",
		"NpcType.Merchantlists.Item.TributeLevels",
		"NpcType.NpcEmotes",
		"NpcType.NpcFactions",
		"NpcType.NpcFactions.NpcFactionEntries",
		"NpcType.NpcSpells",
		"NpcType.NpcSpells.NpcSpellsEntries",
		"NpcType.NpcTypesTint",
		"NpcType.Spawnentries",
		"Spawngroup",
		"Spawngroup.Spawn2",
		"Spawngroup.Spawn2.Spawnentries",
		"Spawngroup.Spawn2.Spawngroup",
	}
}

func (Spawnentry) Connection() string {
    return "eqemu_content"
}
