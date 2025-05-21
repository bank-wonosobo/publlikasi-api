package entities

import "time"

type Status string

const (
	Draft     Status = "draft"
	Published Status = "published"
	Archived  Status = "archived"
)

type Report struct {
	ID           string `gorm:"primaryKey"`
	Title        string `gorm:"varchar(255)"`
	Description  string `gorm:"text"`
	PeriodStart  time.Time
	PeriodEnd    time.Time
	Year         int
	Quarter      *int
	FileUrl      string
	Version      string
	Status       Status
	UploadBy     string
	ApprovedBy   *string
	ReportTypeID uint
	ReportType   ReportType
}
