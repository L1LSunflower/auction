package middlewares

import (
	"github.com/L1LSunflower/auction/internal/responses"
	"strings"

	"github.com/gofiber/fiber/v2"

	"github.com/L1LSunflower/auction/config"
	"github.com/L1LSunflower/auction/internal/tools/errorhandler"
)

func BearerAuth() fiber.Handler {
	bearerToken := config.GetConfig().BearerToken
	return func(ctx *fiber.Ctx) (err error) {
		hv := ctx.Get(fiber.HeaderAuthorization)
		components := strings.SplitN(hv, " ", 2)

		if len(components) < 2 || bearerToken != components[1] {
			ctx.Status(fiber.StatusUnauthorized)
			return responses.NewFailedResponse(ctx, errorhandler.ErrBearerAuth)
		}

		return ctx.Next()
	}
}
