package models

type CharacterPetName struct {
	CharacterId  int    `json:"character_id" gorm:"Column:character_id"`
	Name         string `json:"name" gorm:"Column:name"`
}

func (CharacterPetName) TableName() string {
    return "character_pet_name"
}

func (CharacterPetName) Relationships() []string {
    return []string{}
}

func (CharacterPetName) Connection() string {
    return ""
}
