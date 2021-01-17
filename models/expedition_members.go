package models

type ExpeditionMember struct {
	ID                uint  `json:"id" gorm:"Column:id"`
	ExpeditionId      uint  `json:"expedition_id" gorm:"Column:expedition_id"`
	CharacterId       uint  `json:"character_id" gorm:"Column:character_id"`
	IsCurrentMember   uint8 `json:"is_current_member" gorm:"Column:is_current_member"`
}

func (ExpeditionMember) TableName() string {
    return "expedition_members"
}

func (ExpeditionMember) Relationships() []string {
    return []string{}
}

func (ExpeditionMember) Connection() string {
    return ""
}
