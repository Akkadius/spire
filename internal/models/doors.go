package models

import (
	"github.com/volatiletech/null/v8"
)

type Door struct {
	ID                     int         `json:"id" gorm:"Column:id"`
	Doorid                 int16       `json:"doorid" gorm:"Column:doorid"`
	Zone                   null.String `json:"zone" gorm:"Column:zone"`
	Version                int16       `json:"version" gorm:"Column:version"`
	Name                   string      `json:"name" gorm:"Column:name"`
	PosY                   float32     `json:"pos_y" gorm:"Column:pos_y"`
	PosX                   float32     `json:"pos_x" gorm:"Column:pos_x"`
	PosZ                   float32     `json:"pos_z" gorm:"Column:pos_z"`
	Heading                float32     `json:"heading" gorm:"Column:heading"`
	Opentype               int16       `json:"opentype" gorm:"Column:opentype"`
	Guild                  int16       `json:"guild" gorm:"Column:guild"`
	Lockpick               int16       `json:"lockpick" gorm:"Column:lockpick"`
	Keyitem                int         `json:"keyitem" gorm:"Column:keyitem"`
	Nokeyring              uint8       `json:"nokeyring" gorm:"Column:nokeyring"`
	Triggerdoor            int16       `json:"triggerdoor" gorm:"Column:triggerdoor"`
	Triggertype            int16       `json:"triggertype" gorm:"Column:triggertype"`
	DisableTimer           int8        `json:"disable_timer" gorm:"Column:disable_timer"`
	Doorisopen             int16       `json:"doorisopen" gorm:"Column:doorisopen"`
	DoorParam              int         `json:"door_param" gorm:"Column:door_param"`
	DestZone               null.String `json:"dest_zone" gorm:"Column:dest_zone"`
	DestInstance           uint        `json:"dest_instance" gorm:"Column:dest_instance"`
	DestX                  float32     `json:"dest_x" gorm:"Column:dest_x"`
	DestY                  float32     `json:"dest_y" gorm:"Column:dest_y"`
	DestZ                  float32     `json:"dest_z" gorm:"Column:dest_z"`
	DestHeading            float32     `json:"dest_heading" gorm:"Column:dest_heading"`
	InvertState            int         `json:"invert_state" gorm:"Column:invert_state"`
	Incline                int         `json:"incline" gorm:"Column:incline"`
	Size                   uint16      `json:"size" gorm:"Column:size"`
	Buffer                 float32     `json:"buffer" gorm:"Column:buffer"`
	ClientVersionMask      uint        `json:"client_version_mask" gorm:"Column:client_version_mask"`
	IsLdonDoor             int16       `json:"is_ldon_door" gorm:"Column:is_ldon_door"`
	MinExpansion           int8        `json:"min_expansion" gorm:"Column:min_expansion"`
	MaxExpansion           int8        `json:"max_expansion" gorm:"Column:max_expansion"`
	ContentFlags           null.String `json:"content_flags" gorm:"Column:content_flags"`
	ContentFlagsDisabled   null.String `json:"content_flags_disabled" gorm:"Column:content_flags_disabled"`
	Item                   *Item       `json:"item,omitempty" gorm:"foreignKey:keyitem;references:id"`
}

func (Door) TableName() string {
    return "doors"
}

