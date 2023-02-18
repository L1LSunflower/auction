package redis_repository

import (
	userInterface "github.com/L1LSunflower/auction/internal/domain/repositories/redis_repository/users"
)

var (
	UserInterface userInterface.UserInterface
)

func init() {
	UserInterface = userInterface.GetUsesInterface()
}
