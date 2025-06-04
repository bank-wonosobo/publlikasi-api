package repositories

import (
	"context"

	"github.com/bank-wonosobo/publlikasi-api.git/internal/entities"
	"gorm.io/gorm"
)

type ReportTypeRepository interface {
	Save(ctx context.Context, reportType *entities.ReportType) (*entities.ReportType, error)
	Delete(ctx context.Context, reportType *entities.ReportType) error
	FindAll(ctx context.Context) ([]entities.ReportType, error)
	FindByID(ctx context.Context, id int) (*entities.ReportType, error)
	FindByName(ctx context.Context, name string) (*entities.ReportType, error)
}

type reportTypeRepository struct {
	db *gorm.DB
}

func NewReportType(db *gorm.DB) ReportTypeRepository {
	return &reportTypeRepository{
		db: db,
	}
}

// Save implements ReportTypeRepository.
func (r *reportTypeRepository) Save(ctx context.Context, reportType *entities.ReportType) (*entities.ReportType, error) {
	err := r.db.WithContext(ctx).Save(&reportType).Error
	if err != nil {
		return nil, err
	}

	return reportType, nil
}

// Delete implements ReportTypeRepository.
func (r *reportTypeRepository) Delete(ctx context.Context, reportType *entities.ReportType) error {
	err := r.db.WithContext(ctx).Delete(&reportType).Error
	if err != nil {
		return err
	}

	return nil
}

// FindAll implements ReportTypeRepository.
func (r *reportTypeRepository) FindAll(ctx context.Context) (result []entities.ReportType, err error) {
	err = r.db.WithContext(ctx).Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

// FindByID implements ReportTypeRepository.
func (r *reportTypeRepository) FindByID(ctx context.Context, id int) (result *entities.ReportType, err error) {
	err = r.db.WithContext(ctx).Where("id = ?", id).First(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

// FindByName implements ReportTypeRepository.
func (r *reportTypeRepository) FindByName(ctx context.Context, name string) (result *entities.ReportType, err error) {
	err = r.db.WithContext(ctx).Where("name = ?", name).First(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}
