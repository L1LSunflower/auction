package users

import (
	"github.com/L1LSunflower/auction/internal/middlewares/validator"
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
	if err := requests.ParseRequest(ctx, request); err != nil {
		return err
	}

	if err := validator.ValidateRequest(request); err != nil {
		return responses.NewValidationErrResponse(ctx, err)
	}

	ctx.Locals(requests.RequestKey, request)

	return ctx.Next()
}

func SignInValidator(ctx *fiber.Ctx) error {
	request := &userRequest.SignIn{}
	if err := requests.ParseRequest(ctx, request); err != nil {
		return err
	}

	if err := validator.ValidateRequest(request); err != nil {
		return responses.NewValidationErrResponse(ctx, err)
	}

	ctx.Locals(requests.RequestKey, request)

	return ctx.Next()
}

func ConfirmValidator(ctx *fiber.Ctx) error {
	request := &userRequest.Confirm{}
	if err := requests.ParseRequest(ctx, request); err != nil {
		return err
	}

	if request.ID = ctx.Params("id"); len(request.ID) != uuidLength {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}

	if err := validator.ValidateRequest(request); err != nil {
		return responses.NewValidationErrResponse(ctx, err)
	}

	ctx.Locals(requests.RequestKey, request)

	return ctx.Next()
}

func GetUserValidator(ctx *fiber.Ctx) error {
	id, ok := ctx.Locals(requests.UserIDCtx).(string)
	if !ok || len(id) != uuidLength {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}

	request := &userRequest.User{ID: id}

	if err := validator.ValidateRequest(request); err != nil {
		return responses.NewValidationErrResponse(ctx, err)
	}

	ctx.Locals(requests.RequestKey, request)

	return ctx.Next()
}

func RestoreValidator(ctx *fiber.Ctx) error {
	request := &userRequest.RestorePassword{}
	if err := requests.ParseRequest(ctx, request); err != nil {
		return err
	}

	if err := validator.ValidateRequest(request); err != nil {
		return responses.NewValidationErrResponse(ctx, err)
	}

	ctx.Locals(requests.RequestKey, request)

	return ctx.Next()
}

func RefreshValidator(ctx *fiber.Ctx) error {
	var ok bool
	request := &userRequest.Tokens{}

	if request.AccessToken = ctx.Get("access"); request.AccessToken == "" {
		return responses.NewFailedResponse(ctx, errorhandler.AuthRequired)
	}

	if request.RefreshToken = ctx.Get("refresh"); request.RefreshToken == "" {
		return responses.NewFailedResponse(ctx, errorhandler.AuthRequired)
	}

	if request.ID, ok = ctx.Locals(requests.UserIDCtx).(string); !ok || request.ID == "" {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}

	if err := validator.ValidateRequest(request); err != nil {
		return responses.NewValidationErrResponse(ctx, err)
	}

	ctx.Locals(requests.RequestKey, request)
	return ctx.Next()
}

func ChangePasswordValidator(ctx *fiber.Ctx) error {
	request := &userRequest.ChangePassword{}

	if err := requests.ParseRequest(ctx, request); err != nil {
		return err
	}

	if err := validator.ValidateRequest(request); err != nil {
		return responses.NewValidationErrResponse(ctx, err)
	}

	ctx.Locals(requests.RequestKey, request)

	return ctx.Next()
}

func UpdateValidator(ctx *fiber.Ctx) error {
	var ok bool
	request := &userRequest.Update{}

	if err := requests.ParseRequest(ctx, request); err != nil {
		return err
	}

	if request.ID, ok = ctx.Locals(requests.UserIDCtx).(string); !ok || request.ID == "" {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}

	if err := validator.ValidateRequest(request); err != nil {
		return responses.NewValidationErrResponse(ctx, err)
	}

	ctx.Locals(requests.RequestKey, request)

	return ctx.Next()
}

func DeleteValidator(ctx *fiber.Ctx) error {
	var ok bool
	request := &userRequest.Delete{}

	if request.ID, ok = ctx.Locals(requests.UserIDCtx).(string); !ok || request.ID == "" {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}

	if err := validator.ValidateRequest(request); err != nil {
		return responses.NewValidationErrResponse(ctx, err)
	}

	ctx.Locals(requests.RequestKey, request)

	return ctx.Next()
}

func ProfileValidator(ctx *fiber.Ctx) error {
	var ok bool
	request := &userRequest.User{}

	if request.ID, ok = ctx.Locals(requests.UserIDCtx).(string); !ok || request.ID == "" {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}

	if err := validator.ValidateRequest(request); err != nil {
		return responses.NewValidationErrResponse(ctx, err)
	}

	ctx.Locals(requests.RequestKey, request)

	return ctx.Next()
}

func ProfileCompleted(ctx *fiber.Ctx) error {
	var ok bool
	request := &userRequest.User{}

	if request.ID, ok = ctx.Locals(requests.UserIDCtx).(string); !ok || request.ID == "" {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}

	if err := validator.ValidateRequest(request); err != nil {
		return responses.NewValidationErrResponse(ctx, err)
	}

	ctx.Locals(requests.RequestKey, request)

	return ctx.Next()
}
