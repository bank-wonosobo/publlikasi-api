package services

import (
	"context"
	"errors"
	"mime/multipart"

	"github.com/bank-wonosobo/publlikasi-api.git/internal/delivery/http/dto/request"
	"github.com/bank-wonosobo/publlikasi-api.git/internal/delivery/http/dto/response"
	"github.com/bank-wonosobo/publlikasi-api.git/internal/entities"
	"github.com/bank-wonosobo/publlikasi-api.git/internal/repositories"
	"github.com/bank-wonosobo/publlikasi-api.git/pkg/storage"
	"gorm.io/gorm"
)

type NewsService interface {
	Create(ctx context.Context, request *request.NewsCreateRequest, file *multipart.FileHeader) (*response.NewsResponse, error)
}

type newsService struct {
	db       *gorm.DB
	newsRepo repositories.NewsRepository
	s3       storage.S3Storage
}

func NewNews(db *gorm.DB, newsRepo repositories.NewsRepository, s3 storage.S3Storage) NewsService {
	return &newsService{
		db:       db,
		newsRepo: newsRepo,
		s3:       s3,
	}
}

// Create implements NewsService.
func (n *newsService) Create(ctx context.Context, req *request.NewsCreateRequest, file *multipart.FileHeader) (*response.NewsResponse, error) {
	// check if title exist
	_, err := n.newsRepo.FindByTitle(ctx, n.db, req.Title)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("berita sudah ada")
	}

	// upload file
	imageUrl, err := n.s3.UploadFileRename(file, "news/images/", nil)
	if err != nil {
		return nil, err
	}

	// create report type
	report := entities.News{
		Title:    req.Title,
		Slug:     req.Slug,
		Content:  req.Content,
		Author:   "user",
		Status:   entities.Draft,
		ImageUrl: imageUrl,
	}
	result, err := n.newsRepo.Save(ctx, n.db, report)
	if err != nil {
		return nil, err
	}

	// return result
	newsReponse := toNewsResponse(*result)
	return &newsReponse, nil
}

func toNewsResponse(news entities.News) (result response.NewsResponse) {
	result = response.NewsResponse{
		ID:          news.ID,
		Title:       news.Title,
		Slug:        news.Slug,
		Content:     news.Content,
		Author:      news.Author,
		ImageUrl:    news.ImageUrl,
		PublishedAt: news.PublishedAt,
		Status:      string(news.Status),
		ApprovedBy:  news.ApprovedBy,
		CreatedAt:   news.CreatedAt,
		UpdatedAt:   news.UpdatedAt,
	}

	return result
}
