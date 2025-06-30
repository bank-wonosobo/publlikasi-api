package services

import (
	"context"
	"errors"

	"github.com/bank-wonosobo/publlikasi-api.git/internal/delivery/http/dto"
	"github.com/bank-wonosobo/publlikasi-api.git/internal/entities"
	"github.com/bank-wonosobo/publlikasi-api.git/internal/repositories"
	"gorm.io/gorm"
)

type ComplaintTypeService interface {
	Index(ctx context.Context) ([]dto.ComplaintTypeResponse, error)
	Create(ctx context.Context, req *dto.ComplaintTypeCreateRequest) (*dto.ComplaintTypeResponse, error)
	Update(ctx context.Context, req *dto.ComplaintTypeCreateRequest, id int) (*dto.ComplaintTypeResponse, error)
	Delete(ctx context.Context, id int) error
}

type complaintTypeService struct {
	repository repositories.ComplaintTypeRepository
}

// Create implements ComplaintTypeService.
func (c *complaintTypeService) Create(ctx context.Context, req *dto.ComplaintTypeCreateRequest) (*dto.ComplaintTypeResponse, error) {
	// check if name exist
	_, err := c.repository.FindByName(ctx, req.Name)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("jenis laporan sudah ada")
	}
	// create complaint type
	complaintType := entities.ComplaintType{
		Name:        req.Name,
		Description: req.Description,
	}
	result, err := c.repository.Save(ctx, &complaintType)
	if err != nil {
		return nil, err
	}

	// return result
	complaintTypeResponse := toComplaintTypeResponse(*result)

	return &complaintTypeResponse, nil
}

// Delete implements ComplaintTypeService.
func (c *complaintTypeService) Delete(ctx context.Context, id int) error {
	// get report type by id
	complaintType, err := c.repository.FindByID(ctx, id)
	if err != nil {
		return errors.New("complaint type tidak ditemukan")
	}

	err = c.repository.Delete(ctx, complaintType)
	if err != nil {
		return err
	}

	return nil
}

// Index implements ComplaintTypeService.
func (c *complaintTypeService) Index(ctx context.Context) ([]dto.ComplaintTypeResponse, error) {
	// get all
	result, err := c.repository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	// return result
	complaintTypeResponse := toComplaintTypeResponses(result)

	return complaintTypeResponse, nil
}

// Update implements ComplaintTypeService.
func (c *complaintTypeService) Update(ctx context.Context, req *dto.ComplaintTypeCreateRequest, id int) (*dto.ComplaintTypeResponse, error) {
	// get report type by id
	complaintType, err := c.repository.FindByID(ctx, id)
	if err != nil {
		return nil, errors.New("complaint type tidak ditemukan")
	}

	// update report type
	complaintType.Name = req.Name
	complaintType.Description = req.Description
	result, err := c.repository.Save(ctx, complaintType)
	if err != nil {
		return nil, err
	}

	// return result
	complaintTypeResponse := toComplaintTypeResponse(*result)

	return &complaintTypeResponse, nil
}

func NewComplaintType(repository repositories.ComplaintTypeRepository) ComplaintTypeService {
	return &complaintTypeService{
		repository: repository,
	}
}

func toComplaintTypeResponse(reportType entities.ComplaintType) (result dto.ComplaintTypeResponse) {
	result = dto.ComplaintTypeResponse{
		ID:          reportType.ID,
		Name:        reportType.Name,
		Description: reportType.Description,
		CreatedAt:   reportType.CreatedAt.String(),
		UpdatedAt:   reportType.UpdatedAt.String(),
	}

	return result
}

func toComplaintTypeResponses(reportTypes []entities.ComplaintType) []dto.ComplaintTypeResponse {
	var complaintTypeResponse []dto.ComplaintTypeResponse
	for _, reportType := range reportTypes {
		complaintTypeResponse = append(complaintTypeResponse, toComplaintTypeResponse(reportType))
	}

	return complaintTypeResponse
}
