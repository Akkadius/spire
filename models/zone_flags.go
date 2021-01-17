package models

type ZoneFlag struct {
	CharID int `json:"char_id" gorm:"Column:charID"`
	ZoneID int `json:"zone_id" gorm:"Column:zoneID"`
}

func (ZoneFlag) TableName() string {
    return "zone_flags"
}

func (ZoneFlag) Relationships() []string {
    return []string{}
}

func (ZoneFlag) Connection() string {
    return ""
}
