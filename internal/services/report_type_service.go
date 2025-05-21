package services

import (
	"context"
	"errors"

	"github.com/bank-wonosobo/publlikasi-api.git/internal/delivery/http/dto/request"
	"github.com/bank-wonosobo/publlikasi-api.git/internal/delivery/http/dto/response"
	"github.com/bank-wonosobo/publlikasi-api.git/internal/entities"
	"github.com/bank-wonosobo/publlikasi-api.git/internal/repositories"
	"gorm.io/gorm"
)

type ReportTypeService interface {
	Index(ctx context.Context) (*response.ReportTypeResponse, error)
	Create(ctx context.Context, req *request.ReportTypeCreateRequest) (*response.ReportTypeResponse, error)
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
func (r *reportTypeService) Index(ctx context.Context) (*response.ReportTypeResponse, error) {
	panic("unimplemented")
}

// Create implements ReportTypeService.
func (r *reportTypeService) Create(ctx context.Context, req *request.ReportTypeCreateRequest) (*response.ReportTypeResponse, error) {
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
	reportTypeResponse := response.ReportTypeResponse{
		ID:          result.ID,
		Name:        result.Name,
		Description: result.Description,
		CreatedAt:   result.CreatedAt.String(),
		UpdatedAt:   result.UpdatedAt.String(),
	}

	return &reportTypeResponse, nil
}
