package models

type Title struct {
	ID              uint   `json:"id" gorm:"Column:id"`
	SkillId         int8   `json:"skill_id" gorm:"Column:skill_id"`
	MinSkillValue   int32  `json:"min_skill_value" gorm:"Column:min_skill_value"`
	MaxSkillValue   int32  `json:"max_skill_value" gorm:"Column:max_skill_value"`
	MinAaPoints     int32  `json:"min_aa_points" gorm:"Column:min_aa_points"`
	MaxAaPoints     int32  `json:"max_aa_points" gorm:"Column:max_aa_points"`
	Class           int8   `json:"class" gorm:"Column:class"`
	Gender          int8   `json:"gender" gorm:"Column:gender"`
	CharId          int    `json:"char_id" gorm:"Column:char_id"`
	Status          int    `json:"status" gorm:"Column:status"`
	ItemId          int    `json:"item_id" gorm:"Column:item_id"`
	Prefix          string `json:"prefix" gorm:"Column:prefix"`
	Suffix          string `json:"suffix" gorm:"Column:suffix"`
	TitleSet        int    `json:"title_set" gorm:"Column:title_set"`
}

func (Title) TableName() string {
    return "titles"
}

func (Title) Relationships() []string {
    return []string{}
}

func (Title) Connection() string {
    return ""
}
