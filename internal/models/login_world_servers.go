package models

import (
	"github.com/volatiletech/null/v8"
)

type LoginWorldServer struct {
	ID                        uint        `json:"id" gorm:"Column:id"`
	LongName                  string      `json:"long_name" gorm:"Column:long_name"`
	ShortName                 string      `json:"short_name" gorm:"Column:short_name"`
	TagDescription            string      `json:"tag_description" gorm:"Column:tag_description"`
	LoginServerListTypeId     int         `json:"login_server_list_type_id" gorm:"Column:login_server_list_type_id"`
	LastLoginDate             null.Time   `json:"last_login_date" gorm:"Column:last_login_date"`
	LastIpAddress             null.String `json:"last_ip_address" gorm:"Column:last_ip_address"`
	LoginServerAdminId        int         `json:"login_server_admin_id" gorm:"Column:login_server_admin_id"`
	IsServerTrusted           int         `json:"is_server_trusted" gorm:"Column:is_server_trusted"`
	Note                      null.String `json:"note" gorm:"Column:note"`
}

func (LoginWorldServer) TableName() string {
    return "login_world_servers"
}

func (LoginWorldServer) Relationships() []string {
    return []string{}
}

func (LoginWorldServer) Connection() string {
    return ""
}
