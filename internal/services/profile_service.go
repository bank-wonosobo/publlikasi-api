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

type ProfileService interface {
	Create(ctx context.Context, req *dto.ProfileCreateReq, file *multipart.FileHeader) (*dto.ProfileResponse, error)
	Index(ctx context.Context, params *dto.ProfileGetQueryParams) ([]dto.ProfileResponse, int64, error)
	Update(ctx context.Context, req *dto.ProfileUpdateReq, file *multipart.FileHeader, id string) (*dto.ProfileResponse, error)
	Delete(ctx context.Context, id string) error
}

type profileService struct {
	profileRepo repositories.ProfileRepository
	s3          storage.S3Storage
}

func NewProfile(profileRepo repositories.ProfileRepository, s3 storage.S3Storage) ProfileService {
	return &profileService{
		profileRepo: profileRepo,
		s3:          s3,
	}
}

func (p *profileService) Create(ctx context.Context, req *dto.ProfileCreateReq, file *multipart.FileHeader) (*dto.ProfileResponse, error) {
	if req.Title != "" {
		_, err := p.profileRepo.FindByTitle(ctx, req.Title)
		if err == nil {
			return nil, errors.New("profile sudah ada")
		}
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}

	var imageUrl *string
	if file != nil {
		uploadedURL, err := p.s3.UploadFileRename(file, "profiles/images/", nil)
		if err != nil {
			return nil, err
		}
		imageUrl = &uploadedURL
	}

	profile := entities.Profile{
		Title:       req.Title,
		Description: req.Description,
		ImageUrl:    imageUrl,
	}

	result, err := p.profileRepo.Save(ctx, &profile)
	if err != nil {
		return nil, err
	}

	response := toProfileResponse(*result)
	return &response, nil
}

func (p *profileService) Index(ctx context.Context, params *dto.ProfileGetQueryParams) ([]dto.ProfileResponse, int64, error) {
	if params.Page < 1 {
		params.Page = 1
	}
	if params.Limit < 1 {
		params.Limit = 10
	}

	offset := (params.Page - 1) * params.Limit

	result, total, err := p.profileRepo.GetAll(ctx, params, offset)
	if err != nil {
		return nil, 0, err
	}

	return toProfileResponses(result), total, nil
}

func (p *profileService) Update(ctx context.Context, req *dto.ProfileUpdateReq, file *multipart.FileHeader, id string) (*dto.ProfileResponse, error) {
	profile, err := p.profileRepo.FindByID(ctx, id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("profile id tidak ditemukan")
	}
	if err != nil {
		return nil, err
	}

	profile.Title = req.Title
	profile.Description = req.Description

	if file != nil {
		uploadedURL, err := p.s3.UploadFileRename(file, "profiles/images/", nil)
		if err != nil {
			return nil, err
		}
		profile.ImageUrl = &uploadedURL
	}

	result, err := p.profileRepo.Save(ctx, profile)
	if err != nil {
		return nil, err
	}

	response := toProfileResponse(*result)
	return &response, nil
}

func (p *profileService) Delete(ctx context.Context, id string) error {
	profile, err := p.profileRepo.FindByID(ctx, id)
	if err != nil {
		return errors.New("profile tidak ditemukan")
	}

	err = p.profileRepo.Delete(ctx, profile)
	if err != nil {
		return err
	}

	return nil
}

func toProfileResponse(profile entities.Profile) dto.ProfileResponse {
	return dto.ProfileResponse{
		ID:          profile.ID,
		Title:       profile.Title,
		Description: profile.Description,
		ImageUrl:    profile.ImageUrl,
		CreatedAt:   profile.CreatedAt,
		UpdatedAt:   profile.UpdatedAt,
	}
}

func toProfileResponses(profiles []entities.Profile) []dto.ProfileResponse {
	var responses []dto.ProfileResponse
	for _, profile := range profiles {
		responses = append(responses, toProfileResponse(profile))
	}

	return responses
}
