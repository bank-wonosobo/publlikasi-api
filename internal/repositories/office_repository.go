package repositories

import (
	"context"

	"github.com/bank-wonosobo/publlikasi-api.git/internal/entities"
	"gorm.io/gorm"
)

type OfficeRepository interface {
	Save(ctx context.Context, office *entities.Office) (*entities.Office, error)
	GetAll(ctx context.Context) ([]entities.Office, error)
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
func (o *officeRepository) GetAll(ctx context.Context) ([]entities.Office, error) {
	panic("unimplemented")
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
