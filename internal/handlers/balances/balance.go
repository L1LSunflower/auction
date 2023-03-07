package balances

import (
	"github.com/gofiber/fiber/v2"
)

func Credit(ctx *fiber.Ctx) error {
	//request, ok := ctx.Locals(requests.RequestKey).(*balanceReq.Credit)
	//if !ok {
	//	return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	//}
	//
	//dbConn := db.SqlInstance(config.GetConfig().DB.DBDriver, config.GetConfig().DB.DBString).DB
	//redisConn := redisdb.RedisInstance().RedisClient

	//contxt, err := context_with_depends.ContextWithDepends(context.Background(), dbConn, redisConn)
	//if err != nil {
	//	return responses.NewFailedResponse(ctx, errorhandler.ErrDependency)
	//}

	//userBalance, err := balanceService.Credit(contxt, request)
	//if err != nil {
	//	return responses.NewFailedResponse(ctx, err)
	//}

	//return responses.CreditBalance(ctx, userBalance)
	return nil
}

func Debit(ctx *fiber.Ctx) error {
	//request, ok := ctx.Locals(requests.RequestKey).(*balanceReq.Debit)
	//if !ok {
	//	return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	//}
	//
	//dbConn := db.SqlInstance(config.GetConfig().DB.DBDriver, config.GetConfig().DB.DBString).DB
	//redisConn := redisdb.RedisInstance().RedisClient
	//
	//contxt, err := context_with_depends.ContextWithDepends(context.Background(), dbConn, redisConn)
	//if err != nil {
	//	return responses.NewFailedResponse(ctx, errorhandler.ErrDependency)
	//}
	//
	//userBalance, err := balanceService.Debit(contxt, request)
	//if err != nil {
	//	return responses.NewFailedResponse(ctx, err)
	//}
	//
	//return responses.DebitBalance(ctx, userBalance)
	return nil
}

func Balance(ctx *fiber.Ctx) error {
	//request, ok := ctx.Locals(requests.RequestKey).(*balanceReq.Balance)
	//if !ok {
	//	return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	//}
	//
	//dbConn := db.SqlInstance(config.GetConfig().DB.DBDriver, config.GetConfig().DB.DBString).DB
	//redisConn := redisdb.RedisInstance().RedisClient
	//
	//contxt, err := context_with_depends.ContextWithDepends(context.Background(), dbConn, redisConn)
	//if err != nil {
	//	return responses.NewFailedResponse(ctx, errorhandler.ErrDependency)
	//}
	//
	//userBalance, err := balanceService.Balance(contxt, request)
	//if err != nil {
	//	return responses.NewFailedResponse(ctx, err)
	//}
	//
	//return responses.Balance(ctx, userBalance)
	return nil
}
