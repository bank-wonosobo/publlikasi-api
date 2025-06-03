package dto

import "time"

// request
type NewsCreateRequest struct {
	Title   string `json:"title" form:"title"`
	Slug    string `json:"slug" form:"slug"`
	Content string `json:"content" form:"content"`
	Image   string `json:"image" form:"image"`
}

type NewsGetQueryParams struct {
	Key    string `query:"key"`
	Status string `query:"status"`
	Page   int    `query:"page"`
	Limit  int    `query:"limit"`
}

type NewsUpdateRequest struct {
	Title   string  `json:"title" form:"title"`
	Slug    string  `json:"slug" form:"slug"`
	Content string  `json:"content" form:"content"`
	Image   *string `json:"image" form:"image"`
}

// response
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
