package response

import "time"

type NewsResponse struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Slug        string     `json:"slug"`
	Content     string     `json:"content"`
	Author      string     `json:"author"`
	ImageUrl    string     `json:"image_url"`
	PublishedAt *time.Time `json:"published_at"`
	Status      string     `json:"status"`
	ApprovedBy  *string    `json:"approved_by"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}
