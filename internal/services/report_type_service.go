package services

import (
	"context"
	"errors"

	"github.com/bank-wonosobo/publlikasi-api.git/internal/delivery/http/dto"
	"github.com/bank-wonosobo/publlikasi-api.git/internal/entities"
	"github.com/bank-wonosobo/publlikasi-api.git/internal/repositories"
	"gorm.io/gorm"
)

type ReportTypeService interface {
	Index(ctx context.Context) ([]dto.ReportTypeResponse, error)
	Create(ctx context.Context, req *dto.ReportTypeCreateRequest) (*dto.ReportTypeResponse, error)
	Update(ctx context.Context, req *dto.ReportTypeCreateRequest, id int) (*dto.ReportTypeResponse, error)
	Delete(ctx context.Context, id int) error
}

type reportTypeService struct {
	db             *gorm.DB
	reportTypeRepo repositories.ReportTypeRepository
}

func NewReportType(db *gorm.DB,
	reportTypeRepo repositories.ReportTypeRepository) ReportTypeService {
	return &reportTypeService{
		db:             db,
		reportTypeRepo: reportTypeRepo,
	}
}

// Index implements ReportTypeService.
func (r *reportTypeService) Index(ctx context.Context) ([]dto.ReportTypeResponse, error) {
	// get all
	result, err := r.reportTypeRepo.FindAll(ctx, r.db)
	if err != nil {
		return nil, err
	}

	// return result
	reportTypeResponse := toReportTypeResponses(result)

	return reportTypeResponse, nil
}

// Create implements ReportTypeService.
func (r *reportTypeService) Create(ctx context.Context, req *dto.ReportTypeCreateRequest) (*dto.ReportTypeResponse, error) {
	// check if name exist
	_, err := r.reportTypeRepo.FindByName(ctx, r.db, req.Name)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("jenis laporan sudah ada")
	}
	// create report type
	reportType := entities.ReportType{
		Name:        req.Name,
		Description: req.Description,
	}
	result, err := r.reportTypeRepo.Save(ctx, r.db, &reportType)
	if err != nil {
		return nil, err
	}

	// return result
	reportTypeResponse := toReportTypeResponse(*result)

	return &reportTypeResponse, nil
}

// Update implements ReportTypeService.
func (r *reportTypeService) Update(ctx context.Context, req *dto.ReportTypeCreateRequest, id int) (*dto.ReportTypeResponse, error) {
	// get report type by id
	reportType, err := r.reportTypeRepo.FindByID(ctx, r.db, id)
	if err != nil {
		return nil, errors.New("report type tidak ditemukan")
	}

	// update report type
	reportType.Name = req.Name
	reportType.Description = req.Description
	result, err := r.reportTypeRepo.Update(ctx, r.db, reportType)
	if err != nil {
		return nil, err
	}

	// return result
	reportTypeResponse := toReportTypeResponse(*result)

	return &reportTypeResponse, nil
}

// Delete implements ReportTypeService.
func (r *reportTypeService) Delete(ctx context.Context, id int) error {
	// get report type by id
	reportType, err := r.reportTypeRepo.FindByID(ctx, r.db, id)
	if err != nil {
		return errors.New("report type tidak ditemukan")
	}

	err = r.reportTypeRepo.Delete(ctx, r.db, reportType)
	if err != nil {
		return err
	}

	return nil
}

func toReportTypeResponse(reportType entities.ReportType) (result dto.ReportTypeResponse) {
	result = dto.ReportTypeResponse{
		ID:          reportType.ID,
		Name:        reportType.Name,
		Description: reportType.Description,
		CreatedAt:   reportType.CreatedAt.String(),
		UpdatedAt:   reportType.UpdatedAt.String(),
	}

	return result
}

func toReportTypeResponses(reportTypes []entities.ReportType) []dto.ReportTypeResponse {
	var reportTypeResponse []dto.ReportTypeResponse
	for _, reportType := range reportTypes {
		reportTypeResponse = append(reportTypeResponse, toReportTypeResponse(reportType))
	}

	return reportTypeResponse
}
