package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Banner struct {
	*gorm.Model
	ID          string `gorm:"primaryKey"`
	Name        string
	ImageUrl    string
	Description string
	IsActive    bool
}

// BeforeCreate hook to set UUID
func (b *Banner) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.NewString()
	return nil
}
