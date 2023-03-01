package errorhandler

import "errors"

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
	// Error for requests
	ErrParseRequest = errors.New("failed to parse request")
	// Internal errors
	ErrDependency = errors.New("error dependency")
	// Unknown error
	UnknownError = errors.New("unknown error")
)
