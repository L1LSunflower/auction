package users

import (
	"context"
	"errors"
	"github.com/L1LSunflower/auction/internal/domain/entities"
	"github.com/L1LSunflower/auction/internal/domain/repositories/db_repository"
	userRequest "github.com/L1LSunflower/auction/internal/requests/structs/users"
	"github.com/L1LSunflower/auction/internal/tools/context_with_depends"
	"github.com/L1LSunflower/auction/internal/tools/errorhandler"
)

func User(ctx context.Context, request *userRequest.User) (*entities.User, error) {
	user, _ := db_repository.UserInterface.User(ctx, request.ID)
	if user == nil {
		return nil, errorhandler.ErrUserExist
	}

	//if user.CreatedAt.IsZero() {
	//	return nil, errorhandler.ErrUserExist
	//}

	return user, nil
}

func Update(ctx context.Context, request *userRequest.Update) (*entities.User, error) {
	if err := context_with_depends.StartDBTx(ctx); err != nil {
		return nil, err
	}
	defer context_with_depends.DBTxRollback(ctx)

	user, err := db_repository.UserInterface.User(ctx, request.ID)
	if err != nil {
		return nil, err
	}

	user = &entities.User{
		ID:        request.ID,
		Email:     request.Email,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Phone:     request.Phone,
		Password:  request.Password,
	}

	if err := db_repository.UserInterface.Update(ctx, user); err != nil {
		return nil, errors.New("failed to update user")
	}

	context_with_depends.DBTxCommit(ctx)

	return user, nil
}

func Delete(ctx context.Context, request *userRequest.Delete) error {
	if err := context_with_depends.StartDBTx(ctx); err != nil {
		return err
	}
	defer context_with_depends.DBTxRollback(ctx)

	if _, err := db_repository.UserInterface.User(ctx, request.ID); err != nil {
		return err
	}

	if err := db_repository.UserInterface.Delete(ctx, request.ID); err != nil {
		return err
	}

	return nil
}
