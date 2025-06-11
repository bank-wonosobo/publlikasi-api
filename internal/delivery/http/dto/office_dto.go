package dto

import "time"

// response
type OfficeResponse struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Address     string    `json:"address"`
	Latitude    string    `json:"latitude"`
	Longitude   string    `json:"longitude"`
	ImageUrl    string    `json:"image_url"`
	MapLink     string    `json:"map_link"`
	PhoneNumber string    `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
	UpdateAt    time.Time `json:"updated_at"`
}

// params
type OfficeGetQueryParams struct {
	Key   string `query:"key"`
	Page  int    `query:"page"`
	Limit int    `query:"limit"`
}

// request
type OfficeCreateReq struct {
	Name        string `json:"name" form:"name" validate:"required"`
	Address     string `json:"address" form:"address" validate:"required"`
	Latitude    string `json:"latitude" form:"latitude" validate:"required"`
	Longitude   string `json:"longitude" form:"longitude" validate:"required"`
	Image       string `json:"image" form:"image"`
	MapLink     string `json:"map_link" form:"map_link" validate:"required"`
	PhoneNumber string `json:"phone_number" form:"phone_number" validate:"required"`
}

type OfficeUpdateReq struct {
	Name        string  `json:"name" form:"name" validate:"required"`
	Address     string  `json:"address" form:"address" validate:"required"`
	Latitude    string  `json:"latitude" form:"latitude" validate:"required"`
	Longitude   string  `json:"longitude" form:"longitude" validate:"required"`
	Image       *string `json:"image" form:"image"`
	MapLink     string  `json:"map_link" form:"map_link" validate:"required"`
	PhoneNumber string  `json:"phone_number" form:"phone_number" validate:"required"`
}
