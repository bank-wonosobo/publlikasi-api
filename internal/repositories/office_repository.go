package repositories

import (
	"context"

	"github.com/bank-wonosobo/publlikasi-api.git/internal/delivery/http/dto"
	"github.com/bank-wonosobo/publlikasi-api.git/internal/entities"
	"gorm.io/gorm"
)

type OfficeRepository interface {
	Save(ctx context.Context, office *entities.Office) (*entities.Office, error)
	GetAll(ctx context.Context, params *dto.OfficeGetQueryParams, offsite int) ([]entities.Office, int64, error)
	Delete(ctx context.Context, office *entities.Office) error
	FindByID(ctx context.Context, id string) (*entities.Office, error)
	FindByName(ctx context.Context, name string) (*entities.Office, error)
}

type officeRepository struct {
	db *gorm.DB
}

// Delete implements OfficeRepository.
func (o *officeRepository) Delete(ctx context.Context, office *entities.Office) error {
	panic("unimplemented")
}

// FindByID implements OfficeRepository.
func (o *officeRepository) FindByID(ctx context.Context, id string) (result *entities.Office, err error) {
	err = o.db.WithContext(ctx).Where("id = ?", id).First(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

// FindByTitle implements OfficeRepository.
func (o *officeRepository) FindByName(ctx context.Context, name string) (result *entities.Office, err error) {
	err = o.db.WithContext(ctx).Where("name = ?", name).First(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetAll implements OfficeRepository.
func (o *officeRepository) GetAll(ctx context.Context, params *dto.OfficeGetQueryParams, offsite int) (result []entities.Office, total int64, err error) {
	query := o.db.Model(&entities.Office{})

	if params.Key != "" {
		query = query.Where("name ILIKE ?", "%"+params.Key+"%").Or("address ILIKE ?", "%"+params.Key+"%")
	}

	query.Count(&total)

	err = query.Limit(params.Limit).Offset(offsite).Find(&result).Error
	if err != nil {
		return nil, 0, err
	}

	return result, total, nil
}

// Save implements OfficeRepository.
func (o *officeRepository) Save(ctx context.Context, office *entities.Office) (*entities.Office, error) {
	err := o.db.WithContext(ctx).Save(&office).Error
	if err != nil {
		return nil, err
	}

	return office, nil
}

func NewOffice(db *gorm.DB) OfficeRepository {
	return &officeRepository{
		db: db,
	}
}
