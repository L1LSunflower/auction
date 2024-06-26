package auctions

import (
	"context"
	"errors"
	"fmt"
	"github.com/L1LSunflower/auction/internal/domain/aggregates"
	"github.com/L1LSunflower/auction/internal/domain/entities"
	"github.com/L1LSunflower/auction/internal/domain/repositories/db_repository"
	"github.com/L1LSunflower/auction/internal/domain/services"
	"github.com/L1LSunflower/auction/internal/requests"
	auctionReq "github.com/L1LSunflower/auction/internal/requests/structs/auctions"
	"github.com/L1LSunflower/auction/internal/tools/context_with_depends"
	"github.com/L1LSunflower/auction/internal/tools/errorhandler"
	"github.com/L1LSunflower/auction/internal/tools/metadata"
	"os"
	"time"
)

func Create(ctx context.Context, request *auctionReq.Create) (*aggregates.AuctionAggregation, error) {
	var err error
	if err = context_with_depends.StartDBTx(ctx); err != nil {
		return nil, err
	}
	defer context_with_depends.DBTxRollback(ctx)

	auctionAgg := &aggregates.AuctionAggregation{}

	if auctionAgg.User, err = db_repository.UserInterface.User(ctx, request.OwnerID); err != nil || auctionAgg.User == nil {
		return nil, errorhandler.ErrUserNotExist
	}

	if auctionAgg.User.CreatedAt.IsZero() {
		return nil, errorhandler.ErrUserNotExist
	}

	// Temporary commented
	//if auctionAgg.Auction, err = db_repository.AuctionInterface.ActiveAuction(ctx, request.OwnerID); err != nil {
	//	return nil, errorhandler.InternalError
	//}
	//
	//if auctionAgg.Auction.Status == entities.ActiveStatus {
	//	return nil, errorhandler.ErrActiveAuctionExist
	//}

	//var auctions int
	//if auctions, err = db_repository.AuctionInterface.CountInactiveAuctions(ctx, request.OwnerID); err != nil && auctions >= 10 {
	//	return nil, errorhandler.ErrCreateLimit
	//}

	auctionAgg.CreateItem(request.ItemTitle, request.ItemDescription)
	if err = db_repository.ItemInterface.Create(ctx, auctionAgg.Item); err != nil {
		return nil, errorhandler.ErrCreateItem
	}

	for _, tagName := range request.ItemTags {
		var tag *entities.Tag
		tag, err = db_repository.TagsInterface.ByName(ctx, tagName)
		if err != nil {
			if tag, err = db_repository.TagsInterface.Create(ctx, tagName); err != nil {
				return nil, errorhandler.ErrCreateTag
			}
		}

		if _, err = db_repository.TagsInterface.CreateLink(ctx, auctionAgg.Item.ID, tag.ID); err != nil {
			return nil, errorhandler.ErrCreateTag
		}
	}

	for _, filename := range request.ItemFiles {
		if _, err = db_repository.FilesInterface.CreateLink(ctx, auctionAgg.Item.ID, filename); err != nil {
			return nil, errorhandler.ErrCreateFile
		}
	}

	auctionAgg.CreateAuction(request.StartPrice, request.MinimalPrice, request.Category, request.ShortDescription, request.StartDate)
	if err = db_repository.AuctionInterface.Create(ctx, auctionAgg.Auction); err != nil {
		return nil, errorhandler.ErrCreateAuction
	}

	context_with_depends.DBTxCommit(ctx)
	return auctionAgg, nil
}

