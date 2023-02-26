package auctions

import "github.com/L1LSunflower/auction/internal/tools/metadata"

type Create struct {
	OwnerID            int     `json:"owner_id"`
	Category           string  `json:"category"`
	Name               string  `json:"name"`
	ItemID             int     `json:"item_id"`
	Title              string  `json:"title"`
	Description        string  `json:"description"`
	StartPrice         float64 `json:"start_price"`
	MinimalPrice       float64 `json:"minimal_price"`
	Tag1               string  `json:"tag1"`
	Tag2               string  `json:"tag2"`
	Tag3               string  `json:"tag3"`
	Tag4               string  `json:"tag4"`
	Tag5               string  `json:"tag5"`
	Tag6               string  `json:"tag6"`
	Tag7               string  `json:"tag7"`
	Tag8               string  `json:"tag8"`
	Tag9               string  `json:"tag9"`
	Tag10              string  `json:"tag10"`
	Images             string  `json:"files"`
	ItemDescription    string  `json:"item_description"`
	AuctionDescription string  `json:"auction_description"`
	Status             string  `json:"status"`
}

type Auction struct {
	ID int
}

type Auctions struct {
	Where    []string
	Metadata *metadata.Metadata
}

type Update struct {
	ID                 int
	Tag1               string  `json:"tag1"`
	Tag2               string  `json:"tag2"`
	Tag3               string  `json:"tag3"`
	Tag4               string  `json:"tag4"`
	Tag5               string  `json:"tag5"`
	Tag6               string  `json:"tag6"`
	Tag7               string  `json:"tag7"`
	Tag8               string  `json:"tag8"`
	Tag9               string  `json:"tag9"`
	Tag10              string  `json:"tag10"`
	Images             string  `json:"files"`
	ItemDescription    string  `json:"item_description"`
	WinnerID           string  `json:"winner_id"`
	Title              string  `json:"title"`
	AuctionDescription string  `json:"auction_description"`
	StartPrice         float64 `json:"start_price"`
	MinimalPrice       float64 `json:"minimal_price"`
	Status             string  `json:"status"`
}

type Start struct {
	ID int
}

type End struct {
	ID int
}

type Delete struct {
	ID int
}
