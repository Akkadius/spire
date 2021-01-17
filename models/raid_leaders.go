package models

type RaidLeader struct {
	Gid            uint   `json:"gid" gorm:"Column:gid"`
	Rid            uint   `json:"rid" gorm:"Column:rid"`
	Marknpc        string `json:"marknpc" gorm:"Column:marknpc"`
	Maintank       string `json:"maintank" gorm:"Column:maintank"`
	Assist         string `json:"assist" gorm:"Column:assist"`
	Puller         string `json:"puller" gorm:"Column:puller"`
	Leadershipaa   []byte `json:"leadershipaa" gorm:"Column:leadershipaa"`
	Mentoree       string `json:"mentoree" gorm:"Column:mentoree"`
	MentorPercent  int    `json:"mentor_percent" gorm:"Column:mentor_percent"`
}

func (RaidLeader) TableName() string {
    return "raid_leaders"
}

func (RaidLeader) Relationships() []string {
    return []string{}
}

func (RaidLeader) Connection() string {
    return ""
}
