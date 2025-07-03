package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ComplaintStatus string

const (
	Pending ComplaintStatus = "pending"
	Process ComplaintStatus = "process"
	Done    ComplaintStatus = "done"
	Reject  ComplaintStatus = "reject"
)

type Complaint struct {
	ID               string `gorm:"primaryKey"`
	ComplaintID      string `gorm:"unique"`
	ReportedName     string
	InsidentLocation string
	InsidentTime     time.Time
	Description      string
	EvidenceUrl      string
	ReporterName     string
	ReporterEmail    string
	ReporterPhone    string
	Status           ComplaintStatus
	ComplaintTypeID  uint
	ComplaintType    ComplaintType
}

// BeforeCreate hook to set UUID
func (c *Complaint) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.NewString()
	return nil
}
