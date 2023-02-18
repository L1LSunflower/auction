package cities

import (
	"github.com/gofiber/fiber/v2"
	"strconv"

	"github.com/L1LSunflower/auction/internal/requests"
	cityRequest "github.com/L1LSunflower/auction/internal/requests/structs/cities"
	"github.com/L1LSunflower/auction/internal/responses"
	"github.com/L1LSunflower/auction/internal/tools/errorhandler"
)

const (
	cityID = "id"
)

func Create(ctx *fiber.Ctx) error {
	request := &cityRequest.CreateCity{}
	if err := ctx.BodyParser(request); err != nil {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}

	ctx.Locals(requests.RequestKey, request)

	return ctx.Next()
}

func City(ctx *fiber.Ctx) error {
	var (
		request = &cityRequest.City{}
		err     error
	)

	if request.ID, err = strconv.Atoi(ctx.Params(cityID)); err != nil {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}

	ctx.Locals(requests.RequestKey, request)

	return ctx.Next()
}

func UpdateCity(ctx *fiber.Ctx) error {
	var (
		request = &cityRequest.CityUpdate{}
		err     error
	)

	if err = ctx.BodyParser(request); err != nil {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}

	if request.ID, err = strconv.Atoi(ctx.Params(cityID)); err != nil {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}

	ctx.Locals(requests.RequestKey, request)

	return ctx.Next()
}

func DeleteCity(ctx *fiber.Ctx) error {
	var (
		request = &cityRequest.City{}
		err     error
	)

	if request.ID, err = strconv.Atoi(ctx.Params(cityID)); err != nil {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}

	ctx.Locals(requests.RequestKey, request)

	return ctx.Next()
}
