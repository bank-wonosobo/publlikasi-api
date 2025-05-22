package repositories

import (
	"context"

	"github.com/bank-wonosobo/publlikasi-api.git/internal/entities"
	"gorm.io/gorm"
)

type ReportRepository interface {
	Save(ctx context.Context, tx *gorm.DB, report entities.Report) (*entities.Report, error)
	GetAll()
	Update()
	Delete()
	FindByID()
	FindByTitle(ctx context.Context, tx *gorm.DB, title string) (*entities.Report, error)
}

type reportRepository struct {
}

func NewReport() ReportRepository {
	return &reportRepository{}
}

// Save implements ReportRepository.
func (r *reportRepository) Save(ctx context.Context, tx *gorm.DB, report entities.Report) (*entities.Report, error) {
	err := tx.WithContext(ctx).Preload("report_types").Create(&report).Error
	if err != nil {
		return nil, err
	}

	return &report, nil
}

// Delete implements ReportRepository.
func (r *reportRepository) Delete() {
	panic("unimplemented")
}

// FindByID implements ReportRepository.
func (r *reportRepository) FindByID() {
	panic("unimplemented")
}

// GetAll implements ReportRepository.
func (r *reportRepository) GetAll() {
	panic("unimplemented")
}

// Update implements ReportRepository.
func (r *reportRepository) Update() {
	panic("unimplemented")
}

// FindByTitle implements ReportRepository.
func (r *reportRepository) FindByTitle(ctx context.Context, tx *gorm.DB, title string) (result *entities.Report, err error) {
	err = tx.WithContext(ctx).Where("title = ?", title).First(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}
