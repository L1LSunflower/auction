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
	// Auction Errors
	ErrUpdateActiveAuction.Error(): {
		StatusCode: 129,
		HttpCode:   fiber.StatusUnavailableForLegalReasons,
	},
	// Balance Errors
	ErrCreateBalance.Error(): {
		StatusCode: 128,
		HttpCode:   fiber.StatusUnprocessableEntity,
	},
	ErrProcessCard.Error(): {
		StatusCode: 201,
		HttpCode:   fiber.StatusUnprocessableEntity,
	},
	ErrCreditBalance.Error(): {
		StatusCode: 202,
		HttpCode:   fiber.StatusUnprocessableEntity,
	},
	ErrDebitBalance.Error(): {
		StatusCode: 203,
		HttpCode:   fiber.StatusUnprocessableEntity,
	},
	ErrGetBalance.Error(): {
		StatusCode: 204,
		HttpCode:   fiber.StatusUnprocessableEntity,
	},
	// Validation Errors
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
	ErrActiveAuctionExist.Error(): {
		StatusCode: 128,
		HttpCode:   fiber.StatusUnprocessableEntity,
	},
	ErrCreateItem.Error(): {
		StatusCode: 129,
		HttpCode:   fiber.StatusUnprocessableEntity,
	},
	ErrCreateTag.Error(): {
		StatusCode: 130,
		HttpCode:   fiber.StatusUnprocessableEntity,
	},
	ErrCreateFile.Error(): {
		StatusCode: 131,
		HttpCode:   fiber.StatusUnprocessableEntity,
	},
	ErrCreateAuction.Error(): {
		StatusCode: 132,
		HttpCode:   fiber.StatusUnprocessableEntity,
	},
	ErrCreateLimit.Error(): {
		StatusCode: 133,
		HttpCode:   fiber.StatusUnprocessableEntity,
	},
	ErrDoesNotExistAuction.Error(): {
		StatusCode: 134,
		HttpCode:   fiber.StatusNotFound,
	},
	ErrDoesNotExistItem.Error(): {
		StatusCode: 135,
		HttpCode:   fiber.StatusNotFound,
	},
	ErrGetFiles.Error(): {
		StatusCode: 136,
		HttpCode:   fiber.StatusNotFound,
	},
	ErrGetTags.Error(): {
		StatusCode: 137,
		HttpCode:   fiber.StatusNotFound,
	},
	ErrGetAuctions.Error(): {
		StatusCode: 138,
		HttpCode:   fiber.StatusNotFound,
	},
	ErrUpdateItem.Error(): {
		StatusCode: 139,
		HttpCode:   fiber.StatusUnprocessableEntity,
	},
	ErrDeleteTags.Error(): {
		StatusCode: 140,
		HttpCode:   fiber.StatusInternalServerError,
	},
	ErrDeleteFile.Error(): {
		StatusCode: 141,
		HttpCode:   fiber.StatusInternalServerError,
	},
	ErrFailedStartAuction.Error(): {
		StatusCode: 142,
		HttpCode:   fiber.StatusUnprocessableEntity,
	},
	ErrDeleteByStatus.Error(): {
		StatusCode: 143,
		HttpCode:   fiber.StatusUnprocessableEntity,
	},
	ErrDeleteAuction.Error(): {
		StatusCode: 144,
		HttpCode:   fiber.StatusInternalServerError,
	},
	ErrDeleteItem.Error(): {
		StatusCode: 145,
		HttpCode:   fiber.StatusInternalServerError,
	},
	ErrDeleteFiles.Error(): {
		StatusCode: 146,
		HttpCode:   fiber.StatusInternalServerError,
	},
}
