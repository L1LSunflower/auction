package cities

import (
	"context"
	"database/sql"

	"github.com/L1LSunflower/auction/internal/domain/entities"
	"github.com/L1LSunflower/auction/internal/domain/repositories/db_repository"
	cityRequest "github.com/L1LSunflower/auction/internal/requests/structs/cities"
	"github.com/L1LSunflower/auction/internal/tools/context_with_tx"
	"github.com/L1LSunflower/auction/internal/tools/errorhandler"
)

func Create(parentCtx context.Context, db *sql.DB, request cityRequest.CreateCity) (*entities.City, error) {
	ctx, err := context_with_tx.ContextWithTx(parentCtx, db)
	if err != nil {
		return nil, err
	}
	defer context_with_tx.TxRollback(ctx)

	city, err := db_repository.CityInterface.Create(ctx, request.Name)
	if err != nil {
		return nil, errorhandler.ErrCreateCity
	}

	context_with_tx.TxCommit(ctx)
	return city, nil
}

func Cities(parentCtx context.Context, db *sql.DB) ([]*entities.City, error) {
	ctx, err := context_with_tx.ContextWithTx(parentCtx, db)
	if err != nil {
		return nil, err
	}
	defer context_with_tx.TxRollback(ctx)

	city, err := db_repository.CityInterface.GetAll(ctx)
	if err != nil {
		return nil, errorhandler.ErrCreateCity
	}

	context_with_tx.TxCommit(ctx)
	return city, nil
}

func City(parentCtx context.Context, db *sql.DB, request cityRequest.City) (*entities.City, error) {
	ctx, err := context_with_tx.ContextWithTx(parentCtx, db)
	if err != nil {
		return nil, err
	}
	defer context_with_tx.TxRollback(ctx)

	city, err := db_repository.CityInterface.Get(ctx, request.ID)
	if err != nil {
		return nil, errorhandler.ErrCreateCity
	}

	context_with_tx.TxCommit(ctx)
	return city, nil
}

func Update(parentCtx context.Context, db *sql.DB, request cityRequest.CityUpdate) error {
	ctx, err := context_with_tx.ContextWithTx(parentCtx, db)
	if err != nil {
		return err
	}
	defer context_with_tx.TxRollback(ctx)

	if _, err = db_repository.CityInterface.Update(ctx, request.ID, request.Name); err != nil {
		return errorhandler.ErrCreateCity
	}

	context_with_tx.TxCommit(ctx)
	return nil
}

func Delete(parentCtx context.Context, db *sql.DB, request cityRequest.City) error {
	ctx, err := context_with_tx.ContextWithTx(parentCtx, db)
	if err != nil {
		return err
	}
	defer context_with_tx.TxRollback(ctx)

	if err = db_repository.CityInterface.Delete(ctx, request.ID); err != nil {
		return errorhandler.ErrCreateCity
	}

	context_with_tx.TxCommit(ctx)
	return nil
}
