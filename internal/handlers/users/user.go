package users

import (
	"context"
	"errors"
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
	user, err := userService.SignUp(contxt, request)
	if err != nil {
		return responses.NewFailedResponse(ctx, err)
	}

	return responses.NewSuccessResponse(ctx, user)
}

func SignIn(ctx *fiber.Ctx) error {
	request, ok := ctx.Locals(requests.RequestKey).(*usersRequest.SignIn)
	if !ok {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}

	dbConn := db.SqlInstance(config.GetConfig().DB.DBDriver, config.GetConfig().DB.DBString).DB
	redisConn := redisdb.RedisInstance().RedisClient
	contxt, err := context_with_depends.ContextWithDepends(context.Background(), dbConn, redisConn)
	tokens, id, err := userService.SignIn(contxt, request)
	if err != nil {
		return responses.NewFailedResponse(ctx, err)
	}

	return responses.NewSuccessResponse(ctx, fiber.Map{
		"status":  "success",
		"access":  tokens.AccessToken,
		"refresh": tokens.RefreshToken,
		"id":      id,
	})
}

func Confirm(ctx *fiber.Ctx) error {
	request, ok := ctx.Locals(requests.RequestKey).(*usersRequest.Confirm)
	if !ok {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}

	dbConn := db.SqlInstance(config.GetConfig().DB.DBDriver, config.GetConfig().DB.DBString).DB
	redisConn := redisdb.RedisInstance().RedisClient
	contxt, err := context_with_depends.ContextWithDepends(context.Background(), dbConn, redisConn)
	tokens, err := userService.Confirm(contxt, request)
	if err != nil {
		return responses.NewFailedResponse(ctx, err)
	}

	return responses.NewSuccessResponse(ctx, tokens)
}

func GetUser(ctx *fiber.Ctx) error {
	request, ok := ctx.Locals(requests.RequestKey).(*usersRequest.User)
	if !ok {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}

	dbConn := db.SqlInstance(config.GetConfig().DB.DBDriver, config.GetConfig().DB.DBString).DB
	redisConn := redisdb.RedisInstance().RedisClient
	contxt, err := context_with_depends.ContextWithDepends(context.Background(), dbConn, redisConn)
	user, err := userService.User(contxt, request)
	if err != nil {
		return responses.NewFailedResponse(ctx, err)
	}

	return responses.NewSuccessResponse(ctx, user)
}

func Refresh(ctx *fiber.Ctx) error {
	request, ok := ctx.Locals(requests.RequestKey).(*usersRequest.Tokens)
	if !ok {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}

	dbConn := db.SqlInstance(config.GetConfig().DB.DBDriver, config.GetConfig().DB.DBString).DB
	redisConn := redisdb.RedisInstance().RedisClient
	contxt, err := context_with_depends.ContextWithDepends(context.Background(), dbConn, redisConn)
	tokens, err := userService.RefreshToken(contxt, request)
	if err != nil {
		return responses.NewFailedResponse(ctx, err)
	}

	return responses.NewSuccessResponse(ctx, tokens)
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
		return responses.NewFailedResponse(ctx, errors.New("failed to start context with dependencies"))
	}
	if err = userService.RestorePassword(contxt, request); err != nil {
		return responses.NewFailedResponse(ctx, errors.New("failed to restore passwrod"))
	}

	return responses.NewSuccessResponse(ctx, fiber.Map{
		"status":  "success",
		"message": "otp sent",
	})
}
