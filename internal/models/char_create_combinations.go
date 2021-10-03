package models

type CharCreateCombination struct {
	AllocationId   uint `json:"allocation_id" gorm:"Column:allocation_id"`
	Race           uint `json:"race" gorm:"Column:race"`
	Class          uint `json:"class" gorm:"Column:class"`
	Deity          uint `json:"deity" gorm:"Column:deity"`
	StartZone      uint `json:"start_zone" gorm:"Column:start_zone"`
	ExpansionsReq  uint `json:"expansions_req" gorm:"Column:expansions_req"`
}

func (CharCreateCombination) TableName() string {
    return "char_create_combinations"
}

func (CharCreateCombination) Relationships() []string {
    return []string{}
}

func (CharCreateCombination) Connection() string {
    return "eqemu_content"
}
