package models

type LoginServerListType struct {
	ID          uint   `json:"id" gorm:"Column:id"`
	Description string `json:"description" gorm:"Column:description"`
}

func (LoginServerListType) TableName() string {
    return "login_server_list_types"
}

func (LoginServerListType) Relationships() []string {
    return []string{}
}

func (LoginServerListType) Connection() string {
    return ""
}
