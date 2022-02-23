package models

type Aura struct {
	Type       int        `json:"type" gorm:"Column:type"`
	NpcType    int        `json:"npc_type" gorm:"Column:npc_type"`
	Name       string     `json:"name" gorm:"Column:name"`
	SpellId    int        `json:"spell_id" gorm:"Column:spell_id"`
	Distance   int        `json:"distance" gorm:"Column:distance"`
	AuraType   int        `json:"aura_type" gorm:"Column:aura_type"`
	SpawnType  int        `json:"spawn_type" gorm:"Column:spawn_type"`
	Movement   int        `json:"movement" gorm:"Column:movement"`
	Duration   int        `json:"duration" gorm:"Column:duration"`
	Icon       int        `json:"icon" gorm:"Column:icon"`
	CastTime   int        `json:"cast_time" gorm:"Column:cast_time"`
	NpcTypeRelation*NpcType   `json:"npc_type,omitempty" gorm:"foreignKey:npc_type;references:id"`
	SpellsNew  *SpellsNew `json:"spells_new,omitempty" gorm:"foreignKey:spell_id;references:id"`
}

func (Aura) TableName() string {
    return "auras"
}

func (Aura) Relationships() []string {
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
		"SpellsNew",
		"SpellsNew.Aura",
		"SpellsNew.BlockedSpells",
		"SpellsNew.Damageshieldtypes",
		"SpellsNew.SpellBuckets",
		"SpellsNew.SpellGlobals",
	}
}

func (Aura) Connection() string {
    return "eqemu_content"
}
