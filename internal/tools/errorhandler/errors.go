package errorhandler

import "errors"

var (
	ErrParseRequestCode = 997
)

var (
	// Error for Users
	ErrBearerAuth     = errors.New("authorization required")
	ErrStoreUser      = errors.New("failed to store user")
	ErrNeedConfirm    = errors.New("need account confirmation")
	ErrFindByPhone    = errors.New("failed to find user by phone")
	ErrUserExist      = errors.New("user with that phone already exist")
	ErrCreateUser     = errors.New("error create user")
	ErrSendOtp        = errors.New("failed to send otp code")
	ErrStoreOtp       = errors.New("failed to store otp code")
	ErrCodeExpired    = errors.New("user code expired")
	WrongCode         = errors.New("wrong code")
	ErrConfirm        = errors.New("failed to confirm otp code")
	ErrUserExpired    = errors.New("that user does not exist")
	ErrStoreToken     = errors.New("failed to create token")
	WrongPassword     = errors.New("wrong password")
	WrongTokens       = errors.New("wrong tokens")
	AuthRequired      = errors.New("authorization required")
	ErrGetTokens      = errors.New("failed to get tokens")
	ErrRestore        = errors.New("failed to restore password")
	ErrUserNotExist   = errors.New("user does not exist")
	ErrUpdatePassword = errors.New("failed to update user password")
	ErrUpdateUser     = errors.New("failed to update user")
	ErrDeleteUser     = errors.New("failed to delete user")
	// Error for Auctions
	ErrActiveAuctionExist  = errors.New("active auction already exist")
	ErrCreateItem          = errors.New("failed to create item")
	ErrCreateTag           = errors.New("failed to create tag for item")
	ErrCreateFile          = errors.New("failed to create file for item")
	ErrCreateAuction       = errors.New("failed to create auction")
	ErrCreateLimit         = errors.New("auction create limit")
	ErrDoesNotExistAuction = errors.New("auction does not exist")
	ErrDoesNotExistItem    = errors.New("item does not exist")
	ErrGetFiles            = errors.New("failed get item files")
	ErrGetTags             = errors.New("failed get item tags")
	ErrGetAuctions         = errors.New("failed to get auctions")
	ErrUpdateItem          = errors.New("failed to update auction item")
	ErrDeleteTags          = errors.New("failed to delete item tags")
	ErrDeleteFile          = errors.New("failed to delete item file")
	ErrFailedStartAuction  = errors.New("failed to start auction")
	ErrDeleteByStatus      = errors.New("failed to delete auction with that status")
	ErrDeleteAuction       = errors.New("failed to delete auction")
	ErrDeleteItem          = errors.New("failed to delete item")
	ErrDeleteFiles         = errors.New("failed to delete item files")
	ErrUpdateActiveAuction = errors.New("can't update auction while active")
	ErrAuctionNotExist     = errors.New("failed to get auction")
	ErrDeleteAuctions      = errors.New("failed to delete user auctions")
	ErrFailedToGetTags     = errors.New("failed to get tags from url params")
	ErrEndAuction          = errors.New("failed to end auction")
	ErrParticipate         = errors.New("failed to add participate user for auction")
	// Error for Balances
	ErrCreateBalance     = errors.New("failed to create user balance")
	NotEnoughBalance     = errors.New("not enough balance")
	ErrProcessCard       = errors.New("failed to process card")
	ErrCreditBalance     = errors.New("failed to credit balance")
	ErrDebitBalance      = errors.New("failed to debit amount from balance")
	ErrGetBalance        = errors.New("failed to get users balance")
	ErrCreateTransaction = errors.New("failed to create transaction")
	// Error for metadata
	ErrGettingPerPage    = errors.New("failed to get per page param")
	ErrGettingPage       = errors.New("failed to get page param")
	ErrWrongPerPageValue = errors.New("wrong per page value")
	// Error for requests
	ErrParseRequest = errors.New("failed to parse request")
	// Internal errors
	ErrDependency = errors.New("error dependency")
	// Unknown error
	InternalError = errors.New("internal error")
	UnknownError  = errors.New("unknown error")
)
