package entities

import "time"

type Balance struct {
	ID        string
	Balance   float64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
