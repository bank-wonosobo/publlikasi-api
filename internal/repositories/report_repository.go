package repositories

import (
	"context"

	"github.com/bank-wonosobo/publlikasi-api.git/internal/delivery/http/dto/request"
	"github.com/bank-wonosobo/publlikasi-api.git/internal/entities"
	"gorm.io/gorm"
)

type ReportRepository interface {
	Save(ctx context.Context, tx *gorm.DB, report *entities.Report) (*entities.Report, error)
	GetAll(ctx context.Context, tx *gorm.DB, params *request.ReportGetQueryParams, offside int) ([]entities.Report, int64, error)
	Update(ctx context.Context, tx *gorm.DB, report *entities.Report) (*entities.Report, error)
	Delete(ctx context.Context, tx *gorm.DB, report *entities.Report) error
	FindByID(ctx context.Context, tx *gorm.DB, id string) (*entities.Report, error)
	FindByTitle(ctx context.Context, tx *gorm.DB, title string) (*entities.Report, error)
}

type reportRepository struct {
}

func NewReport() ReportRepository {
	return &reportRepository{}
}

// GetAll implements ReportRepository.
func (r *reportRepository) GetAll(ctx context.Context, tx *gorm.DB, params *request.ReportGetQueryParams, offside int) (result []entities.Report, total int64, err error) {
	query := tx.Model(&entities.Report{})

	if params.Title != "" {
		query = query.Where("title ILIKE ?", "%"+params.Title+"%")
	}

	if params.Description != "" {
		query = query.Where("description ILIKE ?", "%"+params.Description+"%")
	}

	if params.Year != 0 {
		query = query.Where("year = ?", params.Year)
	}

	query.Count(&total)

	err = query.Preload("ReportType").Limit(params.Limit).Offset(offside).Find(&result).Error
	if err != nil {
		return nil, 0, err
	}

	return result, total, nil
}

// Save implements ReportRepository.
func (r *reportRepository) Save(ctx context.Context, tx *gorm.DB, report *entities.Report) (*entities.Report, error) {
	err := tx.WithContext(ctx).Preload("ReportTypes").Create(&report).Error
	if err != nil {
		return nil, err
	}

	return report, nil
}

// Update implements ReportRepository.
func (r *reportRepository) Update(ctx context.Context, tx *gorm.DB, report *entities.Report) (*entities.Report, error) {
	err := tx.WithContext(ctx).Save(&report).Error
	if err != nil {
		return nil, err
	}

	return report, nil
}

// FindByID implements ReportRepository.
func (r *reportRepository) FindByID(ctx context.Context, tx *gorm.DB, id string) (result *entities.Report, err error) {
	err = tx.WithContext(ctx).Preload("ReportType").Where("id = ?", id).First(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

// FindByTitle implements ReportRepository.
func (r *reportRepository) FindByTitle(ctx context.Context, tx *gorm.DB, title string) (result *entities.Report, err error) {
	err = tx.WithContext(ctx).Preload("ReportType").Where("title = ?", title).First(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Delete implements ReportRepository.
func (r *reportRepository) Delete(ctx context.Context, tx *gorm.DB, report *entities.Report) error {
	err := tx.WithContext(ctx).Delete(&report).Error
	if err != nil {
		return err
	}

	return nil
}
