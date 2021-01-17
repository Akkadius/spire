package models

type NpcTypesTint struct {
	ID            uint   `json:"id" gorm:"Column:id"`
	TintSetName   string `json:"tint_set_name" gorm:"Column:tint_set_name"`
	Red1H         uint8  `json:"red_1_h" gorm:"Column:red1h"`
	Grn1H         uint8  `json:"grn_1_h" gorm:"Column:grn1h"`
	Blu1H         uint8  `json:"blu_1_h" gorm:"Column:blu1h"`
	Red2C         uint8  `json:"red_2_c" gorm:"Column:red2c"`
	Grn2C         uint8  `json:"grn_2_c" gorm:"Column:grn2c"`
	Blu2C         uint8  `json:"blu_2_c" gorm:"Column:blu2c"`
	Red3A         uint8  `json:"red_3_a" gorm:"Column:red3a"`
	Grn3A         uint8  `json:"grn_3_a" gorm:"Column:grn3a"`
	Blu3A         uint8  `json:"blu_3_a" gorm:"Column:blu3a"`
	Red4B         uint8  `json:"red_4_b" gorm:"Column:red4b"`
	Grn4B         uint8  `json:"grn_4_b" gorm:"Column:grn4b"`
	Blu4B         uint8  `json:"blu_4_b" gorm:"Column:blu4b"`
	Red5G         uint8  `json:"red_5_g" gorm:"Column:red5g"`
	Grn5G         uint8  `json:"grn_5_g" gorm:"Column:grn5g"`
	Blu5G         uint8  `json:"blu_5_g" gorm:"Column:blu5g"`
	Red6L         uint8  `json:"red_6_l" gorm:"Column:red6l"`
	Grn6L         uint8  `json:"grn_6_l" gorm:"Column:grn6l"`
	Blu6L         uint8  `json:"blu_6_l" gorm:"Column:blu6l"`
	Red7F         uint8  `json:"red_7_f" gorm:"Column:red7f"`
	Grn7F         uint8  `json:"grn_7_f" gorm:"Column:grn7f"`
	Blu7F         uint8  `json:"blu_7_f" gorm:"Column:blu7f"`
	Red8X         uint8  `json:"red_8_x" gorm:"Column:red8x"`
	Grn8X         uint8  `json:"grn_8_x" gorm:"Column:grn8x"`
	Blu8X         uint8  `json:"blu_8_x" gorm:"Column:blu8x"`
	Red9X         uint8  `json:"red_9_x" gorm:"Column:red9x"`
	Grn9X         uint8  `json:"grn_9_x" gorm:"Column:grn9x"`
	Blu9X         uint8  `json:"blu_9_x" gorm:"Column:blu9x"`
}

func (NpcTypesTint) TableName() string {
    return "npc_types_tint"
}

func (NpcTypesTint) Relationships() []string {
    return []string{}
}

func (NpcTypesTint) Connection() string {
    return "eqemu_content"
}
