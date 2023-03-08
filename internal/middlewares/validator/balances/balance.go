package balances

import (
	"github.com/L1LSunflower/auction/internal/requests"
	"github.com/L1LSunflower/auction/internal/requests/structs/balances"
	"github.com/L1LSunflower/auction/internal/responses"
	"github.com/L1LSunflower/auction/internal/tools/errorhandler"
	"github.com/gofiber/fiber/v2"
)

func Credit(ctx *fiber.Ctx) error {
	var ok bool
	request := &balances.Credit{}
	if err := ctx.BodyParser(request); err != nil {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}

	if request.ID, ok = ctx.Locals(requests.UserIDCtx).(string); !ok {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}

	ctx.Locals(requests.RequestKey, request)

	return ctx.Next()
}

func Debit(ctx *fiber.Ctx) error {
	var ok bool
	request := &balances.Debit{}

	if request.ID, ok = ctx.Locals(requests.UserIDCtx).(string); !ok {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}

	if err := ctx.BodyParser(request); err != nil {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}

	ctx.Locals(requests.RequestKey, request)

	return ctx.Next()
}

func Balance(ctx *fiber.Ctx) error {
	var ok bool
	request := &balances.Balance{}

	if request.ID, ok = ctx.Locals(requests.UserIDCtx).(string); !ok {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}

	ctx.Locals(requests.RequestKey, request)

	return ctx.Next()
}
