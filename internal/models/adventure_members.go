package models

type AdventureMember struct {
	ID     uint `json:"id" gorm:"Column:id"`
	Charid uint `json:"charid" gorm:"Column:charid"`
}

func (AdventureMember) TableName() string {
    return "adventure_members"
}

func (AdventureMember) Relationships() []string {
    return []string{}
}

func (AdventureMember) Connection() string {
    return ""
}
