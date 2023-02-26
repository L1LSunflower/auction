package users

import (
	"context"
	"encoding/json"
	"github.com/L1LSunflower/auction/internal/tools/context_with_depends"
	"time"

	"github.com/L1LSunflower/auction/internal/domain/entities"
)

const (
	timeExpirationUser  = 40 * time.Minute
	timeExpirationCode  = 30 * time.Minute
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
	if err := rClient.Set(rClient.Context(), user.ID, dataBytes, timeExpirationUser).Err(); err != nil {
		return err
	}

	return nil
}

func (r *Repository) User(ctx context.Context, id string) (*entities.User, error) {
	var user *entities.User
	rClient, err := context_with_depends.GetRedis(ctx)
	if err != nil {
		return nil, err
	}

	userString, err := rClient.Get(rClient.Context(), id).Result()
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

	if err := rClient.Set(rClient.Context(), id+"_code", code, timeExpirationCode).Err(); err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetUserCode(ctx context.Context, id string) (string, error) {
	rClient, err := context_with_depends.GetRedis(ctx)
	if err != nil {
		return "", err
	}

	code, err := rClient.Get(rClient.Context(), id+"_code").Result()
	if err != nil {
		return "", err
	}

	return code, nil
}

func (r *Repository) StoreToken(ctx context.Context, tokens *entities.Tokens) error {
	rClient, err := context_with_depends.GetRedis(ctx)
	if err != nil {
		return err
	}

	if err := rClient.Set(rClient.Context(), tokens.AccessToken+":access", tokens.RefreshToken, timeExpirationToken).Err(); err != nil {
		return err
	}

	return nil
}

func (r *Repository) Tokens(ctx context.Context, accessToken string) (*entities.Tokens, error) {
	rClient, err := context_with_depends.GetRedis(ctx)
	if err != nil {
		return nil, err
	}

	refreshToken, err := rClient.Get(rClient.Context(), accessToken).Result()
	if err != nil {
		return nil, err
	}

	return &entities.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (r *Repository) TokenByKey(ctx context.Context, tokenString string) (string, error) {
	rClient, err := context_with_depends.GetRedis(ctx)
	if err != nil {
		return "", err
	}

	token, err := rClient.Get(rClient.Context(), tokenString).Result()

	if err != nil {
		return "", err
	}

	return token, err
}
