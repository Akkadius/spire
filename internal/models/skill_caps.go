package models

type SkillCap struct {
	ID       uint   `json:"id" gorm:"Column:id"`
	SkillId  uint8  `json:"skill_id" gorm:"Column:skill_id"`
	ClassId  uint8  `json:"class_id" gorm:"Column:class_id"`
	Level    uint8  `json:"level" gorm:"Column:level"`
	Cap      uint32 `json:"cap" gorm:"Column:cap"`
	Class2   uint8  `json:"class_" gorm:"Column:class_"`
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
