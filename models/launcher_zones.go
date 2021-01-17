package models

type LauncherZone struct {
	Launcher string `json:"launcher" gorm:"Column:launcher"`
	Zone     string `json:"zone" gorm:"Column:zone"`
	Port     int32  `json:"port" gorm:"Column:port"`
}

func (LauncherZone) TableName() string {
    return "launcher_zones"
}

func (LauncherZone) Relationships() []string {
    return []string{}
}

func (LauncherZone) Connection() string {
    return ""
}
