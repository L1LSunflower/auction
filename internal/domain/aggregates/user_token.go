package aggregates

import "github.com/L1LSunflower/auction/internal/domain/entities"

type UserToken struct {
	User  *entities.User
	Token *entities.Tokens
}
