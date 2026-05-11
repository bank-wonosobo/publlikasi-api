package repositories

import (
	"context"

	"github.com/bank-wonosobo/publlikasi-api.git/internal/delivery/http/dto"
	"github.com/bank-wonosobo/publlikasi-api.git/internal/entities"
	"gorm.io/gorm"
)

type VisionMissionRepository interface {
	Save(ctx context.Context, item *entities.VisionMission) (*entities.VisionMission, error)
	GetAll(ctx context.Context, params *dto.VisionMissionGetQueryParams, offset int) ([]entities.VisionMission, int64, error)
	Delete(ctx context.Context, item *entities.VisionMission) error
	FindByID(ctx context.Context, id string) (*entities.VisionMission, error)
	FindByTitle(ctx context.Context, title string) (*entities.VisionMission, error)
}

type visionMissionRepository struct {
	db *gorm.DB
}

func NewVisionMission(db *gorm.DB) VisionMissionRepository {
	return &visionMissionRepository{db: db}
}

func (v *visionMissionRepository) Save(ctx context.Context, item *entities.VisionMission) (*entities.VisionMission, error) {
	err := v.db.WithContext(ctx).Save(&item).Error
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (v *visionMissionRepository) GetAll(ctx context.Context, params *dto.VisionMissionGetQueryParams, offset int) (result []entities.VisionMission, total int64, err error) {
	query := v.db.WithContext(ctx).Model(&entities.VisionMission{})

	if params.Key != "" {
		query = query.Where("title ILIKE ?", "%"+params.Key+"%").
			Or("vision ILIKE ?", "%"+params.Key+"%").
			Or("mission ILIKE ?", "%"+params.Key+"%")
	}

	query.Count(&total)

	err = query.Limit(params.Limit).Offset(offset).Find(&result).Error
	if err != nil {
		return nil, 0, err
	}

	return result, total, nil
}

func (v *visionMissionRepository) Delete(ctx context.Context, item *entities.VisionMission) error {
	err := v.db.WithContext(ctx).Delete(&item).Error
	if err != nil {
		return err
	}

	return nil
}

func (v *visionMissionRepository) FindByID(ctx context.Context, id string) (result *entities.VisionMission, err error) {
	err = v.db.WithContext(ctx).Where("id = ?", id).First(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (v *visionMissionRepository) FindByTitle(ctx context.Context, title string) (result *entities.VisionMission, err error) {
	err = v.db.WithContext(ctx).Where("title = ?", title).First(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}