func (Door) Relationships() []string {
    return []string{
		"Item",
		"Item.AlternateCurrencies",
		"Item.CharacterCorpseItems",
		"Item.DiscoveredItems",
		"Item.Doors",
		"Item.Fishings",
		"Item.Fishings.Item",
		"Item.Fishings.NpcType",
		"Item.Fishings.NpcType.AlternateCurrency",
		"Item.Fishings.NpcType.Loottable",
		"Item.Fishings.NpcType.Loottable.LoottableEntries",
		"Item.Fishings.NpcType.Loottable.LoottableEntries.LootdropEntries",
		"Item.Fishings.NpcType.Loottable.LoottableEntries.LootdropEntries.Item",
		"Item.Fishings.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop",
		"Item.Fishings.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LootdropEntries",
		"Item.Fishings.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LoottableEntries",
		"Item.Fishings.NpcType.Loottable.LoottableEntries.Loottable",
		"Item.Fishings.NpcType.Loottable.NpcTypes",
		"Item.Fishings.NpcType.Merchantlists",
		"Item.Fishings.NpcType.Merchantlists.NpcType",
		"Item.Fishings.NpcType.NpcEmotes",
		"Item.Fishings.NpcType.NpcFactions",
		"Item.Fishings.NpcType.NpcFactions.NpcFactionEntries",
		"Item.Fishings.NpcType.NpcSpells",
		"Item.Fishings.NpcType.NpcSpells.NpcSpellsEntries",
		"Item.Fishings.NpcType.NpcTypesTint",
		"Item.Fishings.NpcType.Spawnentries",
		"Item.Fishings.NpcType.Spawnentries.NpcType",
		"Item.Fishings.NpcType.Spawnentries.Spawngroup",
		"Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2",
		"Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries",
		"Item.Fishings.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup",
		"Item.Fishings.Zone",
		"Item.Forages",
		"Item.Forages.Item",
		"Item.Forages.Zone",
		"Item.GroundSpawns",
		"Item.GroundSpawns.Zone",
		"Item.ItemTicks",
		"Item.Keyrings",
		"Item.LootdropEntries",
		"Item.LootdropEntries.Item",
		"Item.LootdropEntries.Lootdrop",
		"Item.LootdropEntries.Lootdrop.LootdropEntries",
		"Item.LootdropEntries.Lootdrop.LoottableEntries",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.LootdropEntries",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.LoottableEntries",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.AlternateCurrency",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Loottable",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Merchantlists.NpcType",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcEmotes",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcFactions.NpcFactionEntries",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpells",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcSpells.NpcSpellsEntries",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.NpcTypesTint",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.NpcType",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawnentries",
		"Item.LootdropEntries.Lootdrop.LoottableEntries.Loottable.NpcTypes.Spawnentries.Spawngroup.Spawn2.Spawngroup",
		"Item.Merchantlists",
		"Item.Merchantlists.NpcType",
		"Item.Merchantlists.NpcType.AlternateCurrency",
		"Item.Merchantlists.NpcType.Loottable",
		"Item.Merchantlists.NpcType.Loottable.LoottableEntries",
		"Item.Merchantlists.NpcType.Loottable.LoottableEntries.LootdropEntries",
		"Item.Merchantlists.NpcType.Loottable.LoottableEntries.LootdropEntries.Item",
		"Item.Merchantlists.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop",
		"Item.Merchantlists.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LootdropEntries",
		"Item.Merchantlists.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LoottableEntries",
		"Item.Merchantlists.NpcType.Loottable.LoottableEntries.Loottable",
		"Item.Merchantlists.NpcType.Loottable.NpcTypes",
		"Item.Merchantlists.NpcType.Merchantlists",
		"Item.Merchantlists.NpcType.NpcEmotes",
		"Item.Merchantlists.NpcType.NpcFactions",
		"Item.Merchantlists.NpcType.NpcFactions.NpcFactionEntries",
		"Item.Merchantlists.NpcType.NpcSpells",
		"Item.Merchantlists.NpcType.NpcSpells.NpcSpellsEntries",
		"Item.Merchantlists.NpcType.NpcTypesTint",
		"Item.Merchantlists.NpcType.Spawnentries",
		"Item.Merchantlists.NpcType.Spawnentries.NpcType",
		"Item.Merchantlists.NpcType.Spawnentries.Spawngroup",
		"Item.Merchantlists.NpcType.Spawnentries.Spawngroup.Spawn2",
		"Item.Merchantlists.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries",
		"Item.Merchantlists.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup",
		"Item.ObjectContents",
		"Item.Objects",
		"Item.Objects.Item",
		"Item.Objects.Zone",
		"Item.StartingItems",
		"Item.StartingItems.Item",
		"Item.StartingItems.Zone",
		"Item.Tasks",
		"Item.Tasks.TaskActivities",
		"Item.Tasks.TaskActivities.Goallists",
		"Item.Tasks.TaskActivities.NpcType",
		"Item.Tasks.TaskActivities.NpcType.AlternateCurrency",
		"Item.Tasks.TaskActivities.NpcType.Loottable",
		"Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries",
		"Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries",
		"Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries.Item",
		"Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop",
		"Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LootdropEntries",
		"Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LoottableEntries",
		"Item.Tasks.TaskActivities.NpcType.Loottable.LoottableEntries.Loottable",
		"Item.Tasks.TaskActivities.NpcType.Loottable.NpcTypes",
		"Item.Tasks.TaskActivities.NpcType.Merchantlists",
		"Item.Tasks.TaskActivities.NpcType.Merchantlists.NpcType",
		"Item.Tasks.TaskActivities.NpcType.NpcEmotes",
		"Item.Tasks.TaskActivities.NpcType.NpcFactions",
		"Item.Tasks.TaskActivities.NpcType.NpcFactions.NpcFactionEntries",
		"Item.Tasks.TaskActivities.NpcType.NpcSpells",
		"Item.Tasks.TaskActivities.NpcType.NpcSpells.NpcSpellsEntries",
		"Item.Tasks.TaskActivities.NpcType.NpcTypesTint",
		"Item.Tasks.TaskActivities.NpcType.Spawnentries",
		"Item.Tasks.TaskActivities.NpcType.Spawnentries.NpcType",
		"Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup",
		"Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2",
		"Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries",
		"Item.Tasks.TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup",
		"Item.Tasks.Tasksets",
		"Item.TradeskillRecipeEntries",
		"Item.TradeskillRecipeEntries.TradeskillRecipe",
		"Item.TributeLevels",
	}
}

func (Door) Connection() string {
    return "eqemu_content"
}
