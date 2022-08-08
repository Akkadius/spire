package models

import (
	"github.com/volatiletech/null/v8"
)

type NpcSpell struct {
	ID                       uint             `json:"id" gorm:"Column:id"`
	Name                     null.String      `json:"name" gorm:"Column:name"`
	ParentList               uint             `json:"parent_list" gorm:"Column:parent_list"`
	AttackProc               int16            `json:"attack_proc" gorm:"Column:attack_proc"`
	ProcChance               int8             `json:"proc_chance" gorm:"Column:proc_chance"`
	RangeProc                int16            `json:"range_proc" gorm:"Column:range_proc"`
	RprocChance              int16            `json:"rproc_chance" gorm:"Column:rproc_chance"`
	DefensiveProc            int16            `json:"defensive_proc" gorm:"Column:defensive_proc"`
	DprocChance              int16            `json:"dproc_chance" gorm:"Column:dproc_chance"`
	FailRecast               uint             `json:"fail_recast" gorm:"Column:fail_recast"`
	EngagedNoSpRecastMin     uint             `json:"engaged_no_sp_recast_min" gorm:"Column:engaged_no_sp_recast_min"`
	EngagedNoSpRecastMax     uint             `json:"engaged_no_sp_recast_max" gorm:"Column:engaged_no_sp_recast_max"`
	EngagedBSelfChance       uint8            `json:"engaged_b_self_chance" gorm:"Column:engaged_b_self_chance"`
	EngagedBOtherChance      uint8            `json:"engaged_b_other_chance" gorm:"Column:engaged_b_other_chance"`
	EngagedDChance           uint8            `json:"engaged_d_chance" gorm:"Column:engaged_d_chance"`
	PursueNoSpRecastMin      uint             `json:"pursue_no_sp_recast_min" gorm:"Column:pursue_no_sp_recast_min"`
	PursueNoSpRecastMax      uint             `json:"pursue_no_sp_recast_max" gorm:"Column:pursue_no_sp_recast_max"`
	PursueDChance            uint8            `json:"pursue_d_chance" gorm:"Column:pursue_d_chance"`
	IdleNoSpRecastMin        uint             `json:"idle_no_sp_recast_min" gorm:"Column:idle_no_sp_recast_min"`
	IdleNoSpRecastMax        uint             `json:"idle_no_sp_recast_max" gorm:"Column:idle_no_sp_recast_max"`
	IdleBChance              uint8            `json:"idle_b_chance" gorm:"Column:idle_b_chance"`
	NpcSpellsEntries         []NpcSpellsEntry `json:"npc_spells_entries,omitempty" gorm:"foreignKey:npc_spells_id;references:id"`
	NpcSpell                 *NpcSpell        `json:"npc_spell,omitempty" gorm:"foreignKey:parent_list;references:id"`
}

func (NpcSpell) TableName() string {
    return "npc_spells"
}

