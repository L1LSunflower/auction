package structs

type CreateAuction struct {
	Status string `json:"status"`
	ID     int    `json:"id"`
}

type Auction struct {
	Status           string   `json:"status"`
	Member           bool     `json:"member"`
	ID               int      `json:"id,required"`
	AuctionStatus    string   `json:"auction_status"`
	Phone            string   `json:"phone"`
	Category         string   `json:"category,required"`
	WinnerID         string   `json:"winner_id,omitempty"`
	Title            string   `json:"title,required"`
	ShortDescription string   `json:"short_description,required"`
	Description      string   `json:"description,omitempty"`
	StartPrice       float64  `json:"start_price,omitempty"`
	MinimalPrice     float64  `json:"minimal_price,omitempty"`
	StartDate        string   `json:"start_date,omitempty"`
	EndedAt          string   `json:"end_date,omitempty"`
	Files            []string `json:"files,omitempty"`
	Tags             []string `json:"tags,omitempty"`
}

type AuctionsWithFile struct {
	Status      string            `json:"status,required"`
	CurrentPage int               `json:"current_page"`
	Total       int               `json:"total"`
	LastPage    int               `json:"last_page"`
	Auctions    []AuctionWithFile `json:"auctions,required"`
}

type AuctionWithFile struct {
	Status           string `json:"status,required"`
	ID               int    `json:"id,required"`
	AuctionStatus    string `json:"auction_status"`
	ShortDescription string `json:"short_description,required"`
	Files            string `json:"media,required"`
	Category         string `json:"category"`
}

type Update struct {
	Status           string   `json:"status,required"`
	ID               int      `json:"id,required"`
	Category         string   `json:"category,required"`
	WinnerID         string   `json:"winner_id,omitempty"`
	Title            string   `json:"title,required"`
	ShortDescription string   `json:"short_description,required"`
	Description      string   `json:"description,omitempty"`
	AuctionStatus    string   `json:"auction_status"`
	StartPrice       float64  `json:"start_price,omitempty"`
	MinimalPrice     float64  `json:"minimal_price,omitempty"`
	StartDate        string   `json:"start_date,omitempty"`
	EndedAt          string   `json:"ended_at,omitempty"`
	Files            []string `json:"files,omitempty"`
}

type Delete struct {
	Status string `json:"status"`
	Date   string `json:"date"`
}

type Start struct {
	Status string `json:"status"`
	Date   string `json:"date"`
}

type End struct {
	Status   string `json:"status"`
	Date     string `json:"date"`
	WinnerID string `json:"winner_id"`
}

type Participate struct {
	Status string `json:"status"`
	Date   string `json:"date"`
}

type SetPrice struct {
	Status string  `json:"status,required"`
	Price  float64 `json:"price,required"`
	Date   string  `json:"date"`
}
