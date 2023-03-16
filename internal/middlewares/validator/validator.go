package validator

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
)

const (
	errSeparator = ", "
)

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func (e *ErrorResponse) Err() string {
	return fmt.Sprintf("%s: is %s, got value: '%s'", e.FailedField, e.Tag, e.Value)
}

func ValidateRequest(request any) error {
	var (
		validate = validator.New()
		errs     []*ErrorResponse
		err      error
	)

	if err = validate.Struct(request); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errs = append(errs, &element)
		}
	}

	if len(errs) > 0 {
		listErrs := MakeFlatStringError(errs)
		return errors.New(strings.Join(listErrs, errSeparator))
	}

	return nil
}

func MakeFlatStringError(errors []*ErrorResponse) []string {
	var flatStringErrs []string
	for _, err := range errors {
		flatStringErrs = append(flatStringErrs, err.Err())
	}
	return flatStringErrs
}
