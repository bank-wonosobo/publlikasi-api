package repositories

import (
	"context"

	"github.com/bank-wonosobo/publlikasi-api.git/internal/entities"
	"gorm.io/gorm"
)

type ReportTypeRepository interface {
	Save(ctx context.Context, tx *gorm.DB, reportType *entities.ReportType) (*entities.ReportType, error)
	Update(ctx context.Context, tx *gorm.DB, reportType *entities.ReportType) (*entities.ReportType, error)
	Delete(ctx context.Context, tx *gorm.DB, reportType *entities.ReportType) error
	FindAll(ctx context.Context, tx *gorm.DB) ([]entities.ReportType, error)
	FindByID(ctx context.Context, tx *gorm.DB, id int) (*entities.ReportType, error)
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

// Update implements ReportTypeRepository.
func (r *reportTypeRepository) Update(ctx context.Context, tx *gorm.DB, reportType *entities.ReportType) (*entities.ReportType, error) {
	err := tx.WithContext(ctx).Save(&reportType).Error
	if err != nil {
		return nil, err
	}

	return reportType, nil
}

// Delete implements ReportTypeRepository.
func (r *reportTypeRepository) Delete(ctx context.Context, tx *gorm.DB, reportType *entities.ReportType) error {
	err := tx.WithContext(ctx).Delete(&reportType).Error
	if err != nil {
		return err
	}

	return nil
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
func (r *reportTypeRepository) FindByID(ctx context.Context, tx *gorm.DB, id int) (result *entities.ReportType, err error) {
	err = tx.WithContext(ctx).Where("id = ?", id).First(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

// FindByName implements ReportTypeRepository.
func (r *reportTypeRepository) FindByName(ctx context.Context, tx *gorm.DB, name string) (result *entities.ReportType, err error) {
	err = tx.WithContext(ctx).Where("name = ?", name).First(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}
