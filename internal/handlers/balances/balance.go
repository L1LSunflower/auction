package balances

import (
	"context"

	"github.com/gofiber/fiber/v2"

	"github.com/L1LSunflower/auction/config"
	balanceService "github.com/L1LSunflower/auction/internal/domain/services/balances"
	"github.com/L1LSunflower/auction/internal/requests"
	balanceReq "github.com/L1LSunflower/auction/internal/requests/structs/balances"
	"github.com/L1LSunflower/auction/internal/responses"
	"github.com/L1LSunflower/auction/internal/tools/context_with_depends"
	"github.com/L1LSunflower/auction/internal/tools/errorhandler"
	"github.com/L1LSunflower/auction/pkg/db"
	"github.com/L1LSunflower/auction/pkg/redisdb"
)

func Credit(ctx *fiber.Ctx) error {
	request, ok := ctx.Locals(requests.RequestKey).(*balanceReq.Credit)
	if !ok {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}

	dbConn := db.SqlInstance(config.GetConfig().DB.DBDriver, config.GetConfig().DB.DBString).DB
	redisConn := redisdb.RedisInstance().RedisClient

	contxt, err := context_with_depends.ContextWithDepends(context.Background(), dbConn, redisConn)
	if err != nil {
		return responses.NewFailedResponse(ctx, errorhandler.ErrDependency)
	}

	userBalance, err := balanceService.Credit(contxt, request)
	if err != nil {
		return responses.NewFailedResponse(ctx, err)
	}

	return responses.Credit(ctx, userBalance)
}

func Debit(ctx *fiber.Ctx) error {
	request, ok := ctx.Locals(requests.RequestKey).(*balanceReq.Debit)
	if !ok {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}

	dbConn := db.SqlInstance(config.GetConfig().DB.DBDriver, config.GetConfig().DB.DBString).DB
	redisConn := redisdb.RedisInstance().RedisClient

	contxt, err := context_with_depends.ContextWithDepends(context.Background(), dbConn, redisConn)
	if err != nil {
		return responses.NewFailedResponse(ctx, errorhandler.ErrDependency)
	}

	userBalance, err := balanceService.Debit(contxt, request)
	if err != nil {
		return responses.NewFailedResponse(ctx, err)
	}

	return responses.Debit(ctx, userBalance)
}

func Balance(ctx *fiber.Ctx) error {
	request, ok := ctx.Locals(requests.RequestKey).(*balanceReq.Balance)
	if !ok {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}

	dbConn := db.SqlInstance(config.GetConfig().DB.DBDriver, config.GetConfig().DB.DBString).DB
	redisConn := redisdb.RedisInstance().RedisClient

	contxt, err := context_with_depends.ContextWithDepends(context.Background(), dbConn, redisConn)
	if err != nil {
		return responses.NewFailedResponse(ctx, errorhandler.ErrDependency)
	}

	userBalance, err := balanceService.Balance(contxt, request)
	if err != nil {
		return responses.NewFailedResponse(ctx, err)
	}

	return responses.Balance(ctx, userBalance)
}
