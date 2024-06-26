package users

import (
	"context"
	"database/sql"
	"github.com/L1LSunflower/auction/internal/domain/aggregates"
	"github.com/L1LSunflower/auction/internal/domain/entities"
	"github.com/L1LSunflower/auction/internal/domain/repositories/db_repository"
	userRequest "github.com/L1LSunflower/auction/internal/requests/structs/users"
	"github.com/L1LSunflower/auction/internal/tools/context_with_depends"
	"github.com/L1LSunflower/auction/internal/tools/errorhandler"
)

func User(ctx context.Context, request *userRequest.User) (*entities.User, error) {
	user, err := db_repository.UserInterface.User(ctx, request.ID)
	if err != nil && err != sql.ErrNoRows {
		return nil, errorhandler.InternalError
	}

	if user == nil || user.CreatedAt.IsZero() {
		return nil, errorhandler.ErrUserExist
	}

	return user, nil
}

func Update(ctx context.Context, request *userRequest.Update) (*entities.User, error) {
	if err := context_with_depends.StartDBTx(ctx); err != nil {
		return nil, err
	}
	defer context_with_depends.DBTxRollback(ctx)

	user, err := db_repository.UserInterface.User(ctx, request.ID)
	if err != nil || user == nil {
		return nil, errorhandler.ErrUserNotExist
	}

	if len(request.Email) > 0 {
		user.Email = request.Email
	}

	if len(request.FirstName) > 0 {
		user.FirstName = request.FirstName
	}

	if len(request.LastName) > 0 {
		user.LastName = request.LastName
	}

	if len(request.Password) > 0 {
		user.Password = request.Password
	}

	if err = db_repository.UserInterface.Update(ctx, user); err != nil {
		return nil, errorhandler.ErrUpdateUser
	}

	context_with_depends.DBTxCommit(ctx)

	return user, nil
}

func Delete(ctx context.Context, request *userRequest.Delete) (*entities.User, error) {
	if err := context_with_depends.StartDBTx(ctx); err != nil {
		return nil, err
	}
	defer context_with_depends.DBTxRollback(ctx)

	user, err := db_repository.UserInterface.User(ctx, request.ID)
	if err != nil || user == nil {
		return nil, err
	}

	if err = db_repository.AuctionInterface.DeleteByOwnerID(ctx, user.ID); err != nil {
		return nil, errorhandler.ErrDeleteAuctions
	}

	if err = db_repository.UserInterface.Delete(ctx, request.ID); err != nil {
		return nil, err
	}

	context_with_depends.DBTxCommit(ctx)

	return user, nil
}

func Profile(ctx context.Context, request *userRequest.User) (*aggregates.ProfileAggregation, error) {
	var err error
	userProfile := &aggregates.ProfileAggregation{}

	if userProfile.User, err = db_repository.UserInterface.User(ctx, request.ID); err != nil || userProfile.User == nil {
		return nil, errorhandler.ErrUserNotExist
	}

	if userProfile.Balance, err = db_repository.BalanceInterface.Balance(ctx, userProfile.User.ID); err != nil || userProfile.Balance == nil {
		return nil, errorhandler.ErrGetBalance
	}

	auctions, err := db_repository.AuctionInterface.ActiveAuctions(ctx, userProfile.User.ID)
	if err != nil {
		return nil, errorhandler.ErrGetAuctions
	}

	for _, auction := range auctions {
		var files []*entities.File
		auctionWithFile := &aggregates.AuctionFile{}

		if files, err = db_repository.FilesInterface.Files(ctx, auction.ItemID); err != nil {
			return nil, err
		}

		auctionWithFile.Auction = &entities.Auction{ID: auction.ID, ShortDescription: auction.ShortDescription, Status: auction.Status, Category: auction.Category}

		for _, file := range files {
			auctionWithFile.Files = append(auctionWithFile.Files, file.Name)
		}
		userProfile.Auctions = append(userProfile.Auctions, auctionWithFile)
	}

	return userProfile, nil
}

func OwnersAuctions(ctx context.Context, request *userRequest.User) (*aggregates.ProfileHistoryAggregation, error) {
	var err error
	userProfile := &aggregates.ProfileHistoryAggregation{}

	if userProfile.User, err = db_repository.UserInterface.User(ctx, request.ID); err != nil || userProfile.User == nil {
		return nil, errorhandler.ErrUserNotExist
	}

	auctions, err := db_repository.AuctionInterface.ByOwnerID(ctx, userProfile.User.ID)
	if err != nil {
		return nil, errorhandler.ErrGetAuctions
	}

	for _, auction := range auctions {
		var files []*entities.File
		auctionWithFile := &aggregates.AuctionFile{}

		if files, err = db_repository.FilesInterface.Files(ctx, auction.ItemID); err != nil {
			return nil, err
		}

		auctionWithFile.Auction = &entities.Auction{ID: auction.ID, ShortDescription: auction.ShortDescription, Status: auction.Status, Category: auction.Category}

		for _, file := range files {
			auctionWithFile.Files = append(auctionWithFile.Files, file.Name)
		}
		userProfile.Auctions = append(userProfile.Auctions, auctionWithFile)
	}

	return userProfile, nil
}

func CompletedAuctions(ctx context.Context, request *userRequest.User) (*aggregates.ProfileHistoryAggregation, error) {
	var err error
	userProfile := &aggregates.ProfileHistoryAggregation{}

	if userProfile.User, err = db_repository.UserInterface.User(ctx, request.ID); err != nil || userProfile.User == nil {
		return nil, errorhandler.ErrUserNotExist
	}

	auctions, err := db_repository.AuctionInterface.Completed(ctx, userProfile.User.ID)
	if err != nil {
		return nil, errorhandler.ErrGetAuctions
	}

	for _, auction := range auctions {
		var files []*entities.File
		auctionWithFile := &aggregates.AuctionFile{}

		if files, err = db_repository.FilesInterface.Files(ctx, auction.ItemID); err != nil {
			return nil, err
		}

		auctionWithFile.Auction = &entities.Auction{ID: auction.ID, ShortDescription: auction.ShortDescription, Status: auction.Status, Category: auction.Category}

		for _, file := range files {
			auctionWithFile.Files = append(auctionWithFile.Files, file.Name)
		}
		userProfile.Auctions = append(userProfile.Auctions, auctionWithFile)
	}

	return userProfile, nil
}
