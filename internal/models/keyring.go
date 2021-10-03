package models

type Keyring struct {
	CharId  int `json:"char_id" gorm:"Column:char_id"`
	ItemId  int `json:"item_id" gorm:"Column:item_id"`
}

func (Keyring) TableName() string {
    return "keyring"
}

func (Keyring) Relationships() []string {
    return []string{}
}

func (Keyring) Connection() string {
    return ""
}
