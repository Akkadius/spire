package models

type Launcher struct {
	Name     string `json:"name" gorm:"Column:name"`
	Dynamics uint8  `json:"dynamics" gorm:"Column:dynamics"`
}

func (Launcher) TableName() string {
    return "launcher"
}

func (Launcher) Relationships() []string {
    return []string{}
}

func (Launcher) Connection() string {
    return ""
}
