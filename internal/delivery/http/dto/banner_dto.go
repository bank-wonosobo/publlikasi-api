package dto

import "time"

type BannerCreateRequest struct {
	Name        string `json:"name" form:"name"`
	Image       string `json:"image" form:"image"`
	Description string `json:"description" form:"description"`
}

type BannerGetQueryParams struct {
	Key   string `query:"key"`
	Page  int    `query:"page"`
	Limit int    `query:"limit"`
}

// response

type BannerResponse struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	ImageUrl    string    `json:"image_url"`
	Description string    `json:"description"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
