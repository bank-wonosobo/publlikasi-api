package services

import (
	"context"

	"github.com/bank-wonosobo/publlikasi-api.git/internal/delivery/http/dto"
	"github.com/bank-wonosobo/publlikasi-api.git/internal/entities"
	"github.com/bank-wonosobo/publlikasi-api.git/internal/repositories"
	"gorm.io/gorm"
)

type AnnouncementService interface {
	Index(ctx context.Context, param *dto.AnnouncementGetQueryParams) ([]dto.AnnouncementResponse, int64, error)
}

type announcementService struct {
	db               *gorm.DB
	announcementRepo repositories.AnnouncementRepository
}

func NewAnnouncement(db *gorm.DB,
	announcementRepo repositories.AnnouncementRepository) AnnouncementService {
	return &announcementService{
		db:               db,
		announcementRepo: announcementRepo,
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
