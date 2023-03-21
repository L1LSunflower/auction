package users

import (
	"context"
	"fmt"
	"github.com/L1LSunflower/auction/internal/domain/aggregates"
	"github.com/L1LSunflower/auction/internal/domain/entities"
	"github.com/L1LSunflower/auction/internal/domain/repositories/db_repository"
	"github.com/L1LSunflower/auction/internal/domain/repositories/redis_repository"
	"github.com/L1LSunflower/auction/internal/domain/services"
	userRequest "github.com/L1LSunflower/auction/internal/requests/structs/users"
	"github.com/L1LSunflower/auction/internal/tools/context_with_depends"
	"github.com/L1LSunflower/auction/internal/tools/errorhandler"
	"github.com/L1LSunflower/auction/pkg/logger"
	"github.com/L1LSunflower/auction/pkg/logger/message"
	"github.com/go-redis/redis/v8"
	"github.com/gofrs/uuid"
)

func SignUp(ctx context.Context, request *userRequest.SignUp) (*entities.User, error) {
	var (
		user = entities.NewUser()
		err  error
	)

	user, err = redis_repository.UserInterface.User(ctx, request.Phone)
	if err != nil && err != redis.Nil {
		logger.Log.Error(message.NewMessage(fmt.Sprintf("failed to store user with error: %s", err.Error())))
		return nil, errorhandler.ErrStoreUser
	}

	if user != nil {
		return nil, errorhandler.ErrNeedConfirm
	}

	user, err = db_repository.UserInterface.UserByPhone(ctx, request.Phone)
	if err != nil {
		logger.Log.Error(message.NewMessage(fmt.Sprintf("failed to get user by phone with error: %s", err.Error())))
		return nil, errorhandler.InternalError
	}

	if !user.CreatedAt.IsZero() {
		return nil, errorhandler.ErrUserExist
	}

	uid, err := uuid.NewV7()
	if err != nil {
		return nil, errorhandler.InternalError
	}

	user = entities.NewUserFromRequest(uid.String(), request)
	if err = redis_repository.UserInterface.Create(ctx, user); err != nil {
		return nil, errorhandler.ErrCreateUser
	}

	code := services.GenerateRandomCode()
	//if err = sms.SendSMS(request.Phone, code); err != nil {
	//logger.Log.Error(message.NewMessage(fmt.Sprintf("failed to send otp code on phone: %s with error: %s", request.Phone, err.Error())))
	//	return nil, errorhandler.ErrSendOtp
	//}
	if err = redis_repository.UserInterface.StoreUserCode(ctx, uid.String(), code); err != nil {
		logger.Log.Error(message.NewMessage(fmt.Sprintf("failed to store user: %s, otp code with error: %s", uid.String(), err.Error())))
		return nil, errorhandler.ErrStoreOtp
	}

	return user, nil
}

func Confirm(ctx context.Context, request *userRequest.Confirm) (*aggregates.UserToken, error) {
	var (
		userToken = &aggregates.UserToken{}
		err       error
	)

	code, err := redis_repository.UserInterface.GetUserCode(ctx, request.ID)
	if err != nil {
		return nil, errorhandler.ErrCodeExpired
	}

	if request.Code != code {
		return nil, errorhandler.WrongCode
	}

	if err = context_with_depends.StartDBTx(ctx); err != nil {
		return nil, errorhandler.ErrDependency
	}
	defer context_with_depends.DBTxRollback(ctx)

	if userToken.User, err = db_repository.UserInterface.UserByPhone(ctx, request.Phone); userToken.User == nil {
		return nil, errorhandler.ErrConfirm
	}

	if userToken.User, err = redis_repository.UserInterface.User(ctx, request.Phone); err != nil {
		return nil, errorhandler.ErrUserExpired
	}

	if err = db_repository.UserInterface.Create(ctx, userToken.User); err != nil {
		return nil, errorhandler.ErrStoreUser
	}

	if _, err = db_repository.BalanceInterface.Create(ctx, userToken.User.ID); err != nil {
		return nil, errorhandler.ErrCreateBalance
	}

	userToken.Token = services.GenerateToken()
	if err = redis_repository.UserInterface.StoreToken(ctx, userToken.User.ID, userToken.Token); err != nil {
		return nil, errorhandler.ErrStoreToken
	}

	context_with_depends.DBTxCommit(ctx)

	return userToken, nil
}

