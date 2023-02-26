package auctions

import (
	"github.com/L1LSunflower/auction/internal/requests"
	"github.com/L1LSunflower/auction/internal/responses"
	"github.com/L1LSunflower/auction/internal/tools/errorhandler"
	"github.com/gofiber/fiber/v2"

	requestAuction "github.com/L1LSunflower/auction/internal/requests/structs/auctions"
)

func Create(ctx *fiber.Ctx) error {
	request := &requestAuction.Create{}
	if err := ctx.BodyParser(request); err != nil {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}
	ctx.Locals(requests.RequestKey, request)

	return ctx.Next()
}

func Auction(ctx *fiber.Ctx) error {
	var err error
	request := &requestAuction.Auction{}

	if request.ID, err = ctx.ParamsInt("id"); err != nil {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}
	ctx.Locals(requests.RequestKey, request)

	return ctx.Next()
}

func Auctions(ctx *fiber.Ctx) error {

	return ctx.Next()
}

func Update(ctx *fiber.Ctx) error {
	var err error
	request := &requestAuction.Update{}

	if request.ID, err = ctx.ParamsInt("id"); err != nil {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}

	if err = ctx.BodyParser(request); err != nil {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}

	return ctx.Next()
}

func Start(ctx *fiber.Ctx) error {
	var err error
	request := &requestAuction.Start{}

	if request.ID, err = ctx.ParamsInt("id"); err != nil {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}
	ctx.Locals(requests.RequestKey, request)

	return ctx.Next()
}

func End(ctx *fiber.Ctx) error {
	var err error
	request := &requestAuction.End{}

	if request.ID, err = ctx.ParamsInt("id"); err != nil {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}
	ctx.Locals(requests.RequestKey, request)

	return ctx.Next()
}

func Delete(ctx *fiber.Ctx) error {
	var err error
	request := &requestAuction.Delete{}

	if request.ID, err = ctx.ParamsInt("id"); err != nil {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}
	ctx.Locals(requests.RequestKey, request)

	return ctx.Next()
}
