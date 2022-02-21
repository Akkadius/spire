package models

type SharedTaskDynamicZone struct {
	SharedTaskId    int64 `json:"shared_task_id" gorm:"Column:shared_task_id"`
	DynamicZoneId   uint  `json:"dynamic_zone_id" gorm:"Column:dynamic_zone_id"`
}

func (SharedTaskDynamicZone) TableName() string {
    return "shared_task_dynamic_zones"
}

func (SharedTaskDynamicZone) Relationships() []string {
    return []string{}
}

func (SharedTaskDynamicZone) Connection() string {
    return ""
}
