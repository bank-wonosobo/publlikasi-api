package response

import "time"

type NewsResponse struct {
	ID          string `gorm:"primaryKey"`
	Title       string
	Slug        string
	Content     string
	Author      string
	ImageUrl    string
	PublishedAt *time.Time
	Status      string
	ApprovedBy  *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
