package repositories

import (
	"context"

	"github.com/bank-wonosobo/publlikasi-api.git/internal/delivery/http/dto"
	"github.com/bank-wonosobo/publlikasi-api.git/internal/entities"
	"gorm.io/gorm"
)

type ProfileRepository interface {
	Save(ctx context.Context, profile *entities.Profile) (*entities.Profile, error)
	GetAll(ctx context.Context, params *dto.ProfileGetQueryParams, offset int) ([]entities.Profile, int64, error)
	Delete(ctx context.Context, profile *entities.Profile) error
	FindByID(ctx context.Context, id string) (*entities.Profile, error)
	FindByTitle(ctx context.Context, title string) (*entities.Profile, error)
}

type profileRepository struct {
	db *gorm.DB
}

func NewProfile(db *gorm.DB) ProfileRepository {
	return &profileRepository{
		db: db,
	}
}

func (p *profileRepository) Save(ctx context.Context, profile *entities.Profile) (*entities.Profile, error) {
	err := p.db.WithContext(ctx).Save(&profile).Error
	if err != nil {
		return nil, err
	}

	return profile, nil
}

func (p *profileRepository) GetAll(ctx context.Context, params *dto.ProfileGetQueryParams, offset int) (result []entities.Profile, total int64, err error) {
	query := p.db.WithContext(ctx).Model(&entities.Profile{})

	if params.Key != "" {
		query = query.Where("title ILIKE ?", "%"+params.Key+"%").Or("description ILIKE ?", "%"+params.Key+"%")
	}

	query.Count(&total)

	err = query.Limit(params.Limit).Offset(offset).Find(&result).Error
	if err != nil {
		return nil, 0, err
	}

	return result, total, nil
}

func (p *profileRepository) Delete(ctx context.Context, profile *entities.Profile) error {
	err := p.db.WithContext(ctx).Delete(&profile).Error
	if err != nil {
		return err
	}

	return nil
}

func (p *profileRepository) FindByID(ctx context.Context, id string) (result *entities.Profile, err error) {
	err = p.db.WithContext(ctx).Where("id = ?", id).First(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (p *profileRepository) FindByTitle(ctx context.Context, title string) (result *entities.Profile, err error) {
	err = p.db.WithContext(ctx).Where("title = ?", title).First(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}
