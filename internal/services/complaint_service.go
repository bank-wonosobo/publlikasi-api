package services

import (
	"context"
	"errors"
	"mime/multipart"
	"strconv"
	"time"

	"github.com/bank-wonosobo/publlikasi-api.git/internal/delivery/http/dto"
	"github.com/bank-wonosobo/publlikasi-api.git/internal/entities"
	"github.com/bank-wonosobo/publlikasi-api.git/internal/repositories"
	"github.com/bank-wonosobo/publlikasi-api.git/pkg/storage"
	"gorm.io/gorm"
)

type ComplaintService interface {
	Index(ctx context.Context, params *dto.ComplaintGetQueryParams) ([]dto.ComplaintResponse, int64, error)
	Create(ctx context.Context, req *dto.ComplaintCreateRequest, file *multipart.FileHeader) (*dto.ComplaintResponse, error)
	Process(ctx context.Context, complaintID string) (*dto.ComplaintResponse, error)
	Done(ctx context.Context, complaintID string) (*dto.ComplaintResponse, error)
}

type complaintService struct {
	complaintRepo     repositories.ComplaintRepository
	complaintTypeRepo repositories.ComplaintTypeRepository
	s3                storage.S3Storage
}

// Done implements ComplaintService.
func (c *complaintService) Done(ctx context.Context, complaintID string) (*dto.ComplaintResponse, error) {
	// get news type by id
	complaint, err := c.complaintRepo.FindByID(ctx, complaintID)
	if err != nil {
		return nil, errors.New("complaint tidak ditemukan")
	}

	// update complaint
	complaint.Status = entities.Done
	result, err := c.complaintRepo.Save(ctx, complaint)
	if err != nil {
		return nil, err
	}

	// return result
	complaintResponse := toComplaintResponse(*result)
	return &complaintResponse, nil
}

// Process implements ComplaintService.
func (c *complaintService) Process(ctx context.Context, complaintID string) (*dto.ComplaintResponse, error) {
	// get news type by id
	complaint, err := c.complaintRepo.FindByID(ctx, complaintID)
	if err != nil {
		return nil, errors.New("complaint tidak ditemukan")
	}

	// update complaint
	complaint.Status = entities.Process
	result, err := c.complaintRepo.Save(ctx, complaint)
	if err != nil {
		return nil, err
	}

	// return result
	complaintResponse := toComplaintResponse(*result)
	return &complaintResponse, nil
}

// Index implements ComplaintService.
func (c *complaintService) Index(ctx context.Context, params *dto.ComplaintGetQueryParams) ([]dto.ComplaintResponse, int64, error) {
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
	result, total, err := c.complaintRepo.GetAll(ctx, params, offset)
	if err != nil {
		return nil, 0, err
	}

	// return result
	complaintResponse := toComplaintResponses(result)

	return complaintResponse, total, nil
}

// Create implements ComplaintService.
func (c *complaintService) Create(ctx context.Context, req *dto.ComplaintCreateRequest, file *multipart.FileHeader) (*dto.ComplaintResponse, error) {

	// check report type id
	complaintType, err := c.complaintTypeRepo.FindByName(ctx, req.ComplaintType)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("complaint type tidak ditemukan")
	}
	// upload file
	evidenceUrl, err := c.s3.UploadFileRename(file, "complaint/evidence/", nil)
	if err != nil {
		return nil, err
	}

	// create report type
	report := entities.Complaint{
		ComplaintID:      generateComplaintID(),
		EvidenceUrl:      evidenceUrl,
		ComplaintType:    *complaintType,
		ReportedName:     req.ReportedName,
		InsidentLocation: req.InsidentLocation,
		InsidentTime:     req.InsidentTime,
		Description:      req.Description,
		ReporterName:     req.ReportedName,
		ReporterEmail:    req.ReporterEmail,
		ReporterPhone:    req.ReporterPhone,
		Status:           entities.Pending,
	}
	result, err := c.complaintRepo.Save(ctx, &report)
	if err != nil {
		return nil, err
	}

	// return result
	complaintResponse := toComplaintResponse(*result)
	return &complaintResponse, nil
}

func NewComplaint(complaintRepo repositories.ComplaintRepository,
	s3 storage.S3Storage,
	complaintTypeRepo repositories.ComplaintTypeRepository) ComplaintService {
	return &complaintService{
		complaintRepo:     complaintRepo,
		s3:                s3,
		complaintTypeRepo: complaintTypeRepo,
	}
}

func toComplaintResponse(complaint entities.Complaint) (result dto.ComplaintResponse) {
	result = dto.ComplaintResponse{
		ID:               complaint.ID,
		ComplaintID:      complaint.ComplaintID,
		ReportedName:     complaint.ReportedName,
		Email:            complaint.ReporterEmail,
		InsidentLocation: complaint.InsidentLocation,
		InsidentTime:     complaint.InsidentTime,
		Description:      complaint.Description,
		EvidenceUrl:      complaint.EvidenceUrl,
		ReporterName:     complaint.ReporterName,
		ReporterPhone:    complaint.ReporterPhone,
		Status:           string(complaint.Status),
		ComplaintType:    complaint.ComplaintType.Name,
	}

	return result
}

func toComplaintResponses(complaints []entities.Complaint) []dto.ComplaintResponse {
	var complaintResponse []dto.ComplaintResponse
	for _, item := range complaints {
		complaintResponse = append(complaintResponse, toComplaintResponse(item))
	}

	return complaintResponse
}

func generateComplaintID() string {
	uniqueNumber := time.Now().UnixNano()
	uniqueNumberStr := strconv.FormatInt(uniqueNumber, 10)
	return "PGD" + uniqueNumberStr
}
