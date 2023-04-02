package users

import (
	"context"
	"github.com/L1LSunflower/auction/internal/tools/context_with_depends"
	"github.com/L1LSunflower/auction/pkg/redisdb"
	"github.com/gofiber/fiber/v2"

	"github.com/L1LSunflower/auction/config"
	userService "github.com/L1LSunflower/auction/internal/domain/services/users"
	"github.com/L1LSunflower/auction/internal/requests"
	usersRequest "github.com/L1LSunflower/auction/internal/requests/structs/users"
	"github.com/L1LSunflower/auction/internal/responses"
	"github.com/L1LSunflower/auction/internal/tools/errorhandler"
	"github.com/L1LSunflower/auction/pkg/db"
)

func SignUp(ctx *fiber.Ctx) error {
	request, ok := ctx.Locals(requests.RequestKey).(*usersRequest.SignUp)
	if !ok {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}

	dbConn := db.SqlInstance(config.GetConfig().DB.DBDriver, config.GetConfig().DB.DBString).DB
	redisConn := redisdb.RedisInstance().RedisClient

	contxt, err := context_with_depends.ContextWithDepends(context.Background(), dbConn, redisConn)
	if err != nil {
		return responses.NewFailedResponse(ctx, err)
	}

	user, err := userService.SignUp(contxt, request)
	if err != nil {
		return responses.NewFailedResponse(ctx, err)
	}

	return responses.SuccessSignUp(ctx, user)
}

func SignIn(ctx *fiber.Ctx) error {
	request, ok := ctx.Locals(requests.RequestKey).(*usersRequest.SignIn)
	if !ok {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}

	dbConn := db.SqlInstance(config.GetConfig().DB.DBDriver, config.GetConfig().DB.DBString).DB
	redisConn := redisdb.RedisInstance().RedisClient

	contxt, err := context_with_depends.ContextWithDepends(context.Background(), dbConn, redisConn)
	if err != nil {
		return responses.NewFailedResponse(ctx, errorhandler.ErrDependency)
	}

	userToken, err := userService.SignIn(contxt, request)
	if err != nil {
		return responses.NewFailedResponse(ctx, err)
	}

	return responses.SuccessSignIn(ctx, userToken)
}

func Confirm(ctx *fiber.Ctx) error {
	request, ok := ctx.Locals(requests.RequestKey).(*usersRequest.Confirm)
	if !ok {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}

	dbConn := db.SqlInstance(config.GetConfig().DB.DBDriver, config.GetConfig().DB.DBString).DB
	redisConn := redisdb.RedisInstance().RedisClient

	contxt, err := context_with_depends.ContextWithDepends(context.Background(), dbConn, redisConn)
	if err != nil {
		return responses.NewFailedResponse(ctx, errorhandler.ErrDependency)
	}

	userToken, err := userService.Confirm(contxt, request)
	if err != nil {
		return responses.NewFailedResponse(ctx, err)
	}

	return responses.SuccessConfirm(ctx, userToken)
}

func GetUser(ctx *fiber.Ctx) error {
	request, ok := ctx.Locals(requests.RequestKey).(*usersRequest.User)
	if !ok {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}

	dbConn := db.SqlInstance(config.GetConfig().DB.DBDriver, config.GetConfig().DB.DBString).DB
	redisConn := redisdb.RedisInstance().RedisClient

	contxt, err := context_with_depends.ContextWithDepends(context.Background(), dbConn, redisConn)
	if err != nil {
		return responses.NewFailedResponse(ctx, errorhandler.ErrDependency)
	}

	user, err := userService.User(contxt, request)
	if err != nil {
		return responses.NewFailedResponse(ctx, err)
	}

	return responses.SuccessGetUser(ctx, user)
}

func Refresh(ctx *fiber.Ctx) error {
	request, ok := ctx.Locals(requests.RequestKey).(*usersRequest.Tokens)
	if !ok {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}

	dbConn := db.SqlInstance(config.GetConfig().DB.DBDriver, config.GetConfig().DB.DBString).DB
	redisConn := redisdb.RedisInstance().RedisClient

	contxt, err := context_with_depends.ContextWithDepends(context.Background(), dbConn, redisConn)
	if err != nil {
		return responses.NewFailedResponse(ctx, errorhandler.ErrDependency)
	}

	tokens, err := userService.RefreshToken(contxt, request)
	if err != nil {
		return responses.NewFailedResponse(ctx, err)
	}

	return responses.RefreshTokens(ctx, tokens)
}

func Restore(ctx *fiber.Ctx) error {
	request, ok := ctx.Locals(requests.RequestKey).(*usersRequest.RestorePassword)
	if !ok {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}

	dbConn := db.SqlInstance(config.GetConfig().DB.DBDriver, config.GetConfig().DB.DBString).DB
	redisConn := redisdb.RedisInstance().RedisClient

	contxt, err := context_with_depends.ContextWithDepends(context.Background(), dbConn, redisConn)
	if err != nil {
		return responses.NewFailedResponse(ctx, errorhandler.ErrDependency)
	}

	if err = userService.SendRestoreCode(contxt, request); err != nil {
		return responses.NewFailedResponse(ctx, errorhandler.ErrRestore)
	}

	return responses.SuccessSendOtp(ctx, "success", "otp sent")
}

