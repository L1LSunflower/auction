package responses

import (
	"github.com/L1LSunflower/auction/internal/domain/entities"
	"github.com/L1LSunflower/auction/internal/tools/metadata"
	"github.com/gofiber/fiber/v2"
	"strings"
	"time"

	"github.com/L1LSunflower/auction/internal/domain/aggregates"
	"github.com/L1LSunflower/auction/internal/responses/structs"
)

func CreateAuction(ctx *fiber.Ctx, auction *aggregates.AuctionAggregation) error {
	return ctx.JSON(&structs.CreateAuction{
		Status: successStatus,
		ID:     auction.Auction.ID,
	})
}

func Auction(ctx *fiber.Ctx, auction *aggregates.AuctionAggregation) error {
	var (
		files          []string
		tags           []string
		visitStartDate string
		visitEndDate   string
	)

	if len(auction.ItemFiles) > 0 {
		for _, file := range auction.ItemFiles {
			files = append(files, file.Name)
		}
	}

	if len(auction.Tags) > 0 {
		for _, tag := range auction.Tags {
			tags = append(tags, tag.Name)
		}
	}

	if !auction.Auction.VisitStartDate.IsZero() && !auction.Auction.VisitEndDate.IsZero() {
		visitStartDate = auction.Auction.VisitStartDate.Format(entities.DateFormat)
		visitEndDate = auction.Auction.VisitEndDate.Format(entities.DateFormat)
	}

	return ctx.JSON(&structs.Auction{
		Status:           successStatus,
		Member:           auction.Member,
		ID:               auction.Auction.ID,
		AuctionStatus:    auction.Auction.Status,
		Phone:            auction.OwnerUser.Phone,
		Category:         auction.Auction.Category,
		WinnerID:         auction.Auction.WinnerID,
		Title:            auction.Item.Name,
		ShortDescription: auction.Auction.ShortDescription,
		Description:      auction.Item.Description,
		StartPrice:       auction.Auction.StartPrice,
		MinimalPrice:     auction.Auction.MinPrice,
		StartDate:        auction.Auction.StartedAt.Format(entities.DateFormat),
		EndedAt:          auction.Auction.EndedAt.Format(entities.DateFormat),
		Files:            files,
		Tags:             tags,
		Visitor:          auction.AuctionVisitor,
		VisitStartDate:   visitStartDate,
		VisitEndDate:     visitEndDate,
		VisitorsCount:    auction.VisitorsCount,
	})
}

func Auctions(ctx *fiber.Ctx, auctions []*aggregates.AuctionFile, metadata *metadata.Metadata) error {
	auctionsWithFiles := &structs.AuctionsWithFile{Status: successStatus, CurrentPage: metadata.CurrentPage, Total: metadata.Total, LastPage: metadata.LastPage}
	if len(auctions) <= 0 {
		return ctx.JSON(auctionsWithFiles)
	}

	for _, auction := range auctions {
		file := GetFirstVideoOrImage(auction.Files)
		auctionsWithFiles.Auctions = append(auctionsWithFiles.Auctions, structs.AuctionWithFile{
			ID:               auction.Auction.ID,
			Status:           auction.Auction.Status,
			AuctionStatus:    auction.Auction.Status,
			ShortDescription: auction.Auction.ShortDescription,
			Category:         auction.Auction.Category,
			Files:            file,
		})
	}

	return ctx.JSON(auctionsWithFiles)
}

func UpdateAuction(ctx *fiber.Ctx, auction *aggregates.AuctionAggregation) error {
	var files []string
	if len(auction.ItemFiles) > 0 {
		for _, file := range auction.ItemFiles {
			files = append(files, file.Name)
		}
	}

	return ctx.JSON(&structs.Update{
		Status:           successStatus,
		ID:               auction.Auction.ID,
		Category:         auction.Auction.Category,
		WinnerID:         auction.Auction.WinnerID,
		Title:            auction.Item.Name,
		ShortDescription: auction.Auction.ShortDescription,
		Description:      auction.Item.Description,
		AuctionStatus:    auction.Auction.Status,
		StartPrice:       auction.Auction.StartPrice,
		MinimalPrice:     auction.Auction.MinPrice,
		StartDate:        auction.Auction.StartedAt.Format(entities.DateFormat),
		EndedAt:          auction.Auction.EndedAt.Format(entities.DateFormat),
		Files:            files,
	})
}

