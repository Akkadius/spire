package models

type GuildMember struct {
	CharId         int    `json:"char_id" gorm:"Column:char_id"`
	GuildId        uint32 `json:"guild_id" gorm:"Column:guild_id"`
	Rank           uint8  `json:"rank" gorm:"Column:rank"`
	TributeEnable  uint8  `json:"tribute_enable" gorm:"Column:tribute_enable"`
	TotalTribute   uint   `json:"total_tribute" gorm:"Column:total_tribute"`
	LastTribute    uint   `json:"last_tribute" gorm:"Column:last_tribute"`
	Banker         uint8  `json:"banker" gorm:"Column:banker"`
	PublicNote     string `json:"public_note" gorm:"Column:public_note"`
	Alt            uint8  `json:"alt" gorm:"Column:alt"`
}

func (GuildMember) TableName() string {
    return "guild_members"
}

func (GuildMember) Relationships() []string {
    return []string{}
}

func (GuildMember) Connection() string {
    return ""
}
