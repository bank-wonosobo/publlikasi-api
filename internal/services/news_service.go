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
	Index(ctx context.Context, params *request.NewsGetQueryParams) ([]response.NewsResponse, int64, error)
	Update(ctx context.Context, request *request.NewsUpdateRequest, file *multipart.FileHeader, id string) (*response.NewsResponse, error)
	Delete(ctx context.Context, id string) error
	Approve(ctx context.Context, id string) (*response.NewsResponse, error)
	Archive(ctx context.Context, id string) (*response.NewsResponse, error)
	Detail(ctx context.Context, id string) (*response.NewsResponse, error)
	DetailBySlug(ctx context.Context, slug string) (*response.NewsResponse, error)
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
	result, err := n.newsRepo.Save(ctx, n.db, &report)
	if err != nil {
		return nil, err
	}

	// return result
	newsReponse := toNewsResponse(*result)
	return &newsReponse, nil
}

// Index implements NewsService.
func (n *newsService) Index(ctx context.Context, params *request.NewsGetQueryParams) ([]response.NewsResponse, int64, error) {
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
	result, total, err := n.newsRepo.GetAll(ctx, n.db, params, offset)
	if err != nil {
		return nil, 0, err
	}

	// return result
	newsResponses := toNewsResponses(result)

	return newsResponses, total, nil
}

// Update implements NewsService.
func (n *newsService) Update(ctx context.Context, request *request.NewsUpdateRequest, file *multipart.FileHeader, id string) (*response.NewsResponse, error) {
	// check news type id
	news, err := n.newsRepo.FindByID(ctx, n.db, id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("news id tidak ditemukan")
	}

	// update report type
	news.Title = request.Title
	news.Author = request.Author
	news.Content = request.Content
	news.Slug = request.Slug
	if file != nil {
		// upload file
		imageUrl, err := n.s3.UploadFileRename(file, "news/images/", nil)
		if err != nil {
			return nil, err
		}
		news.ImageUrl = imageUrl
	}

	result, err := n.newsRepo.Update(ctx, n.db, news)
	if err != nil {
		return nil, err
	}

	// return result
	newsResponse := toNewsResponse(*result)
	return &newsResponse, nil
}

// Approve implements NewsService.
func (n *newsService) Approve(ctx context.Context, id string) (*response.NewsResponse, error) {
	// get news type by id
	news, err := n.newsRepo.FindByID(ctx, n.db, id)
	if err != nil {
		return nil, errors.New("berita tidak ditemukan")
	}

	// update news
	news.Status = entities.Published
	result, err := n.newsRepo.Update(ctx, n.db, news)
	if err != nil {
		return nil, err
	}

	// return result
	reportResponse := toNewsResponse(*result)
	return &reportResponse, nil
}

// Archive implements NewsService.
func (n *newsService) Archive(ctx context.Context, id string) (*response.NewsResponse, error) {
	// get news type by id
	news, err := n.newsRepo.FindByID(ctx, n.db, id)
	if err != nil {
		return nil, errors.New("news tidak ditemukan")
	}

	// update news
	news.Status = entities.Archived
	result, err := n.newsRepo.Update(ctx, n.db, news)
	if err != nil {
		return nil, err
	}

	// return result
	reportResponse := toNewsResponse(*result)
	return &reportResponse, nil
}

// Delete implements NewsService.
func (n *newsService) Delete(ctx context.Context, id string) error {
	// get news  by id
	news, err := n.newsRepo.FindByID(ctx, n.db, id)
	if err != nil {
		return errors.New("news tidak ditemukan")
	}

	err = n.newsRepo.Delete(ctx, n.db, news)
	if err != nil {
		return err
	}

	return nil
}

// Detail implements NewsService.
func (n *newsService) Detail(ctx context.Context, id string) (*response.NewsResponse, error) {
	panic("unimplemented")
}

// DetailBySlug implements NewsService.
func (n *newsService) DetailBySlug(ctx context.Context, slug string) (*response.NewsResponse, error) {
	panic("unimplemented")
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

func toNewsResponses(news []entities.News) []response.NewsResponse {
	var newsReponses []response.NewsResponse
	for _, newsItem := range news {
		newsReponses = append(newsReponses, toNewsResponse(newsItem))
	}

	return newsReponses
}
