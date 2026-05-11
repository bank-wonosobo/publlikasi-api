package dto

import "time"

type VisionMissionResponse struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Vision    string    `json:"vision"`
	Mission   string    `json:"mission"`
	ImageUrl  *string   `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type VisionMissionGetQueryParams struct {
	Key   string `query:"key"`
	Page  int    `query:"page"`
	Limit int    `query:"limit"`
}

type VisionMissionCreateReq struct {
	Title   string `json:"title" form:"title"`
	Vision  string `json:"vision" form:"vision" validate:"required"`
	Mission string `json:"mission" form:"mission" validate:"required"`
	Image   string `json:"image" form:"image"`
}

type VisionMissionUpdateReq struct {
	Title   string  `json:"title" form:"title"`
	Vision  string  `json:"vision" form:"vision" validate:"required"`
	Mission string  `json:"mission" form:"mission" validate:"required"`
	Image   *string `json:"image" form:"image"`
}
