package auctions

import (
	"fmt"
	"github.com/L1LSunflower/auction/internal/requests"
	"github.com/L1LSunflower/auction/internal/responses"
	"github.com/L1LSunflower/auction/internal/tools/errorhandler"
	"github.com/L1LSunflower/auction/internal/tools/metadata"
	"github.com/gofiber/fiber/v2"
	"strings"

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

	if tags := ctx.Get("tags"); tags != "" {
		if strings.Index(tags, "") >= 0 {
			request.Tags = strings.Split(tags[1:len(tags)-1-1], ",")
		} else {
			request.Tags = append(request.Tags, fmt.Sprintf("'%s'", tags[1:len(tags)-1-1]))
		}
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

	if err = ctx.BodyParser(request); err != nil {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
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
