package models

type GuildTribute struct {
	GuildId           uint `json:"guild_id" gorm:"Column:guild_id"`
	TributeId1        uint `json:"tribute_id_1" gorm:"Column:tribute_id_1"`
	TributeId1Tier    uint `json:"tribute_id_1_tier" gorm:"Column:tribute_id_1_tier"`
	TributeId2        uint `json:"tribute_id_2" gorm:"Column:tribute_id_2"`
	TributeId2Tier    uint `json:"tribute_id_2_tier" gorm:"Column:tribute_id_2_tier"`
	TimeRemaining     uint `json:"time_remaining" gorm:"Column:time_remaining"`
	Enabled           uint `json:"enabled" gorm:"Column:enabled"`
}

func (GuildTribute) TableName() string {
    return "guild_tributes"
}

func (GuildTribute) Relationships() []string {
    return []string{}
}

func (GuildTribute) Connection() string {
    return ""
}
