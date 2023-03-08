package aggregates

import "github.com/L1LSunflower/auction/internal/domain/entities"

type UserBalance struct {
	User    *entities.User
	Balance *entities.Balance
}
