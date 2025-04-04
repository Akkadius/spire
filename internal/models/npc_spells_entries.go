package models

import (
	"github.com/volatiletech/null/v8"
)

type NpcSpellsEntry struct {
	ID                     uint        `json:"id" gorm:"Column:id"`
	NpcSpellsId            int         `json:"npc_spells_id" gorm:"Column:npc_spells_id"`
	Spellid                uint16      `json:"spellid" gorm:"Column:spellid"`
	Type                   uint        `json:"type" gorm:"Column:type"`
	Minlevel               uint8       `json:"minlevel" gorm:"Column:minlevel"`
	Maxlevel               uint8       `json:"maxlevel" gorm:"Column:maxlevel"`
	Manacost               int16       `json:"manacost" gorm:"Column:manacost"`
	RecastDelay            int         `json:"recast_delay" gorm:"Column:recast_delay"`
	Priority               int16       `json:"priority" gorm:"Column:priority"`
	ResistAdjust           null.Int    `json:"resist_adjust" gorm:"Column:resist_adjust"`
	MinHp                  null.Int16  `json:"min_hp" gorm:"Column:min_hp"`
	MaxHp                  null.Int16  `json:"max_hp" gorm:"Column:max_hp"`
	MinExpansion           int8        `json:"min_expansion" gorm:"Column:min_expansion"`
	MaxExpansion           int8        `json:"max_expansion" gorm:"Column:max_expansion"`
	ContentFlags           null.String `json:"content_flags" gorm:"Column:content_flags"`
	ContentFlagsDisabled   null.String `json:"content_flags_disabled" gorm:"Column:content_flags_disabled"`
	SpellsNew              *SpellsNew  `json:"spells_new,omitempty" gorm:"foreignKey:spellid;references:id"`
}

func (NpcSpellsEntry) TableName() string {
    return "npc_spells_entries"
}

func (NpcSpellsEntry) Relationships() []string {
    return []string{
		"SpellsNew",
		"SpellsNew.Aura",
		"SpellsNew.Aura.SpellsNew",
		"SpellsNew.BlockedSpells",
		"SpellsNew.BotSpellsEntries",
		"SpellsNew.BotSpellsEntries.NpcSpell",
		"SpellsNew.BotSpellsEntries.NpcSpell.BotSpellsEntries",
		"SpellsNew.BotSpellsEntries.NpcSpell.NpcSpell",
		"SpellsNew.BotSpellsEntries.NpcSpell.NpcSpellsEntries",
		"SpellsNew.BotSpellsEntries.SpellsNew",
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
		"SpellsNew.Items.Fishings.NpcType.NpcSpell.BotSpellsEntries.NpcSpell",
		"SpellsNew.Items.Fishings.NpcType.NpcSpell.BotSpellsEntries.SpellsNew",
		"SpellsNew.Items.Fishings.NpcType.NpcSpell.NpcSpell",
		"SpellsNew.Items.Fishings.NpcType.NpcSpell.NpcSpellsEntries",
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
		"SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.BotSpellsEntries.NpcSpell",
		"SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew",
		"SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpell",
		"SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries",
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
		"SpellsNew.Items.Merchantlists.NpcTypes.NpcSpell.BotSpellsEntries.NpcSpell",
		"SpellsNew.Items.Merchantlists.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew",
		"SpellsNew.Items.Merchantlists.NpcTypes.NpcSpell.NpcSpell",
		"SpellsNew.Items.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries",
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
		"SpellsNew.SpellBuckets",
		"SpellsNew.SpellGlobals",
	}
}

func (NpcSpellsEntry) Connection() string {
    return "eqemu_content"
}
