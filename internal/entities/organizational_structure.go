package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrganizationalStructure struct {
	*gorm.Model
	ID          string `gorm:"primaryKey"`
	Title       string
	Description string
	ImageUrl    *string
}

// BeforeCreate hook to set UUID
func (o *OrganizationalStructure) BeforeCreate(tx *gorm.DB) (err error) {
	o.ID = uuid.NewString()
	return nil
}
