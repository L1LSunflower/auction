package entities

import "time"

type Item struct {
	ID          int
	UserID      int
	Category    string
	Name        string
	Tag1        string
	Tag2        string
	Tag3        string
	Tag4        string
	Tag5        string
	Tag6        string
	Tag7        string
	Tag8        string
	Tag9        string
	Tag10       string
	Images      string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}
