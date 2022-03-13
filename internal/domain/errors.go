package domain

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gocraft/dbr"
	"github.com/lib/pq"
	"net"
)

type ErrCode uint8

const (
	ErrCodeInternal        ErrCode = iota + 1 // 1
	ErrCodeNotFound                           // 2
	ErrCodeValidationErr                      // 3
	ErrCodeBadRequest                         // 4
	ErrCodeDatabaseFailure                    // 5
	ErrCodeInvalidParam                       // 6
	ErrCodeDatabaseError                      // 7
	ErrCodeNotAuthorized                      // 8
	ErrCodeForbidden                          // 9
	ErrCodeDuplicateKey                       // 10
)

const (
	DetailsUnknown = "Details unknown"
)

var (
	ErrInternal        = errors.New("internal error, report to developers")
	ErrNotFound        = errors.New("not found")
	ErrValidation      = errors.New("some fields failed the validation")
	ErrBadRequest      = errors.New("invalid request")
	ErrDatabaseFailure = errors.New("cannot connect to the database")
	ErrInvalidParam    = errors.New("invalid params")
	ErrDatabaseError   = errors.New("the database cannot process the query")
	ErrNotAuthorized   = errors.New("invalid login or password")
	ErrForbidden       = errors.New("action is forbidden")
	ErrDuplicateKey    = errors.New("entity already exists")
)

var errorMessages = map[ErrCode]error{
	ErrCodeInternal:        ErrInternal,
	ErrCodeNotFound:        ErrNotFound,
	ErrCodeValidationErr:   ErrValidation,
	ErrCodeBadRequest:      ErrBadRequest,
	ErrCodeDatabaseFailure: ErrDatabaseFailure,
	ErrCodeInvalidParam:    ErrInvalidParam,
	ErrCodeDatabaseError:   ErrDatabaseError,
	ErrCodeNotAuthorized:   ErrNotAuthorized,
	ErrCodeForbidden:       ErrForbidden,
	ErrCodeDuplicateKey:    ErrDuplicateKey,
}

type Error struct {
	UUID    string      `json:"uuid"`
	UserID  int64       `json:"user_id"`
	Code    ErrCode     `json:"code"`
	Details interface{} `json:"details"`
	text    string
}

type ValidationErrorDetails struct {
	ActualTag       string      `json:"actual_tag"`
	Error           string      `json:"error"`
	Field           string      `json:"field"`
	Kind            string      `json:"kind"`
	Namespace       string      `json:"namespace"`
	Param           string      `json:"param"`
	StructField     string      `json:"struct_field"`
	StructNamespace string      `json:"struct_namespace"`
	Tag             string      `json:"tag"`
	Type            string      `json:"type"`
	Value           interface{} `json:"value"`
}

func (e *Error) Error() string {
	return e.text
}

func NewValidationError(err error) *Error {
	validationErr, ok := err.(validator.ValidationErrors)

	if !ok {
		return NewErrorWithDetails(ErrCodeValidationErr, DetailsUnknown)
	}

	details := make([]ValidationErrorDetails, len(validationErr))

	for k, v := range validationErr {
		details[k] = ValidationErrorDetails{
			ActualTag:       v.ActualTag(),
			Error:           v.Error(),
			Field:           v.Field(),
			Kind:            v.Kind().String(),
			Namespace:       v.Namespace(),
			Param:           v.Param(),
			StructField:     v.StructField(),
			StructNamespace: v.StructNamespace(),
			Tag:             v.Tag(),
			Type:            v.Type().String(),
			Value:           v.Value(),
		}
	}

	return NewErrorWithDetails(ErrCodeValidationErr, details)
}

// NewErrorWrap wrap an external error by custom.
func NewErrorWrap(err error, code ErrCode, params ...interface{}) *Error {
	return &Error{
		Details: err.Error(),
		text:    fmt.Sprintf("%s %v", errorMessages[code], params),
		Code:    code,
	}
}

func NewErrorWithDetails(code ErrCode, details interface{}) *Error {
	return &Error{
		Details: details,
		Code:    code,
		text:    errorMessages[code].Error(),
	}
}

// NewError creates a new custom error of our format.
func NewError2(code ErrCode) *Error {
	return NewErrorWithDetails(code, nil)
}

// NewError creates a new custom error of our format.
func NewError(code ErrCode, params ...interface{}) *Error {
	err := errorMessages[code]

	return &Error{
		Details: err.Error(),
		text:    fmt.Sprintf("%s %v", err, params),
		Code:    code,
	}
}

func NewDBErrorWrap(err error) error {
	switch err.(type) { // nolint:errorlint
	case *net.OpError:
		return NewErrorWithDetails(ErrCodeDatabaseFailure, err)
	case *pq.Error:
		return NewErrorWithDetails(ErrCodeDatabaseError, err)
	}

	if errors.Is(err, dbr.ErrNotFound) {
		return NewError2(ErrCodeNotFound)
	}

	return NewError2(ErrCodeInternal)
}

func (e Error) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"error_data": e,
	}
}
