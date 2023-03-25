package auctions

import (
	"github.com/L1LSunflower/auction/internal/tools/metadata"
	"time"
)

type Create struct {
	OwnerID          string    `json:"owner_id" validate:"required"`
	ShortDescription string    `json:"short_description" validate:"omitempty"`
	StartDate        time.Time `json:"start_date" validate:"required"`
	StartPrice       float64   `json:"start_price" validate:"required"`
	MinimalPrice     float64   `json:"minimal_price" validate:"required"`
	Category         string    `json:"category" validate:"required"`
	ItemTitle        string    `json:"title" validate:"required"`
	ItemDescription  string    `json:"description" validate:"omitempty"`
	ItemFiles        []string  `json:"files" validate:"omitempty"`
	ItemTags         []string  `json:"tags" validate:"omitempty"`
}

type Auction struct {
	ID     int    `validate:"required"`
	UserID string `validate:"required"`
}

type Auctions struct {
	Where    []string
	Tags     string
	OrderBy  string
	Metadata *metadata.Metadata
}

type Update struct {
	ID               int       `validate:"required"`
	Title            string    `json:"title" validate:"required"`
	Description      string    `json:"description" validate:"required"`
	ShortDescription string    `json:"short_description" validate:"required"`
	StartPrice       float64   `json:"start_price" validate:"required"`
	MinimalPrice     float64   `json:"minimal_price" validate:"required"`
	StartDate        time.Time `json:"start_date" validate:"required"`
	Category         string    `json:"category" validate:"required"`
	ItemFiles        []string  `json:"files" validate:"required"`
	ItemTags         []string  `json:"tags" validate:"required"`
}

type Start struct {
	ID      int       `validate:"required"`
	UserID  string    `validate:"required"`
	EndedAt time.Time `json:"end_date" validate:"required"`
}

type End struct {
	ID     int    `validate:"required"`
	UserID string `validate:"required"`
}

type Delete struct {
	ID int `validate:"required"`
}

type Participate struct {
	AuctionID int    `validate:"required"`
	UserID    string `validate:"required"`
}

type SetPrice struct {
	UserID    string  `validate:"required"`
	AuctionID int     `validate:"required"`
	Price     float64 `json:"price" validate:"required"`
}
