package errorhandler

import "errors"

var (
	// Error for Authentication
	ErrBearerAuth = errors.New("authorization required")
	// Error for requests
	ErrParseRequest = errors.New("failed to parse request")
	// Error for cities
	ErrCreateCity = errors.New("failed to create city")
	ErrDeleteCity = errors.New("faield to delete city")
)
