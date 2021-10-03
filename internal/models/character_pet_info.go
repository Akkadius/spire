package models

type CharacterPetInfo struct {
	CharId   int     `json:"char_id" gorm:"Column:char_id"`
	Pet      int     `json:"pet" gorm:"Column:pet"`
	Petname  string  `json:"petname" gorm:"Column:petname"`
	Petpower int     `json:"petpower" gorm:"Column:petpower"`
	SpellId  int     `json:"spell_id" gorm:"Column:spell_id"`
	Hp       int     `json:"hp" gorm:"Column:hp"`
	Mana     int     `json:"mana" gorm:"Column:mana"`
	Size     float32 `json:"size" gorm:"Column:size"`
	Taunting int8    `json:"taunting" gorm:"Column:taunting"`
}

func (CharacterPetInfo) TableName() string {
    return "character_pet_info"
}

func (CharacterPetInfo) Relationships() []string {
    return []string{}
}

func (CharacterPetInfo) Connection() string {
    return ""
}
