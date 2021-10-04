package models

type GuildRank struct {
	GuildId      uint32 `json:"guild_id" gorm:"Column:guild_id"`
	Rank         uint8  `json:"rank" gorm:"Column:rank"`
	Title        string `json:"title" gorm:"Column:title"`
	CanHear      uint8  `json:"can_hear" gorm:"Column:can_hear"`
	CanSpeak     uint8  `json:"can_speak" gorm:"Column:can_speak"`
	CanInvite    uint8  `json:"can_invite" gorm:"Column:can_invite"`
	CanRemove    uint8  `json:"can_remove" gorm:"Column:can_remove"`
	CanPromote   uint8  `json:"can_promote" gorm:"Column:can_promote"`
	CanDemote    uint8  `json:"can_demote" gorm:"Column:can_demote"`
	CanMotd      uint8  `json:"can_motd" gorm:"Column:can_motd"`
	CanWarpeace  uint8  `json:"can_warpeace" gorm:"Column:can_warpeace"`
}

func (GuildRank) TableName() string {
    return "guild_ranks"
}

func (GuildRank) Relationships() []string {
    return []string{}
}

func (GuildRank) Connection() string {
    return ""
}
