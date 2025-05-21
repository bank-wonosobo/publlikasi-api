package repositories

import (
	"context"

	"github.com/bank-wonosobo/publlikasi-api.git/internal/entities"
	"gorm.io/gorm"
)

type ReportTypeRepository interface {
	Save(ctx context.Context, tx *gorm.DB, reportType *entities.ReportType) (*entities.ReportType, error)
	Update()
	Delete()
	FindAll(ctx context.Context, tx *gorm.DB) ([]entities.ReportType, error)
	FindByID()
	FindByName(ctx context.Context, tx *gorm.DB, name string) (*entities.ReportType, error)
}

type reportTypeRepository struct {
}

func NewReportType() ReportTypeRepository {
	return &reportTypeRepository{}
}

// Save implements ReportTypeRepository.
func (r *reportTypeRepository) Save(ctx context.Context, tx *gorm.DB, reportType *entities.ReportType) (*entities.ReportType, error) {
	err := tx.WithContext(ctx).Create(&reportType).Error
	if err != nil {
		return nil, err
	}

	return reportType, nil
}

// Delete implements ReportTypeRepository.
func (r *reportTypeRepository) Delete() {
	panic("unimplemented")
}

// FindAll implements ReportTypeRepository.
func (r *reportTypeRepository) FindAll(ctx context.Context, tx *gorm.DB) (result []entities.ReportType, err error) {
	err = tx.WithContext(ctx).Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

// FindByID implements ReportTypeRepository.
func (r *reportTypeRepository) FindByID() {
	panic("unimplemented")
}

// Update implements ReportTypeRepository.
func (r *reportTypeRepository) Update() {
	panic("unimplemented")
}

// FindByName implements ReportTypeRepository.
func (r *reportTypeRepository) FindByName(ctx context.Context, tx *gorm.DB, name string) (result *entities.ReportType, err error) {
	err = tx.WithContext(ctx).Where("name = ?", name).First(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}
