package models

type Grid struct {
	ID          int         `json:"id" gorm:"Column:id"`
	Zoneid      int         `json:"zoneid" gorm:"Column:zoneid"`
	Type        int         `json:"type" gorm:"Column:type"`
	Type2       int         `json:"type_2" gorm:"Column:type2"`
	GridEntries []GridEntry `json:"grid_entries,omitempty" gorm:"foreignKey:gridid;references:id"`
}

func (Grid) TableName() string {
    return "grid"
}

func (Grid) Relationships() []string {
    return []string{
		"GridEntries",
	}
}

func (Grid) Connection() string {
    return "eqemu_content"
}
