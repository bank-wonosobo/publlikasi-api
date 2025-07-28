package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Auction struct {
	*gorm.Model
	ID          string `gorm:"primaryKey"`
	Title       string
	Description string
	ExternalUrl string
	ImageUrl    string
	StartTime   time.Time
	EndTime     time.Time
}

// BeforeCreate hook to set UUID
func (b *Auction) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.NewString()
	return nil
}
