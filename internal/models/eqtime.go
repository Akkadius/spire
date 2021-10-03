package models

type Eqtime struct {
	Minute   int8 `json:"minute" gorm:"Column:minute"`
	Hour     int8 `json:"hour" gorm:"Column:hour"`
	Day      int8 `json:"day" gorm:"Column:day"`
	Month    int8 `json:"month" gorm:"Column:month"`
	Year     int  `json:"year" gorm:"Column:year"`
	Realtime int  `json:"realtime" gorm:"Column:realtime"`
}

func (Eqtime) TableName() string {
    return "eqtime"
}

func (Eqtime) Relationships() []string {
    return []string{}
}

func (Eqtime) Connection() string {
    return ""
}
