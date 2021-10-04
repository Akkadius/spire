package models

type Tribute struct {
	ID      uint   `json:"id" gorm:"Column:id"`
	Unknown uint   `json:"unknown" gorm:"Column:unknown"`
	Name    string `json:"name" gorm:"Column:name"`
	Descr   string `json:"descr" gorm:"Column:descr"`
	Isguild int8   `json:"isguild" gorm:"Column:isguild"`
}

func (Tribute) TableName() string {
    return "tributes"
}

func (Tribute) Relationships() []string {
    return []string{}
}

func (Tribute) Connection() string {
    return "eqemu_content"
}
