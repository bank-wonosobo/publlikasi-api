package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type VisionMission struct {
	*gorm.Model
	ID       string `gorm:"primaryKey"`
	Title    string
	Vision   string
	Mission  string
	ImageUrl *string
}

// BeforeCreate hook to set UUID
func (v *VisionMission) BeforeCreate(tx *gorm.DB) (err error) {
	v.ID = uuid.NewString()
	return nil
}
