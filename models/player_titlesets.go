package models

type PlayerTitleset struct {
	ID        uint `json:"id" gorm:"Column:id"`
	CharId    uint `json:"char_id" gorm:"Column:char_id"`
	TitleSet  uint `json:"title_set" gorm:"Column:title_set"`
}

func (PlayerTitleset) TableName() string {
    return "player_titlesets"
}

func (PlayerTitleset) Relationships() []string {
    return []string{}
}

func (PlayerTitleset) Connection() string {
    return ""
}
