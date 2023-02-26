package entities

import "time"

type User struct {
	ID        string    `db:"id,required"`
	Email     string    `db:"email,required"`
	FirstName string    `db:"first_name,required"`
	LastName  string    `db:"last_name,required"`
	Phone     string    `db:"phone,required"`
	City      int       `db:"cities,omitempty"`
	Password  string    `db:"password,required"`
	IsActive  int       `db:"is_active,required"`
	CreatedAt time.Time `db:"created_at,required"`
	UpdatedAt time.Time `db:"updated_at,omitempty"`
	DeleteAt  time.Time `db:"deleted_at,omitempty"`
}
