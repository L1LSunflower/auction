package errorhandler

import "github.com/gofiber/fiber/v2"

type TypeError struct {
	StatusCode int
	HttpCode   int
}

var TypeErrors = map[string]*TypeError{
	// Error for users
	ErrBearerAuth.Error(): {
		StatusCode: 110,
		HttpCode:   fiber.StatusForbidden,
	},
	ErrStoreUser.Error(): {
		StatusCode: 111,
		HttpCode:   fiber.StatusUnprocessableEntity,
	},
	ErrNeedConfirm.Error(): {
		StatusCode: 112,
		HttpCode:   fiber.StatusBadRequest,
	},
	ErrFindByPhone.Error(): {
		StatusCode: 113,
		HttpCode:   fiber.StatusNotFound,
	},
	ErrUserExist.Error(): {
		StatusCode: 114,
		HttpCode:   fiber.StatusUnprocessableEntity,
	},
	ErrCreateUser.Error(): {
		StatusCode: 115,
		HttpCode:   fiber.StatusUnprocessableEntity,
	},
	ErrSendOtp.Error(): {
		StatusCode: 116,
		HttpCode:   fiber.StatusUnprocessableEntity,
	},
	ErrStoreOtp.Error(): {
		StatusCode: 117,
		HttpCode:   fiber.StatusUnprocessableEntity,
	},
	ErrCodeExpired.Error(): {
		StatusCode: 118,
		HttpCode:   fiber.StatusNotFound,
	},
	WrongCode.Error(): {
		StatusCode: 119,
		HttpCode:   fiber.StatusUnprocessableEntity,
	},
	ErrConfirm.Error(): {
		StatusCode: 120,
		HttpCode:   fiber.StatusUnprocessableEntity,
	},
	ErrUserExpired.Error(): {
		StatusCode: 121,
		HttpCode:   fiber.StatusNotFound,
	},
	ErrStoreToken.Error(): {
		StatusCode: 122,
		HttpCode:   fiber.StatusUnprocessableEntity,
	},
	WrongPassword.Error(): {
		StatusCode: 123,
		HttpCode:   fiber.StatusBadRequest,
	},
	WrongTokens.Error(): {
		StatusCode: 124,
		HttpCode:   fiber.StatusForbidden,
	},
	AuthRequired.Error(): {
		StatusCode: 125,
		HttpCode:   fiber.StatusForbidden,
	},
	ErrGetTokens.Error(): {
		StatusCode: 126,
		HttpCode:   fiber.StatusForbidden,
	},
	ErrUserNotExist.Error(): {
		StatusCode: 127,
		HttpCode:   fiber.StatusNotFound,
	},
	// Validation Error
	ErrParseRequest.Error(): {
		StatusCode: 997,
		HttpCode:   fiber.StatusBadRequest,
	},
	// Internal Errors
	ErrDependency.Error(): {
		StatusCode: 998,
		HttpCode:   fiber.StatusFailedDependency,
	},
	// Unknown Error
	UnknownError.Error(): {
		StatusCode: 999,
		HttpCode:   fiber.StatusInternalServerError,
	},
}
