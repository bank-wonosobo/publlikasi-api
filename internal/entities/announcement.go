package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TargetAudience string

const (
	All      TargetAudience = "all"
	Employee TargetAudience = "employee"
	Customer TargetAudience = "customer"
)

type Announcement struct {
	*gorm.Model
	ID             string `gorm:"primaryKey"`
	Title          string
	Content        string
	Author         string
	TargetAudience TargetAudience
	StartDate      string  // tanggal mulai berlaku
	EndDate        string  // tanngal selesai berlaku
	AttachmentUrl  *string // file attachment (pdf / image)
	IsActive       bool
	Status         Status
}

// BeforeCreate hook to set UUID
func (a *Announcement) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = uuid.NewString()
	return nil
}
