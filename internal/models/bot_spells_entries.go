package models

type BotSpellsEntry struct {
	ID                uint       `json:"id" gorm:"Column:id"`
	NpcSpellsId       int        `json:"npc_spells_id" gorm:"Column:npc_spells_id"`
	SpellId           uint16     `json:"spell_id" gorm:"Column:spell_id"`
	Type              uint       `json:"type" gorm:"Column:type"`
	Minlevel          uint8      `json:"minlevel" gorm:"Column:minlevel"`
	Maxlevel          uint8      `json:"maxlevel" gorm:"Column:maxlevel"`
	Manacost          int16      `json:"manacost" gorm:"Column:manacost"`
	RecastDelay       int        `json:"recast_delay" gorm:"Column:recast_delay"`
	Priority          int16      `json:"priority" gorm:"Column:priority"`
	ResistAdjust      int        `json:"resist_adjust" gorm:"Column:resist_adjust"`
	MinHp             int16      `json:"min_hp" gorm:"Column:min_hp"`
	MaxHp             int16      `json:"max_hp" gorm:"Column:max_hp"`
	BucketName        string     `json:"bucket_name" gorm:"Column:bucket_name"`
	BucketValue       string     `json:"bucket_value" gorm:"Column:bucket_value"`
	BucketComparison  uint8      `json:"bucket_comparison" gorm:"Column:bucket_comparison"`
	NpcSpell          *NpcSpell  `json:"npc_spell,omitempty" gorm:"foreignKey:npc_spells_id;references:id"`
	SpellsNew         *SpellsNew `json:"spells_new,omitempty" gorm:"foreignKey:spell_id;references:id"`
}

func (BotSpellsEntry) TableName() string {
    return "bot_spells_entries"
}

func (BotSpellsEntry) Relationships() []string {
    return []string{
		"NpcSpell",
		"NpcSpell.BotSpellsEntries",
		"NpcSpell.NpcSpell",
		"NpcSpell.NpcSpellsEntries",
		"NpcSpell.NpcSpellsEntries.SpellsNew",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Aura",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew",
		"NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells",
		"NpcSpell.NpcSpellsEntries.SpellsNew.BotSpellsEntries",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.AlternateCurrencies",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.AlternateCurrencies.Item",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.CharacterCorpseItems",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.DiscoveredItems",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.Doors",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.Doors.Item",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.Item",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.AlternateCurrency",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.AlternateCurrency.Item",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.Loottable",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.Loottable.LoottableEntries",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Lootdrop",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LoottableEntries",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.Loottable.LoottableEntries.Loottable",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.Loottable.NpcTypes",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.Merchantlists",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.Merchantlists.Items",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.Merchantlists.NpcTypes",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.NpcEmotes",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.NpcFactions",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.NpcFactions.NpcFactionEntries",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.NpcFactions.NpcFactionEntries.FactionList",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.NpcSpell",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.NpcTypesTint",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.Spawnentries",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.Spawnentries.NpcType",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.Spawnentries.Spawngroup",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.Zone",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.Forages",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.Forages.Item",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.Forages.Zone",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.GroundSpawns",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.GroundSpawns.Zone",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.ItemTicks",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.Keyrings",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Item",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LootdropEntries",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Lootdrop",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.LoottableEntries",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Loottable",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.NpcTypes",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcEmotes",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions.NpcFactionEntries",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions.NpcFactionEntries.FactionList",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcTypesTint",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.NpcType",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists.Items",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.AlternateCurrency",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.AlternateCurrency.Item",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.Loottable",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.Loottable.LoottableEntries",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Lootdrop",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LoottableEntries",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.Loottable.LoottableEntries.Loottable",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.Loottable.NpcTypes",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.Merchantlists",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.NpcEmotes",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.NpcFactions",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.NpcFactions.NpcFactionEntries",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.NpcFactions.NpcFactionEntries.FactionList",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.NpcSpell",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.NpcTypesTint",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.Spawnentries",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.Spawnentries.NpcType",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.Spawnentries.Spawngroup",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.ObjectContents",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.Objects",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.Objects.Item",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.Objects.Zone",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.TradeskillRecipeEntries",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.TradeskillRecipeEntries.TradeskillRecipe",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.TradeskillRecipeEntries.TradeskillRecipe.TradeskillRecipeEntries",
		"NpcSpell.NpcSpellsEntries.SpellsNew.Items.TributeLevels",
		"NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries",
		"NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets",
		"NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals",
		"SpellsNew",
		"SpellsNew.Aura",
		"SpellsNew.Aura.SpellsNew",
		"SpellsNew.BlockedSpells",
		"SpellsNew.BotSpellsEntries",
		"SpellsNew.Damageshieldtypes",
		"SpellsNew.Items",
		"SpellsNew.Items.AlternateCurrencies",
		"SpellsNew.Items.AlternateCurrencies.Item",
		"SpellsNew.Items.CharacterCorpseItems",
		"SpellsNew.Items.DiscoveredItems",
		"SpellsNew.Items.Doors",
		"SpellsNew.Items.Doors.Item",
		"SpellsNew.Items.Fishings",
		"SpellsNew.Items.Fishings.Item",
		"SpellsNew.Items.Fishings.NpcType",
		"SpellsNew.Items.Fishings.NpcType.AlternateCurrency",
		"SpellsNew.Items.Fishings.NpcType.AlternateCurrency.Item",
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
		"SpellsNew.Items.Fishings.NpcType.NpcSpell.BotSpellsEntries",
		"SpellsNew.Items.Fishings.NpcType.NpcSpell.NpcSpell",
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
		"SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item",
		"SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Loottable",
		"SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists",
		"SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items",
		"SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.NpcTypes",
		"SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcEmotes",
		"SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions",
		"SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions.NpcFactionEntries",
		"SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions.NpcFactionEntries.FactionList",
		"SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell",
		"SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.BotSpellsEntries",
		"SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpell",
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
		"SpellsNew.Items.Merchantlists.NpcTypes.AlternateCurrency.Item",
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
		"SpellsNew.Items.Merchantlists.NpcTypes.NpcSpell.BotSpellsEntries",
		"SpellsNew.Items.Merchantlists.NpcTypes.NpcSpell.NpcSpell",
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
		"SpellsNew.Items.TradeskillRecipeEntries",
		"SpellsNew.Items.TradeskillRecipeEntries.TradeskillRecipe",
		"SpellsNew.Items.TradeskillRecipeEntries.TradeskillRecipe.TradeskillRecipeEntries",
		"SpellsNew.Items.TributeLevels",
		"SpellsNew.NpcSpellsEntries",
		"SpellsNew.NpcSpellsEntries.SpellsNew",
		"SpellsNew.SpellBuckets",
		"SpellsNew.SpellGlobals",
	}
}

func (BotSpellsEntry) Connection() string {
    return ""
}
