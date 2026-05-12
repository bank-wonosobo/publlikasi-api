package services

import (
	"context"
	"errors"
	"mime/multipart"
	"time"

	"github.com/bank-wonosobo/publlikasi-api.git/internal/delivery/http/dto"
	"github.com/bank-wonosobo/publlikasi-api.git/internal/entities"
	"github.com/bank-wonosobo/publlikasi-api.git/internal/repositories"
	"github.com/bank-wonosobo/publlikasi-api.git/pkg/storage"
	"gorm.io/gorm"
)

type ReportService interface {
	Index(ctx context.Context, params *dto.ReportGetQueryParams) ([]dto.ReportResponse, int64, error)
	Create(ctx context.Context, req *dto.ReportCreateRequest, file *multipart.FileHeader) (*dto.ReportResponse, error)
	UploadFile(ctx context.Context, file *multipart.FileHeader, id string) (*dto.ReportResponse, error)
	Update(ctx context.Context, req *dto.ReportUpdateRequest, id string) (*dto.ReportResponse, error)
	Delete(ctx context.Context, id string) error
	Approve(ctx context.Context, id string) (*dto.ReportResponse, error)
	Archive(ctx context.Context, id string) (*dto.ReportResponse, error)
	GetByReportType(ctx context.Context, params *dto.ReportGetQueryParams, reportTypeID int) (*dto.ReportPaginateResponse, error)
}

type reportService struct {
	reportRepo     repositories.ReportRepository
	reportTypeRepo repositories.ReportTypeRepository
	s3             storage.S3Storage
}

func NewReport(reportRepo repositories.ReportRepository,
	reportTypeRepo repositories.ReportTypeRepository,
	s3 storage.S3Storage,
) ReportService {
	return &reportService{
		reportRepo:     reportRepo,
		reportTypeRepo: reportTypeRepo,
		s3:             s3,
	}
}

// Index implements ReportService.
func (r *reportService) Index(ctx context.Context, params *dto.ReportGetQueryParams) ([]dto.ReportResponse, int64, error) {
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
	result, total, err := r.reportRepo.GetAll(ctx, params, offset)
	if err != nil {
		return nil, 0, err
	}

	// return result
	reportResponse := toReportResponses(result)

	return reportResponse, total, nil
}

// Create implements ReportService.
func (r *reportService) Create(ctx context.Context, req *dto.ReportCreateRequest, file *multipart.FileHeader) (*dto.ReportResponse, error) {
	// check if title exist
	_, err := r.reportRepo.FindByTitle(ctx, req.Title)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("judul laporan sudah ada")
	}

	// check report type id
	reportType, err := r.reportTypeRepo.FindByName(ctx, req.ReportType)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("report type tidak ditemukan")
	}
	// upload file
	fileUrl, err := r.s3.UploadFileRename(file, "reports/"+reportType.Name, nil)
	if err != nil {
		return nil, err
	}

	periodStart, err := time.Parse("2006-01-02", req.PeriodStart)
	if err != nil {
		return nil, errors.New("invalid period start format")
	}

	var periodEnd *time.Time
	if req.PeriodEnd != nil {
		t, err := time.Parse("2006-01-02", *req.PeriodEnd)
		if err != nil {
			return nil, errors.New("invalid period end format")
		}
		periodEnd = &t
	}

	report := entities.Report{
		Title:       req.Title,
		Description: req.Description,
		PeriodStart: periodStart,
		PeriodEnd:   periodEnd,
		Year:        req.Year,
		Quarter:     req.Quarter,
		Version:     req.Version,
		ReportType:  *reportType,
		Status:      entities.Draft,
		UploadBy:    "user",
		FileUrl:     &fileUrl,
	}
	result, err := r.reportRepo.Save(ctx, &report)
	if err != nil {
		return nil, err
	}

	// return result
	reportResponse := toReportResponse(*result)
	return &reportResponse, nil
}

// UploadFile implements ReportService.
func (r *reportService) UploadFile(ctx context.Context, file *multipart.FileHeader, id string) (*dto.ReportResponse, error) {
	// check report id
	report, err := r.reportRepo.FindByID(ctx, id)
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
	result, err := r.reportRepo.Save(ctx, report)
	if err != nil {
		return nil, err
	}

	// return result
	reportResponse := toReportResponse(*result)
	return &reportResponse, nil
}

// Update implements ReportService.
func (r *reportService) Update(ctx context.Context, req *dto.ReportUpdateRequest, id string) (*dto.ReportResponse, error) {
	// get report type by id
	report, err := r.reportRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// check report type id
	reportType, err := r.reportTypeRepo.FindByName(ctx, req.ReportType)
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

	result, err := r.reportRepo.Save(ctx, report)
	if err != nil {
		return nil, err
	}

	// return result
	reportResponse := toReportResponse(*result)
	return &reportResponse, nil
}

// Delete implements ReportService.
func (r *reportService) Delete(ctx context.Context, id string) error {
	// get report type by id
	report, err := r.reportRepo.FindByID(ctx, id)
	if err != nil {
		return errors.New("report tidak ditemukan")
	}

	err = r.reportRepo.Delete(ctx, report)
	if err != nil {
		return err
	}

	return nil
}

// Archive implements ReportService.
func (r *reportService) Archive(ctx context.Context, id string) (*dto.ReportResponse, error) {
	// get report type by id
	report, err := r.reportRepo.FindByID(ctx, id)
	if err != nil {
		return nil, errors.New("report tidak ditemukan")
	}

	// update report
	report.Status = entities.Archived
	result, err := r.reportRepo.Save(ctx, report)
	if err != nil {
		return nil, err
	}

	// return result
	reportResponse := toReportResponse(*result)
	return &reportResponse, nil
}

// Upprove implements ReportService.
func (r *reportService) Approve(ctx context.Context, id string) (*dto.ReportResponse, error) {
	// get report type by id
	report, err := r.reportRepo.FindByID(ctx, id)
	if err != nil {
		return nil, errors.New("report tidak ditemukan")
	}

	// update report
	userApprover := "user approver"
	report.Status = entities.Published
	report.ApprovedBy = &userApprover
	result, err := r.reportRepo.Save(ctx, report)
	if err != nil {
		return nil, err
	}

	// return result
	reportResponse := toReportResponse(*result)
	return &reportResponse, nil
}

// GetByReportType implements ReportService.
func (r *reportService) GetByReportType(ctx context.Context, params *dto.ReportGetQueryParams, reportTypeID int) (*dto.ReportPaginateResponse, error) {
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
	result, total, err := r.reportRepo.GetByReportType(ctx, params, reportTypeID, offset)
	if err != nil {
		return nil, err
	}

	// return result
	reportResponse := toReportResponses(result)

	reportPaginateResponse := dto.ReportPaginateResponse{
		Reports:   reportResponse,
		Page:      params.Page,
		Limit:     params.Limit,
		Total:     total,
		TotalPage: (total + int64(params.Limit) - 1) / int64(params.Limit),
	}

	return &reportPaginateResponse, nil
}

func toReportResponse(report entities.Report) (result dto.ReportResponse) {
	result = dto.ReportResponse{
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
		FileUrl:     report.FileUrl,
	}

	return result
}

func toReportResponses(reports []entities.Report) []dto.ReportResponse {
	var reportResponse []dto.ReportResponse
	for _, report := range reports {
		reportResponse = append(reportResponse, toReportResponse(report))
	}

	return reportResponse
}
