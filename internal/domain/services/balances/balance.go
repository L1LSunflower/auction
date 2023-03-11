package balances

import (
	"context"
	"github.com/L1LSunflower/auction/internal/domain/aggregates"
	"github.com/L1LSunflower/auction/internal/domain/entities"
	"github.com/L1LSunflower/auction/internal/domain/repositories/db_repository"
	"github.com/L1LSunflower/auction/internal/tools/context_with_depends"
	"github.com/L1LSunflower/auction/internal/tools/errorhandler"

	balanceReq "github.com/L1LSunflower/auction/internal/requests/structs/balances"
)

var legCreds = map[string]string{
	"5738 7710 7131 8780": "832",
	"5746 3882 5878 3536": "975",
	"4380 1622 7050 5613": "376",
}

func Credit(ctx context.Context, credit *balanceReq.Credit) (*aggregates.UserBalance, error) {
	var err error
	if err = context_with_depends.StartDBTx(ctx); err != nil {
		return nil, errorhandler.ErrDependency
	}
	defer context_with_depends.DBTxRollback(ctx)

	userBalance := &aggregates.UserBalance{}
	if userBalance.User, err = db_repository.UserInterface.User(ctx, credit.ID); err != nil {
		return nil, errorhandler.ErrUserNotExist
	}

	if len(credit.Pan) <= 0 || len(credit.CVV) <= 0 {
		return nil, errorhandler.ErrProcessCard
	}

	if legCreds[credit.Pan] != credit.CVV {
		return nil, errorhandler.ErrProcessCard
	}

	if userBalance.Balance, err = db_repository.BalanceInterface.Credit(ctx, userBalance.User.ID, credit.Amount); err != nil {
		return nil, errorhandler.ErrCreditBalance
	}

	if _, err = db_repository.TransactionsInterface.Create(ctx, userBalance.User.ID, entities.CreditType, credit.Amount); err != nil {
		return nil, errorhandler.ErrCreateTransaction
	}

	context_with_depends.DBTxCommit(ctx)

	return userBalance, nil
}

func Debit(ctx context.Context, debit *balanceReq.Debit) (*aggregates.UserBalance, error) {
	var err error
	if err = context_with_depends.StartDBTx(ctx); err != nil {
		return nil, errorhandler.ErrDependency
	}
	defer context_with_depends.DBTxRollback(ctx)

	userBalance := &aggregates.UserBalance{}
	if userBalance.User, err = db_repository.UserInterface.User(ctx, debit.ID); err != nil {
		return nil, errorhandler.ErrUserNotExist
	}

	if userBalance.Balance, err = db_repository.BalanceInterface.Debit(ctx, userBalance.User.ID, debit.Amount); err != nil {
		return nil, errorhandler.ErrDebitBalance
	}

	if _, err = db_repository.TransactionsInterface.Create(ctx, userBalance.User.ID, entities.DebitType, debit.Amount); err != nil {
		return nil, errorhandler.ErrCreateTransaction
	}

	context_with_depends.DBTxCommit(ctx)

	return userBalance, nil
}

func Balance(ctx context.Context, balance *balanceReq.Balance) (*aggregates.UserBalance, error) {
	var err error
	userBalance := &aggregates.UserBalance{}
	if userBalance.User, err = db_repository.UserInterface.User(ctx, balance.ID); err != nil {
		return nil, errorhandler.ErrUserNotExist
	}

	if userBalance.Balance, err = db_repository.BalanceInterface.Balance(ctx, userBalance.User.ID); err != nil {
		return nil, errorhandler.ErrGetBalance
	}

	return userBalance, nil
}
