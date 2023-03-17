package auctions

import (
	"fmt"
	"github.com/L1LSunflower/auction/internal/middlewares/validator"
	"github.com/L1LSunflower/auction/internal/requests"
	requestAuction "github.com/L1LSunflower/auction/internal/requests/structs/auctions"
	"github.com/L1LSunflower/auction/internal/responses"
	"github.com/L1LSunflower/auction/internal/tools/errorhandler"
	"github.com/L1LSunflower/auction/internal/tools/metadata"
	"github.com/gofiber/fiber/v2"
)

func Create(ctx *fiber.Ctx) error {
	request := &requestAuction.Create{}

	if err := requests.ParseRequest(ctx, request); err != nil {
		return err
	}

	if err := validator.ValidateRequest(request); err != nil {
		return responses.NewValidationErrResponse(ctx, err)
	}

	ctx.Locals(requests.RequestKey, request)

	return ctx.Next()
}

func Auction(ctx *fiber.Ctx) error {
	var (
		err error
		ok  bool
	)
	request := &requestAuction.Auction{}

	if request.ID, err = ctx.ParamsInt("id"); err != nil {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}
	ctx.Locals(requests.RequestKey, request)

	if request.UserID, ok = ctx.Locals(requests.UserIDCtx).(string); !ok {
		return responses.NewFailedResponse(ctx, errorhandler.ErrUserNotExist)
	}

	if err = validator.ValidateRequest(request); err != nil {
		return responses.NewValidationErrResponse(ctx, err)
	}

	ctx.Locals(requests.RequestKey, request)

	return ctx.Next()
}

func Auctions(ctx *fiber.Ctx) error {
	var err error
	request := &requestAuction.Auctions{}

	if request.Metadata, err = metadata.GetParams(ctx); err != nil {
		return responses.NewFailedResponse(ctx, err)
	}

	// Group by
	price := ctx.Get("price")
	if price == "high to low" {
		request.OrderBy = "a.start_price asc"
	} else if price == "low to high" {
		request.OrderBy = "a.start_price desc"
	}

	if price == "" {
		date := ctx.Get("date")
		if date == "newer" {
			request.OrderBy = "a.created_at desc"
		} else if date == "older" {
			request.OrderBy = "a.created_at asc"
		}
	}

	// Where
	if category := ctx.Get("category"); category != "" {
		request.Where = append(request.Where, fmt.Sprintf("category='%s'", category))
	}

	if status := ctx.Get("status"); status != "" {
		request.Where = append(request.Where, fmt.Sprintf("status='%s'", status))
	}

	tags := metadata.ParseParams(ctx.Request().URI().QueryArgs(), "tags")
	if len(tags) > 0 {
		request.Tags = metadata.PrepareTags(tags)
	}

	if err = validator.ValidateRequest(request); err != nil {
		return responses.NewValidationErrResponse(ctx, err)
	}

	ctx.Locals(requests.RequestKey, request)

	return ctx.Next()
}

func Update(ctx *fiber.Ctx) error {
	var err error
	request := &requestAuction.Update{}

	if request.ID, err = ctx.ParamsInt("id"); err != nil {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}

	if err = requests.ParseRequest(ctx, request); err != nil {
		return err
	}

	if err = validator.ValidateRequest(request); err != nil {
		return responses.NewValidationErrResponse(ctx, err)
	}

	ctx.Locals(requests.RequestKey, request)

	return ctx.Next()
}

func Start(ctx *fiber.Ctx) error {
	var err error
	request := &requestAuction.Start{}

	if request.ID, err = ctx.ParamsInt("id"); err != nil {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}
	ctx.Locals(requests.RequestKey, request)

	if err = validator.ValidateRequest(request); err != nil {
		return responses.NewValidationErrResponse(ctx, err)
	}

	return ctx.Next()
}

func End(ctx *fiber.Ctx) error {
	var err error
	request := &requestAuction.End{}

	if request.ID, err = ctx.ParamsInt("id"); err != nil {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}

	if err = validator.ValidateRequest(request); err != nil {
		return responses.NewValidationErrResponse(ctx, err)
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

	if err = validator.ValidateRequest(request); err != nil {
		return responses.NewValidationErrResponse(ctx, err)
	}

	ctx.Locals(requests.RequestKey, request)

	return ctx.Next()
}

func Participate(ctx *fiber.Ctx) error {
	var (
		ok  bool
		err error
	)
	request := &requestAuction.Participate{}

	if request.UserID, ok = ctx.Locals(requests.UserIDCtx).(string); !ok || request.UserID == "" {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}

	if request.AuctionID, err = ctx.ParamsInt("id"); err != nil || request.AuctionID <= 0 {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}

	if err = validator.ValidateRequest(request); err != nil {
		return responses.NewValidationErrResponse(ctx, err)
	}

	ctx.Locals(requests.RequestKey, request)

	return ctx.Next()
}
