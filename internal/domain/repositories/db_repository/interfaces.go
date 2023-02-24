package db_repository

import (
	auctionsRepo "github.com/L1LSunflower/auction/internal/domain/repositories/db_repository/auctions"
	cityRepo "github.com/L1LSunflower/auction/internal/domain/repositories/db_repository/cities"
	usersRepo "github.com/L1LSunflower/auction/internal/domain/repositories/db_repository/users"
)

var (
	UserInterface    usersRepo.UserInterface
	CityInterface    cityRepo.CityInterface
	AuctionInterface auctionsRepo.AuctionInterface
)

func init() {
	UserInterface = usersRepo.NewRepository()
	CityInterface = cityRepo.NewRepository()
	AuctionInterface = auctionsRepo.NewRepository()
}
