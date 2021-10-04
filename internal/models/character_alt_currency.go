package models

type CharacterAltCurrency struct {
	CharId      uint `json:"char_id" gorm:"Column:char_id"`
	CurrencyId  uint `json:"currency_id" gorm:"Column:currency_id"`
	Amount      uint `json:"amount" gorm:"Column:amount"`
}

func (CharacterAltCurrency) TableName() string {
    return "character_alt_currency"
}

func (CharacterAltCurrency) Relationships() []string {
    return []string{}
}

func (CharacterAltCurrency) Connection() string {
    return ""
}
