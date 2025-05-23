package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Status string

const (
	Draft     Status = "draft"
	Published Status = "published"
	Archived  Status = "archived"
)

type Report struct {
	*gorm.Model
	ID           string `gorm:"primaryKey"`
	Title        string `gorm:"varchar(255)"`
	Description  string `gorm:"text"`
	PeriodStart  time.Time
	PeriodEnd    time.Time
	Year         int
	Quarter      *int
	FileUrl      *string
	Version      string
	Status       Status
	UploadBy     string
	ApprovedBy   *string
	ReportTypeID uint
	ReportType   ReportType
}

// BeforeCreate hook to set UUID
func (r *Report) BeforeCreate(tx *gorm.DB) (err error) {
	r.ID = uuid.NewString()
	return nil
}
