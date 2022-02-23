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
		"TaskActivities.NpcType.Merchantlists",
		"TaskActivities.NpcType.Merchantlists.Item",
		"TaskActivities.NpcType.Merchantlists.Item.AlternateCurrencies",
		"TaskActivities.NpcType.Merchantlists.Item.CharacterCorpseItems",
		"TaskActivities.NpcType.Merchantlists.Item.DiscoveredItems",
		"TaskActivities.NpcType.Merchantlists.Item.Doors",
		"TaskActivities.NpcType.Merchantlists.Item.Doors.Item",
		"TaskActivities.NpcType.Merchantlists.Item.Doors.Zone",
		"TaskActivities.NpcType.Merchantlists.Item.Fishings",
		"TaskActivities.NpcType.Merchantlists.Item.Fishings.Item",
		"TaskActivities.NpcType.Merchantlists.Item.Fishings.NpcType",
		"TaskActivities.NpcType.Merchantlists.Item.Fishings.Zone",
		"TaskActivities.NpcType.Merchantlists.Item.Forages",
		"TaskActivities.NpcType.Merchantlists.Item.Forages.Item",
		"TaskActivities.NpcType.Merchantlists.Item.Forages.Zone",
		"TaskActivities.NpcType.Merchantlists.Item.GroundSpawns",
		"TaskActivities.NpcType.Merchantlists.Item.GroundSpawns.Item",
		"TaskActivities.NpcType.Merchantlists.Item.GroundSpawns.Zone",
		"TaskActivities.NpcType.Merchantlists.Item.ItemTicks",
		"TaskActivities.NpcType.Merchantlists.Item.Keyrings",
		"TaskActivities.NpcType.Merchantlists.Item.LootdropEntries",
		"TaskActivities.NpcType.Merchantlists.Item.LootdropEntries.Item",
		"TaskActivities.NpcType.Merchantlists.Item.LootdropEntries.Lootdrop",
		"TaskActivities.NpcType.Merchantlists.Item.LootdropEntries.Lootdrop.LootdropEntries",
		"TaskActivities.NpcType.Merchantlists.Item.LootdropEntries.Lootdrop.LoottableEntries",
		"TaskActivities.NpcType.Merchantlists.Item.LootdropEntries.Lootdrop.LoottableEntries.LootdropEntries",
		"TaskActivities.NpcType.Merchantlists.Item.Merchantlists",
		"TaskActivities.NpcType.Merchantlists.Item.ObjectContents",
		"TaskActivities.NpcType.Merchantlists.Item.Objects",
		"TaskActivities.NpcType.Merchantlists.Item.Objects.Item",
		"TaskActivities.NpcType.Merchantlists.Item.Objects.Zone",
		"TaskActivities.NpcType.Merchantlists.Item.StartingItems",
		"TaskActivities.NpcType.Merchantlists.Item.StartingItems.Item",
		"TaskActivities.NpcType.Merchantlists.Item.StartingItems.Zone",
		"TaskActivities.NpcType.Merchantlists.Item.TradeskillRecipeEntries",
		"TaskActivities.NpcType.Merchantlists.Item.TradeskillRecipeEntries.TradeskillRecipe",
		"TaskActivities.NpcType.Merchantlists.Item.TributeLevels",
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
