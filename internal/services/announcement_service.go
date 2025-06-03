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

type AnnouncementService interface {
	Index(ctx context.Context, param *dto.AnnouncementGetQueryParams) ([]dto.AnnouncementResponse, int64, error)
	Create(ctx context.Context, request *dto.AnnouncementCreateReq, file *multipart.FileHeader) (*dto.AnnouncementResponse, error)
	Update(ctx context.Context, request *dto.AnnouncementUpdateReq, file *multipart.FileHeader, id string) (*dto.AnnouncementResponse, error)
	Delete(ctx context.Context, id string) error
}

type announcementService struct {
	db               *gorm.DB
	announcementRepo repositories.AnnouncementRepository
	s3               storage.S3Storage
}

func NewAnnouncement(db *gorm.DB,
	announcementRepo repositories.AnnouncementRepository,
	s3 storage.S3Storage) AnnouncementService {
	return &announcementService{
		db:               db,
		announcementRepo: announcementRepo,
		s3:               s3,
	}
}

// Index implements AnnouncementService.
func (a *announcementService) Index(ctx context.Context, params *dto.AnnouncementGetQueryParams) ([]dto.AnnouncementResponse, int64, error) {
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
	result, total, err := a.announcementRepo.GetAll(ctx, a.db, params, offset)
	if err != nil {
		return nil, 0, err
	}

	// return result
	announcementRes := toAnnouncementResponses(result)

	return announcementRes, total, nil
}

// Create implements AnnouncementService.
func (a *announcementService) Create(ctx context.Context, request *dto.AnnouncementCreateReq, file *multipart.FileHeader) (*dto.AnnouncementResponse, error) {
	// check if title exist
	_, err := a.announcementRepo.FindByTitle(ctx, a.db, request.Title)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("pengumuman sudah ada")
	}

	// upload file
	attachmentUrl, err := a.s3.UploadFileRename(file, "announcements/attachment/", nil)
	if err != nil {
		return nil, err
	}

	// create report type
	report := entities.Announcement{
		Title:          request.Title,
		Content:        request.Content,
		Author:         "user login",
		StartDate:      request.StartDate,
		EndDate:        request.EndDate,
		TargetAudience: entities.TargetAudience(request.TargetAudience),
		AttachmentUrl:  &attachmentUrl,
		IsActive:       true,
		Status:         entities.Draft,
	}
	result, err := a.announcementRepo.Save(ctx, a.db, &report)
	if err != nil {
		return nil, err
	}

	// return result
	announcementRes := toAnnouncementResponse(*result)
	return &announcementRes, nil
}

// Update implements AnnouncementService.
func (a *announcementService) Update(ctx context.Context, request *dto.AnnouncementUpdateReq, file *multipart.FileHeader, id string) (*dto.AnnouncementResponse, error) {
	// check news type id
	announcement, err := a.announcementRepo.FindByID(ctx, a.db, id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("announcement id tidak ditemukan")
	}

	// update report type
	announcement.Title = request.Title
	announcement.Author = "user login edited"
	announcement.Content = request.Content
	announcement.StartDate = request.StartDate
	announcement.EndDate = request.EndDate

	if file != nil {
		// upload file
		attachmentUrl, err := a.s3.UploadFileRename(file, "announcements/attachment/", nil)
		if err != nil {
			return nil, err
		}
		announcement.AttachmentUrl = &attachmentUrl
	}

	result, err := a.announcementRepo.Update(ctx, a.db, announcement)
	if err != nil {
		return nil, err
	}

	// return result
	announcementResult := toAnnouncementResponse(*result)
	return &announcementResult, nil
}

// Delete implements AnnouncementService.
func (a *announcementService) Delete(ctx context.Context, id string) error {
	// get news  by id
	news, err := a.announcementRepo.FindByID(ctx, a.db, id)
	if err != nil {
		return errors.New("pengumuman tidak ditemukan")
	}

	err = a.announcementRepo.Delete(ctx, a.db, news)
	if err != nil {
		return err
	}

	return nil
}

func toAnnouncementResponse(a entities.Announcement) dto.AnnouncementResponse {
	return dto.AnnouncementResponse{
		ID:             a.ID,
		Title:          a.Title,
		Content:        a.Content,
		Author:         a.Author,
		TargetAudience: string(a.TargetAudience),
		StartDate:      a.StartDate,
		EndDate:        a.EndDate,
		AttachmentUrl:  a.AttachmentUrl,
		IsActive:       a.IsActive,
		Status:         string(a.Status),
	}
}

func toAnnouncementResponses(a []entities.Announcement) []dto.AnnouncementResponse {
	var announcementsRes []dto.AnnouncementResponse
	for _, item := range a {
		announcementsRes = append(announcementsRes, toAnnouncementResponse(item))
	}

	return announcementsRes
}
