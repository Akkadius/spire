package models

type Expedition struct {
	ID                 uint  `json:"id" gorm:"Column:id"`
	DynamicZoneId      uint  `json:"dynamic_zone_id" gorm:"Column:dynamic_zone_id"`
	AddReplayOnJoin    uint8 `json:"add_replay_on_join" gorm:"Column:add_replay_on_join"`
	IsLocked           uint8 `json:"is_locked" gorm:"Column:is_locked"`
}

func (Expedition) TableName() string {
    return "expeditions"
}

func (Expedition) Relationships() []string {
    return []string{}
}

func (Expedition) Connection() string {
    return ""
}
