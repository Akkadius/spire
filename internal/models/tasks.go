package models

type Task struct {
	ID                      uint           `json:"id" gorm:"Column:id"`
	Type                    int8           `json:"type" gorm:"Column:type"`
	Duration                uint           `json:"duration" gorm:"Column:duration"`
	DurationCode            int8           `json:"duration_code" gorm:"Column:duration_code"`
	Title                   string         `json:"title" gorm:"Column:title"`
	Description             string         `json:"description" gorm:"Column:description"`
	Reward                  string         `json:"reward" gorm:"Column:reward"`
	Rewardid                uint           `json:"rewardid" gorm:"Column:rewardid"`
	Cashreward              uint           `json:"cashreward" gorm:"Column:cashreward"`
	Xpreward                int            `json:"xpreward" gorm:"Column:xpreward"`
	Rewardmethod            uint8          `json:"rewardmethod" gorm:"Column:rewardmethod"`
	RewardRadiantCrystals   uint           `json:"reward_radiant_crystals" gorm:"Column:reward_radiant_crystals"`
	RewardEbonCrystals      uint           `json:"reward_ebon_crystals" gorm:"Column:reward_ebon_crystals"`
	Minlevel                uint8          `json:"minlevel" gorm:"Column:minlevel"`
	Maxlevel                uint8          `json:"maxlevel" gorm:"Column:maxlevel"`
	LevelSpread             uint           `json:"level_spread" gorm:"Column:level_spread"`
	MinPlayers              uint           `json:"min_players" gorm:"Column:min_players"`
	MaxPlayers              uint           `json:"max_players" gorm:"Column:max_players"`
	Repeatable              uint8          `json:"repeatable" gorm:"Column:repeatable"`
	FactionReward           int            `json:"faction_reward" gorm:"Column:faction_reward"`
	CompletionEmote         string         `json:"completion_emote" gorm:"Column:completion_emote"`
	ReplayTimerSeconds      uint           `json:"replay_timer_seconds" gorm:"Column:replay_timer_seconds"`
	RequestTimerSeconds     uint           `json:"request_timer_seconds" gorm:"Column:request_timer_seconds"`
	TaskActivities          []TaskActivity `json:"task_activities,omitempty" gorm:"foreignKey:taskid;references:id"`
	Tasksets                []Taskset      `json:"tasksets,omitempty" gorm:"foreignKey:taskid;references:id"`
}

func (Task) TableName() string {
    return "tasks"
}

func (Task) Relationships() []string {
    return []string{
		"TaskActivities",
		"TaskActivities.Goallists",
		"TaskActivities.NpcType",
		"TaskActivities.NpcType.AlternateCurrency",
		"TaskActivities.NpcType.Loottable",
		"TaskActivities.NpcType.Loottable.LoottableEntries",
		"TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries",
		"TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries.Item",
		"TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.AlternateCurrencies",
		"TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.CharacterCorpseItems",
		"TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.DiscoveredItems",
		"TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Doors",
		"TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Doors.Item",
		"TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Fishings",
		"TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Fishings.Item",
		"TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Fishings.NpcType",
		"TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Fishings.Zone",
		"TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Forages",
		"TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Forages.Item",
		"TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Forages.Zone",
		"TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.GroundSpawns",
		"TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.GroundSpawns.Zone",
		"TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.ItemTicks",
		"TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Keyrings",
		"TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.LootdropEntries",
		"TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Merchantlists",
		"TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Merchantlists.NpcType",
		"TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.ObjectContents",
		"TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Objects",
		"TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Objects.Item",
		"TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Objects.Zone",
		"TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.StartingItems",
		"TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.StartingItems.Item",
		"TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.StartingItems.Zone",
		"TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.Tasks",
		"TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.TradeskillRecipeEntries",
		"TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.TradeskillRecipeEntries.TradeskillRecipe",
		"TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries.Item.TributeLevels",
		"TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop",
		"TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LootdropEntries",
		"TaskActivities.NpcType.Loottable.LoottableEntries.LootdropEntries.Lootdrop.LoottableEntries",
		"TaskActivities.NpcType.Loottable.LoottableEntries.Loottable",
		"TaskActivities.NpcType.Loottable.NpcTypes",
		"TaskActivities.NpcType.Merchantlists",
		"TaskActivities.NpcType.Merchantlists.NpcType",
		"TaskActivities.NpcType.NpcEmotes",
		"TaskActivities.NpcType.NpcFactions",
		"TaskActivities.NpcType.NpcFactions.NpcFactionEntries",
		"TaskActivities.NpcType.NpcSpells",
		"TaskActivities.NpcType.NpcSpells.NpcSpellsEntries",
		"TaskActivities.NpcType.NpcTypesTint",
		"TaskActivities.NpcType.Spawnentries",
		"TaskActivities.NpcType.Spawnentries.NpcType",
		"TaskActivities.NpcType.Spawnentries.Spawngroup",
		"TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2",
		"TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2.Spawnentries",
		"TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2.Spawngroup",
		"Tasksets",
	}
}

func (Task) Connection() string {
    return "eqemu_content"
}
