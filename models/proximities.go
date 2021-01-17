package models

type Proximity struct {
	Zoneid    uint    `json:"zoneid" gorm:"Column:zoneid"`
	Exploreid uint    `json:"exploreid" gorm:"Column:exploreid"`
	Minx      float32 `json:"minx" gorm:"Column:minx"`
	Maxx      float32 `json:"maxx" gorm:"Column:maxx"`
	Miny      float32 `json:"miny" gorm:"Column:miny"`
	Maxy      float32 `json:"maxy" gorm:"Column:maxy"`
	Minz      float32 `json:"minz" gorm:"Column:minz"`
	Maxz      float32 `json:"maxz" gorm:"Column:maxz"`
}

func (Proximity) TableName() string {
    return "proximities"
}

func (Proximity) Relationships() []string {
    return []string{}
}

func (Proximity) Connection() string {
    return "eqemu_content"
}
