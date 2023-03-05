/*
Package middlewares that package contains attempts (requests|responses),
jwt authentication, jwt validation and write attempts logs.
*/
package middlewares

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofrs/uuid"
	"github.com/valyala/fasthttp"

	"github.com/L1LSunflower/auction/pkg/logger"
	"github.com/L1LSunflower/auction/pkg/logger/message"
)

const (
	// AttemptBefore message for log.
	AttemptBefore = "ATTEMPT:BEFORE"
	// AttemptAfter message for log.
	AttemptAfter = "ATTEMPT:AFTER"
)

/*
Attempts attempts for log (requests|responses), returns fiber.Handler.
*/
func Attempts() fiber.Handler {
	return func(ctx *fiber.Ctx) (err error) {
		request := parseRequest(ctx)

		uid, err := uuid.NewV4()
		if err != nil {
			logger.Log.Error(&message.LogMessage{
				Message: fmt.Sprintf("failed to generat uuid v4 for attempts with error_response: %s", err.Error()),
			})
		}

		requestId := uid.String()

		beforeAttempt(ctx, request, ctx.Request().Header.String(), requestId)

		err = ctx.Next()

		res := parseResponse(ctx)

		afterAttempt(ctx, res, ctx.Response().Header.String(), requestId)

		return err
	}
}

/*
beforeAttempt that write log about requests.
*/
func beforeAttempt(ctx *fiber.Ctx, request, headers, requestId string) {
	extra := map[string]any{
		"endpoint":       ctx.Request().URI().String(),
		"request_id":     requestId,
		"request_method": string(ctx.Request().Header.Method()),
		"headers":        headers,
	}

	msg := &message.LogMessage{
		Message:     AttemptBefore,
		FullMessage: &request,
		Extra:       &extra,
	}

	logger.Log.Info(msg)
}

/*
afterAttempt that write log about responses.
*/
func afterAttempt(ctx *fiber.Ctx, response, headers, requestId string) {
	extra := map[string]any{
		"endpoint":       ctx.Request().URI().String(),
		"request_id":     requestId,
		"request_method": string(ctx.Request().Header.Method()),
		"headers":        headers,
	}

	msg := &message.LogMessage{
		Message:     AttemptAfter,
		FullMessage: &response,
		Extra:       &extra,
	}

	logger.Log.Info(msg)
}

/*
cleanHeaders that cleans header, returns string.
*/
func cleanHeaders(headers *fasthttp.RequestHeader) string {
	var result []string

	headers.VisitAllInOrder(func(key, value []byte) {
		k := string(key)
		v := string(value)

		if k == "Authorization" && len(v) > 10 {
			v = fmt.Sprintf("%s-%s", v[:15], strings.Repeat("*", 10))
		}

		result = append(result, fmt.Sprintf("%s: %s", k, v))
	})

	return strings.Join(result, "\n")
}

/*
parseRequest parsing requests from bytes and turn into string, returns string.
*/
func parseRequest(ctx *fiber.Ctx) (req string) {
	if string(ctx.Request().Header.Method()) == fiber.MethodPost {
		req = string(ctx.Request().Body())
	} else {
		req = ctx.Request().URI().QueryArgs().String()
	}

	return
}

/*
parseResponse parsing responses, returns string.
*/
func parseResponse(ctx *fiber.Ctx) (res string) {
	res = string(ctx.Response().Body())

	return
}
