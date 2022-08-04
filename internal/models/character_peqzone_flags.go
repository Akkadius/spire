package models

type CharacterPeqzoneFlag struct {
	ID      int `json:"id" gorm:"Column:id"`
	ZoneId  int `json:"zone_id" gorm:"Column:zone_id"`
}

func (CharacterPeqzoneFlag) TableName() string {
    return "character_peqzone_flags"
}

func (CharacterPeqzoneFlag) Relationships() []string {
    return []string{}
}

func (CharacterPeqzoneFlag) Connection() string {
    return ""
}
