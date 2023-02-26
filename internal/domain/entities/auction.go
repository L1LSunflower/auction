package entities

import "time"

type Auction struct {
	ID          int
	OwnerID     int
	WinnerID    string
	ItemID      int
	Title       string
	Description string
	StartPrice  float64
	MinPrice    float64
	Status      string
	StartedAt   time.Time
	EndedAt     time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}
