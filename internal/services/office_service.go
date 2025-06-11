package services

import (
	"context"
	"errors"
	"mime/multipart"

	"github.com/bank-wonosobo/publlikasi-api.git/internal/delivery/http/dto"
	"github.com/bank-wonosobo/publlikasi-api.git/internal/entities"
	"github.com/bank-wonosobo/publlikasi-api.git/internal/repositories"
	"github.com/bank-wonosobo/publlikasi-api.git/pkg/storage"
	"gorm.io/gorm"
)

type OfficeService interface {
	Create(ctx context.Context, req *dto.OfficeCreateReq, file *multipart.FileHeader) (*dto.OfficeResponse, error)
	Index(ctx context.Context, params *dto.OfficeGetQueryParams) ([]dto.OfficeResponse, int64, error)
	Update(ctx context.Context, request *dto.OfficeUpdateReq, file *multipart.FileHeader, id string) (*dto.OfficeResponse, error)
	Delete(ctx context.Context, id string) error
}

type officeService struct {
	officeRepo repositories.OfficeRepository
	s3         storage.S3Storage
}

// Delete implements OfficeService.
func (o *officeService) Delete(ctx context.Context, id string) error {
	// get news  by id
	news, err := o.officeRepo.FindByID(ctx, id)
	if err != nil {
		return errors.New("news tidak ditemukan")
	}

	err = o.officeRepo.Delete(ctx, news)
	if err != nil {
		return err
	}

	return nil
}

// Update implements OfficeService.
func (o *officeService) Update(ctx context.Context, request *dto.OfficeUpdateReq, file *multipart.FileHeader, id string) (*dto.OfficeResponse, error) {
	// check news type id
	office, err := o.officeRepo.FindByID(ctx, id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("office id tidak ditemukan")
	}

	// update report type
	office.Address = request.Address
	office.Latitude = request.Latitude
	office.Longitude = request.Longitude
	office.Name = request.Name
	office.PhoneNumber = request.PhoneNumber
	office.MapLink = request.MapLink
	if file != nil {
		// upload file
		imageUrl, err := o.s3.UploadFileRename(file, "office/images/", nil)
		if err != nil {
			return nil, err
		}
		office.ImageUrl = imageUrl
	}

	result, err := o.officeRepo.Save(ctx, office)
	if err != nil {
		return nil, err
	}

	// return result
	officeResponse := toOfficeResponse(*result)
	return &officeResponse, nil
}

// Index implements OfficeService.
func (o *officeService) Index(ctx context.Context, params *dto.OfficeGetQueryParams) ([]dto.OfficeResponse, int64, error) {
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
	result, total, err := o.officeRepo.GetAll(ctx, params, offset)
	if err != nil {
		return nil, 0, err
	}

	// return result
	newsResponses := toOfficeResponses(result)

	return newsResponses, total, nil
}

// Create implements OfficeService.
func (o *officeService) Create(ctx context.Context, req *dto.OfficeCreateReq, file *multipart.FileHeader) (*dto.OfficeResponse, error) {
	// check if title exist
	_, err := o.officeRepo.FindByName(ctx, req.Name)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("kantor sudah ada")
	}

	// upload file
	imageUrl, err := o.s3.UploadFileRename(file, "office/images/", nil)
	if err != nil {
		return nil, err
	}

	// create report type
	office := entities.Office{
		Name:        req.Name,
		Address:     req.Address,
		Latitude:    req.Latitude,
		Longitude:   req.Longitude,
		MapLink:     req.MapLink,
		PhoneNumber: req.PhoneNumber,
		ImageUrl:    imageUrl,
	}
	result, err := o.officeRepo.Save(ctx, &office)
	if err != nil {
		return nil, err
	}

	// return result
	officeReponse := toOfficeResponse(*result)
	return &officeReponse, nil
}

func NewOffice(officeRepo repositories.OfficeRepository,
	s3 storage.S3Storage) OfficeService {
	return &officeService{
		officeRepo: officeRepo,
		s3:         s3,
	}
}

func toOfficeResponse(office entities.Office) dto.OfficeResponse {
	result := dto.OfficeResponse{
		ID:          office.ID,
		Name:        office.Name,
		Address:     office.Address,
		Latitude:    office.Latitude,
		Longitude:   office.Longitude,
		ImageUrl:    office.ImageUrl,
		MapLink:     office.MapLink,
		PhoneNumber: office.PhoneNumber,
		CreatedAt:   office.CreatedAt,
		UpdateAt:    office.UpdatedAt,
	}

	return result
}

func toOfficeResponses(offices []entities.Office) []dto.OfficeResponse {
	var officeResponses []dto.OfficeResponse
	for _, office := range offices {
		officeResponses = append(officeResponses, toOfficeResponse(office))
	}

	return officeResponses
}
