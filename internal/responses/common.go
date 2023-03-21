package responses

import (
	"github.com/L1LSunflower/auction/internal/responses/structs"
	"github.com/gofiber/fiber/v2"

	"github.com/L1LSunflower/auction/internal/tools/errorhandler"
)

const (
	successStatus = "success"
)

func NewFailedResponse(ctx *fiber.Ctx, err error) error {
	var (
		typeError *errorhandler.TypeError
		ok        bool
	)

	if typeError, ok = errorhandler.TypeErrors[err.Error()]; !ok {
		typeError = &errorhandler.TypeError{
			StatusCode: 999,
			HttpCode:   fiber.StatusInternalServerError,
		}
	}

	ctx.Status(typeError.HttpCode)
	return ctx.JSON(&structs.ErrorResponse{
		StatusCode: typeError.StatusCode,
		Message:    err.Error(),
	})
}

func NewSuccessResponse(ctx *fiber.Ctx, data any) error {
	return ctx.JSON(&structs.SuccessResponse{
		Status:  successStatus,
		Message: data,
	})
}

func NewValidationErrResponse(ctx *fiber.Ctx, err error) error {
	ctx.Status(fiber.StatusBadRequest)
	return ctx.JSON(&structs.ErrorResponse{
		StatusCode: errorhandler.ErrParseRequestCode,
		Message:    err.Error(),
	})
}

func NewErrorResponse(statusCode int, err error) any {
	return &structs.ErrorResponse{
		StatusCode: statusCode,
		Message:    err.Error(),
	}
}
