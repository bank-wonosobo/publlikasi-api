package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Office struct {
	*gorm.Model
	ID          string `gorm:"primaryKey"`
	Name        string
	Address     string
	Latitude    string
	Longitude   string
	ImageUrl    string
	MapLink     string
	PhoneNumber string
}

// BeforeCreate hook to set UUID
func (o *Office) BeforeCreate(tx *gorm.DB) (err error) {
	o.ID = uuid.NewString()
	return nil
}
