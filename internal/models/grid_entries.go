package models

type GridEntry struct {
	Gridid      int     `json:"gridid" gorm:"Column:gridid"`
	Zoneid      int     `json:"zoneid" gorm:"Column:zoneid"`
	Number      int     `json:"number" gorm:"Column:number"`
	X           float32 `json:"x" gorm:"Column:x"`
	Y           float32 `json:"y" gorm:"Column:y"`
	Z           float32 `json:"z" gorm:"Column:z"`
	Heading     float32 `json:"heading" gorm:"Column:heading"`
	Pause       int     `json:"pause" gorm:"Column:pause"`
	Centerpoint int8    `json:"centerpoint" gorm:"Column:centerpoint"`
}

func (GridEntry) TableName() string {
    return "grid_entries"
}

func (GridEntry) Relationships() []string {
    return []string{}
}

func (GridEntry) Connection() string {
    return "eqemu_content"
}
