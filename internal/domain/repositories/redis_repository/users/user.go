package users

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/L1LSunflower/auction/internal/tools/context_with_depends"
	"time"

	"github.com/L1LSunflower/auction/internal/domain/entities"
)

const (
	userKey             = "users:data"
	userOtpKey          = "users:otp"
	userTokens          = "users:tokens"
	timeExpirationUser  = 15 * time.Minute
	timeExpirationCode  = 15 * time.Minute
	timeExpirationToken = 1 * time.Hour
)

type Repository struct{}

func (r *Repository) Create(ctx context.Context, user *entities.User) error {
	rClient, err := context_with_depends.GetRedis(ctx)
	if err != nil {
		return err
	}

	dataBytes, err := json.Marshal(user)
	if err != nil {
		return err
	}
	if err := rClient.Set(rClient.Context(), fmt.Sprintf("%s:%s", userKey, user.Phone), dataBytes, timeExpirationUser).Err(); err != nil {
		return err
	}

	return nil
}

func (r *Repository) ByPhone(ctx context.Context, id string) (*entities.User, error) {
	var user *entities.User
	rClient, err := context_with_depends.GetRedis(ctx)
	if err != nil {
		return nil, err
	}

	userString, err := rClient.Get(rClient.Context(), userKey).Result()

	if err = json.Unmarshal([]byte(userString), &user); err != nil {
		return nil, err
	}

	return user, nil
}

func (r *Repository) User(ctx context.Context, phone string) (*entities.User, error) {
	var user *entities.User
	rClient, err := context_with_depends.GetRedis(ctx)
	if err != nil {
		return nil, err
	}

	userString, err := rClient.Get(rClient.Context(), fmt.Sprintf("%s:%s", userKey, phone)).Result()
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal([]byte(userString), &user); err != nil {
		return nil, err
	}

	return user, nil
}

func (r *Repository) StoreUserCode(ctx context.Context, id, code string) error {
	rClient, err := context_with_depends.GetRedis(ctx)
	if err != nil {
		return err
	}

	if err := rClient.Set(rClient.Context(), fmt.Sprintf("%s:%s%s", userOtpKey, id, "otp"), code, timeExpirationCode).Err(); err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetUserCode(ctx context.Context, id string) (string, error) {
	rClient, err := context_with_depends.GetRedis(ctx)
	if err != nil {
		return "", err
	}

	code, err := rClient.Get(rClient.Context(), fmt.Sprintf("%s:%s%s", userOtpKey, id, "otp")).Result()
	if err != nil {
		return "", err
	}

	return code, nil
}

func (r *Repository) StoreToken(ctx context.Context, userID string, tokens *entities.Tokens) error {
	rClient, err := context_with_depends.GetRedis(ctx)
	if err != nil {
		return err
	}

	tokenBytes, err := json.Marshal(tokens)
	if err != nil {
		return err
	}

	if err := rClient.Set(rClient.Context(), fmt.Sprintf("%s:%s%s", userTokens, userID, "tokens"), tokenBytes, timeExpirationToken).Err(); err != nil {
		return err
	}

	return nil
}

func (r *Repository) Tokens(ctx context.Context, userID string) (*entities.Tokens, error) {
	tokens := &entities.Tokens{}
	rClient, err := context_with_depends.GetRedis(ctx)
	if err != nil {
		return nil, err
	}

	tokensString, err := rClient.Get(rClient.Context(), fmt.Sprintf("%s:%s%s", userTokens, userID, "tokens")).Result()
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal([]byte(tokensString), tokens); err != nil {
		return nil, err
	}

	return tokens, nil
}

func (r *Repository) StoreRestoreCode(ctx context.Context, code, id string) error {
	rClient, err := context_with_depends.GetRedis(ctx)
	if err != nil {
		return err
	}

	if err := rClient.Set(rClient.Context(), fmt.Sprintf("%s:%s%s", userOtpKey, id, "otp"), code, timeExpirationToken).Err(); err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetRestoreCode(ctx context.Context, id string) (string, error) {
	rClient, err := context_with_depends.GetRedis(ctx)
	if err != nil {
		return "", err
	}

	code, err := rClient.Get(rClient.Context(), fmt.Sprintf("%s:%s%s", userOtpKey, id, "otp")).Result()
	if err != nil {
		return "", err
	}

	return code, nil
}
