package models

type Pet struct {
	ID           int      `json:"id" gorm:"Column:id"`
	Type         string   `json:"type" gorm:"Column:type"`
	Petpower     int      `json:"petpower" gorm:"Column:petpower"`
	NpcID        int      `json:"npc_id" gorm:"Column:npcID"`
	Temp         int8     `json:"temp" gorm:"Column:temp"`
	Petcontrol   int8     `json:"petcontrol" gorm:"Column:petcontrol"`
	Petnaming    int8     `json:"petnaming" gorm:"Column:petnaming"`
	Monsterflag  int8     `json:"monsterflag" gorm:"Column:monsterflag"`
	Equipmentset int      `json:"equipmentset" gorm:"Column:equipmentset"`
	NpcType      *NpcType `json:"npc_type,omitempty" gorm:"foreignKey:npcID;references:id"`
}

func (Pet) TableName() string {
    return "pets"
}

func (Pet) Relationships() []string {
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
		"NpcType.Spawnentries.NpcType",
		"NpcType.Spawnentries.Spawngroup",
		"NpcType.Spawnentries.Spawngroup.Spawn2",
		"NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries",
		"NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup",
	}
}

func (Pet) Connection() string {
    return "eqemu_content"
}
