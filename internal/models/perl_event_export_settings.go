package models

import (
	"github.com/volatiletech/null/v8"
)

type PerlEventExportSetting struct {
	EventId           int         `json:"event_id" gorm:"Column:event_id"`
	EventDescription  null.String `json:"event_description" gorm:"Column:event_description"`
	ExportQglobals    null.Int16  `json:"export_qglobals" gorm:"Column:export_qglobals"`
	ExportMob         null.Int16  `json:"export_mob" gorm:"Column:export_mob"`
	ExportZone        null.Int16  `json:"export_zone" gorm:"Column:export_zone"`
	ExportItem        null.Int16  `json:"export_item" gorm:"Column:export_item"`
	ExportEvent       null.Int16  `json:"export_event" gorm:"Column:export_event"`
}

func (PerlEventExportSetting) TableName() string {
    return "perl_event_export_settings"
}

func (PerlEventExportSetting) Relationships() []string {
    return []string{}
}

func (PerlEventExportSetting) Connection() string {
    return ""
}
