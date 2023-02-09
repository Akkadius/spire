package models

import (
	"github.com/volatiletech/null/v8"
)

type BotInventory struct {
	InventoriesIndex    uint        `json:"inventories_index" gorm:"Column:inventories_index"`
	BotId               uint        `json:"bot_id" gorm:"Column:bot_id"`
	SlotId              uint32      `json:"slot_id" gorm:"Column:slot_id"`
	ItemId              null.Uint   `json:"item_id" gorm:"Column:item_id"`
	InstCharges         null.Uint8  `json:"inst_charges" gorm:"Column:inst_charges"`
	InstColor           uint        `json:"inst_color" gorm:"Column:inst_color"`
	InstNoDrop          uint8       `json:"inst_no_drop" gorm:"Column:inst_no_drop"`
	InstCustomData      null.String `json:"inst_custom_data" gorm:"Column:inst_custom_data"`
	OrnamentIcon        uint        `json:"ornament_icon" gorm:"Column:ornament_icon"`
	OrnamentIdFile      uint        `json:"ornament_id_file" gorm:"Column:ornament_id_file"`
	OrnamentHeroModel   int         `json:"ornament_hero_model" gorm:"Column:ornament_hero_model"`
	Augment1            uint32      `json:"augment_1" gorm:"Column:augment_1"`
	Augment2            uint32      `json:"augment_2" gorm:"Column:augment_2"`
	Augment3            uint32      `json:"augment_3" gorm:"Column:augment_3"`
	Augment4            uint32      `json:"augment_4" gorm:"Column:augment_4"`
	Augment5            uint32      `json:"augment_5" gorm:"Column:augment_5"`
	Augment6            uint32      `json:"augment_6" gorm:"Column:augment_6"`
}

func (BotInventory) TableName() string {
    return "bot_inventories"
}

func (BotInventory) Relationships() []string {
    return []string{}
}

func (BotInventory) Connection() string {
    return ""
}
