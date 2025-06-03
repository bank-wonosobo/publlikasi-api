package repositories

import (
	"context"

	"github.com/bank-wonosobo/publlikasi-api.git/internal/delivery/http/dto"
	"github.com/bank-wonosobo/publlikasi-api.git/internal/entities"
	"gorm.io/gorm"
)

type AnnouncementRepository interface {
	Save(ctx context.Context, tx *gorm.DB, announcement *entities.Announcement) (*entities.Announcement, error)
	GetAll(ctx context.Context, tx *gorm.DB, params *dto.AnnouncementGetQueryParams, offsite int) ([]entities.Announcement, int64, error)
	Update(ctx context.Context, tx *gorm.DB, news *entities.Announcement) (*entities.Announcement, error)
	FindByID(ctx context.Context, tx *gorm.DB, id string) (*entities.Announcement, error)
	FindByTitle(ctx context.Context, tx *gorm.DB, title string) (*entities.Announcement, error)
}

type announcementRepository struct {
}

func NewAnnouncement() AnnouncementRepository {
	return &announcementRepository{}
}

// Save implements AnnouncementRepository.
func (a *announcementRepository) Save(ctx context.Context, tx *gorm.DB, announcement *entities.Announcement) (*entities.Announcement, error) {
	err := tx.WithContext(ctx).Create(&announcement).Error
	if err != nil {
		return nil, err
	}

	return announcement, nil
}

// FindByTitle implements AnnouncementRepository.
func (a *announcementRepository) FindByTitle(ctx context.Context, tx *gorm.DB, title string) (result *entities.Announcement, err error) {
	err = tx.WithContext(ctx).Where("title = ?", title).First(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

// FindByID implements AnnouncementRepository.
func (a *announcementRepository) FindByID(ctx context.Context, tx *gorm.DB, id string) (result *entities.Announcement, err error) {
	err = tx.WithContext(ctx).Where("id = ?", id).First(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetAll implements AnnouncementRepository.
func (a *announcementRepository) GetAll(ctx context.Context, tx *gorm.DB, params *dto.AnnouncementGetQueryParams, offsite int) (result []entities.Announcement, total int64, err error) {
	query := tx.Model(&entities.Announcement{})

	if params.Key != "" {
		query = query.Where("title ILIKE ?", "%"+params.Key+"%").Or("content ILIKE ?", "%"+params.Key+"%")
	}

	if params.TargetAudience != "" {
		query = query.Where("target_audience = ?", params.TargetAudience)
	}

	if params.IsActive != 0 {
		query = query.Where("is_active = ?", params.IsActive)
	}

	if params.Status != "" {
		query = query.Where("status = ?", params.Status)
	}

	query.Count(&total)

	err = query.Limit(params.Limit).Offset(offsite).Find(&result).Error
	if err != nil {
		return nil, 0, err
	}

	return result, total, nil
}

// Update implements AnnouncementRepository.
func (a *announcementRepository) Update(ctx context.Context, tx *gorm.DB, announcement *entities.Announcement) (*entities.Announcement, error) {
	err := tx.WithContext(ctx).Save(&announcement).Error
	if err != nil {
		return nil, err
	}

	return announcement, nil
}
