package transactions

import (
	"context"

	"github.com/L1LSunflower/auction/internal/domain/entities"
)

type TransactionsInterface interface {
	Create(ctx context.Context, userID, trType string, amount float64) (*entities.Transaction, error)
	Transaction(ctx context.Context, userID, typeTransaction string) (*entities.Transaction, error)
}

func GetTransactionsInterface() TransactionsInterface {
	return &Repository{}
}
