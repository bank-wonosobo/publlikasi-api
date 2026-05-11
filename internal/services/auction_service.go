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

type AuctionService interface {
	Create(ctx context.Context, req *dto.AuctionCreateRequest, file *multipart.FileHeader) (*dto.AuctionResponse, error)
	Index(ctx context.Context, params *dto.AuctionGetQueryParams) ([]dto.AuctionResponse, int64, error)
	Update(ctx context.Context, req *dto.AuctionUpdateRequest, file *multipart.FileHeader, id string) (*dto.AuctionResponse, error)
	Delete(ctx context.Context, id string) error
}

type auctionService struct {
	auctionRepo repositories.AuctionRepsitory
	s3          storage.S3Storage
}

// Create implements AuctionService.
func (a *auctionService) Create(ctx context.Context, req *dto.AuctionCreateRequest, file *multipart.FileHeader) (*dto.AuctionResponse, error) {

	// upload file
	imageUrl, err := a.s3.UploadFileRename(file, "auction/images/", nil)
	if err != nil {
		return nil, err
	}

	// create product type
	auction := entities.Auction{
		Title:       req.Title,
		Description: req.Description,
		ImageUrl:    imageUrl,
		Link:        req.Link,
	}

	result, err := a.auctionRepo.Save(ctx, &auction)
	if err != nil {
		return nil, err
	}

	// return result
	auctionResponse := toAuctionResponse(*result)
	return &auctionResponse, nil
}

// Index implements AuctionService.
func (a *auctionService) Index(ctx context.Context, params *dto.AuctionGetQueryParams) ([]dto.AuctionResponse, int64, error) {
	if params.Page < 1 {
		params.Page = 1
	}
	if params.Limit < 1 {
		params.Limit = 10
	}

	offset := (params.Page - 1) * params.Limit

	result, total, err := a.auctionRepo.GetAll(ctx, params, offset)
	if err != nil {
		return nil, 0, err
	}

	return toAuctionResponses(result), total, nil
}

// Update implements AuctionService.
func (a *auctionService) Update(ctx context.Context, req *dto.AuctionUpdateRequest, file *multipart.FileHeader, id string) (*dto.AuctionResponse, error) {
	auction, err := a.auctionRepo.FindByID(ctx, id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("lelang tidak ditemukan")
	}
	if err != nil {
		return nil, err
	}

	auction.Title = req.Title
	auction.Description = req.Description
	auction.Link = req.Link

	if file != nil {
		imageUrl, err := a.s3.UploadFileRename(file, "auction/images/", nil)
		if err != nil {
			return nil, err
		}
		auction.ImageUrl = imageUrl
	}

	result, err := a.auctionRepo.Save(ctx, auction)
	if err != nil {
		return nil, err
	}

	auctionResponse := toAuctionResponse(*result)
	return &auctionResponse, nil
}

// Delete implements AuctionService.
func (a *auctionService) Delete(ctx context.Context, id string) error {
	auction, err := a.auctionRepo.FindByID(ctx, id)
	if err != nil {
		return errors.New("lelang tidak ditemukan")
	}

	return a.auctionRepo.Delete(ctx, auction)
}

func NewAuction(auctionRepo repositories.AuctionRepsitory, s3 storage.S3Storage) AuctionService {
	return &auctionService{
		auctionRepo: auctionRepo,
		s3:          s3,
	}
}

func toAuctionResponse(auction entities.Auction) (result dto.AuctionResponse) {
	result = dto.AuctionResponse{
		ID:          auction.ID,
		Title:       auction.Title,
		Description: auction.Description,
		ImageUrl:    auction.ImageUrl,
		Link:        auction.Link,
	}

	return result
}

func toAuctionResponses(banners []entities.Auction) []dto.AuctionResponse {
	var auctionReponses []dto.AuctionResponse
	for _, auctionItem := range banners {
		auctionReponses = append(auctionReponses, toAuctionResponse(auctionItem))
	}

	return auctionReponses
}
