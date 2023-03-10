package tags

import (
	"github.com/gofiber/fiber/v2"

	"github.com/L1LSunflower/auction/internal/requests"
	tagRequest "github.com/L1LSunflower/auction/internal/requests/structs/tags"
	"github.com/L1LSunflower/auction/internal/responses"
	"github.com/L1LSunflower/auction/internal/tools/errorhandler"
)

func ByPattern(ctx *fiber.Ctx) error {
	request := &tagRequest.Tag{}
	if err := ctx.BodyParser(request); err != nil {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}

	ctx.Locals(requests.RequestKey, request)

	return ctx.Next()
}
