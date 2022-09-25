package models

type FactionAssociation struct {
	ID     int     `json:"id" gorm:"Column:id"`
	Id1    int     `json:"id_1" gorm:"Column:id_1"`
	Mod1   float32 `json:"mod_1" gorm:"Column:mod_1"`
	Id2    int     `json:"id_2" gorm:"Column:id_2"`
	Mod2   float32 `json:"mod_2" gorm:"Column:mod_2"`
	Id3    int     `json:"id_3" gorm:"Column:id_3"`
	Mod3   float32 `json:"mod_3" gorm:"Column:mod_3"`
	Id4    int     `json:"id_4" gorm:"Column:id_4"`
	Mod4   float32 `json:"mod_4" gorm:"Column:mod_4"`
	Id5    int     `json:"id_5" gorm:"Column:id_5"`
	Mod5   float32 `json:"mod_5" gorm:"Column:mod_5"`
	Id6    int     `json:"id_6" gorm:"Column:id_6"`
	Mod6   float32 `json:"mod_6" gorm:"Column:mod_6"`
	Id7    int     `json:"id_7" gorm:"Column:id_7"`
	Mod7   float32 `json:"mod_7" gorm:"Column:mod_7"`
	Id8    int     `json:"id_8" gorm:"Column:id_8"`
	Mod8   float32 `json:"mod_8" gorm:"Column:mod_8"`
	Id9    int     `json:"id_9" gorm:"Column:id_9"`
	Mod9   float32 `json:"mod_9" gorm:"Column:mod_9"`
	Id10   int     `json:"id_10" gorm:"Column:id_10"`
	Mod10  float32 `json:"mod_10" gorm:"Column:mod_10"`
}

func (FactionAssociation) TableName() string {
    return "faction_association"
}

func (FactionAssociation) Relationships() []string {
    return []string{}
}

func (FactionAssociation) Connection() string {
    return ""
}
