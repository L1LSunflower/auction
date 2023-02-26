package aggregates

import "github.com/L1LSunflower/auction/internal/domain/entities"

type AuctionItem struct {
	Auction *entities.Auction
	Item    *entities.Item
}

type AuctionsItem struct {
	AuctsItem []*AuctItem `json:"auctions"`
}

type AuctItem struct {
	ID     int
	Status string
	Files  string
}

func NewAuctionItem() *AuctionItem {
	return &AuctionItem{}
}
