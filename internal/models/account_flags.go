package models

type AccountFlag struct {
	PAccid  uint   `json:"p_accid" gorm:"Column:p_accid"`
	PFlag   string `json:"p_flag" gorm:"Column:p_flag"`
	PValue  string `json:"p_value" gorm:"Column:p_value"`
}

func (AccountFlag) TableName() string {
    return "account_flags"
}

func (AccountFlag) Relationships() []string {
    return []string{}
}

func (AccountFlag) Connection() string {
    return ""
}
