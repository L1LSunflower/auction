package entities

import "time"

type City struct {
	ID        int       `db:"id"`
	Name      string    `db:"name"`
	IsActive  int       `db:"is_active"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"created_at"`
	DeletedAt time.Time `db:"deleted_at"`
}
