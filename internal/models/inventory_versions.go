package models

type InventoryVersion struct {
	Version  uint `json:"version" gorm:"Column:version"`
	Step     uint `json:"step" gorm:"Column:step"`
	BotStep  uint `json:"bot_step" gorm:"Column:bot_step"`
}

func (InventoryVersion) TableName() string {
    return "inventory_versions"
}

func (InventoryVersion) Relationships() []string {
    return []string{}
}

func (InventoryVersion) Connection() string {
    return ""
}
