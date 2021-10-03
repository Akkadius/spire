package models

type RaidMember struct {
	Raidid        int    `json:"raidid" gorm:"Column:raidid"`
	Charid        int    `json:"charid" gorm:"Column:charid"`
	Groupid       uint   `json:"groupid" gorm:"Column:groupid"`
	Class         int8   `json:"_class" gorm:"Column:_class"`
	Level         int8   `json:"level" gorm:"Column:level"`
	Name          string `json:"name" gorm:"Column:name"`
	Isgroupleader int8   `json:"isgroupleader" gorm:"Column:isgroupleader"`
	Israidleader  int8   `json:"israidleader" gorm:"Column:israidleader"`
	Islooter      int8   `json:"islooter" gorm:"Column:islooter"`
}

func (RaidMember) TableName() string {
    return "raid_members"
}

func (RaidMember) Relationships() []string {
    return []string{}
}

func (RaidMember) Connection() string {
    return ""
}
