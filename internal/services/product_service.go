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

type ProductService interface {
	Create(ctx context.Context, req *dto.ProductCreateReq, file *multipart.FileHeader) (*dto.ProductResponse, error)
	Index(ctx context.Context, params *dto.ProductGetQueryParams) ([]dto.ProductResponse, int64, error)
	Update(ctx context.Context, req *dto.ProductUpdateReq, file *multipart.FileHeader, id string) (*dto.ProductResponse, error)
	Delete(ctx context.Context, id string) error
	Detail(ctx context.Context, id string) (*dto.ProductResponse, error)
}

type productService struct {
	productRepo repositories.ProductRepository
	s3          storage.S3Storage
}

// Update implements ProductService.
func (p *productService) Update(ctx context.Context, req *dto.ProductUpdateReq, file *multipart.FileHeader, id string) (*dto.ProductResponse, error) {
	// check news type id
	product, err := p.productRepo.FindByID(ctx, id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("news id tidak ditemukan")
	}

	// update report type
	product.Name = req.Name
	product.Description = req.Description
	product.Tagline = req.Tagline
	product.ProductCategory = req.ProductCategory
	if file != nil {
		// upload file
		imageUrl, err := p.s3.UploadFileRename(file, "news/images/", nil)
		if err != nil {
			return nil, err
		}
		product.ImageUrl = imageUrl
	}

	result, err := p.productRepo.Save(ctx, product)
	if err != nil {
		return nil, err
	}

	// return result
	productResponse := toProductResponse(*result)
	return &productResponse, nil
}

// Index implements ProductService.
func (p *productService) Index(ctx context.Context, params *dto.ProductGetQueryParams) ([]dto.ProductResponse, int64, error) {
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
	result, total, err := p.productRepo.GetAll(ctx, params, offset)
	if err != nil {
		return nil, 0, err
	}

	// return result
	productResponse := toProductResponses(result)

	return productResponse, total, nil
}

// Create implements ProductService.
func (p *productService) Create(ctx context.Context, req *dto.ProductCreateReq, file *multipart.FileHeader) (*dto.ProductResponse, error) {
	// check if title exist
	_, err := p.productRepo.FindByName(ctx, req.Name)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("product sudah ada")
	}

	// upload file
	imageUrl, err := p.s3.UploadFileRename(file, "products/images/", nil)
	if err != nil {
		return nil, err
	}

	// create product type
	product := entities.Product{
		Name:            req.Name,
		Description:     req.Description,
		Tagline:         req.Tagline,
		ProductCategory: req.ProductCategory,
		ImageUrl:        imageUrl,
	}

	result, err := p.productRepo.Save(ctx, &product)
	if err != nil {
		return nil, err
	}

	// return result
	productResponse := toProductResponse(*result)
	return &productResponse, nil
}

// Delete implements ProductService.
func (p *productService) Delete(ctx context.Context, id string) error {
	// get product by id
	product, err := p.productRepo.FindByID(ctx, id)
	if err != nil {
		return errors.New("product tidak ditemukan")
	}

	err = p.productRepo.Delete(ctx, product)
	if err != nil {
		return err
	}

	return nil
}

// Detail implements ProductService.
func (p *productService) Detail(ctx context.Context, id string) (*dto.ProductResponse, error) {
	product, err := p.productRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// return result
	productReponse := toProductResponse(*product)

	return &productReponse, nil
}

func NewProduct(productRepo repositories.ProductRepository, s3 storage.S3Storage) ProductService {
	return &productService{
		productRepo: productRepo,
		s3:          s3,
	}
}

func toProductResponse(product entities.Product) (result dto.ProductResponse) {
	result = dto.ProductResponse{
		ID:              product.ID,
		Name:            product.Name,
		Description:     product.Description,
		Tagline:         product.Tagline,
		ProductCategory: string(product.ProductCategory),
		ImageUrl:        product.ImageUrl,
		CreatedAt:       product.CreatedAt,
		UpdatedAt:       product.UpdatedAt,
	}

	return result
}

func toProductResponses(news []entities.Product) []dto.ProductResponse {
	var newsReponses []dto.ProductResponse
	for _, newsItem := range news {
		newsReponses = append(newsReponses, toProductResponse(newsItem))
	}

	return newsReponses
}
