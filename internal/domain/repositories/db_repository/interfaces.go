package db_repository

import (
	auctionsRepo "github.com/L1LSunflower/auction/internal/domain/repositories/db_repository/auctions"
	itemRepo "github.com/L1LSunflower/auction/internal/domain/repositories/db_repository/items"
	usersRepo "github.com/L1LSunflower/auction/internal/domain/repositories/db_repository/users"
)

var (
	UserInterface    usersRepo.UserInterface
	AuctionInterface auctionsRepo.AuctionInterface
	ItemInterface    itemRepo.ItemInterface
)

func init() {
	UserInterface = usersRepo.NewRepository()
	AuctionInterface = auctionsRepo.NewRepository()
	ItemInterface = itemRepo.GetItemInterface()
}
