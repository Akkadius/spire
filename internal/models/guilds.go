package models

type Guild struct {
	ID           int           `json:"id" gorm:"Column:id"`
	Name         string        `json:"name" gorm:"Column:name"`
	Leader       int           `json:"leader" gorm:"Column:leader"`
	Minstatus    int16         `json:"minstatus" gorm:"Column:minstatus"`
	Motd         string        `json:"motd" gorm:"Column:motd"`
	Tribute      uint          `json:"tribute" gorm:"Column:tribute"`
	MotdSetter   string        `json:"motd_setter" gorm:"Column:motd_setter"`
	Channel      string        `json:"channel" gorm:"Column:channel"`
	Url          string        `json:"url" gorm:"Column:url"`
	Favor        uint          `json:"favor" gorm:"Column:favor"`
	GuildBanks   []GuildBank   `json:"guild_banks,omitempty" gorm:"foreignKey:guildid;references:id"`
	GuildRanks   []GuildRank   `json:"guild_ranks,omitempty" gorm:"foreignKey:guild_id;references:id"`
	GuildMembers []GuildMember `json:"guild_members,omitempty" gorm:"foreignKey:guild_id;references:id"`
}

func (Guild) TableName() string {
    return "guilds"
}

func (Guild) Relationships() []string {
    return []string{
		"GuildBanks",
		"GuildMembers",
		"GuildRanks",
	}
}

func (Guild) Connection() string {
    return ""
}
