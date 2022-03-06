package models

type LootdropEntry struct {
	LootdropId        uint      `json:"lootdrop_id" gorm:"Column:lootdrop_id"`
	ItemId            int       `json:"item_id" gorm:"Column:item_id"`
	ItemCharges       uint16    `json:"item_charges" gorm:"Column:item_charges"`
	EquipItem         uint8     `json:"equip_item" gorm:"Column:equip_item"`
	Chance            float32   `json:"chance" gorm:"Column:chance"`
	DisabledChance    float32   `json:"disabled_chance" gorm:"Column:disabled_chance"`
	TrivialMinLevel   uint16    `json:"trivial_min_level" gorm:"Column:trivial_min_level"`
	TrivialMaxLevel   uint16    `json:"trivial_max_level" gorm:"Column:trivial_max_level"`
	Multiplier        uint8     `json:"multiplier" gorm:"Column:multiplier"`
	NpcMinLevel       uint16    `json:"npc_min_level" gorm:"Column:npc_min_level"`
	NpcMaxLevel       uint16    `json:"npc_max_level" gorm:"Column:npc_max_level"`
	Item              *Item     `json:"item,omitempty" gorm:"foreignKey:item_id;references:id"`
	Lootdrop          *Lootdrop `json:"lootdrop,omitempty" gorm:"foreignKey:lootdrop_id;references:id"`
}

func (LootdropEntry) TableName() string {
    return "lootdrop_entries"
}

func (LootdropEntry) Relationships() []string {
    return []string{
		"Item",
		"Item.AlternateCurrencies",
		"Item.CharacterCorpseItems",
		"Item.DiscoveredItems",
		"Item.Doors",
		"Item.Doors.Item",
		"Item.Fishings",
		"Item.Fishings.Item",
		"Item.Fishings.NpcType",
		"Item.Fishings.NpcType.AlternateCurrency",
		"Item.Fishings.NpcType.Loottable",
		"Item.Fishings.NpcType.Loottable.LoottableEntries",
		"Item.Fishings.NpcType.Loottable.LoottableEntries.LootdropEntries",
		"Item.Fishings.NpcType.Loottable.LoottableEntries.Loottable",
		"Item.Fishings.NpcType.Loottable.NpcTypes",
		"Item.Fishings.NpcType.Merchantlists",
		"Item.Fishings.NpcType.Merchantlists.NpcType",
		"Item.Fishings.NpcType.NpcEmotes",
		"Item.Fishings.NpcType.NpcFactions",
		"Item.Fishings.NpcType.NpcFactions.NpcFactionEntries",
		"Item.Fishings.NpcType.NpcSpells",
		"Item.Fishings.NpcType.NpcSpells.NpcSpellsEntries",
		"Item.Fishings.NpcType.NpcTypesTint",
		"Item.Fishings.NpcType.Spawnentries",
		"Item.Fishings.NpcType.Spawnentries.NpcType",
		"Item.Fishings.NpcType.Spawnentries.Spawngroup",
		"Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2",
		"Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries",
		"Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup",
		"Item.Fishings.Zone",
		"Item.Forages",
		"Item.Forages.Item",
		"Item.Forages.Zone",
		"Item.GroundSpawns",
		"Item.GroundSpawns.Zone",
		"Item.ItemTicks",
		"Item.Keyrings",
		"Item.LootdropEntries",
		"Item.Merchantlists",
		"Item.Merchantlists.NpcType",
		"Item.Merchantlists.NpcType.AlternateCurrency",
		"Item.Merchantlists.NpcType.Loottable",
		"Item.Merchantlists.NpcType.Loottable.LoottableEntries",
		"Item.Merchantlists.NpcType.Loottable.LoottableEntries.LootdropEntries",
		"Item.Merchantlists.NpcType.Loottable.LoottableEntries.Loottable",
		"Item.Merchantlists.NpcType.Loottable.NpcTypes",
		"Item.Merchantlists.NpcType.Merchantlists",
		"Item.Merchantlists.NpcType.NpcEmotes",
		"Item.Merchantlists.NpcType.NpcFactions",
		"Item.Merchantlists.NpcType.NpcFactions.NpcFactionEntries",
		"Item.Merchantlists.NpcType.NpcSpells",
		"Item.Merchantlists.NpcType.NpcSpells.NpcSpellsEntries",
		"Item.Merchantlists.NpcType.NpcTypesTint",
		"Item.Merchantlists.NpcType.Spawnentries",
		"Item.Merchantlists.NpcType.Spawnentries.NpcType",
		"Item.Merchantlists.NpcType.Spawnentries.Spawngroup",
		"Item.Merchantlists.NpcType.Spawnentries.Spawngroup.Spawn2",
		"Item.Merchantlists.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries",
		"Item.Merchantlists.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup",
		"Item.ObjectContents",
		"Item.Objects",
		"Item.Objects.Item",
		"Item.Objects.Zone",
		"Item.StartingItems",
		"Item.StartingItems.Item",
		"Item.StartingItems.Zone",
		"Item.Tasks",
		"Item.Tasks.TaskActivities",
		"Item.Tasks.TaskActivities.Goallists",
		"Item.Tasks.TaskActivities.NpcType",
		"Item.Tasks.TaskActivities.NpcType.AlternateCurrency",
		"Item.Tasks.TaskActivities.NpcType.Loottable",
		"Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries",
		"Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries",
		"Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.Loottable",
		"Item.Tasks.TaskActivities.NpcType.Loottable.NpcTypes",
		"Item.Tasks.TaskActivities.NpcType.Merchantlists",
		"Item.Tasks.TaskActivities.NpcType.Merchantlists.NpcType",
		"Item.Tasks.TaskActivities.NpcType.NpcEmotes",
		"Item.Tasks.TaskActivities.NpcType.NpcFactions",
		"Item.Tasks.TaskActivities.NpcType.NpcFactions.NpcFactionEntries",
		"Item.Tasks.TaskActivities.NpcType.NpcSpells",
		"Item.Tasks.TaskActivities.NpcType.NpcSpells.NpcSpellsEntries",
		"Item.Tasks.TaskActivities.NpcType.NpcTypesTint",
		"Item.Tasks.TaskActivities.NpcType.Spawnentries",
		"Item.Tasks.TaskActivities.NpcType.Spawnentries.NpcType",
		"Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup",
		"Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2",
		"Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries",
		"Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup",
		"Item.Tasks.Tasksets",
		"Item.TradeskillRecipeEntries",
		"Item.TradeskillRecipeEntries.TradeskillRecipe",
		"Item.TributeLevels",
		"Lootdrop",
		"Lootdrop.LootdropEntries",
		"Lootdrop.LoottableEntries",
		"Lootdrop.LoottableEntries.LootdropEntries",
		"Lootdrop.LoottableEntries.Loottable",
		"Lootdrop.LoottableEntries.Loottable.LoottableEntries",
		"Lootdrop.LoottableEntries.Loottable.NpcTypes",
		"Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency",
		"Lootdrop.LoottableEntries.Loottable.NpcTypes.Loottable",
		"Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists",
		"Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.NpcType",
		"Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcEmotes",
		"Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions",
		"Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions.NpcFactionEntries",
		"Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpells",
		"Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpells.NpcSpellsEntries",
		"Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcTypesTint",
		"Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries",
		"Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.NpcType",
		"Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup",
		"Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2",
		"Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries",
		"Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup",
	}
}

func (LootdropEntry) Connection() string {
    return "eqemu_content"
}
