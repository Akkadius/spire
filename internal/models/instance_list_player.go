package models

type InstanceListPlayer struct {
	ID     uint `json:"id" gorm:"Column:id"`
	Charid uint `json:"charid" gorm:"Column:charid"`
}

func (InstanceListPlayer) TableName() string {
    return "instance_list_player"
}

func (InstanceListPlayer) Relationships() []string {
    return []string{}
}

func (InstanceListPlayer) Connection() string {
    return ""
}
