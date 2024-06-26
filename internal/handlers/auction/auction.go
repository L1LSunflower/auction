package auction

import (
	"context"
	"github.com/L1LSunflower/auction/config"
	"github.com/L1LSunflower/auction/internal/tools/context_with_depends"
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
	if err != nil {
		return responses.NewFailedResponse(ctx, errorhandler.ErrDependency)
	}

	auction, err := auctionService.Create(contxt, request)
	if err != nil {
		return responses.NewFailedResponse(ctx, err)
	}

	return responses.CreateAuction(ctx, auction)
}

func Auction(ctx *fiber.Ctx) error {
	request, ok := ctx.Locals(requests.RequestKey).(*auctionReq.Auction)
	if !ok {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}

	dbConn := db.SqlInstance(config.GetConfig().DB.DBDriver, config.GetConfig().DB.DBString).DB
	redisConn := redisdb.RedisInstance().RedisClient

	contxt, err := context_with_depends.ContextWithDepends(context.Background(), dbConn, redisConn)
	if err != nil {
		return responses.NewFailedResponse(ctx, errorhandler.ErrDependency)
	}

	auction, err := auctionService.Auction(contxt, request)
	if err != nil {
		return responses.NewFailedResponse(ctx, err)
	}

	return responses.Auction(ctx, auction)
}

func Auctions(ctx *fiber.Ctx) error {
	request, ok := ctx.Locals(requests.RequestKey).(*auctionReq.Auctions)
	if !ok {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}

	dbConn := db.SqlInstance(config.GetConfig().DB.DBDriver, config.GetConfig().DB.DBString).DB
	redisConn := redisdb.RedisInstance().RedisClient

	contxt, err := context_with_depends.ContextWithDepends(context.Background(), dbConn, redisConn)
	if err != nil {
		return responses.NewFailedResponse(ctx, errorhandler.ErrDependency)
	}

	auctions, err := auctionService.Auctions(contxt, request)
	if err != nil {
		return responses.NewFailedResponse(ctx, err)
	}

	return responses.Auctions(ctx, auctions, request.Metadata)
}

func Update(ctx *fiber.Ctx) error {
	request, ok := ctx.Locals(requests.RequestKey).(*auctionReq.Update)
	if !ok {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}

	dbConn := db.SqlInstance(config.GetConfig().DB.DBDriver, config.GetConfig().DB.DBString).DB
	redisConn := redisdb.RedisInstance().RedisClient

	contxt, err := context_with_depends.ContextWithDepends(context.Background(), dbConn, redisConn)
	if err != nil {
		return responses.NewFailedResponse(ctx, errorhandler.ErrDependency)
	}

	auction, err := auctionService.Update(contxt, request)
	if err != nil {
		return responses.NewFailedResponse(ctx, err)
	}

	return responses.UpdateAuction(ctx, auction)
}

func Start(ctx *fiber.Ctx) error {
	request, ok := ctx.Locals(requests.RequestKey).(*auctionReq.Start)
	if !ok {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}

	dbConn := db.SqlInstance(config.GetConfig().DB.DBDriver, config.GetConfig().DB.DBString).DB
	redisConn := redisdb.RedisInstance().RedisClient

	contxt, err := context_with_depends.ContextWithDepends(context.Background(), dbConn, redisConn)
	if err != nil {
		return responses.NewFailedResponse(ctx, errorhandler.ErrDependency)
	}

	auction, err := auctionService.Start(contxt, request)
	if err != nil {
		return responses.NewFailedResponse(ctx, err)
	}

	return responses.StartAuction(ctx, auction)
}

func End(ctx *fiber.Ctx) error {
	request, ok := ctx.Locals(requests.RequestKey).(*auctionReq.End)
	if !ok {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}

	dbConn := db.SqlInstance(config.GetConfig().DB.DBDriver, config.GetConfig().DB.DBString).DB
	redisConn := redisdb.RedisInstance().RedisClient

	contxt, err := context_with_depends.ContextWithDepends(context.Background(), dbConn, redisConn)
	if err != nil {
		return responses.NewFailedResponse(ctx, errorhandler.ErrDependency)
	}

	auction, err := auctionService.End(contxt, request)
	if err != nil {
		return responses.NewFailedResponse(ctx, err)
	}

	return responses.EndAuction(ctx, auction)
}

func Delete(ctx *fiber.Ctx) error {
	request, ok := ctx.Locals(requests.RequestKey).(*auctionReq.Delete)
	if !ok {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}

	dbConn := db.SqlInstance(config.GetConfig().DB.DBDriver, config.GetConfig().DB.DBString).DB
	redisConn := redisdb.RedisInstance().RedisClient

	contxt, err := context_with_depends.ContextWithDepends(context.Background(), dbConn, redisConn)
	if err != nil {
		return responses.NewFailedResponse(ctx, errorhandler.ErrDependency)
	}

	auction, err := auctionService.Delete(contxt, request)
	if err != nil {
		return responses.NewFailedResponse(ctx, err)
	}

	return responses.DeleteAuction(ctx, auction)
}

