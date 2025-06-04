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
	FindByID(ctx context.Context, id string) (*entities.Product, error)
	Delete(ctx context.Context, product *entities.Product) error
}

type productRepository struct {
	db *gorm.DB
}

// Delete implements ProductRepository.
func (p *productRepository) Delete(ctx context.Context, product *entities.Product) error {
	err := p.db.WithContext(ctx).Delete(&product).Error
	if err != nil {
		return err
	}

	return nil
}

// FindByID implements ProductRepository.
func (p *productRepository) FindByID(ctx context.Context, id string) (result *entities.Product, err error) {
	err = p.db.WithContext(ctx).Where("id = ?", id).First(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetAll implements ProductRepository.
func (p *productRepository) GetAll(ctx context.Context, params *dto.ProductGetQueryParams, offsite int) (result []entities.Product, total int64, err error) {
	query := p.db.Model(&entities.Product{})

	if params.Category != "" {
		query = query.Where("product_category = ?", params.Category)
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
	err := p.db.WithContext(ctx).Save(&product).Error
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
