package users

import (
	"context"
	"database/sql"
	"errors"
	"github.com/L1LSunflower/auction/internal/domain/entities"
	"github.com/L1LSunflower/auction/internal/domain/repositories/db_repository"
	"github.com/L1LSunflower/auction/internal/domain/repositories/redis_repository"
	"github.com/L1LSunflower/auction/internal/domain/services"
	userRequest "github.com/L1LSunflower/auction/internal/requests/structs/users"
	"github.com/L1LSunflower/auction/internal/tools/context_with_tx"
	"github.com/gofrs/uuid"
)

func SignUp(request *userRequest.SignUp) (*entities.User, error) {
	var (
		user = &entities.User{}
		err  error
	)

	//user, _ = redis_repository.UserInterface.User(request.Phone)
	//; err != nil {
	//	return nil, err
	//}

	//if user != nil {
	//	return user, nil
	//}

	code := services.GenerateRandomCode()
	//if err = sms.SendSMS(request.Phone, code); err != nil {
	//	return nil, err
	//}

	uid, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	if err = redis_repository.UserInterface.StoreUserCode(uid.String(), code); err != nil {
		return nil, err
	}

	user = &entities.User{
		ID:        uid.String(),
		Email:     request.Email,
		FirstName: request.Email,
		LastName:  request.LastName,
		Phone:     request.Phone,
		Password:  request.Password,
		//IsActive:  1,
	}

	if err = redis_repository.UserInterface.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func Confirm(parentCtx context.Context, db *sql.DB, request *userRequest.Confirm) (*entities.Tokens, error) {
	var (
		user = &entities.User{}
		err  error
	)

	code, err := redis_repository.UserInterface.GetUserCode(request.ID)
	if err != nil {
		return nil, err
	}

	if request.Code != code {
		return nil, errors.New("wrong code")
	}

	ctx, err := context_with_tx.ContextWithTx(parentCtx, db)
	if err != nil {
		return nil, err
	}
	defer context_with_tx.TxRollback(ctx)

	if user, err = redis_repository.UserInterface.User(request.ID); err != nil {
		return nil, err
	}

	if err = db_repository.UserInterface.Create(ctx, user); err != nil {
		return nil, err
	}

	tokens := services.GenerateToken()
	if err = redis_repository.UserInterface.StoreToken(tokens); err != nil {
		return nil, err
	}

	context_with_tx.TxCommit(ctx)

	return tokens, nil
}

func SignIn(parentCtx context.Context, db *sql.DB, request *userRequest.SignIn) (*entities.Tokens, error) {
	ctx, err := context_with_tx.ContextWithTx(parentCtx, db)
	if err != nil {
		return nil, err
	}
	defer context_with_tx.TxRollback(ctx)

	user, err := db_repository.UserInterface.UserByPhone(ctx, request.Phone)
	if err != nil {
		return nil, err
	}

	if request.Password != user.Password {
		return nil, errors.New("wrong password")
	}

	tokens := services.GenerateToken()
	if err = redis_repository.UserInterface.StoreToken(tokens); err != nil {
		return nil, err
	}

	context_with_tx.TxCommit(ctx)
	return tokens, nil
}

func RefreshToken(request *userRequest.Tokens) (*entities.Tokens, error) {
	tokens, err := redis_repository.UserInterface.Tokens(request.AccessToken)
	if err != nil {
		return nil, err
	}

	if request.AccessToken+request.RefreshToken != tokens.AccessToken+tokens.RefreshToken {
		return nil, errors.New("wrong token")
	}

	newTokens := services.GenerateToken()
	if err = redis_repository.UserInterface.StoreToken(newTokens); err != nil {
		return nil, err
	}

	return newTokens, nil
}

func User(parentCtx context.Context, db *sql.DB, request *userRequest.User) (*entities.User, error) {
	ctx, err := context_with_tx.ContextWithTx(parentCtx, db)
	if err != nil {
		return nil, err
	}
	defer context_with_tx.TxRollback(ctx)

	user, err := db_repository.UserInterface.User(ctx, request.ID)
	if err != nil {
		return nil, err
	}

	context_with_tx.TxCommit(ctx)
	return user, nil
}
