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

type ReportService interface {
	Create(ctx context.Context, req *request.ReportCreateRequest) (*response.ReportResponse, error)
}

type reportService struct {
	db             *gorm.DB
	reportRepo     repositories.ReportRepository
	reportTypeRepo repositories.ReportTypeRepository
}

func NewReport(db *gorm.DB,
	reportRepo repositories.ReportRepository,
	reportTypeRepo repositories.ReportTypeRepository) ReportService {
	return &reportService{
		db:             db,
		reportRepo:     reportRepo,
		reportTypeRepo: reportTypeRepo,
	}
}

// Create implements ReportService.
func (r *reportService) Create(ctx context.Context, req *request.ReportCreateRequest) (*response.ReportResponse, error) {
	// check if title exist
	_, err := r.reportRepo.FindByTitle(ctx, r.db, req.Title)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("judul laporan sudah ada")
	}

	// check report type id
	reportType, err := r.reportTypeRepo.FindByName(ctx, r.db, req.ReportType)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("report type tidak ditemukan")
	}

	// create report type
	report := entities.Report{
		Title:       req.Title,
		Description: req.Description,
		PeriodStart: req.PeriodStart,
		PeriodEnd:   req.PeriodEnd,
		Year:        req.Year,
		Quarter:     req.Quarter,
		Version:     req.Version,
		ReportType:  *reportType,
		FileUrl:     "http://fileurl",
		Status:      entities.Draft,
		UploadBy:    "user",
	}
	result, err := r.reportRepo.Save(ctx, r.db, report)
	if err != nil {
		return nil, err
	}

	// return result
	reportResponse := response.ReportResponse{
		ID:          result.ID,
		Title:       result.Title,
		Description: result.Description,
		PeriodStart: result.PeriodStart,
		PeriodEnd:   result.PeriodEnd,
		Year:        result.Year,
		Quarter:     result.Quarter,
		FileUrl:     result.FileUrl,
		Version:     result.Version,
		Status:      string(result.Status),
		UploadBy:    result.UploadBy,
		ApprovedBy:  result.ApprovedBy,
		ReportType:  result.ReportType.Name,
	}
	return &reportResponse, nil
}
