package models

type DynamicZone struct {
	ID                  uint    `json:"id" gorm:"Column:id"`
	InstanceId          int     `json:"instance_id" gorm:"Column:instance_id"`
	Type                uint8   `json:"type" gorm:"Column:type"`
	Uuid                string  `json:"uuid" gorm:"Column:uuid"`
	Name                string  `json:"name" gorm:"Column:name"`
	LeaderId            uint    `json:"leader_id" gorm:"Column:leader_id"`
	MinPlayers          uint    `json:"min_players" gorm:"Column:min_players"`
	MaxPlayers          uint    `json:"max_players" gorm:"Column:max_players"`
	DzSwitchId          int     `json:"dz_switch_id" gorm:"Column:dz_switch_id"`
	CompassZoneId       uint    `json:"compass_zone_id" gorm:"Column:compass_zone_id"`
	CompassX            float32 `json:"compass_x" gorm:"Column:compass_x"`
	CompassY            float32 `json:"compass_y" gorm:"Column:compass_y"`
	CompassZ            float32 `json:"compass_z" gorm:"Column:compass_z"`
	SafeReturnZoneId    uint    `json:"safe_return_zone_id" gorm:"Column:safe_return_zone_id"`
	SafeReturnX         float32 `json:"safe_return_x" gorm:"Column:safe_return_x"`
	SafeReturnY         float32 `json:"safe_return_y" gorm:"Column:safe_return_y"`
	SafeReturnZ         float32 `json:"safe_return_z" gorm:"Column:safe_return_z"`
	SafeReturnHeading   float32 `json:"safe_return_heading" gorm:"Column:safe_return_heading"`
	ZoneInX             float32 `json:"zone_in_x" gorm:"Column:zone_in_x"`
	ZoneInY             float32 `json:"zone_in_y" gorm:"Column:zone_in_y"`
	ZoneInZ             float32 `json:"zone_in_z" gorm:"Column:zone_in_z"`
	ZoneInHeading       float32 `json:"zone_in_heading" gorm:"Column:zone_in_heading"`
	HasZoneIn           uint8   `json:"has_zone_in" gorm:"Column:has_zone_in"`
	IsLocked            int8    `json:"is_locked" gorm:"Column:is_locked"`
	AddReplay           int8    `json:"add_replay" gorm:"Column:add_replay"`
}

func (DynamicZone) TableName() string {
    return "dynamic_zones"
}

func (DynamicZone) Relationships() []string {
    return []string{}
}

func (DynamicZone) Connection() string {
    return ""
}
