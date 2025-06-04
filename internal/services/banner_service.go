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
}

type bannerService struct {
	bannerRepo repositories.BannerRepository
	s3         storage.S3Storage
}

// Create implements BannerService.
func (b *bannerService) Create(ctx context.Context, req *dto.BannerCreateRequest, file *multipart.FileHeader) (*dto.BannerResponse, error) {
	// check if title exist
	_, err := b.bannerRepo.FindByName(ctx, req.Name)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("banner sudah ada")
	}

	// upload file
	imageUrl, err := b.s3.UploadFileRename(file, "products/images/", nil)
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

func toBannerResponses(news []entities.Product) []dto.ProductResponse {
	var newsReponses []dto.ProductResponse
	for _, newsItem := range news {
		newsReponses = append(newsReponses, toProductResponse(newsItem))
	}

	return newsReponses
}
