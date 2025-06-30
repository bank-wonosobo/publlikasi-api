package entities

import "time"

type ComplaintStatus string

const (
	Pending TargetAudience = "pending"
	Process TargetAudience = "process"
	Done    TargetAudience = "done"
	Reject  TargetAudience = "reject"
)

type Complaint struct {
	ReportedName     string          `gorm:"reported_name"`
	Email            string          `gorm:"email"`
	InsidentLocation string          `gorm:"incident_location"`
	InsidentDate     time.Time       `gorm:"incident_date"`
	InsidentTime     time.Time       `gorm:"incident_date"`
	Description      string          `gorm:"description"`
	EvidenceUrl      string          `gorm:"evidence_url"`
	ReporterName     string          `gorm:"reporter_name"`
	ReporterPhone    string          `gorm:"reporter_phone"`
	Status           ComplaintStatus `gorm:"status"`
	ComplaintTypeID  uint
	ComplaintType    ComplaintType
}
