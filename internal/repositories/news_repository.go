package repositories

import (
	"context"

	"github.com/bank-wonosobo/publlikasi-api.git/internal/entities"
	"gorm.io/gorm"
)

type NewsRepository interface {
	Save(ctx context.Context, tx *gorm.DB, news entities.News) (*entities.News, error)
	GetAll() ([]entities.News, error)
	Update() (*entities.News, error)
	Delete() (*entities.News, error)
	FindByID() (*entities.News, error)
	FindBySlug() (*entities.News, error)
	FindByTitle(ctx context.Context, tx *gorm.DB, title string) (*entities.News, error)
}

type newsRepository struct {
}

func NewNews() NewsRepository {
	return &newsRepository{}
}

// Create implements NewsRepository.
func (n *newsRepository) Save(ctx context.Context, tx *gorm.DB, news entities.News) (*entities.News, error) {
	err := tx.WithContext(ctx).Create(&news).Error
	if err != nil {
		return nil, err
	}

	return &news, nil
}

// Delete implements NewsRepository.
func (n *newsRepository) Delete() (*entities.News, error) {
	panic("unimplemented")
}

// FindByID implements NewsRepository.
func (n *newsRepository) FindByID() (*entities.News, error) {
	panic("unimplemented")
}

// FindBySlug implements NewsRepository.
func (n *newsRepository) FindBySlug() (*entities.News, error) {
	panic("unimplemented")
}

// FindByTitle implements NewsRepository.
func (n *newsRepository) FindByTitle(ctx context.Context, tx *gorm.DB, title string) (result *entities.News, err error) {
	err = tx.WithContext(ctx).Where("title = ?", title).First(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetAll implements NewsRepository.
func (n *newsRepository) GetAll() ([]entities.News, error) {
	panic("unimplemented")
}

// Update implements NewsRepository.
func (n *newsRepository) Update() (*entities.News, error) {
	panic("unimplemented")
}
