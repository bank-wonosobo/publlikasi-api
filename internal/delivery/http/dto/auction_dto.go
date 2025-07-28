package dto

import "time"

// response

type AuctionResponse struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	ImageUrl    string    `json:"image_url"`
	Description string    `json:"description"`
	ExternalUrl string    `json:"external_url"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type AuctionGetQueryParams struct {
	Key      string `query:"key"`
	IsActive *bool  `query:"is_active"`
	Page     int    `query:"page"`
	Limit    int    `query:"limit"`
}

type AuctionCreateRequest struct {
	Title       string    `json:"title" form:"title" validate:"required"`
	Image       string    `json:"image" form:"image"`
	Description string    `json:"description" form:"description" validate:"required"`
	ExternalUrl string    `json:"external_url" form:"external_url" validate:"required"`
	StartTime   time.Time `json:"start_time" form:"start_time" validate:"required"`
	EndTime     time.Time `json:"end_time" form:"end_time" validate:"required"`
}
