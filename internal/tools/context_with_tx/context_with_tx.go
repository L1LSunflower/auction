package context_with_tx

import (
	"context"
	"database/sql"
	"errors"
)

const (
	txKey = "tx"
)

func ContextWithTx(parentCtx context.Context, db *sql.DB) (context.Context, error) {
	tx, err := db.BeginTx(parentCtx, nil)
	if err != nil {
		return nil, err
	}
	return context.WithValue(parentCtx, txKey, tx), nil
}

func TxFromContext(ctx context.Context) (*sql.Tx, error) {
	tx, ok := ctx.Value("tx").(*sql.Tx)
	if !ok {
		return nil, errors.New("failed to get tx from context")
	}
	return tx, nil
}

func TxRollback(ctx context.Context) error {
	tx, err := TxFromContext(ctx)
	if err != nil {
		return err
	}
	tx.Rollback()
	ctx.Done()
	return nil
}

func TxCommit(ctx context.Context) error {
	tx, err := TxFromContext(ctx)
	if err != nil {
		return err
	}
	tx.Commit()
	ctx.Done()
	return nil
}
