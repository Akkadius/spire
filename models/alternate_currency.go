package models

type AlternateCurrency struct {
	ID      int `json:"id" gorm:"Column:id"`
	ItemId  int `json:"item_id" gorm:"Column:item_id"`
}

func (AlternateCurrency) TableName() string {
    return "alternate_currency"
}

func (AlternateCurrency) Relationships() []string {
    return []string{}
}

func (AlternateCurrency) Connection() string {
    return "eqemu_content"
}
