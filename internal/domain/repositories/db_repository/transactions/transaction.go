package transactions

import (
	"context"
	"fmt"
	"github.com/L1LSunflower/auction/internal/tools/context_with_depends"
	"strings"
	"time"

	"github.com/L1LSunflower/auction/internal/domain/entities"
)

const (
	fieldsSeparator = ","
)

type Repository struct{}

func (r *Repository) Create(ctx context.Context, userID, trType string, amount float64) (*entities.Transaction, error) {
	tx, err := context_with_depends.TxFromContext(ctx)
	if err != nil {
		return nil, err
	}

	fields := []string{
		"user_id",
		"amount",
		"type",
		"created_at",
	}

	tr := &entities.Transaction{ID: userID, Amount: amount, Type: trType, CreatedAt: time.Now()}
	query := fmt.Sprintf("insert into auction_transactions (%s) values (?, ?, ?, ?)", strings.Join(fields, fieldsSeparator))
	if _, err = tx.Exec(query, tr.ID, tr.Amount, tr.Type, tr.CreatedAt); err != nil {
		return nil, err
	}

	return tr, nil
}

func (r *Repository) Transaction(ctx context.Context, userID, typeTransaction string) (*entities.Transaction, error) {
	db, err := context_with_depends.GetDb(ctx)
	if err != nil {
		return nil, err
	}

	fields := []string{
		"user_id",
		"amount",
		"type",
		"created_at",
	}

	tr := &entities.Transaction{}
	query := fmt.Sprintf("select %s from auction_transactions where user_id=? and type=?", strings.Join(fields, fieldsSeparator))
	if err = db.QueryRow(query, userID, typeTransaction).Scan(&tr.ID, &tr.Amount, &tr.Type, &tr.CreatedAt); err != nil {
		return nil, err
	}

	return tr, nil
}
