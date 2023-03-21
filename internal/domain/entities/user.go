package entities

import (
	"time"

	"github.com/L1LSunflower/auction/internal/requests/structs/users"
)

type User struct {
	ID        string    `db:"id,required"`
	Phone     string    `db:"phone,required"`
	Email     string    `db:"email,required"`
	Password  string    `db:"password,required"`
	FirstName string    `db:"first_name,required"`
	LastName  string    `db:"last_name,required"`
	City      string    `db:"city,omitempty"`
	CreatedAt time.Time `db:"created_at,omitempty"`
	UpdatedAt time.Time `db:"updated_at,omitempty"`
}

func NewUser() *User {
	return &User{}
}

func NewUserFromRequest(uuid string, request *users.SignUp) *User {
	return &User{
		ID:        uuid,
		Phone:     request.Phone,
		Email:     request.Email,
		Password:  request.Password,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		City:      request.City,
	}
}
