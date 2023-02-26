package middlewares

import (
	"context"
	"github.com/L1LSunflower/auction/internal/domain/repositories/redis_repository"
	"github.com/L1LSunflower/auction/internal/tools/context_with_depends"
	"github.com/L1LSunflower/auction/pkg/redisdb"
	"github.com/gofiber/fiber/v2"
)

func Auth() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		hv := ctx.Get("access")

		redisConn := redisdb.RedisInstance().RedisClient
		contxt, err := context_with_depends.ContextWithDepends(context.Background(), nil, redisConn)
		if err != nil {
			ctx.Status(fiber.StatusForbidden)
			return ctx.JSON(fiber.Map{
				"status":  "error",
				"message": "authorization required",
			})
		}

		if _, err = redis_repository.UserInterface.TokenByKey(contxt, hv+":access"); err != nil {
			ctx.Status(fiber.StatusForbidden)
			return ctx.JSON(fiber.Map{
				"status":  "error",
				"message": "authorization required",
			})
		}

		return ctx.Next()

	}
}
