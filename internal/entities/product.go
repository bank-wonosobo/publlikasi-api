package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductCategory string

const (
	Deposito ProductCategory = "deposito"
	Tabungan ProductCategory = "tabungan"
	Kredit   ProductCategory = "kredit"
	Digital  ProductCategory = "digital"
)

type Product struct {
	*gorm.Model
	ID              string `gorm:"primaryKey"`
	Name            string
	Description     string
	ImageUrl        string
	Tagline         string
	ProductCategory ProductCategory
}

// BeforeCreate hook to set UUID
func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = uuid.NewString()
	return nil
}