func (NpcSpell) Relationships() []string {
    return []string{
		"NpcSpell",
		"NpcSpellsEntries",
		"NpcSpellsEntries.SpellsNew",
		"NpcSpellsEntries.SpellsNew.Aura",
		"NpcSpellsEntries.SpellsNew.Aura.SpellsNew",
		"NpcSpellsEntries.SpellsNew.BlockedSpells",
		"NpcSpellsEntries.SpellsNew.Damageshieldtypes",
		"NpcSpellsEntries.SpellsNew.Items",
		"NpcSpellsEntries.SpellsNew.Items.AlternateCurrencies",
		"NpcSpellsEntries.SpellsNew.Items.AlternateCurrencies.Item",
		"NpcSpellsEntries.SpellsNew.Items.CharacterCorpseItems",
		"NpcSpellsEntries.SpellsNew.Items.DiscoveredItems",
		"NpcSpellsEntries.SpellsNew.Items.Doors",
		"NpcSpellsEntries.SpellsNew.Items.Doors.Item",
		"NpcSpellsEntries.SpellsNew.Items.Fishings",
		"NpcSpellsEntries.SpellsNew.Items.Fishings.Item",
		"NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType",
		"NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.AlternateCurrency",
		"NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.AlternateCurrency.Item",
		"NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.Loottable",
		"NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.Loottable.LoottableEntries",
		"NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop",
		"NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries",
		"NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item",
		"NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Lootdrop",
		"NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LoottableEntries",
		"NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.Loottable.LoottableEntries.Loottable",
		"NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.Loottable.NpcTypes",
		"NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.Merchantlists",
		"NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.Merchantlists.Items",
		"NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.Merchantlists.NpcTypes",
		"NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.NpcEmotes",
		"NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.NpcFactions",
		"NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.NpcFactions.NpcFactionEntries",
		"NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.NpcFactions.NpcFactionEntries.FactionList",
		"NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.NpcSpell",
		"NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.NpcTypesTint",
		"NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.Spawnentries",
		"NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.Spawnentries.NpcType",
		"NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.Spawnentries.Spawngroup",
		"NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2",
		"NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries",
		"NpcSpellsEntries.SpellsNew.Items.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup",
		"NpcSpellsEntries.SpellsNew.Items.Fishings.Zone",
		"NpcSpellsEntries.SpellsNew.Items.Forages",
		"NpcSpellsEntries.SpellsNew.Items.Forages.Item",
		"NpcSpellsEntries.SpellsNew.Items.Forages.Zone",
		"NpcSpellsEntries.SpellsNew.Items.GroundSpawns",
		"NpcSpellsEntries.SpellsNew.Items.GroundSpawns.Zone",
		"NpcSpellsEntries.SpellsNew.Items.ItemTicks",
		"NpcSpellsEntries.SpellsNew.Items.Keyrings",
		"NpcSpellsEntries.SpellsNew.Items.LootdropEntries",
		"NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Item",
		"NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop",
		"NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LootdropEntries",
		"NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries",
		"NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Lootdrop",
		"NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable",
		"NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.LoottableEntries",
		"NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes",
		"NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency",
		"NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item",
		"NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Loottable",
		"NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists",
		"NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items",
		"NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.NpcTypes",
		"NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcEmotes",
		"NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions",
		"NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions.NpcFactionEntries",
		"NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions.NpcFactionEntries.FactionList",
		"NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell",
		"NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcTypesTint",
		"NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries",
		"NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.NpcType",
		"NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup",
		"NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2",
		"NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries",
		"NpcSpellsEntries.SpellsNew.Items.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup",
		"NpcSpellsEntries.SpellsNew.Items.Merchantlists",
		"NpcSpellsEntries.SpellsNew.Items.Merchantlists.Items",
		"NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes",
		"NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.AlternateCurrency",
		"NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.AlternateCurrency.Item",
		"NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.Loottable",
		"NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.Loottable.LoottableEntries",
		"NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop",
		"NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries",
		"NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item",
		"NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Lootdrop",
		"NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LoottableEntries",
		"NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.Loottable.LoottableEntries.Loottable",
		"NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.Loottable.NpcTypes",
		"NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.Merchantlists",
		"NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.NpcEmotes",
		"NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.NpcFactions",
		"NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.NpcFactions.NpcFactionEntries",
		"NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.NpcFactions.NpcFactionEntries.FactionList",
		"NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.NpcSpell",
		"NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.NpcTypesTint",
		"NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.Spawnentries",
		"NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.Spawnentries.NpcType",
		"NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.Spawnentries.Spawngroup",
		"NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2",
		"NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries",
		"NpcSpellsEntries.SpellsNew.Items.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup",
		"NpcSpellsEntries.SpellsNew.Items.ObjectContents",
		"NpcSpellsEntries.SpellsNew.Items.Objects",
		"NpcSpellsEntries.SpellsNew.Items.Objects.Item",
		"NpcSpellsEntries.SpellsNew.Items.Objects.Zone",
		"NpcSpellsEntries.SpellsNew.Items.StartingItems",
		"NpcSpellsEntries.SpellsNew.Items.StartingItems.Item",
		"NpcSpellsEntries.SpellsNew.Items.StartingItems.Zone",
		"NpcSpellsEntries.SpellsNew.Items.Tasks",
		"NpcSpellsEntries.SpellsNew.Items.Tasks.AlternateCurrency",
		"NpcSpellsEntries.SpellsNew.Items.Tasks.AlternateCurrency.Item",
		"NpcSpellsEntries.SpellsNew.Items.Tasks.TaskActivities",
		"NpcSpellsEntries.SpellsNew.Items.Tasks.TaskActivities.Goallists",
		"NpcSpellsEntries.SpellsNew.Items.Tasks.TaskActivities.NpcType",
		"NpcSpellsEntries.SpellsNew.Items.Tasks.TaskActivities.NpcType.AlternateCurrency",
		"NpcSpellsEntries.SpellsNew.Items.Tasks.TaskActivities.NpcType.AlternateCurrency.Item",
		"NpcSpellsEntries.SpellsNew.Items.Tasks.TaskActivities.NpcType.Loottable",
		"NpcSpellsEntries.SpellsNew.Items.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries",
		"NpcSpellsEntries.SpellsNew.Items.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.Lootdrop",
		"NpcSpellsEntries.SpellsNew.Items.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries",
		"NpcSpellsEntries.SpellsNew.Items.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item",
		"NpcSpellsEntries.SpellsNew.Items.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Lootdrop",
		"NpcSpellsEntries.SpellsNew.Items.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.Lootdrop.LoottableEntries",
		"NpcSpellsEntries.SpellsNew.Items.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.Loottable",
		"NpcSpellsEntries.SpellsNew.Items.Tasks.TaskActivities.NpcType.Loottable.NpcTypes",
		"NpcSpellsEntries.SpellsNew.Items.Tasks.TaskActivities.NpcType.Merchantlists",
		"NpcSpellsEntries.SpellsNew.Items.Tasks.TaskActivities.NpcType.Merchantlists.Items",
		"NpcSpellsEntries.SpellsNew.Items.Tasks.TaskActivities.NpcType.Merchantlists.NpcTypes",
		"NpcSpellsEntries.SpellsNew.Items.Tasks.TaskActivities.NpcType.NpcEmotes",
		"NpcSpellsEntries.SpellsNew.Items.Tasks.TaskActivities.NpcType.NpcFactions",
		"NpcSpellsEntries.SpellsNew.Items.Tasks.TaskActivities.NpcType.NpcFactions.NpcFactionEntries",
		"NpcSpellsEntries.SpellsNew.Items.Tasks.TaskActivities.NpcType.NpcFactions.NpcFactionEntries.FactionList",
		"NpcSpellsEntries.SpellsNew.Items.Tasks.TaskActivities.NpcType.NpcSpell",
		"NpcSpellsEntries.SpellsNew.Items.Tasks.TaskActivities.NpcType.NpcTypesTint",
		"NpcSpellsEntries.SpellsNew.Items.Tasks.TaskActivities.NpcType.Spawnentries",
		"NpcSpellsEntries.SpellsNew.Items.Tasks.TaskActivities.NpcType.Spawnentries.NpcType",
		"NpcSpellsEntries.SpellsNew.Items.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup",
		"NpcSpellsEntries.SpellsNew.Items.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2",
		"NpcSpellsEntries.SpellsNew.Items.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries",
		"NpcSpellsEntries.SpellsNew.Items.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup",
		"NpcSpellsEntries.SpellsNew.Items.Tasks.Tasksets",
		"NpcSpellsEntries.SpellsNew.Items.TradeskillRecipeEntries",
		"NpcSpellsEntries.SpellsNew.Items.TradeskillRecipeEntries.TradeskillRecipe",
		"NpcSpellsEntries.SpellsNew.Items.TributeLevels",
		"NpcSpellsEntries.SpellsNew.NpcSpellsEntries",
		"NpcSpellsEntries.SpellsNew.SpellBuckets",
		"NpcSpellsEntries.SpellsNew.SpellGlobals",
	}
}

func (NpcSpell) Connection() string {
    return "eqemu_content"
}
