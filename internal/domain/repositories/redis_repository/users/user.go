package users

import (
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"time"

	"github.com/L1LSunflower/auction/internal/domain/entities"
)

const (
	timeExpirationUser  = 40 * time.Minute
	timeExpirationCode  = 30 * time.Minute
	timeExpirationToken = 1 * time.Hour
)

type Repository struct {
	redisClient *redis.Client
}

func (r *Repository) Create(user *entities.User) error {
	dataBytes, err := json.Marshal(user)
	if err != nil {
		return err
	}
	if err := r.redisClient.Set(r.redisClient.Context(), user.ID, dataBytes, timeExpirationUser).Err(); err != nil {
		return err
	}

	return nil
}

func (r *Repository) User(id string) (*entities.User, error) {
	var (
		user *entities.User
		err  error
	)

	userString, err := r.redisClient.Get(r.redisClient.Context(), id).Result()
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal([]byte(userString), &user); err != nil {
		return nil, err
	}

	return user, nil
}

func (r *Repository) StoreUserCode(id, code string) error {
	if err := r.redisClient.Set(r.redisClient.Context(), id+"_code", code, timeExpirationCode).Err(); err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetUserCode(id string) (string, error) {
	code, err := r.redisClient.Get(r.redisClient.Context(), id+"_code").Result()
	if err != nil {
		return "", err
	}

	return code, nil
}

func (r *Repository) StoreToken(tokens *entities.Tokens) error {
	if err := r.redisClient.Set(r.redisClient.Context(), tokens.AccessToken, tokens.RefreshToken, timeExpirationToken).Err(); err != nil {
		return err
	}

	return nil
}

func (r *Repository) Tokens(accessToken string) (*entities.Tokens, error) {
	refreshToken, err := r.redisClient.Get(r.redisClient.Context(), accessToken).Result()
	if err != nil {
		return nil, err
	}

	return &entities.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
