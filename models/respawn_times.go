package models

type RespawnTime struct {
	ID          int   `json:"id" gorm:"Column:id"`
	Start       int   `json:"start" gorm:"Column:start"`
	Duration    int   `json:"duration" gorm:"Column:duration"`
	InstanceId  int16 `json:"instance_id" gorm:"Column:instance_id"`
}

func (RespawnTime) TableName() string {
    return "respawn_times"
}

func (RespawnTime) Relationships() []string {
    return []string{}
}

func (RespawnTime) Connection() string {
    return ""
}
