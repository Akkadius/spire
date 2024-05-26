package models

type GuildPermission struct {
	ID         int `json:"id" gorm:"Column:id"`
	PermId     int `json:"perm_id" gorm:"Column:perm_id"`
	GuildId    int `json:"guild_id" gorm:"Column:guild_id"`
	Permission int `json:"permission" gorm:"Column:permission"`
}

func (GuildPermission) TableName() string {
    return "guild_permissions"
}

func (GuildPermission) Relationships() []string {
    return []string{}
}

func (GuildPermission) Connection() string {
    return ""
}
