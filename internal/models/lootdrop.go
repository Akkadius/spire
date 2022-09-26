package models

import (
	"github.com/volatiletech/null/v8"
)

type Lootdrop struct {
	ID                     uint             `json:"id" gorm:"Column:id"`
	Name                   string           `json:"name" gorm:"Column:name"`
	MinExpansion           int8             `json:"min_expansion" gorm:"Column:min_expansion"`
	MaxExpansion           int8             `json:"max_expansion" gorm:"Column:max_expansion"`
	ContentFlags           null.String      `json:"content_flags" gorm:"Column:content_flags"`
	ContentFlagsDisabled   null.String      `json:"content_flags_disabled" gorm:"Column:content_flags_disabled"`
	LootdropEntries        []LootdropEntry  `json:"lootdrop_entries,omitempty" gorm:"foreignKey:lootdrop_id;references:id"`
	LoottableEntries       []LoottableEntry `json:"loottable_entries,omitempty" gorm:"foreignKey:lootdrop_id;references:id"`
}

func (Lootdrop) TableName() string {
    return "lootdrop"
}

func (Lootdrop) Relationships() []string {
    return []string{
		"LootdropEntries",
		"LootdropEntries.Item",
		"LootdropEntries.Item.AlternateCurrencies",
		"LootdropEntries.Item.AlternateCurrencies.Item",
		"LootdropEntries.Item.CharacterCorpseItems",
		"LootdropEntries.Item.DiscoveredItems",
		"LootdropEntries.Item.Doors",
		"LootdropEntries.Item.Doors.Item",
		"LootdropEntries.Item.Fishings",
		"LootdropEntries.Item.Fishings.Item",
		"LootdropEntries.Item.Fishings.NpcType",
		"LootdropEntries.Item.Fishings.NpcType.AlternateCurrency",
		"LootdropEntries.Item.Fishings.NpcType.AlternateCurrency.Item",
		"LootdropEntries.Item.Fishings.NpcType.Loottable",
		"LootdropEntries.Item.Fishings.NpcType.Loottable.LoottableEntries",
		"LootdropEntries.Item.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop",
		"LootdropEntries.Item.Fishings.NpcType.Loottable.LoottableEntries.Loottable",
		"LootdropEntries.Item.Fishings.NpcType.Loottable.NpcTypes",
		"LootdropEntries.Item.Fishings.NpcType.Merchantlists",
		"LootdropEntries.Item.Fishings.NpcType.Merchantlists.Items",
		"LootdropEntries.Item.Fishings.NpcType.Merchantlists.NpcTypes",
		"LootdropEntries.Item.Fishings.NpcType.NpcEmotes",
		"LootdropEntries.Item.Fishings.NpcType.NpcFactions",
		"LootdropEntries.Item.Fishings.NpcType.NpcFactions.NpcFactionEntries",
		"LootdropEntries.Item.Fishings.NpcType.NpcFactions.NpcFactionEntries.FactionList",
		"LootdropEntries.Item.Fishings.NpcType.NpcSpell",
		"LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpell",
		"LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries",
		"LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew",
		"LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Aura",
		"LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew",
		"LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells",
		"LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes",
		"LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items",
		"LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries",
		"LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets",
		"LootdropEntries.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals",
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
		"LootdropEntries.Item.GroundSpawns",
		"LootdropEntries.Item.GroundSpawns.Zone",
		"LootdropEntries.Item.ItemTicks",
		"LootdropEntries.Item.Keyrings",
		"LootdropEntries.Item.LootdropEntries",
		"LootdropEntries.Item.Merchantlists",
		"LootdropEntries.Item.Merchantlists.Items",
		"LootdropEntries.Item.Merchantlists.NpcTypes",
		"LootdropEntries.Item.Merchantlists.NpcTypes.AlternateCurrency",
		"LootdropEntries.Item.Merchantlists.NpcTypes.AlternateCurrency.Item",
		"LootdropEntries.Item.Merchantlists.NpcTypes.Loottable",
		"LootdropEntries.Item.Merchantlists.NpcTypes.Loottable.LoottableEntries",
		"LootdropEntries.Item.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop",
		"LootdropEntries.Item.Merchantlists.NpcTypes.Loottable.LoottableEntries.Loottable",
		"LootdropEntries.Item.Merchantlists.NpcTypes.Loottable.NpcTypes",
		"LootdropEntries.Item.Merchantlists.NpcTypes.Merchantlists",
		"LootdropEntries.Item.Merchantlists.NpcTypes.NpcEmotes",
		"LootdropEntries.Item.Merchantlists.NpcTypes.NpcFactions",
		"LootdropEntries.Item.Merchantlists.NpcTypes.NpcFactions.NpcFactionEntries",
		"LootdropEntries.Item.Merchantlists.NpcTypes.NpcFactions.NpcFactionEntries.FactionList",
		"LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell",
		"LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpell",
		"LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries",
		"LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew",
		"LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura",
		"LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew",
		"LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells",
		"LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes",
		"LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items",
		"LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries",
		"LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets",
		"LootdropEntries.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals",
		"LootdropEntries.Item.Merchantlists.NpcTypes.NpcTypesTint",
		"LootdropEntries.Item.Merchantlists.NpcTypes.Spawnentries",
		"LootdropEntries.Item.Merchantlists.NpcTypes.Spawnentries.NpcType",
		"LootdropEntries.Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup",
		"LootdropEntries.Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2",
		"LootdropEntries.Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries",
		"LootdropEntries.Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup",
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
		"LoottableEntries",
		"LoottableEntries.Lootdrop",
		"LoottableEntries.Loottable",
		"LoottableEntries.Loottable.LoottableEntries",
		"LoottableEntries.Loottable.NpcTypes",
		"LoottableEntries.Loottable.NpcTypes.AlternateCurrency",
		"LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item",
		"LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.AlternateCurrencies",
		"LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.CharacterCorpseItems",
		"LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.DiscoveredItems",
		"LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.Doors",
		"LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.Doors.Item",
		"LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.Fishings",
		"LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.Fishings.Item",
		"LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.Fishings.NpcType",
		"LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.Fishings.Zone",
		"LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.Forages",
		"LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.Forages.Item",
		"LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.Forages.Zone",
		"LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.GroundSpawns",
		"LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.GroundSpawns.Zone",
		"LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.ItemTicks",
		"LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.Keyrings",
		"LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.LootdropEntries",
		"LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.LootdropEntries.Item",
		"LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.LootdropEntries.Lootdrop",
		"LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.Merchantlists",
		"LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.Merchantlists.Items",
		"LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.Merchantlists.NpcTypes",
		"LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.ObjectContents",
		"LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.Objects",
		"LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.Objects.Item",
		"LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.Objects.Zone",
		"LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.StartingItems",
		"LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.StartingItems.Item",
		"LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.StartingItems.Zone",
		"LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.TradeskillRecipeEntries",
		"LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.TradeskillRecipeEntries.TradeskillRecipe",
		"LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item.TributeLevels",
		"LoottableEntries.Loottable.NpcTypes.Loottable",
		"LoottableEntries.Loottable.NpcTypes.Merchantlists",
		"LoottableEntries.Loottable.NpcTypes.Merchantlists.Items",
		"LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.AlternateCurrencies",
		"LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.AlternateCurrencies.Item",
		"LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.CharacterCorpseItems",
		"LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.DiscoveredItems",
		"LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.Doors",
		"LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.Doors.Item",
		"LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.Fishings",
		"LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.Fishings.Item",
		"LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.Fishings.NpcType",
		"LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.Fishings.Zone",
		"LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.Forages",
		"LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.Forages.Item",
		"LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.Forages.Zone",
		"LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.GroundSpawns",
		"LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.GroundSpawns.Zone",
		"LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.ItemTicks",
		"LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.Keyrings",
		"LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.LootdropEntries",
		"LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.LootdropEntries.Item",
		"LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.LootdropEntries.Lootdrop",
		"LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.Merchantlists",
		"LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.ObjectContents",
		"LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.Objects",
		"LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.Objects.Item",
		"LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.Objects.Zone",
		"LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.StartingItems",
		"LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.StartingItems.Item",
		"LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.StartingItems.Zone",
		"LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.TradeskillRecipeEntries",
		"LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.TradeskillRecipeEntries.TradeskillRecipe",
		"LoottableEntries.Loottable.NpcTypes.Merchantlists.Items.TributeLevels",
		"LoottableEntries.Loottable.NpcTypes.Merchantlists.NpcTypes",
		"LoottableEntries.Loottable.NpcTypes.NpcEmotes",
		"LoottableEntries.Loottable.NpcTypes.NpcFactions",
		"LoottableEntries.Loottable.NpcTypes.NpcFactions.NpcFactionEntries",
		"LoottableEntries.Loottable.NpcTypes.NpcFactions.NpcFactionEntries.FactionList",
		"LoottableEntries.Loottable.NpcTypes.NpcSpell",
		"LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpell",
		"LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries",
		"LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew",
		"LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura",
		"LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew",
		"LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells",
		"LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes",
		"LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items",
		"LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.AlternateCurrencies",
		"LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.AlternateCurrencies.Item",
		"LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.CharacterCorpseItems",
		"LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.DiscoveredItems",
		"LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Doors",
		"LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Doors.Item",
		"LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings",
		"LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.Item",
		"LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType",
		"LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Fishings.Zone",
		"LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Forages",
		"LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Forages.Item",
		"LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Forages.Zone",
		"LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.GroundSpawns",
		"LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.GroundSpawns.Zone",
		"LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.ItemTicks",
		"LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Keyrings",
		"LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries",
		"LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Item",
		"LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop",
		"LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists",
		"LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists.Items",
		"LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes",
		"LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.ObjectContents",
		"LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Objects",
		"LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Objects.Item",
		"LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.Objects.Zone",
		"LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.StartingItems",
		"LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.StartingItems.Item",
		"LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.StartingItems.Zone",
		"LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.TradeskillRecipeEntries",
		"LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.TradeskillRecipeEntries.TradeskillRecipe",
		"LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items.TributeLevels",
		"LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries",
		"LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets",
		"LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals",
		"LoottableEntries.Loottable.NpcTypes.NpcTypesTint",
		"LoottableEntries.Loottable.NpcTypes.Spawnentries",
		"LoottableEntries.Loottable.NpcTypes.Spawnentries.NpcType",
		"LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup",
		"LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2",
		"LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries",
		"LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup",
	}
}

func (Lootdrop) Connection() string {
    return "eqemu_content"
}