func Auction(ctx context.Context, request *auctionReq.Auction) (*aggregates.AuctionAggregation, error) {
	var err error
	auctionAgg := &aggregates.AuctionAggregation{}

	if auctionAgg.User, err = db_repository.UserInterface.User(ctx, request.UserID); err != nil {
		return nil, errorhandler.ErrUserNotExist
	}

	if auctionAgg.Auction, err = db_repository.AuctionInterface.Auction(ctx, request.ID); err != nil {
		return nil, errorhandler.ErrDoesNotExistAuction
	}

	if auctionAgg.Item, err = db_repository.ItemInterface.Item(ctx, auctionAgg.Auction.ItemID); err != nil {
		return nil, errorhandler.ErrDoesNotExistItem
	}

	if auctionAgg.ItemFiles, err = db_repository.FilesInterface.Files(ctx, auctionAgg.Item.ID); err != nil {
		return nil, errorhandler.ErrGetFiles
	}

	if auctionAgg.Tags, err = db_repository.TagsInterface.Tags(ctx, auctionAgg.Item.ID); err != nil {
		return nil, errorhandler.ErrGetTags
	}

	if auctionAgg.OwnerUser, err = db_repository.UserInterface.User(ctx, auctionAgg.Auction.OwnerID); err != nil {
		return nil, errorhandler.ErrUserExist
	}

	member, err := db_repository.AuctionInterface.Member(ctx, auctionAgg.Auction.ID, auctionAgg.User.ID)
	if err != nil || member == nil {
		auctionAgg.Member = false
	}

	if member != nil {
		if len(member.ParticipantID) > 0 {
			auctionAgg.Member = true
		}
	}

	if auctionAgg.Auction.OwnerID == auctionAgg.User.ID {
		if auctionAgg.VisitorsCount, err = db_repository.AuctionInterface.VisitorsCount(ctx, auctionAgg.Auction.ID); err != nil {
			return nil, err
		}
	} else {
		visitor, err := db_repository.AuctionInterface.Visitor(ctx, auctionAgg.Auction.ID, auctionAgg.User.ID)
		if err != nil {
			return nil, errorhandler.ErrAuctionVisitor
		}

		if len(visitor.UserID) > 0 {
			auctionAgg.AuctionVisitor = true
		}
	}

	return auctionAgg, nil
}

func Auctions(ctx context.Context, request *auctionReq.Auctions) ([]*aggregates.AuctionFile, error) {
	var (
		auctionsWithFiles []*aggregates.AuctionFile
		err               error
	)

	request.Metadata.Total, err = db_repository.AuctionInterface.Count(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get count with error: %s", err.Error())
	}

	if request.Metadata.Total <= 0 {
		request.Metadata.CurrentPage = 0
		request.Metadata.PerPage = 0
		return auctionsWithFiles, nil
	}

	if err = services.GetLimitAndOffset(request.Metadata); err != nil {
		return nil, err
	}
	whereString := metadata.ConcatStrings(request.Where, " and ")

	auctions, err := db_repository.AuctionInterface.Auctions(ctx, whereString, request.Tags, request.OrderBy, request.Metadata)
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
		auctionsWithFiles = append(auctionsWithFiles, auctionWithFile)
	}

	return auctionsWithFiles, nil
}

