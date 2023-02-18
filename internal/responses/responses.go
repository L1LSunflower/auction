package responses

import (
	"github.com/L1LSunflower/auction/internal/responses/structs"
	"github.com/gofiber/fiber/v2"

	"github.com/L1LSunflower/auction/internal/tools/errorhandler"
)

const (
	errorStatus   = "error"
	successStatus = "success"
)

func NewFailedResponse(ctx *fiber.Ctx, err error) error {
	var (
		typeError *errorhandler.TypeError
		ok        bool
	)

	if typeError, ok = errorhandler.TypeErrors[err.Error()]; !ok {
		typeError = &errorhandler.TypeError{
			StatusCode: fiber.StatusInternalServerError,
		}
	}

	ctx.Status(typeError.StatusCode)
	return ctx.JSON(&structs.ErrorResponse{
		Status:  errorStatus,
		Message: err.Error(),
	})
}

func NewSuccessResponse(ctx *fiber.Ctx, data any) error {
	return ctx.JSON(&structs.SuccessResponse{
		Status:  successStatus,
		Message: data,
	})
}
