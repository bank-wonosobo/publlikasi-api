package services

import (
	"context"
	"errors"
	"mime/multipart"

	"github.com/bank-wonosobo/publlikasi-api.git/internal/delivery/http/dto/request"
	"github.com/bank-wonosobo/publlikasi-api.git/internal/delivery/http/dto/response"
	"github.com/bank-wonosobo/publlikasi-api.git/internal/entities"
	"github.com/bank-wonosobo/publlikasi-api.git/internal/repositories"
	"github.com/bank-wonosobo/publlikasi-api.git/pkg/storage"
	"gorm.io/gorm"
)

type ReportService interface {
	Index(ctx context.Context, params *request.ReportGetQueryParams) (*response.ReportPaginateResponse, error)
	Create(ctx context.Context, req *request.ReportCreateRequest) (*response.ReportResponse, error)
	UploadFile(ctx context.Context, file *multipart.FileHeader, id string) (*response.ReportResponse, error)
	Update(ctx context.Context, req *request.ReportUpdateRequest, id string) (*response.ReportResponse, error)
	Delete(ctx context.Context, id string) error
	Upprove(ctx context.Context, id string) (*response.ReportResponse, error)
	Archive(ctx context.Context, id string) (*response.ReportResponse, error)
}

type reportService struct {
	db             *gorm.DB
	reportRepo     repositories.ReportRepository
	reportTypeRepo repositories.ReportTypeRepository
	s3             storage.S3Storage
}

func NewReport(db *gorm.DB,
	reportRepo repositories.ReportRepository,
	reportTypeRepo repositories.ReportTypeRepository,
	s3 storage.S3Storage,
) ReportService {
	return &reportService{
		db:             db,
		reportRepo:     reportRepo,
		reportTypeRepo: reportTypeRepo,
		s3:             s3,
	}
}

// Index implements ReportService.
func (r *reportService) Index(ctx context.Context, params *request.ReportGetQueryParams) (*response.ReportPaginateResponse, error) {
	// Set default values
	if params.Page < 1 {
		params.Page = 1
	}
	if params.Limit < 1 {
		params.Limit = 10
	}

	// set offset
	offset := (params.Page - 1) * params.Limit

	// get all data with total
	result, total, err := r.reportRepo.GetAll(ctx, r.db, params, offset)
	if err != nil {
		return nil, err
	}

	// return result
	var reportResponse []response.ReportResponse
	for _, report := range result {
		reportResponse = append(reportResponse, response.ReportResponse{
			ID:          report.ID,
			Title:       report.Title,
			Description: report.Description,
			PeriodStart: report.PeriodStart,
			PeriodEnd:   report.PeriodEnd,
			Year:        report.Year,
			Quarter:     report.Quarter,
			Version:     report.Version,
			Status:      string(report.Status),
			UploadBy:    report.UploadBy,
			ApprovedBy:  report.ApprovedBy,
			ReportType:  report.ReportType.Name,
		})
	}

	reportPaginateResponse := response.ReportPaginateResponse{
		Reports:   reportResponse,
		Page:      params.Page,
		Limit:     params.Limit,
		Total:     total,
		TotalPage: (total + int64(params.Limit) - 1) / int64(params.Limit),
	}

	return &reportPaginateResponse, nil

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
		Status:      entities.Draft,
		UploadBy:    "user",
	}
	result, err := r.reportRepo.Save(ctx, r.db, &report)
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
		Version:     result.Version,
		Status:      string(result.Status),
		UploadBy:    result.UploadBy,
		ApprovedBy:  result.ApprovedBy,
		ReportType:  result.ReportType.Name,
	}
	return &reportResponse, nil
}

// UploadFile implements ReportService.
func (r *reportService) UploadFile(ctx context.Context, file *multipart.FileHeader, id string) (*response.ReportResponse, error) {
	// check report id
	report, err := r.reportRepo.FindByID(ctx, r.db, id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("report tidak ditemukan")
	}

	if report.FileUrl != nil {
		return nil, errors.New("file sudah ada")
	}

	// upload file
	fileUrl, err := r.s3.UploadFileRename(file, "reports/"+report.ReportType.Name, nil)
	if err != nil {
		return nil, err
	}

	// update report
	report.FileUrl = &fileUrl
	result, err := r.reportRepo.Update(ctx, r.db, report)
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
		Version:     result.Version,
		Status:      string(result.Status),
		FileUrl:     fileUrl,
		UploadBy:    result.UploadBy,
		ApprovedBy:  result.ApprovedBy,
		ReportType:  result.ReportType.Name,
	}
	return &reportResponse, nil
}

// Update implements ReportService.
func (r *reportService) Update(ctx context.Context, req *request.ReportUpdateRequest, id string) (*response.ReportResponse, error) {
	// get report type by id
	report, err := r.reportRepo.FindByID(ctx, r.db, id)
	if err != nil {
		return nil, err
	}

	// check report type id
	reportType, err := r.reportTypeRepo.FindByName(ctx, r.db, req.ReportType)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("report type tidak ditemukan")
	}

	// update report type
	report.Title = req.Title
	report.Description = req.Description
	report.PeriodStart = req.PeriodStart
	report.PeriodEnd = req.PeriodEnd
	report.Quarter = req.Quarter
	report.Year = req.Year
	report.Version = req.Version
	report.ReportType = *reportType
	report.Status = entities.Status(req.Status)

	result, err := r.reportRepo.Update(ctx, r.db, report)
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
		Version:     result.Version,
		Status:      string(result.Status),
		UploadBy:    result.UploadBy,
		ApprovedBy:  result.ApprovedBy,
		ReportType:  result.ReportType.Name,
	}

	return &reportResponse, nil
}

// Delete implements ReportService.
func (r *reportService) Delete(ctx context.Context, id string) error {
	// get report type by id
	report, err := r.reportRepo.FindByID(ctx, r.db, id)
	if err != nil {
		return errors.New("report tidak ditemukan")
	}

	err = r.reportRepo.Delete(ctx, r.db, report)
	if err != nil {
		return err
	}

	return nil
}

// Archive implements ReportService.
func (r *reportService) Archive(ctx context.Context, id string) (*response.ReportResponse, error) {
	panic("unimplemented")
}

// Upprove implements ReportService.
func (r *reportService) Upprove(ctx context.Context, id string) (*response.ReportResponse, error) {
	panic("unimplemented")
}
