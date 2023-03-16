package requests

import (
	"github.com/L1LSunflower/auction/internal/responses"
	"github.com/L1LSunflower/auction/internal/tools/errorhandler"
	"github.com/gofiber/fiber/v2"
)

const (
	RequestKey = "request"
	UserIDCtx  = "user_id"
)

func ParseRequest(ctx *fiber.Ctx, request any) error {
	if err := ctx.BodyParser(request); err != nil {
		return responses.NewFailedResponse(ctx, errorhandler.ErrParseRequest)
	}
	return nil
}
