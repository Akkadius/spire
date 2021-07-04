package models

type Spawngroup struct {
	ID            int     `json:"id" gorm:"Column:id"`
	Name          string  `json:"name" gorm:"Column:name"`
	SpawnLimit    int8    `json:"spawn_limit" gorm:"Column:spawn_limit"`
	Dist          float32 `json:"dist" gorm:"Column:dist"`
	MaxX          float32 `json:"max_x" gorm:"Column:max_x"`
	MinX          float32 `json:"min_x" gorm:"Column:min_x"`
	MaxY          float32 `json:"max_y" gorm:"Column:max_y"`
	MinY          float32 `json:"min_y" gorm:"Column:min_y"`
	Delay         int     `json:"delay" gorm:"Column:delay"`
	Mindelay      int     `json:"mindelay" gorm:"Column:mindelay"`
	Despawn       int8    `json:"despawn" gorm:"Column:despawn"`
	DespawnTimer  int     `json:"despawn_timer" gorm:"Column:despawn_timer"`
	WpSpawns      uint8   `json:"wp_spawns" gorm:"Column:wp_spawns"`
}

func (Spawngroup) TableName() string {
    return "spawngroup"
}

func (Spawngroup) Relationships() []string {
    return []string{}
}

func (Spawngroup) Connection() string {
    return "eqemu_content"
}
