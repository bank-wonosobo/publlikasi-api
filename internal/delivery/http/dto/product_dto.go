package dto

import (
	"time"

	"github.com/bank-wonosobo/publlikasi-api.git/internal/entities"
)

type ProductCreateReq struct {
	Name            string                   `json:"name" validate:"required" form:"name"`
	Description     string                   `json:"description" validate:"required" form:"description"`
	Image           string                   `json:"image" form:"image"`
	Tagline         string                   `json:"tagline" validate:"required" form:"tagline"`
	ProductCategory entities.ProductCategory `json:"product_category" validate:"required,must_category_product" form:"product_category"`
}

type ProductUpdateReq struct {
	Name            string                   `json:"name" validate:"required" form:"name"`
	Description     string                   `json:"description" validate:"required" form:"description"`
	Image           *string                  `json:"image" form:"image"`
	Tagline         string                   `json:"tagline" validate:"required" form:"tagline"`
	ProductCategory entities.ProductCategory `json:"product_category" validate:"required,must_category_product" form:"product_category"`
}
type ProductGetQueryParams struct {
	Key      string `query:"key"`
	Category string `query:"category"`
	Page     int    `query:"page"`
	Limit    int    `query:"limit"`
}

// response
type ProductResponse struct {
	ID              string    `json:"id"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	ImageUrl        string    `json:"image_url"`
	Tagline         string    `json:"tagline"`
	ProductCategory string    `json:"product_category"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
