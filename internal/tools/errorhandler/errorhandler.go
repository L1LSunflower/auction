package errorhandler

import "github.com/gofiber/fiber/v2"

type TypeError struct {
	StatusCode int
}

var TypeErrors = map[string]*TypeError{
	ErrBearerAuth.Error(): {
		StatusCode: fiber.StatusForbidden,
	},
	ErrParseRequest.Error(): {
		StatusCode: fiber.StatusBadRequest,
	},
}
