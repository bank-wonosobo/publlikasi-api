package services

import (
	"context"
	"mime/multipart"

	"github.com/bank-wonosobo/publlikasi-api.git/internal/delivery/http/dto"
	"github.com/bank-wonosobo/publlikasi-api.git/internal/entities"
	"github.com/bank-wonosobo/publlikasi-api.git/internal/repositories"
	"github.com/bank-wonosobo/publlikasi-api.git/pkg/storage"
)

type AuctionService interface {
	Create(ctx context.Context, req *dto.AuctionCreateRequest, file *multipart.FileHeader) (*dto.AuctionResponse, error)
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
		StartTime:   req.StartTime,
		EndTime:     req.EndTime,
		ExternalUrl: req.ExternalUrl,
	}

	result, err := a.auctionRepo.Save(ctx, &auction)
	if err != nil {
		return nil, err
	}

	// return result
	auctionResponse := toAuctionResponse(*result)
	return &auctionResponse, nil
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
		ExternalUrl: auction.ExternalUrl,
		CreatedAt:   auction.CreatedAt,
		StartTime:   auction.StartTime,
		EndTime:     auction.EndTime,
		UpdatedAt:   auction.UpdatedAt,
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
