package models

type SkillCap struct {
	SkillID uint8  `json:"skill_id" gorm:"Column:skillID"`
	Class   uint8  `json:"class" gorm:"Column:class"`
	Level   uint8  `json:"level" gorm:"Column:level"`
	Cap     uint32 `json:"cap" gorm:"Column:cap"`
	Class2  uint8  `json:"class_" gorm:"Column:class_"`
}

func (SkillCap) TableName() string {
    return "skill_caps"
}

func (SkillCap) Relationships() []string {
    return []string{}
}

func (SkillCap) Connection() string {
    return "eqemu_content"
}
