package repositories

import (
	"context"

	"github.com/bank-wonosobo/publlikasi-api.git/internal/entities"
	"gorm.io/gorm"
)

type ReportRepository interface {
	Save(ctx context.Context, tx *gorm.DB, report *entities.Report) (*entities.Report, error)
	GetAll()
	Update(ctx context.Context, tx *gorm.DB, report *entities.Report) (*entities.Report, error)
	Delete()
	FindByID(ctx context.Context, tx *gorm.DB, id string) (*entities.Report, error)
	FindByTitle(ctx context.Context, tx *gorm.DB, title string) (*entities.Report, error)
}

type reportRepository struct {
}

func NewReport() ReportRepository {
	return &reportRepository{}
}

// Save implements ReportRepository.
func (r *reportRepository) Save(ctx context.Context, tx *gorm.DB, report *entities.Report) (*entities.Report, error) {
	err := tx.WithContext(ctx).Preload("ReportTypes").Create(&report).Error
	if err != nil {
		return nil, err
	}

	return report, nil
}

// GetAll implements ReportRepository.
func (r *reportRepository) GetAll() {
	panic("unimplemented")
}

// Update implements ReportRepository.
func (r *reportRepository) Update(ctx context.Context, tx *gorm.DB, report *entities.Report) (*entities.Report, error) {
	err := tx.WithContext(ctx).Save(&report).Error
	if err != nil {
		return nil, err
	}

	return report, nil
}

// Delete implements ReportRepository.
func (r *reportRepository) Delete() {
	panic("unimplemented")
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
