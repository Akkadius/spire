package models

import "time"

type CrashReport struct {
	ID              uint      `json:"id" gorm:"primary_key,AUTO_INCREMENT"`
	PlatformName    string    `json:"platform_name" gorm:"platform_name;type:varchar(20)"`
	OriginationInfo string    `json:"origination_info" gorm:"origination_info;type:varchar(150)"`
	CompileDate     string    `json:"compile_date" gorm:"compile_date;type:varchar(20)"`
	CompileTime     string    `json:"compile_time" gorm:"compile_time;type:varchar(20)"`
	Cpus            int       `json:"cpus" gorm:"cpus"`
	CrashReport     string    `json:"crash_report" gorm:"crash_report"`
	OsMachine       string    `json:"os_machine" gorm:"os_machine;type:varchar(200)"`
	OsRelease       string    `json:"os_release" gorm:"os_release;type:varchar(200)"`
	OsSysname       string    `json:"os_sysname" gorm:"os_sysname;type:varchar(200)"`
	OsVersion       string    `json:"os_version" gorm:"os_version;type:varchar(200)"`
	ProcessID       int       `json:"process_id" gorm:"process_id"`
	RssMemory       float64   `json:"rss_memory" gorm:"rss_memory"`
	ServerName      string    `json:"server_name" gorm:"server_name;type:varchar(200)"`
	ServerShortName string    `json:"server_short_name" gorm:"server_short_name;type:varchar(200)"`
	ServerVersion   string    `json:"server_version" gorm:"server_version;type:varchar(50);index:version"`
	Fingerprint     string    `json:"fingerprint" gorm:"fingerprint;type:varchar(100);index:fingerprint"`
	Uptime          int       `json:"uptime" gorm:"uptime"`
	CreatedAt       time.Time `json:"created_at" gorm:"Column:created_at"`
}

func (CrashReport) TableName() string {
	return "spire_crash_reports"
}

func (CrashReport) Relationships() []string {
	return []string{}
}

func (CrashReport) Connection() string {
	return "spire"
}
