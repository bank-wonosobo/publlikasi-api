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

type OrganizationalStructureService interface {
	Create(ctx context.Context, req *dto.OrganizationalStructureCreateReq, file *multipart.FileHeader) (*dto.OrganizationalStructureResponse, error)
	Index(ctx context.Context, params *dto.OrganizationalStructureGetQueryParams) ([]dto.OrganizationalStructureResponse, int64, error)
	Update(ctx context.Context, req *dto.OrganizationalStructureUpdateReq, file *multipart.FileHeader, id string) (*dto.OrganizationalStructureResponse, error)
	Delete(ctx context.Context, id string) error
}

type organizationalStructureService struct {
	repo repositories.OrganizationalStructureRepository
	s3   storage.S3Storage
}

func NewOrganizationalStructure(repo repositories.OrganizationalStructureRepository, s3 storage.S3Storage) OrganizationalStructureService {
	return &organizationalStructureService{
		repo: repo,
		s3:   s3,
	}
}

func (o *organizationalStructureService) Create(ctx context.Context, req *dto.OrganizationalStructureCreateReq, file *multipart.FileHeader) (*dto.OrganizationalStructureResponse, error) {
	_, err := o.repo.FindByTitle(ctx, req.Title)
	if err == nil {
		return nil, errors.New("struktur organisasi sudah ada")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	var imageUrl *string
	if file != nil {
		uploadedURL, err := o.s3.UploadFileRename(file, "organizational-structures/images/", nil)
		if err != nil {
			return nil, err
		}
		imageUrl = &uploadedURL
	}

	item := entities.OrganizationalStructure{
		Title:       req.Title,
		Description: req.Description,
		ImageUrl:    imageUrl,
	}

	result, err := o.repo.Save(ctx, &item)
	if err != nil {
		return nil, err
	}

	response := toOrganizationalStructureResponse(*result)
	return &response, nil
}

func (o *organizationalStructureService) Index(ctx context.Context, params *dto.OrganizationalStructureGetQueryParams) ([]dto.OrganizationalStructureResponse, int64, error) {
	if params.Page < 1 {
		params.Page = 1
	}
	if params.Limit < 1 {
		params.Limit = 10
	}

	offset := (params.Page - 1) * params.Limit

	result, total, err := o.repo.GetAll(ctx, params, offset)
	if err != nil {
		return nil, 0, err
	}

	return toOrganizationalStructureResponses(result), total, nil
}

func (o *organizationalStructureService) Update(ctx context.Context, req *dto.OrganizationalStructureUpdateReq, file *multipart.FileHeader, id string) (*dto.OrganizationalStructureResponse, error) {
	item, err := o.repo.FindByID(ctx, id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("struktur organisasi id tidak ditemukan")
	}
	if err != nil {
		return nil, err
	}

	item.Title = req.Title
	item.Description = req.Description

	if file != nil {
		uploadedURL, err := o.s3.UploadFileRename(file, "organizational-structures/images/", nil)
		if err != nil {
			return nil, err
		}
		item.ImageUrl = &uploadedURL
	}

	result, err := o.repo.Save(ctx, item)
	if err != nil {
		return nil, err
	}

	response := toOrganizationalStructureResponse(*result)
	return &response, nil
}

func (o *organizationalStructureService) Delete(ctx context.Context, id string) error {
	item, err := o.repo.FindByID(ctx, id)
	if err != nil {
		return errors.New("struktur organisasi tidak ditemukan")
	}

	err = o.repo.Delete(ctx, item)
	if err != nil {
		return err
	}

	return nil
}

func toOrganizationalStructureResponse(item entities.OrganizationalStructure) dto.OrganizationalStructureResponse {
	return dto.OrganizationalStructureResponse{
		ID:          item.ID,
		Title:       item.Title,
		Description: item.Description,
		ImageUrl:    item.ImageUrl,
		CreatedAt:   item.CreatedAt,
		UpdatedAt:   item.UpdatedAt,
	}
}

func toOrganizationalStructureResponses(items []entities.OrganizationalStructure) []dto.OrganizationalStructureResponse {
	var responses []dto.OrganizationalStructureResponse
	for _, item := range items {
		responses = append(responses, toOrganizationalStructureResponse(item))
	}

	return responses
}
