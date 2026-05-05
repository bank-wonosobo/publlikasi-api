package repositories

import (
	"context"

	"github.com/bank-wonosobo/publlikasi-api.git/internal/delivery/http/dto"
	"github.com/bank-wonosobo/publlikasi-api.git/internal/entities"
	"gorm.io/gorm"
)

type OrganizationalStructureRepository interface {
	Save(ctx context.Context, item *entities.OrganizationalStructure) (*entities.OrganizationalStructure, error)
	GetAll(ctx context.Context, params *dto.OrganizationalStructureGetQueryParams, offset int) ([]entities.OrganizationalStructure, int64, error)
	Delete(ctx context.Context, item *entities.OrganizationalStructure) error
	FindByID(ctx context.Context, id string) (*entities.OrganizationalStructure, error)
	FindByTitle(ctx context.Context, title string) (*entities.OrganizationalStructure, error)
}

type organizationalStructureRepository struct {
	db *gorm.DB
}

func NewOrganizationalStructure(db *gorm.DB) OrganizationalStructureRepository {
	return &organizationalStructureRepository{db: db}
}

func (o *organizationalStructureRepository) Save(ctx context.Context, item *entities.OrganizationalStructure) (*entities.OrganizationalStructure, error) {
	err := o.db.WithContext(ctx).Save(&item).Error
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (o *organizationalStructureRepository) GetAll(ctx context.Context, params *dto.OrganizationalStructureGetQueryParams, offset int) (result []entities.OrganizationalStructure, total int64, err error) {
	query := o.db.WithContext(ctx).Model(&entities.OrganizationalStructure{})

	if params.Key != "" {
		query = query.Where("title ILIKE ?", "%"+params.Key+"%").Or("description ILIKE ?", "%"+params.Key+"%")
	}

	query.Count(&total)

	err = query.Limit(params.Limit).Offset(offset).Find(&result).Error
	if err != nil {
		return nil, 0, err
	}

	return result, total, nil
}

func (o *organizationalStructureRepository) Delete(ctx context.Context, item *entities.OrganizationalStructure) error {
	err := o.db.WithContext(ctx).Delete(&item).Error
	if err != nil {
		return err
	}

	return nil
}

func (o *organizationalStructureRepository) FindByID(ctx context.Context, id string) (result *entities.OrganizationalStructure, err error) {
	err = o.db.WithContext(ctx).Where("id = ?", id).First(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (o *organizationalStructureRepository) FindByTitle(ctx context.Context, title string) (result *entities.OrganizationalStructure, err error) {
	err = o.db.WithContext(ctx).Where("title = ?", title).First(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}
