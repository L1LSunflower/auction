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

const dateFormat = "2006-01-02 15:04:05"

func CreateAuction(ctx *fiber.Ctx, auction *aggregates.AuctionAggregation) error {
	return ctx.JSON(&structs.CreateAuction{
		Status: successStatus,
		ID:     auction.Auction.ID,
	})
}

func Auction(ctx *fiber.Ctx, auction *aggregates.AuctionAggregation) error {
	var (
		files []string
		tags  []string
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

	return ctx.JSON(&structs.Auction{
		Status:           successStatus,
		ID:               auction.Auction.ID,
		Phone:            auction.User.Phone,
		Category:         auction.Auction.Category,
		WinnerID:         auction.Auction.WinnerID,
		Title:            auction.Item.Name,
		ShortDescription: auction.Auction.ShortDescription,
		Description:      auction.Item.Description,
		StartPrice:       auction.Auction.StartPrice,
		MinimalPrice:     auction.Auction.MinPrice,
		StartDate:        auction.Auction.StartedAt.Format(dateFormat),
		EndedAt:          auction.Auction.EndedAt.Format(dateFormat),
		Files:            files,
		Tags:             tags,
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
			ShortDescription: auction.Auction.ShortDescription,
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
		StartDate:        auction.Auction.StartedAt.Format(dateFormat),
		EndedAt:          auction.Auction.EndedAt.Format(dateFormat),
		Files:            files,
	})
}

func StartAuction(ctx *fiber.Ctx, auction *entities.Auction) error {
	return ctx.JSON(&structs.Start{
		Status: successStatus,
		Date:   time.Now().Format(dateFormat),
	})
}

func EndAuction(ctx *fiber.Ctx, auction *entities.Auction) error {
	return ctx.JSON(&structs.End{
		Status:   successStatus,
		Date:     time.Now().Format(dateFormat),
		WinnerID: auction.WinnerID,
	})
}

func DeleteAuction(ctx *fiber.Ctx, auction *aggregates.AuctionAggregation) error {
	return ctx.JSON(&structs.Delete{
		Status: successStatus,
		Date:   time.Now().Format(dateFormat),
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
