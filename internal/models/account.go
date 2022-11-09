package models

import (
	"github.com/volatiletech/null/v8"
)

type Account struct {
	ID             int             `json:"id" gorm:"Column:id"`
	Name           string          `json:"name" gorm:"Column:name"`
	Charname       string          `json:"charname" gorm:"Column:charname"`
	Sharedplat     uint            `json:"sharedplat" gorm:"Column:sharedplat"`
	Password       string          `json:"password" gorm:"Column:password"`
	Status         int             `json:"status" gorm:"Column:status"`
	LsId           string     `json:"ls_id" gorm:"Column:ls_id"`
	LsaccountId    uint       `json:"lsaccount_id" gorm:"Column:lsaccount_id"`
	Gmspeed        uint8           `json:"gmspeed" gorm:"Column:gmspeed"`
	Invulnerable   int8       `json:"invulnerable" gorm:"Column:invulnerable"`
	Flymode        int8       `json:"flymode" gorm:"Column:flymode"`
	IgnoreTells    int8       `json:"ignore_tells" gorm:"Column:ignore_tells"`
	Revoked        uint8           `json:"revoked" gorm:"Column:revoked"`
	Karma          uint            `json:"karma" gorm:"Column:karma"`
	MiniloginIp    string          `json:"minilogin_ip" gorm:"Column:minilogin_ip"`
	Hideme         int8            `json:"hideme" gorm:"Column:hideme"`
	Rulesflag      uint8           `json:"rulesflag" gorm:"Column:rulesflag"`
	Suspendeduntil time.Time       `json:"suspendeduntil" gorm:"Column:suspendeduntil"`
	TimeCreation   uint            `json:"time_creation" gorm:"Column:time_creation"`
	BanReason      string     `json:"ban_reason" gorm:"Column:ban_reason"`
	SuspendReason  string     `json:"suspend_reason" gorm:"Column:suspend_reason"`
	CrcEqgame      string     `json:"crc_eqgame" gorm:"Column:crc_eqgame"`
	CrcSkillcaps   string     `json:"crc_skillcaps" gorm:"Column:crc_skillcaps"`
	CrcBasedata    string     `json:"crc_basedata" gorm:"Column:crc_basedata"`
	AccountFlags   []AccountFlag   `json:"account_flags,omitempty" gorm:"foreignKey:p_accid;references:id"`
	AccountRewards []AccountReward `json:"account_rewards,omitempty" gorm:"foreignKey:account_id;references:id"`
	Sharedbanks    []Sharedbank    `json:"sharedbanks,omitempty" gorm:"foreignKey:acctid;references:id"`
	BugReports     []BugReport     `json:"bug_reports,omitempty" gorm:"foreignKey:account_id;references:id"`
	AccountIps     []AccountIp     `json:"account_ips,omitempty" gorm:"foreignKey:accid;references:id"`
}

func (Account) TableName() string {
    return "account"
}

func (Account) Relationships() []string {
    return []string{
		"AccountFlags",
		"AccountIps",
		"AccountRewards",
		"BugReports",
		"Sharedbanks",
	}
}

func (Account) Connection() string {
    return ""
}