func Participate(ctx *fiber.Ctx) error {
	request, ok := ctx.Locals(requests.RequestKey).(*auctionReq.Participate)
	if !ok {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}

	dbConn := db.SqlInstance(config.GetConfig().DB.DBDriver, config.GetConfig().DB.DBString).DB
	redisConn := redisdb.RedisInstance().RedisClient

	contxt, err := context_with_depends.ContextWithDepends(context.Background(), dbConn, redisConn)
	if err != nil {
		return responses.NewFailedResponse(ctx, errorhandler.ErrDependency)
	}

	auction, err := auctionService.Participate(contxt, request)
	if err != nil {
		return responses.NewFailedResponse(ctx, err)
	}

	return responses.Participate(ctx, auction)
}

func SetPrice(ctx *fiber.Ctx) error {
	request, ok := ctx.Locals(requests.RequestKey).(*auctionReq.SetPrice)
	if !ok {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}

	dbConn := db.SqlInstance(config.GetConfig().DB.DBDriver, config.GetConfig().DB.DBString).DB
	redisConn := redisdb.RedisInstance().RedisClient

	contxt, err := context_with_depends.ContextWithDepends(context.Background(), dbConn, redisConn)
	if err != nil {
		return responses.NewFailedResponse(ctx, errorhandler.ErrDependency)
	}

	price, err := auctionService.SetPrice(contxt, request)
	if err != nil {
		return responses.NewFailedResponse(ctx, err)
	}

	return responses.SetPrice(ctx, price)
}

func SetVisit(ctx *fiber.Ctx) error {
	request, ok := ctx.Locals(requests.RequestKey).(*auctionReq.SetVisit)
	if !ok {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}

	dbConn := db.SqlInstance(config.GetConfig().DB.DBDriver, config.GetConfig().DB.DBString).DB
	redisConn := redisdb.RedisInstance().RedisClient

	contxt, err := context_with_depends.ContextWithDepends(context.Background(), dbConn, redisConn)
	if err != nil {
		return responses.NewFailedResponse(ctx, errorhandler.ErrDependency)
	}

	settedVisit, err := auctionService.SetVisit(contxt, request)
	if err != nil {
		return responses.NewFailedResponse(ctx, err)
	}

	return responses.SettedVisit(ctx, settedVisit)
}

func Visit(ctx *fiber.Ctx) error {
	request, ok := ctx.Locals(requests.RequestKey).(*auctionReq.Visit)
	if !ok {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}

	dbConn := db.SqlInstance(config.GetConfig().DB.DBDriver, config.GetConfig().DB.DBString).DB
	redisConn := redisdb.RedisInstance().RedisClient

	contxt, err := context_with_depends.ContextWithDepends(context.Background(), dbConn, redisConn)
	if err != nil {
		return responses.NewFailedResponse(ctx, errorhandler.ErrDependency)
	}

	if err = auctionService.Visit(contxt, request); err != nil {
		return responses.NewFailedResponse(ctx, err)
	}

	return responses.Visit(ctx)
}

func Unvisit(ctx *fiber.Ctx) error {
	request, ok := ctx.Locals(requests.RequestKey).(*auctionReq.Unvisit)
	if !ok {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}

	dbConn := db.SqlInstance(config.GetConfig().DB.DBDriver, config.GetConfig().DB.DBString).DB
	redisConn := redisdb.RedisInstance().RedisClient

	contxt, err := context_with_depends.ContextWithDepends(context.Background(), dbConn, redisConn)
	if err != nil {
		return responses.NewFailedResponse(ctx, errorhandler.ErrDependency)
	}

	if err = auctionService.Unvisit(contxt, request); err != nil {
		return responses.NewFailedResponse(ctx, err)
	}

	return responses.Unvisit(ctx)
}

func Visitors(ctx *fiber.Ctx) error {
	request, ok := ctx.Locals(requests.RequestKey).(*auctionReq.Visitor)
	if !ok {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}

	dbConn := db.SqlInstance(config.GetConfig().DB.DBDriver, config.GetConfig().DB.DBString).DB
	redisConn := redisdb.RedisInstance().RedisClient

	contxt, err := context_with_depends.ContextWithDepends(context.Background(), dbConn, redisConn)
	if err != nil {
		return responses.NewFailedResponse(ctx, errorhandler.ErrDependency)
	}

	auctionVisitors, err := auctionService.Visitors(contxt, request)
	if err != nil {
		return responses.NewFailedResponse(ctx, err)
	}

	return responses.Visitors(ctx, auctionVisitors)
}

func UpdateVisit(ctx *fiber.Ctx) error {
	request, ok := ctx.Locals(requests.RequestKey).(*auctionReq.UpdateVisit)
	if !ok {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}

	dbConn := db.SqlInstance(config.GetConfig().DB.DBDriver, config.GetConfig().DB.DBString).DB
	redisConn := redisdb.RedisInstance().RedisClient

	contxt, err := context_with_depends.ContextWithDepends(context.Background(), dbConn, redisConn)
	if err != nil {
		return responses.NewFailedResponse(ctx, errorhandler.ErrDependency)
	}

	updatedVisit, err := auctionService.UpdateVisit(contxt, request)
	if err != nil {
		return responses.NewFailedResponse(ctx, err)
	}

	return responses.UpdateVisit(ctx, updatedVisit)
}
