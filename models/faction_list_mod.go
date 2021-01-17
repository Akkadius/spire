package models

type FactionListMod struct {
	ID         uint   `json:"id" gorm:"Column:id"`
	FactionId  uint   `json:"faction_id" gorm:"Column:faction_id"`
	Mod        int16  `json:"mod" gorm:"Column:mod"`
	ModName    string `json:"mod_name" gorm:"Column:mod_name"`
}

func (FactionListMod) TableName() string {
    return "faction_list_mod"
}

func (FactionListMod) Relationships() []string {
    return []string{}
}

func (FactionListMod) Connection() string {
    return "eqemu_content"
}
