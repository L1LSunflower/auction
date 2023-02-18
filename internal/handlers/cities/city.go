package cities

import (
	"context"

	"github.com/gofiber/fiber/v2"

	"github.com/L1LSunflower/auction/config"
	cityService "github.com/L1LSunflower/auction/internal/domain/services/cities"
	"github.com/L1LSunflower/auction/internal/requests"
	cityRequest "github.com/L1LSunflower/auction/internal/requests/structs/cities"
	"github.com/L1LSunflower/auction/internal/responses"
	"github.com/L1LSunflower/auction/internal/tools/errorhandler"
	"github.com/L1LSunflower/auction/pkg/db"
)

func Create(ctx *fiber.Ctx) error {
	request, ok := ctx.Locals(requests.RequestKey).(cityRequest.CreateCity)
	if !ok {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}

	dbConn := db.SqlInstance(config.GetConfig().DB.DBDriver, config.GetConfig().DB.DBString).DB
	city, err := cityService.Create(context.Background(), dbConn, request)
	if err != nil {
		return responses.NewFailedResponse(ctx, err)
	}

	return responses.NewSuccessResponse(ctx, city)
}

func Cities(ctx *fiber.Ctx) error {
	dbConn := db.SqlInstance(config.GetConfig().DB.DBDriver, config.GetConfig().DB.DBString).DB
	cities, err := cityService.Cities(context.Background(), dbConn)
	if err != nil {
		return responses.NewFailedResponse(ctx, err)
	}

	return responses.NewSuccessResponse(ctx, cities)
}

func City(ctx *fiber.Ctx) error {
	request, ok := ctx.Locals(requests.RequestKey).(cityRequest.City)
	if !ok {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}

	dbConn := db.SqlInstance(config.GetConfig().DB.DBDriver, config.GetConfig().DB.DBString).DB
	city, err := cityService.City(context.Background(), dbConn, request)
	if err != nil {
		return responses.NewFailedResponse(ctx, err)
	}

	return responses.NewSuccessResponse(ctx, city)
}

func Update(ctx *fiber.Ctx) error {
	request, ok := ctx.Locals(requests.RequestKey).(cityRequest.CityUpdate)
	if !ok {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}

	dbConn := db.SqlInstance(config.GetConfig().DB.DBDriver, config.GetConfig().DB.DBString).DB
	if err := cityService.Update(context.Background(), dbConn, request); err != nil {
		return responses.NewFailedResponse(ctx, err)
	}

	return responses.NewSuccessResponse(ctx, "success")
}

func Delete(ctx *fiber.Ctx) error {
	request, ok := ctx.Locals(requests.RequestKey).(cityRequest.City)
	if !ok {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}

	dbConn := db.SqlInstance(config.GetConfig().DB.DBDriver, config.GetConfig().DB.DBString).DB
	if err := cityService.Delete(context.Background(), dbConn, request); err != nil {
		return responses.NewFailedResponse(ctx, err)
	}

	return responses.NewSuccessResponse(ctx, "success")
}
