package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type News struct {
	*gorm.Model
	ID          string `gorm:"primaryKey"`
	Title       string
	Slug        string
	Content     string
	Author      string
	ImageUrl    string
	PublishedAt *time.Time
	Status      Status
	ApprovedBy  *string
}

// BeforeCreate hook to set UUID
func (n *News) BeforeCreate(tx *gorm.DB) (err error) {
	n.ID = uuid.NewString()
	return nil
}
