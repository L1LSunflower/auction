package users

import (
	"context"

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

	user, err := userService.SignUp(request)
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
	tokens, err := userService.SignIn(context.Background(), dbConn, request)
	if err != nil {
		return responses.NewFailedResponse(ctx, err)
	}

	return responses.NewSuccessResponse(ctx, tokens)
}

func Confirm(ctx *fiber.Ctx) error {
	request, ok := ctx.Locals(requests.RequestKey).(*usersRequest.Confirm)
	if !ok {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}

	dbConn := db.SqlInstance(config.GetConfig().DB.DBDriver, config.GetConfig().DB.DBString).DB
	tokens, err := userService.Confirm(context.Background(), dbConn, request)
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
	user, err := userService.User(context.Background(), dbConn, request)
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

	tokens, err := userService.RefreshToken(request)
	if err != nil {
		return responses.NewFailedResponse(ctx, err)
	}

	return responses.NewSuccessResponse(ctx, tokens)
}
