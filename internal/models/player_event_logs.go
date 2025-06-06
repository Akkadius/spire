package models

import (
	"github.com/volatiletech/null/v8"
)

type PlayerEventLog struct {
	ID              int64           `json:"id" gorm:"Column:id"`
	AccountId       null.Int64      `json:"account_id" gorm:"Column:account_id"`
	CharacterId     null.Int64      `json:"character_id" gorm:"Column:character_id"`
	ZoneId          null.Int        `json:"zone_id" gorm:"Column:zone_id"`
	InstanceId      null.Int        `json:"instance_id" gorm:"Column:instance_id"`
	X               null.Float32    `json:"x" gorm:"Column:x"`
	Y               null.Float32    `json:"y" gorm:"Column:y"`
	Z               null.Float32    `json:"z" gorm:"Column:z"`
	Heading         null.Float32    `json:"heading" gorm:"Column:heading"`
	EventTypeId     null.Int        `json:"event_type_id" gorm:"Column:event_type_id"`
	EventTypeName   null.String     `json:"event_type_name" gorm:"Column:event_type_name"`
	EventData       null.String     `json:"event_data" gorm:"Column:event_data"`
	EtlTableId      int64           `json:"etl_table_id" gorm:"Column:etl_table_id"`
	CreatedAt       null.Time       `json:"created_at" gorm:"Column:created_at"`
	Account         *Account        `json:"account,omitempty" gorm:"foreignKey:account_id;references:id"`
	CharacterDatum  *CharacterDatum `json:"character_datum,omitempty" gorm:"foreignKey:character_id;references:id"`
	Zone            *Zone           `json:"zone,omitempty" gorm:"foreignKey:zone_id;references:zoneidnumber"`
}

func (PlayerEventLog) TableName() string {
    return "player_event_logs"
}

