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

type BannerService interface {
	Create(ctx context.Context, req *dto.BannerCreateRequest, file *multipart.FileHeader) (*dto.BannerResponse, error)
	Index(ctx context.Context, params *dto.BannerGetQueryParams) ([]dto.BannerResponse, int64, error)
	Update(ctx context.Context, request *dto.BannerUpdateRequest, file *multipart.FileHeader, id string) (*dto.BannerResponse, error)
}

type bannerService struct {
	bannerRepo repositories.BannerRepository
	s3         storage.S3Storage
}

// Update implements BannerService.
func (b *bannerService) Update(ctx context.Context, request *dto.BannerUpdateRequest, file *multipart.FileHeader, id string) (*dto.BannerResponse, error) {
	// check news type id
	product, err := b.bannerRepo.FindByID(ctx, id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("banner id tidak ditemukan")
	}

	// update report type
	product.Name = request.Name
	product.Description = request.Description
	if file != nil {
		// upload file
		imageUrl, err := b.s3.UploadFileRename(file, "banner/images/", nil)
		if err != nil {
			return nil, err
		}
		product.ImageUrl = imageUrl
	}

	result, err := b.bannerRepo.Save(ctx, product)
	if err != nil {
		return nil, err
	}

	// return result
	bannerResponse := toBannerResponse(*result)
	return &bannerResponse, nil
}

// Index implements BannerService.
func (b *bannerService) Index(ctx context.Context, params *dto.BannerGetQueryParams) ([]dto.BannerResponse, int64, error) {
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
	result, total, err := b.bannerRepo.GetAll(ctx, params, offset)
	if err != nil {
		return nil, 0, err
	}

	// return result
	productResponse := toBannerResponses(result)

	return productResponse, total, nil
}

// Create implements BannerService.
func (b *bannerService) Create(ctx context.Context, req *dto.BannerCreateRequest, file *multipart.FileHeader) (*dto.BannerResponse, error) {
	// check if title exist
	_, err := b.bannerRepo.FindByName(ctx, req.Name)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("banner sudah ada")
	}

	// upload file
	imageUrl, err := b.s3.UploadFileRename(file, "banner/images/", nil)
	if err != nil {
		return nil, err
	}

	// create product type
	product := entities.Banner{
		Name:        req.Name,
		Description: req.Description,
		IsActive:    false,
		ImageUrl:    imageUrl,
	}

	result, err := b.bannerRepo.Save(ctx, &product)
	if err != nil {
		return nil, err
	}

	// return result
	bannerResponse := toBannerResponse(*result)
	return &bannerResponse, nil
}

func NewBanner(bannerRepo repositories.BannerRepository, s3 storage.S3Storage) BannerService {
	return &bannerService{
		bannerRepo: bannerRepo,
		s3:         s3,
	}
}

func toBannerResponse(banner entities.Banner) (result dto.BannerResponse) {
	result = dto.BannerResponse{
		ID:          banner.ID,
		Name:        banner.Name,
		Description: banner.Description,
		ImageUrl:    banner.ImageUrl,
		CreatedAt:   banner.CreatedAt,
		UpdatedAt:   banner.UpdatedAt,
	}

	return result
}

func toBannerResponses(banners []entities.Banner) []dto.BannerResponse {
	var bannerReponses []dto.BannerResponse
	for _, bannerItem := range banners {
		bannerReponses = append(bannerReponses, toBannerResponse(bannerItem))
	}

	return bannerReponses
}
