package repositories

import (
	"context"

	"github.com/bank-wonosobo/publlikasi-api.git/internal/delivery/http/dto"
	"github.com/bank-wonosobo/publlikasi-api.git/internal/entities"
	"gorm.io/gorm"
)

type NewsRepository interface {
	Save(ctx context.Context, tx *gorm.DB, news *entities.News) (*entities.News, error)
	GetAll(ctx context.Context, tx *gorm.DB, params *dto.NewsGetQueryParams, offsite int) ([]entities.News, int64, error)
	Update(ctx context.Context, tx *gorm.DB, news *entities.News) (*entities.News, error)
	Delete(ctx context.Context, tx *gorm.DB, news *entities.News) error
	FindByID(ctx context.Context, tx *gorm.DB, id string) (*entities.News, error)
	FindBySlug(ctx context.Context, tx *gorm.DB, slug string) (*entities.News, error)
	FindByTitle(ctx context.Context, tx *gorm.DB, title string) (*entities.News, error)
}

type newsRepository struct {
}

func NewNews() NewsRepository {
	return &newsRepository{}
}

// Create implements NewsRepository.
func (n *newsRepository) Save(ctx context.Context, tx *gorm.DB, news *entities.News) (*entities.News, error) {
	err := tx.WithContext(ctx).Create(&news).Error
	if err != nil {
		return nil, err
	}

	return news, nil
}

// GetAll implements NewsRepository.
func (n *newsRepository) GetAll(ctx context.Context, tx *gorm.DB, params *dto.NewsGetQueryParams, offsite int) (result []entities.News, total int64, err error) {
	query := tx.Model(&entities.News{})

	if params.Status != "" {
		query = query.Where("status = ?", params.Status)
	}

	if params.Key != "" {
		query = query.Where("title ILIKE ?", "%"+params.Key+"%").Or("content ILIKE ?", "%"+params.Key+"%")
	}

	query.Count(&total)

	err = query.Limit(params.Limit).Offset(offsite).Find(&result).Error
	if err != nil {
		return nil, 0, err
	}

	return result, total, nil
}

// FindByID implements NewsRepository.
func (n *newsRepository) FindByID(ctx context.Context, tx *gorm.DB, id string) (result *entities.News, err error) {
	err = tx.WithContext(ctx).Where("id = ?", id).First(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

// FindByTitle implements NewsRepository.
func (n *newsRepository) FindByTitle(ctx context.Context, tx *gorm.DB, title string) (result *entities.News, err error) {
	err = tx.WithContext(ctx).Where("title = ?", title).First(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Update implements NewsRepository.
func (n *newsRepository) Update(ctx context.Context, tx *gorm.DB, news *entities.News) (*entities.News, error) {
	err := tx.WithContext(ctx).Preload("ReportType").Save(&news).Error
	if err != nil {
		return nil, err
	}

	return news, nil
}

// Delete implements NewsRepository.
func (n *newsRepository) Delete(ctx context.Context, tx *gorm.DB, news *entities.News) error {
	err := tx.WithContext(ctx).Delete(&news).Error
	if err != nil {
		return err
	}

	return nil
}

// FindBySlug implements NewsRepository.
func (n *newsRepository) FindBySlug(ctx context.Context, tx *gorm.DB, slug string) (result *entities.News, err error) {
	err = tx.WithContext(ctx).Where("slug = ?", slug).First(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}
