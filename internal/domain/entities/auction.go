package entities

import "time"

const (
	InactiveStatus  = "inactive"
	ActiveStatus    = "active"
	CompletedStatus = "completed"
)

type Auction struct {
	ID               int
	Category         string
	OwnerID          string
	WinnerID         string
	ItemID           int
	ShortDescription string
	StartPrice       float64
	MinPrice         float64
	Status           string
	StartedAt        time.Time
	EndedAt          time.Time
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        time.Time
}