func ChangePassword(ctx *fiber.Ctx) error {
	request, ok := ctx.Locals(requests.RequestKey).(*usersRequest.ChangePassword)
	if !ok {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}

	dbConn := db.SqlInstance(config.GetConfig().DB.DBDriver, config.GetConfig().DB.DBString).DB
	redisConn := redisdb.RedisInstance().RedisClient

	contxt, err := context_with_depends.ContextWithDepends(context.Background(), dbConn, redisConn)
	if err != nil {
		return responses.NewFailedResponse(ctx, errorhandler.ErrDependency)
	}

	tokens, err := userService.ChangePassword(contxt, request)
	if err != nil {
		return responses.NewFailedResponse(ctx, err)
	}

	return responses.SuccessChangePassword(ctx, tokens)
}

func Update(ctx *fiber.Ctx) error {
	request, ok := ctx.Locals(requests.RequestKey).(*usersRequest.Update)
	if !ok {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}

	dbConn := db.SqlInstance(config.GetConfig().DB.DBDriver, config.GetConfig().DB.DBString).DB
	redisConn := redisdb.RedisInstance().RedisClient

	contxt, err := context_with_depends.ContextWithDepends(context.Background(), dbConn, redisConn)
	if err != nil {
		return responses.NewFailedResponse(ctx, errorhandler.ErrDependency)
	}

	user, err := userService.Update(contxt, request)
	if err != nil {
		return responses.NewFailedResponse(ctx, err)
	}

	return responses.UpdateUser(ctx, user)
}

func Delete(ctx *fiber.Ctx) error {
	request, ok := ctx.Locals(requests.RequestKey).(*usersRequest.Delete)
	if !ok {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}

	dbConn := db.SqlInstance(config.GetConfig().DB.DBDriver, config.GetConfig().DB.DBString).DB
	redisConn := redisdb.RedisInstance().RedisClient

	contxt, err := context_with_depends.ContextWithDepends(context.Background(), dbConn, redisConn)
	if err != nil {
		return responses.NewFailedResponse(ctx, errorhandler.ErrDependency)
	}

	user, err := userService.Delete(contxt, request)
	if err != nil {
		return responses.NewFailedResponse(ctx, err)
	}

	return responses.DeleteUser(ctx, user)
}

func Profile(ctx *fiber.Ctx) error {
	request, ok := ctx.Locals(requests.RequestKey).(*usersRequest.User)
	if !ok {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}

	dbConn := db.SqlInstance(config.GetConfig().DB.DBDriver, config.GetConfig().DB.DBString).DB
	redisConn := redisdb.RedisInstance().RedisClient

	contxt, err := context_with_depends.ContextWithDepends(context.Background(), dbConn, redisConn)
	if err != nil {
		return responses.NewFailedResponse(ctx, errorhandler.ErrDependency)
	}

	userProfile, err := userService.Profile(contxt, request)
	if err != nil {
		return responses.NewFailedResponse(ctx, err)
	}

	return responses.UserProfile(ctx, userProfile)
}

func ProfileHistory(ctx *fiber.Ctx) error {
	request, ok := ctx.Locals(requests.RequestKey).(*usersRequest.User)
	if !ok {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}

	dbConn := db.SqlInstance(config.GetConfig().DB.DBDriver, config.GetConfig().DB.DBString).DB
	redisConn := redisdb.RedisInstance().RedisClient

	contxt, err := context_with_depends.ContextWithDepends(context.Background(), dbConn, redisConn)
	if err != nil {
		return responses.NewFailedResponse(ctx, errorhandler.ErrDependency)
	}

	userProfile, err := userService.OwnersAuctions(contxt, request)
	if err != nil {
		return responses.NewFailedResponse(ctx, err)
	}

	return responses.UserProfileHistory(ctx, userProfile)
}

func ProfileCompleted(ctx *fiber.Ctx) error {
	request, ok := ctx.Locals(requests.RequestKey).(*usersRequest.User)
	if !ok {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}

	dbConn := db.SqlInstance(config.GetConfig().DB.DBDriver, config.GetConfig().DB.DBString).DB
	redisConn := redisdb.RedisInstance().RedisClient

	contxt, err := context_with_depends.ContextWithDepends(context.Background(), dbConn, redisConn)
	if err != nil {
		return responses.NewFailedResponse(ctx, errorhandler.ErrDependency)
	}

	userCompleted, err := userService.CompletedAuctions(contxt, request)
	if err != nil {
		return responses.NewFailedResponse(ctx, err)
	}

	return responses.CompletedAuctions(ctx, userCompleted)
}
