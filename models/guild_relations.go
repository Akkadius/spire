package models

type GuildRelation struct {
	Guild1   uint32 `json:"guild_1" gorm:"Column:guild1"`
	Guild2   uint32 `json:"guild_2" gorm:"Column:guild2"`
	Relation int8   `json:"relation" gorm:"Column:relation"`
}

func (GuildRelation) TableName() string {
    return "guild_relations"
}

func (GuildRelation) Relationships() []string {
    return []string{}
}

func (GuildRelation) Connection() string {
    return ""
}
