package dto

// response

type AuctionResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	ImageUrl    string `json:"image_url"`
	Description string `json:"description"`
	Link        string `json:"link"`
}

type AuctionGetQueryParams struct {
	Key   string `query:"key"`
	Page  int    `query:"page"`
	Limit int    `query:"limit"`
}

type AuctionCreateRequest struct {
	Title       string `json:"title" form:"title" validate:"required"`
	Image       string `json:"image" form:"image"`
	Description string `json:"description" form:"description" validate:"required"`
	Link        string `json:"link" form:"link" validate:"required"`
}

type AuctionUpdateRequest struct {
	Title       string  `json:"title" form:"title" validate:"required"`
	Image       *string `json:"image" form:"image"`
	Description string  `json:"description" form:"description" validate:"required"`
	Link        string  `json:"link" form:"link" validate:"required"`
}
