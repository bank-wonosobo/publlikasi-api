package repositories

import (
	"context"

	"github.com/bank-wonosobo/publlikasi-api.git/internal/delivery/http/dto"
	"github.com/bank-wonosobo/publlikasi-api.git/internal/entities"
	"gorm.io/gorm"
)

type AuctionRepsitory interface {
	GetAll(ctx context.Context, params *dto.AuctionGetQueryParams, offsite int) ([]entities.Auction, int64, error)
	Save(ctx context.Context, auction *entities.Auction) (*entities.Auction, error)
	FindByID(ctx context.Context, id string) (*entities.Auction, error)
	Delete(ctx context.Context, auction *entities.Auction) error
}

type auctionRepository struct {
	db *gorm.DB
}

func NewAuction(db *gorm.DB) AuctionRepsitory {
	return &auctionRepository{
		db: db,
	}
}

// Delete implements AuctionRepsitory.
func (a *auctionRepository) Delete(ctx context.Context, auction *entities.Auction) error {
	err := a.db.WithContext(ctx).Delete(&auction).Error
	if err != nil {
		return err
	}

	return nil
}

// FindByID implements AuctionRepsitory.
func (a *auctionRepository) FindByID(ctx context.Context, id string) (*entities.Auction, error) {
	var auction entities.Auction
	err := a.db.WithContext(ctx).Where("id = ?", id).First(&auction).Error
	if err != nil {
		return nil, err
	}

	return &auction, nil
}

// GetAll implements AuctionRepsitory.
func (a *auctionRepository) GetAll(ctx context.Context, params *dto.AuctionGetQueryParams, offsite int) ([]entities.Auction, int64, error) {
	var total int64

	query := a.db.WithContext(ctx).Model(&entities.Auction{})

	if params.Key != "" {
		query = query.Where("title ILIKE ?", "%"+params.Key+"%").Or("description ILIKE ?", "%"+params.Key+"%")
	}

	query.Count(&total)

	var auctions []entities.Auction
	err := query.Offset(offsite).Limit(params.Limit).Find(&auctions).Error
	if err != nil {
		return nil, 0, err
	}

	return auctions, total, nil
}

// Save implements AuctionRepsitory.
func (a *auctionRepository) Save(ctx context.Context, auction *entities.Auction) (*entities.Auction, error) {
	err := a.db.WithContext(ctx).Save(&auction).Error
	if err != nil {
		return nil, err
	}

	return auction, nil
}
