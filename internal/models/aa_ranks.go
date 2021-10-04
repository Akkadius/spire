package models

type AaRank struct {
	ID               uint `json:"id" gorm:"Column:id"`
	UpperHotkeySid   int  `json:"upper_hotkey_sid" gorm:"Column:upper_hotkey_sid"`
	LowerHotkeySid   int  `json:"lower_hotkey_sid" gorm:"Column:lower_hotkey_sid"`
	TitleSid         int  `json:"title_sid" gorm:"Column:title_sid"`
	DescSid          int  `json:"desc_sid" gorm:"Column:desc_sid"`
	Cost             int  `json:"cost" gorm:"Column:cost"`
	LevelReq         int  `json:"level_req" gorm:"Column:level_req"`
	Spell            int  `json:"spell" gorm:"Column:spell"`
	SpellType        int  `json:"spell_type" gorm:"Column:spell_type"`
	RecastTime       int  `json:"recast_time" gorm:"Column:recast_time"`
	Expansion        int  `json:"expansion" gorm:"Column:expansion"`
	PrevId           int  `json:"prev_id" gorm:"Column:prev_id"`
	NextId           int  `json:"next_id" gorm:"Column:next_id"`
}

func (AaRank) TableName() string {
    return "aa_ranks"
}

func (AaRank) Relationships() []string {
    return []string{}
}

func (AaRank) Connection() string {
    return "eqemu_content"
}
