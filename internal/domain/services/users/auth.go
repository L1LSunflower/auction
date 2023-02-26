package users

import (
	"context"
	"errors"

	"github.com/gofrs/uuid"

	"github.com/L1LSunflower/auction/internal/domain/entities"
	"github.com/L1LSunflower/auction/internal/domain/repositories/db_repository"
	"github.com/L1LSunflower/auction/internal/domain/repositories/redis_repository"
	"github.com/L1LSunflower/auction/internal/domain/services"
	userRequest "github.com/L1LSunflower/auction/internal/requests/structs/users"
	"github.com/L1LSunflower/auction/internal/tools/context_with_depends"
	"github.com/L1LSunflower/auction/pkg/sms"
)

func SignUp(ctx context.Context, request *userRequest.SignUp) (*entities.User, error) {
	var (
		user = &entities.User{}
		err  error
	)

	user, _ = redis_repository.UserInterface.User(ctx, request.Phone)
	if user != nil {
		return nil, errors.New("need account confirmation")
	}

	code := services.GenerateRandomCode()
	if err = sms.SendSMS(request.Phone, code); err != nil {
		return nil, err
	}

	uid, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	if err = redis_repository.UserInterface.StoreUserCode(ctx, uid.String(), code); err != nil {
		return nil, err
	}

	user = &entities.User{
		ID:        uid.String(),
		Email:     request.Email,
		FirstName: request.Email,
		LastName:  request.LastName,
		Phone:     request.Phone,
		Password:  request.Password,
	}

	if err = redis_repository.UserInterface.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func Confirm(ctx context.Context, request *userRequest.Confirm) (*entities.Tokens, error) {
	var (
		user = &entities.User{}
		err  error
	)

	code, err := redis_repository.UserInterface.GetUserCode(ctx, request.ID)
	if err != nil {
		return nil, err
	}

	if request.Code != code {
		return nil, errors.New("wrong code")
	}

	if err = context_with_depends.StartDBTx(ctx); err != nil {
		return nil, err
	}
	defer context_with_depends.DBTxRollback(ctx)

	if user, err = redis_repository.UserInterface.User(ctx, request.ID); err != nil {
		return nil, err
	}

	if err = db_repository.UserInterface.Create(ctx, user); err != nil {
		return nil, err
	}

	tokens := services.GenerateToken()
	if err = redis_repository.UserInterface.StoreToken(ctx, tokens); err != nil {
		return nil, err
	}

	context_with_depends.DBTxCommit(ctx)

	return tokens, nil
}

func SignIn(ctx context.Context, request *userRequest.SignIn) (*entities.Tokens, string, error) {
	if err := context_with_depends.StartDBTx(ctx); err != nil {
		return nil, "", err
	}
	defer context_with_depends.DBTxRollback(ctx)

	user, err := db_repository.UserInterface.UserByPhone(ctx, request.Phone)
	if err != nil {
		return nil, "", err
	}

	if user.CreatedAt.IsZero() {
		return nil, "", errors.New("that user does not exist")
	}

	if request.Password != user.Password {
		return nil, "", errors.New("wrong password")
	}

	tokens := services.GenerateToken()
	if err = redis_repository.UserInterface.StoreToken(ctx, tokens); err != nil {
		return nil, "", err
	}

	context_with_depends.DBTxCommit(ctx)
	return tokens, user.ID, nil
}

func RefreshToken(ctx context.Context, request *userRequest.Tokens) (*entities.Tokens, error) {
	tokens, err := redis_repository.UserInterface.Tokens(ctx, request.AccessToken)
	if err != nil {
		return nil, err
	}

	if request.AccessToken+request.RefreshToken != tokens.AccessToken+tokens.RefreshToken {
		return nil, errors.New("wrong token")
	}

	newTokens := services.GenerateToken()
	if err = redis_repository.UserInterface.StoreToken(ctx, newTokens); err != nil {
		return nil, err
	}

	return newTokens, nil
}

func RestorePassword(ctx context.Context, request *userRequest.RestorePassword) error {
	user, err := db_repository.UserInterface.UserByPhone(ctx, request.Phone)
	if err != nil {
		return errors.New("that user does not exist")
	}

	if !user.CreatedAt.IsZero() {
		return errors.New("that user does not exist")
	}

	code := services.GenerateRandomCode()
	if err := sms.SendSMS(request.Phone, code); err != nil {
		return err
	}

	return nil
}
