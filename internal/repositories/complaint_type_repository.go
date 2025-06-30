package repositories

import (
	"context"

	"github.com/bank-wonosobo/publlikasi-api.git/internal/entities"
	"gorm.io/gorm"
)

type ComplaintTypeRepository interface {
	Save(ctx context.Context, complaintType *entities.ComplaintType) (*entities.ComplaintType, error)
	Delete(ctx context.Context, complaintType *entities.ComplaintType) error
	FindAll(ctx context.Context) ([]entities.ComplaintType, error)
	FindByID(ctx context.Context, id int) (*entities.ComplaintType, error)
	FindByName(ctx context.Context, name string) (*entities.ComplaintType, error)
}

type complaintTypeRepository struct {
	db *gorm.DB
}

// Delete implements ComplaintTypeRepository.
func (c *complaintTypeRepository) Delete(ctx context.Context, complaintType *entities.ComplaintType) error {
	err := c.db.WithContext(ctx).Delete(&complaintType).Error
	if err != nil {
		return err
	}

	return nil
}

// FindAll implements ComplaintTypeRepository.
func (c *complaintTypeRepository) FindAll(ctx context.Context) (result []entities.ComplaintType, err error) {
	err = c.db.WithContext(ctx).Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

// FindByID implements ComplaintTypeRepository.
func (c *complaintTypeRepository) FindByID(ctx context.Context, id int) (result *entities.ComplaintType, err error) {
	err = c.db.WithContext(ctx).Where("id = ?", id).First(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

// FindByName implements ComplaintTypeRepository.
func (c *complaintTypeRepository) FindByName(ctx context.Context, name string) (result *entities.ComplaintType, err error) {
	err = c.db.WithContext(ctx).Where("name = ?", name).First(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Save implements ComplaintTypeRepository.
func (c *complaintTypeRepository) Save(ctx context.Context, complaintType *entities.ComplaintType) (*entities.ComplaintType, error) {
	err := c.db.WithContext(ctx).Save(&complaintType).Error
	if err != nil {
		return nil, err
	}

	return complaintType, nil
}

func NewComplaintType(db *gorm.DB) ComplaintTypeRepository {
	return &complaintTypeRepository{
		db: db,
	}
}
