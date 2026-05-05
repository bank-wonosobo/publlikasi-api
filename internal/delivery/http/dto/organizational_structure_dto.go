package dto

import "time"

type OrganizationalStructureResponse struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ImageUrl    *string   `json:"image_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type OrganizationalStructureGetQueryParams struct {
	Key   string `query:"key"`
	Page  int    `query:"page"`
	Limit int    `query:"limit"`
}

type OrganizationalStructureCreateReq struct {
	Title       string `json:"title" form:"title"`
	Description string `json:"description" form:"description" validate:"required"`
	Image       string `json:"image" form:"image"`
}

type OrganizationalStructureUpdateReq struct {
	Title       string  `json:"title" form:"title"`
	Description string  `json:"description" form:"description" validate:"required"`
	Image       *string `json:"image" form:"image"`
}