func StartAuction(ctx *fiber.Ctx, auction *entities.Auction) error {
	return ctx.JSON(&structs.Start{
		Status: successStatus,
		Date:   time.Now().Format(entities.DateFormat),
	})
}

func EndAuction(ctx *fiber.Ctx, auction *entities.Auction) error {
	return ctx.JSON(&structs.End{
		Status:   successStatus,
		Date:     time.Now().Format(entities.DateFormat),
		WinnerID: auction.WinnerID,
	})
}

func DeleteAuction(ctx *fiber.Ctx, auction *aggregates.AuctionAggregation) error {
	return ctx.JSON(&structs.Delete{
		Status: successStatus,
		Date:   time.Now().Format(entities.DateFormat),
	})
}

func Participate(ctx *fiber.Ctx, auctionMember *entities.AuctionMember) error {
	return ctx.JSON(&structs.Participate{
		Status: successStatus,
		Date:   time.Now().Format(entities.DateFormat),
	})
}

func SetPrice(ctx *fiber.Ctx, price float64) error {
	return ctx.JSON(&structs.SetPrice{
		Status: successStatus,
		Price:  price,
		Date:   time.Now().Format(entities.DateFormat),
	})
}

func GetFirstVideoOrImage(files []string) string {
	if len(files) <= 0 {
		return ""
	}

	for _, file := range files {
		if splitedFile := strings.Split(file, "."); len(splitedFile) > 1 && splitedFile[len(splitedFile)-1] == "mp4" {
			return file
		}
	}

	return files[0]
}

func SettedVisit(ctx *fiber.Ctx, settedVisit *entities.AuctionVisit) error {
	return ctx.JSON(&structs.SettedVisit{
		Status:    successStatus,
		StartDate: settedVisit.StartDate.Format(entities.DateFormat),
		EndDate:   settedVisit.EndDate.Format(entities.DateFormat),
	})
}

func Visit(ctx *fiber.Ctx) error {
	return ctx.JSON(&structs.SuccessResponse{
		Status: successStatus,
	})
}

func Unvisit(ctx *fiber.Ctx) error {
	return ctx.JSON(&structs.SuccessResponse{
		Status: successStatus,
	})
}

func Visitors(ctx *fiber.Ctx, visitors []*entities.AuctionVisitor) error {
	response := &structs.Visitors{Status: successStatus, Count: len(visitors)}
	for _, visitor := range visitors {
		response.AuctionVisitors = append(response.AuctionVisitors, &structs.Visitor{
			FirstName: visitor.FirstName,
			LastName:  visitor.LastName,
			Phone:     visitor.Phone,
		})
	}

	return ctx.JSON(response)
}

func UpdateVisit(ctx *fiber.Ctx, updatedVisit *entities.AuctionVisit) error {
	return ctx.JSON(&structs.UpdateVisit{
		Status:    successStatus,
		StartDate: updatedVisit.StartDate.Format(entities.DateFormat),
		EndDate:   updatedVisit.EndDate.Format(entities.DateFormat),
	})
}

func AuctionMembers(members []*entities.AuctionMember, userAndPrice map[string]any) fiber.Map {
	var (
		membersSlice []fiber.Map
		userID       string
		price        float64
		ok           bool
	)

	for _, member := range members {
		membersSlice = append(membersSlice, fiber.Map{
			"price":      member.Price,
			"first_name": member.FirstName,
			"last_name":  member.LastName,
		})
	}

	userID, ok = userAndPrice["user_id"].(string)
	if !ok {
		userID = ""
	}

	price, ok = userAndPrice["price"].(float64)
	if !ok {
		price = 0
	}

	return fiber.Map{
		"status":        successStatus,
		"user_id":       userID,
		"price":         price,
		"members_count": len(membersSlice),
		"members":       membersSlice,
	}
}
