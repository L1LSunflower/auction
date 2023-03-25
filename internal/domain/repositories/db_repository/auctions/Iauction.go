package auctions

import (
	"context"
	"time"

	"github.com/L1LSunflower/auction/internal/domain/entities"
	"github.com/L1LSunflower/auction/internal/tools/metadata"
)

type AuctionInterface interface {
	Create(ctx context.Context, auction *entities.Auction) error
	Auction(ctx context.Context, id int) (*entities.Auction, error)
	Auctions(ctx context.Context, where, tags, groupBy string, metadata *metadata.Metadata) ([]*entities.Auction, error)
	ByOwnerID(ctx context.Context, ownerID string) ([]*entities.Auction, error)
	Update(ctx context.Context, auction *entities.Auction) error
	Start(ctx context.Context, id int, endedDate time.Time) error
	End(ctx context.Context, id int) error
	Delete(ctx context.Context, auction *entities.Auction) error
	Count(ctx context.Context) (int, error)
	ActiveAuction(ctx context.Context, ownerID string) (*entities.Auction, error)
	ActiveAuctions(ctx context.Context, ownerID string) ([]*entities.Auction, error)
	CountInactiveAuctions(ctx context.Context, ownerID string) (int, error)
	DeleteByOwnerID(ctx context.Context, ownerID string) error
	UpdatePrice(ctx context.Context, id int, price float64) error
	CreateMember(ctx context.Context, auctionID int, userID string) (*entities.AuctionMember, error)
	Member(ctx context.Context, auctionID int, userID string) (*entities.AuctionMember, error)
	ActivateAuctions(ctx context.Context) error
	EndAuctions(ctx context.Context) error
	SetPrice(ctx context.Context, auctionID int, price float64) error
}

func NewRepository() AuctionInterface {
	return &Repository{}
}
