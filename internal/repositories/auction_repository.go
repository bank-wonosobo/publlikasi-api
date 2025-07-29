package repositories

import (
	"context"

	"github.com/bank-wonosobo/publlikasi-api.git/internal/delivery/http/dto"
	"github.com/bank-wonosobo/publlikasi-api.git/internal/entities"
	"gorm.io/gorm"
)

type AuctionRepsitory interface {
	GetAll(ctx context.Context, params *dto.AuctionGetQueryParams, offsite int) ([]entities.Auction, int64, error)
	FindByName(ctx context.Context, name string) (*entities.Auction, error)
	Save(ctx context.Context, auction *entities.Auction) (*entities.Auction, error)
	FindByID(ctx context.Context, id string) (*entities.Auction, error)
	Delete(ctx context.Context, auction *entities.Auction) error
}

type auctionRepository struct {
	db *gorm.DB
}

// Delete implements AuctionRepsitory.
func (a *auctionRepository) Delete(ctx context.Context, auction *entities.Auction) error {
	panic("unimplemented")
}

// FindByID implements AuctionRepsitory.
func (a *auctionRepository) FindByID(ctx context.Context, id string) (*entities.Auction, error) {
	panic("unimplemented")
}

// FindByName implements AuctionRepsitory.
func (a *auctionRepository) FindByName(ctx context.Context, name string) (*entities.Auction, error) {
	panic("unimplemented")
}

// GetAll implements AuctionRepsitory.
func (a *auctionRepository) GetAll(ctx context.Context, params *dto.AuctionGetQueryParams, offsite int) ([]entities.Auction, int64, error) {
	panic("unimplemented")
}

// Save implements AuctionRepsitory.
func (a *auctionRepository) Save(ctx context.Context, auction *entities.Auction) (*entities.Auction, error) {
	err := a.db.WithContext(ctx).Save(&auction).Error
	if err != nil {
		return nil, err
	}

	return auction, nil
}

func NewAuction(db *gorm.DB) AuctionRepsitory {
	return &auctionRepository{
		db: db,
	}
}
