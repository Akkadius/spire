package models

type DynamicZoneMember struct {
	ID              uint `json:"id" gorm:"Column:id"`
	DynamicZoneId   uint `json:"dynamic_zone_id" gorm:"Column:dynamic_zone_id"`
	CharacterId     uint `json:"character_id" gorm:"Column:character_id"`
}

func (DynamicZoneMember) TableName() string {
    return "dynamic_zone_members"
}

func (DynamicZoneMember) Relationships() []string {
    return []string{}
}

func (DynamicZoneMember) Connection() string {
    return ""
}
