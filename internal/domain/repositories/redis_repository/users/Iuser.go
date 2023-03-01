package users

import (
	"context"
	"github.com/L1LSunflower/auction/internal/domain/entities"
)

type UserInterface interface {
	Create(ctx context.Context, user *entities.User) error
	User(ctx context.Context, phone string) (*entities.User, error)
	StoreUserCode(ctx context.Context, id, code string) error
	GetUserCode(ctx context.Context, id string) (string, error)
	StoreToken(ctx context.Context, userID string, tokens *entities.Tokens) error
	Tokens(ctx context.Context, userID string) (*entities.Tokens, error)
	GetRestoreCode(ctx context.Context, id string) (string, error)
}

func GetUsesInterface() UserInterface {
	return &Repository{}
}
