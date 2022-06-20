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
	SpellsNew  *SpellsNew `json:"spells_new,omitempty" gorm:"foreignKey:spell_id;references:id"`
}

func (Aura) TableName() string {
    return "auras"
}

func (Aura) Relationships() []string {
    return []string{
		"SpellsNew",
		"SpellsNew.Aura",
		"SpellsNew.BlockedSpells",
		"SpellsNew.Damageshieldtypes",
		"SpellsNew.Items",
		"SpellsNew.Items.AlternateCurrencies",
		"SpellsNew.Items.CharacterCorpseItems",
		"SpellsNew.Items.DiscoveredItems",
		"SpellsNew.Items.Doors",
		"SpellsNew.Items.Doors.Item",
		"SpellsNew.Items.Fishings",
		"SpellsNew.Items.Fishings.Item",
		"SpellsNew.Items.Fishings.NpcType",
		"SpellsNew.Items.Fishings.NpcType.AlternateCurrency",
		"SpellsNew.Items.Fishings.NpcType.Loottable",
		"SpellsNew.Items.Fishings.NpcType.Loottable.LoottableEntries",
		"SpellsNew.Items.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop",
		"SpellsNew.Items.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries",
		"SpellsNew.Items.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item",
		"SpellsNew.Items.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Lootdrop",
		"SpellsNew.Items.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LoottableEntries",
		"SpellsNew.Items.Fishings.NpcType.Loottable.LoottableEntries.Loottable",
		"SpellsNew.Items.Fishings.NpcType.Loottable.NpcTypes",
		"SpellsNew.Items.Fishings.NpcType.Merchantlists",
		"SpellsNew.Items.Fishings.NpcType.Merchantlists.Items",
		"SpellsNew.Items.Fishings.NpcType.Merchantlists.NpcTypes",
		"SpellsNew.Items.Fishings.NpcType.NpcEmotes",
		"SpellsNew.Items.Fishings.NpcType.NpcFactions",
		"SpellsNew.Items.Fishings.NpcType.NpcFactions.NpcFactionEntries",
		"SpellsNew.Items.Fishings.NpcType.NpcFactions.NpcFactionEntries.FactionList",
		"SpellsNew.Items.Fishings.NpcType.NpcSpell",
		"SpellsNew.Items.Fishings.NpcType.NpcSpell.NpcSpellsEntries",
		"SpellsNew.Items.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew",
		"SpellsNew.Items.Fishings.NpcType.NpcTypesTint",
		"SpellsNew.Items.Fishings.NpcType.Spawnentries",
		"SpellsNew.Items.Fishings.NpcType.Spawnentries.NpcType",
		"SpellsNew.Items.Fishings.NpcType.Spawnentries.Spawngroup",
		"SpellsNew.Items.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2",
		"SpellsNew.Items.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries",
		"SpellsNew.Items.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup",
		"SpellsNew.Items.Fishings.Zone",
		"SpellsNew.Items.Forages",
		"SpellsNew.Items.Forages.Item",
		"SpellsNew.Items.Forages.Zone",
		"SpellsNew.Items.GroundSpawns",
		"SpellsNew.Items.GroundSpawns.Zone",
		"SpellsNew.Items.ItemTicks",
		"SpellsNew.Items.Keyrings",
		"SpellsNew.Items.LootdropEntries",
		"SpellsNew.Items.LootdropEntries.Item",
		"SpellsNew.Items.LootdropEntries.Lootdrop",
		"SpellsNew.Items.LootdropEntries.Lootdrop.LootdropEntries",
		"SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries",
		"SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Lootdrop",
		"SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable",
		"SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.LoottableEntries",
		"SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes",
		"SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency",
		"SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Loottable",
		"SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists",
		"SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items",
		"SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.NpcTypes",
		"SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcEmotes",
		"SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions",
		"SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions.NpcFactionEntries",
		"SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions.NpcFactionEntries.FactionList",
		"SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell",
		"SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries",
		"SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew",
		"SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcTypesTint",
		"SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries",
		"SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.NpcType",
		"SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup",
		"SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2",
		"SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries",
		"SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup",
		"SpellsNew.Items.Merchantlists",
		"SpellsNew.Items.Merchantlists.Items",
		"SpellsNew.Items.Merchantlists.NpcTypes",
		"SpellsNew.Items.Merchantlists.NpcTypes.AlternateCurrency",
		"SpellsNew.Items.Merchantlists.NpcTypes.Loottable",
		"SpellsNew.Items.Merchantlists.NpcTypes.Loottable.LoottableEntries",
		"SpellsNew.Items.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop",
		"SpellsNew.Items.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries",
		"SpellsNew.Items.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item",
		"SpellsNew.Items.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Lootdrop",
		"SpellsNew.Items.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LoottableEntries",
		"SpellsNew.Items.Merchantlists.NpcTypes.Loottable.LoottableEntries.Loottable",
		"SpellsNew.Items.Merchantlists.NpcTypes.Loottable.NpcTypes",
		"SpellsNew.Items.Merchantlists.NpcTypes.Merchantlists",
		"SpellsNew.Items.Merchantlists.NpcTypes.NpcEmotes",
		"SpellsNew.Items.Merchantlists.NpcTypes.NpcFactions",
		"SpellsNew.Items.Merchantlists.NpcTypes.NpcFactions.NpcFactionEntries",
		"SpellsNew.Items.Merchantlists.NpcTypes.NpcFactions.NpcFactionEntries.FactionList",
		"SpellsNew.Items.Merchantlists.NpcTypes.NpcSpell",
		"SpellsNew.Items.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries",
		"SpellsNew.Items.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew",
		"SpellsNew.Items.Merchantlists.NpcTypes.NpcTypesTint",
		"SpellsNew.Items.Merchantlists.NpcTypes.Spawnentries",
		"SpellsNew.Items.Merchantlists.NpcTypes.Spawnentries.NpcType",
		"SpellsNew.Items.Merchantlists.NpcTypes.Spawnentries.Spawngroup",
		"SpellsNew.Items.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2",
		"SpellsNew.Items.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries",
		"SpellsNew.Items.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup",
		"SpellsNew.Items.ObjectContents",
		"SpellsNew.Items.Objects",
		"SpellsNew.Items.Objects.Item",
		"SpellsNew.Items.Objects.Zone",
		"SpellsNew.Items.StartingItems",
		"SpellsNew.Items.StartingItems.Item",
		"SpellsNew.Items.StartingItems.Zone",
		"SpellsNew.Items.Tasks",
		"SpellsNew.Items.Tasks.TaskActivities",
		"SpellsNew.Items.Tasks.TaskActivities.Goallists",
		"SpellsNew.Items.Tasks.TaskActivities.NpcType",
		"SpellsNew.Items.Tasks.TaskActivities.NpcType.AlternateCurrency",
		"SpellsNew.Items.Tasks.TaskActivities.NpcType.Loottable",
		"SpellsNew.Items.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries",
		"SpellsNew.Items.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.Lootdrop",
		"SpellsNew.Items.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries",
		"SpellsNew.Items.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item",
		"SpellsNew.Items.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Lootdrop",
		"SpellsNew.Items.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.Lootdrop.LoottableEntries",
		"SpellsNew.Items.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.Loottable",
		"SpellsNew.Items.Tasks.TaskActivities.NpcType.Loottable.NpcTypes",
		"SpellsNew.Items.Tasks.TaskActivities.NpcType.Merchantlists",
		"SpellsNew.Items.Tasks.TaskActivities.NpcType.Merchantlists.Items",
		"SpellsNew.Items.Tasks.TaskActivities.NpcType.Merchantlists.NpcTypes",
		"SpellsNew.Items.Tasks.TaskActivities.NpcType.NpcEmotes",
		"SpellsNew.Items.Tasks.TaskActivities.NpcType.NpcFactions",
		"SpellsNew.Items.Tasks.TaskActivities.NpcType.NpcFactions.NpcFactionEntries",
		"SpellsNew.Items.Tasks.TaskActivities.NpcType.NpcFactions.NpcFactionEntries.FactionList",
		"SpellsNew.Items.Tasks.TaskActivities.NpcType.NpcSpell",
		"SpellsNew.Items.Tasks.TaskActivities.NpcType.NpcSpell.NpcSpellsEntries",
		"SpellsNew.Items.Tasks.TaskActivities.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew",
		"SpellsNew.Items.Tasks.TaskActivities.NpcType.NpcTypesTint",
		"SpellsNew.Items.Tasks.TaskActivities.NpcType.Spawnentries",
		"SpellsNew.Items.Tasks.TaskActivities.NpcType.Spawnentries.NpcType",
		"SpellsNew.Items.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup",
		"SpellsNew.Items.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2",
		"SpellsNew.Items.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries",
		"SpellsNew.Items.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup",
		"SpellsNew.Items.Tasks.Tasksets",
		"SpellsNew.Items.TradeskillRecipeEntries",
		"SpellsNew.Items.TradeskillRecipeEntries.TradeskillRecipe",
		"SpellsNew.Items.TributeLevels",
		"SpellsNew.NpcSpellsEntries",
		"SpellsNew.NpcSpellsEntries.SpellsNew",
		"SpellsNew.SpellBuckets",
		"SpellsNew.SpellGlobals",
	}
}

func (Aura) Connection() string {
    return "eqemu_content"
}
