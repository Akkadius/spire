package models

type CharacterInstanceSafereturn struct {
	ID               uint    `json:"id" gorm:"Column:id"`
	CharacterId      uint    `json:"character_id" gorm:"Column:character_id"`
	InstanceZoneId   int     `json:"instance_zone_id" gorm:"Column:instance_zone_id"`
	InstanceId       int     `json:"instance_id" gorm:"Column:instance_id"`
	SafeZoneId       int     `json:"safe_zone_id" gorm:"Column:safe_zone_id"`
	SafeX            float32 `json:"safe_x" gorm:"Column:safe_x"`
	SafeY            float32 `json:"safe_y" gorm:"Column:safe_y"`
	SafeZ            float32 `json:"safe_z" gorm:"Column:safe_z"`
	SafeHeading      float32 `json:"safe_heading" gorm:"Column:safe_heading"`
}

func (CharacterInstanceSafereturn) TableName() string {
    return "character_instance_safereturns"
}

func (CharacterInstanceSafereturn) Relationships() []string {
    return []string{}
}

func (CharacterInstanceSafereturn) Connection() string {
    return ""
}