func (PlayerEventLog) Relationships() []string {
    return []string{
		"Account",
		"Account.AccountFlags",
		"Account.AccountIps",
		"Account.AccountRewards",
		"Account.BugReports",
		"CharacterDatum",
		"CharacterDatum.AdventureStats",
		"CharacterDatum.Buyers",
		"CharacterDatum.CharRecipeLists",
		"CharacterDatum.CharacterActivities",
		"CharacterDatum.CharacterAltCurrencies",
		"CharacterDatum.CharacterAlternateAbilities",
		"CharacterDatum.CharacterAuras",
		"CharacterDatum.CharacterBandoliers",
		"CharacterDatum.CharacterBinds",
		"CharacterDatum.CharacterBuffs",
		"CharacterDatum.CharacterCorpses",
		"CharacterDatum.CharacterCurrencies",
		"CharacterDatum.CharacterDisciplines",
		"CharacterDatum.CharacterEnabledtasks",
		"CharacterDatum.CharacterInspectMessages",
		"CharacterDatum.CharacterItemRecasts",
		"CharacterDatum.CharacterLanguages",
		"CharacterDatum.CharacterLeadershipAbilities",
		"CharacterDatum.CharacterMaterials",
		"CharacterDatum.CharacterMemmedSpells",
		"CharacterDatum.CharacterPetBuffs",
		"CharacterDatum.CharacterPetInfos",
		"CharacterDatum.CharacterPetInventories",
		"CharacterDatum.CharacterPotionbelts",
		"CharacterDatum.CharacterSkills",
		"CharacterDatum.CharacterSpells",
		"CharacterDatum.CharacterTasks",
		"CharacterDatum.CharacterTributes",
		"CharacterDatum.CompletedTasks",
		"CharacterDatum.DataBuckets",
		"CharacterDatum.FactionValues",
		"CharacterDatum.Friends",
		"CharacterDatum.Guild",
		"CharacterDatum.Guild.GuildBanks",
		"CharacterDatum.Guild.GuildMembers",
		"CharacterDatum.Guild.GuildRanks",
		"CharacterDatum.GuildMembers",
		"CharacterDatum.InstanceListPlayers",
		"CharacterDatum.Inventories",
		"CharacterDatum.Inventories.Item",
		"CharacterDatum.Inventories.Item.AlternateCurrencies",
		"CharacterDatum.Inventories.Item.AlternateCurrencies.Item",
		"CharacterDatum.Inventories.Item.CharacterCorpseItems",
		"CharacterDatum.Inventories.Item.DiscoveredItems",
		"CharacterDatum.Inventories.Item.Doors",
		"CharacterDatum.Inventories.Item.Doors.Item",
		"CharacterDatum.Inventories.Item.Fishings",
		"CharacterDatum.Inventories.Item.Fishings.Item",
		"CharacterDatum.Inventories.Item.Fishings.NpcType",
		"CharacterDatum.Inventories.Item.Fishings.NpcType.AlternateCurrency",
		"CharacterDatum.Inventories.Item.Fishings.NpcType.AlternateCurrency.Item",
		"CharacterDatum.Inventories.Item.Fishings.NpcType.Loottable",
		"CharacterDatum.Inventories.Item.Fishings.NpcType.Loottable.LoottableEntries",
		"CharacterDatum.Inventories.Item.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop",
		"CharacterDatum.Inventories.Item.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries",
		"CharacterDatum.Inventories.Item.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item",
		"CharacterDatum.Inventories.Item.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Lootdrop",
		"CharacterDatum.Inventories.Item.Fishings.NpcType.Loottable.LoottableEntries.Lootdrop.LoottableEntries",
		"CharacterDatum.Inventories.Item.Fishings.NpcType.Loottable.LoottableEntries.Loottable",
		"CharacterDatum.Inventories.Item.Fishings.NpcType.Loottable.NpcTypes",
		"CharacterDatum.Inventories.Item.Fishings.NpcType.Merchantlists",
		"CharacterDatum.Inventories.Item.Fishings.NpcType.Merchantlists.Items",
		"CharacterDatum.Inventories.Item.Fishings.NpcType.Merchantlists.NpcTypes",
		"CharacterDatum.Inventories.Item.Fishings.NpcType.NpcEmotes",
		"CharacterDatum.Inventories.Item.Fishings.NpcType.NpcFactions",
		"CharacterDatum.Inventories.Item.Fishings.NpcType.NpcFactions.NpcFactionEntries",
		"CharacterDatum.Inventories.Item.Fishings.NpcType.NpcFactions.NpcFactionEntries.FactionList",
		"CharacterDatum.Inventories.Item.Fishings.NpcType.NpcSpell",
		"CharacterDatum.Inventories.Item.Fishings.NpcType.NpcSpell.BotSpellsEntries",
		"CharacterDatum.Inventories.Item.Fishings.NpcType.NpcSpell.BotSpellsEntries.NpcSpell",
		"CharacterDatum.Inventories.Item.Fishings.NpcType.NpcSpell.BotSpellsEntries.SpellsNew",
		"CharacterDatum.Inventories.Item.Fishings.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Aura",
		"CharacterDatum.Inventories.Item.Fishings.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Aura.SpellsNew",
		"CharacterDatum.Inventories.Item.Fishings.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.BlockedSpells",
		"CharacterDatum.Inventories.Item.Fishings.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.BotSpellsEntries",
		"CharacterDatum.Inventories.Item.Fishings.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Damageshieldtypes",
		"CharacterDatum.Inventories.Item.Fishings.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.Items",
		"CharacterDatum.Inventories.Item.Fishings.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.NpcSpellsEntries",
		"CharacterDatum.Inventories.Item.Fishings.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.NpcSpellsEntries.SpellsNew",
		"CharacterDatum.Inventories.Item.Fishings.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.SpellBuckets",
		"CharacterDatum.Inventories.Item.Fishings.NpcType.NpcSpell.BotSpellsEntries.SpellsNew.SpellGlobals",
		"CharacterDatum.Inventories.Item.Fishings.NpcType.NpcSpell.NpcSpell",
		"CharacterDatum.Inventories.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries",
		"CharacterDatum.Inventories.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew",
		"CharacterDatum.Inventories.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Aura",
		"CharacterDatum.Inventories.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew",
		"CharacterDatum.Inventories.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells",
		"CharacterDatum.Inventories.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.BotSpellsEntries",
		"CharacterDatum.Inventories.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.BotSpellsEntries.NpcSpell",
		"CharacterDatum.Inventories.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.BotSpellsEntries.SpellsNew",
		"CharacterDatum.Inventories.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes",
		"CharacterDatum.Inventories.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.Items",
		"CharacterDatum.Inventories.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries",
		"CharacterDatum.Inventories.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets",
		"CharacterDatum.Inventories.Item.Fishings.NpcType.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals",
		"CharacterDatum.Inventories.Item.Fishings.NpcType.NpcTypesTint",
		"CharacterDatum.Inventories.Item.Fishings.NpcType.Spawnentries",
		"CharacterDatum.Inventories.Item.Fishings.NpcType.Spawnentries.NpcType",
		"CharacterDatum.Inventories.Item.Fishings.NpcType.Spawnentries.Spawngroup",
		"CharacterDatum.Inventories.Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2",
		"CharacterDatum.Inventories.Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries",
		"CharacterDatum.Inventories.Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup",
		"CharacterDatum.Inventories.Item.Fishings.Zone",
		"CharacterDatum.Inventories.Item.Forages",
		"CharacterDatum.Inventories.Item.Forages.Item",
		"CharacterDatum.Inventories.Item.Forages.Zone",
		"CharacterDatum.Inventories.Item.GroundSpawns",
		"CharacterDatum.Inventories.Item.GroundSpawns.Zone",
		"CharacterDatum.Inventories.Item.ItemTicks",
		"CharacterDatum.Inventories.Item.Keyrings",
		"CharacterDatum.Inventories.Item.LootdropEntries",
		"CharacterDatum.Inventories.Item.LootdropEntries.Item",
		"CharacterDatum.Inventories.Item.LootdropEntries.Lootdrop",
		"CharacterDatum.Inventories.Item.LootdropEntries.Lootdrop.LootdropEntries",
		"CharacterDatum.Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries",
		"CharacterDatum.Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Lootdrop",
		"CharacterDatum.Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable",
		"CharacterDatum.Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.LoottableEntries",
		"CharacterDatum.Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes",
		"CharacterDatum.Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency",
		"CharacterDatum.Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency.Item",
		"CharacterDatum.Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Loottable",
		"CharacterDatum.Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists",
		"CharacterDatum.Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.Items",
		"CharacterDatum.Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.NpcTypes",
		"CharacterDatum.Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcEmotes",
		"CharacterDatum.Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions",
		"CharacterDatum.Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions.NpcFactionEntries",
		"CharacterDatum.Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions.NpcFactionEntries.FactionList",
		"CharacterDatum.Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell",
		"CharacterDatum.Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.BotSpellsEntries",
		"CharacterDatum.Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.BotSpellsEntries.NpcSpell",
		"CharacterDatum.Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew",
		"CharacterDatum.Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Aura",
		"CharacterDatum.Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Aura.SpellsNew",
		"CharacterDatum.Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.BlockedSpells",
		"CharacterDatum.Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.BotSpellsEntries",
		"CharacterDatum.Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Damageshieldtypes",
		"CharacterDatum.Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items",
		"CharacterDatum.Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.NpcSpellsEntries",
		"CharacterDatum.Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.NpcSpellsEntries.SpellsNew",
		"CharacterDatum.Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.SpellBuckets",
		"CharacterDatum.Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.SpellGlobals",
		"CharacterDatum.Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpell",
		"CharacterDatum.Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries",
		"CharacterDatum.Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew",
		"CharacterDatum.Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura",
		"CharacterDatum.Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew",
		"CharacterDatum.Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells",
		"CharacterDatum.Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.BotSpellsEntries",
		"CharacterDatum.Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.BotSpellsEntries.NpcSpell",
		"CharacterDatum.Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.BotSpellsEntries.SpellsNew",
		"CharacterDatum.Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes",
		"CharacterDatum.Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items",
		"CharacterDatum.Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries",
		"CharacterDatum.Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets",
		"CharacterDatum.Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals",
		"CharacterDatum.Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcTypesTint",
		"CharacterDatum.Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries",
		"CharacterDatum.Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.NpcType",
		"CharacterDatum.Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup",
		"CharacterDatum.Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2",
		"CharacterDatum.Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries",
		"CharacterDatum.Inventories.Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup",
		"CharacterDatum.Inventories.Item.Merchantlists",
		"CharacterDatum.Inventories.Item.Merchantlists.Items",
		"CharacterDatum.Inventories.Item.Merchantlists.NpcTypes",
		"CharacterDatum.Inventories.Item.Merchantlists.NpcTypes.AlternateCurrency",
		"CharacterDatum.Inventories.Item.Merchantlists.NpcTypes.AlternateCurrency.Item",
		"CharacterDatum.Inventories.Item.Merchantlists.NpcTypes.Loottable",
		"CharacterDatum.Inventories.Item.Merchantlists.NpcTypes.Loottable.LoottableEntries",
		"CharacterDatum.Inventories.Item.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop",
		"CharacterDatum.Inventories.Item.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries",
		"CharacterDatum.Inventories.Item.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Item",
		"CharacterDatum.Inventories.Item.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LootdropEntries.Lootdrop",
		"CharacterDatum.Inventories.Item.Merchantlists.NpcTypes.Loottable.LoottableEntries.Lootdrop.LoottableEntries",
		"CharacterDatum.Inventories.Item.Merchantlists.NpcTypes.Loottable.LoottableEntries.Loottable",
		"CharacterDatum.Inventories.Item.Merchantlists.NpcTypes.Loottable.NpcTypes",
		"CharacterDatum.Inventories.Item.Merchantlists.NpcTypes.Merchantlists",
		"CharacterDatum.Inventories.Item.Merchantlists.NpcTypes.NpcEmotes",
		"CharacterDatum.Inventories.Item.Merchantlists.NpcTypes.NpcFactions",
		"CharacterDatum.Inventories.Item.Merchantlists.NpcTypes.NpcFactions.NpcFactionEntries",
		"CharacterDatum.Inventories.Item.Merchantlists.NpcTypes.NpcFactions.NpcFactionEntries.FactionList",
		"CharacterDatum.Inventories.Item.Merchantlists.NpcTypes.NpcSpell",
		"CharacterDatum.Inventories.Item.Merchantlists.NpcTypes.NpcSpell.BotSpellsEntries",
		"CharacterDatum.Inventories.Item.Merchantlists.NpcTypes.NpcSpell.BotSpellsEntries.NpcSpell",
		"CharacterDatum.Inventories.Item.Merchantlists.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew",
		"CharacterDatum.Inventories.Item.Merchantlists.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Aura",
		"CharacterDatum.Inventories.Item.Merchantlists.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Aura.SpellsNew",
		"CharacterDatum.Inventories.Item.Merchantlists.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.BlockedSpells",
		"CharacterDatum.Inventories.Item.Merchantlists.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.BotSpellsEntries",
		"CharacterDatum.Inventories.Item.Merchantlists.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Damageshieldtypes",
		"CharacterDatum.Inventories.Item.Merchantlists.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.Items",
		"CharacterDatum.Inventories.Item.Merchantlists.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.NpcSpellsEntries",
		"CharacterDatum.Inventories.Item.Merchantlists.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.NpcSpellsEntries.SpellsNew",
		"CharacterDatum.Inventories.Item.Merchantlists.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.SpellBuckets",
		"CharacterDatum.Inventories.Item.Merchantlists.NpcTypes.NpcSpell.BotSpellsEntries.SpellsNew.SpellGlobals",
		"CharacterDatum.Inventories.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpell",
		"CharacterDatum.Inventories.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries",
		"CharacterDatum.Inventories.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew",
		"CharacterDatum.Inventories.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura",
		"CharacterDatum.Inventories.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Aura.SpellsNew",
		"CharacterDatum.Inventories.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.BlockedSpells",
		"CharacterDatum.Inventories.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.BotSpellsEntries",
		"CharacterDatum.Inventories.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.BotSpellsEntries.NpcSpell",
		"CharacterDatum.Inventories.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.BotSpellsEntries.SpellsNew",
		"CharacterDatum.Inventories.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Damageshieldtypes",
		"CharacterDatum.Inventories.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.Items",
		"CharacterDatum.Inventories.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.NpcSpellsEntries",
		"CharacterDatum.Inventories.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellBuckets",
		"CharacterDatum.Inventories.Item.Merchantlists.NpcTypes.NpcSpell.NpcSpellsEntries.SpellsNew.SpellGlobals",
		"CharacterDatum.Inventories.Item.Merchantlists.NpcTypes.NpcTypesTint",
		"CharacterDatum.Inventories.Item.Merchantlists.NpcTypes.Spawnentries",
		"CharacterDatum.Inventories.Item.Merchantlists.NpcTypes.Spawnentries.NpcType",
		"CharacterDatum.Inventories.Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup",
		"CharacterDatum.Inventories.Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2",
		"CharacterDatum.Inventories.Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries",
		"CharacterDatum.Inventories.Item.Merchantlists.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup",
		"CharacterDatum.Inventories.Item.ObjectContents",
		"CharacterDatum.Inventories.Item.Objects",
		"CharacterDatum.Inventories.Item.Objects.Item",
		"CharacterDatum.Inventories.Item.Objects.Zone",
		"CharacterDatum.Inventories.Item.TradeskillRecipeEntries",
		"CharacterDatum.Inventories.Item.TradeskillRecipeEntries.TradeskillRecipe",
		"CharacterDatum.Inventories.Item.TradeskillRecipeEntries.TradeskillRecipe.TradeskillRecipeEntries",
		"CharacterDatum.Inventories.Item.TributeLevels",
		"CharacterDatum.Keyrings",
		"CharacterDatum.Mail",
		"CharacterDatum.PlayerTitlesets",
		"CharacterDatum.QuestGlobals",
		"CharacterDatum.Timers",
		"CharacterDatum.Titles",
		"CharacterDatum.Traders",
		"CharacterDatum.ZoneFlags",
		"Zone",
	}
}

func (PlayerEventLog) Connection() string {
    return "eqemu_logs"
}
