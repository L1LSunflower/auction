package entities

import "time"

const (
	InactiveStatus  = "inactive"
	ActiveStatus    = "active"
	CompletedStatus = "completed"
	DateFormat      = "2006-01-02 15:04:05"
	// Auction visit status
	VisitNotSet = "not set"
	VisitSet    = "set"
	VisitOpened = "opened"
	VisitClosed = "closed"
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
	Price            float64
	StartedAt        time.Time
	EndedAt          time.Time
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        time.Time
	VisitStatus      string
	VisitStartDate   time.Time
	VisitEndDate     time.Time
}

type AuctionMember struct {
	AuctionID     int
	ParticipantID string
	Price         float64
	FirstName     string
	LastName      string
}

type AuctionVisit struct {
	AuctionID int
	StartDate time.Time
	EndDate   time.Time
}
