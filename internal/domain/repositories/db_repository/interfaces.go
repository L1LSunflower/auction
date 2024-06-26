package db_repository

import (
	auctionsRepo "github.com/L1LSunflower/auction/internal/domain/repositories/db_repository/auctions"
	balancesRepo "github.com/L1LSunflower/auction/internal/domain/repositories/db_repository/balances"
	filesRepo "github.com/L1LSunflower/auction/internal/domain/repositories/db_repository/files"
	itemRepo "github.com/L1LSunflower/auction/internal/domain/repositories/db_repository/items"
	tagsRepo "github.com/L1LSunflower/auction/internal/domain/repositories/db_repository/tags"
	transactionsRepo "github.com/L1LSunflower/auction/internal/domain/repositories/db_repository/transactions"
	usersRepo "github.com/L1LSunflower/auction/internal/domain/repositories/db_repository/users"
)

var (
	UserInterface         usersRepo.UserInterface
	AuctionInterface      auctionsRepo.AuctionInterface
	ItemInterface         itemRepo.ItemInterface
	TagsInterface         tagsRepo.TagsInterface
	FilesInterface        filesRepo.FilesInterface
	BalanceInterface      balancesRepo.BalanceInterface
	TransactionsInterface transactionsRepo.TransactionsInterface
)

func init() {
	UserInterface = usersRepo.NewRepository()
	AuctionInterface = auctionsRepo.NewRepository()
	ItemInterface = itemRepo.GetItemInterface()
	TagsInterface = tagsRepo.GetTagsInterface()
	FilesInterface = filesRepo.GetFilesInterface()
	BalanceInterface = balancesRepo.GetBalanceInterface()
	TransactionsInterface = transactionsRepo.GetTransactionsInterface()
}