func Update(ctx context.Context, request *auctionReq.Update) (*aggregates.AuctionAggregation, error) {
	var err error
	if err = context_with_depends.StartDBTx(ctx); err != nil {
		return nil, err
	}
	defer context_with_depends.DBTxRollback(ctx)

	auctionAgg := &aggregates.AuctionAggregation{}

	if auctionAgg.User, err = db_repository.UserInterface.User(ctx, requests.UserIDCtx); err != nil {
		return nil, err
	}

	if auctionAgg.Auction, err = db_repository.AuctionInterface.Auction(ctx, request.ID); err != nil {
		return nil, errorhandler.ErrDoesNotExistAuction
	}

	if auctionAgg.Auction.Status != entities.InactiveStatus {
		return nil, errorhandler.ErrUpdateActiveAuction
	}

	if auctionAgg.Item, err = db_repository.ItemInterface.Item(ctx, auctionAgg.Auction.ItemID); err != nil {
		return nil, errorhandler.ErrDoesNotExistItem
	}

	auctionAgg.Item.Name = request.Title
	auctionAgg.Item.Description = request.Description

	if err = db_repository.ItemInterface.Update(ctx, auctionAgg.Item); err != nil {
		return nil, errorhandler.ErrUpdateItem
	}

	tempTags, err := db_repository.TagsInterface.Tags(ctx, auctionAgg.Item.ID)
	if err != nil {
		return nil, errorhandler.ErrGetTags
	}

	tagsMap := services.CreateMapFromTags(tempTags)
	for _, tag := range request.ItemTags {
		if ok := tagsMap[tag]; ok {
			delete(tagsMap, tag)
			request.ItemTags = request.ItemTags[1:]
		}
	}

	var tagEn *entities.Tag
	for tag := range tagsMap {
		if tagEn, err = db_repository.TagsInterface.ByName(ctx, tag); err != nil {
			return nil, errorhandler.ErrGetTags
		}

		if err = db_repository.TagsInterface.DeleteItemTags(ctx, auctionAgg.Item.ID, tagEn.ID); err != nil {
			return nil, errorhandler.ErrDeleteTags
		}
	}

	for _, tagName := range request.ItemTags {
		var tag *entities.Tag
		tag, err = db_repository.TagsInterface.ByName(ctx, tagName)
		if err != nil {
			if tag, err = db_repository.TagsInterface.Create(ctx, tagName); err != nil {
				return nil, errorhandler.ErrCreateTag
			}
		}

		if _, err = db_repository.TagsInterface.CreateLink(ctx, auctionAgg.Item.ID, tag.ID); err != nil {
			return nil, errorhandler.ErrCreateTag
		}
	}

	tempFiles, err := db_repository.FilesInterface.Files(ctx, auctionAgg.Item.ID)
	if err != nil {
		return nil, errorhandler.ErrGetFiles
	}

	filesMap := services.CreateMapFromFiles(tempFiles)
	for _, file := range request.ItemFiles {
		if ok := filesMap[file]; ok {
			delete(filesMap, file)
			request.ItemFiles = request.ItemFiles[1:]
		}
	}

	for file := range filesMap {
		if err = db_repository.FilesInterface.Delete(ctx, auctionAgg.Item.ID, file); err != nil {
			return nil, errorhandler.ErrDeleteFile
		}

		if err = os.Remove(fmt.Sprintf("./static/%s", file)); err != nil {
			return nil, errorhandler.ErrDeleteFile
		}
	}

	for _, filename := range request.ItemFiles {
		if _, err = db_repository.FilesInterface.CreateLink(ctx, auctionAgg.Item.ID, filename); err != nil {
			return nil, errorhandler.ErrCreateFile
		}
	}

	auctionAgg.Auction.StartPrice = request.StartPrice
	auctionAgg.Auction.MinPrice = request.MinimalPrice
	auctionAgg.Auction.ShortDescription = request.ShortDescription
	auctionAgg.Auction.StartedAt = request.StartDate

	if err = db_repository.AuctionInterface.Update(ctx, auctionAgg.Auction); err != nil {
		return nil, errorhandler.ErrCreateAuction
	}

	context_with_depends.DBTxCommit(ctx)
	return auctionAgg, nil
}

