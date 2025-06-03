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
	Index(ctx context.Context)
	Update(ctx context.Context)
	Delete(ctx context.Context)
	Detail(ctx context.Context)
}

type productService struct {
	productRepo repositories.ProductRepository
	s3          storage.S3Storage
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
func (p *productService) Delete(ctx context.Context) {
	panic("unimplemented")
}

// Detail implements ProductService.
func (p *productService) Detail(ctx context.Context) {
	panic("unimplemented")
}

// Index implements ProductService.
func (p *productService) Index(ctx context.Context) {
	panic("unimplemented")
}

// Update implements ProductService.
func (p *productService) Update(ctx context.Context) {
	panic("unimplemented")
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

// func toProductResponses(news []entities.Product) []dto.NewsResponse {
// 	var newsReponses []dto.NewsResponse
// 	for _, newsItem := range news {
// 		newsReponses = append(newsReponses, toNewsResponse(newsItem))
// 	}

// 	return newsReponses
// }
