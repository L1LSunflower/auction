package auctions

import (
	"github.com/L1LSunflower/auction/internal/tools/metadata"
	"time"
)

type Create struct {
	OwnerID          string    `json:"owner_id"`
	ShortDescription string    `json:"short_description"`
	StartDate        time.Time `json:"start_date"`
	StartPrice       float64   `json:"start_price"`
	MinimalPrice     float64   `json:"minimal_price"`
	Category         string    `json:"category"`
	ItemTitle        string    `json:"title"`
	ItemDescription  string    `json:"description"`
	ItemFiles        []string  `json:"files"`
	ItemTags         []string  `json:"tags"`
}

type Auction struct {
	ID     int
	UserID string
}

type Auctions struct {
	Where    []string
	Tags     []string
	OrderBy  string
	Metadata *metadata.Metadata
}

type Update struct {
	ID               int
	Title            string    `json:"title"`
	Description      string    `json:"description"`
	ShortDescription string    `json:"short_description"`
	StartPrice       float64   `json:"start_price"`
	MinimalPrice     float64   `json:"minimal_price"`
	StartDate        time.Time `json:"start_date"`
	Category         string    `json:"category"`
	ItemFiles        []string  `json:"files"`
	ItemTags         []string  `json:"tags"`
}

type Start struct {
	ID      int
	EndedAt time.Time `json:"ended_at"`
}

type End struct {
	ID int
}

type Delete struct {
	ID int
}

type Participate struct {
	AuctionID int
	UserID    string
}
