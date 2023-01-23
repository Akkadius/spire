package models

type CrashReport struct {
	ID              uint    `json:"id" gorm:"primary_key,AUTO_INCREMENT"`
	PlatformName    string  `json:"platform_name" gorm:"platform_name"`
	CompileDate     string  `json:"compile_date" gorm:"compile_date"`
	CompileTime     string  `json:"compile_time" gorm:"compile_time"`
	Cpus            int     `json:"cpus" gorm:"cpus"`
	CrashReport     string  `json:"crash_report" gorm:"crash_report"`
	OsMachine       string  `json:"os_machine" gorm:"os_machine"`
	OsRelease       string  `json:"os_release" gorm:"os_release"`
	OsSysname       string  `json:"os_sysname" gorm:"os_sysname"`
	OsVersion       string  `json:"os_version" gorm:"os_version"`
	ProcessID       int     `json:"process_id" gorm:"process_id"`
	RssMemory       float64 `json:"rss_memory" gorm:"rss_memory"`
	ServerName      string  `json:"server_name" gorm:"server_name"`
	ServerShortName string  `json:"server_short_name" gorm:"server_short_name"`
	ServerVersion   string  `json:"server_version" gorm:"server_version"`
	Uptime          int     `json:"uptime" gorm:"uptime"`
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