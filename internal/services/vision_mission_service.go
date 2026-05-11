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

type VisionMissionService interface {
	Create(ctx context.Context, req *dto.VisionMissionCreateReq, file *multipart.FileHeader) (*dto.VisionMissionResponse, error)
	Index(ctx context.Context, params *dto.VisionMissionGetQueryParams) ([]dto.VisionMissionResponse, int64, error)
	Update(ctx context.Context, req *dto.VisionMissionUpdateReq, file *multipart.FileHeader, id string) (*dto.VisionMissionResponse, error)
	Delete(ctx context.Context, id string) error
}

type visionMissionService struct {
	repo repositories.VisionMissionRepository
	s3   storage.S3Storage
}

func NewVisionMission(repo repositories.VisionMissionRepository, s3 storage.S3Storage) VisionMissionService {
	return &visionMissionService{
		repo: repo,
		s3:   s3,
	}
}

func (v *visionMissionService) Create(ctx context.Context, req *dto.VisionMissionCreateReq, file *multipart.FileHeader) (*dto.VisionMissionResponse, error) {
	_, err := v.repo.FindByTitle(ctx, req.Title)
	if err == nil {
		return nil, errors.New("visi misi sudah ada")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	var imageUrl *string
	if file != nil {
		uploadedURL, err := v.s3.UploadFileRename(file, "vision-missions/images/", nil)
		if err != nil {
			return nil, err
		}
		imageUrl = &uploadedURL
	}

	item := entities.VisionMission{
		Title:    req.Title,
		Vision:   req.Vision,
		Mission:  req.Mission,
		ImageUrl: imageUrl,
	}

	result, err := v.repo.Save(ctx, &item)
	if err != nil {
		return nil, err
	}

	response := toVisionMissionResponse(*result)
	return &response, nil
}

func (v *visionMissionService) Index(ctx context.Context, params *dto.VisionMissionGetQueryParams) ([]dto.VisionMissionResponse, int64, error) {
	if params.Page < 1 {
		params.Page = 1
	}
	if params.Limit < 1 {
		params.Limit = 10
	}

	offset := (params.Page - 1) * params.Limit

	result, total, err := v.repo.GetAll(ctx, params, offset)
	if err != nil {
		return nil, 0, err
	}

	return toVisionMissionResponses(result), total, nil
}

func (v *visionMissionService) Update(ctx context.Context, req *dto.VisionMissionUpdateReq, file *multipart.FileHeader, id string) (*dto.VisionMissionResponse, error) {
	item, err := v.repo.FindByID(ctx, id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("visi misi id tidak ditemukan")
	}
	if err != nil {
		return nil, err
	}

	item.Title = req.Title
	item.Vision = req.Vision
	item.Mission = req.Mission

	if file != nil {
		uploadedURL, err := v.s3.UploadFileRename(file, "vision-missions/images/", nil)
		if err != nil {
			return nil, err
		}
		item.ImageUrl = &uploadedURL
	}

	result, err := v.repo.Save(ctx, item)
	if err != nil {
		return nil, err
	}

	response := toVisionMissionResponse(*result)
	return &response, nil
}

func (v *visionMissionService) Delete(ctx context.Context, id string) error {
	item, err := v.repo.FindByID(ctx, id)
	if err != nil {
		return errors.New("visi misi tidak ditemukan")
	}

	err = v.repo.Delete(ctx, item)
	if err != nil {
		return err
	}

	return nil
}

func toVisionMissionResponse(item entities.VisionMission) dto.VisionMissionResponse {
	return dto.VisionMissionResponse{
		ID:        item.ID,
		Title:     item.Title,
		Vision:    item.Vision,
		Mission:   item.Mission,
		ImageUrl:  item.ImageUrl,
		CreatedAt: item.CreatedAt,
		UpdatedAt: item.UpdatedAt,
	}
}

func toVisionMissionResponses(items []entities.VisionMission) []dto.VisionMissionResponse {
	var responses []dto.VisionMissionResponse
	for _, item := range items {
		responses = append(responses, toVisionMissionResponse(item))
	}

	return responses
}
