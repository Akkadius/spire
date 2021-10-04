package models

import (
	"time"
)

type BugReport struct {
	ID                  uint      `json:"id" gorm:"Column:id"`
	Zone                string    `json:"zone" gorm:"Column:zone"`
	ClientVersionId     uint      `json:"client_version_id" gorm:"Column:client_version_id"`
	ClientVersionName   string    `json:"client_version_name" gorm:"Column:client_version_name"`
	AccountId           uint      `json:"account_id" gorm:"Column:account_id"`
	CharacterId         uint      `json:"character_id" gorm:"Column:character_id"`
	CharacterName       string    `json:"character_name" gorm:"Column:character_name"`
	ReporterSpoof       int8      `json:"reporter_spoof" gorm:"Column:reporter_spoof"`
	CategoryId          uint      `json:"category_id" gorm:"Column:category_id"`
	CategoryName        string    `json:"category_name" gorm:"Column:category_name"`
	ReporterName        string    `json:"reporter_name" gorm:"Column:reporter_name"`
	UiPath              string    `json:"ui_path" gorm:"Column:ui_path"`
	PosX                float32   `json:"pos_x" gorm:"Column:pos_x"`
	PosY                float32   `json:"pos_y" gorm:"Column:pos_y"`
	PosZ                float32   `json:"pos_z" gorm:"Column:pos_z"`
	Heading             uint      `json:"heading" gorm:"Column:heading"`
	TimePlayed          uint      `json:"time_played" gorm:"Column:time_played"`
	TargetId            uint      `json:"target_id" gorm:"Column:target_id"`
	TargetName          string    `json:"target_name" gorm:"Column:target_name"`
	OptionalInfoMask    uint      `json:"optional_info_mask" gorm:"Column:optional_info_mask"`
	CanDuplicate        int8      `json:"_can_duplicate" gorm:"Column:_can_duplicate"`
	CrashBug            int8      `json:"_crash_bug" gorm:"Column:_crash_bug"`
	TargetInfo          int8      `json:"_target_info" gorm:"Column:_target_info"`
	CharacterFlags      int8      `json:"_character_flags" gorm:"Column:_character_flags"`
	UnknownValue        int8      `json:"_unknown_value" gorm:"Column:_unknown_value"`
	BugReport           string    `json:"bug_report" gorm:"Column:bug_report"`
	SystemInfo          string    `json:"system_info" gorm:"Column:system_info"`
	ReportDatetime      time.Time `json:"report_datetime" gorm:"Column:report_datetime"`
	BugStatus           uint8     `json:"bug_status" gorm:"Column:bug_status"`
	LastReview          time.Time `json:"last_review" gorm:"Column:last_review"`
	LastReviewer        string    `json:"last_reviewer" gorm:"Column:last_reviewer"`
	ReviewerNotes       string    `json:"reviewer_notes" gorm:"Column:reviewer_notes"`
}

func (BugReport) TableName() string {
    return "bug_reports"
}

func (BugReport) Relationships() []string {
    return []string{}
}

func (BugReport) Connection() string {
    return ""
}
