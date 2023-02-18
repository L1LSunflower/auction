package db_repository

import (
	cityRepo "github.com/L1LSunflower/auction/internal/domain/repositories/db_repository/cities"
	usersRepo "github.com/L1LSunflower/auction/internal/domain/repositories/db_repository/users"
)

var (
	UserInterface usersRepo.UserInterface
	CityInterface cityRepo.CityInterface
)

func init() {
	UserInterface = usersRepo.NewRepository()
	CityInterface = cityRepo.NewRepository()
}
