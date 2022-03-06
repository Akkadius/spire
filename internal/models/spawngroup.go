package models

type Spawngroup struct {
	ID            int     `json:"id" gorm:"Column:id"`
	Name          string  `json:"name" gorm:"Column:name"`
	SpawnLimit    int8    `json:"spawn_limit" gorm:"Column:spawn_limit"`
	Dist          float32 `json:"dist" gorm:"Column:dist"`
	MaxX          float32 `json:"max_x" gorm:"Column:max_x"`
	MinX          float32 `json:"min_x" gorm:"Column:min_x"`
	MaxY          float32 `json:"max_y" gorm:"Column:max_y"`
	MinY          float32 `json:"min_y" gorm:"Column:min_y"`
	Delay         int     `json:"delay" gorm:"Column:delay"`
	Mindelay      int     `json:"mindelay" gorm:"Column:mindelay"`
	Despawn       int8    `json:"despawn" gorm:"Column:despawn"`
	DespawnTimer  int     `json:"despawn_timer" gorm:"Column:despawn_timer"`
	WpSpawns      uint8   `json:"wp_spawns" gorm:"Column:wp_spawns"`
	Spawn2        *Spawn2 `json:"spawn_2,omitempty" gorm:"foreignKey:id;references:spawngroupID"`
}

func (Spawngroup) TableName() string {
    return "spawngroup"
}

func (Spawngroup) Relationships() []string {
    return []string{
		"Spawn2",
		"Spawn2.Spawnentries",
		"Spawn2.Spawnentries.NpcType",
		"Spawn2.Spawnentries.NpcType.AlternateCurrency",
		"Spawn2.Spawnentries.NpcType.Loottable",
		"Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries",
		"Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries",
		"Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item",
		"Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.AlternateCurrencies",
		"Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.CharacterCorpseItems",
		"Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.DiscoveredItems",
		"Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Doors",
		"Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Doors.Item",
		"Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Fishings",
		"Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Fishings.Item",
		"Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Fishings.NpcType",
		"Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Fishings.Zone",
		"Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Forages",
		"Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Forages.Item",
		"Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Forages.Zone",
		"Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.GroundSpawns",
		"Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.GroundSpawns.Zone",
		"Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.ItemTicks",
		"Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Keyrings",
		"Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.LootdropEntries",
		"Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Merchantlists",
		"Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType",
		"Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.ObjectContents",
		"Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Objects",
		"Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Objects.Item",
		"Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Objects.Zone",
		"Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.StartingItems",
		"Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.StartingItems.Item",
		"Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.StartingItems.Zone",
		"Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Tasks",
		"Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities",
		"Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.Goallists",
		"Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Tasks.TaskActivities.NpcType",
		"Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Tasks.Tasksets",
		"Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.TradeskillRecipeEntries",
		"Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.TradeskillRecipeEntries.TradeskillRecipe",
		"Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.TributeLevels",
		"Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop",
		"Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LootdropEntries",
		"Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LoottableEntries",
		"Spawn2.Spawnentries.NpcType.Loottable.LoottableEntries.Loottable",
		"Spawn2.Spawnentries.NpcType.Loottable.NpcTypes",
		"Spawn2.Spawnentries.NpcType.Merchantlists",
		"Spawn2.Spawnentries.NpcType.Merchantlists.NpcType",
		"Spawn2.Spawnentries.NpcType.NpcEmotes",
		"Spawn2.Spawnentries.NpcType.NpcFactions",
		"Spawn2.Spawnentries.NpcType.NpcFactions.NpcFactionEntries",
		"Spawn2.Spawnentries.NpcType.NpcSpells",
		"Spawn2.Spawnentries.NpcType.NpcSpells.NpcSpellsEntries",
		"Spawn2.Spawnentries.NpcType.NpcTypesTint",
		"Spawn2.Spawnentries.NpcType.Spawnentries",
		"Spawn2.Spawnentries.Spawngroup",
		"Spawn2.Spawngroup",
	}
}

func (Spawngroup) Connection() string {
    return "eqemu_content"
}
