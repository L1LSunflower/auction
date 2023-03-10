package tags

import (
	"context"
	"github.com/gofiber/fiber/v2"

	"github.com/L1LSunflower/auction/config"
	tagsService "github.com/L1LSunflower/auction/internal/domain/services/tags"
	"github.com/L1LSunflower/auction/internal/requests"
	tagsRequest "github.com/L1LSunflower/auction/internal/requests/structs/tags"
	"github.com/L1LSunflower/auction/internal/responses"
	"github.com/L1LSunflower/auction/internal/tools/context_with_depends"
	"github.com/L1LSunflower/auction/internal/tools/errorhandler"
	"github.com/L1LSunflower/auction/pkg/db"
	"github.com/L1LSunflower/auction/pkg/redisdb"
)

func ByPattern(ctx *fiber.Ctx) error {
	request, ok := ctx.Locals(requests.RequestKey).(*tagsRequest.Tag)
	if !ok {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}

	dbConn := db.SqlInstance(config.GetConfig().DB.DBDriver, config.GetConfig().DB.DBString).DB
	redisConn := redisdb.RedisInstance().RedisClient
	contxt, err := context_with_depends.ContextWithDepends(context.Background(), dbConn, redisConn)
	if err != nil {
		return responses.NewFailedResponse(ctx, errorhandler.ErrDependency)
	}

	tags, err := tagsService.ByPattern(contxt, request)
	if err != nil {
		return responses.NewFailedResponse(ctx, err)
	}

	return responses.Tags(ctx, tags)
}
