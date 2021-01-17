package models

type CharacterSkill struct {
	ID       uint   `json:"id" gorm:"Column:id"`
	SkillId  uint16 `json:"skill_id" gorm:"Column:skill_id"`
	Value    uint16 `json:"value" gorm:"Column:value"`
}

func (CharacterSkill) TableName() string {
    return "character_skills"
}

func (CharacterSkill) Relationships() []string {
    return []string{}
}

func (CharacterSkill) Connection() string {
    return ""
}
