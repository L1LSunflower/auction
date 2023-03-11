package aggregates

import (
	"github.com/L1LSunflower/auction/internal/domain/entities"
	"time"
)

type AuctionAggregation struct {
	User      *entities.User
	Auction   *entities.Auction
	Item      *entities.Item
	Tags      []*entities.Tag
	ItemFiles []*entities.File
	Member    bool
}

func (a *AuctionAggregation) CreateItem(name, description string) {
	a.Item = &entities.Item{
		UserID:      a.User.ID,
		Name:        name,
		Description: description,
	}
}

func (a *AuctionAggregation) CreateAuction(startPrice, minPrice float64, category, shortDescription string, startDate time.Time) {
	a.Auction = &entities.Auction{
		Category:         category,
		OwnerID:          a.User.ID,
		ItemID:           a.Item.ID,
		StartedAt:        startDate,
		StartPrice:       startPrice,
		MinPrice:         minPrice,
		ShortDescription: shortDescription,
		Status:           entities.InactiveStatus,
	}
}

type ProfileAggregation struct {
	User     *entities.User
	Balance  *entities.Balance
	Auctions []*AuctionFile
}

type ProfileHistoryAggregation struct {
	User     *entities.User
	Auctions []*AuctionFile
}

type AuctionFile struct {
	Auction *entities.Auction
	Files   []string
}

type AuctionsItem struct {
	AuctsItem []*AuctItem
}

type AuctItem struct {
	ID     int
	Status string
	Files  string
}