func Start(ctx context.Context, request *auctionReq.Start) (*entities.Auction, error) {
	if err := context_with_depends.StartDBTx(ctx); err != nil {
		return nil, err
	}
	defer context_with_depends.DBTxRollback(ctx)

	auction, err := db_repository.AuctionInterface.Auction(ctx, request.ID)
	if err != nil || auction == nil {
		return nil, errorhandler.ErrAuctionNotExist
	}

	if auction.OwnerID != request.UserID {
		return nil, errorhandler.ErrFailedStartAuction
	}

	if auction.Status != entities.InactiveStatus {
		return nil, errorhandler.ErrFailedStartAuction
	}

	if err = db_repository.AuctionInterface.Start(ctx, auction.ID, request.EndedAt); err != nil {
		return nil, errorhandler.ErrFailedStartAuction
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
	if err != nil || auction == nil {
		return nil, errorhandler.ErrAuctionNotExist
	}

	if auction.OwnerID != request.UserID {
		return nil, errorhandler.ErrFailedStartAuction
	}

	if auction.Status != entities.ActiveStatus {
		return nil, errorhandler.ErrEndAuction
	}

	if err := db_repository.AuctionInterface.End(ctx, auction.ID); err != nil {
		return nil, errors.New("failed to end auction")
	}

	context_with_depends.DBTxCommit(ctx)
	return auction, nil
}

func Delete(ctx context.Context, request *auctionReq.Delete) (*aggregates.AuctionAggregation, error) {
	var err error
	if err = context_with_depends.StartDBTx(ctx); err != nil {
		return nil, err
	}
	defer context_with_depends.DBTxRollback(ctx)

	auctionAgg := &aggregates.AuctionAggregation{}

	if auctionAgg.Auction, err = db_repository.AuctionInterface.Auction(ctx, request.ID); err != nil {
		return nil, errorhandler.ErrDoesNotExistAuction
	}

	if auctionAgg.Auction.Status == entities.ActiveStatus || auctionAgg.Auction.Status == entities.CompletedStatus {
		return nil, errorhandler.ErrDeleteByStatus
	}

	if err = db_repository.AuctionInterface.Delete(ctx, auctionAgg.Auction); err != nil {
		return nil, errorhandler.ErrDeleteAuction
	}

	if auctionAgg.Item, err = db_repository.ItemInterface.Item(ctx, auctionAgg.Auction.ItemID); err != nil {
		return nil, errorhandler.ErrDoesNotExistItem
	}

	if err = db_repository.ItemInterface.Delete(ctx, auctionAgg.Item.ID); err != nil {
		return nil, errorhandler.ErrDeleteItem
	}

	if auctionAgg.ItemFiles, err = db_repository.FilesInterface.Files(ctx, auctionAgg.Item.ID); err != nil {
		return nil, errorhandler.ErrGetFiles
	}

	if err = db_repository.FilesInterface.DeleteAll(ctx, auctionAgg.Item.ID); err != nil {
		return nil, errorhandler.ErrDeleteFiles
	}

	for _, filename := range auctionAgg.ItemFiles {
		if err = os.Remove(fmt.Sprintf("./images/%s", filename.Name)); err != nil {
			return nil, errorhandler.ErrDeleteFile
		}
	}

	if err = db_repository.TagsInterface.DeleteItemLinks(ctx, auctionAgg.Item.ID); err != nil {
		return nil, errorhandler.ErrDeleteTags
	}

	context_with_depends.DBTxCommit(ctx)
	return auctionAgg, nil
}

func Participate(ctx context.Context, request *auctionReq.Participate) (*entities.AuctionMember, error) {
	var err error
	if err = context_with_depends.StartDBTx(ctx); err != nil {
		return nil, err
	}
	defer context_with_depends.DBTxRollback(ctx)

	user, err := db_repository.UserInterface.User(ctx, request.UserID)
	if err != nil || user == nil {
		return nil, errorhandler.ErrUserNotExist
	}

	auction, err := db_repository.AuctionInterface.Auction(ctx, request.AuctionID)
	if err != nil || auction == nil {
		return nil, errorhandler.ErrAuctionNotExist
	}

	if auction.Status != entities.ActiveStatus {
		return nil, errorhandler.ErrActiveAuctionNotExist
	}

	auctionMember, err := db_repository.AuctionInterface.Member(ctx, auction.ID, user.ID)
	if err != nil || auctionMember == nil {
		return nil, errorhandler.ErrAddAuctionMember
	}

	if auctionMember.AuctionID > 0 && len(auctionMember.ParticipantID) > 0 {
		return nil, errorhandler.ErrAlreadyAuctionMember
	}

	member := &entities.AuctionMember{
		AuctionID:     auction.ID,
		ParticipantID: user.ID,
		Price:         0,
		FirstName:     user.FirstName,
		LastName:      user.LastName,
	}

	if err = db_repository.AuctionInterface.CreateMember(ctx, member); err != nil {
		return nil, errorhandler.ErrParticipate
	}

	context_with_depends.DBTxCommit(ctx)
	return auctionMember, nil
}

func SetPrice(ctx context.Context, request *auctionReq.SetPrice) (float64, error) {
	var err error
	if err = context_with_depends.StartDBTx(ctx); err != nil {
		return 0, err
	}
	defer context_with_depends.DBTxRollback(ctx)

	user, err := db_repository.UserInterface.User(ctx, request.UserID)
	if err != nil || user == nil {
		return 0, errorhandler.ErrUserNotExist
	}

	auction, err := db_repository.AuctionInterface.Auction(ctx, request.AuctionID)
	if err != nil || auction == nil {
		return 0, errorhandler.ErrAuctionNotExist
	}

	if auction.Price > request.Price {
		return 0, errorhandler.ErrPriceLow
	}

	if err = db_repository.AuctionInterface.SetPrice(ctx, request.AuctionID, user.ID, request.Price); err != nil {
		return 0, errorhandler.ErrSetPrice
	}

	context_with_depends.DBTxCommit(ctx)

	go SendPrice(request.AuctionID, user.ID, request.Price)

	return request.Price, nil
}

func SetVisit(ctx context.Context, request *auctionReq.SetVisit) (*entities.AuctionVisit, error) {
	var err error
	if err = context_with_depends.StartDBTx(ctx); err != nil {
		return nil, err
	}
	defer context_with_depends.DBTxRollback(ctx)

	user, err := db_repository.UserInterface.User(ctx, request.UserID)
	if err != nil || user == nil {
		return nil, errorhandler.ErrUserNotExist
	}

	auction, err := db_repository.AuctionInterface.Auction(ctx, request.AuctionID)
	if err != nil || auction.CreatedAt.IsZero() || auction.VisitStatus != entities.VisitNotSet {
		return nil, errorhandler.ErrAuctionNotExist
	}

	visit := &entities.AuctionVisit{
		AuctionID: auction.ID,
		StartDate: request.StartDate,
		EndDate:   request.EndDate,
	}

	if err = db_repository.AuctionInterface.SetVisit(ctx, visit); err != nil {
		return nil, errorhandler.ErrSetVisit
	}

	context_with_depends.DBTxCommit(ctx)

	return visit, nil
}

func Visit(ctx context.Context, request *auctionReq.Visit) error {
	var err error
	if err = context_with_depends.StartDBTx(ctx); err != nil {
		return err
	}
	defer context_with_depends.DBTxRollback(ctx)

	user, err := db_repository.UserInterface.User(ctx, request.UserID)
	if err != nil || user == nil {
		return errorhandler.ErrUserNotExist
	}

	auction, err := db_repository.AuctionInterface.Auction(ctx, request.AuctionID)
	if err != nil || auction.CreatedAt.IsZero() || auction.VisitStatus != entities.VisitSet && auction.VisitStatus != entities.VisitOpened {
		return errorhandler.ErrAuctionNotExist
	}

	visitor, err := db_repository.AuctionInterface.Visitor(ctx, auction.ID, user.ID)
	if err != nil {
		return errorhandler.ErrAuctionVisitor
	}

	if len(visitor.UserID) > 0 {
		return errorhandler.ErrAlreadyVisitor
	}

	auctionVisitor := &entities.AuctionVisitor{
		AuctionID: auction.ID,
		UserID:    user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Phone:     user.Phone,
	}

	if err = db_repository.AuctionInterface.Visit(ctx, auctionVisitor); err != nil {
		return errorhandler.ErrVisit
	}

	context_with_depends.DBTxCommit(ctx)

	return nil
}

func Unvisit(ctx context.Context, request *auctionReq.Unvisit) error {
	var err error
	if err = context_with_depends.StartDBTx(ctx); err != nil {
		return err
	}
	defer context_with_depends.DBTxRollback(ctx)

	auctionVisitor, err := db_repository.AuctionInterface.Visitor(ctx, request.AuctionID, request.UserID)
	if err != nil || auctionVisitor == nil {
		return errorhandler.ErrUnvisit
	}

	if len(auctionVisitor.UserID) <= 0 {
		return errorhandler.ErrUnvisit
	}

	if err = db_repository.AuctionInterface.Unvisit(ctx, request.AuctionID, request.UserID); err != nil {
		return errorhandler.ErrUnvisit
	}

	context_with_depends.DBTxCommit(ctx)

	return nil
}

func Visitors(ctx context.Context, request *auctionReq.Visitor) ([]*entities.AuctionVisitor, error) {
	var err error

	user, err := db_repository.UserInterface.User(ctx, request.UserID)
	if err != nil || user == nil {
		return nil, errorhandler.ErrUserNotExist
	}

	if len(user.ID) <= 0 {
		return nil, errorhandler.ErrUserNotExist
	}

	auction, err := db_repository.AuctionInterface.Auction(ctx, request.AuctionID)
	if err != nil || auction == nil {
		return nil, errorhandler.ErrAuctionNotExist
	}

	if auction.ID <= 0 {
		return nil, errorhandler.ErrAuctionNotExist
	}

	if user.ID != auction.OwnerID {
		return nil, errorhandler.ErrNotAuctionOwner
	}

	auctionVisitors, err := db_repository.AuctionInterface.Visitors(ctx, request.AuctionID)
	if err != nil {
		return nil, errorhandler.ErrAuctionVisitors
	}

	return auctionVisitors, nil
}

func UpdateVisit(ctx context.Context, request *auctionReq.UpdateVisit) (*entities.AuctionVisit, error) {
	var err error
	if err = context_with_depends.StartDBTx(ctx); err != nil {
		return nil, err
	}
	defer context_with_depends.DBTxRollback(ctx)

	user, err := db_repository.UserInterface.User(ctx, request.UserID)
	if err != nil || user == nil {
		return nil, errorhandler.ErrUserNotExist
	}

	if len(user.ID) <= 0 {
		return nil, errorhandler.ErrUserNotExist
	}

	auction, err := db_repository.AuctionInterface.Auction(ctx, request.AuctionID)
	if err != nil || auction == nil {
		return nil, errorhandler.ErrAuctionNotExist
	}

	if auction.ID <= 0 {
		return nil, errorhandler.ErrAuctionNotExist
	}

	if user.ID != auction.OwnerID {
		return nil, errorhandler.ErrNotAuctionOwner
	}

	if auction.VisitStatus == entities.VisitOpened || auction.VisitStatus == entities.VisitClosed {
		return nil, errorhandler.ErrUpdateVisit
	}

	auction.VisitStartDate = request.StartDate
	auction.VisitEndDate = request.EndDate

	err = db_repository.AuctionInterface.Update(ctx, auction)
	if err != nil {
		return nil, errorhandler.ErrUpdateAuction
	}

	context_with_depends.DBTxCommit(ctx)

	return &entities.AuctionVisit{
		AuctionID: auction.ID,
		StartDate: request.StartDate,
		EndDate:   request.EndDate,
	}, nil
}

func SendPrice(auctionID int, userID string, price float64) {
	for {
		if !CheckEvent(auctionID) {
			break
		}
		time.Sleep(1 * time.Second)
	}

	if CheckAuction(auctionID) {
		RegisterNewEvent(auctionID, userID, price)
		for {
			if DataSent(auctionID) {
				DeleteEvent(auctionID)
				break
			}
			time.Sleep(3 * time.Second)
		}
	}
}
