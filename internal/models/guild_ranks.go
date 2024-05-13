package models

type GuildRank struct {
	GuildId  uint32 `json:"guild_id" gorm:"Column:guild_id"`
	Rank     uint8  `json:"rank" gorm:"Column:rank"`
	Title    string `json:"title" gorm:"Column:title"`
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
