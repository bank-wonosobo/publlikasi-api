package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Post struct {
	*gorm.Model
	ID         string `gorm:"primaryKey"`
	Title      string
	Content    string `gorm:"text"`
	Summary    *string
	ImageUrl   *string
	Attachment string
	Status     Status
	PublishAt  time.Time
	CreatedBy  string
	ApprovedBy *string
	ExpiredAt  *time.Time
	PostTypeID string
	PostType   PostType
}

// BeforeCreate hook to set UUID
func (p *Post) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = uuid.NewString()
	return nil
}
