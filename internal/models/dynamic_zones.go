package models

type DynamicZone struct {
	ID                  uint    `json:"id" gorm:"Column:id"`
	InstanceId          int     `json:"instance_id" gorm:"Column:instance_id"`
	Type                uint8   `json:"type" gorm:"Column:type"`
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
