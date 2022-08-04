package models

type DynamicZoneTemplate struct {
	ID               uint    `json:"id" gorm:"Column:id"`
	ZoneId           int     `json:"zone_id" gorm:"Column:zone_id"`
	ZoneVersion      int     `json:"zone_version" gorm:"Column:zone_version"`
	Name             string  `json:"name" gorm:"Column:name"`
	MinPlayers       int     `json:"min_players" gorm:"Column:min_players"`
	MaxPlayers       int     `json:"max_players" gorm:"Column:max_players"`
	DurationSeconds  int     `json:"duration_seconds" gorm:"Column:duration_seconds"`
	DzSwitchId       int     `json:"dz_switch_id" gorm:"Column:dz_switch_id"`
	CompassZoneId    int     `json:"compass_zone_id" gorm:"Column:compass_zone_id"`
	CompassX         float32 `json:"compass_x" gorm:"Column:compass_x"`
	CompassY         float32 `json:"compass_y" gorm:"Column:compass_y"`
	CompassZ         float32 `json:"compass_z" gorm:"Column:compass_z"`
	ReturnZoneId     int     `json:"return_zone_id" gorm:"Column:return_zone_id"`
	ReturnX          float32 `json:"return_x" gorm:"Column:return_x"`
	ReturnY          float32 `json:"return_y" gorm:"Column:return_y"`
	ReturnZ          float32 `json:"return_z" gorm:"Column:return_z"`
	ReturnH          float32 `json:"return_h" gorm:"Column:return_h"`
	OverrideZoneIn   int8    `json:"override_zone_in" gorm:"Column:override_zone_in"`
	ZoneInX          float32 `json:"zone_in_x" gorm:"Column:zone_in_x"`
	ZoneInY          float32 `json:"zone_in_y" gorm:"Column:zone_in_y"`
	ZoneInZ          float32 `json:"zone_in_z" gorm:"Column:zone_in_z"`
	ZoneInH          float32 `json:"zone_in_h" gorm:"Column:zone_in_h"`
}

func (DynamicZoneTemplate) TableName() string {
    return "dynamic_zone_templates"
}

func (DynamicZoneTemplate) Relationships() []string {
    return []string{}
}

func (DynamicZoneTemplate) Connection() string {
    return ""
}
