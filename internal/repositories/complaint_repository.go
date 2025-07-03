package repositories

import (
	"context"

	"github.com/bank-wonosobo/publlikasi-api.git/internal/delivery/http/dto"
	"github.com/bank-wonosobo/publlikasi-api.git/internal/entities"
	"gorm.io/gorm"
)

type ComplaintRepository interface {
	Save(ctx context.Context, news *entities.Complaint) (*entities.Complaint, error)
	GetAll(ctx context.Context, params *dto.ComplaintGetQueryParams, offsite int) ([]entities.Complaint, int64, error)
	Delete(ctx context.Context, news *entities.Complaint) error
	FindByID(ctx context.Context, id string) (*entities.Complaint, error)
	FindByTitle(ctx context.Context, title string) (*entities.Complaint, error)
}

type complaintRepository struct {
	db *gorm.DB
}

// Delete implements ComplaintRepository.
func (c *complaintRepository) Delete(ctx context.Context, news *entities.Complaint) error {
	panic("unimplemented")
}

// FindByID implements ComplaintRepository.
func (c *complaintRepository) FindByID(ctx context.Context, id string) (result *entities.Complaint, err error) {
	err = c.db.WithContext(ctx).Preload("ComplaintType").Where("id = ?", id).First(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

// FindByTitle implements ComplaintRepository.
func (c *complaintRepository) FindByTitle(ctx context.Context, title string) (*entities.Complaint, error) {
	panic("unimplemented")
}

// GetAll implements ComplaintRepository.
func (c *complaintRepository) GetAll(ctx context.Context, params *dto.ComplaintGetQueryParams, offsite int) (result []entities.Complaint, total int64, err error) {
	query := c.db.Model(&entities.Complaint{})

	if params.Type != "" {
		query = query.Where("complaint_type_id = ?", params.Type)
	}

	if params.Status != "" {
		query = query.Where("status = ?", params.Status)
	}

	if params.Key != "" {
		query = query.Preload("ComplaintType").Where("reported_name ILIKE ?", "%"+params.Key+"%").Or("description ILIKE ?", "%"+params.Key+"%").Or("insident_location ILIKE ?", "%"+params.Key+"%")
	}

	query.Count(&total)

	err = query.Limit(params.Limit).Offset(offsite).Find(&result).Error
	if err != nil {
		return nil, 0, err
	}

	return result, total, nil
}

// Save implements ComplaintRepository.
func (c *complaintRepository) Save(ctx context.Context, complaint *entities.Complaint) (*entities.Complaint, error) {
	err := c.db.Preload("ComplaintType").WithContext(ctx).Save(&complaint).Error
	if err != nil {
		return nil, err
	}

	return complaint, nil
}

func NewComplaint(db *gorm.DB) ComplaintRepository {
	return &complaintRepository{
		db: db,
	}
}
