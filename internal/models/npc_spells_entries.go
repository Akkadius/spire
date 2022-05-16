package models

import (
	"github.com/volatiletech/null/v8"
)

type NpcSpellsEntry struct {
	ID            uint       `json:"id" gorm:"Column:id"`
	NpcSpellsId   int        `json:"npc_spells_id" gorm:"Column:npc_spells_id"`
	Spellid       uint16     `json:"spellid" gorm:"Column:spellid"`
	Type          uint       `json:"type" gorm:"Column:type"`
	Minlevel      uint8      `json:"minlevel" gorm:"Column:minlevel"`
	Maxlevel      uint8      `json:"maxlevel" gorm:"Column:maxlevel"`
	Manacost      int16      `json:"manacost" gorm:"Column:manacost"`
	RecastDelay   int        `json:"recast_delay" gorm:"Column:recast_delay"`
	Priority      int16      `json:"priority" gorm:"Column:priority"`
	ResistAdjust  null.Int   `json:"resist_adjust" gorm:"Column:resist_adjust"`
	MinHp         null.Int16 `json:"min_hp" gorm:"Column:min_hp"`
	MaxHp         null.Int16 `json:"max_hp" gorm:"Column:max_hp"`
	SpellsNew     *SpellsNew `json:"spells_new,omitempty" gorm:"foreignKey:spellid;references:id"`
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
		"SpellsNew.Items.Fishings.NpcType.Loottable.LoottableEntries.LootdropEntries",
		"SpellsNew.Items.Fishings.NpcType.Loottable.LoottableEntries.LootdropEntries.Item",
		"SpellsNew.Items.Fishings.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop",
		"SpellsNew.Items.Fishings.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LootdropEntries",
		"SpellsNew.Items.Fishings.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LoottableEntries",
		"SpellsNew.Items.Fishings.NpcType.Loottable.LoottableEntries.Loottable",
		"SpellsNew.Items.Fishings.NpcType.Loottable.NpcTypes",
		"SpellsNew.Items.Fishings.NpcType.Merchantlists",
		"SpellsNew.Items.Fishings.NpcType.Merchantlists.NpcType",
		"SpellsNew.Items.Fishings.NpcType.NpcEmotes",
		"SpellsNew.Items.Fishings.NpcType.NpcFactions",
		"SpellsNew.Items.Fishings.NpcType.NpcFactions.NpcFactionEntries",
		"SpellsNew.Items.Fishings.NpcType.NpcSpell",
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
		"SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.LootdropEntries",
		"SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable",
		"SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.LoottableEntries",
		"SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes",
		"SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency",
		"SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Loottable",
		"SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists",
		"SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.NpcType",
		"SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcEmotes",
		"SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions",
		"SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions.NpcFactionEntries",
		"SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell",
		"SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries",
		"SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcTypesTint",
		"SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries",
		"SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.NpcType",
		"SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup",
		"SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2",
		"SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries",
		"SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup",
		"SpellsNew.Items.Merchantlists",
		"SpellsNew.Items.Merchantlists.NpcType",
		"SpellsNew.Items.Merchantlists.NpcType.AlternateCurrency",
		"SpellsNew.Items.Merchantlists.NpcType.Loottable",
		"SpellsNew.Items.Merchantlists.NpcType.Loottable.LoottableEntries",
		"SpellsNew.Items.Merchantlists.NpcType.Loottable.LoottableEntries.LootdropEntries",
		"SpellsNew.Items.Merchantlists.NpcType.Loottable.LoottableEntries.LootdropEntries.Item",
		"SpellsNew.Items.Merchantlists.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop",
		"SpellsNew.Items.Merchantlists.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LootdropEntries",
		"SpellsNew.Items.Merchantlists.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LoottableEntries",
		"SpellsNew.Items.Merchantlists.NpcType.Loottable.LoottableEntries.Loottable",
		"SpellsNew.Items.Merchantlists.NpcType.Loottable.NpcTypes",
		"SpellsNew.Items.Merchantlists.NpcType.Merchantlists",
		"SpellsNew.Items.Merchantlists.NpcType.NpcEmotes",
		"SpellsNew.Items.Merchantlists.NpcType.NpcFactions",
		"SpellsNew.Items.Merchantlists.NpcType.NpcFactions.NpcFactionEntries",
		"SpellsNew.Items.Merchantlists.NpcType.NpcSpell",
		"SpellsNew.Items.Merchantlists.NpcType.NpcSpell.NpcSpellsEntries",
		"SpellsNew.Items.Merchantlists.NpcType.NpcTypesTint",
		"SpellsNew.Items.Merchantlists.NpcType.Spawnentries",
		"SpellsNew.Items.Merchantlists.NpcType.Spawnentries.NpcType",
		"SpellsNew.Items.Merchantlists.NpcType.Spawnentries.Spawngroup",
		"SpellsNew.Items.Merchantlists.NpcType.Spawnentries.Spawngroup.Spawn2",
		"SpellsNew.Items.Merchantlists.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries",
		"SpellsNew.Items.Merchantlists.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup",
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
		"SpellsNew.Items.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries",
		"SpellsNew.Items.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries.Item",
		"SpellsNew.Items.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop",
		"SpellsNew.Items.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LootdropEntries",
		"SpellsNew.Items.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LoottableEntries",
		"SpellsNew.Items.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.Loottable",
		"SpellsNew.Items.Tasks.TaskActivities.NpcType.Loottable.NpcTypes",
		"SpellsNew.Items.Tasks.TaskActivities.NpcType.Merchantlists",
		"SpellsNew.Items.Tasks.TaskActivities.NpcType.Merchantlists.NpcType",
		"SpellsNew.Items.Tasks.TaskActivities.NpcType.NpcEmotes",
		"SpellsNew.Items.Tasks.TaskActivities.NpcType.NpcFactions",
		"SpellsNew.Items.Tasks.TaskActivities.NpcType.NpcFactions.NpcFactionEntries",
		"SpellsNew.Items.Tasks.TaskActivities.NpcType.NpcSpell",
		"SpellsNew.Items.Tasks.TaskActivities.NpcType.NpcSpell.NpcSpellsEntries",
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
		"SpellsNew.SpellBuckets",
		"SpellsNew.SpellGlobals",
	}
}

func (NpcSpellsEntry) Connection() string {
    return "eqemu_content"
}
