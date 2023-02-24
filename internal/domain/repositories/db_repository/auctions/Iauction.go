package auctions

import (
	"context"

	"github.com/L1LSunflower/auction/internal/domain/entities"
	"github.com/L1LSunflower/auction/internal/tools/metadata"
)

type AuctionInterface interface {
	Create(ctx context.Context, auction *entities.Auction) error
	Auction(ctx context.Context, id int) (*entities.Auction, error)
	Auctions(ctx context.Context, where []string, metadata *metadata.Metadata) ([]*entities.Auction, error)
	ByOwnerID(ctx context.Context, ownerID int) (*entities.Auction, error)
	Update(ctx context.Context, auction *entities.Auction) error
	Start(ctx context.Context, auction *entities.Auction) error
	End(ctx context.Context, auction *entities.Auction) error
	Delete(ctx context.Context, auction *entities.Auction) error
}

func NewRepository() AuctionInterface {
	return &Repository{}
}
