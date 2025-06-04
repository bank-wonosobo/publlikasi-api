package repositories

import (
	"context"

	"github.com/bank-wonosobo/publlikasi-api.git/internal/delivery/http/dto"
	"github.com/bank-wonosobo/publlikasi-api.git/internal/entities"
	"gorm.io/gorm"
)

type AnnouncementRepository interface {
	Save(ctx context.Context, announcement *entities.Announcement) (*entities.Announcement, error)
	GetAll(ctx context.Context, params *dto.AnnouncementGetQueryParams, offsite int) ([]entities.Announcement, int64, error)
	Delete(ctx context.Context, news *entities.Announcement) error
	FindByID(ctx context.Context, id string) (*entities.Announcement, error)
	FindByTitle(ctx context.Context, title string) (*entities.Announcement, error)
}

type announcementRepository struct {
	db *gorm.DB
}

func NewAnnouncement(db *gorm.DB) AnnouncementRepository {
	return &announcementRepository{
		db: db,
	}
}

// Save implements AnnouncementRepository.
func (a *announcementRepository) Save(ctx context.Context, announcement *entities.Announcement) (*entities.Announcement, error) {
	err := a.db.WithContext(ctx).Save(&announcement).Error
	if err != nil {
		return nil, err
	}

	return announcement, nil
}

// FindByTitle implements AnnouncementRepository.
func (a *announcementRepository) FindByTitle(ctx context.Context, title string) (result *entities.Announcement, err error) {
	err = a.db.WithContext(ctx).Where("title = ?", title).First(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

// FindByID implements AnnouncementRepository.
func (a *announcementRepository) FindByID(ctx context.Context, id string) (result *entities.Announcement, err error) {
	err = a.db.WithContext(ctx).Where("id = ?", id).First(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetAll implements AnnouncementRepository.
func (a *announcementRepository) GetAll(ctx context.Context, params *dto.AnnouncementGetQueryParams, offsite int) (result []entities.Announcement, total int64, err error) {
	query := a.db.Model(&entities.Announcement{})

	if params.TargetAudience != "" {
		query = query.Where("target_audience = ?", params.TargetAudience)
	}

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

// Delete implements AnnouncementRepository.
func (a *announcementRepository) Delete(ctx context.Context, news *entities.Announcement) error {
	err := a.db.WithContext(ctx).Delete(&news).Error
	if err != nil {
		return err
	}

	return nil
}
