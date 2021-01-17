package models

type Lfguild struct {
	Type       uint8  `json:"type" gorm:"Column:type"`
	Name       string `json:"name" gorm:"Column:name"`
	Comment    string `json:"comment" gorm:"Column:comment"`
	Fromlevel  uint8  `json:"fromlevel" gorm:"Column:fromlevel"`
	Tolevel    uint8  `json:"tolevel" gorm:"Column:tolevel"`
	Classes    uint   `json:"classes" gorm:"Column:classes"`
	Aacount    uint   `json:"aacount" gorm:"Column:aacount"`
	Timezone   uint   `json:"timezone" gorm:"Column:timezone"`
	Timeposted uint   `json:"timeposted" gorm:"Column:timeposted"`
}

func (Lfguild) TableName() string {
    return "lfguild"
}

func (Lfguild) Relationships() []string {
    return []string{}
}

func (Lfguild) Connection() string {
    return ""
}
