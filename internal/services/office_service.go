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
}

type officeService struct {
	officeRepo repositories.OfficeRepository
	s3         storage.S3Storage
}

// Create implements OfficeService.
func (o *officeService) Create(ctx context.Context, req *dto.OfficeCreateReq, file *multipart.FileHeader) (*dto.OfficeResponse, error) {
	// check if title exist
	_, err := o.officeRepo.FindByName(ctx, req.Name)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("berita sudah ada")
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
