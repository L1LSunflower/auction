package users

import (
	"github.com/L1LSunflower/auction/internal/requests"
	userRequest "github.com/L1LSunflower/auction/internal/requests/structs/users"
	"github.com/L1LSunflower/auction/internal/responses"
	"github.com/L1LSunflower/auction/internal/tools/errorhandler"
	"github.com/gofiber/fiber/v2"
)

const (
	uuidLength = 36
)

func SignUpValidator(ctx *fiber.Ctx) error {
	request := &userRequest.SignUp{}
	if err := ctx.BodyParser(request); err != nil {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}

	ctx.Locals(requests.RequestKey, request)

	return ctx.Next()
}

func SignInValidator(ctx *fiber.Ctx) error {
	request := &userRequest.SignIn{}
	if err := ctx.BodyParser(request); err != nil {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}

	ctx.Locals(requests.RequestKey, request)

	return ctx.Next()
}

func ConfirmValidator(ctx *fiber.Ctx) error {
	request := &userRequest.Confirm{}
	if err := ctx.BodyParser(request); err != nil {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}

	request.ID = ctx.Params("id")
	if len(request.ID) != uuidLength {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}

	ctx.Locals(requests.RequestKey, request)

	return ctx.Next()
}

func GetUserValidator(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if len(id) != uuidLength {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}

	request := &userRequest.User{ID: id}
	ctx.Locals(requests.RequestKey, request)

	return ctx.Next()
}
