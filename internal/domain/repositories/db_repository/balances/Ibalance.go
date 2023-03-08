package balances

import (
	"context"

	"github.com/L1LSunflower/auction/internal/domain/entities"
)

type BalanceInterface interface {
	Create(ctx context.Context, userID string) (*entities.Balance, error)
	Balance(ctx context.Context, userID string) (*entities.Balance, error)
	Credit(ctx context.Context, userID string, amount float64) (*entities.Balance, error)
	Debit(ctx context.Context, userID string, amount float64) (*entities.Balance, error)
}

func GetBalanceInterface() BalanceInterface {
	return &Repository{}
}
