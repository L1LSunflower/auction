package users

import (
	"github.com/L1LSunflower/auction/internal/domain/entities"
)

type UserInterface interface {
	Create(user *entities.User) error
	User(uuid string) (*entities.User, error)
	StoreUserCode(id, code string) error
	GetUserCode(id string) (string, error)
	StoreToken(tokens *entities.Tokens) error
	Tokens(accessToken string) (*entities.Tokens, error)
}

func GetUsesInterface() UserInterface {
	//return &Repository{redisClient: redisConn.RedisInstance().RedisClient}
	return &Repository{}
}
