package models

type AaAbility struct {
	ID                 uint   `json:"id" gorm:"Column:id"`
	Name               string `json:"name" gorm:"Column:name"`
	Category           int    `json:"category" gorm:"Column:category"`
	Classes            int    `json:"classes" gorm:"Column:classes"`
	Races              int    `json:"races" gorm:"Column:races"`
	DrakkinHeritage    int    `json:"drakkin_heritage" gorm:"Column:drakkin_heritage"`
	Deities            int    `json:"deities" gorm:"Column:deities"`
	Status             int    `json:"status" gorm:"Column:status"`
	Type               int    `json:"type" gorm:"Column:type"`
	Charges            int    `json:"charges" gorm:"Column:charges"`
	GrantOnly          int8   `json:"grant_only" gorm:"Column:grant_only"`
	FirstRankId        int    `json:"first_rank_id" gorm:"Column:first_rank_id"`
	Enabled            uint8  `json:"enabled" gorm:"Column:enabled"`
	ResetOnDeath       int8   `json:"reset_on_death" gorm:"Column:reset_on_death"`
	AutoGrantEnabled   int8   `json:"auto_grant_enabled" gorm:"Column:auto_grant_enabled"`
}

func (AaAbility) TableName() string {
    return "aa_ability"
}

func (AaAbility) Relationships() []string {
    return []string{}
}

func (AaAbility) Connection() string {
    return "eqemu_content"
}
