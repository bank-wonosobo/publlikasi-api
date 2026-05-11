package repositories

import (
	"context"

	"github.com/bank-wonosobo/publlikasi-api.git/internal/delivery/http/dto"
	"github.com/bank-wonosobo/publlikasi-api.git/internal/entities"
	"gorm.io/gorm"
)

type NewsRepository interface {
	Save(ctx context.Context, news *entities.News) (*entities.News, error)
	GetAll(ctx context.Context, params *dto.NewsGetQueryParams, offsite int) ([]entities.News, int64, error)
	Delete(ctx context.Context, news *entities.News) error
	FindByID(ctx context.Context, id string) (*entities.News, error)
	FindBySlug(ctx context.Context, slug string) (*entities.News, error)
	FindByTitle(ctx context.Context, title string) (*entities.News, error)
}

type newsRepository struct {
	db *gorm.DB
}

func NewNews(db *gorm.DB) NewsRepository {
	return &newsRepository{
		db: db,
	}
}

// Create implements NewsRepository.
func (n *newsRepository) Save(ctx context.Context, news *entities.News) (*entities.News, error) {
	err := n.db.WithContext(ctx).Save(&news).Error
	if err != nil {
		return nil, err
	}

	return news, nil
}

// GetAll implements NewsRepository.
func (n *newsRepository) GetAll(ctx context.Context, params *dto.NewsGetQueryParams, offsite int) (result []entities.News, total int64, err error) {
	query := n.db.Model(&entities.News{})

	if params.Status != "" {
		query = query.Where("status = ?", params.Status)
	}

	if params.Key != "" {
		query = query.Where("title ILIKE ?", "%"+params.Key+"%").Or("content ILIKE ?", "%"+params.Key+"%")
	}

	query.Count(&total)

	err = query.Order("published_at DESC").Limit(params.Limit).Offset(offsite).Find(&result).Error
	if err != nil {
		return nil, 0, err
	}

	return result, total, nil
}

// FindByID implements NewsRepository.
func (n *newsRepository) FindByID(ctx context.Context, id string) (result *entities.News, err error) {
	err = n.db.WithContext(ctx).Where("id = ?", id).First(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

// FindByTitle implements NewsRepository.
func (n *newsRepository) FindByTitle(ctx context.Context, title string) (result *entities.News, err error) {
	err = n.db.WithContext(ctx).Where("title = ?", title).First(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Delete implements NewsRepository.
func (n *newsRepository) Delete(ctx context.Context, news *entities.News) error {
	err := n.db.WithContext(ctx).Delete(&news).Error
	if err != nil {
		return err
	}

	return nil
}

// FindBySlug implements NewsRepository.
func (n *newsRepository) FindBySlug(ctx context.Context, slug string) (result *entities.News, err error) {
	err = n.db.WithContext(ctx).Where("slug = ?", slug).First(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}
