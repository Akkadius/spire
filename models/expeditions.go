package models

type Expedition struct {
	ID                 uint   `json:"id" gorm:"Column:id"`
	Uuid               string `json:"uuid" gorm:"Column:uuid"`
	DynamicZoneId      uint   `json:"dynamic_zone_id" gorm:"Column:dynamic_zone_id"`
	ExpeditionName     string `json:"expedition_name" gorm:"Column:expedition_name"`
	LeaderId           uint   `json:"leader_id" gorm:"Column:leader_id"`
	MinPlayers         uint8  `json:"min_players" gorm:"Column:min_players"`
	MaxPlayers         uint8  `json:"max_players" gorm:"Column:max_players"`
	AddReplayOnJoin    uint8  `json:"add_replay_on_join" gorm:"Column:add_replay_on_join"`
	IsLocked           uint8  `json:"is_locked" gorm:"Column:is_locked"`
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
