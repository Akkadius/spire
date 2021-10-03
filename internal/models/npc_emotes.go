package models

type NpcEmote struct {
	ID      int    `json:"id" gorm:"Column:id"`
	Emoteid uint   `json:"emoteid" gorm:"Column:emoteid"`
	Event2  int8   `json:"event_" gorm:"Column:event_"`
	Type    int8   `json:"type" gorm:"Column:type"`
	Text    string `json:"text" gorm:"Column:text"`
}

func (NpcEmote) TableName() string {
    return "npc_emotes"
}

func (NpcEmote) Relationships() []string {
    return []string{}
}

func (NpcEmote) Connection() string {
    return "eqemu_content"
}
