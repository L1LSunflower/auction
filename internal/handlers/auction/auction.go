package auction

import (
	"context"
	"github.com/L1LSunflower/auction/config"
	"github.com/L1LSunflower/auction/internal/tools/context_with_depends"
	"github.com/L1LSunflower/auction/internal/tools/metadata"
	"github.com/L1LSunflower/auction/pkg/db"
	"github.com/L1LSunflower/auction/pkg/redisdb"
	"github.com/gofiber/fiber/v2"

	auctionService "github.com/L1LSunflower/auction/internal/domain/services/auctions"
	"github.com/L1LSunflower/auction/internal/requests"
	auctionReq "github.com/L1LSunflower/auction/internal/requests/structs/auctions"
	"github.com/L1LSunflower/auction/internal/responses"
	"github.com/L1LSunflower/auction/internal/tools/errorhandler"
)

func Create(ctx *fiber.Ctx) error {
	request, ok := ctx.Locals(requests.RequestKey).(*auctionReq.Create)
	if !ok {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}

	dbConn := db.SqlInstance(config.GetConfig().DB.DBDriver, config.GetConfig().DB.DBString).DB
	redisConn := redisdb.RedisInstance().RedisClient
	contxt, err := context_with_depends.ContextWithDepends(context.Background(), dbConn, redisConn)
	auction, err := auctionService.Create(contxt, request)
	if err != nil {
		return responses.NewFailedResponse(ctx, err)
	}

	return responses.NewSuccessResponse(ctx, auction)
}

func Auction(ctx *fiber.Ctx) error {
	request, ok := ctx.Locals(requests.RequestKey).(*auctionReq.Auction)
	if !ok {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}

	dbConn := db.SqlInstance(config.GetConfig().DB.DBDriver, config.GetConfig().DB.DBString).DB
	redisConn := redisdb.RedisInstance().RedisClient
	contxt, err := context_with_depends.ContextWithDepends(context.Background(), dbConn, redisConn)
	auction, err := auctionService.Auction(contxt, request)
	if err != nil {
		return responses.NewFailedResponse(ctx, err)
	}

	return responses.NewSuccessResponse(ctx, auction)
}

func Auctions(ctx *fiber.Ctx) error {
	request, ok := ctx.Locals(requests.RequestKey).(*auctionReq.Auctions)
	if !ok {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}

	mdata, err := metadata.GetParams(ctx)
	if err != nil {
		return responses.NewFailedResponse(ctx, err)
	}

	where, _ := metadata.Filter(ctx)

	dbConn := db.SqlInstance(config.GetConfig().DB.DBDriver, config.GetConfig().DB.DBString).DB
	redisConn := redisdb.RedisInstance().RedisClient
	contxt, err := context_with_depends.ContextWithDepends(context.Background(), dbConn, redisConn)

	auctions, err := auctionService.Auctions(contxt, request, mdata, where)
	if err != nil {
		return responses.NewFailedResponse(ctx, err)
	}

	return responses.NewSuccessResponse(ctx, auctions)
}

func Update(ctx *fiber.Ctx) error {
	request, ok := ctx.Locals(requests.RequestKey).(*auctionReq.Update)
	if !ok {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}

	dbConn := db.SqlInstance(config.GetConfig().DB.DBDriver, config.GetConfig().DB.DBString).DB
	redisConn := redisdb.RedisInstance().RedisClient
	contxt, err := context_with_depends.ContextWithDepends(context.Background(), dbConn, redisConn)
	auction, err := auctionService.Update(contxt, request)
	if err != nil {
		return responses.NewFailedResponse(ctx, err)
	}

	return responses.NewSuccessResponse(ctx, auction)
}

func Start(ctx *fiber.Ctx) error {
	request, ok := ctx.Locals(requests.RequestKey).(*auctionReq.Start)
	if !ok {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}

	auction, err := auctionService.Start(context.Background(), request)
	if err != nil {
		return responses.NewFailedResponse(ctx, err)
	}

	return responses.NewSuccessResponse(ctx, auction)
}

func End(ctx *fiber.Ctx) error {
	request, ok := ctx.Locals(requests.RequestKey).(*auctionReq.End)
	if !ok {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}

	dbConn := db.SqlInstance(config.GetConfig().DB.DBDriver, config.GetConfig().DB.DBString).DB
	redisConn := redisdb.RedisInstance().RedisClient
	contxt, err := context_with_depends.ContextWithDepends(context.Background(), dbConn, redisConn)
	auction, err := auctionService.End(contxt, request)
	if err != nil {
		return responses.NewFailedResponse(ctx, err)
	}

	return responses.NewSuccessResponse(ctx, auction)
}

func Delete(ctx *fiber.Ctx) error {
	request, ok := ctx.Locals(requests.RequestKey).(*auctionReq.Delete)
	if !ok {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}

	dbConn := db.SqlInstance(config.GetConfig().DB.DBDriver, config.GetConfig().DB.DBString).DB
	redisConn := redisdb.RedisInstance().RedisClient
	contxt, err := context_with_depends.ContextWithDepends(context.Background(), dbConn, redisConn)
	auction, err := auctionService.Delete(contxt, request)
	if err != nil {
		return responses.NewFailedResponse(ctx, err)
	}

	return responses.NewSuccessResponse(ctx, auction)
}
