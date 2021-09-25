package models

type Task struct {
	ID               uint           `json:"id" gorm:"Column:id"`
	Type             int8           `json:"type" gorm:"Column:type"`
	Duration         uint           `json:"duration" gorm:"Column:duration"`
	DurationCode     int8           `json:"duration_code" gorm:"Column:duration_code"`
	Title            string         `json:"title" gorm:"Column:title"`
	Description      string         `json:"description" gorm:"Column:description"`
	Reward           string         `json:"reward" gorm:"Column:reward"`
	Rewardid         uint           `json:"rewardid" gorm:"Column:rewardid"`
	Cashreward       uint           `json:"cashreward" gorm:"Column:cashreward"`
	Xpreward         int            `json:"xpreward" gorm:"Column:xpreward"`
	Rewardmethod     uint8          `json:"rewardmethod" gorm:"Column:rewardmethod"`
	Minlevel         uint8          `json:"minlevel" gorm:"Column:minlevel"`
	Maxlevel         uint8          `json:"maxlevel" gorm:"Column:maxlevel"`
	Repeatable       uint8          `json:"repeatable" gorm:"Column:repeatable"`
	FactionReward    int            `json:"faction_reward" gorm:"Column:faction_reward"`
	CompletionEmote  string         `json:"completion_emote" gorm:"Column:completion_emote"`
	TaskActivities   []TaskActivity `json:"task_activities,omitempty" gorm:"foreignKey:taskid;references:id"`
	Tasksets         []Taskset      `json:"tasksets,omitempty" gorm:"foreignKey:taskid;references:id"`
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
		"TaskActivities.NpcType.NpcEmotes",
		"TaskActivities.NpcType.NpcFactions",
		"TaskActivities.NpcType.NpcFactions.NpcFactionEntries",
		"TaskActivities.NpcType.NpcSpells",
		"TaskActivities.NpcType.NpcSpells.NpcSpellsEntries",
		"TaskActivities.NpcType.NpcTypesTint",
		"TaskActivities.NpcType.Spawnentries",
		"TaskActivities.NpcType.Spawnentries.Spawngroup",
		"TaskActivities.NpcType.Spawnentries.Spawngroup.Spawn2",
		"Tasksets",
	}
}

func (Task) Connection() string {
    return "eqemu_content"
}
