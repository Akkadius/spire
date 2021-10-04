package models

type AdventureDetail struct {
	ID                uint   `json:"id" gorm:"Column:id"`
	AdventureId       uint16 `json:"adventure_id" gorm:"Column:adventure_id"`
	InstanceId        int    `json:"instance_id" gorm:"Column:instance_id"`
	Count             uint16 `json:"count" gorm:"Column:count"`
	AssassinateCount  uint16 `json:"assassinate_count" gorm:"Column:assassinate_count"`
	Status            uint8  `json:"status" gorm:"Column:status"`
	TimeCreated       uint   `json:"time_created" gorm:"Column:time_created"`
	TimeZoned         uint   `json:"time_zoned" gorm:"Column:time_zoned"`
	TimeCompleted     uint   `json:"time_completed" gorm:"Column:time_completed"`
}

func (AdventureDetail) TableName() string {
    return "adventure_details"
}

func (AdventureDetail) Relationships() []string {
    return []string{}
}

func (AdventureDetail) Connection() string {
    return ""
}
