package repositories

import (
	"context"

	"github.com/bank-wonosobo/publlikasi-api.git/internal/delivery/http/dto"
	"github.com/bank-wonosobo/publlikasi-api.git/internal/entities"
	"gorm.io/gorm"
)

type BannerRepository interface {
	GetAll(ctx context.Context, params *dto.BannerGetQueryParams, offsite int) ([]entities.Banner, int64, error)
	FindByName(ctx context.Context, name string) (*entities.Banner, error)
	Save(ctx context.Context, banner *entities.Banner) (*entities.Banner, error)
	FindByID(ctx context.Context, id string) (*entities.Banner, error)
	Delete(ctx context.Context, banner *entities.Banner) error
}

type bannerRepository struct {
	db *gorm.DB
}

// Delete implements BannerRepository.
func (b *bannerRepository) Delete(ctx context.Context, banner *entities.Banner) error {
	panic("unimplemented")
}

// FindByID implements BannerRepository.
func (b *bannerRepository) FindByID(ctx context.Context, id string) (*entities.Banner, error) {
	panic("unimplemented")
}

// FindByName implements BannerRepository.
func (b *bannerRepository) FindByName(ctx context.Context, name string) (result *entities.Banner, err error) {
	err = b.db.WithContext(ctx).Where("name = ?", name).First(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetAll implements BannerRepository.
func (b *bannerRepository) GetAll(ctx context.Context, params *dto.BannerGetQueryParams, offsite int) ([]entities.Banner, int64, error) {
	panic("unimplemented")
}

// Save implements BannerRepository.
func (b *bannerRepository) Save(ctx context.Context, banner *entities.Banner) (*entities.Banner, error) {
	err := b.db.WithContext(ctx).Create(&banner).Error
	if err != nil {
		return nil, err
	}

	return banner, nil
}

func NewBanner(db *gorm.DB) BannerRepository {
	return &bannerRepository{
		db: db,
	}
}
