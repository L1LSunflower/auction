package auctions

import (
	"context"
	"errors"
	"fmt"
	"github.com/L1LSunflower/auction/internal/domain/aggregates"
	"github.com/L1LSunflower/auction/internal/domain/entities"
	"github.com/L1LSunflower/auction/internal/domain/repositories/db_repository"
	"github.com/L1LSunflower/auction/internal/domain/services"
	auctionReq "github.com/L1LSunflower/auction/internal/requests/structs/auctions"
	"github.com/L1LSunflower/auction/internal/tools/context_with_depends"
	"github.com/L1LSunflower/auction/internal/tools/metadata"
)

func Create(ctx context.Context, request *auctionReq.Create) (*entities.Auction, error) {
	if err := context_with_depends.StartDBTx(ctx); err != nil {
		return nil, err
	}
	defer context_with_depends.DBTxRollback(ctx)

	auction, err := db_repository.AuctionInterface.ByOwnerID(ctx, request.OwnerID)
	if err != nil {
		return nil, err
	}

	if !auction.CreatedAt.IsZero() {
		return nil, errors.New("auction limit on auction owner id")
	}

	item := &entities.Item{
		UserID:   request.OwnerID,
		Category: request.Category,
		Name:     request.Name,
		Tag1:     request.Tag1,
		Tag2:     request.Tag2,
		Tag3:     request.Tag3,
		Tag4:     request.Tag4,
		Tag5:     request.Tag5,
		Tag6:     request.Tag6,
		Tag7:     request.Tag7,
		Tag8:     request.Tag8,
		Tag9:     request.Tag9,
		Tag10:    request.Tag10,
		Images:   request.Images,
	}
	if err = db_repository.ItemInterface.Create(ctx, item); err != nil {
		return nil, errors.New("failed to create item for auction")
	}

	auction = &entities.Auction{
		OwnerID:     request.OwnerID,
		ItemID:      item.ID,
		Title:       request.Title,
		Description: request.Description,
		Status:      "inactive",
		StartPrice:  request.StartPrice,
		MinPrice:    request.MinimalPrice,
	}
	if err = db_repository.AuctionInterface.Create(ctx, auction); err != nil {
		return nil, err
	}

	context_with_depends.DBTxCommit(ctx)
	return auction, nil
}

func Auction(ctx context.Context, request *auctionReq.Auction) (*aggregates.AuctionItem, error) {
	var err error
	aItem := aggregates.NewAuctionItem()

	if aItem.Auction, err = db_repository.AuctionInterface.Auction(ctx, request.ID); err != nil {
		return nil, err
	}

	if aItem.Auction.CreatedAt.IsZero() {
		return nil, errors.New("that auction does not exist")
	}

	if aItem.Item, err = db_repository.ItemInterface.Item(ctx, aItem.Auction.ItemID); err != nil {
		return nil, err
	}

	return aItem, nil
}

func Auctions(ctx context.Context, request *auctionReq.Auctions, mdata *metadata.Metadata, where []string) (*aggregates.AuctionsItem, error) {
	var err error

	mdata.Total, err = db_repository.AuctionInterface.Count(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get count with error: %s", err.Error())
	}

	if mdata.Total <= 0 {
		mdata.CurrentPage = 0
		mdata.PerPage = 0
		return &aggregates.AuctionsItem{}, nil
	}

	if err = services.GetLimitAndOffset(mdata); err != nil {
		return nil, err
	}

	auctions, err := db_repository.AuctionInterface.Auctions(ctx, where, request.Metadata)
	if err != nil {
		return nil, err
	}

	return auctions, nil
}

