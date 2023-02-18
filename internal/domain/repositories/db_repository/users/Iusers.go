package users

import (
	"context"

	"github.com/L1LSunflower/auction/internal/domain/entities"
)

type UserInterface interface {
	Create(ctx context.Context, user *entities.User) error
	User(ctx context.Context, id string) (*entities.User, error)
	UserByPhone(ctx context.Context, phone string) (*entities.User, error)
}

func NewRepository() UserInterface {
	return &Repository{}
}
