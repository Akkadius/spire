package models

type Chatchannel struct {
	ID        int    `json:"id" gorm:"Column:id"`
	Name      string `json:"name" gorm:"Column:name"`
	Owner     string `json:"owner" gorm:"Column:owner"`
	Password  string `json:"password" gorm:"Column:password"`
	Minstatus int    `json:"minstatus" gorm:"Column:minstatus"`
}

func (Chatchannel) TableName() string {
    return "chatchannels"
}

func (Chatchannel) Relationships() []string {
    return []string{}
}

func (Chatchannel) Connection() string {
    return ""
}
