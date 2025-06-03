package repositories

import (
	"context"

	"github.com/bank-wonosobo/publlikasi-api.git/internal/delivery/http/dto"
	"github.com/bank-wonosobo/publlikasi-api.git/internal/entities"
	"gorm.io/gorm"
)

type ProductRepository interface {
	GetAll(ctx context.Context, params *dto.ProductGetQueryParams, offsite int) ([]entities.Product, int64, error)
	FindByName(ctx context.Context, name string) (*entities.Product, error)
	Save(ctx context.Context, product *entities.Product) (*entities.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

// GetAll implements ProductRepository.
func (p *productRepository) GetAll(ctx context.Context, params *dto.ProductGetQueryParams, offsite int) (result []entities.Product, total int64, err error) {
	query := p.db.Model(&entities.Product{})

	if params.Category != "" {
		query = query.Where("category_product = ?", params.Category)
	}

	if params.Key != "" {
		query = query.Where("name ILIKE ?", "%"+params.Key+"%").Or("description ILIKE ?", "%"+params.Key+"%")
	}

	query.Count(&total)

	err = query.Limit(params.Limit).Offset(offsite).Find(&result).Error
	if err != nil {
		return nil, 0, err
	}

	return result, total, nil
}

// FindByName implements ProductRepository.
func (p *productRepository) FindByName(ctx context.Context, name string) (result *entities.Product, err error) {
	err = p.db.WithContext(ctx).Where("name = ?", name).First(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Save implements ProductRepository.
func (p *productRepository) Save(ctx context.Context, product *entities.Product) (*entities.Product, error) {
	err := p.db.WithContext(ctx).Create(&product).Error
	if err != nil {
		return nil, err
	}

	return product, nil
}

func NewProduct(db *gorm.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}
