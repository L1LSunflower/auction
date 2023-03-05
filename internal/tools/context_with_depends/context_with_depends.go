package context_with_depends

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
)

const (
	depends = "depends"
)

type DBAndRedis struct {
	DB        *sql.DB
	Redis     *redis.Client
	DBTx      *sql.Tx
	RedisDBTx *redis.Tx
}

func ContextWithDepends(parentCtx context.Context, db *sql.DB, redis *redis.Client) (context.Context, error) {
	dep := &DBAndRedis{
		DB:    db,
		Redis: redis,
	}

	return context.WithValue(parentCtx, depends, dep), nil
}

func StartDBTx(ctx context.Context) error {
	var err error
	dep, ok := ctx.Value(depends).(*DBAndRedis)
	if !ok {
		return errors.New("failed to get dependencies")
	}

	if dep.DBTx, err = dep.DB.BeginTx(context.Background(), nil); err != nil {
		return errors.New(fmt.Sprintf("failed to start db tx with error: %s", err.Error()))
	}

	return nil
}

func GetDb(ctx context.Context) (*sql.DB, error) {
	dep, ok := ctx.Value(depends).(*DBAndRedis)
	if !ok {
		return nil, errors.New("failed to get dependency db")
	}

	return dep.DB, nil
}

func GetRedis(ctx context.Context) (*redis.Client, error) {
	dep, ok := ctx.Value(depends).(*DBAndRedis)
	if !ok {
		return nil, errors.New("failed to get dependency redis")
	}

	return dep.Redis, nil
}

func TxFromContext(ctx context.Context) (*sql.Tx, error) {
	dep, ok := ctx.Value(depends).(*DBAndRedis)
	if !ok {
		return nil, errors.New("failed to get tx from context")
	}
	return dep.DBTx, nil
}

func DBTxRollback(ctx context.Context) error {
	tx, err := TxFromContext(ctx)
	if err != nil {
		return err
	}

	tx.Rollback()
	ctx.Done()
	return nil
}

func DBTxCommit(ctx context.Context) error {
	tx, err := TxFromContext(ctx)
	if err != nil {
		return err
	}

	tx.Commit()
	ctx.Done()
	return nil
}
