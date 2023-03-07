package balances

import (
	"context"
	"fmt"
	"github.com/L1LSunflower/auction/internal/domain/entities"
	"github.com/L1LSunflower/auction/internal/tools/context_with_depends"
	"github.com/L1LSunflower/auction/internal/tools/errorhandler"
	"strings"
	"time"
)

const fieldsSeparator = ","

type Repository struct{}

func (r *Repository) Create(ctx context.Context, userID string) (*entities.Balance, error) {
	tx, err := context_with_depends.TxFromContext(ctx)
	if err != nil {
		return nil, err
	}

	fields := []string{
		"id",
		"balance",
		"created_at",
		"updated_at",
	}

	now := time.Now()
	accBalance := &entities.Balance{ID: userID, Balance: 0, CreatedAt: now, UpdatedAt: now}
	query := fmt.Sprintf("insert into balance (%s) values (?, ?, ?, ?)", strings.Join(fields, fieldsSeparator))
	if _, err = tx.Exec(query, accBalance.ID, accBalance.Balance, accBalance.CreatedAt, accBalance.UpdatedAt); err != nil {
		return nil, err
	}

	return accBalance, nil
}

func (r *Repository) Balance(ctx context.Context, userID string) (*entities.Balance, error) {
	db, err := context_with_depends.GetDb(ctx)
	if err != nil {
		return nil, err
	}

	fields := []string{
		"id",
		"balance",
		"created_at",
		"updated_at",
	}

	accBalance := &entities.Balance{}
	query := fmt.Sprintf("select %s from balance where id=? and deleted_at is null", strings.Join(fields, fieldsSeparator))
	if err = db.QueryRow(query, userID).Scan(&accBalance.ID, &accBalance.Balance, &accBalance.CreatedAt, &accBalance.UpdatedAt); err != nil {
		return nil, err
	}

	return accBalance, nil
}

func (r *Repository) Credit(ctx context.Context, userID string, amount float64) (*entities.Balance, error) {
	tx, err := context_with_depends.TxFromContext(ctx)
	if err != nil {
		return nil, err
	}

	if _, err = tx.Exec("update balance set balance = balance + ? where id=?", userID, amount); err != nil {
		return nil, err
	}

	fields := []string{
		"id",
		"balance",
		"created_at",
		"updated_at",
	}

	accBalance := &entities.Balance{}
	query := fmt.Sprintf("select %s from balance where id=?", strings.Join(fields, fieldsSeparator))
	if err = tx.QueryRow(query, userID).Scan(&accBalance.ID, &accBalance.Balance, &accBalance.CreatedAt, &accBalance.UpdatedAt); err != nil {
		return nil, err
	}

	return accBalance, nil
}

func (r *Repository) Debit(ctx context.Context, userID string, amount float64) (*entities.Balance, error) {
	tx, err := context_with_depends.TxFromContext(ctx)
	if err != nil {
		return nil, err
	}

	fields := []string{
		"id",
		"balance",
		"created_at",
		"updated_at",
	}

	accBalance := &entities.Balance{}
	query := fmt.Sprintf("select %s from balance where id=?", strings.Join(fields, fieldsSeparator))
	if err = tx.QueryRow(query, userID).Scan(&accBalance.ID, &accBalance.Balance, &accBalance.CreatedAt, &accBalance.UpdatedAt); err != nil {
		return nil, err
	}

	if accBalance.Balance < amount {
		return nil, errorhandler.NotEnoughBalance
	}

	if _, err = tx.Exec("update balance set balance = balance - ? where id=?", userID, amount); err != nil {
		return nil, err
	}

	return accBalance, nil
}
