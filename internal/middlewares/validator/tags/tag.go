package tags

import (
	"github.com/L1LSunflower/auction/internal/middlewares/validator"
	"github.com/L1LSunflower/auction/internal/requests"
	tagRequest "github.com/L1LSunflower/auction/internal/requests/structs/tags"
	"github.com/L1LSunflower/auction/internal/responses"
	"github.com/L1LSunflower/auction/internal/tools/errorhandler"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

func ByPattern(ctx *fiber.Ctx) error {
	var err error
	request := &tagRequest.Tag{}
	args := ctx.Request().URI().QueryArgs()
	request.Pattern, err = parseParam(args, "pattern")

	if err != nil {
		return responses.NewFailedResponse(ctx, err)
	}

	if err = validator.ValidateRequest(request); err != nil {
		return responses.NewValidationErrResponse(ctx, err)
	}

	ctx.Locals(requests.RequestKey, request)

	return ctx.Next()
}

func parseParam(args *fasthttp.Args, nameParams string) (string, error) {
	if v := string(args.Peek(nameParams)); v != "" {
		return v, nil
	}
	return "", errorhandler.ErrFailedToGetTags
}