func Update(ctx context.Context, request *auctionReq.Update) (*aggregates.AuctionItem, error) {
	if err := context_with_depends.StartDBTx(ctx); err != nil {
		return nil, err
	}
	defer context_with_depends.DBTxRollback(ctx)

	var err error
	aItem := aggregates.NewAuctionItem()

	if aItem.Auction, err = db_repository.AuctionInterface.Auction(ctx, request.ID); err != nil {
		return nil, err
	}

	if aItem.Auction.CreatedAt.IsZero() {
		return nil, errors.New("that auction does not exist")
	}

	if aItem.Item, err = db_repository.ItemInterface.Item(ctx, aItem.Auction.ItemID); err != nil || aItem.Item.CreatedAt.IsZero() {
		return nil, errors.New("auction item does not exist")
	}

	aItem.Item.Tag1 = request.Tag1
	aItem.Item.Tag2 = request.Tag2
	aItem.Item.Tag3 = request.Tag3
	aItem.Item.Tag4 = request.Tag4
	aItem.Item.Tag5 = request.Tag5
	aItem.Item.Tag6 = request.Tag6
	aItem.Item.Tag7 = request.Tag7
	aItem.Item.Tag8 = request.Tag8
	aItem.Item.Tag9 = request.Tag9
	aItem.Item.Tag10 = request.Tag10
	aItem.Item.Images = request.Images
	aItem.Item.Description = request.ItemDescription

	if err = db_repository.ItemInterface.Update(ctx, aItem.Item); err != nil {
		return nil, errors.New("failed to update auction item")
	}

	aItem.Auction.WinnerID = request.WinnerID
	aItem.Auction.Title = request.Title
	aItem.Auction.Description = request.AuctionDescription
	aItem.Auction.StartPrice = request.StartPrice
	aItem.Auction.MinPrice = request.MinimalPrice
	aItem.Auction.Status = request.Status

	if err = db_repository.AuctionInterface.Update(ctx, aItem.Auction); err != nil {
		return nil, errors.New("failed to update auction")
	}

	context_with_depends.DBTxCommit(ctx)
	return aItem, nil
}

func Start(ctx context.Context, request *auctionReq.Start) (*entities.Auction, error) {
	if err := context_with_depends.StartDBTx(ctx); err != nil {
		return nil, err
	}
	defer context_with_depends.DBTxRollback(ctx)

	auction, err := db_repository.AuctionInterface.Auction(ctx, request.ID)
	if err != nil {
		return nil, err
	}

	if auction.CreatedAt.IsZero() {
		return nil, errors.New("that auction does not exist")
	}

	if err = db_repository.AuctionInterface.Start(ctx, auction); err != nil {
		return nil, errors.New("failed to start auction")
	}

	context_with_depends.DBTxCommit(ctx)
	return auction, nil
}

func End(ctx context.Context, request *auctionReq.End) (*entities.Auction, error) {
	if err := context_with_depends.StartDBTx(ctx); err != nil {
		return nil, err
	}
	defer context_with_depends.DBTxRollback(ctx)

	auction, err := db_repository.AuctionInterface.Auction(ctx, request.ID)
	if err != nil {
		return nil, err
	}

	if auction.CreatedAt.IsZero() {
		return nil, errors.New("that auction does not exist")
	}

	if err = db_repository.AuctionInterface.End(ctx, auction); err != nil {
		return nil, errors.New("failed to end auction")
	}

	context_with_depends.DBTxCommit(ctx)
	return auction, nil
}

func Delete(ctx context.Context, request *auctionReq.Delete) (*entities.Auction, error) {
	if err := context_with_depends.StartDBTx(ctx); err != nil {
		return nil, err
	}
	defer context_with_depends.DBTxRollback(ctx)

	auction, err := db_repository.AuctionInterface.Auction(ctx, request.ID)
	if err != nil {
		return nil, err
	}

	if auction.CreatedAt.IsZero() {
		return nil, errors.New("that auction does not exist")
	}

	if err = db_repository.ItemInterface.Delete(ctx, auction.ItemID); err != nil {
		return nil, errors.New("failed to delete item")
	}

	if err = db_repository.AuctionInterface.End(ctx, auction); err != nil {
		return nil, errors.New("failed to end auction")
	}

	context_with_depends.DBTxCommit(ctx)
	return auction, nil
}
