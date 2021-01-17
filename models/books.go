package models

type Book struct {
	Name     string `json:"name" gorm:"Column:name"`
	Txtfile  string `json:"txtfile" gorm:"Column:txtfile"`
	Language int    `json:"language" gorm:"Column:language"`
}

func (Book) TableName() string {
    return "books"
}

func (Book) Relationships() []string {
    return []string{}
}

func (Book) Connection() string {
    return "eqemu_content"
}
