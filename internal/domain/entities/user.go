package entities

import "time"

type User struct {
	ID        string    `db:"id,required"`
	Phone     string    `db:"phone,required"`
	Email     string    `db:"email,required"`
	Password  string    `db:"password,required"`
	FirstName string    `db:"first_name,required"`
	LastName  string    `db:"last_name,required"`
	City      string    `db:"cities,omitempty"`
	CreatedAt time.Time `db:"created_at,required"`
	UpdatedAt time.Time `db:"updated_at,omitempty"`
}