func SignIn(ctx context.Context, request *userRequest.SignIn) (*aggregates.UserToken, error) {
	var (
		userToken = &aggregates.UserToken{}
		err       error
	)

	if userToken.User, err = db_repository.UserInterface.UserByPhone(ctx, request.Phone); err != nil || userToken.User == nil {
		return nil, errorhandler.ErrFindByPhone
	}

	if userToken.User.CreatedAt.IsZero() {
		return nil, errorhandler.ErrFindByPhone
	}

	if request.Password != userToken.User.Password {
		return nil, errorhandler.WrongPassword
	}

	userToken.Token = services.GenerateToken()
	if err = redis_repository.UserInterface.StoreToken(ctx, userToken.User.ID, userToken.Token); err != nil {
		return nil, errorhandler.ErrStoreToken
	}

	return userToken, nil
}

func RefreshToken(ctx context.Context, request *userRequest.Tokens) (*entities.Tokens, error) {
	tokens, err := redis_repository.UserInterface.Tokens(ctx, request.ID)
	if err != nil {
		return nil, errorhandler.ErrGetTokens
	}

	if request.AccessToken+request.RefreshToken != tokens.AccessToken+tokens.RefreshToken {
		return nil, errorhandler.WrongTokens
	}

	newTokens := services.GenerateToken()
	if err = redis_repository.UserInterface.StoreToken(ctx, request.ID, newTokens); err != nil {
		return nil, errorhandler.ErrStoreToken
	}

	return newTokens, nil
}

func SendRestoreCode(ctx context.Context, request *userRequest.RestorePassword) error {
	user, err := db_repository.UserInterface.UserByPhone(ctx, request.Phone)
	if err != nil || user == nil {
		return errorhandler.ErrUserExist
	}

	code := services.GenerateRandomCode()
	//if err := sms.SendSMS(request.Phone, code); err != nil {
	//	return err
	//}

	if err = redis_repository.UserInterface.StoreUserCode(ctx, user.Phone, code); err != nil {
		return errorhandler.ErrStoreOtp
	}

	return nil
}

func ChangePassword(ctx context.Context, request *userRequest.ChangePassword) (*aggregates.UserToken, error) {
	var err error

	if err = context_with_depends.StartDBTx(ctx); err != nil {
		return nil, errorhandler.ErrDependency
	}
	defer context_with_depends.DBTxRollback(ctx)

	userToken := &aggregates.UserToken{}
	if userToken.User, err = db_repository.UserInterface.UserByPhone(ctx, request.Phone); err != nil || userToken.User == nil {
		return nil, errorhandler.ErrUserNotExist
	}

	correctCode, err := redis_repository.UserInterface.GetUserCode(ctx, request.Phone)
	if err != nil {
		return nil, errorhandler.ErrCodeExpired
	}

	if correctCode != request.Code {
		return nil, errorhandler.WrongCode
	}

	userToken.User.Password = request.Password
	if err = db_repository.UserInterface.UpdatePassword(ctx, userToken.User); err != nil {
		return nil, errorhandler.ErrUpdatePassword
	}

	userToken.Token = services.GenerateToken()
	if err = redis_repository.UserInterface.StoreToken(ctx, userToken.User.ID, userToken.Token); err != nil {
		return nil, errorhandler.ErrStoreToken
	}

	context_with_depends.DBTxCommit(ctx)

	return userToken, nil
}
