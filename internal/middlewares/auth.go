package middlewares

import (
	"context"
	"github.com/L1LSunflower/auction/internal/requests"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"

	"github.com/L1LSunflower/auction/internal/domain/repositories/redis_repository"
	"github.com/L1LSunflower/auction/internal/responses"
	"github.com/L1LSunflower/auction/internal/tools/context_with_depends"
	"github.com/L1LSunflower/auction/internal/tools/errorhandler"
	"github.com/L1LSunflower/auction/pkg/redisdb"
)

func Auth() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Get("id")
		at := ctx.Get("access")

		redisConn := redisdb.RedisInstance().RedisClient
		contxt, err := context_with_depends.ContextWithDepends(context.Background(), nil, redisConn)
		if err != nil {
			return responses.NewFailedResponse(ctx, errorhandler.AuthRequired)
		}

		tokens, err := redis_repository.UserInterface.Tokens(contxt, id)
		if err != nil && err != redis.Nil {
			return responses.NewFailedResponse(ctx, errorhandler.AuthRequired)
		}

		if tokens == nil {
			return responses.NewFailedResponse(ctx, errorhandler.AuthRequired)
		}

		if tokens.AccessToken != at {
			return responses.NewSuccessResponse(ctx, errorhandler.AuthRequired)
		}

		ctx.Locals(requests.UserIDCtx, id)
		return ctx.Next()

	}
}
